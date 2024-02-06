package engine

import (
	"bytes"
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
	"sigs.k8s.io/yaml"

	"go.indent.com/indent-go/pkg/auditv1"
	multipartyv1 "multiparty.ai/api/multiparty/v1"
)

var (
	marshaller = jsonpb.Marshaler{}
)

func Eval(in []byte) ([]byte, error) {
	jsonIn, err := yaml.YAMLToJSONStrict(in)
	if err != nil {
		return nil, fmt.Errorf("failed to convert yaml to json: %w", err)
	}

	req := new(multipartyv1.EvalRequest)
	var out bytes.Buffer
	if err := jsonpb.Unmarshal(bytes.NewReader(jsonIn), req); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	} else if res, err := EvalPolicy(req); err != nil {
		return nil, fmt.Errorf("failed to eval policy: %w", err)
	} else if err = marshaller.Marshal(&out, res); err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	yamlOut, err := yaml.JSONToYAML(out.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to convert json to yaml: %w", err)
	}
	return yamlOut, nil
}

func EvalPolicy(req *multipartyv1.EvalRequest) (*multipartyv1.EvalResponse, error) {
	cfg := req.GetConfig()
	results := make([]*multipartyv1.Result, len(req.GetTests()))
	for i, rule := range req.GetTests() {
		var err error
		if results[i], err = eval(cfg, rule); err != nil {
			return nil, err
		}
	}

	return &multipartyv1.EvalResponse{
		Results: results,
	}, nil
}

func eval(cfg *multipartyv1.Config, test *multipartyv1.Test) (*multipartyv1.Result, error) {
	result := new(multipartyv1.Result)
	for i, rule := range cfg.GetRules() {
		actor, resource, err := loadPetitionFromIn(test.GetIn())
		if err != nil {
			return nil, fmt.Errorf("failed to load petition: %w", err)
		}

		if prg, err := setupCEL(rule.GetIf()); err != nil {
			return nil, fmt.Errorf("rule %d (%q) failed to compile: %w", i, rule.GetIf(), err)
		} else if out, _, err := prg.Eval(map[string]interface{}{
			"actor":    actor,
			"resource": resource,
		}); err != nil {
			return nil, fmt.Errorf("rule %d (%q) failed to evaluate: %w", i, rule.GetIf(), err)
		} else if pass, ok := out.Value().(bool); !ok {
			return nil, fmt.Errorf("rule %d (%q) returned non-boolean: %v", i, rule.GetIf(), out.Value())
		} else if !pass {
			continue
		}
		outcome := rule.GetOutcome()
		result.Rules = append(result.Rules, rule)
		result.Events = append(result.Events, outcome)

		for i, check := range test.GetChecks() {
			if env, err := setupTestCEL(check.GetIf()); err != nil {
				return nil, fmt.Errorf("check %d (%q) failed to compile: %w", i, check.GetIf(), err)
			} else if out, _, err := env.Eval(map[string]interface{}{
				"out": celVar(outcome),
			}); err != nil {
				return nil, fmt.Errorf("check %d (%q) failed to evaluate: %w", i, check.GetIf(), err)
			} else if pass, ok := out.Value().(bool); !ok {
				return nil, fmt.Errorf("rule %d (%q) returned non-boolean: %v", i, rule.GetIf(), out.Value())
			} else if !pass {
				result.Failures = append(result.Failures, &multipartyv1.Failure{
					Check: check,
					Error: "Did not pass",
				})
			}
		}
	}

	if len(result.GetEvents()) == 0 && len(test.GetChecks()) > 0 {
		result.Failures = append(result.Failures, &multipartyv1.Failure{
			Error: "No matching rules",
		})
	}
	return result, nil
}

func setupCEL(txt string) (cel.Program, error) {
	pReflect := proto.MessageReflect(new(auditv1.Resource))
	resourceType := decls.NewObjectType(string(pReflect.Descriptor().FullName()))
	env, err := cel.NewEnv(
		cel.Types(pReflect.Interface()),
		cel.Declarations(
			decls.NewVar("resource", resourceType),
			decls.NewVar("actor", resourceType),
		),
		cel.Macros(cel.StandardMacros...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create cel environment: %w", err)
	}
	return compileCEL(env, txt)
}

func setupTestCEL(txt string) (cel.Program, error) {
	pReflect := proto.MessageReflect(new(multipartyv1.Outcome))
	outType := decls.NewObjectType(string(pReflect.Descriptor().FullName()))
	env, err := cel.NewEnv(
		cel.Types(pReflect.Interface()),
		cel.Declarations(
			decls.NewVar("out", outType),
		),
		cel.Macros(cel.StandardMacros...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create cel environment: %w", err)
	}
	return compileCEL(env, txt)
}

func compileCEL(env *cel.Env, txt string) (cel.Program, error) {
	if ast, issues := env.Compile(txt); issues != nil && issues.Err() != nil {
		return nil, fmt.Errorf("failed to compile: %w", issues.Err())
	} else if prg, err := env.Program(ast); err != nil {
		return nil, fmt.Errorf("failed to setup program: %w", err)
	} else {
		return prg, nil
	}
}

func loadPetitionFromIn(in *structpb.Value) (actor, resource protoreflect.ProtoMessage, err error) {
	strct := in.GetStructValue()
	if strct == nil {
		return nil, nil, fmt.Errorf("input must be a struct")
	}

	if actor, err = resourceFromValue(strct.Fields["actor"]); err != nil {
		return nil, nil, fmt.Errorf("failed to load actor: %w", err)
	} else if resource, err = resourceFromValue(strct.Fields["resource"]); err != nil {
		return nil, nil, fmt.Errorf("failed to load resource: %w", err)
	}
	return actor, resource, nil
}

func resourceFromValue(val *structpb.Value) (protoreflect.ProtoMessage, error) {
	if val == nil {
		return nil, nil
	}

	var buf bytes.Buffer
	if err := marshaller.Marshal(&buf, val); err != nil {
		return nil, fmt.Errorf("failed to marshal value: %w", err)
	}

	resource := new(auditv1.Resource)
	if err := jsonpb.Unmarshal(&buf, resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}
	return celVar(resource), nil
}

func celVar(msg proto.Message) protoreflect.ProtoMessage {
	return proto.MessageReflect(msg).Interface()
}

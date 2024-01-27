package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"sigs.k8s.io/yaml"

	multipartyv1 "multiparty.ai/api/multiparty/v1"
)

func TestPolicy(t *testing.T) {
	files, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		t.Run(file.Name(), func(t *testing.T) {
			t.Parallel()

			resp := new(multipartyv1.EvalResponse)
			data, err := os.ReadFile("testdata/" + file.Name())
			if err != nil {
				t.Fatal(err)
			} else if out, err := Eval(data); err != nil {
				t.Fatal(err)
			} else if jsonData, err := yaml.YAMLToJSON(out); err != nil {
				t.Fatal(err)
			} else if err = jsonpb.Unmarshal(bytes.NewReader(jsonData), resp); err != nil {
				t.Fatal(err)
			}

			for _, result := range resp.GetResults() {
				if len(result.GetFailures()) > 0 {
					t.Error("There were failures")
					for _, failure := range result.GetFailures() {
						t.Log(failure)
					}
				}
			}
		})
	}
}

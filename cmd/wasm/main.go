//go:build js && wasm

package main

import (
	"errors"
	"syscall/js"

	"multiparty.ai/pkg/engine"
)

func main() {
	evalFunc := js.FuncOf(evalWrapper)
	// lintFunc := js.FuncOf(lintWrapper)
	js.Global().Set("mp", evalFunc)
	// js.Global().Set("clint", lintFunc)
	defer evalFunc.Release()
	<-make(chan bool)
}

// lintWrapper wraps the validate function with `syscall/js` parameters
// func lintWrapper(_ js.Value, args []js.Value) any {
// 	if len(args) < 1 {
// 		return response("", errors.New("missing CEL expression"))
// 	}
// 	exp := args[0].String()
// 	output, err := eval.Lint(exp)
// 	if err != nil {
// 		return response(false, err)
// 	}
// 	return response(output, nil)
// }

func log(msg string) {
	js.Global().Get("console").Call("log", msg)
}

// evalWrapper wraps the eval function with `syscall/js` parameters
func evalWrapper(_ js.Value, args []js.Value) (res any) {
	if len(args) < 1 {
		return response("", errors.New("invalid arguments"))
	}
	defer func() {
		// recover from panic if one occurred. Set err to nil otherwise.
		if recover() != nil {
			res = map[string]any{"output": "panic error", "error": true}
		}
	}()

	input := args[0].String()
	output, err := engine.Eval([]byte(input))
	if err != nil {
		return response("", err)
	}
	return response(string(output), nil)
}

func response(out any, err error) any {
	if err != nil {
		out = err.Error()
	}
	return map[string]any{"output": out, "error": err != nil}
}

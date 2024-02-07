//go:build js && wasm

package main

import (
	"errors"
	"syscall/js"

	"multiparty.ai/pkg/engine"
)

func main() {
	evalFunc := js.FuncOf(evalWrapper)
	js.Global().Set("MP", evalFunc)
	// js.Global().Set("@indent/engine", js.ValueOf(make(map[string]interface{})))
	// module := js.Global().Get("@indent/engine")
	// module.Set("eval", evalFunc)
	<-make(chan struct{})
}

// evalWrapper wraps the eval function with `syscall/js` parameters
func evalWrapper(this js.Value, args []js.Value) (res any) {
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

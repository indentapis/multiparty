package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"multiparty.ai/pkg/engine"
)

func main() {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	input := string(stdin)
	output, err := engine.Eval([]byte(input))
	if err != nil {
		response("", err)
		return
	}
	response(string(output), nil)
}

func response(out string, err error) {
	if err != nil {
		out = err.Error()
	}
	r := map[string]interface{}{"output": out, "error": err != nil}
	jsonData, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.TrimSuffix(string(jsonData), "\n"))
}

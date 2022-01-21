package main

import (
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

func prepareExpr(code string) *vm.Program {
	program, err := expr.Compile(code)
	if err != nil {
		panic(err)
	}

	_, err = expr.Run(program, nil)
	if err != nil {
		panic(err)
	}
	// fmt.Println(output)

	return program
}

func prepareTengo(code string, vars map[string]interface{}) *tengo.Compiled {
	scr := tengo.NewScript([]byte(code))
	scr.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))

	for name, value := range vars {
		scr.Add(name, value)
	}

	program, err := scr.Run()
	if err != nil {
		panic(err)
	}

	return program
}

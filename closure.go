package main

import (
	"github.com/d5/tengo/v2"
	"github.com/k0kubun/pp"
)

func closure() {
	code := `
fmt := import("fmt")
v := 888
f := func(a, b) {
    // fmt.println("Here it is my closure: ", v)
    return a + b
}
fmt.println("Running my closure:")
fmt.println(f(1, 2))
`
	symbolTable := tengo.NewSymbolTable()
	globals := make([]tengo.Object, tengo.GlobalsSize)

	c := buildCompiler(code, symbolTable, nil)
	vm := tengo.NewVM(c.Bytecode(), globals, -1)
	if err := vm.Run(); err != nil {
		panic(err)
	}

	// f, _, _ := symbolTable.Resolve("f", false)
	// pp.Println(f)
	pp.Println(symbolTable)
	pp.Println(globals)

	code = `
//fmt := import("fmt")
fmt.println(f)
fmt.println(is_callable(f))
fmt.println(is_undefined(f))
fmt.println("Running my closure:")
fmt.println(f(3, 4))
`
	c = buildCompiler(code, symbolTable, nil)
	vm = tengo.NewVM(c.Bytecode(), globals, -1)
	if err := vm.Run(); err != nil {
		panic(err)
	}
}

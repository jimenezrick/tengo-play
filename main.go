package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/parser"
	"github.com/d5/tengo/v2/stdlib"
	"github.com/k0kubun/pp"
)

type Bar struct {
	Xxx string
	Yyy bool
}

type Foo struct {
	A int
	B string
	C *Bar
}

var l = []interface{}{1, 2, 3, 4}
var b = true
var o = map[string]interface{}{
	"foo":  123,
	"bar":  true,
	"zzz":  map[string]interface{}{"hey": 1, "bye": 2},
	"list": l,
	"bool": b,
}

func main() {
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "burn-true":
		burnTengoTrue()
	case "burn-lookup":
		burnTengo_50_array_lookup()
	default:
		dump()
	}
}

func burnTengoTrue() {
	code := `true`
	program := prepareTengo(code, nil)
	for {
		program.Clone().Run()
	}
}

func burnTengo_50_array_lookup() {
	code := `
fmt := import("fmt")
enum := import("enum")

l := [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50]
r := enum.any(l, func(k, v) { return v == 1001})
`
	program := prepareTengo(code, nil)
	for {
		program.Clone().Run()
	}
}

func dump() {
	code := `
fmt := import("fmt")

f("xxx", 1, {a: true});
v := 123
v = g()

counter = counter + 1

fmt.println(x)
fmt.println(o)
fmt.println(u)
fmt.println(v)
fmt.println("counter: ", counter)
fmt.println("--- END ---------------")
`

	uu := Foo{
		A: 123,
		B: "hey",
		C: &Bar{
			Xxx: "no",
			Yyy: true,
		},
	}

	b, _ := json.Marshal(uu)
	var u map[string]interface{}
	if err := json.Unmarshal(b, &u); err != nil {
		panic(err)
	}

	scr := tengo.NewScript([]byte(code))
	scr.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
	if err := scr.Add("counter", 0); err != nil {
		panic(err)
	}
	if err := scr.Add("o", o); err != nil {
		panic(err)
	}
	if err := scr.Add("u", u); err != nil {
		panic(err)
	}
	if err := scr.Add("x", 123); err != nil {
		panic(err)
	}
	scr.Add("f", &tengo.UserFunction{
		Value: func(args ...tengo.Object) (tengo.Object, error) {
			fmt.Println("Calling Go function here with args:", args)
			return nil, nil
		},
	})
	scr.Add("g", &tengo.UserFunction{
		Value: func(args ...tengo.Object) (tengo.Object, error) {
			return &tengo.String{Value: "here"}, nil
		},
	},
	)
	prog, err := scr.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("* Read globals:")
	pp.Println(prog.GetAll())
	fmt.Println("* VM bytecode:")
	fmt.Println(prog)

	fmt.Println("=== RUN PROGRAM AGAIN ============")
	err = prog.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("counter:", prog.Get("counter"))

	fmt.Println("=== COMPILER =====================")
	code = `
123
456
{a: 12.34, b: "bar"}
[1, false, "foo"]
`
	fileSet := parser.NewFileSet()
	srcFile := fileSet.AddFile("(main)", -1, len(code))
	p := parser.NewParser(srcFile, []byte(code), nil)
	file, err := p.ParseFile()
	if err != nil {
		panic(err)
	}

	comp := tengo.NewCompiler(srcFile, nil, nil, nil, nil)
	if err := comp.Compile(file); err != nil {
		panic(err)
	}

	fmt.Println("* VM instructions:")
	pp.Println(comp.Bytecode().FormatInstructions())
	fmt.Println("* VM constants table:")
	pp.Println(comp.Bytecode().FormatConstants())
	fmt.Println("==================================")
}

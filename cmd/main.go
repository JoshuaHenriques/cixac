package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"syscall/js"

	"github.com/joshuahenriques/cixac/evaluator"
	"github.com/joshuahenriques/cixac/lexer"
	"github.com/joshuahenriques/cixac/object"
	"github.com/joshuahenriques/cixac/parser"
	"github.com/joshuahenriques/cixac/repl"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	BuildVersion string = "0.1-alpha"
	BuildDate    string = "Aug 20 2024"
)

func main() {
	eFlag := flag.String("e", "", "Execute inline code: Specifies a string of code to be directly executed by the program")
	wFlag := flag.Bool("w", false, "WASM environment.")
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Printf("Cixac Version: %s (%s) on %s\n", BuildVersion, BuildDate, runtime.GOOS)
		fmt.Printf("Type \"quit()\" to exit the REPL\n")
		repl.Start(os.Stdin, os.Stdout)
	}

	switch {
	case isFlagPassed("e"):
		runProgram(*eFlag)
	case *wFlag:
		js.Global().Set("runScript", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) == 0 {
				return "Invalid number of arguments passed"
			}

			return runProgram(args[0].String())
		}))
		select {}
	default:
		file, err := os.ReadFile(os.Args[1])
		check(err)
		runProgram(string(file))
	}
}

func runProgram(code string) string {
	l := lexer.New(code)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(os.Stdout, p.Errors())
	}

	evaluated := evaluator.Eval(program, object.NewEnvironment())
	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
		return evaluated.Inspect() + "\n\n\n"
	}
	return ""
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

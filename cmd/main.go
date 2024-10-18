package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

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
	BuildVersion string = "0.5-beta"
	BuildDate    string = "Oct 03 2024"
)

func main() {
	eFlag := flag.String("e", "", "Execute inline code: Specifies a string of code to be directly executed by the program")
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Printf("Cixac Version: %s (%s) on %s\n", BuildVersion, BuildDate, runtime.GOOS)
		fmt.Printf("Use '\\' at the end of a line for multi-line input\n")
		fmt.Printf("Type \"quit()\" to exit the REPL\n")
		repl.Start(os.Stdin, os.Stdout)
	}

	if isFlagPassed("e") {
		runProgram(*eFlag)
	} else {
		file, err := os.ReadFile(os.Args[1])
		check(err)
		runProgram(string(file))
	}
}

func runProgram(code string) {
	l := lexer.New(code)
	p := parser.New(l)

	program := p.ParseProgram()

	evaluated := evaluator.Eval(program, object.NewEnvironment())
	if evaluated != nil && evaluated.Type() != object.EMPTY_OBJ {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	}

	if len(p.Errors()) != 0 {
		printParserErrors(os.Stdout, p.Errors())
	}
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

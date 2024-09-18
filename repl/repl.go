package repl

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/joshuahenriques/cixac/evaluator"
	"github.com/joshuahenriques/cixac/lexer"
	"github.com/joshuahenriques/cixac/object"
	"github.com/joshuahenriques/cixac/parser"
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func Start(in io.Reader, out io.Writer) {
	l, err := readline.NewEx(&readline.Config{
		Prompt:              "\033[31m»\033[0m ",
		HistoryFile:         "/tmp/readline.tmp",
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	l.CaptureExitSignal()

	env := object.NewEnvironment()

	log.SetOutput(l.Stderr())

	var multiLineBuffer strings.Builder
	isMultiLine := false

	for {
		var line string
		var err error

		if isMultiLine {
			l.SetPrompt("... ")
		} else {
			l.SetPrompt("\033[31m»\033[0m ")
		}

		line, err = l.Readline()

		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		switch {
		case line == "quit()":
			os.Exit(0)
		case strings.HasSuffix(line, `\`):
			multiLineBuffer.WriteString(strings.TrimSuffix(line, `\`) + "\n")
			isMultiLine = true
			continue
		}

		if isMultiLine {
			multiLineBuffer.WriteString(line + "\n")
			line = multiLineBuffer.String()
			multiLineBuffer.Reset()
			isMultiLine = false
		}

		line = strings.TrimSpace(line)
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil && evaluated.Type() != object.EMPTY_OBJ {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

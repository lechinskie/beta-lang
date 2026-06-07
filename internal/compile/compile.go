package compile

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/eramoss/b/internal/codegen"
	"github.com/eramoss/b/internal/parser"
	"github.com/eramoss/b/internal/sema"
)

func Parse(source string, errors *sema.ErrorList) (bool, antlr.ParseTree) {
	input := antlr.NewInputStream(source)

	lexer := parser.NewbLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(sema.NewLexerErrorListener(errors))

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewbParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(sema.NewParserErrorListener(errors))

	if errors.Count() > 0 {
		return false, nil
	}

	tree := p.Program()

	if errors.Count() > 0 {
		return false, tree
	}

	return true, tree
}

func Main() {
	showSt := false
	showAssembly := false
	fileArg := ""

	for _, arg := range os.Args[1:] {
		if arg == "--exp-st" {
			showSt = true
		} else if !strings.HasPrefix(arg, "-") {
			fileArg = arg
		} else if  arg == "--S" {
			showAssembly = true
		}
	}

	if fileArg == "" {
		fmt.Fprintln(os.Stderr, "usage: b [--exp-st] [--S] <file.b>")
		os.Exit(1)
	}

	data, err := os.ReadFile(fileArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	source := string(data)
	errors := sema.NewErrorList()

	ok, tree := Parse(source, errors)

	if errors.Count() > 0 {
		fmt.Print(errors.Format(source))
		os.Exit(1)
	}

	if !ok {
		fmt.Println("parse failed")
		os.Exit(1)
	}

	st := sema.CheckSemantic(source, errors, tree)

	if errors.Count() > 0 {
		fmt.Print(errors.Format(source))
		os.Exit(1)
	}

	out := codegen.Generate(source, errors, tree, st)

	if errors.Count() > 0 {
		fmt.Print(errors.Format(source))
		os.Exit(1)
	}
	
	if (showAssembly){
			fmt.Print(out.String())
	}

	if showSt {
		fmt.Print(st.String())
	}
}

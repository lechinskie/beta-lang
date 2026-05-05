package analysis

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/eramoss/b/internal/parser"
)

type SymbolKind int

const (
	KindVar SymbolKind = iota
	KindParam
	KindFunc
	KindExtrn
	KindArray
	KindLabel
)

func (k SymbolKind) String() string {
	switch k {
	case KindVar:
		return "var"
	case KindParam:
		return "param"
	case KindFunc:
		return "func"
	case KindExtrn:
		return "extrn"
	case KindArray:
		return "array"
	case KindLabel:
		return "label"
	default:
		return "unknown"
	}
}

type Symbol struct {
	Name       string
	Type       string
	Scope      int
	IsDeclared bool
	Kind       SymbolKind
	ParamPos   int // position in parameter list (-1 if not a parameter)
}

type SymbolTable struct {
	scopes []map[string]*Symbol
	depth  int
}

func NewSymbolTable() *SymbolTable {
	st := &SymbolTable{
		scopes: make([]map[string]*Symbol, 0),
		depth:  0,
	}
	st.EnterScope()
	return st
}

func (st *SymbolTable) EnterScope() {
	st.scopes = append(st.scopes, make(map[string]*Symbol))
	st.depth++
}

func (st *SymbolTable) ExitScope() {
	if st.depth > 0 {
		st.scopes = st.scopes[:st.depth-1]
		st.depth--
	}
}

func (st *SymbolTable) Insert(name string, sym *Symbol) error {
	current := st.scopes[st.depth-1]
	if existing, ok := current[name]; ok && existing.IsDeclared {
		return fmt.Errorf("%s:%d: variable '%s' already declared in this scope", sym.Kind, sym.Scope, name)
	}
	current[name] = sym
	return nil
}

func (st *SymbolTable) Lookup(name string) *Symbol {
	for i := st.depth - 1; i >= 0; i-- {
		if sym, ok := st.scopes[i][name]; ok {
			return sym
		}
	}
	return nil
}

func (st *SymbolTable) CurrentDepth() int {
	return st.depth
}

func Main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: b <file.b>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	source := string(data)
	errors := NewErrorList()

	ok, tree := Parse(source, errors)

	if errors.Count() > 0 {
		fmt.Print(errors.Format(source))
		os.Exit(1)
	}

	if !ok {
		fmt.Println("parse failed")
		os.Exit(1)
	}

	CheckSemantic(source, errors, tree)

	if errors.Count() > 0 {
		fmt.Print(errors.Format(source))
		os.Exit(1)
	}
}

type ErrorKind int

const (
	KindLexer ErrorKind = iota
	KindParser
)

type Error struct {
	Kind   ErrorKind
	Line   int
	Column int
	Text   string
}

func (e *Error) String() string {
	loc := fmt.Sprintf("%d:%d", e.Line, e.Column)
	switch e.Kind {
	case KindLexer:
		return fmt.Sprintf("lexer error at %s: %s", loc, e.Text)
	case KindParser:
		return fmt.Sprintf("syntax error at %s: %s", loc, e.Text)
	default:
		return fmt.Sprintf("error at %s: %s", loc, e.Text)
	}
}

type ErrorList struct {
	errors []*Error
}

func NewErrorList() *ErrorList {
	return &ErrorList{errors: make([]*Error, 0)}
}

func (el *ErrorList) Add(err *Error) {
	el.errors = append(el.errors, err)
}

func (el *ErrorList) AddLexer(line, col int, text string) {
	el.Add(&Error{Kind: KindLexer, Line: line, Column: col, Text: text})
}

func (el *ErrorList) AddParser(line, col int, text string) {
	el.Add(&Error{Kind: KindParser, Line: line, Column: col, Text: text})
}

func (el *ErrorList) Count() int {
	return len(el.errors)
}

func (el *ErrorList) Format(source string) string {
	if len(el.errors) == 0 {
		return ""
	}

	var b strings.Builder
	fmt.Fprintf(&b, "found %d error(s):\n\n", len(el.errors))

	srcLines := strings.Split(source, "\n")

	for i, e := range el.errors {
		fmt.Fprintf(&b, "%d: %s\n", i+1, e.String())

		if e.Line > 0 && e.Line <= len(srcLines) {
			line := srcLines[e.Line-1]
			fmt.Fprintf(&b, "    | %s\n", line)

			if e.Column > 0 && e.Column <= len(line) {
				spaces := strings.Repeat(" ", e.Column-1)
				fmt.Fprintf(&b, "    | %s^", spaces)

				b.WriteByte('\n')
			}
		}
		b.WriteByte('\n')
	}

	return b.String()
}

// --- Lexer error definitions ---
type lexerErrorListener struct {
	errors *ErrorList
}

func newLexerErrorListener(errors *ErrorList) *lexerErrorListener {
	return &lexerErrorListener{errors: errors}
}

func (l *lexerErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	l.errors.AddLexer(line, column, msg)
}

func (l *lexerErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}
func (l *lexerErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}
func (l *lexerErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}

// --- Parser error definitions ---
type parserErrorListener struct {
	errors *ErrorList
}

func newParserErrorListener(errors *ErrorList) *parserErrorListener {
	return &parserErrorListener{errors: errors}
}

func (l *parserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	l.errors.AddParser(line, column, msg)
}

func (l *parserErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}
func (l *parserErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}
func (l *parserErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}

func Parse(source string, errors *ErrorList) (bool, antlr.ParseTree) {
	input := antlr.NewInputStream(source)

	lexer := parser.NewbLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(newLexerErrorListener(errors))

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewbParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(newParserErrorListener(errors))

	if errors.Count() > 0 {
		return false, nil
	}

	tree := p.Program()

	if errors.Count() > 0 {
		return false, tree
	}

	return true, tree
}

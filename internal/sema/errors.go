package sema

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

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

type LexerErrorListener struct {
	errors *ErrorList
}

func NewLexerErrorListener(errors *ErrorList) *LexerErrorListener {
	return &LexerErrorListener{errors: errors}
}

func (l *LexerErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	l.errors.AddLexer(line, column, msg)
}

func (l *LexerErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (l *LexerErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (l *LexerErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}

type ParserErrorListener struct {
	errors *ErrorList
}

func NewParserErrorListener(errors *ErrorList) *ParserErrorListener {
	return &ParserErrorListener{errors: errors}
}

func (l *ParserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	l.errors.AddParser(line, column, msg)
}

func (l *ParserErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (l *ParserErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (l *ParserErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}

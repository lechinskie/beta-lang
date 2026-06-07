package sema

import (
	"fmt"
	"strings"
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
	ParamPos   int
}

type SymbolTable struct {
	scopes []map[string]*Symbol
	depth  int
	all    []*Symbol
}

func NewSymbolTable() *SymbolTable {
	st := &SymbolTable{
		scopes: make([]map[string]*Symbol, 0),
		depth:  0,
		all:    make([]*Symbol, 0),
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
	st.all = append(st.all, sym)
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

func typeDisplay(t string) string {
	switch t {
	case "word":
		return "int"
	case "":
		return "auto"
	default:
		return t
	}
}

func (st *SymbolTable) String() string {
	var b strings.Builder
	for _, sym := range st.all {
		fmt.Fprintf(&b, "  %-8s %s", sym.Kind.String(), sym.Name)
		fmt.Fprintf(&b, " (type=%s", typeDisplay(sym.Type))
		if sym.ParamPos >= 0 {
			fmt.Fprintf(&b, ", scope=%d, paramPos=%d", sym.Scope, sym.ParamPos)
		} else {
			fmt.Fprintf(&b, ", scope=%d", sym.Scope)
		}
		fmt.Fprintf(&b, ")\n")
	}
	return b.String()
}

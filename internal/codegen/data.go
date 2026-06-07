package codegen

import (
	"strings"

	"github.com/eramoss/b/internal/parser"
)

func (g *Generator) emitData() {
	g.emitDirective(".data")
	g.emitDirective("__nl__: .asciz \"\\n\"")

	for name, size := range g.arraySizes {
		var line strings.Builder
		line.WriteString(name)
		line.WriteString(": .word")
		for i := range size {
			if i > 0 {
				line.WriteString(",")
			}
			line.WriteString(" 0")
		}
		g.emitDirective(line.String())
	}
}

func (g *Generator) EnterExt_def(ctx *parser.Ext_defContext) {
	if ctx.Statement() == nil {
		return
	}

	g.inFunc = true
	g.funcName = ctx.ID().GetText()

	g.emitDirective(".globl %s", g.funcName)
	g.emitLabel(g.funcName)
	if g.frameSize > 0 {
		g.emit("addi sp, sp, -%d", g.frameSize)
	}
}

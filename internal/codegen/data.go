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
		line.WriteString(mangle(name))
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
	g.isFirst = g.funcName == "main"

	g.localOffs = g.funcLocals[g.funcName]
	paramCount := g.funcParams[g.funcName]

	locals := len(g.localOffs)
	needed := max((locals+maxTemps)*4, 16)
	needed = (needed + 15) & ^15

	for name := range g.localOffs {
		g.localOffs[name] += maxTemps * 4
	}

	g.frameSize = needed

	g.emitDirective(".globl %s", mangle(g.funcName))
	g.emitLabel(mangle(g.funcName))
	g.emit("addi sp, sp, -%d", g.frameSize)

	if !g.isFirst {
		g.emit("sw ra, %d(sp)", g.frameSize-4)
	}

	if paramCount > 0 {
		for _, off := range g.localOffs {
			if off < maxTemps*4+paramCount*4 {
				argIdx := (off - maxTemps*4) / 4
				g.emit("addi t0, a%d, 0", argIdx)
				g.emit("sw t0, %d(sp)", off)
			}
		}
	}
}

func (g *Generator) ExitExt_def(ctx *parser.Ext_defContext) {
	if ctx.Statement() == nil {
		return
	}

	if g.isFirst {
		g.emitExit()
	} else {
		g.emit("lw ra, %d(sp)", g.frameSize-4)
		g.emit("addi sp, sp, %d", g.frameSize)
		g.emit("jalr zero, ra, 0")
	}

	g.inFunc = false
	g.funcName = ""
	g.localOffs = nil
	g.frameSize = 0
}

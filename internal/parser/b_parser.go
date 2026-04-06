// Code generated from b.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // b

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type bParser struct {
	*antlr.BaseParser
}

var BParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bParserInit() {
	staticData := &BParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "','", "'__asm__'", "'('", "')'", "'extrn'", "'__variadic__'",
		"':'", "'{'", "'}'", "'return'", "'goto'", "'switch'", "'while'", "'if'",
		"'else'", "'case'", "'auto'", "'?'", "'&'", "'='", "'++'", "'--'", "'-'",
		"'!'", "'|'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'<<'",
		"'>>'", "'+'", "'%'", "'*'", "'/'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "NAME", "INT", "STRING1", "STRING2", "LINECOMMENT",
		"BLOCKCOMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "topLevel", "definition", "extrndecl", "variadicdecl", "ival",
		"statement", "nullstmt", "expressionstmt", "blockstmt", "returnstmt",
		"gotostmt", "switchstmt", "whilestmt", "ifstmt", "casestmt", "externsmt",
		"autosmt", "asmstmt", "stringlist", "rvalue", "ternary", "comparison",
		"assignment", "expression", "functioninvocation", "functionparameters",
		"assign", "incdec", "unary", "binary", "lvalue", "constant", "name",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 339, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 5, 0, 70, 8, 0, 10, 0, 12, 0, 73,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 81, 8, 1, 1, 2, 1, 2, 3,
		2, 85, 8, 2, 1, 2, 1, 2, 1, 2, 5, 2, 90, 8, 2, 10, 2, 12, 2, 93, 9, 2,
		5, 2, 95, 8, 2, 10, 2, 12, 2, 98, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 114, 8, 2, 10,
		2, 12, 2, 117, 9, 2, 3, 2, 119, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 124, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 130, 8, 3, 10, 3, 12, 3, 133, 9, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5,
		147, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 165, 8, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 174, 8, 9, 10, 9, 12, 9, 177, 9, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 186, 8, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 211, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 5, 16, 222, 8, 16, 10, 16, 12, 16, 225, 9, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 3, 17, 232, 8, 17, 1, 17, 1, 17, 1, 17, 3, 17, 237,
		8, 17, 5, 17, 239, 8, 17, 10, 17, 12, 17, 242, 9, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 255,
		8, 19, 10, 19, 12, 19, 258, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 264,
		8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 3, 24, 298, 8, 24, 1, 25, 1, 25, 1, 25, 3, 25, 303, 8, 25, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 26, 5, 26, 310, 8, 26, 10, 26, 12, 26, 313,
		9, 26, 1, 27, 1, 27, 3, 27, 317, 8, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31,
		333, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 0, 0, 34, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 0, 4, 1, 0, 22, 23, 1, 0, 24,
		25, 3, 0, 20, 20, 24, 24, 26, 38, 1, 0, 42, 44, 352, 0, 71, 1, 0, 0, 0,
		2, 80, 1, 0, 0, 0, 4, 123, 1, 0, 0, 0, 6, 125, 1, 0, 0, 0, 8, 136, 1, 0,
		0, 0, 10, 146, 1, 0, 0, 0, 12, 164, 1, 0, 0, 0, 14, 166, 1, 0, 0, 0, 16,
		168, 1, 0, 0, 0, 18, 171, 1, 0, 0, 0, 20, 180, 1, 0, 0, 0, 22, 189, 1,
		0, 0, 0, 24, 193, 1, 0, 0, 0, 26, 197, 1, 0, 0, 0, 28, 203, 1, 0, 0, 0,
		30, 212, 1, 0, 0, 0, 32, 217, 1, 0, 0, 0, 34, 228, 1, 0, 0, 0, 36, 245,
		1, 0, 0, 0, 38, 251, 1, 0, 0, 0, 40, 263, 1, 0, 0, 0, 42, 265, 1, 0, 0,
		0, 44, 271, 1, 0, 0, 0, 46, 275, 1, 0, 0, 0, 48, 297, 1, 0, 0, 0, 50, 299,
		1, 0, 0, 0, 52, 306, 1, 0, 0, 0, 54, 314, 1, 0, 0, 0, 56, 318, 1, 0, 0,
		0, 58, 320, 1, 0, 0, 0, 60, 322, 1, 0, 0, 0, 62, 332, 1, 0, 0, 0, 64, 334,
		1, 0, 0, 0, 66, 336, 1, 0, 0, 0, 68, 70, 3, 2, 1, 0, 69, 68, 1, 0, 0, 0,
		70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74, 1,
		0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 75, 5, 0, 0, 1, 75, 1, 1, 0, 0, 0, 76,
		81, 3, 4, 2, 0, 77, 81, 3, 6, 3, 0, 78, 81, 3, 8, 4, 0, 79, 81, 5, 1, 0,
		0, 80, 76, 1, 0, 0, 0, 80, 77, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 79,
		1, 0, 0, 0, 81, 3, 1, 0, 0, 0, 82, 84, 3, 66, 33, 0, 83, 85, 3, 64, 32,
		0, 84, 83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 96, 1, 0, 0, 0, 86, 91,
		3, 10, 5, 0, 87, 88, 5, 2, 0, 0, 88, 90, 3, 10, 5, 0, 89, 87, 1, 0, 0,
		0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 95,
		1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 86, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0,
		96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 99, 1, 0, 0, 0, 98, 96, 1,
		0, 0, 0, 99, 100, 5, 1, 0, 0, 100, 124, 1, 0, 0, 0, 101, 102, 3, 66, 33,
		0, 102, 103, 5, 3, 0, 0, 103, 104, 5, 4, 0, 0, 104, 105, 3, 38, 19, 0,
		105, 106, 5, 5, 0, 0, 106, 107, 5, 1, 0, 0, 107, 124, 1, 0, 0, 0, 108,
		109, 3, 66, 33, 0, 109, 118, 5, 4, 0, 0, 110, 115, 3, 66, 33, 0, 111, 112,
		5, 2, 0, 0, 112, 114, 3, 66, 33, 0, 113, 111, 1, 0, 0, 0, 114, 117, 1,
		0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 119, 1, 0, 0,
		0, 117, 115, 1, 0, 0, 0, 118, 110, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		120, 1, 0, 0, 0, 120, 121, 5, 5, 0, 0, 121, 122, 3, 12, 6, 0, 122, 124,
		1, 0, 0, 0, 123, 82, 1, 0, 0, 0, 123, 101, 1, 0, 0, 0, 123, 108, 1, 0,
		0, 0, 124, 5, 1, 0, 0, 0, 125, 126, 5, 6, 0, 0, 126, 131, 3, 66, 33, 0,
		127, 128, 5, 2, 0, 0, 128, 130, 3, 66, 33, 0, 129, 127, 1, 0, 0, 0, 130,
		133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134,
		1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 135, 5, 1, 0, 0, 135, 7, 1, 0, 0,
		0, 136, 137, 5, 7, 0, 0, 137, 138, 5, 4, 0, 0, 138, 139, 3, 66, 33, 0,
		139, 140, 5, 2, 0, 0, 140, 141, 5, 42, 0, 0, 141, 142, 5, 5, 0, 0, 142,
		143, 5, 1, 0, 0, 143, 9, 1, 0, 0, 0, 144, 147, 3, 64, 32, 0, 145, 147,
		3, 66, 33, 0, 146, 144, 1, 0, 0, 0, 146, 145, 1, 0, 0, 0, 147, 11, 1, 0,
		0, 0, 148, 165, 3, 32, 16, 0, 149, 165, 3, 34, 17, 0, 150, 151, 3, 66,
		33, 0, 151, 152, 5, 8, 0, 0, 152, 153, 3, 12, 6, 0, 153, 165, 1, 0, 0,
		0, 154, 165, 3, 30, 15, 0, 155, 165, 3, 18, 9, 0, 156, 165, 3, 28, 14,
		0, 157, 165, 3, 26, 13, 0, 158, 165, 3, 24, 12, 0, 159, 165, 3, 22, 11,
		0, 160, 165, 3, 20, 10, 0, 161, 165, 3, 36, 18, 0, 162, 165, 3, 16, 8,
		0, 163, 165, 3, 14, 7, 0, 164, 148, 1, 0, 0, 0, 164, 149, 1, 0, 0, 0, 164,
		150, 1, 0, 0, 0, 164, 154, 1, 0, 0, 0, 164, 155, 1, 0, 0, 0, 164, 156,
		1, 0, 0, 0, 164, 157, 1, 0, 0, 0, 164, 158, 1, 0, 0, 0, 164, 159, 1, 0,
		0, 0, 164, 160, 1, 0, 0, 0, 164, 161, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0,
		164, 163, 1, 0, 0, 0, 165, 13, 1, 0, 0, 0, 166, 167, 5, 1, 0, 0, 167, 15,
		1, 0, 0, 0, 168, 169, 3, 40, 20, 0, 169, 170, 5, 1, 0, 0, 170, 17, 1, 0,
		0, 0, 171, 175, 5, 9, 0, 0, 172, 174, 3, 12, 6, 0, 173, 172, 1, 0, 0, 0,
		174, 177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176,
		178, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178, 179, 5, 10, 0, 0, 179, 19,
		1, 0, 0, 0, 180, 185, 5, 11, 0, 0, 181, 182, 5, 4, 0, 0, 182, 183, 3, 40,
		20, 0, 183, 184, 5, 5, 0, 0, 184, 186, 1, 0, 0, 0, 185, 181, 1, 0, 0, 0,
		185, 186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 5, 1, 0, 0, 188,
		21, 1, 0, 0, 0, 189, 190, 5, 12, 0, 0, 190, 191, 3, 40, 20, 0, 191, 192,
		5, 1, 0, 0, 192, 23, 1, 0, 0, 0, 193, 194, 5, 13, 0, 0, 194, 195, 3, 40,
		20, 0, 195, 196, 3, 12, 6, 0, 196, 25, 1, 0, 0, 0, 197, 198, 5, 14, 0,
		0, 198, 199, 5, 4, 0, 0, 199, 200, 3, 40, 20, 0, 200, 201, 5, 5, 0, 0,
		201, 202, 3, 12, 6, 0, 202, 27, 1, 0, 0, 0, 203, 204, 5, 15, 0, 0, 204,
		205, 5, 4, 0, 0, 205, 206, 3, 40, 20, 0, 206, 207, 5, 5, 0, 0, 207, 210,
		3, 12, 6, 0, 208, 209, 5, 16, 0, 0, 209, 211, 3, 12, 6, 0, 210, 208, 1,
		0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 29, 1, 0, 0, 0, 212, 213, 5, 17, 0,
		0, 213, 214, 3, 64, 32, 0, 214, 215, 5, 8, 0, 0, 215, 216, 3, 12, 6, 0,
		216, 31, 1, 0, 0, 0, 217, 218, 5, 6, 0, 0, 218, 223, 3, 66, 33, 0, 219,
		220, 5, 2, 0, 0, 220, 222, 3, 66, 33, 0, 221, 219, 1, 0, 0, 0, 222, 225,
		1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 226, 1, 0,
		0, 0, 225, 223, 1, 0, 0, 0, 226, 227, 5, 1, 0, 0, 227, 33, 1, 0, 0, 0,
		228, 229, 5, 18, 0, 0, 229, 231, 3, 66, 33, 0, 230, 232, 3, 64, 32, 0,
		231, 230, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 240, 1, 0, 0, 0, 233,
		234, 5, 2, 0, 0, 234, 236, 3, 66, 33, 0, 235, 237, 3, 64, 32, 0, 236, 235,
		1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 239, 1, 0, 0, 0, 238, 233, 1, 0,
		0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0,
		241, 243, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 244, 5, 1, 0, 0, 244,
		35, 1, 0, 0, 0, 245, 246, 5, 3, 0, 0, 246, 247, 5, 4, 0, 0, 247, 248, 3,
		38, 19, 0, 248, 249, 5, 5, 0, 0, 249, 250, 5, 1, 0, 0, 250, 37, 1, 0, 0,
		0, 251, 256, 5, 43, 0, 0, 252, 253, 5, 2, 0, 0, 253, 255, 5, 43, 0, 0,
		254, 252, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 256,
		257, 1, 0, 0, 0, 257, 39, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 259, 264, 3,
		48, 24, 0, 260, 264, 3, 44, 22, 0, 261, 264, 3, 42, 21, 0, 262, 264, 3,
		46, 23, 0, 263, 259, 1, 0, 0, 0, 263, 260, 1, 0, 0, 0, 263, 261, 1, 0,
		0, 0, 263, 262, 1, 0, 0, 0, 264, 41, 1, 0, 0, 0, 265, 266, 3, 48, 24, 0,
		266, 267, 5, 19, 0, 0, 267, 268, 3, 40, 20, 0, 268, 269, 5, 8, 0, 0, 269,
		270, 3, 40, 20, 0, 270, 43, 1, 0, 0, 0, 271, 272, 3, 48, 24, 0, 272, 273,
		3, 60, 30, 0, 273, 274, 3, 40, 20, 0, 274, 45, 1, 0, 0, 0, 275, 276, 3,
		66, 33, 0, 276, 277, 3, 54, 27, 0, 277, 278, 3, 40, 20, 0, 278, 47, 1,
		0, 0, 0, 279, 280, 5, 4, 0, 0, 280, 281, 3, 40, 20, 0, 281, 282, 5, 5,
		0, 0, 282, 298, 1, 0, 0, 0, 283, 298, 3, 66, 33, 0, 284, 298, 3, 64, 32,
		0, 285, 286, 3, 56, 28, 0, 286, 287, 3, 66, 33, 0, 287, 298, 1, 0, 0, 0,
		288, 289, 3, 66, 33, 0, 289, 290, 3, 56, 28, 0, 290, 298, 1, 0, 0, 0, 291,
		292, 3, 58, 29, 0, 292, 293, 3, 40, 20, 0, 293, 298, 1, 0, 0, 0, 294, 295,
		5, 20, 0, 0, 295, 298, 3, 66, 33, 0, 296, 298, 3, 50, 25, 0, 297, 279,
		1, 0, 0, 0, 297, 283, 1, 0, 0, 0, 297, 284, 1, 0, 0, 0, 297, 285, 1, 0,
		0, 0, 297, 288, 1, 0, 0, 0, 297, 291, 1, 0, 0, 0, 297, 294, 1, 0, 0, 0,
		297, 296, 1, 0, 0, 0, 298, 49, 1, 0, 0, 0, 299, 300, 3, 66, 33, 0, 300,
		302, 5, 4, 0, 0, 301, 303, 3, 52, 26, 0, 302, 301, 1, 0, 0, 0, 302, 303,
		1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305, 5, 5, 0, 0, 305, 51, 1, 0,
		0, 0, 306, 311, 3, 40, 20, 0, 307, 308, 5, 2, 0, 0, 308, 310, 3, 40, 20,
		0, 309, 307, 1, 0, 0, 0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 311,
		312, 1, 0, 0, 0, 312, 53, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 316, 5,
		21, 0, 0, 315, 317, 3, 60, 30, 0, 316, 315, 1, 0, 0, 0, 316, 317, 1, 0,
		0, 0, 317, 55, 1, 0, 0, 0, 318, 319, 7, 0, 0, 0, 319, 57, 1, 0, 0, 0, 320,
		321, 7, 1, 0, 0, 321, 59, 1, 0, 0, 0, 322, 323, 7, 2, 0, 0, 323, 61, 1,
		0, 0, 0, 324, 333, 3, 66, 33, 0, 325, 326, 5, 37, 0, 0, 326, 333, 3, 40,
		20, 0, 327, 328, 3, 40, 20, 0, 328, 329, 5, 39, 0, 0, 329, 330, 3, 40,
		20, 0, 330, 331, 5, 40, 0, 0, 331, 333, 1, 0, 0, 0, 332, 324, 1, 0, 0,
		0, 332, 325, 1, 0, 0, 0, 332, 327, 1, 0, 0, 0, 333, 63, 1, 0, 0, 0, 334,
		335, 7, 3, 0, 0, 335, 65, 1, 0, 0, 0, 336, 337, 5, 41, 0, 0, 337, 67, 1,
		0, 0, 0, 25, 71, 80, 84, 91, 96, 115, 118, 123, 131, 146, 164, 175, 185,
		210, 223, 231, 236, 240, 256, 263, 297, 302, 311, 316, 332,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// bParserInit initializes any static state used to implement bParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewbParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BParserInit() {
	staticData := &BParserStaticData
	staticData.once.Do(bParserInit)
}

// NewbParser produces a new parser instance for the optional input antlr.TokenStream.
func NewbParser(input antlr.TokenStream) *bParser {
	BParserInit()
	this := new(bParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "b.g4"

	return this
}

// bParser tokens.
const (
	bParserEOF          = antlr.TokenEOF
	bParserT__0         = 1
	bParserT__1         = 2
	bParserT__2         = 3
	bParserT__3         = 4
	bParserT__4         = 5
	bParserT__5         = 6
	bParserT__6         = 7
	bParserT__7         = 8
	bParserT__8         = 9
	bParserT__9         = 10
	bParserT__10        = 11
	bParserT__11        = 12
	bParserT__12        = 13
	bParserT__13        = 14
	bParserT__14        = 15
	bParserT__15        = 16
	bParserT__16        = 17
	bParserT__17        = 18
	bParserT__18        = 19
	bParserT__19        = 20
	bParserT__20        = 21
	bParserT__21        = 22
	bParserT__22        = 23
	bParserT__23        = 24
	bParserT__24        = 25
	bParserT__25        = 26
	bParserT__26        = 27
	bParserT__27        = 28
	bParserT__28        = 29
	bParserT__29        = 30
	bParserT__30        = 31
	bParserT__31        = 32
	bParserT__32        = 33
	bParserT__33        = 34
	bParserT__34        = 35
	bParserT__35        = 36
	bParserT__36        = 37
	bParserT__37        = 38
	bParserT__38        = 39
	bParserT__39        = 40
	bParserNAME         = 41
	bParserINT          = 42
	bParserSTRING1      = 43
	bParserSTRING2      = 44
	bParserLINECOMMENT  = 45
	bParserBLOCKCOMMENT = 46
	bParserWS           = 47
)

// bParser rules.
const (
	bParserRULE_program            = 0
	bParserRULE_topLevel           = 1
	bParserRULE_definition         = 2
	bParserRULE_extrndecl          = 3
	bParserRULE_variadicdecl       = 4
	bParserRULE_ival               = 5
	bParserRULE_statement          = 6
	bParserRULE_nullstmt           = 7
	bParserRULE_expressionstmt     = 8
	bParserRULE_blockstmt          = 9
	bParserRULE_returnstmt         = 10
	bParserRULE_gotostmt           = 11
	bParserRULE_switchstmt         = 12
	bParserRULE_whilestmt          = 13
	bParserRULE_ifstmt             = 14
	bParserRULE_casestmt           = 15
	bParserRULE_externsmt          = 16
	bParserRULE_autosmt            = 17
	bParserRULE_asmstmt            = 18
	bParserRULE_stringlist         = 19
	bParserRULE_rvalue             = 20
	bParserRULE_ternary            = 21
	bParserRULE_comparison         = 22
	bParserRULE_assignment         = 23
	bParserRULE_expression         = 24
	bParserRULE_functioninvocation = 25
	bParserRULE_functionparameters = 26
	bParserRULE_assign             = 27
	bParserRULE_incdec             = 28
	bParserRULE_unary              = 29
	bParserRULE_binary             = 30
	bParserRULE_lvalue             = 31
	bParserRULE_constant           = 32
	bParserRULE_name               = 33
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllTopLevel() []ITopLevelContext
	TopLevel(i int) ITopLevelContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(bParserEOF, 0)
}

func (s *ProgramContext) AllTopLevel() []ITopLevelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelContext); ok {
			tst[i] = t.(ITopLevelContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) TopLevel(i int) ITopLevelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *bParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, bParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199023255746) != 0 {
		{
			p.SetState(68)
			p.TopLevel()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(74)
		p.Match(bParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopLevelContext is an interface to support dynamic dispatch.
type ITopLevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Definition() IDefinitionContext
	Extrndecl() IExtrndeclContext
	Variadicdecl() IVariadicdeclContext

	// IsTopLevelContext differentiates from other interfaces.
	IsTopLevelContext()
}

type TopLevelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelContext() *TopLevelContext {
	var p = new(TopLevelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_topLevel
	return p
}

func InitEmptyTopLevelContext(p *TopLevelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_topLevel
}

func (*TopLevelContext) IsTopLevelContext() {}

func NewTopLevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelContext {
	var p = new(TopLevelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_topLevel

	return p
}

func (s *TopLevelContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelContext) Definition() IDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *TopLevelContext) Extrndecl() IExtrndeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtrndeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtrndeclContext)
}

func (s *TopLevelContext) Variadicdecl() IVariadicdeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariadicdeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariadicdeclContext)
}

func (s *TopLevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterTopLevel(s)
	}
}

func (s *TopLevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitTopLevel(s)
	}
}

func (p *bParser) TopLevel() (localctx ITopLevelContext) {
	localctx = NewTopLevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, bParserRULE_topLevel)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Definition()
		}

	case bParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Extrndecl()
		}

	case bParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Variadicdecl()
		}

	case bParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(79)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllName() []INameContext
	Name(i int) INameContext
	Constant() IConstantContext
	AllIval() []IIvalContext
	Ival(i int) IIvalContext
	Stringlist() IStringlistContext
	Statement() IStatementContext

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_definition
	return p
}

func InitEmptyDefinitionContext(p *DefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_definition
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *DefinitionContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DefinitionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *DefinitionContext) AllIval() []IIvalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIvalContext); ok {
			len++
		}
	}

	tst := make([]IIvalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIvalContext); ok {
			tst[i] = t.(IIvalContext)
			i++
		}
	}

	return tst
}

func (s *DefinitionContext) Ival(i int) IIvalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIvalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIvalContext)
}

func (s *DefinitionContext) Stringlist() IStringlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringlistContext)
}

func (s *DefinitionContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *bParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, bParserRULE_definition)
	var _la int

	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Name()
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(83)
				p.Constant()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32985348833280) != 0 {
			{
				p.SetState(86)
				p.Ival()
			}
			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == bParserT__1 {
				{
					p.SetState(87)
					p.Match(bParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(88)
					p.Ival()
				}

				p.SetState(93)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(99)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Name()
		}
		{
			p.SetState(102)
			p.Match(bParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Stringlist()
		}
		{
			p.SetState(105)
			p.Match(bParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Name()
		}
		{
			p.SetState(109)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == bParserNAME {
			{
				p.SetState(110)
				p.Name()
			}
			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == bParserT__1 {
				{
					p.SetState(111)
					p.Match(bParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(112)
					p.Name()
				}

				p.SetState(117)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(120)
			p.Match(bParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Statement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtrndeclContext is an interface to support dynamic dispatch.
type IExtrndeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllName() []INameContext
	Name(i int) INameContext

	// IsExtrndeclContext differentiates from other interfaces.
	IsExtrndeclContext()
}

type ExtrndeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtrndeclContext() *ExtrndeclContext {
	var p = new(ExtrndeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_extrndecl
	return p
}

func InitEmptyExtrndeclContext(p *ExtrndeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_extrndecl
}

func (*ExtrndeclContext) IsExtrndeclContext() {}

func NewExtrndeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtrndeclContext {
	var p = new(ExtrndeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_extrndecl

	return p
}

func (s *ExtrndeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtrndeclContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *ExtrndeclContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExtrndeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtrndeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtrndeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExtrndecl(s)
	}
}

func (s *ExtrndeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExtrndecl(s)
	}
}

func (p *bParser) Extrndecl() (localctx IExtrndeclContext) {
	localctx = NewExtrndeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, bParserRULE_extrndecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(bParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Name()
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__1 {
		{
			p.SetState(127)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Name()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(bParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariadicdeclContext is an interface to support dynamic dispatch.
type IVariadicdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	INT() antlr.TerminalNode

	// IsVariadicdeclContext differentiates from other interfaces.
	IsVariadicdeclContext()
}

type VariadicdeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariadicdeclContext() *VariadicdeclContext {
	var p = new(VariadicdeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_variadicdecl
	return p
}

func InitEmptyVariadicdeclContext(p *VariadicdeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_variadicdecl
}

func (*VariadicdeclContext) IsVariadicdeclContext() {}

func NewVariadicdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariadicdeclContext {
	var p = new(VariadicdeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_variadicdecl

	return p
}

func (s *VariadicdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariadicdeclContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VariadicdeclContext) INT() antlr.TerminalNode {
	return s.GetToken(bParserINT, 0)
}

func (s *VariadicdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariadicdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariadicdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterVariadicdecl(s)
	}
}

func (s *VariadicdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitVariadicdecl(s)
	}
}

func (p *bParser) Variadicdecl() (localctx IVariadicdeclContext) {
	localctx = NewVariadicdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, bParserRULE_variadicdecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(bParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(bParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Name()
	}
	{
		p.SetState(139)
		p.Match(bParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(bParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(bParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(bParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIvalContext is an interface to support dynamic dispatch.
type IIvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constant() IConstantContext
	Name() INameContext

	// IsIvalContext differentiates from other interfaces.
	IsIvalContext()
}

type IvalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIvalContext() *IvalContext {
	var p = new(IvalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_ival
	return p
}

func InitEmptyIvalContext(p *IvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_ival
}

func (*IvalContext) IsIvalContext() {}

func NewIvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IvalContext {
	var p = new(IvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_ival

	return p
}

func (s *IvalContext) GetParser() antlr.Parser { return s.parser }

func (s *IvalContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *IvalContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *IvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterIval(s)
	}
}

func (s *IvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitIval(s)
	}
}

func (p *bParser) Ival() (localctx IIvalContext) {
	localctx = NewIvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, bParserRULE_ival)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bParserINT, bParserSTRING1, bParserSTRING2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Constant()
		}

	case bParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Name()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Externsmt() IExternsmtContext
	Autosmt() IAutosmtContext
	Name() INameContext
	Statement() IStatementContext
	Casestmt() ICasestmtContext
	Blockstmt() IBlockstmtContext
	Ifstmt() IIfstmtContext
	Whilestmt() IWhilestmtContext
	Switchstmt() ISwitchstmtContext
	Gotostmt() IGotostmtContext
	Returnstmt() IReturnstmtContext
	Asmstmt() IAsmstmtContext
	Expressionstmt() IExpressionstmtContext
	Nullstmt() INullstmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Externsmt() IExternsmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExternsmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExternsmtContext)
}

func (s *StatementContext) Autosmt() IAutosmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAutosmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAutosmtContext)
}

func (s *StatementContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *StatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) Casestmt() ICasestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasestmtContext)
}

func (s *StatementContext) Blockstmt() IBlockstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockstmtContext)
}

func (s *StatementContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *StatementContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *StatementContext) Switchstmt() ISwitchstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstmtContext)
}

func (s *StatementContext) Gotostmt() IGotostmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGotostmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGotostmtContext)
}

func (s *StatementContext) Returnstmt() IReturnstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnstmtContext)
}

func (s *StatementContext) Asmstmt() IAsmstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsmstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsmstmtContext)
}

func (s *StatementContext) Expressionstmt() IExpressionstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionstmtContext)
}

func (s *StatementContext) Nullstmt() INullstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INullstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INullstmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *bParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, bParserRULE_statement)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Externsmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Autosmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)
			p.Name()
		}
		{
			p.SetState(151)
			p.Match(bParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(154)
			p.Casestmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(155)
			p.Blockstmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(156)
			p.Ifstmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(157)
			p.Whilestmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(158)
			p.Switchstmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(159)
			p.Gotostmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(160)
			p.Returnstmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(161)
			p.Asmstmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(162)
			p.Expressionstmt()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(163)
			p.Nullstmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INullstmtContext is an interface to support dynamic dispatch.
type INullstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNullstmtContext differentiates from other interfaces.
	IsNullstmtContext()
}

type NullstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullstmtContext() *NullstmtContext {
	var p = new(NullstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_nullstmt
	return p
}

func InitEmptyNullstmtContext(p *NullstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_nullstmt
}

func (*NullstmtContext) IsNullstmtContext() {}

func NewNullstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullstmtContext {
	var p = new(NullstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_nullstmt

	return p
}

func (s *NullstmtContext) GetParser() antlr.Parser { return s.parser }
func (s *NullstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterNullstmt(s)
	}
}

func (s *NullstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitNullstmt(s)
	}
}

func (p *bParser) Nullstmt() (localctx INullstmtContext) {
	localctx = NewNullstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, bParserRULE_nullstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(bParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionstmtContext is an interface to support dynamic dispatch.
type IExpressionstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Rvalue() IRvalueContext

	// IsExpressionstmtContext differentiates from other interfaces.
	IsExpressionstmtContext()
}

type ExpressionstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionstmtContext() *ExpressionstmtContext {
	var p = new(ExpressionstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_expressionstmt
	return p
}

func InitEmptyExpressionstmtContext(p *ExpressionstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_expressionstmt
}

func (*ExpressionstmtContext) IsExpressionstmtContext() {}

func NewExpressionstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionstmtContext {
	var p = new(ExpressionstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_expressionstmt

	return p
}

func (s *ExpressionstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionstmtContext) Rvalue() IRvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *ExpressionstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExpressionstmt(s)
	}
}

func (s *ExpressionstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExpressionstmt(s)
	}
}

func (p *bParser) Expressionstmt() (localctx IExpressionstmtContext) {
	localctx = NewExpressionstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, bParserRULE_expressionstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Rvalue()
	}
	{
		p.SetState(169)
		p.Match(bParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockstmtContext is an interface to support dynamic dispatch.
type IBlockstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockstmtContext differentiates from other interfaces.
	IsBlockstmtContext()
}

type BlockstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockstmtContext() *BlockstmtContext {
	var p = new(BlockstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_blockstmt
	return p
}

func InitEmptyBlockstmtContext(p *BlockstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_blockstmt
}

func (*BlockstmtContext) IsBlockstmtContext() {}

func NewBlockstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockstmtContext {
	var p = new(BlockstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_blockstmt

	return p
}

func (s *BlockstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockstmtContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockstmtContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterBlockstmt(s)
	}
}

func (s *BlockstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitBlockstmt(s)
	}
}

func (p *bParser) Blockstmt() (localctx IBlockstmtContext) {
	localctx = NewBlockstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, bParserRULE_blockstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(bParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32985413253722) != 0 {
		{
			p.SetState(172)
			p.Statement()
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(178)
		p.Match(bParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnstmtContext is an interface to support dynamic dispatch.
type IReturnstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Rvalue() IRvalueContext

	// IsReturnstmtContext differentiates from other interfaces.
	IsReturnstmtContext()
}

type ReturnstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnstmtContext() *ReturnstmtContext {
	var p = new(ReturnstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_returnstmt
	return p
}

func InitEmptyReturnstmtContext(p *ReturnstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_returnstmt
}

func (*ReturnstmtContext) IsReturnstmtContext() {}

func NewReturnstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnstmtContext {
	var p = new(ReturnstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_returnstmt

	return p
}

func (s *ReturnstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnstmtContext) Rvalue() IRvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *ReturnstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterReturnstmt(s)
	}
}

func (s *ReturnstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitReturnstmt(s)
	}
}

func (p *bParser) Returnstmt() (localctx IReturnstmtContext) {
	localctx = NewReturnstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, bParserRULE_returnstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(bParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == bParserT__3 {
		{
			p.SetState(181)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Rvalue()
		}
		{
			p.SetState(183)
			p.Match(bParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(187)
		p.Match(bParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGotostmtContext is an interface to support dynamic dispatch.
type IGotostmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Rvalue() IRvalueContext

	// IsGotostmtContext differentiates from other interfaces.
	IsGotostmtContext()
}

type GotostmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotostmtContext() *GotostmtContext {
	var p = new(GotostmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_gotostmt
	return p
}

func InitEmptyGotostmtContext(p *GotostmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_gotostmt
}

func (*GotostmtContext) IsGotostmtContext() {}

func NewGotostmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotostmtContext {
	var p = new(GotostmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_gotostmt

	return p
}

func (s *GotostmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GotostmtContext) Rvalue() IRvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *GotostmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotostmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotostmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterGotostmt(s)
	}
}

func (s *GotostmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitGotostmt(s)
	}
}

func (p *bParser) Gotostmt() (localctx IGotostmtContext) {
	localctx = NewGotostmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, bParserRULE_gotostmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(bParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Rvalue()
	}
	{
		p.SetState(191)
		p.Match(bParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchstmtContext is an interface to support dynamic dispatch.
type ISwitchstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Rvalue() IRvalueContext
	Statement() IStatementContext

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchstmtContext() *SwitchstmtContext {
	var p = new(SwitchstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_switchstmt
	return p
}

func InitEmptySwitchstmtContext(p *SwitchstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_switchstmt
}

func (*SwitchstmtContext) IsSwitchstmtContext() {}

func NewSwitchstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstmtContext {
	var p = new(SwitchstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_switchstmt

	return p
}

func (s *SwitchstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstmtContext) Rvalue() IRvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *SwitchstmtContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SwitchstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitSwitchstmt(s)
	}
}

func (p *bParser) Switchstmt() (localctx ISwitchstmtContext) {
	localctx = NewSwitchstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, bParserRULE_switchstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(bParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Rvalue()
	}
	{
		p.SetState(195)
		p.Statement()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Rvalue() IRvalueContext
	Statement() IStatementContext

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) Rvalue() IRvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *WhilestmtContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (p *bParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, bParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(bParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Match(bParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Rvalue()
	}
	{
		p.SetState(200)
		p.Match(bParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Statement()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Rvalue() IRvalueContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) Rvalue() IRvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *IfstmtContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfstmtContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *bParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, bParserRULE_ifstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(bParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(bParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Rvalue()
	}
	{
		p.SetState(206)
		p.Match(bParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Statement()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(208)
			p.Match(bParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Statement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICasestmtContext is an interface to support dynamic dispatch.
type ICasestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constant() IConstantContext
	Statement() IStatementContext

	// IsCasestmtContext differentiates from other interfaces.
	IsCasestmtContext()
}

type CasestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasestmtContext() *CasestmtContext {
	var p = new(CasestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_casestmt
	return p
}

func InitEmptyCasestmtContext(p *CasestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_casestmt
}

func (*CasestmtContext) IsCasestmtContext() {}

func NewCasestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasestmtContext {
	var p = new(CasestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_casestmt

	return p
}

func (s *CasestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CasestmtContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *CasestmtContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CasestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterCasestmt(s)
	}
}

func (s *CasestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitCasestmt(s)
	}
}

func (p *bParser) Casestmt() (localctx ICasestmtContext) {
	localctx = NewCasestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, bParserRULE_casestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(bParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Constant()
	}
	{
		p.SetState(214)
		p.Match(bParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Statement()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExternsmtContext is an interface to support dynamic dispatch.
type IExternsmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllName() []INameContext
	Name(i int) INameContext

	// IsExternsmtContext differentiates from other interfaces.
	IsExternsmtContext()
}

type ExternsmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternsmtContext() *ExternsmtContext {
	var p = new(ExternsmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_externsmt
	return p
}

func InitEmptyExternsmtContext(p *ExternsmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_externsmt
}

func (*ExternsmtContext) IsExternsmtContext() {}

func NewExternsmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternsmtContext {
	var p = new(ExternsmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_externsmt

	return p
}

func (s *ExternsmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternsmtContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *ExternsmtContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExternsmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternsmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternsmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExternsmt(s)
	}
}

func (s *ExternsmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExternsmt(s)
	}
}

func (p *bParser) Externsmt() (localctx IExternsmtContext) {
	localctx = NewExternsmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, bParserRULE_externsmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(bParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Name()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__1 {
		{
			p.SetState(219)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.Name()
		}

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(226)
		p.Match(bParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAutosmtContext is an interface to support dynamic dispatch.
type IAutosmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllName() []INameContext
	Name(i int) INameContext
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext

	// IsAutosmtContext differentiates from other interfaces.
	IsAutosmtContext()
}

type AutosmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAutosmtContext() *AutosmtContext {
	var p = new(AutosmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_autosmt
	return p
}

func InitEmptyAutosmtContext(p *AutosmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_autosmt
}

func (*AutosmtContext) IsAutosmtContext() {}

func NewAutosmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AutosmtContext {
	var p = new(AutosmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_autosmt

	return p
}

func (s *AutosmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AutosmtContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *AutosmtContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AutosmtContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *AutosmtContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *AutosmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AutosmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AutosmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterAutosmt(s)
	}
}

func (s *AutosmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitAutosmt(s)
	}
}

func (p *bParser) Autosmt() (localctx IAutosmtContext) {
	localctx = NewAutosmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, bParserRULE_autosmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(bParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Name()
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30786325577728) != 0 {
		{
			p.SetState(230)
			p.Constant()
		}

	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__1 {
		{
			p.SetState(233)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Name()
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30786325577728) != 0 {
			{
				p.SetState(235)
				p.Constant()
			}

		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(243)
		p.Match(bParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsmstmtContext is an interface to support dynamic dispatch.
type IAsmstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Stringlist() IStringlistContext

	// IsAsmstmtContext differentiates from other interfaces.
	IsAsmstmtContext()
}

type AsmstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsmstmtContext() *AsmstmtContext {
	var p = new(AsmstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_asmstmt
	return p
}

func InitEmptyAsmstmtContext(p *AsmstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_asmstmt
}

func (*AsmstmtContext) IsAsmstmtContext() {}

func NewAsmstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsmstmtContext {
	var p = new(AsmstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_asmstmt

	return p
}

func (s *AsmstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AsmstmtContext) Stringlist() IStringlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringlistContext)
}

func (s *AsmstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsmstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsmstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterAsmstmt(s)
	}
}

func (s *AsmstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitAsmstmt(s)
	}
}

func (p *bParser) Asmstmt() (localctx IAsmstmtContext) {
	localctx = NewAsmstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, bParserRULE_asmstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(bParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Match(bParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Stringlist()
	}
	{
		p.SetState(248)
		p.Match(bParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(bParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringlistContext is an interface to support dynamic dispatch.
type IStringlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING1() []antlr.TerminalNode
	STRING1(i int) antlr.TerminalNode

	// IsStringlistContext differentiates from other interfaces.
	IsStringlistContext()
}

type StringlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringlistContext() *StringlistContext {
	var p = new(StringlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_stringlist
	return p
}

func InitEmptyStringlistContext(p *StringlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_stringlist
}

func (*StringlistContext) IsStringlistContext() {}

func NewStringlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringlistContext {
	var p = new(StringlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_stringlist

	return p
}

func (s *StringlistContext) GetParser() antlr.Parser { return s.parser }

func (s *StringlistContext) AllSTRING1() []antlr.TerminalNode {
	return s.GetTokens(bParserSTRING1)
}

func (s *StringlistContext) STRING1(i int) antlr.TerminalNode {
	return s.GetToken(bParserSTRING1, i)
}

func (s *StringlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterStringlist(s)
	}
}

func (s *StringlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitStringlist(s)
	}
}

func (p *bParser) Stringlist() (localctx IStringlistContext) {
	localctx = NewStringlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, bParserRULE_stringlist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(bParserSTRING1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__1 {
		{
			p.SetState(252)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.Match(bParserSTRING1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRvalueContext is an interface to support dynamic dispatch.
type IRvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Comparison() IComparisonContext
	Ternary() ITernaryContext
	Assignment() IAssignmentContext

	// IsRvalueContext differentiates from other interfaces.
	IsRvalueContext()
}

type RvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRvalueContext() *RvalueContext {
	var p = new(RvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_rvalue
	return p
}

func InitEmptyRvalueContext(p *RvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_rvalue
}

func (*RvalueContext) IsRvalueContext() {}

func NewRvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RvalueContext {
	var p = new(RvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_rvalue

	return p
}

func (s *RvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *RvalueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RvalueContext) Comparison() IComparisonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *RvalueContext) Ternary() ITernaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITernaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITernaryContext)
}

func (s *RvalueContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *RvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterRvalue(s)
	}
}

func (s *RvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitRvalue(s)
	}
}

func (p *bParser) Rvalue() (localctx IRvalueContext) {
	localctx = NewRvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, bParserRULE_rvalue)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Comparison()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(261)
			p.Ternary()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(262)
			p.Assignment()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITernaryContext is an interface to support dynamic dispatch.
type ITernaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AllRvalue() []IRvalueContext
	Rvalue(i int) IRvalueContext

	// IsTernaryContext differentiates from other interfaces.
	IsTernaryContext()
}

type TernaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTernaryContext() *TernaryContext {
	var p = new(TernaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_ternary
	return p
}

func InitEmptyTernaryContext(p *TernaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_ternary
}

func (*TernaryContext) IsTernaryContext() {}

func NewTernaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TernaryContext {
	var p = new(TernaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_ternary

	return p
}

func (s *TernaryContext) GetParser() antlr.Parser { return s.parser }

func (s *TernaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TernaryContext) AllRvalue() []IRvalueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRvalueContext); ok {
			len++
		}
	}

	tst := make([]IRvalueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRvalueContext); ok {
			tst[i] = t.(IRvalueContext)
			i++
		}
	}

	return tst
}

func (s *TernaryContext) Rvalue(i int) IRvalueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *TernaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TernaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterTernary(s)
	}
}

func (s *TernaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitTernary(s)
	}
}

func (p *bParser) Ternary() (localctx ITernaryContext) {
	localctx = NewTernaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, bParserRULE_ternary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Expression()
	}
	{
		p.SetState(266)
		p.Match(bParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Rvalue()
	}
	{
		p.SetState(268)
		p.Match(bParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Rvalue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Binary() IBinaryContext
	Rvalue() IRvalueContext

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_comparison
	return p
}

func InitEmptyComparisonContext(p *ComparisonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_comparison
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparisonContext) Binary() IBinaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryContext)
}

func (s *ComparisonContext) Rvalue() IRvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (p *bParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, bParserRULE_comparison)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Expression()
	}
	{
		p.SetState(272)
		p.Binary()
	}
	{
		p.SetState(273)
		p.Rvalue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Assign() IAssignContext
	Rvalue() IRvalueContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AssignmentContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *AssignmentContext) Rvalue() IRvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *bParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, bParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Name()
	}
	{
		p.SetState(276)
		p.Assign()
	}
	{
		p.SetState(277)
		p.Rvalue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Rvalue() IRvalueContext
	Name() INameContext
	Constant() IConstantContext
	Incdec() IIncdecContext
	Unary() IUnaryContext
	Functioninvocation() IFunctioninvocationContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rvalue() IRvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *ExpressionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExpressionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionContext) Incdec() IIncdecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncdecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncdecContext)
}

func (s *ExpressionContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *ExpressionContext) Functioninvocation() IFunctioninvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctioninvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctioninvocationContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *bParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, bParserRULE_expression)
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(279)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Rvalue()
		}
		{
			p.SetState(281)
			p.Match(bParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.Name()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(284)
			p.Constant()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(285)
			p.Incdec()
		}
		{
			p.SetState(286)
			p.Name()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(288)
			p.Name()
		}
		{
			p.SetState(289)
			p.Incdec()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(291)
			p.Unary()
		}
		{
			p.SetState(292)
			p.Rvalue()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(294)
			p.Match(bParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.Name()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(296)
			p.Functioninvocation()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctioninvocationContext is an interface to support dynamic dispatch.
type IFunctioninvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Functionparameters() IFunctionparametersContext

	// IsFunctioninvocationContext differentiates from other interfaces.
	IsFunctioninvocationContext()
}

type FunctioninvocationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctioninvocationContext() *FunctioninvocationContext {
	var p = new(FunctioninvocationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_functioninvocation
	return p
}

func InitEmptyFunctioninvocationContext(p *FunctioninvocationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_functioninvocation
}

func (*FunctioninvocationContext) IsFunctioninvocationContext() {}

func NewFunctioninvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctioninvocationContext {
	var p = new(FunctioninvocationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_functioninvocation

	return p
}

func (s *FunctioninvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctioninvocationContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FunctioninvocationContext) Functionparameters() IFunctionparametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionparametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionparametersContext)
}

func (s *FunctioninvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctioninvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctioninvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterFunctioninvocation(s)
	}
}

func (s *FunctioninvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitFunctioninvocation(s)
	}
}

func (p *bParser) Functioninvocation() (localctx IFunctioninvocationContext) {
	localctx = NewFunctioninvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, bParserRULE_functioninvocation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Name()
	}
	{
		p.SetState(300)
		p.Match(bParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32985412796432) != 0 {
		{
			p.SetState(301)
			p.Functionparameters()
		}

	}
	{
		p.SetState(304)
		p.Match(bParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionparametersContext is an interface to support dynamic dispatch.
type IFunctionparametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRvalue() []IRvalueContext
	Rvalue(i int) IRvalueContext

	// IsFunctionparametersContext differentiates from other interfaces.
	IsFunctionparametersContext()
}

type FunctionparametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionparametersContext() *FunctionparametersContext {
	var p = new(FunctionparametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_functionparameters
	return p
}

func InitEmptyFunctionparametersContext(p *FunctionparametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_functionparameters
}

func (*FunctionparametersContext) IsFunctionparametersContext() {}

func NewFunctionparametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionparametersContext {
	var p = new(FunctionparametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_functionparameters

	return p
}

func (s *FunctionparametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionparametersContext) AllRvalue() []IRvalueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRvalueContext); ok {
			len++
		}
	}

	tst := make([]IRvalueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRvalueContext); ok {
			tst[i] = t.(IRvalueContext)
			i++
		}
	}

	return tst
}

func (s *FunctionparametersContext) Rvalue(i int) IRvalueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *FunctionparametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionparametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionparametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterFunctionparameters(s)
	}
}

func (s *FunctionparametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitFunctionparameters(s)
	}
}

func (p *bParser) Functionparameters() (localctx IFunctionparametersContext) {
	localctx = NewFunctionparametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, bParserRULE_functionparameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Rvalue()
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__1 {
		{
			p.SetState(307)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.Rvalue()
		}

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Binary() IBinaryContext

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) Binary() IBinaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *bParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, bParserRULE_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(bParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(315)
			p.Binary()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncdecContext is an interface to support dynamic dispatch.
type IIncdecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIncdecContext differentiates from other interfaces.
	IsIncdecContext()
}

type IncdecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncdecContext() *IncdecContext {
	var p = new(IncdecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_incdec
	return p
}

func InitEmptyIncdecContext(p *IncdecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_incdec
}

func (*IncdecContext) IsIncdecContext() {}

func NewIncdecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncdecContext {
	var p = new(IncdecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_incdec

	return p
}

func (s *IncdecContext) GetParser() antlr.Parser { return s.parser }
func (s *IncdecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncdecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncdecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterIncdec(s)
	}
}

func (s *IncdecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitIncdec(s)
	}
}

func (p *bParser) Incdec() (localctx IIncdecContext) {
	localctx = NewIncdecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, bParserRULE_incdec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		_la = p.GetTokenStream().LA(1)

		if !(_la == bParserT__21 || _la == bParserT__22) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_unary
	return p
}

func InitEmptyUnaryContext(p *UnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_unary
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (p *bParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, bParserRULE_unary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		_la = p.GetTokenStream().LA(1)

		if !(_la == bParserT__23 || _la == bParserT__24) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryContext is an interface to support dynamic dispatch.
type IBinaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBinaryContext differentiates from other interfaces.
	IsBinaryContext()
}

type BinaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryContext() *BinaryContext {
	var p = new(BinaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_binary
	return p
}

func InitEmptyBinaryContext(p *BinaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_binary
}

func (*BinaryContext) IsBinaryContext() {}

func NewBinaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryContext {
	var p = new(BinaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_binary

	return p
}

func (s *BinaryContext) GetParser() antlr.Parser { return s.parser }
func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (p *bParser) Binary() (localctx IBinaryContext) {
	localctx = NewBinaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, bParserRULE_binary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549706530816) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	AllRvalue() []IRvalueContext
	Rvalue(i int) IRvalueContext

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_lvalue
	return p
}

func InitEmptyLvalueContext(p *LvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_lvalue
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LvalueContext) AllRvalue() []IRvalueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRvalueContext); ok {
			len++
		}
	}

	tst := make([]IRvalueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRvalueContext); ok {
			tst[i] = t.(IRvalueContext)
			i++
		}
	}

	return tst
}

func (s *LvalueContext) Rvalue(i int) IRvalueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (p *bParser) Lvalue() (localctx ILvalueContext) {
	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, bParserRULE_lvalue)
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(324)
			p.Name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(325)
			p.Match(bParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.Rvalue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(327)
			p.Rvalue()
		}
		{
			p.SetState(328)
			p.Match(bParserT__38)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Rvalue()
		}
		{
			p.SetState(330)
			p.Match(bParserT__39)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	STRING1() antlr.TerminalNode
	STRING2() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) INT() antlr.TerminalNode {
	return s.GetToken(bParserINT, 0)
}

func (s *ConstantContext) STRING1() antlr.TerminalNode {
	return s.GetToken(bParserSTRING1, 0)
}

func (s *ConstantContext) STRING2() antlr.TerminalNode {
	return s.GetToken(bParserSTRING2, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *bParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, bParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30786325577728) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(bParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *bParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, bParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		p.Match(bParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

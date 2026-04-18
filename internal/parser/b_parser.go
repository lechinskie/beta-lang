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
		"", "';'", "'('", "','", "')'", "'['", "']'", "':'", "'{'", "'}'", "'*'",
		"'&'", "'-'", "'!'", "'~'", "'/'", "'%'", "'+'", "'<'", "'>'", "'^'",
		"'|'", "'?'", "'auto'", "'break'", "'case'", "'default'", "'else'",
		"'extrn'", "'goto'", "'if'", "'return'", "'switch'", "'while'", "'__asm__'",
		"'__variadic__'", "'++'", "'--'", "'<<'", "'>>'", "'<='", "'>='", "'=='",
		"'!='", "'=*'", "'=/'", "'=%'", "'=+'", "'=-'", "'=<<'", "'=>>'", "'=<='",
		"'=<'", "'=>='", "'=>'", "'==='", "'=!='", "'=&'", "'=^'", "'=|'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "AUTO", "BREAK", "CASE", "DEFAULT", "ELSE",
		"EXTRN", "GOTO", "IF", "RETURN", "SWITCH", "WHILE", "ASM", "VARIADIC",
		"INC", "DEC", "SHL", "SHR", "LE", "GE", "EQ", "NE", "ASS_MUL", "ASS_DIV",
		"ASS_MOD", "ASS_ADD", "ASS_SUB", "ASS_SHL", "ASS_SHR", "ASS_LE", "ASS_LT",
		"ASS_GE", "ASS_GT", "ASS_EQ", "ASS_NE", "ASS_AND", "ASS_XOR", "ASS_OR",
		"ASSIGN", "ID", "DECIMAL", "OCTAL", "CHAR", "STRING", "LINECOMMENT",
		"COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "ext_def", "ival", "ival_list", "arg_list", "name_list",
		"string_list", "statement", "compound_stmt", "auto_decl", "auto_def",
		"extrn_decl", "expr", "expr_list", "assign_op",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 68, 285, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 4, 0,
		32, 8, 0, 11, 0, 12, 0, 33, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 60, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		3, 1, 67, 8, 1, 1, 1, 1, 1, 3, 1, 71, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 80, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 87,
		8, 3, 10, 3, 12, 3, 90, 9, 3, 1, 4, 1, 4, 1, 4, 5, 4, 95, 8, 4, 10, 4,
		12, 4, 98, 9, 4, 1, 5, 1, 5, 1, 5, 5, 5, 103, 8, 5, 10, 5, 12, 5, 106,
		9, 5, 1, 6, 1, 6, 1, 6, 5, 6, 111, 8, 6, 10, 6, 12, 6, 114, 9, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 124, 8, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		175, 8, 7, 1, 8, 1, 8, 5, 8, 179, 8, 8, 10, 8, 12, 8, 182, 9, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 190, 8, 9, 10, 9, 12, 9, 193, 9, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 203, 8, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 221, 8, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 265,
		8, 12, 1, 12, 1, 12, 1, 12, 5, 12, 270, 8, 12, 10, 12, 12, 12, 273, 9,
		12, 1, 13, 1, 13, 1, 13, 5, 13, 278, 8, 13, 10, 13, 12, 13, 281, 9, 13,
		1, 14, 1, 14, 1, 14, 0, 1, 24, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 0, 8, 2, 0, 10, 14, 36, 37, 2, 0, 10, 10, 15, 16, 2, 0,
		12, 12, 17, 17, 1, 0, 38, 39, 2, 0, 18, 19, 40, 41, 1, 0, 42, 43, 1, 0,
		36, 37, 1, 0, 44, 60, 323, 0, 31, 1, 0, 0, 0, 2, 79, 1, 0, 0, 0, 4, 81,
		1, 0, 0, 0, 6, 83, 1, 0, 0, 0, 8, 91, 1, 0, 0, 0, 10, 99, 1, 0, 0, 0, 12,
		107, 1, 0, 0, 0, 14, 174, 1, 0, 0, 0, 16, 176, 1, 0, 0, 0, 18, 185, 1,
		0, 0, 0, 20, 202, 1, 0, 0, 0, 22, 204, 1, 0, 0, 0, 24, 220, 1, 0, 0, 0,
		26, 274, 1, 0, 0, 0, 28, 282, 1, 0, 0, 0, 30, 32, 3, 2, 1, 0, 31, 30, 1,
		0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34,
		35, 1, 0, 0, 0, 35, 36, 5, 0, 0, 1, 36, 1, 1, 0, 0, 0, 37, 38, 5, 28, 0,
		0, 38, 39, 3, 10, 5, 0, 39, 40, 5, 1, 0, 0, 40, 80, 1, 0, 0, 0, 41, 42,
		5, 35, 0, 0, 42, 43, 5, 2, 0, 0, 43, 44, 5, 61, 0, 0, 44, 45, 5, 3, 0,
		0, 45, 46, 3, 24, 12, 0, 46, 47, 5, 4, 0, 0, 47, 48, 5, 1, 0, 0, 48, 80,
		1, 0, 0, 0, 49, 50, 5, 61, 0, 0, 50, 51, 5, 34, 0, 0, 51, 52, 5, 2, 0,
		0, 52, 53, 3, 12, 6, 0, 53, 54, 5, 4, 0, 0, 54, 55, 5, 1, 0, 0, 55, 80,
		1, 0, 0, 0, 56, 57, 5, 61, 0, 0, 57, 59, 5, 2, 0, 0, 58, 60, 3, 8, 4, 0,
		59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 5,
		4, 0, 0, 62, 80, 3, 14, 7, 0, 63, 64, 5, 61, 0, 0, 64, 66, 5, 5, 0, 0,
		65, 67, 3, 24, 12, 0, 66, 65, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 1,
		0, 0, 0, 68, 70, 5, 6, 0, 0, 69, 71, 3, 6, 3, 0, 70, 69, 1, 0, 0, 0, 70,
		71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 80, 5, 1, 0, 0, 73, 74, 5, 61,
		0, 0, 74, 75, 3, 4, 2, 0, 75, 76, 5, 1, 0, 0, 76, 80, 1, 0, 0, 0, 77, 78,
		5, 61, 0, 0, 78, 80, 5, 1, 0, 0, 79, 37, 1, 0, 0, 0, 79, 41, 1, 0, 0, 0,
		79, 49, 1, 0, 0, 0, 79, 56, 1, 0, 0, 0, 79, 63, 1, 0, 0, 0, 79, 73, 1,
		0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 3, 1, 0, 0, 0, 81, 82, 3, 24, 12, 0, 82,
		5, 1, 0, 0, 0, 83, 88, 3, 4, 2, 0, 84, 85, 5, 3, 0, 0, 85, 87, 3, 4, 2,
		0, 86, 84, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89,
		1, 0, 0, 0, 89, 7, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 96, 5, 61, 0, 0,
		92, 93, 5, 3, 0, 0, 93, 95, 5, 61, 0, 0, 94, 92, 1, 0, 0, 0, 95, 98, 1,
		0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 9, 1, 0, 0, 0, 98,
		96, 1, 0, 0, 0, 99, 104, 5, 61, 0, 0, 100, 101, 5, 3, 0, 0, 101, 103, 5,
		61, 0, 0, 102, 100, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102, 1, 0, 0,
		0, 104, 105, 1, 0, 0, 0, 105, 11, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 107,
		112, 5, 65, 0, 0, 108, 109, 5, 3, 0, 0, 109, 111, 5, 65, 0, 0, 110, 108,
		1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0,
		0, 0, 113, 13, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 175, 3, 16, 8, 0,
		116, 117, 5, 30, 0, 0, 117, 118, 5, 2, 0, 0, 118, 119, 3, 24, 12, 0, 119,
		120, 5, 4, 0, 0, 120, 123, 3, 14, 7, 0, 121, 122, 5, 27, 0, 0, 122, 124,
		3, 14, 7, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 175, 1, 0,
		0, 0, 125, 126, 5, 33, 0, 0, 126, 127, 5, 2, 0, 0, 127, 128, 3, 24, 12,
		0, 128, 129, 5, 4, 0, 0, 129, 130, 3, 14, 7, 0, 130, 175, 1, 0, 0, 0, 131,
		132, 5, 32, 0, 0, 132, 133, 5, 2, 0, 0, 133, 134, 3, 24, 12, 0, 134, 135,
		5, 4, 0, 0, 135, 136, 3, 14, 7, 0, 136, 175, 1, 0, 0, 0, 137, 138, 5, 25,
		0, 0, 138, 139, 3, 24, 12, 0, 139, 140, 5, 7, 0, 0, 140, 141, 3, 14, 7,
		0, 141, 175, 1, 0, 0, 0, 142, 143, 5, 26, 0, 0, 143, 144, 5, 7, 0, 0, 144,
		175, 3, 14, 7, 0, 145, 146, 5, 61, 0, 0, 146, 147, 5, 7, 0, 0, 147, 175,
		3, 14, 7, 0, 148, 149, 5, 29, 0, 0, 149, 150, 3, 24, 12, 0, 150, 151, 5,
		1, 0, 0, 151, 175, 1, 0, 0, 0, 152, 153, 5, 24, 0, 0, 153, 175, 5, 1, 0,
		0, 154, 155, 5, 31, 0, 0, 155, 175, 5, 1, 0, 0, 156, 157, 5, 31, 0, 0,
		157, 158, 5, 2, 0, 0, 158, 159, 3, 24, 12, 0, 159, 160, 5, 4, 0, 0, 160,
		161, 5, 1, 0, 0, 161, 175, 1, 0, 0, 0, 162, 163, 5, 34, 0, 0, 163, 164,
		5, 2, 0, 0, 164, 165, 3, 12, 6, 0, 165, 166, 5, 4, 0, 0, 166, 167, 5, 1,
		0, 0, 167, 175, 1, 0, 0, 0, 168, 175, 3, 18, 9, 0, 169, 175, 3, 22, 11,
		0, 170, 171, 3, 24, 12, 0, 171, 172, 5, 1, 0, 0, 172, 175, 1, 0, 0, 0,
		173, 175, 5, 1, 0, 0, 174, 115, 1, 0, 0, 0, 174, 116, 1, 0, 0, 0, 174,
		125, 1, 0, 0, 0, 174, 131, 1, 0, 0, 0, 174, 137, 1, 0, 0, 0, 174, 142,
		1, 0, 0, 0, 174, 145, 1, 0, 0, 0, 174, 148, 1, 0, 0, 0, 174, 152, 1, 0,
		0, 0, 174, 154, 1, 0, 0, 0, 174, 156, 1, 0, 0, 0, 174, 162, 1, 0, 0, 0,
		174, 168, 1, 0, 0, 0, 174, 169, 1, 0, 0, 0, 174, 170, 1, 0, 0, 0, 174,
		173, 1, 0, 0, 0, 175, 15, 1, 0, 0, 0, 176, 180, 5, 8, 0, 0, 177, 179, 3,
		14, 7, 0, 178, 177, 1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0,
		0, 180, 181, 1, 0, 0, 0, 181, 183, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183,
		184, 5, 9, 0, 0, 184, 17, 1, 0, 0, 0, 185, 186, 5, 23, 0, 0, 186, 191,
		3, 20, 10, 0, 187, 188, 5, 3, 0, 0, 188, 190, 3, 20, 10, 0, 189, 187, 1,
		0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0,
		0, 192, 194, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 195, 5, 1, 0, 0, 195,
		19, 1, 0, 0, 0, 196, 203, 5, 61, 0, 0, 197, 198, 5, 61, 0, 0, 198, 199,
		5, 5, 0, 0, 199, 200, 3, 24, 12, 0, 200, 201, 5, 6, 0, 0, 201, 203, 1,
		0, 0, 0, 202, 196, 1, 0, 0, 0, 202, 197, 1, 0, 0, 0, 203, 21, 1, 0, 0,
		0, 204, 205, 5, 28, 0, 0, 205, 206, 3, 10, 5, 0, 206, 207, 5, 1, 0, 0,
		207, 23, 1, 0, 0, 0, 208, 209, 6, 12, -1, 0, 209, 210, 7, 0, 0, 0, 210,
		221, 3, 24, 12, 17, 211, 212, 5, 2, 0, 0, 212, 213, 3, 24, 12, 0, 213,
		214, 5, 4, 0, 0, 214, 221, 1, 0, 0, 0, 215, 221, 5, 61, 0, 0, 216, 221,
		5, 62, 0, 0, 217, 221, 5, 63, 0, 0, 218, 221, 5, 64, 0, 0, 219, 221, 5,
		65, 0, 0, 220, 208, 1, 0, 0, 0, 220, 211, 1, 0, 0, 0, 220, 215, 1, 0, 0,
		0, 220, 216, 1, 0, 0, 0, 220, 217, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 220,
		219, 1, 0, 0, 0, 221, 271, 1, 0, 0, 0, 222, 223, 10, 16, 0, 0, 223, 224,
		7, 1, 0, 0, 224, 270, 3, 24, 12, 17, 225, 226, 10, 15, 0, 0, 226, 227,
		7, 2, 0, 0, 227, 270, 3, 24, 12, 16, 228, 229, 10, 14, 0, 0, 229, 230,
		7, 3, 0, 0, 230, 270, 3, 24, 12, 15, 231, 232, 10, 13, 0, 0, 232, 233,
		7, 4, 0, 0, 233, 270, 3, 24, 12, 14, 234, 235, 10, 12, 0, 0, 235, 236,
		7, 5, 0, 0, 236, 270, 3, 24, 12, 13, 237, 238, 10, 11, 0, 0, 238, 239,
		5, 11, 0, 0, 239, 270, 3, 24, 12, 12, 240, 241, 10, 10, 0, 0, 241, 242,
		5, 20, 0, 0, 242, 270, 3, 24, 12, 11, 243, 244, 10, 9, 0, 0, 244, 245,
		5, 21, 0, 0, 245, 270, 3, 24, 12, 10, 246, 247, 10, 8, 0, 0, 247, 248,
		5, 22, 0, 0, 248, 249, 3, 24, 12, 0, 249, 250, 5, 7, 0, 0, 250, 251, 3,
		24, 12, 8, 251, 270, 1, 0, 0, 0, 252, 253, 10, 7, 0, 0, 253, 254, 3, 28,
		14, 0, 254, 255, 3, 24, 12, 7, 255, 270, 1, 0, 0, 0, 256, 257, 10, 20,
		0, 0, 257, 258, 5, 5, 0, 0, 258, 259, 3, 24, 12, 0, 259, 260, 5, 6, 0,
		0, 260, 270, 1, 0, 0, 0, 261, 262, 10, 19, 0, 0, 262, 264, 5, 2, 0, 0,
		263, 265, 3, 26, 13, 0, 264, 263, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265,
		266, 1, 0, 0, 0, 266, 270, 5, 4, 0, 0, 267, 268, 10, 18, 0, 0, 268, 270,
		7, 6, 0, 0, 269, 222, 1, 0, 0, 0, 269, 225, 1, 0, 0, 0, 269, 228, 1, 0,
		0, 0, 269, 231, 1, 0, 0, 0, 269, 234, 1, 0, 0, 0, 269, 237, 1, 0, 0, 0,
		269, 240, 1, 0, 0, 0, 269, 243, 1, 0, 0, 0, 269, 246, 1, 0, 0, 0, 269,
		252, 1, 0, 0, 0, 269, 256, 1, 0, 0, 0, 269, 261, 1, 0, 0, 0, 269, 267,
		1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0,
		0, 0, 272, 25, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 279, 3, 24, 12, 0,
		275, 276, 5, 3, 0, 0, 276, 278, 3, 24, 12, 0, 277, 275, 1, 0, 0, 0, 278,
		281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 27, 1,
		0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 283, 7, 7, 0, 0, 283, 29, 1, 0, 0,
		0, 19, 33, 59, 66, 70, 79, 88, 96, 104, 112, 123, 174, 180, 191, 202, 220,
		264, 269, 271, 279,
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
	bParserEOF         = antlr.TokenEOF
	bParserT__0        = 1
	bParserT__1        = 2
	bParserT__2        = 3
	bParserT__3        = 4
	bParserT__4        = 5
	bParserT__5        = 6
	bParserT__6        = 7
	bParserT__7        = 8
	bParserT__8        = 9
	bParserT__9        = 10
	bParserT__10       = 11
	bParserT__11       = 12
	bParserT__12       = 13
	bParserT__13       = 14
	bParserT__14       = 15
	bParserT__15       = 16
	bParserT__16       = 17
	bParserT__17       = 18
	bParserT__18       = 19
	bParserT__19       = 20
	bParserT__20       = 21
	bParserT__21       = 22
	bParserAUTO        = 23
	bParserBREAK       = 24
	bParserCASE        = 25
	bParserDEFAULT     = 26
	bParserELSE        = 27
	bParserEXTRN       = 28
	bParserGOTO        = 29
	bParserIF          = 30
	bParserRETURN      = 31
	bParserSWITCH      = 32
	bParserWHILE       = 33
	bParserASM         = 34
	bParserVARIADIC    = 35
	bParserINC         = 36
	bParserDEC         = 37
	bParserSHL         = 38
	bParserSHR         = 39
	bParserLE          = 40
	bParserGE          = 41
	bParserEQ          = 42
	bParserNE          = 43
	bParserASS_MUL     = 44
	bParserASS_DIV     = 45
	bParserASS_MOD     = 46
	bParserASS_ADD     = 47
	bParserASS_SUB     = 48
	bParserASS_SHL     = 49
	bParserASS_SHR     = 50
	bParserASS_LE      = 51
	bParserASS_LT      = 52
	bParserASS_GE      = 53
	bParserASS_GT      = 54
	bParserASS_EQ      = 55
	bParserASS_NE      = 56
	bParserASS_AND     = 57
	bParserASS_XOR     = 58
	bParserASS_OR      = 59
	bParserASSIGN      = 60
	bParserID          = 61
	bParserDECIMAL     = 62
	bParserOCTAL       = 63
	bParserCHAR        = 64
	bParserSTRING      = 65
	bParserLINECOMMENT = 66
	bParserCOMMENT     = 67
	bParserWS          = 68
)

// bParser rules.
const (
	bParserRULE_program       = 0
	bParserRULE_ext_def       = 1
	bParserRULE_ival          = 2
	bParserRULE_ival_list     = 3
	bParserRULE_arg_list      = 4
	bParserRULE_name_list     = 5
	bParserRULE_string_list   = 6
	bParserRULE_statement     = 7
	bParserRULE_compound_stmt = 8
	bParserRULE_auto_decl     = 9
	bParserRULE_auto_def      = 10
	bParserRULE_extrn_decl    = 11
	bParserRULE_expr          = 12
	bParserRULE_expr_list     = 13
	bParserRULE_assign_op     = 14
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllExt_def() []IExt_defContext
	Ext_def(i int) IExt_defContext

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

func (s *ProgramContext) AllExt_def() []IExt_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExt_defContext); ok {
			len++
		}
	}

	tst := make([]IExt_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExt_defContext); ok {
			tst[i] = t.(IExt_defContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Ext_def(i int) IExt_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExt_defContext); ok {
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

	return t.(IExt_defContext)
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2305843043841867776) != 0) {
		{
			p.SetState(30)
			p.Ext_def()
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
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

// IExt_defContext is an interface to support dynamic dispatch.
type IExt_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTRN() antlr.TerminalNode
	Name_list() IName_listContext
	VARIADIC() antlr.TerminalNode
	ID() antlr.TerminalNode
	Expr() IExprContext
	ASM() antlr.TerminalNode
	String_list() IString_listContext
	Statement() IStatementContext
	Arg_list() IArg_listContext
	Ival_list() IIval_listContext
	Ival() IIvalContext

	// IsExt_defContext differentiates from other interfaces.
	IsExt_defContext()
}

type Ext_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExt_defContext() *Ext_defContext {
	var p = new(Ext_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_ext_def
	return p
}

func InitEmptyExt_defContext(p *Ext_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_ext_def
}

func (*Ext_defContext) IsExt_defContext() {}

func NewExt_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ext_defContext {
	var p = new(Ext_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_ext_def

	return p
}

func (s *Ext_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Ext_defContext) EXTRN() antlr.TerminalNode {
	return s.GetToken(bParserEXTRN, 0)
}

func (s *Ext_defContext) Name_list() IName_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IName_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IName_listContext)
}

func (s *Ext_defContext) VARIADIC() antlr.TerminalNode {
	return s.GetToken(bParserVARIADIC, 0)
}

func (s *Ext_defContext) ID() antlr.TerminalNode {
	return s.GetToken(bParserID, 0)
}

func (s *Ext_defContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Ext_defContext) ASM() antlr.TerminalNode {
	return s.GetToken(bParserASM, 0)
}

func (s *Ext_defContext) String_list() IString_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_listContext)
}

func (s *Ext_defContext) Statement() IStatementContext {
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

func (s *Ext_defContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *Ext_defContext) Ival_list() IIval_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIval_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIval_listContext)
}

func (s *Ext_defContext) Ival() IIvalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIvalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIvalContext)
}

func (s *Ext_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ext_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExt_def(s)
	}
}

func (s *Ext_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExt_def(s)
	}
}

func (p *bParser) Ext_def() (localctx IExt_defContext) {
	localctx = NewExt_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, bParserRULE_ext_def)
	var _la int

	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(bParserEXTRN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Name_list()
		}
		{
			p.SetState(39)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.Match(bParserVARIADIC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(bParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.expr(0)
		}
		{
			p.SetState(46)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(bParserASM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.String_list()
		}
		{
			p.SetState(53)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == bParserID {
			{
				p.SetState(58)
				p.Arg_list()
			}

		}
		{
			p.SetState(61)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(bParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-576460700763807999) != 0 {
			{
				p.SetState(65)
				p.expr(0)
			}

		}
		{
			p.SetState(68)
			p.Match(bParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-576460700763807999) != 0 {
			{
				p.SetState(69)
				p.Ival_list()
			}

		}
		{
			p.SetState(72)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(73)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Ival()
		}
		{
			p.SetState(75)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(77)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(bParserT__0)
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

// IIvalContext is an interface to support dynamic dispatch.
type IIvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

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

func (s *IvalContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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
	p.EnterRule(localctx, 4, bParserRULE_ival)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.expr(0)
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

// IIval_listContext is an interface to support dynamic dispatch.
type IIval_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIval() []IIvalContext
	Ival(i int) IIvalContext

	// IsIval_listContext differentiates from other interfaces.
	IsIval_listContext()
}

type Ival_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIval_listContext() *Ival_listContext {
	var p = new(Ival_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_ival_list
	return p
}

func InitEmptyIval_listContext(p *Ival_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_ival_list
}

func (*Ival_listContext) IsIval_listContext() {}

func NewIval_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ival_listContext {
	var p = new(Ival_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_ival_list

	return p
}

func (s *Ival_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Ival_listContext) AllIval() []IIvalContext {
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

func (s *Ival_listContext) Ival(i int) IIvalContext {
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

func (s *Ival_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ival_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ival_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterIval_list(s)
	}
}

func (s *Ival_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitIval_list(s)
	}
}

func (p *bParser) Ival_list() (localctx IIval_listContext) {
	localctx = NewIval_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, bParserRULE_ival_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Ival()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__2 {
		{
			p.SetState(84)
			p.Match(bParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Ival()
		}

		p.SetState(90)
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

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_arg_list
	return p
}

func InitEmptyArg_listContext(p *Arg_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_arg_list
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(bParserID)
}

func (s *Arg_listContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(bParserID, i)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arg_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterArg_list(s)
	}
}

func (s *Arg_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitArg_list(s)
	}
}

func (p *bParser) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, bParserRULE_arg_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(bParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__2 {
		{
			p.SetState(92)
			p.Match(bParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(98)
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

// IName_listContext is an interface to support dynamic dispatch.
type IName_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsName_listContext differentiates from other interfaces.
	IsName_listContext()
}

type Name_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyName_listContext() *Name_listContext {
	var p = new(Name_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_name_list
	return p
}

func InitEmptyName_listContext(p *Name_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_name_list
}

func (*Name_listContext) IsName_listContext() {}

func NewName_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Name_listContext {
	var p = new(Name_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_name_list

	return p
}

func (s *Name_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Name_listContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(bParserID)
}

func (s *Name_listContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(bParserID, i)
}

func (s *Name_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Name_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Name_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterName_list(s)
	}
}

func (s *Name_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitName_list(s)
	}
}

func (p *bParser) Name_list() (localctx IName_listContext) {
	localctx = NewName_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, bParserRULE_name_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(bParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__2 {
		{
			p.SetState(100)
			p.Match(bParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(106)
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

// IString_listContext is an interface to support dynamic dispatch.
type IString_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode

	// IsString_listContext differentiates from other interfaces.
	IsString_listContext()
}

type String_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_listContext() *String_listContext {
	var p = new(String_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_string_list
	return p
}

func InitEmptyString_listContext(p *String_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_string_list
}

func (*String_listContext) IsString_listContext() {}

func NewString_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_listContext {
	var p = new(String_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_string_list

	return p
}

func (s *String_listContext) GetParser() antlr.Parser { return s.parser }

func (s *String_listContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(bParserSTRING)
}

func (s *String_listContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(bParserSTRING, i)
}

func (s *String_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterString_list(s)
	}
}

func (s *String_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitString_list(s)
	}
}

func (p *bParser) String_list() (localctx IString_listContext) {
	localctx = NewString_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, bParserRULE_string_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(bParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__2 {
		{
			p.SetState(108)
			p.Match(bParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(bParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(114)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Compound_stmt() ICompound_stmtContext
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	ELSE() antlr.TerminalNode
	WHILE() antlr.TerminalNode
	SWITCH() antlr.TerminalNode
	CASE() antlr.TerminalNode
	DEFAULT() antlr.TerminalNode
	ID() antlr.TerminalNode
	GOTO() antlr.TerminalNode
	BREAK() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	ASM() antlr.TerminalNode
	String_list() IString_listContext
	Auto_decl() IAuto_declContext
	Extrn_decl() IExtrn_declContext

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

func (s *StatementContext) Compound_stmt() ICompound_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompound_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompound_stmtContext)
}

func (s *StatementContext) IF() antlr.TerminalNode {
	return s.GetToken(bParserIF, 0)
}

func (s *StatementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatementContext) AllStatement() []IStatementContext {
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

func (s *StatementContext) Statement(i int) IStatementContext {
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

func (s *StatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(bParserELSE, 0)
}

func (s *StatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(bParserWHILE, 0)
}

func (s *StatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(bParserSWITCH, 0)
}

func (s *StatementContext) CASE() antlr.TerminalNode {
	return s.GetToken(bParserCASE, 0)
}

func (s *StatementContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(bParserDEFAULT, 0)
}

func (s *StatementContext) ID() antlr.TerminalNode {
	return s.GetToken(bParserID, 0)
}

func (s *StatementContext) GOTO() antlr.TerminalNode {
	return s.GetToken(bParserGOTO, 0)
}

func (s *StatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(bParserBREAK, 0)
}

func (s *StatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(bParserRETURN, 0)
}

func (s *StatementContext) ASM() antlr.TerminalNode {
	return s.GetToken(bParserASM, 0)
}

func (s *StatementContext) String_list() IString_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_listContext)
}

func (s *StatementContext) Auto_decl() IAuto_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuto_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuto_declContext)
}

func (s *StatementContext) Extrn_decl() IExtrn_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtrn_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtrn_declContext)
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
	p.EnterRule(localctx, 14, bParserRULE_statement)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Compound_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Match(bParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.expr(0)
		}
		{
			p.SetState(119)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Statement()
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(121)
				p.Match(bParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(122)
				p.Statement()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Match(bParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.expr(0)
		}
		{
			p.SetState(128)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.Match(bParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.expr(0)
		}
		{
			p.SetState(134)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(137)
			p.Match(bParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.expr(0)
		}
		{
			p.SetState(139)
			p.Match(bParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Statement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(142)
			p.Match(bParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(bParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Statement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(145)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(bParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Statement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(148)
			p.Match(bParserGOTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.expr(0)
		}
		{
			p.SetState(150)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(152)
			p.Match(bParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(154)
			p.Match(bParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(156)
			p.Match(bParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.expr(0)
		}
		{
			p.SetState(159)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(162)
			p.Match(bParserASM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.String_list()
		}
		{
			p.SetState(165)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(168)
			p.Auto_decl()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(169)
			p.Extrn_decl()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(170)
			p.expr(0)
		}
		{
			p.SetState(171)
			p.Match(bParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(173)
			p.Match(bParserT__0)
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

// ICompound_stmtContext is an interface to support dynamic dispatch.
type ICompound_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsCompound_stmtContext differentiates from other interfaces.
	IsCompound_stmtContext()
}

type Compound_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompound_stmtContext() *Compound_stmtContext {
	var p = new(Compound_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_compound_stmt
	return p
}

func InitEmptyCompound_stmtContext(p *Compound_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_compound_stmt
}

func (*Compound_stmtContext) IsCompound_stmtContext() {}

func NewCompound_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compound_stmtContext {
	var p = new(Compound_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_compound_stmt

	return p
}

func (s *Compound_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Compound_stmtContext) AllStatement() []IStatementContext {
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

func (s *Compound_stmtContext) Statement(i int) IStatementContext {
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

func (s *Compound_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compound_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Compound_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterCompound_stmt(s)
	}
}

func (s *Compound_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitCompound_stmt(s)
	}
}

func (p *bParser) Compound_stmt() (localctx ICompound_stmtContext) {
	localctx = NewCompound_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, bParserRULE_compound_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(bParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2305842768838099706) != 0) || _la == bParserCHAR || _la == bParserSTRING {
		{
			p.SetState(177)
			p.Statement()
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(183)
		p.Match(bParserT__8)
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

// IAuto_declContext is an interface to support dynamic dispatch.
type IAuto_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AUTO() antlr.TerminalNode
	AllAuto_def() []IAuto_defContext
	Auto_def(i int) IAuto_defContext

	// IsAuto_declContext differentiates from other interfaces.
	IsAuto_declContext()
}

type Auto_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuto_declContext() *Auto_declContext {
	var p = new(Auto_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_auto_decl
	return p
}

func InitEmptyAuto_declContext(p *Auto_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_auto_decl
}

func (*Auto_declContext) IsAuto_declContext() {}

func NewAuto_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Auto_declContext {
	var p = new(Auto_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_auto_decl

	return p
}

func (s *Auto_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Auto_declContext) AUTO() antlr.TerminalNode {
	return s.GetToken(bParserAUTO, 0)
}

func (s *Auto_declContext) AllAuto_def() []IAuto_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAuto_defContext); ok {
			len++
		}
	}

	tst := make([]IAuto_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAuto_defContext); ok {
			tst[i] = t.(IAuto_defContext)
			i++
		}
	}

	return tst
}

func (s *Auto_declContext) Auto_def(i int) IAuto_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuto_defContext); ok {
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

	return t.(IAuto_defContext)
}

func (s *Auto_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Auto_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Auto_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterAuto_decl(s)
	}
}

func (s *Auto_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitAuto_decl(s)
	}
}

func (p *bParser) Auto_decl() (localctx IAuto_declContext) {
	localctx = NewAuto_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, bParserRULE_auto_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(bParserAUTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Auto_def()
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__2 {
		{
			p.SetState(187)
			p.Match(bParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Auto_def()
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
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

// IAuto_defContext is an interface to support dynamic dispatch.
type IAuto_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsAuto_defContext differentiates from other interfaces.
	IsAuto_defContext()
}

type Auto_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuto_defContext() *Auto_defContext {
	var p = new(Auto_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_auto_def
	return p
}

func InitEmptyAuto_defContext(p *Auto_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_auto_def
}

func (*Auto_defContext) IsAuto_defContext() {}

func NewAuto_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Auto_defContext {
	var p = new(Auto_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_auto_def

	return p
}

func (s *Auto_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Auto_defContext) ID() antlr.TerminalNode {
	return s.GetToken(bParserID, 0)
}

func (s *Auto_defContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Auto_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Auto_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Auto_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterAuto_def(s)
	}
}

func (s *Auto_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitAuto_def(s)
	}
}

func (p *bParser) Auto_def() (localctx IAuto_defContext) {
	localctx = NewAuto_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, bParserRULE_auto_def)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(bParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.expr(0)
		}
		{
			p.SetState(200)
			p.Match(bParserT__5)
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

// IExtrn_declContext is an interface to support dynamic dispatch.
type IExtrn_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTRN() antlr.TerminalNode
	Name_list() IName_listContext

	// IsExtrn_declContext differentiates from other interfaces.
	IsExtrn_declContext()
}

type Extrn_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtrn_declContext() *Extrn_declContext {
	var p = new(Extrn_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_extrn_decl
	return p
}

func InitEmptyExtrn_declContext(p *Extrn_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_extrn_decl
}

func (*Extrn_declContext) IsExtrn_declContext() {}

func NewExtrn_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Extrn_declContext {
	var p = new(Extrn_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_extrn_decl

	return p
}

func (s *Extrn_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Extrn_declContext) EXTRN() antlr.TerminalNode {
	return s.GetToken(bParserEXTRN, 0)
}

func (s *Extrn_declContext) Name_list() IName_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IName_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IName_listContext)
}

func (s *Extrn_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Extrn_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Extrn_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExtrn_decl(s)
	}
}

func (s *Extrn_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExtrn_decl(s)
	}
}

func (p *bParser) Extrn_decl() (localctx IExtrn_declContext) {
	localctx = NewExtrn_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, bParserRULE_extrn_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(bParserEXTRN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Name_list()
	}
	{
		p.SetState(206)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	INC() antlr.TerminalNode
	DEC() antlr.TerminalNode
	ID() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode
	OCTAL() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	SHL() antlr.TerminalNode
	SHR() antlr.TerminalNode
	LE() antlr.TerminalNode
	GE() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	Assign_op() IAssign_opContext
	Expr_list() IExpr_listContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprContext) INC() antlr.TerminalNode {
	return s.GetToken(bParserINC, 0)
}

func (s *ExprContext) DEC() antlr.TerminalNode {
	return s.GetToken(bParserDEC, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(bParserID, 0)
}

func (s *ExprContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(bParserDECIMAL, 0)
}

func (s *ExprContext) OCTAL() antlr.TerminalNode {
	return s.GetToken(bParserOCTAL, 0)
}

func (s *ExprContext) CHAR() antlr.TerminalNode {
	return s.GetToken(bParserCHAR, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(bParserSTRING, 0)
}

func (s *ExprContext) SHL() antlr.TerminalNode {
	return s.GetToken(bParserSHL, 0)
}

func (s *ExprContext) SHR() antlr.TerminalNode {
	return s.GetToken(bParserSHR, 0)
}

func (s *ExprContext) LE() antlr.TerminalNode {
	return s.GetToken(bParserLE, 0)
}

func (s *ExprContext) GE() antlr.TerminalNode {
	return s.GetToken(bParserGE, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(bParserEQ, 0)
}

func (s *ExprContext) NE() antlr.TerminalNode {
	return s.GetToken(bParserNE, 0)
}

func (s *ExprContext) Assign_op() IAssign_opContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_opContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_opContext)
}

func (s *ExprContext) Expr_list() IExpr_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_listContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *bParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *bParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, bParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case bParserT__9, bParserT__10, bParserT__11, bParserT__12, bParserT__13, bParserINC, bParserDEC:
		{
			p.SetState(209)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&206158461952) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(210)
			p.expr(17)
		}

	case bParserT__1:
		{
			p.SetState(211)
			p.Match(bParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.expr(0)
		}
		{
			p.SetState(213)
			p.Match(bParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case bParserID:
		{
			p.SetState(215)
			p.Match(bParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case bParserDECIMAL:
		{
			p.SetState(216)
			p.Match(bParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case bParserOCTAL:
		{
			p.SetState(217)
			p.Match(bParserOCTAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case bParserCHAR:
		{
			p.SetState(218)
			p.Match(bParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case bParserSTRING:
		{
			p.SetState(219)
			p.Match(bParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(222)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(223)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&99328) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(224)
					p.expr(17)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(225)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(226)
					_la = p.GetTokenStream().LA(1)

					if !(_la == bParserT__11 || _la == bParserT__16) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(227)
					p.expr(16)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(229)
					_la = p.GetTokenStream().LA(1)

					if !(_la == bParserSHL || _la == bParserSHR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(230)
					p.expr(15)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(232)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3298535669760) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(233)
					p.expr(14)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(234)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(235)
					_la = p.GetTokenStream().LA(1)

					if !(_la == bParserEQ || _la == bParserNE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(236)
					p.expr(13)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(238)
					p.Match(bParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(239)
					p.expr(12)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(240)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(241)
					p.Match(bParserT__19)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(242)
					p.expr(11)
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(243)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(244)
					p.Match(bParserT__20)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(245)
					p.expr(10)
				}

			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(247)
					p.Match(bParserT__21)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(248)
					p.expr(0)
				}
				{
					p.SetState(249)
					p.Match(bParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(250)
					p.expr(8)
				}

			case 10:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(253)
					p.Assign_op()
				}
				{
					p.SetState(254)
					p.expr(7)
				}

			case 11:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(256)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(257)
					p.Match(bParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(258)
					p.expr(0)
				}
				{
					p.SetState(259)
					p.Match(bParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 12:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(262)
					p.Match(bParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(264)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&-576460700763807999) != 0 {
					{
						p.SetState(263)
						p.Expr_list()
					}

				}
				{
					p.SetState(266)
					p.Match(bParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 13:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, bParserRULE_expr)
				p.SetState(267)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(268)
					_la = p.GetTokenStream().LA(1)

					if !(_la == bParserINC || _la == bParserDEC) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_listContext is an interface to support dynamic dispatch.
type IExpr_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsExpr_listContext differentiates from other interfaces.
	IsExpr_listContext()
}

type Expr_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_listContext() *Expr_listContext {
	var p = new(Expr_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_expr_list
	return p
}

func InitEmptyExpr_listContext(p *Expr_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_expr_list
}

func (*Expr_listContext) IsExpr_listContext() {}

func NewExpr_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_listContext {
	var p = new(Expr_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_expr_list

	return p
}

func (s *Expr_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_listContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Expr_listContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *Expr_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterExpr_list(s)
	}
}

func (s *Expr_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitExpr_list(s)
	}
}

func (p *bParser) Expr_list() (localctx IExpr_listContext) {
	localctx = NewExpr_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, bParserRULE_expr_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.expr(0)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == bParserT__2 {
		{
			p.SetState(275)
			p.Match(bParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.expr(0)
		}

		p.SetState(281)
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

// IAssign_opContext is an interface to support dynamic dispatch.
type IAssign_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	ASS_MUL() antlr.TerminalNode
	ASS_DIV() antlr.TerminalNode
	ASS_MOD() antlr.TerminalNode
	ASS_ADD() antlr.TerminalNode
	ASS_SUB() antlr.TerminalNode
	ASS_SHL() antlr.TerminalNode
	ASS_SHR() antlr.TerminalNode
	ASS_LT() antlr.TerminalNode
	ASS_LE() antlr.TerminalNode
	ASS_GT() antlr.TerminalNode
	ASS_GE() antlr.TerminalNode
	ASS_EQ() antlr.TerminalNode
	ASS_NE() antlr.TerminalNode
	ASS_AND() antlr.TerminalNode
	ASS_XOR() antlr.TerminalNode
	ASS_OR() antlr.TerminalNode

	// IsAssign_opContext differentiates from other interfaces.
	IsAssign_opContext()
}

type Assign_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_opContext() *Assign_opContext {
	var p = new(Assign_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_assign_op
	return p
}

func InitEmptyAssign_opContext(p *Assign_opContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = bParserRULE_assign_op
}

func (*Assign_opContext) IsAssign_opContext() {}

func NewAssign_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_opContext {
	var p = new(Assign_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = bParserRULE_assign_op

	return p
}

func (s *Assign_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_opContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(bParserASSIGN, 0)
}

func (s *Assign_opContext) ASS_MUL() antlr.TerminalNode {
	return s.GetToken(bParserASS_MUL, 0)
}

func (s *Assign_opContext) ASS_DIV() antlr.TerminalNode {
	return s.GetToken(bParserASS_DIV, 0)
}

func (s *Assign_opContext) ASS_MOD() antlr.TerminalNode {
	return s.GetToken(bParserASS_MOD, 0)
}

func (s *Assign_opContext) ASS_ADD() antlr.TerminalNode {
	return s.GetToken(bParserASS_ADD, 0)
}

func (s *Assign_opContext) ASS_SUB() antlr.TerminalNode {
	return s.GetToken(bParserASS_SUB, 0)
}

func (s *Assign_opContext) ASS_SHL() antlr.TerminalNode {
	return s.GetToken(bParserASS_SHL, 0)
}

func (s *Assign_opContext) ASS_SHR() antlr.TerminalNode {
	return s.GetToken(bParserASS_SHR, 0)
}

func (s *Assign_opContext) ASS_LT() antlr.TerminalNode {
	return s.GetToken(bParserASS_LT, 0)
}

func (s *Assign_opContext) ASS_LE() antlr.TerminalNode {
	return s.GetToken(bParserASS_LE, 0)
}

func (s *Assign_opContext) ASS_GT() antlr.TerminalNode {
	return s.GetToken(bParserASS_GT, 0)
}

func (s *Assign_opContext) ASS_GE() antlr.TerminalNode {
	return s.GetToken(bParserASS_GE, 0)
}

func (s *Assign_opContext) ASS_EQ() antlr.TerminalNode {
	return s.GetToken(bParserASS_EQ, 0)
}

func (s *Assign_opContext) ASS_NE() antlr.TerminalNode {
	return s.GetToken(bParserASS_NE, 0)
}

func (s *Assign_opContext) ASS_AND() antlr.TerminalNode {
	return s.GetToken(bParserASS_AND, 0)
}

func (s *Assign_opContext) ASS_XOR() antlr.TerminalNode {
	return s.GetToken(bParserASS_XOR, 0)
}

func (s *Assign_opContext) ASS_OR() antlr.TerminalNode {
	return s.GetToken(bParserASS_OR, 0)
}

func (s *Assign_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.EnterAssign_op(s)
	}
}

func (s *Assign_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bListener); ok {
		listenerT.ExitAssign_op(s)
	}
}

func (p *bParser) Assign_op() (localctx IAssign_opContext) {
	localctx = NewAssign_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, bParserRULE_assign_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2305825417027649536) != 0) {
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

func (p *bParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *bParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 18)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

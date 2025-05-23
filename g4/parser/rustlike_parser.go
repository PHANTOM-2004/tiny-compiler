// Code generated from RustLikeParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RustLikeParser

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

type RustLikeParser struct {
	*antlr.BaseParser
}

var RustLikeParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rustlikeparserParserInit() {
	staticData := &RustLikeParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'i32'", "'let'", "'if'", "'else'", "'while'", "'return'",
		"'mut'", "'fn'", "'loop'", "'break'", "'continue'", "", "", "'+'", "'-'",
		"'*'", "'/'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'='", "';'",
		"':'", "','", "'('", "')'", "'['", "']'", "'{'", "'}'", "'->'", "'.'",
		"'..'",
	}
	staticData.SymbolicNames = []string{
		"", "SL_COMMENT", "ML_COMMENT", "WS", "INT32", "LET", "IF", "ELSE",
		"WHILE", "RETURN", "MUT", "FN", "LOOP", "BREAK", "CONTINUE", "ID", "NUMBER",
		"PLUS", "MINUS", "MULT", "DIV", "EQ", "NE", "LT", "LE", "GT", "GE",
		"ASSIGN", "SEMI", "COLON", "COMMA", "LPAREN", "RPAREN", "LBRAC", "RBRAC",
		"LBRACE", "RBRACE", "ARROW", "DOT", "DOT2",
	}
	staticData.RuleNames = []string{
		"prog", "declaration", "expr", "funcCallList", "funcCallParam", "funcDeclaration",
		"funcDeclarationHeader", "funcDeclarationReturn", "funcParamsList",
		"funcParams", "funcParam", "rtype", "block", "stat", "varType", "varInit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 185, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 1, 5, 1, 36, 8, 1, 10, 1, 12, 1, 39, 9, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 50, 8, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 61, 8, 2, 10, 2, 12, 2, 64, 9,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 73, 8, 4, 10, 4, 12,
		4, 76, 9, 4, 1, 4, 3, 4, 79, 8, 4, 1, 5, 1, 5, 3, 5, 83, 8, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 5, 9, 101, 8, 9, 10, 9, 12, 9, 104, 9, 9, 1, 9, 3, 9, 107,
		8, 9, 1, 10, 3, 10, 110, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 5, 12, 120, 8, 12, 10, 12, 12, 12, 123, 9, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 3, 13, 130, 8, 13, 1, 13, 1, 13, 1, 13, 3, 13,
		135, 8, 13, 1, 13, 1, 13, 3, 13, 139, 8, 13, 1, 13, 3, 13, 142, 8, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 161, 8, 13, 10, 13,
		12, 13, 164, 9, 13, 1, 13, 1, 13, 3, 13, 168, 8, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 177, 8, 13, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 0, 1, 4, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 0, 3, 1, 0, 19, 20, 1, 0, 17, 18, 1, 0, 21, 26,
		196, 0, 32, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 49, 1, 0, 0, 0, 6, 65, 1,
		0, 0, 0, 8, 78, 1, 0, 0, 0, 10, 80, 1, 0, 0, 0, 12, 86, 1, 0, 0, 0, 14,
		90, 1, 0, 0, 0, 16, 93, 1, 0, 0, 0, 18, 106, 1, 0, 0, 0, 20, 109, 1, 0,
		0, 0, 22, 115, 1, 0, 0, 0, 24, 117, 1, 0, 0, 0, 26, 176, 1, 0, 0, 0, 28,
		178, 1, 0, 0, 0, 30, 181, 1, 0, 0, 0, 32, 33, 3, 2, 1, 0, 33, 1, 1, 0,
		0, 0, 34, 36, 3, 10, 5, 0, 35, 34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37,
		35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 3, 1, 0, 0, 0, 39, 37, 1, 0, 0,
		0, 40, 41, 6, 2, -1, 0, 41, 42, 5, 31, 0, 0, 42, 43, 3, 4, 2, 0, 43, 44,
		5, 32, 0, 0, 44, 50, 1, 0, 0, 0, 45, 46, 5, 15, 0, 0, 46, 50, 3, 6, 3,
		0, 47, 50, 5, 15, 0, 0, 48, 50, 5, 16, 0, 0, 49, 40, 1, 0, 0, 0, 49, 45,
		1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 62, 1, 0, 0, 0,
		51, 52, 10, 7, 0, 0, 52, 53, 7, 0, 0, 0, 53, 61, 3, 4, 2, 8, 54, 55, 10,
		6, 0, 0, 55, 56, 7, 1, 0, 0, 56, 61, 3, 4, 2, 7, 57, 58, 10, 5, 0, 0, 58,
		59, 7, 2, 0, 0, 59, 61, 3, 4, 2, 6, 60, 51, 1, 0, 0, 0, 60, 54, 1, 0, 0,
		0, 60, 57, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63,
		1, 0, 0, 0, 63, 5, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 66, 5, 31, 0, 0,
		66, 67, 3, 8, 4, 0, 67, 68, 5, 32, 0, 0, 68, 7, 1, 0, 0, 0, 69, 74, 3,
		4, 2, 0, 70, 71, 5, 30, 0, 0, 71, 73, 3, 4, 2, 0, 72, 70, 1, 0, 0, 0, 73,
		76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 79, 1, 0, 0,
		0, 76, 74, 1, 0, 0, 0, 77, 79, 1, 0, 0, 0, 78, 69, 1, 0, 0, 0, 78, 77,
		1, 0, 0, 0, 79, 9, 1, 0, 0, 0, 80, 82, 3, 12, 6, 0, 81, 83, 3, 14, 7, 0,
		82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 3,
		24, 12, 0, 85, 11, 1, 0, 0, 0, 86, 87, 5, 11, 0, 0, 87, 88, 5, 15, 0, 0,
		88, 89, 3, 16, 8, 0, 89, 13, 1, 0, 0, 0, 90, 91, 5, 37, 0, 0, 91, 92, 3,
		22, 11, 0, 92, 15, 1, 0, 0, 0, 93, 94, 5, 31, 0, 0, 94, 95, 3, 18, 9, 0,
		95, 96, 5, 32, 0, 0, 96, 17, 1, 0, 0, 0, 97, 102, 3, 20, 10, 0, 98, 99,
		5, 30, 0, 0, 99, 101, 3, 20, 10, 0, 100, 98, 1, 0, 0, 0, 101, 104, 1, 0,
		0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 107, 1, 0, 0, 0,
		104, 102, 1, 0, 0, 0, 105, 107, 1, 0, 0, 0, 106, 97, 1, 0, 0, 0, 106, 105,
		1, 0, 0, 0, 107, 19, 1, 0, 0, 0, 108, 110, 5, 10, 0, 0, 109, 108, 1, 0,
		0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 5, 15, 0, 0,
		112, 113, 5, 29, 0, 0, 113, 114, 3, 22, 11, 0, 114, 21, 1, 0, 0, 0, 115,
		116, 5, 4, 0, 0, 116, 23, 1, 0, 0, 0, 117, 121, 5, 35, 0, 0, 118, 120,
		3, 26, 13, 0, 119, 118, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1,
		0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123, 121, 1, 0, 0,
		0, 124, 125, 5, 36, 0, 0, 125, 25, 1, 0, 0, 0, 126, 177, 3, 24, 12, 0,
		127, 129, 5, 9, 0, 0, 128, 130, 3, 4, 2, 0, 129, 128, 1, 0, 0, 0, 129,
		130, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 177, 5, 28, 0, 0, 132, 134,
		5, 5, 0, 0, 133, 135, 5, 10, 0, 0, 134, 133, 1, 0, 0, 0, 134, 135, 1, 0,
		0, 0, 135, 136, 1, 0, 0, 0, 136, 138, 5, 15, 0, 0, 137, 139, 3, 28, 14,
		0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0, 140,
		142, 3, 30, 15, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143,
		1, 0, 0, 0, 143, 177, 5, 28, 0, 0, 144, 145, 5, 15, 0, 0, 145, 146, 5,
		27, 0, 0, 146, 147, 3, 4, 2, 0, 147, 148, 5, 28, 0, 0, 148, 177, 1, 0,
		0, 0, 149, 150, 3, 4, 2, 0, 150, 151, 5, 28, 0, 0, 151, 177, 1, 0, 0, 0,
		152, 153, 5, 6, 0, 0, 153, 154, 3, 4, 2, 0, 154, 162, 3, 24, 12, 0, 155,
		156, 5, 7, 0, 0, 156, 157, 5, 6, 0, 0, 157, 158, 3, 4, 2, 0, 158, 159,
		3, 24, 12, 0, 159, 161, 1, 0, 0, 0, 160, 155, 1, 0, 0, 0, 161, 164, 1,
		0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 167, 1, 0, 0,
		0, 164, 162, 1, 0, 0, 0, 165, 166, 5, 7, 0, 0, 166, 168, 3, 24, 12, 0,
		167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 177, 1, 0, 0, 0, 169,
		170, 5, 8, 0, 0, 170, 171, 3, 4, 2, 0, 171, 172, 3, 24, 12, 0, 172, 177,
		1, 0, 0, 0, 173, 174, 5, 12, 0, 0, 174, 177, 3, 24, 12, 0, 175, 177, 5,
		28, 0, 0, 176, 126, 1, 0, 0, 0, 176, 127, 1, 0, 0, 0, 176, 132, 1, 0, 0,
		0, 176, 144, 1, 0, 0, 0, 176, 149, 1, 0, 0, 0, 176, 152, 1, 0, 0, 0, 176,
		169, 1, 0, 0, 0, 176, 173, 1, 0, 0, 0, 176, 175, 1, 0, 0, 0, 177, 27, 1,
		0, 0, 0, 178, 179, 5, 29, 0, 0, 179, 180, 3, 22, 11, 0, 180, 29, 1, 0,
		0, 0, 181, 182, 5, 27, 0, 0, 182, 183, 3, 4, 2, 0, 183, 31, 1, 0, 0, 0,
		18, 37, 49, 60, 62, 74, 78, 82, 102, 106, 109, 121, 129, 134, 138, 141,
		162, 167, 176,
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

// RustLikeParserInit initializes any static state used to implement RustLikeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRustLikeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RustLikeParserInit() {
	staticData := &RustLikeParserParserStaticData
	staticData.once.Do(rustlikeparserParserInit)
}

// NewRustLikeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRustLikeParser(input antlr.TokenStream) *RustLikeParser {
	RustLikeParserInit()
	this := new(RustLikeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RustLikeParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "RustLikeParser.g4"

	return this
}

// RustLikeParser tokens.
const (
	RustLikeParserEOF        = antlr.TokenEOF
	RustLikeParserSL_COMMENT = 1
	RustLikeParserML_COMMENT = 2
	RustLikeParserWS         = 3
	RustLikeParserINT32      = 4
	RustLikeParserLET        = 5
	RustLikeParserIF         = 6
	RustLikeParserELSE       = 7
	RustLikeParserWHILE      = 8
	RustLikeParserRETURN     = 9
	RustLikeParserMUT        = 10
	RustLikeParserFN         = 11
	RustLikeParserLOOP       = 12
	RustLikeParserBREAK      = 13
	RustLikeParserCONTINUE   = 14
	RustLikeParserID         = 15
	RustLikeParserNUMBER     = 16
	RustLikeParserPLUS       = 17
	RustLikeParserMINUS      = 18
	RustLikeParserMULT       = 19
	RustLikeParserDIV        = 20
	RustLikeParserEQ         = 21
	RustLikeParserNE         = 22
	RustLikeParserLT         = 23
	RustLikeParserLE         = 24
	RustLikeParserGT         = 25
	RustLikeParserGE         = 26
	RustLikeParserASSIGN     = 27
	RustLikeParserSEMI       = 28
	RustLikeParserCOLON      = 29
	RustLikeParserCOMMA      = 30
	RustLikeParserLPAREN     = 31
	RustLikeParserRPAREN     = 32
	RustLikeParserLBRAC      = 33
	RustLikeParserRBRAC      = 34
	RustLikeParserLBRACE     = 35
	RustLikeParserRBRACE     = 36
	RustLikeParserARROW      = 37
	RustLikeParserDOT        = 38
	RustLikeParserDOT2       = 39
)

// RustLikeParser rules.
const (
	RustLikeParserRULE_prog                  = 0
	RustLikeParserRULE_declaration           = 1
	RustLikeParserRULE_expr                  = 2
	RustLikeParserRULE_funcCallList          = 3
	RustLikeParserRULE_funcCallParam         = 4
	RustLikeParserRULE_funcDeclaration       = 5
	RustLikeParserRULE_funcDeclarationHeader = 6
	RustLikeParserRULE_funcDeclarationReturn = 7
	RustLikeParserRULE_funcParamsList        = 8
	RustLikeParserRULE_funcParams            = 9
	RustLikeParserRULE_funcParam             = 10
	RustLikeParserRULE_rtype                 = 11
	RustLikeParserRULE_block                 = 12
	RustLikeParserRULE_stat                  = 13
	RustLikeParserRULE_varType               = 14
	RustLikeParserRULE_varInit               = 15
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaration() IDeclarationContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *RustLikeParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RustLikeParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Declaration()
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

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFuncDeclaration() []IFuncDeclarationContext
	FuncDeclaration(i int) IFuncDeclarationContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) AllFuncDeclaration() []IFuncDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFuncDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDeclarationContext); ok {
			tst[i] = t.(IFuncDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationContext) FuncDeclaration(i int) IFuncDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclarationContext); ok {
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

	return t.(IFuncDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *RustLikeParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RustLikeParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RustLikeParserFN {
		{
			p.SetState(34)
			p.FuncDeclaration()
		}

		p.SetState(39)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = RustLikeParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprAddSubContext struct {
	ExprContext
}

func NewExprAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAddSubContext {
	var p = new(ExprAddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAddSubContext) AllExpr() []IExprContext {
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

func (s *ExprAddSubContext) Expr(i int) IExprContext {
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

func (s *ExprAddSubContext) PLUS() antlr.TerminalNode {
	return s.GetToken(RustLikeParserPLUS, 0)
}

func (s *ExprAddSubContext) MINUS() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMINUS, 0)
}

func (s *ExprAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterExprAddSub(s)
	}
}

func (s *ExprAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitExprAddSub(s)
	}
}

type ExprParenContext struct {
	ExprContext
}

func NewExprParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprParenContext {
	var p = new(ExprParenContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLPAREN, 0)
}

func (s *ExprParenContext) Expr() IExprContext {
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

func (s *ExprParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRPAREN, 0)
}

func (s *ExprParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterExprParen(s)
	}
}

func (s *ExprParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitExprParen(s)
	}
}

type ExprNumContext struct {
	ExprContext
}

func NewExprNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprNumContext {
	var p = new(ExprNumContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNumContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RustLikeParserNUMBER, 0)
}

func (s *ExprNumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterExprNum(s)
	}
}

func (s *ExprNumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitExprNum(s)
	}
}

type ExprMulDivContext struct {
	ExprContext
}

func NewExprMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprMulDivContext {
	var p = new(ExprMulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprMulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMulDivContext) AllExpr() []IExprContext {
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

func (s *ExprMulDivContext) Expr(i int) IExprContext {
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

func (s *ExprMulDivContext) MULT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMULT, 0)
}

func (s *ExprMulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(RustLikeParserDIV, 0)
}

func (s *ExprMulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterExprMulDiv(s)
	}
}

func (s *ExprMulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitExprMulDiv(s)
	}
}

type ExprCmpContext struct {
	ExprContext
}

func NewExprCmpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprCmpContext {
	var p = new(ExprCmpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprCmpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprCmpContext) AllExpr() []IExprContext {
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

func (s *ExprCmpContext) Expr(i int) IExprContext {
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

func (s *ExprCmpContext) EQ() antlr.TerminalNode {
	return s.GetToken(RustLikeParserEQ, 0)
}

func (s *ExprCmpContext) NE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserNE, 0)
}

func (s *ExprCmpContext) LT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLT, 0)
}

func (s *ExprCmpContext) GT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserGT, 0)
}

func (s *ExprCmpContext) LE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLE, 0)
}

func (s *ExprCmpContext) GE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserGE, 0)
}

func (s *ExprCmpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterExprCmp(s)
	}
}

func (s *ExprCmpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitExprCmp(s)
	}
}

type ExprFuncCallContext struct {
	ExprContext
}

func NewExprFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFuncCallContext {
	var p = new(ExprFuncCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprFuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFuncCallContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *ExprFuncCallContext) FuncCallList() IFuncCallListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallListContext)
}

func (s *ExprFuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterExprFuncCall(s)
	}
}

func (s *ExprFuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitExprFuncCall(s)
	}
}

type ExprIDContext struct {
	ExprContext
}

func NewExprIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIDContext {
	var p = new(ExprIDContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIDContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *ExprIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterExprID(s)
	}
}

func (s *ExprIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitExprID(s)
	}
}

func (p *RustLikeParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *RustLikeParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, RustLikeParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(41)
			p.Match(RustLikeParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.expr(0)
		}
		{
			p.SetState(43)
			p.Match(RustLikeParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExprFuncCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.FuncCallList()
		}

	case 3:
		localctx = NewExprIDContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExprNumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.Match(RustLikeParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(52)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RustLikeParserMULT || _la == RustLikeParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(53)
					p.expr(8)
				}

			case 2:
				localctx = NewExprAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(54)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(55)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RustLikeParserPLUS || _la == RustLikeParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(56)
					p.expr(7)
				}

			case 3:
				localctx = NewExprCmpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RustLikeParserRULE_expr)
				p.SetState(57)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(58)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132120576) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(59)
					p.expr(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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

// IFuncCallListContext is an interface to support dynamic dispatch.
type IFuncCallListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	FuncCallParam() IFuncCallParamContext
	RPAREN() antlr.TerminalNode

	// IsFuncCallListContext differentiates from other interfaces.
	IsFuncCallListContext()
}

type FuncCallListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallListContext() *FuncCallListContext {
	var p = new(FuncCallListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcCallList
	return p
}

func InitEmptyFuncCallListContext(p *FuncCallListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcCallList
}

func (*FuncCallListContext) IsFuncCallListContext() {}

func NewFuncCallListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallListContext {
	var p = new(FuncCallListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcCallList

	return p
}

func (s *FuncCallListContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLPAREN, 0)
}

func (s *FuncCallListContext) FuncCallParam() IFuncCallParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallParamContext)
}

func (s *FuncCallListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRPAREN, 0)
}

func (s *FuncCallListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFuncCallList(s)
	}
}

func (s *FuncCallListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFuncCallList(s)
	}
}

func (p *RustLikeParser) FuncCallList() (localctx IFuncCallListContext) {
	localctx = NewFuncCallListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RustLikeParserRULE_funcCallList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(RustLikeParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.FuncCallParam()
	}
	{
		p.SetState(67)
		p.Match(RustLikeParserRPAREN)
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

// IFuncCallParamContext is an interface to support dynamic dispatch.
type IFuncCallParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncCallParamContext differentiates from other interfaces.
	IsFuncCallParamContext()
}

type FuncCallParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallParamContext() *FuncCallParamContext {
	var p = new(FuncCallParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcCallParam
	return p
}

func InitEmptyFuncCallParamContext(p *FuncCallParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcCallParam
}

func (*FuncCallParamContext) IsFuncCallParamContext() {}

func NewFuncCallParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallParamContext {
	var p = new(FuncCallParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcCallParam

	return p
}

func (s *FuncCallParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallParamContext) AllExpr() []IExprContext {
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

func (s *FuncCallParamContext) Expr(i int) IExprContext {
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

func (s *FuncCallParamContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserCOMMA)
}

func (s *FuncCallParamContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOMMA, i)
}

func (s *FuncCallParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFuncCallParam(s)
	}
}

func (s *FuncCallParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFuncCallParam(s)
	}
}

func (p *RustLikeParser) FuncCallParam() (localctx IFuncCallParamContext) {
	localctx = NewFuncCallParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RustLikeParserRULE_funcCallParam)
	var _la int

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RustLikeParserID, RustLikeParserNUMBER, RustLikeParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.expr(0)
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RustLikeParserCOMMA {
			{
				p.SetState(70)
				p.Match(RustLikeParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(71)
				p.expr(0)
			}

			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case RustLikeParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

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

// IFuncDeclarationContext is an interface to support dynamic dispatch.
type IFuncDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncDeclarationHeader() IFuncDeclarationHeaderContext
	Block() IBlockContext
	FuncDeclarationReturn() IFuncDeclarationReturnContext

	// IsFuncDeclarationContext differentiates from other interfaces.
	IsFuncDeclarationContext()
}

type FuncDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclarationContext() *FuncDeclarationContext {
	var p = new(FuncDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcDeclaration
	return p
}

func InitEmptyFuncDeclarationContext(p *FuncDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcDeclaration
}

func (*FuncDeclarationContext) IsFuncDeclarationContext() {}

func NewFuncDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclarationContext {
	var p = new(FuncDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcDeclaration

	return p
}

func (s *FuncDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclarationContext) FuncDeclarationHeader() IFuncDeclarationHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclarationHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclarationHeaderContext)
}

func (s *FuncDeclarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclarationContext) FuncDeclarationReturn() IFuncDeclarationReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclarationReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclarationReturnContext)
}

func (s *FuncDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFuncDeclaration(s)
	}
}

func (s *FuncDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFuncDeclaration(s)
	}
}

func (p *RustLikeParser) FuncDeclaration() (localctx IFuncDeclarationContext) {
	localctx = NewFuncDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RustLikeParserRULE_funcDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.FuncDeclarationHeader()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserARROW {
		{
			p.SetState(81)
			p.FuncDeclarationReturn()
		}

	}
	{
		p.SetState(84)
		p.Block()
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

// IFuncDeclarationHeaderContext is an interface to support dynamic dispatch.
type IFuncDeclarationHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FN() antlr.TerminalNode
	ID() antlr.TerminalNode
	FuncParamsList() IFuncParamsListContext

	// IsFuncDeclarationHeaderContext differentiates from other interfaces.
	IsFuncDeclarationHeaderContext()
}

type FuncDeclarationHeaderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclarationHeaderContext() *FuncDeclarationHeaderContext {
	var p = new(FuncDeclarationHeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcDeclarationHeader
	return p
}

func InitEmptyFuncDeclarationHeaderContext(p *FuncDeclarationHeaderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcDeclarationHeader
}

func (*FuncDeclarationHeaderContext) IsFuncDeclarationHeaderContext() {}

func NewFuncDeclarationHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclarationHeaderContext {
	var p = new(FuncDeclarationHeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcDeclarationHeader

	return p
}

func (s *FuncDeclarationHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclarationHeaderContext) FN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserFN, 0)
}

func (s *FuncDeclarationHeaderContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *FuncDeclarationHeaderContext) FuncParamsList() IFuncParamsListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParamsListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncParamsListContext)
}

func (s *FuncDeclarationHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclarationHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclarationHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFuncDeclarationHeader(s)
	}
}

func (s *FuncDeclarationHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFuncDeclarationHeader(s)
	}
}

func (p *RustLikeParser) FuncDeclarationHeader() (localctx IFuncDeclarationHeaderContext) {
	localctx = NewFuncDeclarationHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RustLikeParserRULE_funcDeclarationHeader)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(RustLikeParserFN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.FuncParamsList()
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

// IFuncDeclarationReturnContext is an interface to support dynamic dispatch.
type IFuncDeclarationReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARROW() antlr.TerminalNode
	Rtype() IRtypeContext

	// IsFuncDeclarationReturnContext differentiates from other interfaces.
	IsFuncDeclarationReturnContext()
}

type FuncDeclarationReturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclarationReturnContext() *FuncDeclarationReturnContext {
	var p = new(FuncDeclarationReturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcDeclarationReturn
	return p
}

func InitEmptyFuncDeclarationReturnContext(p *FuncDeclarationReturnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcDeclarationReturn
}

func (*FuncDeclarationReturnContext) IsFuncDeclarationReturnContext() {}

func NewFuncDeclarationReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclarationReturnContext {
	var p = new(FuncDeclarationReturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcDeclarationReturn

	return p
}

func (s *FuncDeclarationReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclarationReturnContext) ARROW() antlr.TerminalNode {
	return s.GetToken(RustLikeParserARROW, 0)
}

func (s *FuncDeclarationReturnContext) Rtype() IRtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRtypeContext)
}

func (s *FuncDeclarationReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclarationReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclarationReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFuncDeclarationReturn(s)
	}
}

func (s *FuncDeclarationReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFuncDeclarationReturn(s)
	}
}

func (p *RustLikeParser) FuncDeclarationReturn() (localctx IFuncDeclarationReturnContext) {
	localctx = NewFuncDeclarationReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RustLikeParserRULE_funcDeclarationReturn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(RustLikeParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Rtype()
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

// IFuncParamsListContext is an interface to support dynamic dispatch.
type IFuncParamsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	FuncParams() IFuncParamsContext
	RPAREN() antlr.TerminalNode

	// IsFuncParamsListContext differentiates from other interfaces.
	IsFuncParamsListContext()
}

type FuncParamsListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamsListContext() *FuncParamsListContext {
	var p = new(FuncParamsListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParamsList
	return p
}

func InitEmptyFuncParamsListContext(p *FuncParamsListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParamsList
}

func (*FuncParamsListContext) IsFuncParamsListContext() {}

func NewFuncParamsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamsListContext {
	var p = new(FuncParamsListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcParamsList

	return p
}

func (s *FuncParamsListContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamsListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLPAREN, 0)
}

func (s *FuncParamsListContext) FuncParams() IFuncParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncParamsContext)
}

func (s *FuncParamsListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRPAREN, 0)
}

func (s *FuncParamsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFuncParamsList(s)
	}
}

func (s *FuncParamsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFuncParamsList(s)
	}
}

func (p *RustLikeParser) FuncParamsList() (localctx IFuncParamsListContext) {
	localctx = NewFuncParamsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RustLikeParserRULE_funcParamsList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(RustLikeParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.FuncParams()
	}
	{
		p.SetState(95)
		p.Match(RustLikeParserRPAREN)
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

// IFuncParamsContext is an interface to support dynamic dispatch.
type IFuncParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFuncParam() []IFuncParamContext
	FuncParam(i int) IFuncParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncParamsContext differentiates from other interfaces.
	IsFuncParamsContext()
}

type FuncParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamsContext() *FuncParamsContext {
	var p = new(FuncParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParams
	return p
}

func InitEmptyFuncParamsContext(p *FuncParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParams
}

func (*FuncParamsContext) IsFuncParamsContext() {}

func NewFuncParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamsContext {
	var p = new(FuncParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcParams

	return p
}

func (s *FuncParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamsContext) AllFuncParam() []IFuncParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncParamContext); ok {
			len++
		}
	}

	tst := make([]IFuncParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncParamContext); ok {
			tst[i] = t.(IFuncParamContext)
			i++
		}
	}

	return tst
}

func (s *FuncParamsContext) FuncParam(i int) IFuncParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParamContext); ok {
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

	return t.(IFuncParamContext)
}

func (s *FuncParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserCOMMA)
}

func (s *FuncParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOMMA, i)
}

func (s *FuncParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFuncParams(s)
	}
}

func (s *FuncParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFuncParams(s)
	}
}

func (p *RustLikeParser) FuncParams() (localctx IFuncParamsContext) {
	localctx = NewFuncParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RustLikeParserRULE_funcParams)
	var _la int

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RustLikeParserMUT, RustLikeParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.FuncParam()
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RustLikeParserCOMMA {
			{
				p.SetState(98)
				p.Match(RustLikeParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(99)
				p.FuncParam()
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case RustLikeParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

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

// IFuncParamContext is an interface to support dynamic dispatch.
type IFuncParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Rtype() IRtypeContext
	MUT() antlr.TerminalNode

	// IsFuncParamContext differentiates from other interfaces.
	IsFuncParamContext()
}

type FuncParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamContext() *FuncParamContext {
	var p = new(FuncParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParam
	return p
}

func InitEmptyFuncParamContext(p *FuncParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_funcParam
}

func (*FuncParamContext) IsFuncParamContext() {}

func NewFuncParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamContext {
	var p = new(FuncParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_funcParam

	return p
}

func (s *FuncParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *FuncParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOLON, 0)
}

func (s *FuncParamContext) Rtype() IRtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRtypeContext)
}

func (s *FuncParamContext) MUT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMUT, 0)
}

func (s *FuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterFuncParam(s)
	}
}

func (s *FuncParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitFuncParam(s)
	}
}

func (p *RustLikeParser) FuncParam() (localctx IFuncParamContext) {
	localctx = NewFuncParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RustLikeParserRULE_funcParam)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RustLikeParserMUT {
		{
			p.SetState(108)
			p.Match(RustLikeParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(111)
		p.Match(RustLikeParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(RustLikeParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Rtype()
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

// IRtypeContext is an interface to support dynamic dispatch.
type IRtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32() antlr.TerminalNode

	// IsRtypeContext differentiates from other interfaces.
	IsRtypeContext()
}

type RtypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRtypeContext() *RtypeContext {
	var p = new(RtypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_rtype
	return p
}

func InitEmptyRtypeContext(p *RtypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_rtype
}

func (*RtypeContext) IsRtypeContext() {}

func NewRtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RtypeContext {
	var p = new(RtypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_rtype

	return p
}

func (s *RtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RtypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(RustLikeParserINT32, 0)
}

func (s *RtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterRtype(s)
	}
}

func (s *RtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitRtype(s)
	}
}

func (p *RustLikeParser) Rtype() (localctx IRtypeContext) {
	localctx = NewRtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RustLikeParserRULE_rtype)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(RustLikeParserINT32)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStat() []IStatContext
	Stat(i int) IStatContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRBRACE, 0)
}

func (s *BlockContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
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

	return t.(IStatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *RustLikeParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RustLikeParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(RustLikeParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36775760736) != 0 {
		{
			p.SetState(118)
			p.Stat()
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(124)
		p.Match(RustLikeParserRBRACE)
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

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyAll(ctx *StatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatBlockContext struct {
	StatContext
}

func NewStatBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatBlockContext {
	var p = new(StatBlockContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStatBlock(s)
	}
}

func (s *StatBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStatBlock(s)
	}
}

type StatVarAssignContext struct {
	StatContext
}

func NewStatVarAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatVarAssignContext {
	var p = new(StatVarAssignContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatVarAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatVarAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *StatVarAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserASSIGN, 0)
}

func (s *StatVarAssignContext) Expr() IExprContext {
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

func (s *StatVarAssignContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatVarAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStatVarAssign(s)
	}
}

func (s *StatVarAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStatVarAssign(s)
	}
}

type StatExprContext struct {
	StatContext
}

func NewStatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatExprContext {
	var p = new(StatExprContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatExprContext) Expr() IExprContext {
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

func (s *StatExprContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStatExpr(s)
	}
}

func (s *StatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStatExpr(s)
	}
}

type StatFuncReturnContext struct {
	StatContext
}

func NewStatFuncReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatFuncReturnContext {
	var p = new(StatFuncReturnContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatFuncReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatFuncReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserRETURN, 0)
}

func (s *StatFuncReturnContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatFuncReturnContext) Expr() IExprContext {
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

func (s *StatFuncReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStatFuncReturn(s)
	}
}

func (s *StatFuncReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStatFuncReturn(s)
	}
}

type StatVarDeclareContext struct {
	StatContext
}

func NewStatVarDeclareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatVarDeclareContext {
	var p = new(StatVarDeclareContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatVarDeclareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatVarDeclareContext) LET() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLET, 0)
}

func (s *StatVarDeclareContext) ID() antlr.TerminalNode {
	return s.GetToken(RustLikeParserID, 0)
}

func (s *StatVarDeclareContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatVarDeclareContext) MUT() antlr.TerminalNode {
	return s.GetToken(RustLikeParserMUT, 0)
}

func (s *StatVarDeclareContext) VarType() IVarTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *StatVarDeclareContext) VarInit() IVarInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarInitContext)
}

func (s *StatVarDeclareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStatVarDeclare(s)
	}
}

func (s *StatVarDeclareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStatVarDeclare(s)
	}
}

type StatLoopContext struct {
	StatContext
}

func NewStatLoopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatLoopContext {
	var p = new(StatLoopContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatLoopContext) LOOP() antlr.TerminalNode {
	return s.GetToken(RustLikeParserLOOP, 0)
}

func (s *StatLoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStatLoop(s)
	}
}

func (s *StatLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStatLoop(s)
	}
}

type StatIfElseContext struct {
	StatContext
}

func NewStatIfElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatIfElseContext {
	var p = new(StatIfElseContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatIfElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatIfElseContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserIF)
}

func (s *StatIfElseContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserIF, i)
}

func (s *StatIfElseContext) AllExpr() []IExprContext {
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

func (s *StatIfElseContext) Expr(i int) IExprContext {
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

func (s *StatIfElseContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *StatIfElseContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *StatIfElseContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(RustLikeParserELSE)
}

func (s *StatIfElseContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(RustLikeParserELSE, i)
}

func (s *StatIfElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStatIfElse(s)
	}
}

func (s *StatIfElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStatIfElse(s)
	}
}

type StatWhileContext struct {
	StatContext
}

func NewStatWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatWhileContext {
	var p = new(StatWhileContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatWhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(RustLikeParserWHILE, 0)
}

func (s *StatWhileContext) Expr() IExprContext {
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

func (s *StatWhileContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStatWhile(s)
	}
}

func (s *StatWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStatWhile(s)
	}
}

type StatEmptyContext struct {
	StatContext
}

func NewStatEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatEmptyContext {
	var p = new(StatEmptyContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *StatEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatEmptyContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RustLikeParserSEMI, 0)
}

func (s *StatEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterStatEmpty(s)
	}
}

func (s *StatEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitStatEmpty(s)
	}
}

func (p *RustLikeParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RustLikeParserRULE_stat)
	var _la int

	var _alt int

	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStatBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Block()
		}

	case 2:
		localctx = NewStatFuncReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Match(RustLikeParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147581952) != 0 {
			{
				p.SetState(128)
				p.expr(0)
			}

		}
		{
			p.SetState(131)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewStatVarDeclareContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Match(RustLikeParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserMUT {
			{
				p.SetState(133)
				p.Match(RustLikeParserMUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(136)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserCOLON {
			{
				p.SetState(137)
				p.VarType()
			}

		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserASSIGN {
			{
				p.SetState(140)
				p.VarInit()
			}

		}
		{
			p.SetState(143)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStatVarAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(144)
			p.Match(RustLikeParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(RustLikeParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.expr(0)
		}
		{
			p.SetState(147)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStatExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(149)
			p.expr(0)
		}
		{
			p.SetState(150)
			p.Match(RustLikeParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStatIfElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(152)
			p.Match(RustLikeParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.expr(0)
		}
		{
			p.SetState(154)
			p.Block()
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(155)
					p.Match(RustLikeParserELSE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(156)
					p.Match(RustLikeParserIF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(157)
					p.expr(0)
				}
				{
					p.SetState(158)
					p.Block()
				}

			}
			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == RustLikeParserELSE {
			{
				p.SetState(165)
				p.Match(RustLikeParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(166)
				p.Block()
			}

		}

	case 7:
		localctx = NewStatWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(169)
			p.Match(RustLikeParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.expr(0)
		}
		{
			p.SetState(171)
			p.Block()
		}

	case 8:
		localctx = NewStatLoopContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(173)
			p.Match(RustLikeParserLOOP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Block()
		}

	case 9:
		localctx = NewStatEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(175)
			p.Match(RustLikeParserSEMI)
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

// IVarTypeContext is an interface to support dynamic dispatch.
type IVarTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	Rtype() IRtypeContext

	// IsVarTypeContext differentiates from other interfaces.
	IsVarTypeContext()
}

type VarTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarTypeContext() *VarTypeContext {
	var p = new(VarTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_varType
	return p
}

func InitEmptyVarTypeContext(p *VarTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_varType
}

func (*VarTypeContext) IsVarTypeContext() {}

func NewVarTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarTypeContext {
	var p = new(VarTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_varType

	return p
}

func (s *VarTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VarTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(RustLikeParserCOLON, 0)
}

func (s *VarTypeContext) Rtype() IRtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRtypeContext)
}

func (s *VarTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterVarType(s)
	}
}

func (s *VarTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitVarType(s)
	}
}

func (p *RustLikeParser) VarType() (localctx IVarTypeContext) {
	localctx = NewVarTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RustLikeParserRULE_varType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(RustLikeParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Rtype()
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

// IVarInitContext is an interface to support dynamic dispatch.
type IVarInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsVarInitContext differentiates from other interfaces.
	IsVarInitContext()
}

type VarInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarInitContext() *VarInitContext {
	var p = new(VarInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_varInit
	return p
}

func InitEmptyVarInitContext(p *VarInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RustLikeParserRULE_varInit
}

func (*VarInitContext) IsVarInitContext() {}

func NewVarInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarInitContext {
	var p = new(VarInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustLikeParserRULE_varInit

	return p
}

func (s *VarInitContext) GetParser() antlr.Parser { return s.parser }

func (s *VarInitContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(RustLikeParserASSIGN, 0)
}

func (s *VarInitContext) Expr() IExprContext {
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

func (s *VarInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.EnterVarInit(s)
	}
}

func (s *VarInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustLikeParserListener); ok {
		listenerT.ExitVarInit(s)
	}
}

func (p *RustLikeParser) VarInit() (localctx IVarInitContext) {
	localctx = NewVarInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RustLikeParserRULE_varInit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(RustLikeParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
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

func (p *RustLikeParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RustLikeParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

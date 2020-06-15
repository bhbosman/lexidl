package lexidl

import (
	"fmt"
	"github.com/bhbosman/lexpragma"
	"github.com/bhbosman/lexpragmaids"
	"github.com/bhbosman/yaccpragma"
	"github.com/bhbosman/yaccpragmaids"
	"strconv"
)

type IYaccPragmaLexer interface {
	yaccpragma.YaccPragmaLexer
	ErrorMessage() string
	ErrorOccurred() bool
	PragmaNode() yaccpragma.IPragmaNode
}

type lexerImpl struct {
	handler      *lexpragma.Handler
	pragmaNode   yaccpragma.IPragmaNode
	error        bool
	errorMessage string
}

func (y lexerImpl) ErrorOccurred() bool {
	return y.error
}

func (y lexerImpl) ErrorMessage() string {
	return y.errorMessage
}

func (y lexerImpl) PragmaNode() yaccpragma.IPragmaNode {
	return y.pragmaNode
}

func (y lexerImpl) Read(s string) (yaccpragmaids.ITypeVersion, error) {
	lex, err := lexpragmaids.NewPragmaLexIdsFromData("", s)
	if err != nil {
		return nil, err
	}
	lexWrapper := yaccpragma.NewIdsLexerWrapper(lex)
	parser := yaccpragmaids.YaccPragmaIdsNewParser()
	v := parser.Parse(lexWrapper)
	if v != 0 || lexWrapper.ErrorOccurred() {
		return nil, fmt.Errorf("error in parsing pagmaids. Error %v: ", lexWrapper.ErrorMessage())
	}
	return lexWrapper.Version(), nil
}

func (y lexerImpl) GetPragma() yaccpragma.IPragmaNode {
	return y.pragmaNode
}

func (y *lexerImpl) SetPragma(node yaccpragma.IPragmaNode) {
	y.pragmaNode = node
}

func NewYaccPragmaLexer(handler *lexpragma.Handler) IYaccPragmaLexer {
	return &lexerImpl{
		handler: handler,
	}
}

func (y lexerImpl) Lex(lval *yaccpragma.YaccPragmaSymType) int {
	for {
		lexem := y.handler.ReadLexem()

		if lexem.Eof {
			return 0
		}

		switch lexem.TypeKind {
		case yaccpragma.WhiteSpace:
			continue
		case yaccpragma.StringLiteral, yaccpragma.Identifier:
			lval.StringLiteral = lexem.Value
			return lexem.TypeKind
		case yaccpragma.IntLiteral:
			atoi, _ := strconv.Atoi(lexem.Value)
			lval.IntLiteral = int64(atoi)
			return yaccpragma.IntLiteral
		}
		return lexem.TypeKind

	}

}

func (y *lexerImpl) Error(s string) {
	y.error = true
	y.errorMessage = s
}

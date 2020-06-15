package lexidl

import (
	yaccToken "github.com/bhbosman/yaccidl"
	"github.com/bhbosman/yaccpragma"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPragmaPrefix(t *testing.T) {
	t.Run("", func(t *testing.T) {
		s := "#pragma  prefix \"omg.com\"\n"
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashPragma, lexem.TypeKind)
		assert.NotNil(t, lexem.PragmaNodeValue)
		if !assert.Implements(t, (*yaccpragma.IPragmaPrefixNode)(nil), lexem.PragmaNodeValue) {
			return
		}
		prefix, ok := lexem.PragmaNodeValue.(yaccpragma.IPragmaPrefixNode)
		assert.True(t, ok)
		assert.Equal(t, "Prefix", prefix.PragmaType())
		assert.Equal(t, "omg.com", prefix.Prefix())
	})

	t.Run("", func(t *testing.T) {
		s := "#pragma  prefix omg.com\n"
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		_, err := handler.ReadLexem()
		assert.Error(t, err)
	})
}

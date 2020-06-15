package lexidl

import (
	yaccToken "github.com/bhbosman/yaccidl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_Integers(t *testing.T) {
	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "1234", nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(1234), lexem.IntValue)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "1234 5678", nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(1234), lexem.IntValue)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(5678), lexem.IntValue)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "1 2 0 ", nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(1), lexem.IntValue)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(2), lexem.IntValue)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(0), lexem.IntValue)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", " 0 ", nil)
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(0), lexem.IntValue)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		s := `
		#ifdef AA
			1234
		#endif
		`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, 0, handler.DefinitionData.ExpressionBlockCount)
	})
}

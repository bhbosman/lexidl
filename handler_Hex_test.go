package lexidl

import (
	yaccToken "github.com/bhbosman/yaccidl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_Hex(t *testing.T) {
	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "0X1234", nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(0x1234), lexem.IntValue)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "0x1234", nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(0x1234), lexem.IntValue)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "/**/", nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		//assert.Equal(t, int64(0x1234), lexem.IntValue)
		assert.Equal(t, yaccToken.SingleComment, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "/* * */", nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		//assert.Equal(t, int64(0x1234), lexem.IntValue)
		assert.Equal(t, yaccToken.SingleComment, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "/*dfsdfsdfsdfsd*/", nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		//assert.Equal(t, int64(0x1234), lexem.IntValue)
		assert.Equal(t, yaccToken.SingleComment, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "/*\n\n\n\n*/", nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		//assert.Equal(t, int64(0x1234), lexem.IntValue)
		assert.Equal(t, yaccToken.SingleComment, lexem.TypeKind)
	})

}

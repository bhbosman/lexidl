package lexidl

import (
	yaccToken "github.com/bhbosman/yaccidl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_LoadFile(t *testing.T) {

	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "12 34", nil)
		assert.NoError(t, e)
		handler.AddPredefinedFile("a.a", "56\t78")
		handler.AddPredefinedFile("b.a", "90\n0")

		assert.NotNil(t, handler.currentContext)
		assert.Equal(t, "(test stream)", handler.currentContext.FileName)
		//assert.Equal(t, byte('1'), handler.current)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)

		handler.loadFile(*handler.currentContext, "a.a")
		handler.closeTopStream()

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewLexIdlHandlerFromData("(test stream)", "12 34", nil)
		assert.NoError(t, e)
		handler.AddPredefinedFile("a.a", "56\t78")
		handler.AddPredefinedFile("b.a", "90\n0")

		assert.NotNil(t, handler.currentContext.ReaderCloser)
		assert.Equal(t, "(test stream)", handler.currentContext.FileName)
		assert.Equal(t, byte(0), handler.current)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(12), lexem.IntValue)
		assert.Equal(t, byte(32), handler.current)

		handler.loadFile(*handler.currentContext, "a.a")
		assert.Equal(t, byte(0), handler.current)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(56), lexem.IntValue)
		assert.Equal(t, byte('\t'), handler.current)

		handler.loadFile(*handler.currentContext, "b.a")
		assert.Equal(t, byte(0), handler.current)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(90), lexem.IntValue)
		assert.Equal(t, byte('\n'), handler.current)
		handler.closeTopStream()

		handler.loadFile(*handler.currentContext, "b.a")
		assert.Equal(t, byte(0), handler.current)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(90), lexem.IntValue)
		assert.Equal(t, byte('\n'), handler.current)

		handler.closeTopStream()

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(78), lexem.IntValue)
		assert.Equal(t, byte(0), handler.current)

		handler.closeTopStream()
		assert.Equal(t, "(test stream)", handler.currentContext.FileName)
		assert.Equal(t, byte(0), handler.current)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, int64(34), lexem.IntValue)
	})
}

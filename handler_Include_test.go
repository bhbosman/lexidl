package lexidl

import (
	yaccToken "github.com/bhbosman/yaccidl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_Include(t *testing.T) {
	t.Run("", func(t *testing.T) {
		s := `
			#include<ddddd.h>
		`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		handler.AddPredefinedFile("ddddd.h", "1245\n12345")

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)

		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashInclude, lexem.TypeKind)
		assert.Equal(t, "ddddd.h", lexem.StringValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(1245), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(12345), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, true, lexem.Eof)
	})

	t.Run("", func(t *testing.T) {
		s := `
			9876
			#include<ddddd.h>
			9876
		`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		handler.AddPredefinedFile("ddddd.h", "1245\n12345")
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)

		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(9876), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashInclude, lexem.TypeKind)
		assert.Equal(t, "ddddd.h", lexem.StringValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(1245), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(12345), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(9876), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, true, lexem.Eof)
	})

	t.Run("", func(t *testing.T) {
		s := `
			9876
			#include<ddddd.h>
			9876
			#include<ddddd.h>

		`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		handler.AddPredefinedFile("ddddd.h", "1245\n12345")
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(9876), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashInclude, lexem.TypeKind)
		assert.Equal(t, "ddddd.h", lexem.StringValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(1245), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(12345), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(9876), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashInclude, lexem.TypeKind)
		assert.Equal(t, "ddddd.h", lexem.StringValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(1245), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(12345), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, true, lexem.Eof)
	})

	t.Run("", func(t *testing.T) {
		s := "#include<a.h>\n#include<b.h>\n"
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		handler.AddPredefinedFile("a.h", "")
		handler.AddPredefinedFile("b.h", "")

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashInclude, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashInclude, lexem.TypeKind)
	})
}

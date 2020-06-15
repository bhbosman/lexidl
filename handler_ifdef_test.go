package lexidl

import (
	yaccToken "github.com/bhbosman/yaccidl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_IfDef(t *testing.T) {
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
		assert.Equal(t, yaccToken.Hashifdef, lexem.TypeKind)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)
		assert.Equal(t, false, lexem.ValidToken)
		assert.Equal(t, int64(1234), lexem.IntValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, 0, handler.DefinitionData.ExpressionBlockCount)
	})

	t.Run("", func(t *testing.T) {
		s := `
		#ifdef AA
			1234
		#else
			12345
		#endif
		`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashifdef, lexem.TypeKind)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(1234), lexem.IntValue)
		assert.Equal(t, false, lexem.ValidToken)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashelse, lexem.TypeKind)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(12345), lexem.IntValue)
		assert.Equal(t, true, lexem.ValidToken)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, 0, handler.DefinitionData.ExpressionBlockCount)

	})

	t.Run("", func(t *testing.T) {
		s := `
		#define AA
		#ifdef AA
			1234
		#else
			12345
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
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(1234), lexem.IntValue)
		assert.Equal(t, true, lexem.ValidToken)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

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
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(12345), lexem.IntValue)
		assert.Equal(t, false, lexem.ValidToken)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, 0, handler.DefinitionData.ExpressionBlockCount)

	})

	t.Run("", func(t *testing.T) {
		s := `
		#define AA
		#ifdef AA
			1234
		#else
			12345
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
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(1234), lexem.IntValue)
		assert.Equal(t, true, lexem.ValidToken)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashelse, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(12345), lexem.IntValue)
		assert.Equal(t, false, lexem.ValidToken)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, 0, handler.DefinitionData.ExpressionBlockCount)

	})

	t.Run("", func(t *testing.T) {
		s := `
		#define AA
		1234
		#ifdef AA
			#define BB
			12345
		#else
			123456
		#endif
		`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		assert.Equal(t, INITIAL, handler.startCond)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashDefine, lexem.TypeKind)
		assert.Equal(t, "AA", lexem.StringValue)
		assert.Equal(t, INITIAL, handler.startCond)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(1234), lexem.IntValue)
		assert.Equal(t, 0, handler.DefinitionData.ExpressionBlockCount)
		assert.Equal(t, INITIAL, handler.startCond)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashifdef, lexem.TypeKind)
		defined, s2 := handler.DefinitionData.isDefined("AA")
		assert.True(t, defined)
		assert.Equal(t, "", s2)
		assert.Equal(t, DataBlock, handler.startCond)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashDefine, lexem.TypeKind)
		assert.Equal(t, "BB", lexem.StringValue)
		assert.Equal(t, DataBlock, handler.startCond)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(12345), lexem.IntValue)
		assert.Equal(t, true, handler.DefinitionData.validBlock())
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)
		assert.Equal(t, DataBlock, handler.startCond)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashelse, lexem.TypeKind)
		assert.Equal(t, DataBlock, handler.startCond)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(123456), lexem.IntValue)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)
		assert.Equal(t, false, handler.DefinitionData.validBlock())
		assert.Equal(t, DataBlock, handler.startCond)
		assert.Equal(t, DataBlock, handler.startCond)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, INITIAL, handler.startCond)
	})

	t.Run("", func(t *testing.T) {
		s := `
		#ifdef AA
			#ifdef BB
			#endif
		#endif
		`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		assert.Equal(t, INITIAL, handler.startCond)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashifdef, lexem.TypeKind)
		assert.Equal(t, DataBlock, handler.startCond)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashifdef, lexem.TypeKind)
		assert.Equal(t, DataBlock, handler.startCond)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, DataBlock, handler.startCond)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, INITIAL, handler.startCond)
	})

	t.Run("", func(t *testing.T) {
		s := "#define AA\n#undefine AA\n"
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashDefine, lexem.TypeKind)
		defined, s2 := handler.DefinitionData.isDefined("AA")
		assert.Equal(t, true, defined)
		assert.Equal(t, "", s2)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashUnDefine, lexem.TypeKind)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)

	})

	t.Run("", func(t *testing.T) {
		s := "#undefine AA\n#undefine AA\n"
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		defined, s2 := handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashUnDefine, lexem.TypeKind)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashUnDefine, lexem.TypeKind)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)
	})

	t.Run("", func(t *testing.T) {
		s := `
			#define AA some definition
			#undefine AA
			`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		defined, s2 := handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashDefine, lexem.TypeKind)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.Equal(t, true, defined)
		assert.Equal(t, "some definition", s2)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.Equal(t, true, defined)
		assert.Equal(t, "some definition", s2)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashUnDefine, lexem.TypeKind)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)
	})

	t.Run("", func(t *testing.T) {
		s := `
			#define AA some definition
			AA
			BB
			`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		defined, s2 := handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashDefine, lexem.TypeKind)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.Equal(t, true, defined)
		assert.Equal(t, "some definition", s2)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashLoadDefinition, lexem.TypeKind)
		assert.Equal(t, "AA=some definition", lexem.StringValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Identifier, lexem.TypeKind)
		assert.Equal(t, "some", lexem.StringValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Identifier, lexem.TypeKind)
		assert.Equal(t, "definition", lexem.StringValue)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Identifier, lexem.TypeKind)
		assert.Equal(t, "BB", lexem.StringValue)

	})

	t.Run("", func(t *testing.T) {
		s := `
			#ifdef AA
			#define AA some definition 0001
			#else
			#define AA some definition 0002
			#endif
			`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		defined, s2 := handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashifdef, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashDefine, lexem.TypeKind)
		assert.False(t, lexem.ValidToken)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.Equal(t, false, defined)
		assert.Equal(t, "", s2)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashelse, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashDefine, lexem.TypeKind)
		assert.True(t, lexem.ValidToken)
		defined, s2 = handler.DefinitionData.isDefined("AA")
		assert.True(t, defined)
		assert.Equal(t, "some definition 0002", s2)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)
	})
}

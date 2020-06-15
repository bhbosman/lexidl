package lexidl

import (
	yaccToken "github.com/bhbosman/yaccidl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_IfNotDef(t *testing.T) {
	t.Run("", func(t *testing.T) {
		s := `
		#ifndef AA
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
		assert.Equal(t, yaccToken.Hashifndef, lexem.TypeKind)
		assert.Equal(t, "AA", lexem.StringValue)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(1234), lexem.IntValue)
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
		1
		#define AA
		#define BB

		#ifndef AA
			2
			#ifndef BB
				3
			#else
				4
			#endif
				5
		#else
			6
			#ifndef BB
				7
			#else
				8
			#endif
		#endif
		`
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(1), lexem.IntValue)
		assert.Equal(t, true, lexem.ValidToken)
		assert.Equal(t, 0, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashDefine, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashDefine, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashifndef, lexem.TypeKind)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(2), lexem.IntValue)
		assert.Equal(t, false, lexem.ValidToken)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashifndef, lexem.TypeKind)
		assert.Equal(t, 2, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(3), lexem.IntValue)
		assert.Equal(t, false, lexem.ValidToken)
		assert.Equal(t, 2, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashelse, lexem.TypeKind)
		assert.Equal(t, 2, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(4), lexem.IntValue)
		assert.Equal(t, false, lexem.ValidToken)
		assert.Equal(t, 2, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(5), lexem.IntValue)
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
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)
		assert.Equal(t, int64(6), lexem.IntValue)
		assert.Equal(t, true, lexem.ValidToken)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashifndef, lexem.TypeKind)
		assert.Equal(t, 2, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(7), lexem.IntValue)
		assert.Equal(t, false, lexem.ValidToken)
		assert.Equal(t, 2, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashelse, lexem.TypeKind)
		assert.Equal(t, 2, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
		assert.Equal(t, int64(8), lexem.IntValue)
		assert.Equal(t, true, lexem.ValidToken)
		assert.Equal(t, 2, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, 1, handler.DefinitionData.ExpressionBlockCount)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.WhiteSpace, lexem.TypeKind)

		lexem, err = handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.Hashendif, lexem.TypeKind)
		assert.Equal(t, 0, handler.DefinitionData.ExpressionBlockCount)
	})
}

package lexidl

import (
	yaccToken "github.com/bhbosman/yaccidl"
	"github.com/bhbosman/yaccpragma"
	"github.com/bhbosman/yaccpragmaids"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPragmaId(t *testing.T) {
	t.Run("error pragma - #pragma  ID  omg.com", func(t *testing.T) {
		s := "#pragma  ID  omg.com\n"
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		_, err := handler.ReadLexem()
		assert.Error(t, err)
	})
	t.Run("IDL Identifier", func(t *testing.T) {
		s := "#pragma  ID A  \"IDL:A:1.1\"\n"
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccToken.HashPragma, lexem.TypeKind)
		assert.NotNil(t, lexem.PragmaNodeValue)
		if !assert.Implements(t, (*yaccpragma.IPragmaIdentifierNode)(nil), lexem.PragmaNodeValue) {
			return
		}
		prefix, ok := lexem.PragmaNodeValue.(yaccpragma.IPragmaIdentifierNode)
		assert.True(t, ok)
		assert.Equal(t, "Id", prefix.PragmaType())
		assert.Equal(t, "A", prefix.Identifier())
		assert.NotNil(t, prefix.Value())
		v := prefix.Value()
		assert.Equal(t, "IDL", v.VersionType())
		if !assert.Implements(t, (*yaccpragmaids.IIdlVersion)(nil), v) {
			return
		}
		idlVersion, ok := v.(yaccpragmaids.IIdlVersion)
		assert.True(t, ok)
		assert.Equal(t, uint64(1), idlVersion.Value().Major())
		assert.Equal(t, uint64(1), idlVersion.Value().Minor())
		assert.Equal(t, "A", idlVersion.Identifier())
		assert.Equal(t, "IDL", idlVersion.VersionType())

	})

	t.Run("IDL Identifier", func(t *testing.T) {
		s := "#pragma  ID A  \"RMI:foo/bar;:1234567812345678\"\n"
		handler, e := NewLexIdlHandlerFromData("(test stream)", s, nil)
		assert.NoError(t, e)
		lexem, err := handler.ReadLexem()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, yaccToken.HashPragma, lexem.TypeKind)
		assert.NotNil(t, lexem.PragmaNodeValue)
		if !assert.Implements(t, (*yaccpragma.IPragmaIdentifierNode)(nil), lexem.PragmaNodeValue) {
			return
		}
		prefix, ok := lexem.PragmaNodeValue.(yaccpragma.IPragmaIdentifierNode)
		assert.True(t, ok)
		assert.Equal(t, "Id", prefix.PragmaType())
		assert.Equal(t, "A", prefix.Identifier())
		assert.NotNil(t, prefix.Value())
		v := prefix.Value()
		assert.Equal(t, "RMI", v.VersionType())
		if !assert.Implements(t, (*yaccpragmaids.IRmiVersion)(nil), v) {
			return
		}
		rmiVersion, ok := v.(yaccpragmaids.IRmiVersion)
		assert.True(t, ok)
		assert.Equal(t, "foo/bar;", rmiVersion.ClassName())
		assert.Equal(t, uint64(0x1234567812345678), rmiVersion.Hash())

	})

}

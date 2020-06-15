package lexidl

import (
	yaccToken "github.com/bhbosman/yaccidl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerReservedWord(t *testing.T) {
	t.Run("Reserved Words", func(t *testing.T) {
		multipleData := []struct {
			token        int
			reservedWord string
		}{
			{token: yaccToken.RWabstract, reservedWord: "abstract"},
			{token: yaccToken.RWany, reservedWord: "any"},
			{token: yaccToken.RWalias, reservedWord: "alias"},
			{token: yaccToken.RWattribute, reservedWord: "attribute"},
			{token: yaccToken.RWbitfield, reservedWord: "bitfield"},
			{token: yaccToken.RWbitmask, reservedWord: "bitmask"},
			{token: yaccToken.RWbitset, reservedWord: "bitset"},
			{token: yaccToken.RWboolean, reservedWord: "boolean"},
			{token: yaccToken.RWcase, reservedWord: "case"},
			{token: yaccToken.RWchar, reservedWord: "char"},
			{token: yaccToken.RWcomponent, reservedWord: "component"},
			{token: yaccToken.RWconnector, reservedWord: "connector"},
			{token: yaccToken.RWconst, reservedWord: "const"},
			{token: yaccToken.RWconsumes, reservedWord: "consumes"},
			{token: yaccToken.RWcontext, reservedWord: "context"},
			{token: yaccToken.RWcustom, reservedWord: "custom"},
			{token: yaccToken.RWdefault, reservedWord: "default"},
			{token: yaccToken.RWdouble, reservedWord: "double"},
			{token: yaccToken.RWexception, reservedWord: "exception"},
			{token: yaccToken.RWemits, reservedWord: "emits"},
			{token: yaccToken.RWenum, reservedWord: "enum"},
			{token: yaccToken.RWeventtype, reservedWord: "eventtype"},
			{token: yaccToken.RWfactory, reservedWord: "factory"},
			{token: yaccToken.RWFALSE, reservedWord: "FALSE"},
			{token: yaccToken.RWfinder, reservedWord: "finder"},
			{token: yaccToken.RWfixed, reservedWord: "fixed"},
			{token: yaccToken.RWfloat, reservedWord: "float"},
			{token: yaccToken.RWgetraises, reservedWord: "getraises"},
			{token: yaccToken.RWhome, reservedWord: "home"},
			{token: yaccToken.RWimport, reservedWord: "import"},
			{token: yaccToken.RWin, reservedWord: "in"},
			{token: yaccToken.RWinout, reservedWord: "inout"},
			{token: yaccToken.RWinterface, reservedWord: "interface"},
			{token: yaccToken.RWlocal, reservedWord: "local"},
			{token: yaccToken.RWlong, reservedWord: "long"},
			{token: yaccToken.RWmanages, reservedWord: "manages"},
			{token: yaccToken.RWmap, reservedWord: "map"},
			{token: yaccToken.RWmirrorport, reservedWord: "mirrorport"},
			{token: yaccToken.RWmodule, reservedWord: "module"},
			{token: yaccToken.RWmultiple, reservedWord: "multiple"},
			{token: yaccToken.RWnative, reservedWord: "native"},
			{token: yaccToken.RWObject, reservedWord: "Object"},
			{token: yaccToken.RWoctet, reservedWord: "octet"},
			{token: yaccToken.RWoneway, reservedWord: "oneway"},
			{token: yaccToken.RWout, reservedWord: "out"},
			{token: yaccToken.RWprimarykey, reservedWord: "primarykey"},
			{token: yaccToken.RWprivate, reservedWord: "private"},
			{token: yaccToken.RWport, reservedWord: "port"},
			{token: yaccToken.RWporttype, reservedWord: "porttype"},
			{token: yaccToken.RWprovides, reservedWord: "provides"},
			{token: yaccToken.RWpublic, reservedWord: "public"},
			{token: yaccToken.RWpublishes, reservedWord: "publishes"},
			{token: yaccToken.RWraises, reservedWord: "raises"},
			{token: yaccToken.RWreadonly, reservedWord: "readonly"},
			{token: yaccToken.RWsetraises, reservedWord: "setraises"},
			{token: yaccToken.RWsequence, reservedWord: "sequence"},
			{token: yaccToken.RWshort, reservedWord: "short"},
			{token: yaccToken.RWstring, reservedWord: "string"},
			{token: yaccToken.RWstruct, reservedWord: "struct"},
			{token: yaccToken.RWsupports, reservedWord: "supports"},
			{token: yaccToken.RWswitch, reservedWord: "switch"},
			{token: yaccToken.RWTRUE, reservedWord: "TRUE"},
			{token: yaccToken.RWtruncatable, reservedWord: "truncatable"},
			{token: yaccToken.RWtypedef, reservedWord: "typedef"},
			{token: yaccToken.RWtypeid, reservedWord: "typeid"},
			{token: yaccToken.RWtypename, reservedWord: "typename"},
			{token: yaccToken.RWtypeprefix, reservedWord: "typeprefix"},
			{token: yaccToken.RWunsigned, reservedWord: "unsigned"},
			{token: yaccToken.RWunion, reservedWord: "union"},
			{token: yaccToken.RWuses, reservedWord: "uses"},
			{token: yaccToken.RWValueBase, reservedWord: "ValueBase"},
			{token: yaccToken.RWvaluetype, reservedWord: "valuetype"},
			{token: yaccToken.RWvoid, reservedWord: "void"},
			{token: yaccToken.RWwchar, reservedWord: "wchar"},
			{token: yaccToken.RWwstring, reservedWord: "wstring"},
			{token: yaccToken.RWint8, reservedWord: "int8"},
			{token: yaccToken.RWuint8, reservedWord: "uint8"},
			{token: yaccToken.RWint16, reservedWord: "int16"},
			{token: yaccToken.RWint32, reservedWord: "int32"},
			{token: yaccToken.RWint64, reservedWord: "int64"},
			{token: yaccToken.RWuint16, reservedWord: "uint16"},
			{token: yaccToken.RWuint32, reservedWord: "uint32"},
			{token: yaccToken.RWuint64, reservedWord: "uint64"},
		}

		for _, data := range multipleData {
			handler, e := NewLexIdlHandlerFromData("(test stream)", data.reservedWord, nil)
			assert.NoError(t, e)
			lexem, err := handler.ReadLexem()
			assert.NoError(t, err)
			assert.Equal(t, data.token, lexem.TypeKind)
		}
	})
	t.Run("integers", func(t *testing.T) {
		multipleData := []struct {
			s string
		}{
			{s: "1"},
			{s: "0"},
			{s: "0x12"},
		}
		for _, data := range multipleData {
			t.Run(data.s, func(t *testing.T) {
				handler, e := NewLexIdlHandlerFromData("(test stream)", data.s, nil)
				assert.NoError(t, e)
				lexem, err := handler.ReadLexem()
				assert.NoError(t, err)
				assert.Equal(t, yaccToken.Integer_literal, lexem.TypeKind)
			})
		}
	})
}

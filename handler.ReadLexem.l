%{
package lexidl

import "strconv"
import "log"
import "fmt"
import "strings"
import "unicode"
import yaccToken "github.com/bhbosman/yaccidl"

	const (
		INITIAL = iota
		DataBlock
		IfdefExpression
		IfndefExpression
		DefineExpression
		UnDefineExpression
		IncludeExpression
		PragmaExpression
		MultiLineComment00
		MultiLineComment01
	)


func (self *Handler) readLexem(ctx yaccToken.CurrentContext) (*yaccToken.LexemValue, error) {
    if !self.currentAssigned {
        self.current = self.GetChar()
        self.currentAssigned = true
    }
%}

%yyc self.current
%yyn self.current = self.GetChar()
%yyt self.startCond

D  [0-9]+
E  [eE][-+]?{D}
F  {D}"."{D}?{E}?|{D}{E}?|"."{D}{E}?

%s datablock
%x IfdefExpression
%x IfndefExpression
%x DefineExpression
%x UnDefineExpression
%x IncludeExpression
%x PragmaExpression
%x MultiLineComment00
%x MultiLineComment01
%%
    self.buf = self.buf[:0]

"/*"
    {
        // Save current StartCondition.
        self.prevCond = self.startCond
        // Change StartCondition to read the Include Expression
        self.startCond = MultiLineComment00
    }
<MultiLineComment00>.
    {
        ch := self.buf
        if len(ch) == 1{
            if ch[0] == '*'{
                self.startCond = MultiLineComment01
            }
        }
     }
<MultiLineComment00,MultiLineComment01>\n
    {
        self.startCond = MultiLineComment00
    }
<MultiLineComment01>.
    {
        ch := self.buf
        if len(ch) == 1{
            switch ch[0]{
                case '*':
                    // do nothing
                    break
                case '/':
                    self.startCond = self.prevCond
                    return yaccToken.NewLexemStringValue(yaccToken.SingleComment, self.TokenName, string(self.buf), ctx, self.DefinitionData.validBlock())
                default:
                    self.startCond = MultiLineComment00
                    break
            }
        }
     }


<INITIAL,datablock>"#"pragma
    {
        // Save current StartCondition.
        self.prevCond = self.startCond
        // Change StartCondition to read the Include Expression
        self.startCond = PragmaExpression
    }
<PragmaExpression>(.*)\n
    {
        self.startCond = self.prevCond
        s := strings.TrimSpace(string(self.buf[:len(self.buf)-1]))
        return self.ReadPragmaExpression(ctx, s)
    }
<INITIAL,datablock>"#"include
    {
        // Save current StartCondition.
        self.prevCond = self.startCond
        // Change StartCondition to read the Include Expression
        self.startCond = IncludeExpression
    }
<IncludeExpression>(.*)\n
    {
        // Change StartCondition to read the Include Expression
        self.startCond = self.prevCond

        // Remove all spaces and '<', '>', '"', '\''
        s := strings.TrimFunc(
			string(self.buf[:len(self.buf)-1]),
			func(r rune) bool {
				if b := unicode.IsSpace(r); b {
					return true
				}
				switch r {
				case '<', '>', '"', '\'':
					return true
				}
				return false
			})

        // Build result
        result, _ := yaccToken.NewLexemStringValue(yaccToken.HashInclude, self.TokenName, s, ctx, self.DefinitionData.validBlock())

        // load file that came from the include statement
        self.loadFile(ctx, s)

        // Return result
        return result, nil
    }
<INITIAL,datablock>"#"ifdef
    {
        self.startCond = IfdefExpression
    }
<INITIAL,datablock>"#"ifndef
    {
        self.startCond = IfndefExpression
    }
<INITIAL,datablock>"#"define
    {
        self.prevCond = self.startCond
        self.startCond = DefineExpression
    }
<INITIAL,datablock>"#"undefine
    {
        self.prevCond = self.startCond
        self.startCond = UnDefineExpression
    }
<DefineExpression>(.*)\n
    {
        self.startCond = self.prevCond
        s := strings.TrimSpace(string(self.buf[:len(self.buf)-1]))
        self.DefinitionData.setDef(s, true)

        return yaccToken.NewLexemStringValue(yaccToken.HashDefine, self.TokenName,s, ctx, self.DefinitionData.validBlock())
    }
<UnDefineExpression>(.*)\n
    {
        self.startCond = self.prevCond
        s := strings.TrimSpace(string(self.buf[:len(self.buf)-1]))
        self.DefinitionData.setDef(s, false)

        return yaccToken.NewLexemStringValue(yaccToken.HashUnDefine, self.TokenName,s, ctx, self.DefinitionData.validBlock())
    }
<IfdefExpression>(.*)\n
    {
        self.DefinitionData.ExpressionBlockCount++
        self.startCond = DataBlock
        s := strings.TrimSpace(string(self.buf[:len(self.buf)-1]))
        self.DefinitionData.addIfDefStack(s, true)

        return yaccToken.NewLexemStringValue(yaccToken.Hashifdef, self.TokenName,s, ctx, self.DefinitionData.validBlock())
    }
<IfndefExpression>(.*)\n
    {
        self.DefinitionData.ExpressionBlockCount++
        self.startCond = DataBlock
        s := strings.TrimSpace(string(self.buf[:len(self.buf)-1]))
        self.DefinitionData.addIfDefStack(s, false)

        return yaccToken.NewLexemStringValue(yaccToken.Hashifndef, self.TokenName,s, ctx, self.DefinitionData.validBlock())
    }
<datablock>"#"endif
    {
        self.DefinitionData.ExpressionBlockCount--
        if self.DefinitionData.ExpressionBlockCount == 0{
            self.startCond = INITIAL
        }else{
            self.startCond = DataBlock
        }
        self.DefinitionData.removeIfDefStack()

        return yaccToken.NewLexemNoValue(yaccToken.Hashendif, self.TokenName, ctx,self.DefinitionData.validBlock())
    }
<datablock>"#"else
    {
        self.DefinitionData.switchBlock()
        return yaccToken.NewLexemNoValue(yaccToken.Hashelse, self.TokenName, ctx, self.DefinitionData.validBlock())
    }
<INITIAL,datablock>[1-9][0-9]*|0
    {
        v, err := strconv.ParseInt(string(self.buf), 10, 64)
         if err != nil {
            log.Fatal(err)
        }
        return yaccToken.NewLexemIntValue(yaccToken.Integer_literal, self.TokenName, v, ctx, self.DefinitionData.validBlock())
    }
<INITIAL,datablock>0[xX][0-9A-fa-f]+
    {
        v, err := strconv.ParseInt(string(self.buf[2:]), 16, 64)
         if err != nil {
            log.Fatal(err)
        }
        return yaccToken.NewLexemIntValue(yaccToken.Integer_literal, self.TokenName, v, ctx, self.DefinitionData.validBlock())
    }
<INITIAL,datablock>{F}
    {
        v, err := strconv.ParseFloat(string(self.buf), 64)
         if err != nil {
            log.Fatal(err)
        }
        return yaccToken.NewLexemFloatValue(yaccToken.Floating_pt_literal, self.TokenName, v, ctx, self.DefinitionData.validBlock())
    }
<datablock,INITIAL>L?\"(\\.|[^\\"])*\"
    {
        s := strings.TrimSpace(string(self.buf[1:len(self.buf)-1]))
        return yaccToken.NewLexemStringValue(yaccToken.String_literal, self.TokenName, s, ctx, self.DefinitionData.validBlock())
    }
<INITIAL,datablock>[ \t\r\n]+       return yaccToken.NewLexemStringValue(yaccToken.WhiteSpace, self.TokenName, string(self.buf), ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>"//".*\n         return yaccToken.NewLexemStringValue(yaccToken.SingleComment, self.TokenName, string(self.buf), ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>abstract         return yaccToken.NewLexemNoValue(yaccToken.RWabstract, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>any              return yaccToken.NewLexemNoValue(yaccToken.RWany, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>alias            return yaccToken.NewLexemNoValue(yaccToken.RWalias, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>attribute        return yaccToken.NewLexemNoValue(yaccToken.RWattribute, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>bitfield         return yaccToken.NewLexemNoValue(yaccToken.RWbitfield, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>bitmask          return yaccToken.NewLexemNoValue(yaccToken.RWbitmask, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>bitset           return yaccToken.NewLexemNoValue(yaccToken.RWbitset, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>boolean          return yaccToken.NewLexemNoValue(yaccToken.RWboolean, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>case             return yaccToken.NewLexemNoValue(yaccToken.RWcase, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>char             return yaccToken.NewLexemNoValue(yaccToken.RWchar, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>component        return yaccToken.NewLexemNoValue(yaccToken.RWcomponent, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>connector        return yaccToken.NewLexemNoValue(yaccToken.RWconnector, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>const            return yaccToken.NewLexemNoValue(yaccToken.RWconst, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>consumes         return yaccToken.NewLexemNoValue(yaccToken.RWconsumes, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>context          return yaccToken.NewLexemNoValue(yaccToken.RWcontext, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>custom           return yaccToken.NewLexemNoValue(yaccToken.RWcustom, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>default          return yaccToken.NewLexemNoValue(yaccToken.RWdefault, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>double           return yaccToken.NewLexemNoValue(yaccToken.RWdouble, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>exception        return yaccToken.NewLexemNoValue(yaccToken.RWexception, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>emits            return yaccToken.NewLexemNoValue(yaccToken.RWemits, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>enum             return yaccToken.NewLexemNoValue(yaccToken.RWenum, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>eventtype        return yaccToken.NewLexemNoValue(yaccToken.RWeventtype, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>factory          return yaccToken.NewLexemNoValue(yaccToken.RWfactory, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>FALSE            return yaccToken.NewLexemNoValue(yaccToken.RWFALSE, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>finder           return yaccToken.NewLexemNoValue(yaccToken.RWfinder, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>fixed            return yaccToken.NewLexemNoValue(yaccToken.RWfixed, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>float            return yaccToken.NewLexemNoValue(yaccToken.RWfloat, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>getraises        return yaccToken.NewLexemNoValue(yaccToken.RWgetraises, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>home             return yaccToken.NewLexemNoValue(yaccToken.RWhome, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>import           return yaccToken.NewLexemNoValue(yaccToken.RWimport, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>in               return yaccToken.NewLexemNoValue(yaccToken.RWin, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>inout            return yaccToken.NewLexemNoValue(yaccToken.RWinout, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>interface        return yaccToken.NewLexemNoValue(yaccToken.RWinterface, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>local            return yaccToken.NewLexemNoValue(yaccToken.RWlocal, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>long             return yaccToken.NewLexemNoValue(yaccToken.RWlong, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>manages          return yaccToken.NewLexemNoValue(yaccToken.RWmanages, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>map              return yaccToken.NewLexemNoValue(yaccToken.RWmap, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>mirrorport       return yaccToken.NewLexemNoValue(yaccToken.RWmirrorport, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>module           return yaccToken.NewLexemNoValue(yaccToken.RWmodule, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>multiple         return yaccToken.NewLexemNoValue(yaccToken.RWmultiple, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>native           return yaccToken.NewLexemNoValue(yaccToken.RWnative, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>Object           return yaccToken.NewLexemNoValue(yaccToken.RWObject, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>octet            return yaccToken.NewLexemNoValue(yaccToken.RWoctet, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>oneway           return yaccToken.NewLexemNoValue(yaccToken.RWoneway, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>out              return yaccToken.NewLexemNoValue(yaccToken.RWout, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>primarykey       return yaccToken.NewLexemNoValue(yaccToken.RWprimarykey, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>private          return yaccToken.NewLexemNoValue(yaccToken.RWprivate, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>port             return yaccToken.NewLexemNoValue(yaccToken.RWport, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>porttype         return yaccToken.NewLexemNoValue(yaccToken.RWporttype, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>provides         return yaccToken.NewLexemNoValue(yaccToken.RWprovides, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>public           return yaccToken.NewLexemNoValue(yaccToken.RWpublic, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>publishes        return yaccToken.NewLexemNoValue(yaccToken.RWpublishes, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>raises           return yaccToken.NewLexemNoValue(yaccToken.RWraises, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>readonly         return yaccToken.NewLexemNoValue(yaccToken.RWreadonly, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>setraises        return yaccToken.NewLexemNoValue(yaccToken.RWsetraises, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>sequence         return yaccToken.NewLexemNoValue(yaccToken.RWsequence, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>short            return yaccToken.NewLexemNoValue(yaccToken.RWshort, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>string           return yaccToken.NewLexemNoValue(yaccToken.RWstring, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>struct           return yaccToken.NewLexemNoValue(yaccToken.RWstruct, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>supports         return yaccToken.NewLexemNoValue(yaccToken.RWsupports, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>switch           return yaccToken.NewLexemNoValue(yaccToken.RWswitch, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>TRUE             return yaccToken.NewLexemNoValue(yaccToken.RWTRUE, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>truncatable      return yaccToken.NewLexemNoValue(yaccToken.RWtruncatable, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>typedef          return yaccToken.NewLexemNoValue(yaccToken.RWtypedef, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>typeid           return yaccToken.NewLexemNoValue(yaccToken.RWtypeid, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>typename         return yaccToken.NewLexemNoValue(yaccToken.RWtypename, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>typeprefix       return yaccToken.NewLexemNoValue(yaccToken.RWtypeprefix, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>unsigned         return yaccToken.NewLexemNoValue(yaccToken.RWunsigned, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>union            return yaccToken.NewLexemNoValue(yaccToken.RWunion, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>uses             return yaccToken.NewLexemNoValue(yaccToken.RWuses, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>ValueBase        return yaccToken.NewLexemNoValue(yaccToken.RWValueBase, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>valuetype        return yaccToken.NewLexemNoValue(yaccToken.RWvaluetype, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>void             return yaccToken.NewLexemNoValue(yaccToken.RWvoid, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>wchar            return yaccToken.NewLexemNoValue(yaccToken.RWwchar, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>wstring          return yaccToken.NewLexemNoValue(yaccToken.RWwstring, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>int8             return yaccToken.NewLexemNoValue(yaccToken.RWint8, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>uint8            return yaccToken.NewLexemNoValue(yaccToken.RWuint8, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>int16            return yaccToken.NewLexemNoValue(yaccToken.RWint16, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>int32            return yaccToken.NewLexemNoValue(yaccToken.RWint32, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>int64            return yaccToken.NewLexemNoValue(yaccToken.RWint64, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>uint16           return yaccToken.NewLexemNoValue(yaccToken.RWuint16, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>uint32           return yaccToken.NewLexemNoValue(yaccToken.RWuint32, self.TokenName, ctx, self.DefinitionData.validBlock())
<INITIAL,datablock>uint64           return yaccToken.NewLexemNoValue(yaccToken.RWuint64, self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"<"              return yaccToken.NewLexemNoValue(int('<'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>">"              return yaccToken.NewLexemNoValue(int('>'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"{"              return yaccToken.NewLexemNoValue(int('{'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"}"              return yaccToken.NewLexemNoValue(int('}'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"("              return yaccToken.NewLexemNoValue(int('('), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>")"              return yaccToken.NewLexemNoValue(int(')'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"["              return yaccToken.NewLexemNoValue(int('['), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"]"              return yaccToken.NewLexemNoValue(int(']'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"+"              return yaccToken.NewLexemNoValue(int('+'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"-"              return yaccToken.NewLexemNoValue(int('-'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"/"              return yaccToken.NewLexemNoValue(int('/'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"*"              return yaccToken.NewLexemNoValue(int('*'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"%"              return yaccToken.NewLexemNoValue(int('%'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>";"              return yaccToken.NewLexemNoValue(int(';'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>","              return yaccToken.NewLexemNoValue(int(','), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"&"              return yaccToken.NewLexemNoValue(int('&'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"="              return yaccToken.NewLexemNoValue(int('='), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>":"              return yaccToken.NewLexemNoValue(int(':'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"|"              return yaccToken.NewLexemNoValue(int('|'), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"::"             return yaccToken.NewLexemNoValue(yaccToken.ScopeOp, self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>"<<"             return yaccToken.NewLexemNoValue(int(yaccToken.ShlOp), self.TokenName, ctx, self.DefinitionData.validBlock())
<datablock,INITIAL>">>"             return yaccToken.NewLexemNoValue(int(yaccToken.ShrOp), self.TokenName, ctx, self.DefinitionData.validBlock())

<INITIAL,datablock>[@]?[a-zA-Z_][a-zA-Z0-9_]*
        s := strings.TrimSpace(string(self.buf[:len(self.buf)]))
		if data, ok := self.DefinitionData.defined[s]; ok && data != "" {
			self.loadFileFromStringData(s, data)
			return yaccToken.NewLexemStringValue(yaccToken.HashLoadDefinition, self.TokenName, fmt.Sprintf("%s=%s", s, data), ctx, self.DefinitionData.validBlock())
		}
		return yaccToken.NewLexemStringValue(yaccToken.Identifier, self.TokenName, s, ctx, self.DefinitionData.validBlock())
%%
    return yaccToken.NewLexemEofValue(ctx)
}






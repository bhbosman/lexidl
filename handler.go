package lexidl

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/bhbosman/gocommon"
	"github.com/bhbosman/lexpragma"
	yaccToken "github.com/bhbosman/yaccidl"
	"github.com/bhbosman/yaccpragma"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Handler struct {
	fileInformation *yaccToken.FileInformation
	currentContext  *yaccToken.CurrentContext
	DefinitionData  IdlTokenizerDefinitionData
	buf             []byte
	current         byte
	startCond       int
	prevCond        int
	definedFiles    map[string]string
	currentAssigned bool
}

func NewLexIdlHandler(
	fileName string,
	definedDefinitions []string,
	byteReaderCloser gocommon.ByteReaderCloser,
	fileInformation *yaccToken.FileInformation) (*Handler, error) {

	path := filepath.Dir(fileName)
	fileName = filepath.Base(fileName)

	context := yaccToken.NewCurrentContext(
		fileName,
		path,
		byteReaderCloser,
		fileInformation.AddPath(path), nil)

	handler := &Handler{
		fileInformation: fileInformation,
		currentContext:  context,
		DefinitionData:  NewIdlTokenizerDefinitionData(definedDefinitions),
		definedFiles:    make(map[string]string),
	}
	return handler, nil
}

func NewLexIdlHandlerFromFileName(
	fileName string,
	definedDefinitions []string) (*Handler, error) {
	f, e := os.Open(fileName)
	if e != nil {
		return nil, e
	}

	fileInformation := yaccToken.NewFileInformation()
	return NewLexIdlHandler(
		fileName,
		definedDefinitions,
		gocommon.NewByteReaderWithCloser(f, bufio.NewReaderSize(f, 4096)),
		fileInformation)
}

func NewLexIdlHandlerFromData(streamName, s string, definedDefinitions []string) (*Handler, error) {
	fileInformation := yaccToken.NewFileInformation()
	return NewLexIdlHandler(
		streamName,
		definedDefinitions, gocommon.NewByteReaderNoCloser(bytes.NewBufferString(s)),
		fileInformation)
}

func (self *Handler) initStream() {
	self.currentAssigned = false
	self.current = 0
}

func (self *Handler) Close() error {
	if self.currentContext.ReaderCloser != nil {
		return self.currentContext.ReaderCloser.Close()
	}
	return nil
}

func (self *Handler) GetChar() byte {
	if self.current != 0 {
		self.buf = append(self.buf, self.current)
	}
	self.current = 0

	if self.currentContext.UnusedByteAssigned {
		self.current = self.currentContext.UnusedByte
		self.currentContext.UnusedByteAssigned = false
	} else {
		var b byte
		var e error
		if b, e = self.currentContext.ReaderCloser.ReadByte(); e == nil {
			self.current = b
		}
		if e == nil {
			if b == byte('\n') {
				self.currentContext.Row++
				self.currentContext.Col = 0
			} else {
				self.currentContext.Col++
			}
		}
	}
	return self.current
}

type Error struct {
	s string
}

func NewError(s string) *Error {
	return &Error{s: s}
}

func (e Error) Error() string {
	return e.s
}

func createError(s string) error {
	return NewError(s)
}

func (self *Handler) loadFile(ctx yaccToken.CurrentContext, fileName string) {
	s, ok := self.definedFiles[fileName]
	if ok {
		self.loadFileFromStringData(fileName, s)
	} else {
		self.loadFileFromResource(ctx, fileName)
	}
}

func (self *Handler) closeTopStream() {
	if self.currentContext.ReaderCloser != nil {
		temp := self.currentContext.ReaderCloser
		defer temp.Close()
		if next := self.currentContext.Next; next != nil {
			//temp.SetNext(nil)
			self.currentContext = next
			self.current = 0
			self.currentAssigned = false
		} else {
			self.currentContext = nil
		}
	}
}

func (self *Handler) AddPredefinedFile(fileName string, stream string) {
	self.definedFiles[fileName] = stream
}

func (self *Handler) ReadLexem() (*yaccToken.LexemValue, error) {
	for {
		lexem, err := self.readLexem(*self.currentContext)
		if err != nil {
			return nil, err
		}
		if !lexem.Eof {
			return lexem, nil
		}
		self.closeTopStream()
		if self.currentContext == nil || self.currentContext.ReaderCloser == nil {
			return lexem, nil
		}
	}
}

func (self *Handler) loadFileFromStringData(fileName string, s string) {
	if self.currentAssigned {
		self.currentContext.SetUnusedByte(self.current)
	}

	if !filepath.IsAbs(fileName) {
		usr, _ := user.Current()
		homeDir := usr.HomeDir
		if strings.HasPrefix(fileName, "~/") {
			fileName = filepath.Join(homeDir, fileName[2:])
		}
	}
	path := filepath.Dir(fileName)
	self.currentContext = yaccToken.NewCurrentContext(
		fileName,
		path,
		gocommon.NewByteReaderNoCloser(bytes.NewBufferString(s)),

		self.fileInformation.AddPath(path),
		self.currentContext)
	self.initStream()
}
func (self Handler) TokenName(tokenId int) string {
	if tokenId < 255 {
		return string(byte(tokenId))
	}

	return yaccToken.YaccIdlTokname(yaccToken.YaccIdlTok2[tokenId-yaccToken.YaccIdlPrivate])
}

func (self *Handler) loadFileFromResource(ctx yaccToken.CurrentContext, fileName string) {
	if self.currentAssigned {
		self.currentContext.SetUnusedByte(self.current)
	}
	absFileName := filepath.Join(ctx.Path, fileName)

	stat, e := os.Stat(absFileName)
	if e != nil {
		panic(e)
	}
	if e != nil {
		panic(e)
	}
	if stat.IsDir() {
		panic(NewError("can not open folder"))
	}
	f, e := os.Open(absFileName)
	if e != nil {
		panic(e)
	}
	path := filepath.Dir(absFileName)

	fileName = filepath.Join(filepath.Dir(self.currentContext.FileName), fileName)

	self.currentContext = yaccToken.NewCurrentContext(
		fileName,
		path,
		gocommon.NewByteReaderWithCloser(f, bufio.NewReader(f)),
		self.fileInformation.AddPath(path),
		self.currentContext)
	self.initStream()
}

func (self *Handler) ReadPragmaExpression(ctx yaccToken.CurrentContext, s string) (*yaccToken.LexemValue, error) {
	s = fmt.Sprintf("#pragma %v", s)
	handler, _ := lexpragma.NewPragmaLexFromData("(pragma stream)", s)
	lexer := NewYaccPragmaLexer(handler)
	parser := yaccpragma.YaccPragmaNewParser()
	v := parser.Parse(lexer)
	if v != 0 {
		return nil, fmt.Errorf("error occured when reading pragma statement. Error: %v", lexer.ErrorMessage())
	}

	if lexer.PragmaNode() == nil {
		return nil, fmt.Errorf("error in pragma statement. no retuned value")
	}
	return yaccToken.NewLexemPragmaNodeValue(yaccToken.HashPragma, self.TokenName, lexer.PragmaNode(), ctx, self.DefinitionData.validBlock())
}

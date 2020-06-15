package lexidl

import "strings"

type IdlTokenizerDefinitionData struct {
	defined              map[string]string
	expressionCountStack []bool
	ExpressionBlockCount int
}

func NewIdlTokenizerDefinitionData(definedDefinitions []string) IdlTokenizerDefinitionData {
	result := IdlTokenizerDefinitionData{
		defined: make(map[string]string),
	}
	for _, d := range definedDefinitions {
		result.defined[d] = d
	}
	return result
}

func (self *IdlTokenizerDefinitionData) setDef(def string, b bool) {
	if self.validBlock() {
		def = strings.TrimSpace(def)
		index := strings.Index(def, " ")
		var key string
		var value string

		if index == -1 {
			key = def
			value = ""
		} else {
			key = strings.TrimSpace(def[:index])
			value = strings.TrimSpace(def[index:])
		}

		if b {
			self.defined[key] = value
		} else {
			delete(self.defined, def)
		}
	}
}

func (self *IdlTokenizerDefinitionData) isDefined(s string) (bool, string) {
	s = strings.TrimSpace(s)
	v, ok := self.defined[s]
	return ok, v
}

func (self *IdlTokenizerDefinitionData) validBlock() bool {
	if self.expressionCountStack == nil {
		return true
	}
	if len(self.expressionCountStack) == 0 {
		return true
	}
	result := true

	for _, b := range self.expressionCountStack {
		result = result && b
		if !result {
			return false
		}
	}
	return true
}

func (self *IdlTokenizerDefinitionData) addIfDefStack(s string, b bool) {
	s = strings.TrimSpace(s)
	_, ok := self.defined[s]
	if !b {
		ok = !ok
	}
	self.expressionCountStack = append(self.expressionCountStack, ok)
}

func (self *IdlTokenizerDefinitionData) removeIfDefStack() {
	self.expressionCountStack = self.expressionCountStack[0 : len(self.expressionCountStack)-1]
}

func (self *IdlTokenizerDefinitionData) switchBlock() {
	if self.expressionCountStack == nil {
		panic(createError("invalid code path"))
	}
	if len(self.expressionCountStack) == 0 {
		panic(createError("invalid code path"))
	}
	v := self.expressionCountStack[len(self.expressionCountStack)-1]
	self.expressionCountStack[len(self.expressionCountStack)-1] = !v
}

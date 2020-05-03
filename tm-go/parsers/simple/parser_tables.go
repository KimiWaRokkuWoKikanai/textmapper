// generated by Textmapper; DO NOT EDIT

package simple

import (
	"fmt"
)

var tmNonterminals = [...]string{
	"input",
}

func symbolName(sym int32) string {
	if sym < int32(NumTokens) {
		return Token(sym).String()
	}
	if i := int(sym) - int(NumTokens); i < len(tmNonterminals) {
		return tmNonterminals[i]
	}
	return fmt.Sprintf("nonterminal(%d)", sym)
}

var tmAction = []int32{
	-1, 0, -1, -2,
}

var tmGoto = []int32{
	0, 2, 2, 4, 6,
}

var tmFromTo = []int8{
	2, 3, 0, 1, 0, 2,
}

var tmRuleLen = []int8{
	1,
}

var tmRuleSymbol = []int32{
	3,
}

var tmRuleType = [...]NodeType{
	0, // input : 'simple'
}

package gen

type file struct {
	name     string
	template string
}

var lexerFiles = []file{
	{"token.go", tokenTpl},
	{"lexer_tables.go", lexerTablesTpl},
	{"lexer.go", lexerTpl},
}

var parserFiles = []file{
	{"parser_tables.go", parserTablesTpl},
}

const sharedDefs = `
{{define "header" -}}
// generated by Textmapper; DO NOT EDIT

{{end}}

`

const tokenTpl = `
{{- template "header" . -}}
package {{.Name}}

// Token is an enum of all terminal symbols of the {{.Name}} language.
type Token int

// Token values.
const (
	UNAVAILABLE Token = iota - 1
{{- range .Tokens}}
	{{.ID}}{{if .Comment}}  // {{.Comment}}{{end}}
{{- end}}

	NumTokens
)

var tokenStr = [...]string{
{{- range .Tokens}}
	{{if .Comment}}{{str_literal .Comment}}{{else}}{{str_literal .ID}}{{end}},
{{- end}}
}

func (tok Token) String() string {
	if tok >= 0 && int(tok) < len(tokenStr) {
		return tokenStr[tok]
	}
	return "fmt".Sprintf("token(%d)", tok)
}
`

const lexerTablesTpl = `
{{- template "header" . -}}
package {{.Name}}

const tmNumClasses = {{.Lexer.Tables.NumSymbols}}

{{$runeType := bits .Lexer.Tables.NumSymbols -}}
{{if gt .Lexer.Tables.LastMapEntry.Start 2048 -}}
type mapRange struct {
	lo         rune
	hi         rune
	defaultVal uint{{$runeType}}
	val        []uint{{$runeType}}
}

func mapRune(c rune) int {
	lo := 0
	hi := len(tmRuneRanges)
	for lo < hi {
		m := lo + (hi-lo)/2
		r := tmRuneRanges[m]
		if c < r.lo {
			hi = m
		} else if c >= r.hi {
			lo = m + 1
		} else {
			i := int(c - r.lo)
			if i < len(r.val) {
				return int(r.val[i])
			}
			return int(r.defaultVal)
		}
	}
	return 1
}

// Latin-1 characters.
var tmRuneClass = []uint{{$runeType}}{
{{- int_array (.Lexer.Tables.SymbolArr 256) "\t" 79 -}}
}

const tmRuneClassLen = 256
const tmFirstRule = {{.Lexer.Tables.ActionStart}}

var tmRuneRanges = []mapRange{
{{range .Lexer.Tables.CompressedMap 256}}	{ {{- .Lo}}, {{.Hi}}, {{.DefaultVal}}, {{if .Vals}}[]uint{{$runeType}}{
{{- int_array .Vals "\t\t" 78}}	}{{else}}nil{{end -}} },
{{end -}}
}

{{else -}}
{{ $runeArr := .Lexer.Tables.SymbolArr 0 -}}
var tmRuneClass = []uint{{$runeType}}{
{{- int_array $runeArr "\t" 79 -}}
}

const tmRuneClassLen = {{len $runeArr}}
const tmFirstRule = {{.Lexer.Tables.ActionStart}}

{{end -}}
var tmStateMap = []int{
{{- int_array .Lexer.Tables.StateMap "\t" 79 -}}
}

{{if .Lexer.RuleToken -}}
var tmToken = []Token{
{{- int_array .Lexer.RuleToken "\t" 79 -}}
}

{{end -}}
var tmLexerAction = []int{{bits_per_element .Lexer.Tables.Dfa}}{
{{- int_array .Lexer.Tables.Dfa "\t" 79 -}}
}

{{- if .Lexer.Tables.Backtrack}}

var tmBacktracking = []int{
{{- range .Lexer.Tables.Backtrack}}
	{{.Action}}, {{.NextState}},{{if .Details}} // {{.Details}}{{end}}
{{- end}}
}
{{- end}}
`

const lexerTpl = `
{{- template "header" . -}}
package {{.Name}}

{{- if gt (len .Lexer.StartConditions) 1}}

// Lexer states.
const (
{{- range $index, $el := .Lexer.StartConditions}}
	State{{title .}} = {{$index}}
{{- end}}
)
{{- end}}
{{block "onBeforeLexer" .}}{{end}}
{{template "lexerType" .}}
{{template "lexerInit" .}}
{{template "lexerNext" .}}
{{template "lexerPos" .}}
{{- if .Options.TokenLine}}
{{template "lexerLine" .}}
{{- end}}
{{- if .Options.TokenColumn}}
{{template "lexerColumn" .}}
{{- end}}
{{template "lexerText" .}}
{{template "lexerValue" .}}
{{template "lexerRewind" .}}
{{- block "onAfterLexer" .}}{{end}}

{{- define "lexerType" -}}
// Lexer uses a generated DFA to scan through a utf-8 encoded input string. If
// the string starts with a BOM character, it gets skipped.
type Lexer struct {
	source string

	ch          rune // current character, -1 means EOI
	offset      int  // character offset
	tokenOffset int  // last token offset
{{- if .Options.TokenLine}}
	line        int  // current line number (1-based)
	tokenLine   int  // last token line
{{- end}}
{{- if or .Options.TokenLineOffset .Options.TokenColumn}}
	lineOffset  int  // current line offset
{{- end}}
{{- if .Options.TokenColumn}}
	tokenColumn int  // last token column (in bytes)
{{- end}}
	scanOffset  int  // scanning offset
	value       interface{}

	State int // lexer state, modifiable
{{- block "stateVars" .}}{{end}}
}
{{end -}}

{{- define "lexerInit" -}}
var bomSeq = "\xef\xbb\xbf"

// Init prepares the lexer l to tokenize source by performing the full reset
// of the internal state.
func (l *Lexer) Init(source string) {
	l.source = source

	l.ch = 0
	l.offset = 0
	l.tokenOffset = 0
{{- if .Options.TokenLine}}
	l.line = 1
	l.tokenLine = 1
{{- end}}
{{- if or .Options.TokenLineOffset .Options.TokenColumn}}
	l.lineOffset = 0
{{- end}}
{{- if .Options.TokenColumn}}
	l.tokenColumn = 1
{{- end}}
	l.State = 0
{{- block "initStateVars" .}}{{end}}

	if "strings".HasPrefix(source, bomSeq) {
		l.offset += len(bomSeq)
	}

	l.rewind(l.offset)
}
{{end -}}

{{- define "lexerNext" -}}
// Next finds and returns the next token in l.source. The source end is
// indicated by Token.EOI.
//
// The token text can be retrieved later by calling the Text() method.
func (l *Lexer) Next() Token {
{{- block "onBeforeNext" .}}{{end}}
{{- $spaceRules := .SpaceActions}}
{{- if or $spaceRules .Lexer.RuleToken }}
restart:
{{- end}}
{{- if .Options.TokenLine}}
	l.tokenLine = l.line
{{- end}}
{{- if .Options.TokenColumn}}
	l.tokenColumn = l.offset-l.lineOffset+1
{{- end}}
	l.tokenOffset = l.offset

	state := tmStateMap[l.State]
{{- if .Lexer.ClassActions}}
	hash := uint32(0)
{{- end}}
{{- if .Lexer.Tables.Backtrack}}
	backup{{if .Lexer.RuleToken}}Rule{{else}}Token{{end}} := -1
	var backupOffset int
{{- if .Lexer.ClassActions}}
	backupHash := hash
{{- end}}
{{- end}}
	for state >= 0 {
		var ch int
		if uint(l.ch) < tmRuneClassLen {
			ch = int(tmRuneClass[l.ch])
		} else if l.ch < 0 {
			state = int(tmLexerAction[state*tmNumClasses])
			continue
		} else {
{{- if gt .Lexer.Tables.LastMapEntry.Start 2048}}
			ch = mapRune(l.ch)
{{- else}}
			ch = 1
{{- end}}
		}
		state = int(tmLexerAction[state*tmNumClasses+ch])
		if state > tmFirstRule {
{{- if .Lexer.Tables.Backtrack}}
			if state < 0 {
				state = (-1 - state) * 2
				backup{{if .Lexer.RuleToken}}Rule{{else}}Token{{end}} = tmBacktracking[state]
				backupOffset = l.offset
{{- if .Lexer.ClassActions}}
				backupHash = hash
{{- end}}
				state = tmBacktracking[state+1]
			}
{{- end}}
{{- if .Lexer.ClassActions}}
			hash = hash*uint32(31) + uint32(l.ch)
{{end}}
{{- if .Options.TokenLine}}
			if l.ch == '\n' {
				l.line++
{{- if or .Options.TokenLineOffset .Options.TokenColumn}}
				l.lineOffset = l.offset
{{- end}}
			}
{{end}}
			// Scan the next character.
			// Note: the following code is inlined to avoid performance implications.
			l.offset = l.scanOffset
			if l.offset < len(l.source) {
				r, w := rune(l.source[l.offset]), 1
				if r >= 0x80 {
					// not ASCII
					r, w = "unicode/utf8".DecodeRuneInString(l.source[l.offset:])
				}
				l.scanOffset += w
				l.ch = r
			} else {
				l.ch = -1 // EOI
			}
		}
	}
{{if .Lexer.RuleToken}}
	rule := tmFirstRule - state
{{- else}}
	token := Token(tmFirstRule - state)
{{- end}}
{{- if .Lexer.Tables.Backtrack}}
recovered:
{{- end}}
{{- if .Lexer.ClassActions}}
	switch {{if .Lexer.RuleToken}}rule{{else}}token{{end}} {
{{- range .Lexer.ClassActions}}
{{- if $.Lexer.RuleToken}}
	case {{sum .Action 2}}:
{{- else}}
	case {{(index $.Syms .Action).ID}}:
{{- end}}
{{- with string_switch .Custom }}
		hh := hash & {{.Mask}}
		switch hh {
{{- range .Cases}}
		case {{.Value}}:
{{- range .Subcases}}
			if hash == {{hex .Hash}} && {{quote .Str}} == l.source[l.tokenOffset:l.offset] {
{{- if $.Lexer.RuleToken}}
				rule = {{sum .Action 2}}
{{- else}}
				token = {{(index $.Syms .Action).ID}}
{{- end}}
				break
			}
{{- end}}
{{- end}}
		}
{{- end}}
{{- end}}
	}
{{- end}}
{{- if .Lexer.RuleToken}}

	token := tmToken[rule]
	var space bool
{{- if .Lexer.Actions}}
	switch rule {
	case 0:
{{- template "handleInvalidToken" .}}
{{- range .Lexer.Actions}}
	case {{sum .Action 2}}:{{if .Comments}} // {{join .Comments ", "}}{{end}}
{{- if .Space }}
		space = true
{{- end}}
{{- if .Code }}
{{lexer_action .Code}}
{{- end}}
{{- end}}
	}
{{- else}}
	if rule == 0 {
{{- template "handleInvalidToken" .}}
	}
{{- end}}
	if space {
		goto restart
	}
{{- else}}
	switch token {
	case {{(index $.Syms .Lexer.InvalidToken).ID}}:
{{- template "handleInvalidToken" .}}
{{- if $spaceRules}}
	case {{range $i, $val := $spaceRules}}{{if gt $i 0}}, {{end}}{{$val}}{{end}}:
		goto restart
{{- end}}
	}
{{- end}}
{{- block "onAfterNext" .}}{{end}}
	return token
}
{{end -}}

{{- define "lexerPos" -}}
// Pos returns the start and end positions of the last token returned by Next().
func (l *Lexer) Pos() (start, end int) {
	start = l.tokenOffset
	end = l.offset
	return
}
{{end -}}

{{- define "lexerLine" -}}
// Line returns the line number of the last token returned by Next() (1-based).
func (l *Lexer) Line() int {
	return l.tokenLine
}
{{end -}}

{{- define "lexerColumn" -}}
// Column returns the column of the last token returned by Next() (in bytes, 1-based).
func (l *Lexer) Column() int {
	return l.tokenColumn
}
{{end -}}

{{- define "lexerText" -}}
// Text returns the substring of the input corresponding to the last token.
func (l *Lexer) Text() string {
	return l.source[l.tokenOffset:l.offset]
}
{{end -}}

{{- define "lexerValue" -}}
// Value returns the value associated with the last returned token.
func (l *Lexer) Value() interface{} {
	return l.value
}
{{end -}}

{{- define "lexerRewind" -}}
// rewind can be used in lexer actions to accept a portion of a scanned token, or to include
// more text into it.
func (l *Lexer) rewind(offset int) {
{{- if .Options.TokenLine}}
	if offset < l.offset {
		l.line -= "strings".Count(l.source[offset:l.offset], "\n")
	} else {
		if offset > len(l.source) {
			offset = len(l.source)
		}
		l.line += "strings".Count(l.source[l.offset:offset], "\n")
	}
{{- if or .Options.TokenLineOffset .Options.TokenColumn}}
	l.lineOffset = 1 + "strings".LastIndexByte(l.source[:offset], '\n')
{{- end}}
{{end}}
	// Scan the next character.
	l.scanOffset = offset
	l.offset = offset
	if l.offset < len(l.source) {
		r, w := rune(l.source[l.offset]), 1
		if r >= 0x80 {
			// not ASCII
			r, w = "unicode/utf8".DecodeRuneInString(l.source[l.offset:])
		}
		l.scanOffset += w
		l.ch = r
	} else {
		l.ch = -1 // EOI
	}
}
{{end -}}

{{- define "handleInvalidToken" -}}
{{if .Lexer.Tables.Backtrack}}
		if backup{{if .Lexer.RuleToken}}Rule{{else}}Token{{end}} >= 0 {
{{- if .Lexer.RuleToken}}
			rule = backupRule
{{- else}}
			token = Token(backupToken)
{{- end}}
{{- if .Lexer.ClassActions}}
			hash = backupHash
{{- end}}
			l.rewind(backupOffset)
		} else if l.offset == l.tokenOffset {
			l.rewind(l.scanOffset)
		}
{{- if .Lexer.RuleToken}}
		if rule != 0 {
{{- else}}
		if token != {{(index $.Syms .Lexer.InvalidToken).ID}} {
{{- end}}
			goto recovered
		}
{{- else}}
		if l.offset == l.tokenOffset {
			l.rewind(l.scanOffset)
		}
{{- end -}}
{{end -}}
`

const bisonTpl = `%{
%}
{{range .Parser.Inputs}}
%start {{(index $.Parser.Nonterms .Nonterm).Name}}{{if .NoEoi}} // no-eoi{{end}}
{{- end}}
{{range .Parser.Prec}}
%{{.Associativity}}{{range .Terminals}} {{(index $.Syms .).ID}}{{end}}
{{- end}}
{{- range slice .Tokens 1}}
%token {{.ID}}
{{- end}}

%%
{{- range .Parser.Nonterms}}

{{.Name}} :
{{- if eq .Value.Kind 2 }}
{{- range $i, $rule := .Value.Sub}}
{{ if eq $i 0}}  {{else}}| {{end}}{{$.ExprString $rule}}
{{- end}}
{{- else }}
  {{$.ExprString .Value}}
{{- end }}
;
{{- end}}

%%

`

const parserTablesTpl = `
{{- template "header" . -}}
package {{.Name}}

{{- range .Parser.Tables.Markers}}
{{if eq (len .States) 1}}
const {{.Name}}State = {{index .States 0}}
{{- else}}
var {{.Name}}States = map[int]bool{
{{- range .States}}
	{{.}}: true,
{{- end}}
}
{{- end}}
{{- end}}

var tmNonterminals = [...]string{
{{- range .Parser.Nonterms}}
	"{{.Name}}",
{{- end}}
}

func symbolName(sym int32) string {
	if sym == noToken {
		return "<no-token>"
	}
	if sym < int32({{ref "NumTokens"}}) {
		return {{ref "Token"}}(sym).String()
	}
	if i := int(sym) - int({{ref "NumTokens"}}); i < len(tmNonterminals) {
		return tmNonterminals[i]
	}
	return "fmt".Sprintf("nonterminal(%d)", sym)
}

var tmAction = []int32{
{{- int_array .Parser.Tables.Action "\t" 79 -}}
}
{{- if .Parser.Tables.Lalr}}

var tmLalr = []int32{
{{- int_array .Parser.Tables.Lalr "\t" 79 -}}
}
{{- end}}

var tmGoto = []int32{
{{- int_array .Parser.Tables.Goto "\t" 79 -}}
}

{{$stateType := bits_per_element .Parser.Tables.FromTo -}}
var tmFromTo = []int{{$stateType}}{
{{- int_array .Parser.Tables.FromTo "\t" 79 -}}
}

var tmRuleLen = []int{{bits_per_element .Parser.Tables.RuleLen}}{
{{- int_array .Parser.Tables.RuleLen "\t" 79 -}}
}

var tmRuleSymbol = []int32{
{{- int_array .Parser.Tables.RuleSymbol "\t" 79 -}}
}

var tmRuleType = [...]{{ref "NodeType"}}{
{{- range .Parser.Rules}}
{{- if ne .Type -1 }}
{{- $val := index $.Parser.RangeTypes .Type }}
	{{if ne $val $.Options.FileNode}}{{$val}}{{else}}0{{end}}, // {{$.RuleString .}}
{{- else }}
	0, // {{$.RuleString .}}
{{- end}}
{{- end}}
}

{{- range .Sets}}

// {{.Expr}} = {{.ValueString $}}
var {{.Name}} = []int32{
{{- if gt (len .Terminals) 0}}
{{- int_array .Terminals "\t" 79 -}}
{{- end}}
}
{{- end}}
`

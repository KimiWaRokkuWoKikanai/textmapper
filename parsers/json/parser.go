// generated by Textmapper; DO NOT EDIT

package json

import (
	"fmt"

	"github.com/inspirer/textmapper/parsers/json/token"
)

// Parser is a table-driven LALR parser for json.
type Parser struct {
	listener Listener

	next symbol

	// Tokens to be reported with the next shift. Only non-empty when next.symbol != noToken.
	pending []symbol
}

type SyntaxError struct {
	Line      int
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line %v", e.Line)
}

type symbol struct {
	symbol    int32
	offset    int
	endoffset int
}

type stackEntry struct {
	sym   symbol
	state int8
	value interface{}
}

func (p *Parser) Init(l Listener) {
	p.listener = l
	if cap(p.pending) < startTokenBufferSize {
		p.pending = make([]symbol, 0, startTokenBufferSize)
	}
}

const (
	startStackSize       = 256
	startTokenBufferSize = 16
	noToken              = int32(token.UNAVAILABLE)
	eoiToken             = int32(token.EOI)
	debugSyntax          = false
)

func (p *Parser) Parse(lexer *Lexer) error {
	return p.parse(0, 44, lexer)
}

func (p *Parser) parse(start, end int8, lexer *Lexer) error {
	p.pending = p.pending[:0]
	state := start

	var alloc [startStackSize]stackEntry
	stack := append(alloc[:0], stackEntry{state: state})
	p.fetchNext(lexer, stack)

	for state != end {
		action := tmAction[state]
		if action > tmActionBase {
			// Lookahead is needed.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			pos := action + p.next.symbol
			if pos >= 0 && pos < tmTableLen && int32(tmCheck[pos]) == p.next.symbol {
				action = int32(tmTable[pos])
			} else {
				action = tmDefAct[state]
			}
		} else {
			action = tmDefAct[state]
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			rhs := stack[len(stack)-ln:]
			stack = stack[:len(stack)-ln]
			if ln == 0 {
				if p.next.symbol == noToken {
					p.fetchNext(lexer, stack)
				}
				entry.sym.offset, entry.sym.endoffset = p.next.offset, p.next.offset
			} else {
				entry.sym.offset = rhs[0].sym.offset
				entry.sym.endoffset = rhs[ln-1].sym.endoffset
			}
			if err := p.applyRule(rule, &entry, rhs, lexer); err != nil {
				return err
			}
			if debugSyntax {
				fmt.Printf("reduced to: %v\n", symbolName(entry.sym.symbol))
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action < -1 {
			// Shift.
			state = int8(-2 - action)
			stack = append(stack, stackEntry{
				sym:   p.next,
				state: state,
				value: lexer.Value(),
			})
			if debugSyntax {
				fmt.Printf("shift: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
			}
			p.flush(p.next)
			if p.next.symbol != eoiToken {
				switch token.Type(p.next.symbol) {
				case token.JSONSTRING:
					p.listener(JSONString, p.next.offset, p.next.endoffset)
				}
				p.next.symbol = noToken
			}
		}

		if action == -1 || state == -1 {
			break
		}
	}

	if state != end {
		if p.next.symbol == noToken {
			p.fetchNext(lexer, stack)
		}
		err := SyntaxError{
			Line:      lexer.Line(),
			Offset:    p.next.offset,
			Endoffset: p.next.endoffset,
		}
		return err
	}

	return nil
}

func gotoState(state int8, symbol int32) int8 {
	const numTokens = 19
	if symbol >= numTokens {
		pos := tmGoto[symbol-numTokens] + int32(state)
		if pos >= 0 && pos < tmTableLen && tmCheck[pos] == int8(state) {
			return int8(tmTable[pos])
		}
		return int8(tmDefGoto[symbol-numTokens])
	}

	// Shifting a token.
	action := tmAction[state]
	if action == tmActionBase {
		return -1
	}
	pos := action + symbol
	if pos >= 0 && pos < tmTableLen && tmCheck[pos] == int8(symbol) {
		action = int32(tmTable[pos])
	} else {
		action = tmDefAct[state]
	}
	if action < -1 {
		return int8(-2 - action)
	}
	return -1
}

func (p *Parser) fetchNext(lexer *Lexer, stack []stackEntry) {
restart:
	tok := lexer.Next()
	switch tok {
	case token.MULTILINECOMMENT, token.INVALID_TOKEN:
		s, e := lexer.Pos()
		tok := symbol{int32(tok), s, e}
		p.pending = append(p.pending, tok)
		goto restart
	}
	p.next.symbol = int32(tok)
	p.next.offset, p.next.endoffset = lexer.Pos()
}

func lookaheadNext(lexer *Lexer) symbol {
restart:
	tok := lexer.Next()
	switch tok {
	case token.MULTILINECOMMENT, token.INVALID_TOKEN:
		goto restart
	}
	s, e := lexer.Pos()
	return symbol{int32(tok), s, e}
}

func AtEmptyObject(lexer *Lexer, next symbol) bool {
	if debugSyntax {
		fmt.Printf("lookahead EmptyObject, next: %v\n", symbolName(next.symbol))
	}
	return lookahead(lexer, next, 1, 43)
}

func lookahead(l *Lexer, next symbol, start, end int8) bool {
	lexer := l.Copy()
	var allocated [64]stackEntry
	state := start
	stack := append(allocated[:0], stackEntry{state: state})

	for state != end {
		action := tmAction[state]
		if action > tmActionBase {
			// Lookahead is needed.
			if next.symbol == noToken {
				next = lookaheadNext(&lexer)
			}
			pos := action + next.symbol
			if pos >= 0 && pos < tmTableLen && int32(tmCheck[pos]) == next.symbol {
				action = int32(tmTable[pos])
			} else {
				action = tmDefAct[state]
			}
		} else {
			action = tmDefAct[state]
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			stack = stack[:len(stack)-ln]
			if debugSyntax {
				fmt.Printf("lookahead reduced to: %v\n", symbolName(entry.sym.symbol))
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action < -1 {
			// Shift.
			state = int8(-2 - action)
			stack = append(stack, stackEntry{
				sym:   next,
				state: state,
			})
			if debugSyntax {
				fmt.Printf("lookahead shift: %v (%s)\n", symbolName(next.symbol), lexer.Text())
			}
			if state != -1 && next.symbol != eoiToken {
				next.symbol = noToken
			}
		}

		if action == -1 || state == -1 {
			break
		}
	}

	if debugSyntax {
		fmt.Printf("lookahead done: %v\n", state == end)
	}
	return state == end
}

func (p *Parser) applyRule(rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer) (err error) {
	switch rule {
	case 32:
		if AtEmptyObject(lexer, p.next) {
			lhs.sym.symbol = 23 /* lookahead_EmptyObject */
		} else {
			lhs.sym.symbol = 25 /* lookahead_notEmptyObject */
		}
		return
	}
	if nt := tmRuleType[rule]; nt != 0 {
		p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}

func (p *Parser) reportIgnoredToken(tok symbol) {
	var t NodeType
	switch token.Type(tok.symbol) {
	case token.MULTILINECOMMENT:
		t = MultiLineComment
	case token.INVALID_TOKEN:
		t = InvalidToken
	default:
		return
	}
	if debugSyntax {
		fmt.Printf("ignored: %v as %v\n", token.Type(tok.symbol), t)
	}
	p.listener(t, tok.offset, tok.endoffset)
}

// flush reports all pending tokens up to a given symbol.
func (p *Parser) flush(sym symbol) {
	if len(p.pending) > 0 && p.listener != nil {
		for i, tok := range p.pending {
			if tok.endoffset > sym.endoffset {
				// Note: this copying should not happen during normal operation, only
				// during error recovery.
				p.pending = append(p.pending[:0], p.pending[i:]...)
				return
			}
			p.reportIgnoredToken(tok)
		}
		p.pending = p.pending[:0]
	}
}

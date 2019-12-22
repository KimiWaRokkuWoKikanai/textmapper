// generated by Textmapper; DO NOT EDIT

package tm

import (
	"fmt"
)

// ErrorHandler is called every time a parser is unable to process some part of the input.
// This handler can return false to abort the parser.
type ErrorHandler func(err SyntaxError) bool

// StopOnFirstError is an error handler that forces the parser to stop on and return the first
// error.
func StopOnFirstError(_ SyntaxError) bool { return false }

// Parser is a table-driven LALR parser for tm.
type Parser struct {
	eh       ErrorHandler
	listener Listener

	next     symbol
	endState int16
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
	state int16
}

func (p *Parser) Init(eh ErrorHandler, l Listener) {
	p.eh = eh
	p.listener = l
}

const (
	startStackSize       = 256
	startTokenBufferSize = 16
	noToken              = int32(UNAVAILABLE)
	eoiToken             = int32(EOI)
	debugSyntax          = false
)

func (p *Parser) Parse(lexer *Lexer) error {
	return p.parse(0, 489, lexer)
}

func (p *Parser) parse(start, end int16, lexer *Lexer) error {
	ignoredTokens := make([]symbol, 0, startTokenBufferSize) // to be reported with the next shift
	state := start
	var lastErr SyntaxError
	recovering := 0

	var alloc [startStackSize]stackEntry
	stack := append(alloc[:0], stackEntry{state: state})
	p.endState = end
	ignoredTokens = p.fetchNext(lexer, stack, ignoredTokens)

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if p.next.symbol == noToken {
				ignoredTokens = p.fetchNext(lexer, stack, ignoredTokens)
			}
			action = lalr(action, p.next.symbol)
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
					ignoredTokens = p.fetchNext(lexer, stack, ignoredTokens)
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
				fmt.Printf("reduced to: %v\n", Symbol(entry.sym.symbol))
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			// Shift.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack, nil)
			}
			state = gotoState(state, p.next.symbol)
			stack = append(stack, stackEntry{
				sym:   p.next,
				state: state,
			})
			if debugSyntax {
				fmt.Printf("shift: %v (%s)\n", Symbol(p.next.symbol), lexer.Text())
			}
			if len(ignoredTokens) > 0 {
				for _, tok := range ignoredTokens {
					p.reportIgnoredToken(tok)
				}
				ignoredTokens = ignoredTokens[:0]
			}
			if state != -1 && p.next.symbol != eoiToken {
				p.next.symbol = noToken
			}
			if recovering > 0 {
				recovering--
			}
		}

		if action == -2 || state == -1 {
			if recovering == 0 {
				if p.next.symbol == noToken {
					ignoredTokens = p.fetchNext(lexer, stack, ignoredTokens)
				}
				lastErr = SyntaxError{
					Line:      lexer.Line(),
					Offset:    p.next.offset,
					Endoffset: p.next.endoffset,
				}
				if !p.eh(lastErr) {
					return lastErr
				}
			}
			if len(ignoredTokens) > 0 {
				for _, tok := range ignoredTokens {
					p.reportIgnoredToken(tok)
				}
				ignoredTokens = ignoredTokens[:0]
			}
			if stack = p.recoverFromError(lexer, stack); stack == nil {
				return lastErr
			}
			state = stack[len(stack)-1].state
			recovering = 4
		}
	}

	return nil
}

const errSymbol = 38

// willShift checks if "symbol" is going to be shifted in the given state.
// This function does not support empty productions and returns false if they occur before "symbol".
func (p *Parser) willShift(stackPos int, state int16, symbol int32, stack []stackEntry) bool {
	if state == -1 {
		return false
	}

	for state != p.endState {
		action := tmAction[state]
		if action < -2 {
			action = lalr(action, symbol)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])
			if ln == 0 {
				// we do not support empty productions
				return false
			}
			stackPos -= ln - 1
			state = gotoState(stack[stackPos-1].state, tmRuleSymbol[rule])
		} else {
			return action == -1 && gotoState(state, symbol) >= 0
		}
	}
	return symbol == eoiToken
}

func (p *Parser) skipBrokenCode(lexer *Lexer, stack []stackEntry, canRecover func(symbol int32) bool) int {
	var e int
	for p.next.symbol != eoiToken && !canRecover(p.next.symbol) {
		e = p.next.endoffset
		p.fetchNext(lexer, stack, nil)
	}
	return e
}

func (p *Parser) recoverFromError(lexer *Lexer, stack []stackEntry) []stackEntry {
	var recoverSyms [1 + NumTokens/8]uint8
	var recoverPos []int

	for size := len(stack); size > 0; size-- {
		if gotoState(stack[size-1].state, errSymbol) == -1 {
			continue
		}
		recoverPos = append(recoverPos, size)
		if recoveryScopeState == stack[size-1].state {
			break
		}
	}
	if len(recoverPos) == 0 {
		return nil
	}

	for _, v := range afterErr {
		recoverSyms[v/8] |= 1 << uint32(v%8)
	}
	canRecover := func(symbol int32) bool {
		return recoverSyms[symbol/8]&(1<<uint32(symbol%8)) != 0
	}
	if p.next.symbol == noToken {
		p.fetchNext(lexer, stack, nil)
	}
	s := p.next.offset
	e := s
	for {
		if endoffset := p.skipBrokenCode(lexer, stack, canRecover); endoffset > e {
			e = endoffset
		}

		var matchingPos int
		for _, pos := range recoverPos {
			if p.willShift(pos, gotoState(stack[pos-1].state, errSymbol), p.next.symbol, stack) {
				matchingPos = pos
				break
			}
		}
		if matchingPos == 0 {
			if p.next.symbol == eoiToken {
				return nil
			}
			recoverSyms[p.next.symbol/8] &^= 1 << uint32(p.next.symbol%8)
			continue
		}

		if matchingPos < len(stack) {
			s = stack[matchingPos].sym.offset
		}
		stack = append(stack[:matchingPos], stackEntry{
			sym:   symbol{errSymbol, s, e},
			state: gotoState(stack[matchingPos-1].state, errSymbol),
		})
		return stack
	}
	return nil
}

func lalr(action, next int32) int32 {
	a := -action - 3
	for ; tmLalr[a] >= 0; a += 2 {
		if tmLalr[a] == next {
			break
		}
	}
	return tmLalr[a+1]
}

func gotoState(state int16, symbol int32) int16 {
	min := tmGoto[symbol]
	max := tmGoto[symbol+1]

	if max-min < 32 {
		for i := min; i < max; i += 2 {
			if tmFromTo[i] == state {
				return tmFromTo[i+1]
			}
		}
	} else {
		for min < max {
			e := (min + max) >> 1 &^ int32(1)
			i := tmFromTo[e]
			if i == state {
				return tmFromTo[e+1]
			} else if i < state {
				min = e + 2
			} else {
				max = e
			}
		}
	}
	return -1
}

func (p *Parser) fetchNext(lexer *Lexer, stack []stackEntry, ignoredTokens []symbol) []symbol {
restart:
	token := lexer.Next()
	switch token {
	case INVALID_TOKEN, MULTILINECOMMENT, COMMENT, TEMPLATES:
		s, e := lexer.Pos()
		tok := symbol{int32(token), s, e}
		if ignoredTokens == nil {
			p.reportIgnoredToken(tok)
		} else {
			ignoredTokens = append(ignoredTokens, tok)
		}
		goto restart
	}
	p.next.symbol = int32(token)
	p.next.offset, p.next.endoffset = lexer.Pos()
	return ignoredTokens
}

func (p *Parser) applyRule(rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer) (err error) {
	switch rule {
	case 189: // directive : '%' 'assert' 'empty' rhsSet ';'
		p.listener(Empty, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 190: // directive : '%' 'assert' 'nonempty' rhsSet ';'
		p.listener(NonEmpty, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 196: // inputref : symref 'no-eoi'
		p.listener(NoEoi, rhs[1].sym.offset, rhs[1].sym.endoffset)
	}
	nt := ruleNodeType[rule]
	if nt != 0 {
		p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}

func (p *Parser) reportIgnoredToken(tok symbol) {
	var t NodeType
	switch Token(tok.symbol) {
	case INVALID_TOKEN:
		t = InvalidToken
	case MULTILINECOMMENT:
		t = MultilineComment
	case COMMENT:
		t = Comment
	case TEMPLATES:
		t = Templates
	default:
		return
	}
	p.listener(t, tok.offset, tok.endoffset)
}

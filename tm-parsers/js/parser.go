// generated by Textmapper; DO NOT EDIT

package js

import (
	"context"
	"fmt"

	"github.com/inspirer/textmapper/tm-parsers/js/token"
)

// ErrorHandler is called every time a parser is unable to process some part of the input.
// This handler can return false to abort the parser.
type ErrorHandler func(err SyntaxError) bool

// StopOnFirstError is an error handler that forces the parser to stop on and return the first
// error.
func StopOnFirstError(_ SyntaxError) bool { return false }

type SyntaxError struct {
	Line      int
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line %v", e.Line)
}

const (
	startStackSize       = 256
	startTokenBufferSize = 16
	noToken              = int32(token.UNAVAILABLE)
	eoiToken             = int32(token.EOI)
	debugSyntax          = false
)

func (p *Parser) ParseModule(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 9, 9135, lexer)
}

func (p *Parser) ParseTypeSnippet(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 10, 9136, lexer)
}

func (p *Parser) ParseExpressionSnippet(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 11, 9137, lexer)
}

func (p *Parser) ParseNamespaceNameSnippet(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 12, 9138, lexer)
}

const errSymbol = 2

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

func (p *Parser) skipBrokenCode(ctx context.Context, lexer *Lexer, stack []stackEntry, canRecover func(symbol int32) bool) int {
	var e int
	for p.next.symbol != eoiToken && !canRecover(p.next.symbol) {
		if debugSyntax {
			fmt.Printf("skipped while recovering: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
		}
		if len(p.pending) > 0 {
			for _, tok := range p.pending {
				p.reportIgnoredToken(ctx, tok)
			}
			p.pending = p.pending[:0]
		}
		switch token.Token(p.next.symbol) {
		case token.NOSUBSTITUTIONTEMPLATE:
			p.listener(NoSubstitutionTemplate, p.next.offset, p.next.endoffset)
		case token.TEMPLATEHEAD:
			p.listener(TemplateHead, p.next.offset, p.next.endoffset)
		case token.TEMPLATEMIDDLE:
			p.listener(TemplateMiddle, p.next.offset, p.next.endoffset)
		case token.TEMPLATETAIL:
			p.listener(TemplateTail, p.next.offset, p.next.endoffset)
		}
		e = p.next.endoffset
		p.fetchNext(lexer, stack)
	}
	return e
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

func lookaheadRule(ctx context.Context, lexer *Lexer, next, rule int32, s *session) (sym int32, err error) {
	switch rule {
	case 5254:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 0, 9122, s); ok {
			sym = 791 /* lookahead_StartOfArrowFunction */
		} else {
			sym = 185 /* lookahead_notStartOfArrowFunction */
		}
		return
	case 5255:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 4, 9126, s); ok {
			sym = 883 /* lookahead_StartOfTypeImport */
		} else {
			sym = 884 /* lookahead_notStartOfTypeImport */
		}
		return
	case 5256:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 1, 9123, s); ok {
			sym = 370 /* lookahead_StartOfParametrizedCall */
		} else {
			sym = 341 /* lookahead_notStartOfParametrizedCall */
		}
		return
	case 5257:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 5, 9127, s); ok {
			sym = 947 /* lookahead_StartOfIs */
		} else {
			sym = 949 /* lookahead_notStartOfIs */
		}
		return
	case 5258:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 7, 9129, s); ok {
			sym = 983 /* lookahead_StartOfMappedType */
		} else {
			sym = 973 /* lookahead_notStartOfMappedType */
		}
		return
	case 5259:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 6, 9128, s); ok {
			sym = 1015 /* lookahead_StartOfFunctionType */
		} else {
			sym = 966 /* lookahead_notStartOfFunctionType */
		}
		return
	case 5260:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 8, 9130, s); ok {
			sym = 987 /* lookahead_StartOfTupleElementName */
		} else {
			sym = 988 /* lookahead_notStartOfTupleElementName */
		}
		return
	case 5261:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 3, 9125, s); ok {
			sym = 852 /* lookahead_StartOfExtendsTypeRef */
		} else {
			sym = 853 /* lookahead_notStartOfExtendsTypeRef */
		}
		return
	case 5262:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 2, 9124, s); ok {
			sym = 369 /* lookahead_StartLParen */
		} else {
			sym = 371 /* lookahead_notStartLParen */
		}
		return
	}
	return 0, nil
}

func AtStartOfArrowFunction(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 0, 9122, s)
}

func AtStartOfParametrizedCall(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 1, 9123, s)
}

func AtStartLParen(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 2, 9124, s)
}

func AtStartOfExtendsTypeRef(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 3, 9125, s)
}

func AtStartOfTypeImport(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 4, 9126, s)
}

func AtStartOfIs(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 5, 9127, s)
}

func AtStartOfFunctionType(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 6, 9128, s)
}

func AtStartOfMappedType(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 7, 9129, s)
}

func AtStartOfTupleElementName(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 8, 9130, s)
}

func lookahead(ctx context.Context, l *Lexer, next int32, start, end int16, s *session) (bool, error) {
	var lexer Lexer
	lexer.source = l.source
	lexer.ch = l.ch
	lexer.offset = l.offset
	lexer.tokenOffset = l.tokenOffset
	lexer.line = l.line
	lexer.tokenLine = l.tokenLine
	lexer.scanOffset = l.scanOffset
	lexer.State = l.State
	lexer.Dialect = l.Dialect
	lexer.token = l.token
	// Note: Stack is intentionally omitted.

	// Use memoization for recursive lookaheads.
	if next == noToken {
		next = lookaheadNext(&lexer, end, nil /*empty stack*/)
	}
	key := uint64(l.tokenOffset) + uint64(end)<<40
	if ret, ok := s.cache[key]; ok {
		return ret, nil
	}

	var allocated [64]stackEntry
	state := start
	stack := append(allocated[:0], stackEntry{state: state})

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if next == noToken {
				next = lookaheadNext(&lexer, end, stack)
			}
			action = lalr(action, next)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			stack = stack[:len(stack)-ln]
			sym, err := lookaheadRule(ctx, &lexer, next, rule, s)
			if err != nil {
				return false, err
			}
			if sym != 0 {
				entry.sym.symbol = sym
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			if s.shiftCounter++; s.shiftCounter&0x1ff == 0 {
				// Note: checking for context cancellation is expensive so we do it from time to time.
				select {
				case <-ctx.Done():
					return false, ctx.Err()
				default:
				}
			}

			// Shift.
			if next == noToken {
				next = lookaheadNext(&lexer, end, stack)
			}
			state = gotoState(state, next)
			stack = append(stack, stackEntry{
				sym:   symbol{symbol: next},
				state: state,
			})
			if state != -1 && next != eoiToken {
				next = noToken
			}
		}

		if action == -2 || state == -1 {
			break
		}
	}

	s.cache[key] = state == end
	return state == end, nil
}

func (p *Parser) applyRule(ctx context.Context, rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer, s *session) (err error) {
	switch rule {
	case 1112: // PrimaryExpression : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1126: // PrimaryExpression_Await : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1140: // PrimaryExpression_Await_NoAsync_NoLet : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1154: // PrimaryExpression_Await_NoLet : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1168: // PrimaryExpression_Await_NoLet_NoObjLiteral_NoFuncClass : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1176: // PrimaryExpression_Await_NoObjLiteral : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1189: // PrimaryExpression_Await_NoObjLiteral_NoFuncClass : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1197: // PrimaryExpression_NoAsync_NoLet : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1211: // PrimaryExpression_NoFuncClass : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1220: // PrimaryExpression_NoLet : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1234: // PrimaryExpression_NoLet_NoObjLiteral_NoFuncClass : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1242: // PrimaryExpression_NoObjLiteral : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1255: // PrimaryExpression_NoObjLiteral_NoFuncClass : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1263: // PrimaryExpression_Yield : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1277: // PrimaryExpression_Yield_Await : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1291: // PrimaryExpression_Yield_Await_NoAsync_NoLet : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1305: // PrimaryExpression_Yield_Await_NoLet : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1319: // PrimaryExpression_Yield_Await_NoLet_NoObjLiteral_NoFuncClass : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1327: // PrimaryExpression_Yield_Await_NoObjLiteral_NoFuncClass : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1335: // PrimaryExpression_Yield_NoAsync_NoLet : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1349: // PrimaryExpression_Yield_NoLet : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1363: // PrimaryExpression_Yield_NoLet_NoObjLiteral_NoFuncClass : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1371: // PrimaryExpression_Yield_NoObjLiteral_NoFuncClass : 'this'
		p.listener(This, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1420: // Elision : ','
		p.listener(NoElement, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1421: // Elision : Elision ','
		p.listener(NoElement, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 1446: // PropertyDefinition : IdentifierReference
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1454: // PropertyDefinition_Await : IdentifierReference_Await
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1462: // PropertyDefinition_Yield : IdentifierReference_Yield
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1470: // PropertyDefinition_Yield_Await : IdentifierReference_Yield_Await
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1489: // LiteralPropertyName : PrivateIdentifier
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1490: // LiteralPropertyName : StringLiteral
		p.listener(Literal, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1491: // LiteralPropertyName : NumericLiteral
		p.listener(Literal, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1493: // LiteralPropertyName_WithoutNew : PrivateIdentifier
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1494: // LiteralPropertyName_WithoutNew : StringLiteral
		p.listener(Literal, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1495: // LiteralPropertyName_WithoutNew : NumericLiteral
		p.listener(Literal, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1613: // MemberExpression_Await_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1678: // MemberExpression_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1743: // MemberExpression_Yield_Await_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1790: // MemberExpression_Yield_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1801: // SuperExpression : 'super'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 2697: // BinaryExpression : BinaryExpression .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2725: // BinaryExpression_Await : BinaryExpression_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2778: // BinaryExpression_Await_NoLet : BinaryExpression_Await_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2806: // BinaryExpression_Await_NoObjLiteral : BinaryExpression_Await_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2835: // BinaryExpression_In : BinaryExpression_In .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2864: // BinaryExpression_In_Await : BinaryExpression_In_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2919: // BinaryExpression_In_Await_NoObjLiteral : BinaryExpression_In_Await_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2948: // BinaryExpression_In_NoFuncClass : BinaryExpression_In_NoFuncClass .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3003: // BinaryExpression_In_NoObjLiteral : BinaryExpression_In_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3032: // BinaryExpression_In_Yield : BinaryExpression_In_Yield .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3061: // BinaryExpression_In_Yield_Await : BinaryExpression_In_Yield_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3166: // BinaryExpression_NoLet : BinaryExpression_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3194: // BinaryExpression_NoObjLiteral : BinaryExpression_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3222: // BinaryExpression_Yield : BinaryExpression_Yield .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3250: // BinaryExpression_Yield_Await : BinaryExpression_Yield_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3303: // BinaryExpression_Yield_Await_NoLet : BinaryExpression_Yield_Await_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3356: // BinaryExpression_Yield_NoLet : BinaryExpression_Yield_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3885: // ElementElision : ','
		p.listener(NoElement, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 3886: // ElementElision : Elision ','
		p.listener(NoElement, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3939: // SingleNameBinding : IdentifierReference Initializeropt_In
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 3940: // SingleNameBinding_Await : IdentifierReference_Await Initializeropt_In_Await
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 3941: // SingleNameBinding_Yield : IdentifierReference_Yield Initializeropt_In_Yield
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 3942: // SingleNameBinding_Yield_Await : IdentifierReference_Yield_Await Initializeropt_In_Yield_Await
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 3962: // IterationStatement : 'for' '(' 'var' VariableDeclarationList ';' .forSC ForCondition ';' .forSC ForFinalExpression ')' Statement
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3966: // IterationStatement : 'for' '(' 'var' ForBinding 'in' Expression_In ')' Statement
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3969: // IterationStatement : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In ')' Statement
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3970: // IterationStatement : 'for' '(' 'var' ForBinding 'of' AssignmentExpression_In ')' Statement
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3977: // IterationStatement_Await : 'for' '(' 'var' VariableDeclarationList_Await ';' .forSC ForCondition_Await ';' .forSC ForFinalExpression_Await ')' Statement_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3981: // IterationStatement_Await : 'for' '(' 'var' ForBinding_Await 'in' Expression_In_Await ')' Statement_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3983: // IterationStatement_Await : 'for' 'await' '(' LeftHandSideExpression_Await_NoAsync_NoLet 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3985: // IterationStatement_Await : 'for' 'await' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(ReferenceIdent, rhs[3].sym.offset, rhs[3].sym.endoffset)
		p.listener(IdentExpr, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 3986: // IterationStatement_Await : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3987: // IterationStatement_Await : 'for' 'await' '(' 'var' ForBinding_Await 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(Var, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 3988: // IterationStatement_Await : 'for' '(' 'var' ForBinding_Await 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3989: // IterationStatement_Await : 'for' 'await' '(' ForDeclaration_Await 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3996: // IterationStatement_Yield : 'for' '(' 'var' VariableDeclarationList_Yield ';' .forSC ForCondition_Yield ';' .forSC ForFinalExpression_Yield ')' Statement_Yield
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4000: // IterationStatement_Yield : 'for' '(' 'var' ForBinding_Yield 'in' Expression_In_Yield ')' Statement_Yield
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4003: // IterationStatement_Yield : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Yield ')' Statement_Yield
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4004: // IterationStatement_Yield : 'for' '(' 'var' ForBinding_Yield 'of' AssignmentExpression_In_Yield ')' Statement_Yield
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4011: // IterationStatement_Yield_Await : 'for' '(' 'var' VariableDeclarationList_Yield_Await ';' .forSC ForCondition_Yield_Await ';' .forSC ForFinalExpression_Yield_Await ')' Statement_Yield_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4015: // IterationStatement_Yield_Await : 'for' '(' 'var' ForBinding_Yield_Await 'in' Expression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4017: // IterationStatement_Yield_Await : 'for' 'await' '(' LeftHandSideExpression_Yield_Await_NoAsync_NoLet 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4019: // IterationStatement_Yield_Await : 'for' 'await' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(ReferenceIdent, rhs[3].sym.offset, rhs[3].sym.endoffset)
		p.listener(IdentExpr, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 4020: // IterationStatement_Yield_Await : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4021: // IterationStatement_Yield_Await : 'for' 'await' '(' 'var' ForBinding_Yield_Await 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(Var, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 4022: // IterationStatement_Yield_Await : 'for' '(' 'var' ForBinding_Yield_Await 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4023: // IterationStatement_Yield_Await : 'for' 'await' '(' ForDeclaration_Yield_Await 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4077: // CaseClause : 'case' Expression_In ':' StatementList
		p.listener(Cond, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4078: // CaseClause : 'case' Expression_In ':'
		p.listener(Cond, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4081: // CaseClause_Await : 'case' Expression_In_Await ':' StatementList_Await
		p.listener(Cond, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4082: // CaseClause_Await : 'case' Expression_In_Await ':'
		p.listener(Cond, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4085: // CaseClause_Yield : 'case' Expression_In_Yield ':' StatementList_Yield
		p.listener(Cond, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4086: // CaseClause_Yield : 'case' Expression_In_Yield ':'
		p.listener(Cond, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4089: // CaseClause_Yield_Await : 'case' Expression_In_Yield_Await ':' StatementList_Yield_Await
		p.listener(Cond, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4090: // CaseClause_Yield_Await : 'case' Expression_In_Yield_Await ':'
		p.listener(Cond, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4476: // ImportDeclaration : 'import' lookahead_StartOfTypeImport 'type' ImportClause FromClause .noLineBreak AssertClause ';'
		p.listener(TsTypeOnly, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4477: // ImportDeclaration : 'import' lookahead_StartOfTypeImport 'type' ImportClause FromClause ';'
		p.listener(TsTypeOnly, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4493: // ImportRequireDeclaration : 'export' 'import' lookahead_notStartOfTypeImport BindingIdentifier '=' 'require' '(' ModuleSpecifier ')' ';'
		p.listener(TsExport, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4508: // NamedImport : IdentifierNameRef
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4513: // ExportDeclaration : 'export' 'type' '*' 'as' ImportedBinding FromClause .noLineBreak AssertClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4514: // ExportDeclaration : 'export' 'type' '*' 'as' ImportedBinding FromClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4515: // ExportDeclaration : 'export' 'type' '*' FromClause .noLineBreak AssertClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4516: // ExportDeclaration : 'export' 'type' '*' FromClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4521: // ExportDeclaration : 'export' 'type' ExportClause FromClause .noLineBreak AssertClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4522: // ExportDeclaration : 'export' 'type' ExportClause FromClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4525: // ExportDeclaration : 'export' 'type' ExportClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4541: // ExportElement : IdentifierNameRef
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4547: // DecoratorMemberExpression : DecoratorMemberExpression '.' IdentifierName
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4648: // TypePredicate : 'asserts' lookahead_StartOfIs 'is' Type_NoQuest
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4650: // TypePredicate_NoQuest : 'asserts' lookahead_StartOfIs 'is' Type_NoQuest
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4652: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs 'this' 'is' Type
		p.listener(This, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4653: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs 'this'
		p.listener(This, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4654: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs IdentifierName_WithoutKeywords_WithoutAs_WithoutSatisfies 'is' Type
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4655: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs IdentifierName_WithoutKeywords_WithoutAs_WithoutSatisfies
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4686: // TypeOperator : 'infer' IdentifierName
		p.listener(NameIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4691: // TypeOperator_NoQuest : 'infer' IdentifierName
		p.listener(NameIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4742: // TypeName : NamespaceName '.' IdentifierReference_WithDefault
		p.listener(TsNamespaceName, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4844: // TupleElementType : '...' lookahead_StartOfTupleElementName IdentifierName '?' ':' Type
		p.listener(RestType, rhs[5].sym.offset, rhs[5].sym.endoffset)
	case 4845: // TupleElementType : '...' lookahead_StartOfTupleElementName IdentifierName ':' Type
		p.listener(RestType, rhs[4].sym.offset, rhs[4].sym.endoffset)
	case 4877: // ConstructorType : 'abstract' 'new' TypeParameters ParameterList '=>' Type
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4878: // ConstructorType : 'abstract' 'new' ParameterList '=>' Type
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4881: // ConstructorType_NoQuest : 'abstract' 'new' TypeParameters ParameterList '=>' Type_NoQuest
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4882: // ConstructorType_NoQuest : 'abstract' 'new' ParameterList '=>' Type_NoQuest
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4888: // ImportTypeStart : 'typeof' 'import' '(' Type ')'
		p.listener(TsTypeOf, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4951: // Parameter : Modifiers BindingIdentifier '?' TypeAnnotation
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4952: // Parameter : Modifiers BindingIdentifier '?'
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4955: // Parameter : BindingIdentifier '?' TypeAnnotation
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4956: // Parameter : BindingIdentifier '?'
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4959: // Parameter : Modifiers BindingPattern '?' TypeAnnotation
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4960: // Parameter : Modifiers BindingPattern '?'
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4963: // Parameter : BindingPattern '?' TypeAnnotation
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4964: // Parameter : BindingPattern '?'
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4981: // Parameter_Await : Modifiers BindingIdentifier '?' TypeAnnotation
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4982: // Parameter_Await : Modifiers BindingIdentifier '?'
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4985: // Parameter_Await : BindingIdentifier '?' TypeAnnotation
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4986: // Parameter_Await : BindingIdentifier '?'
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4989: // Parameter_Await : Modifiers BindingPattern_Await '?' TypeAnnotation
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4990: // Parameter_Await : Modifiers BindingPattern_Await '?'
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4993: // Parameter_Await : BindingPattern_Await '?' TypeAnnotation
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4994: // Parameter_Await : BindingPattern_Await '?'
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5011: // Parameter_Yield : Modifiers BindingIdentifier '?' TypeAnnotation
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 5012: // Parameter_Yield : Modifiers BindingIdentifier '?'
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 5015: // Parameter_Yield : BindingIdentifier '?' TypeAnnotation
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5016: // Parameter_Yield : BindingIdentifier '?'
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5019: // Parameter_Yield : Modifiers BindingPattern_Yield '?' TypeAnnotation
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 5020: // Parameter_Yield : Modifiers BindingPattern_Yield '?'
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 5023: // Parameter_Yield : BindingPattern_Yield '?' TypeAnnotation
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5024: // Parameter_Yield : BindingPattern_Yield '?'
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5041: // Parameter_Yield_Await : Modifiers BindingIdentifier '?' TypeAnnotation
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 5042: // Parameter_Yield_Await : Modifiers BindingIdentifier '?'
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 5045: // Parameter_Yield_Await : BindingIdentifier '?' TypeAnnotation
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5046: // Parameter_Yield_Await : BindingIdentifier '?'
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5049: // Parameter_Yield_Await : Modifiers BindingPattern_Yield_Await '?' TypeAnnotation
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 5050: // Parameter_Yield_Await : Modifiers BindingPattern_Yield_Await '?'
		p.listener(TsOptional, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 5053: // Parameter_Yield_Await : BindingPattern_Yield_Await '?' TypeAnnotation
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5054: // Parameter_Yield_Await : BindingPattern_Yield_Await '?'
		p.listener(TsOptional, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5082: // IndexSignature : Modifiers '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 5083: // IndexSignature : '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5084: // IndexSignature_WithDeclare : Modifiers_WithDeclare '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 5085: // IndexSignature_WithDeclare : '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5090: // MethodSignature : Modifiers 'new' '?' FormalParameters
		p.listener(NameIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(LiteralPropertyName, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5091: // MethodSignature : 'new' '?' FormalParameters
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
		p.listener(LiteralPropertyName, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 5101: // EnumDeclaration : 'const' 'enum' BindingIdentifier EnumBody
		p.listener(TsConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 5128: // AmbientVariableDeclaration : 'var' AmbientBindingList ';'
		p.listener(Var, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 5129: // AmbientVariableDeclaration : 'let' AmbientBindingList ';'
		p.listener(LetOrConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 5130: // AmbientVariableDeclaration : 'const' AmbientBindingList ';'
		p.listener(LetOrConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 5144: // AmbientEnumDeclaration : 'const' 'enum' BindingIdentifier EnumBody
		p.listener(TsConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 5172: // AmbientModuleDeclaration : 'module' StringLiteral AmbientModuleBody
		p.listener(Literal, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5173: // AmbientModuleDeclaration : 'module' StringLiteral ';'
		p.listener(Literal, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 5254:
		var ok bool
		if ok, err = AtStartOfArrowFunction(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 791 /* lookahead_StartOfArrowFunction */
		} else {
			lhs.sym.symbol = 185 /* lookahead_notStartOfArrowFunction */
		}
		return
	case 5255:
		var ok bool
		if ok, err = AtStartOfTypeImport(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 883 /* lookahead_StartOfTypeImport */
		} else {
			lhs.sym.symbol = 884 /* lookahead_notStartOfTypeImport */
		}
		return
	case 5256:
		var ok bool
		if ok, err = AtStartOfParametrizedCall(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 370 /* lookahead_StartOfParametrizedCall */
		} else {
			lhs.sym.symbol = 341 /* lookahead_notStartOfParametrizedCall */
		}
		return
	case 5257:
		var ok bool
		if ok, err = AtStartOfIs(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 947 /* lookahead_StartOfIs */
		} else {
			lhs.sym.symbol = 949 /* lookahead_notStartOfIs */
		}
		return
	case 5258:
		var ok bool
		if ok, err = AtStartOfMappedType(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 983 /* lookahead_StartOfMappedType */
		} else {
			lhs.sym.symbol = 973 /* lookahead_notStartOfMappedType */
		}
		return
	case 5259:
		var ok bool
		if ok, err = AtStartOfFunctionType(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 1015 /* lookahead_StartOfFunctionType */
		} else {
			lhs.sym.symbol = 966 /* lookahead_notStartOfFunctionType */
		}
		return
	case 5260:
		var ok bool
		if ok, err = AtStartOfTupleElementName(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 987 /* lookahead_StartOfTupleElementName */
		} else {
			lhs.sym.symbol = 988 /* lookahead_notStartOfTupleElementName */
		}
		return
	case 5261:
		var ok bool
		if ok, err = AtStartOfExtendsTypeRef(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 852 /* lookahead_StartOfExtendsTypeRef */
		} else {
			lhs.sym.symbol = 853 /* lookahead_notStartOfExtendsTypeRef */
		}
		return
	case 5262:
		var ok bool
		if ok, err = AtStartLParen(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 369 /* lookahead_StartLParen */
		} else {
			lhs.sym.symbol = 371 /* lookahead_notStartLParen */
		}
		return
	}
	if nt := tmRuleType[rule]; nt != 0 {
		p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}

func (p *Parser) reportIgnoredToken(ctx context.Context, tok symbol) {
	var t NodeType
	switch token.Token(tok.symbol) {
	case token.MULTILINECOMMENT:
		t = MultiLineComment
	case token.SINGLELINECOMMENT:
		t = SingleLineComment
	case token.INVALID_TOKEN:
		t = InvalidToken
	default:
		return
	}
	if debugSyntax {
		fmt.Printf("ignored: %v as %v\n", token.Token(tok.symbol), t)
	}
	p.listener(t, tok.offset, tok.endoffset)
}

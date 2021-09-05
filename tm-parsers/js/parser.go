// generated by Textmapper; DO NOT EDIT

package js

import (
	"context"
	"fmt"
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

func (p *Parser) ParseModule(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 8, 8631, lexer)
}

func (p *Parser) ParseTypeSnippet(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 9, 8632, lexer)
}

func (p *Parser) ParseExpressionSnippet(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 10, 8633, lexer)
}

func lookaheadRule(ctx context.Context, lexer *Lexer, next, rule int32, s *session) (sym int32, err error) {
	switch rule {
	case 5072:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 0, 8620, s); ok {
			sym = 789 /* lookahead_StartOfArrowFunction */
		} else {
			sym = 176 /* lookahead_notStartOfArrowFunction */
		}
		return
	case 5073:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 3, 8623, s); ok {
			sym = 875 /* lookahead_StartOfTypeImport */
		} else {
			sym = 876 /* lookahead_notStartOfTypeImport */
		}
		return
	case 5074:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 1, 8621, s); ok {
			sym = 371 /* lookahead_StartOfParametrizedCall */
		} else {
			sym = 343 /* lookahead_notStartOfParametrizedCall */
		}
		return
	case 5075:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 4, 8624, s); ok {
			sym = 935 /* lookahead_StartOfIs */
		} else {
			sym = 938 /* lookahead_notStartOfIs */
		}
		return
	case 5076:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 6, 8626, s); ok {
			sym = 972 /* lookahead_StartOfMappedType */
		} else {
			sym = 962 /* lookahead_notStartOfMappedType */
		}
		return
	case 5077:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 5, 8625, s); ok {
			sym = 984 /* lookahead_StartOfFunctionType */
		} else {
			sym = 955 /* lookahead_notStartOfFunctionType */
		}
		return
	case 5078:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 7, 8627, s); ok {
			sym = 976 /* lookahead_StartOfTupleElementName */
		} else {
			sym = 977 /* lookahead_notStartOfTupleElementName */
		}
		return
	case 5079:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 2, 8622, s); ok {
			sym = 850 /* lookahead_StartOfExtendsTypeRef */
		} else {
			sym = 851 /* lookahead_notStartOfExtendsTypeRef */
		}
		return
	}
	return 0, nil
}

func AtStartOfArrowFunction(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 0, 8620, s)
}

func AtStartOfParametrizedCall(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 1, 8621, s)
}

func AtStartOfExtendsTypeRef(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 2, 8622, s)
}

func AtStartOfTypeImport(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 3, 8623, s)
}

func AtStartOfIs(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 4, 8624, s)
}

func AtStartOfFunctionType(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 5, 8625, s)
}

func AtStartOfMappedType(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 6, 8626, s)
}

func AtStartOfTupleElementName(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 7, 8627, s)
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

func (p *Parser) applyRule(ctx context.Context, rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer, s *session) (err error) {
	switch rule {
	case 1267: // Elision : ','
		p.listener(NoElement, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1268: // Elision : Elision ','
		p.listener(NoElement, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 1336: // LiteralPropertyName : PrivateIdentifier
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1340: // LiteralPropertyName_WithoutNew : PrivateIdentifier
		p.listener(NameIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1473: // MemberExpression_Await_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1574: // MemberExpression_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1657: // MemberExpression_Yield_Await_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1722: // MemberExpression_Yield_StartWithLet : 'let'
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 2610: // BinaryExpression : BinaryExpression .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2637: // BinaryExpression_Await : BinaryExpression_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2689: // BinaryExpression_Await_NoLet : BinaryExpression_Await_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2716: // BinaryExpression_Await_NoObjLiteral : BinaryExpression_Await_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2744: // BinaryExpression_In : BinaryExpression_In .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2772: // BinaryExpression_In_Await : BinaryExpression_In_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2826: // BinaryExpression_In_Await_NoObjLiteral : BinaryExpression_In_Await_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2854: // BinaryExpression_In_NoFuncClass : BinaryExpression_In_NoFuncClass .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2908: // BinaryExpression_In_NoObjLiteral : BinaryExpression_In_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2936: // BinaryExpression_In_Yield : BinaryExpression_In_Yield .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 2964: // BinaryExpression_In_Yield_Await : BinaryExpression_In_Yield_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3068: // BinaryExpression_NoLet : BinaryExpression_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3095: // BinaryExpression_NoObjLiteral : BinaryExpression_NoObjLiteral .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3122: // BinaryExpression_Yield : BinaryExpression_Yield .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3149: // BinaryExpression_Yield_Await : BinaryExpression_Yield_Await .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3201: // BinaryExpression_Yield_Await_NoLet : BinaryExpression_Yield_Await_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3253: // BinaryExpression_Yield_NoLet : BinaryExpression_Yield_NoLet .noLineBreak 'as' 'const'
		p.listener(TsConst, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3773: // ElementElision : ','
		p.listener(NoElement, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 3774: // ElementElision : Elision ','
		p.listener(NoElement, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3850: // IterationStatement : 'for' '(' 'var' VariableDeclarationList ';' .forSC ForCondition ';' .forSC ForFinalExpression ')' Statement
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3854: // IterationStatement : 'for' '(' 'var' ForBinding 'in' Expression_In ')' Statement
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3857: // IterationStatement : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In ')' Statement
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3858: // IterationStatement : 'for' '(' 'var' ForBinding 'of' AssignmentExpression_In ')' Statement
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3865: // IterationStatement_Await : 'for' '(' 'var' VariableDeclarationList_Await ';' .forSC ForCondition_Await ';' .forSC ForFinalExpression_Await ')' Statement_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3869: // IterationStatement_Await : 'for' '(' 'var' ForBinding_Await 'in' Expression_In_Await ')' Statement_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3871: // IterationStatement_Await : 'for' 'await' '(' LeftHandSideExpression_Await_NoAsync_NoLet 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3873: // IterationStatement_Await : 'for' 'await' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(ReferenceIdent, rhs[3].sym.offset, rhs[3].sym.endoffset)
		p.listener(IdentExpr, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 3874: // IterationStatement_Await : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3875: // IterationStatement_Await : 'for' 'await' '(' 'var' ForBinding_Await 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(Var, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 3876: // IterationStatement_Await : 'for' '(' 'var' ForBinding_Await 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3877: // IterationStatement_Await : 'for' 'await' '(' ForDeclaration_Await 'of' AssignmentExpression_In_Await ')' Statement_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3884: // IterationStatement_Yield : 'for' '(' 'var' VariableDeclarationList_Yield ';' .forSC ForCondition_Yield ';' .forSC ForFinalExpression_Yield ')' Statement_Yield
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3888: // IterationStatement_Yield : 'for' '(' 'var' ForBinding_Yield 'in' Expression_In_Yield ')' Statement_Yield
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3891: // IterationStatement_Yield : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Yield ')' Statement_Yield
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3892: // IterationStatement_Yield : 'for' '(' 'var' ForBinding_Yield 'of' AssignmentExpression_In_Yield ')' Statement_Yield
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3899: // IterationStatement_Yield_Await : 'for' '(' 'var' VariableDeclarationList_Yield_Await ';' .forSC ForCondition_Yield_Await ';' .forSC ForFinalExpression_Yield_Await ')' Statement_Yield_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3903: // IterationStatement_Yield_Await : 'for' '(' 'var' ForBinding_Yield_Await 'in' Expression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3905: // IterationStatement_Yield_Await : 'for' 'await' '(' LeftHandSideExpression_Yield_Await_NoAsync_NoLet 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 3907: // IterationStatement_Yield_Await : 'for' 'await' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(ReferenceIdent, rhs[3].sym.offset, rhs[3].sym.endoffset)
		p.listener(IdentExpr, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 3908: // IterationStatement_Yield_Await : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
		p.listener(IdentExpr, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3909: // IterationStatement_Yield_Await : 'for' 'await' '(' 'var' ForBinding_Yield_Await 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
		p.listener(Var, rhs[3].sym.offset, rhs[3].sym.endoffset)
	case 3910: // IterationStatement_Yield_Await : 'for' '(' 'var' ForBinding_Yield_Await 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Var, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3911: // IterationStatement_Yield_Await : 'for' 'await' '(' ForDeclaration_Yield_Await 'of' AssignmentExpression_In_Yield_Await ')' Statement_Yield_Await
		p.listener(Await, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4342: // ImportDeclaration : 'import' lookahead_StartOfTypeImport 'type' ImportClause FromClause ';'
		p.listener(TsTypeOnly, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4349: // ImportRequireDeclaration : 'export' 'import' lookahead_notStartOfTypeImport BindingIdentifier '=' 'require' '(' StringLiteral ')' ';'
		p.listener(TsExport, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4369: // ExportDeclaration : 'export' 'type' '*' 'as' ImportedBinding FromClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4370: // ExportDeclaration : 'export' 'type' '*' FromClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4373: // ExportDeclaration : 'export' 'type' ExportClause FromClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4375: // ExportDeclaration : 'export' 'type' ExportClause ';'
		p.listener(TsTypeOnly, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4397: // DecoratorMemberExpression : DecoratorMemberExpression '.' IdentifierName
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4498: // TypePredicate : 'asserts' lookahead_StartOfIs 'is' Type_NoQuest
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4500: // TypePredicate_NoQuest : 'asserts' lookahead_StartOfIs 'is' Type_NoQuest
		p.listener(ReferenceIdent, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4502: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs 'this' 'is' Type
		p.listener(This, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4503: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs 'this'
		p.listener(This, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4504: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs IdentifierName_WithoutKeywords_WithoutAs 'is' Type
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4505: // AssertsType : 'asserts' .noLineBreak lookahead_notStartOfIs IdentifierName_WithoutKeywords_WithoutAs
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4506: // AssertsType_NoQuest : 'asserts' .noLineBreak lookahead_notStartOfIs 'this' 'is' Type_NoQuest
		p.listener(This, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4507: // AssertsType_NoQuest : 'asserts' .noLineBreak lookahead_notStartOfIs 'this'
		p.listener(This, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4508: // AssertsType_NoQuest : 'asserts' .noLineBreak lookahead_notStartOfIs IdentifierName_WithoutKeywords_WithoutAs 'is' Type_NoQuest
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4509: // AssertsType_NoQuest : 'asserts' .noLineBreak lookahead_notStartOfIs IdentifierName_WithoutKeywords_WithoutAs
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4539: // TypeOperator : 'infer' IdentifierName
		p.listener(ReferenceIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4544: // TypeOperator_NoQuest : 'infer' IdentifierName
		p.listener(ReferenceIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4692: // TupleElementType : '...' lookahead_StartOfTupleElementName IdentifierName '?' ':' Type
		p.listener(RestType, rhs[5].sym.offset, rhs[5].sym.endoffset)
	case 4693: // TupleElementType : '...' lookahead_StartOfTupleElementName IdentifierName ':' Type
		p.listener(RestType, rhs[4].sym.offset, rhs[4].sym.endoffset)
	case 4732: // ConstructorType : 'abstract' 'new' TypeParameters ParameterList '=>' Type
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4733: // ConstructorType : 'abstract' 'new' ParameterList '=>' Type
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4736: // ConstructorType_NoQuest : 'abstract' 'new' TypeParameters ParameterList '=>' Type_NoQuest
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4737: // ConstructorType_NoQuest : 'abstract' 'new' ParameterList '=>' Type_NoQuest
		p.listener(Abstract, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4752: // TypeQueryExpression : TypeQueryExpression '.' IdentifierName
		p.listener(ReferenceIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4929: // IndexSignature : Modifiers '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4930: // IndexSignature : '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4931: // IndexSignature_WithDeclare : Modifiers_WithDeclare '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 4932: // IndexSignature_WithDeclare : '[' IdentifierName ':' Type ']' TypeAnnotation
		p.listener(NameIdent, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 4946: // EnumDeclaration : 'const' 'enum' BindingIdentifier EnumBody
		p.listener(TsConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4972: // AmbientVariableDeclaration : 'var' AmbientBindingList ';'
		p.listener(Var, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4973: // AmbientVariableDeclaration : 'let' AmbientBindingList ';'
		p.listener(LetOrConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4974: // AmbientVariableDeclaration : 'const' AmbientBindingList ';'
		p.listener(LetOrConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4988: // AmbientEnumDeclaration : 'const' 'enum' BindingIdentifier EnumBody
		p.listener(TsConst, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 5072:
		var ok bool
		if ok, err = AtStartOfArrowFunction(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 789 /* lookahead_StartOfArrowFunction */
		} else {
			lhs.sym.symbol = 176 /* lookahead_notStartOfArrowFunction */
		}
		return
	case 5073:
		var ok bool
		if ok, err = AtStartOfTypeImport(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 875 /* lookahead_StartOfTypeImport */
		} else {
			lhs.sym.symbol = 876 /* lookahead_notStartOfTypeImport */
		}
		return
	case 5074:
		var ok bool
		if ok, err = AtStartOfParametrizedCall(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 371 /* lookahead_StartOfParametrizedCall */
		} else {
			lhs.sym.symbol = 343 /* lookahead_notStartOfParametrizedCall */
		}
		return
	case 5075:
		var ok bool
		if ok, err = AtStartOfIs(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 935 /* lookahead_StartOfIs */
		} else {
			lhs.sym.symbol = 938 /* lookahead_notStartOfIs */
		}
		return
	case 5076:
		var ok bool
		if ok, err = AtStartOfMappedType(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 972 /* lookahead_StartOfMappedType */
		} else {
			lhs.sym.symbol = 962 /* lookahead_notStartOfMappedType */
		}
		return
	case 5077:
		var ok bool
		if ok, err = AtStartOfFunctionType(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 984 /* lookahead_StartOfFunctionType */
		} else {
			lhs.sym.symbol = 955 /* lookahead_notStartOfFunctionType */
		}
		return
	case 5078:
		var ok bool
		if ok, err = AtStartOfTupleElementName(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 976 /* lookahead_StartOfTupleElementName */
		} else {
			lhs.sym.symbol = 977 /* lookahead_notStartOfTupleElementName */
		}
		return
	case 5079:
		var ok bool
		if ok, err = AtStartOfExtendsTypeRef(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 850 /* lookahead_StartOfExtendsTypeRef */
		} else {
			lhs.sym.symbol = 851 /* lookahead_notStartOfExtendsTypeRef */
		}
		return
	}
	if nt := tmRuleType[rule]; nt != 0 {
		p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}

const errSymbol = 2

func (p *Parser) skipBrokenCode(lexer *Lexer, stack []stackEntry, canRecover func(symbol int32) bool) int {
	var e int
	for p.next.symbol != eoiToken && !canRecover(p.next.symbol) {
		if debugSyntax {
			fmt.Printf("skipped while recovering: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
		}
		switch Token(p.next.symbol) {
		case NOSUBSTITUTIONTEMPLATE:
			p.listener(NoSubstitutionTemplate, p.next.offset, p.next.endoffset)
		case TEMPLATEHEAD:
			p.listener(TemplateHead, p.next.offset, p.next.endoffset)
		case TEMPLATEMIDDLE:
			p.listener(TemplateMiddle, p.next.offset, p.next.endoffset)
		case TEMPLATETAIL:
			p.listener(TemplateTail, p.next.offset, p.next.endoffset)
		}
		e = p.next.endoffset
		p.fetchNext(lexer, stack, nil)
	}
	return e
}

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

func (p *Parser) reportIgnoredToken(tok symbol) {
	var t NodeType
	switch Token(tok.symbol) {
	case MULTILINECOMMENT:
		t = MultiLineComment
	case SINGLELINECOMMENT:
		t = SingleLineComment
	case INVALID_TOKEN:
		t = InvalidToken
	default:
		return
	}
	if debugSyntax {
		fmt.Printf("ignored: %v as %v\n", Token(tok.symbol), t)
	}
	p.listener(t, tok.offset, tok.endoffset)
}

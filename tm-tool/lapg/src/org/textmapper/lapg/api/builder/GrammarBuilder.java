/**
 * Copyright 2002-2022 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textmapper.lapg.api.builder;

import org.textmapper.lapg.api.*;
import org.textmapper.lapg.api.ast.AstType;
import org.textmapper.lapg.api.regex.RegexPart;
import org.textmapper.lapg.api.rule.*;
import org.textmapper.lapg.api.rule.RhsSet.Operation;

import java.util.Collection;

public interface GrammarBuilder extends GrammarMapper {

	Terminal addTerminal(Name name, AstType type, boolean isSpace, SourceElement origin);

	Nonterminal addNonterminal(Name name, SourceElement origin);

	Nonterminal addAnonymous(String nameHint, Symbol anchor, SourceElement origin);

	Nonterminal addShared(RhsPart part, Symbol anchor, String nameHint);

	Terminal getEoi();

	TemplateParameter addParameter(TemplateParameter.Type type,
								   Name name, Object defaultValue,
								   TemplateParameter.Modifier m, SourceElement origin);

	TemplateEnvironment getRootEnvironment();

	NamedPattern addPattern(Name name, RegexPart regexp, SourceElement origin);

	LexerState addState(Name name, SourceElement origin);

	LexerRule addLexerRule(int kind, Terminal sym, RegexPart regexp, Iterable<LexerState> states,
						   int priority, int order, LexerRule classLexerRule, SourceElement origin);

	NamedSet addSet(Name name, RhsSet set, SourceElement origin);

	RhsArgument argument(TemplateParameter param, TemplateParameter source, Object value,
						 SourceElement origin);

	RhsSymbol symbol(Symbol sym, Collection<RhsArgument> args, SourceElement origin);

	RhsSymbol symbolFwdAll(Symbol sym, SourceElement origin);

	RhsSymbol templateSymbol(TemplateParameter parameter, Collection<RhsArgument> args,
							 SourceElement origin);

	RhsAssignment assignment(String name, RhsPart inner, boolean isAddition, SourceElement origin);

	RhsCast cast(Symbol asSymbol, Collection<RhsArgument> args, RhsPart inner,
				 SourceElement origin);

	RhsChoice choice(Collection<RhsPart> parts, SourceElement origin);

	RhsConditional conditional(RhsPredicate predicate, RhsSequence inner, SourceElement origin);

	RhsPredicate predicate(RhsPredicate.Operation operation, Collection<RhsPredicate> inner,
						   TemplateParameter param, Object value, SourceElement origin);

	LookaheadPredicate lookaheadPredicate(InputRef inputRef, boolean negate);

	Lookahead lookahead(Collection<LookaheadPredicate> predicates,
						Symbol anchor, SourceElement origin);

	RhsSequence sequence(String name, Collection<RhsPart> parts, SourceElement origin);

	RhsSequence asSequence(RhsPart part);

	RhsSequence empty(SourceElement origin);

	RhsOptional optional(RhsPart inner, SourceElement origin);

	RhsList list(RhsSequence inner, RhsPart separator, boolean nonEmpty, SourceElement origin);

	void addParentheses(Terminal opening, Terminal closing);

	RhsIgnored ignored(RhsPart inner, SourceElement origin);

	RhsSet set(Operation operation,
			   Symbol symbol, Collection<RhsArgument> args,
			   Collection<RhsSet> parts, SourceElement origin);

	RhsStateMarker stateMarker(String name, SourceElement origin);

	RhsSequence addPrecedence(RhsPart p, Terminal prec);

	void define(Nonterminal left, RhsRoot rhs);

	void addRule(Nonterminal left, RhsRule rhs);

	void expectSR(int value);
	void expectRR(int value);


	InputRef addInput(Nonterminal inputSymbol, boolean hasEoi, SourceElement origin);

	Prio addPrio(int prio, Collection<Terminal> symbols, SourceElement origin);


	Grammar create();
}

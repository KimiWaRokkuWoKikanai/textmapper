/**
 * Copyright 2010-2017 Evgeny Gryaznov
 * <p>
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * <p>
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * <p>
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see http://www.gnu.org/licenses/.
 */
package org.textmapper.idea.lang.syntax.psi;

import com.intellij.lang.ASTNode;
import com.intellij.psi.util.PsiTreeUtil;
import org.jetbrains.annotations.NotNull;

import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.StreamSupport;

/**
 * Gryaznov Evgeny, 1/26/11
 */
public class TmGrammar extends TmElement {

	public TmGrammar(@NotNull ASTNode node) {
		super(node);
	}

	public TmNamedElement[] getNamedElements() {
		Iterable<TmNamedElement> elements = getElements(TmNamedElement.class);
		List<TmNamedElement> list = StreamSupport.stream(elements.spliterator(), false).collect(Collectors.toList());
		return list.toArray(new TmNamedElement[list.size()]);
	}

	public TmHeader getHeader() {
		return PsiTreeUtil.getChildOfType(this, TmHeader.class);
	}

	public List<TmImport> getImports() {
		return PsiTreeUtil.getChildrenOfTypeAsList(this, TmImport.class);
	}

	public List<TmOption> getOptions() {
		return PsiTreeUtil.getChildrenOfTypeAsList(this, TmOption.class);
	}

	public List<TmLexeme> getLexemes() {
		return PsiTreeUtil.getChildrenOfTypeAsList(this, TmLexeme.class);
	}

	public List<TmNonterm> getNonterms() {
		return PsiTreeUtil.getChildrenOfTypeAsList(this, TmNonterm.class);
	}

	public TmLexerState[] getStates() {
		List<TmLexerState> states = new ArrayList<>();
		Set<String> seen = new HashSet<>();
		for (TmStatesClause selector : getElements(TmStatesClause.class)) {
			for (TmLexerState tmLexerState : selector.getStates()) {
				if (seen.add(tmLexerState.getName())) {
					states.add(tmLexerState);
				}
			}
		}
		return states.toArray(new TmLexerState[states.size()]);
	}

	public List<TmStartConditionsScope> getStartConditionScopes() {
		return PsiTreeUtil.getChildrenOfTypeAsList(this, TmStartConditionsScope.class);
	}

	public TmNamedElement resolve(String name) {
		if (name.endsWith("opt") && name.length() > 3) {
			name = name.substring(0, name.length() - 3);
		}

		for (TmNamedElement named : getElements(TmNamedElement.class)) {
			if (name.equals(named.getName())) {
				return named;
			}
		}
		return null;
	}

	public TmNamedElement resolveState(String name) {
		for (TmStatesClause clause : getElements(TmStatesClause.class)) {
			for (TmLexerState state : clause.getStates()) {
				if (name.equals(state.getName())) {
					return state;
				}
			}
		}
		return null;
	}

	<T extends TmElement> Iterable<T> getElements(Class<T> c) {
		return () -> new Iterator<T>() {
			Stack<Iterator<TmElement>> stack = new Stack<>();
			T next;
			{
				stack.push(PsiTreeUtil.getChildrenOfAnyType(TmGrammar.this, TmStartConditionsScope.class, c).iterator());
				fetch();
			}

			private void fetch() {
				next = null;
				while (!stack.empty()) {
					if (!stack.peek().hasNext()) {
						stack.pop();
						continue;
					}
					TmElement next = stack.peek().next();
					if (c.isInstance(next)) {
						this.next = (T) next;
						return;
					}
					if (next instanceof TmStartConditionsScope) {
						stack.push(PsiTreeUtil.getChildrenOfAnyType(next, TmStartConditionsScope.class, c).iterator());
					}
				}
			}

			@Override
			public boolean hasNext() {
				return next != null;
			}

			@Override
			public T next() {
				T r = next;
				fetch();
				return r;
			}
		};
	}
}

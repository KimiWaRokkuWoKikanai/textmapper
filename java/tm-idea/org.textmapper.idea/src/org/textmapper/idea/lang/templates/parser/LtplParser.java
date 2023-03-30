/**
 * Copyright 2010-2017 Evgeny Gryaznov
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see http://www.gnu.org/licenses/.
 */
package org.textmapper.idea.lang.templates.parser;

import com.intellij.lang.ASTNode;
import com.intellij.lang.PsiBuilder;
import com.intellij.lang.PsiBuilder.Marker;
import com.intellij.lang.PsiParser;
import com.intellij.psi.TokenType;
import com.intellij.psi.tree.IElementType;
import org.jetbrains.annotations.NotNull;
import org.textmapper.idea.lang.templates.lexer.LtplElementType;
import org.textmapper.templates.ast.TemplatesLexer;
import org.textmapper.templates.ast.TemplatesLexer.ErrorReporter;
import org.textmapper.templates.ast.TemplatesLexer.Span;
import org.textmapper.templates.ast.TemplatesLexer.Tokens;
import org.textmapper.templates.ast.TemplatesParser;
import org.textmapper.templates.ast.TemplatesParser.ParseException;
import org.textmapper.templates.ast.TemplatesParser.Nonterminals;

import java.io.IOException;
import java.io.Reader;
import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

/**
 * evgeny, 3/3/12
 */
public class LtplParser implements PsiParser {

	private static final Map<Integer, IElementType> types = initTypes();

	private static Map<Integer, IElementType> initTypes() {
		Map<Integer, IElementType> result = new HashMap<>();
		result.put(Nonterminals.syntax_problem, TokenType.ERROR_ELEMENT);
		for (IElementType t : LtplElementTypes.allElements) {
			result.put(((LtplElementType) t).getSymbol(), t);
		}
		for (int expr : LtplElementTypes.allExpressions) {
			result.put(expr, LtplElementTypes.EXPRESSION);
		}
		return result;
	}

	private static IElementType reduceType(int token, int rule) {
		IElementType type = types.get(token);
		if (type != null) {
			return type;
		}

		return null;
	}

	@NotNull
	public ASTNode parse(@NotNull IElementType root, @NotNull PsiBuilder builder) {
		final PsiBuilder.Marker file = builder.mark();
		parseBundle(builder);
		file.done(root);
		return builder.getTreeBuilt();
	}

	@NotNull
	public ASTNode parseBody(IElementType root, PsiBuilder builder) {
		final PsiBuilder.Marker file = builder.mark();
		parseBody(builder);
		file.done(root);
		return builder.getTreeBuilt();
	}

	private void parseBundle(PsiBuilder builder) {
		Marker grammar = builder.mark();

		LtplParserEx parser = new LtplParserEx(builder);
		try {
			parser.parseInput(new LtplLexerEx(builder));
		} catch (IOException e) {
			/* cannot happen */
		} catch (ParseException e) {
			/* syntax error, ok */
		}

		boolean cannotRecover = !parser.markers.isEmpty();
		while (!parser.markers.isEmpty()) {
			parser.markers.pop().drop();
		}

		while (!builder.eof()) {
			builder.advanceLexer();
		}
		grammar.done(LtplElementTypes.BUNDLE);
	}

	private void parseBody(PsiBuilder builder) {
		Marker grammar = builder.mark();

		LtplParserEx parser = new LtplParserEx(builder);
		try {
			parser.parseBody(new LtplLexerEx(builder));
		} catch (IOException e) {
			/* cannot happen */
		} catch (ParseException e) {
			/* syntax error, ok */
		}

		boolean cannotRecover = !parser.markers.isEmpty();
		while (!parser.markers.isEmpty()) {
			parser.markers.pop().drop();
		}

		while (!builder.eof()) {
			builder.advanceLexer();
		}
		grammar.done(LtplElementTypes.TEMPLATE_BODY);
	}

	private static class LtplParserEx extends TemplatesParser {

		private final PsiBuilder myBuilder;
		private final Stack<Marker> markers = new Stack<>();

		public LtplParserEx(PsiBuilder builder) {
			super((message, line, offset, endoffset) -> {
				// ignore, errors are reported as syntax_problem productions
			});
			myBuilder = builder;
		}

		private Marker mark() {
			Marker m = myBuilder.mark();
			markers.push(m);
			return m;
		}

		private Marker clone(Marker inner) {
			Marker outer = inner.precede();
			free(inner, false);
			markers.push(outer);
			return outer;
		}

		private void free(Marker m, boolean drop) {
			assert m == markers.pop();
			if (drop) {
				m.drop();
			}
		}

		private void drop(Span sym) {
			if (sym.value != null) {
				Marker m = (Marker) sym.value;
				free(m, true);
				sym.value = null;
			}
		}

		@Override
		protected void shift() throws IOException {
			Marker marker = tmNext.symbol != Tokens.eoi ? mark() : null;
			super.shift();
			tmStack[tmHead].value = marker;
		}

		@Override
		protected void applyRule(Span left, int ruleIndex, int ruleLength) {
			for (int i = 0; i < ruleLength - 1; i++) {
				drop(tmStack[tmHead - i]);
			}
			if (ruleLength > 0) {
				tmStack[tmHead - (ruleLength - 1)].value = null;
			}

			Marker m = (Marker) left.value;
			if (m != null) {
				IElementType elementType = reduceType(left.symbol, ruleIndex);
				if (elementType != null) {
					left.value = clone(m);

					if (left.symbol == Nonterminals.syntax_problem) {
						m.error("syntax error");
					} else {
						m.done(elementType);
					}
				}
			}
			if (left.symbol == Nonterminals.input || left.symbol == Nonterminals.body) {
				drop(left);
			}
		}

		@Override
		protected boolean restore() {
			boolean restored = super.restore();
			if (restored) {
				/* restored after syntax error - mark the location */
				tmStack[tmHead].value = mark();
			}
			return restored;
		}

		@Override
		protected void dispose(Span sym) {
			drop(sym);
		}

		@Override
		protected void cleanup(Span sym) {
			assert sym.value == null;
		}
	}

	private static class LtplLexerEx extends TemplatesLexer {
		private final PsiBuilder myBuilder;
		private Span next;

		public LtplLexerEx(PsiBuilder builder) throws IOException {
			super(null, null);
			myBuilder = builder;
		}

		@Override
		public Span next() throws IOException {
			return nextInternal();
		}

		@Override
		public void reset(CharSequence input) throws IOException {
		}

		private Span nextInternal() {
			if (next != null && !myBuilder.eof()) {
				myBuilder.advanceLexer();
				while (!myBuilder.eof() && myBuilder.getTokenType() == TokenType.BAD_CHARACTER) {
					myBuilder.advanceLexer();
				}
			}
			next = new Span();
			if (myBuilder.eof()) {
				next.symbol = Tokens.eoi;
			} else {
				LtplElementType tokenType = (LtplElementType) myBuilder.getTokenType();
				next.symbol = tokenType.getSymbol();
			}
			return next;
		}
	}
}

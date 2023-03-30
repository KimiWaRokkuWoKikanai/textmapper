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
package org.textmapper.tool.bootstrap.a;

import java.io.IOException;
import java.text.MessageFormat;
import java.util.ArrayList;
import java.util.List;
import org.textmapper.tool.bootstrap.a.SampleALexer.ErrorReporter;
import org.textmapper.tool.bootstrap.a.SampleALexer.Span;
import org.textmapper.tool.bootstrap.a.SampleALexer.Tokens;
import org.textmapper.tool.bootstrap.a.ast.AstClassdef;
import org.textmapper.tool.bootstrap.a.ast.AstClassdeflistItem;
import org.textmapper.tool.bootstrap.a.ast.IAstClassdefNoEoi;

public class SampleAParser {

	public static class ParseException extends Exception {
		private static final long serialVersionUID = 1L;

		public ParseException() {
		}
	}

	private final ErrorReporter reporter;

	public SampleAParser(ErrorReporter reporter) {
		this.reporter = reporter;
	}

	private static final boolean DEBUG_SYNTAX = false;
	private static final int[] tmAction = SampleALexer.unpack_int(15,
		"\uffff\uffff\uffff\uffff\uffff\uffff\0\0\uffff\uffff\ufffd\uffff\4\0\2\0\ufff5\uffff" +
		"\uffff\uffff\3\0\1\0\ufffe\uffff\uffff\uffff\ufffe\uffff");

	private static final int[] tmLalr = SampleALexer.unpack_int(14,
		"\3\0\uffff\uffff\6\0\uffff\uffff\5\0\6\0\uffff\uffff\ufffe\uffff\3\0\uffff\uffff" +
		"\5\0\5\0\uffff\uffff\ufffe\uffff");

	private static final int[] tmGoto = SampleALexer.unpack_int(12,
		"\0\0\2\0\4\0\4\0\14\0\16\0\20\0\22\0\24\0\34\0\36\0\40\0");

	private static final int[] tmFromTo = SampleALexer.unpack_int(32,
		"\15\0\16\0\2\0\4\0\0\0\2\0\1\0\2\0\5\0\2\0\10\0\2\0\4\0\5\0\11\0\13\0\5\0\6\0\0\0" +
		"\14\0\0\0\3\0\1\0\15\0\5\0\7\0\10\0\12\0\5\0\10\0\5\0\11\0");

	private static final int[] tmRuleLen = SampleALexer.unpack_int(7,
		"\1\0\5\0\1\0\2\0\1\0\1\0\0\0");

	private static final int[] tmRuleSymbol = SampleALexer.unpack_int(7,
		"\7\0\10\0\11\0\11\0\11\0\12\0\12\0");

	protected static final String[] tmSymbolNames = new String[] {
		"eoi",
		"identifier",
		"_skip",
		"Lclass",
		"'{'",
		"'}'",
		"error",
		"classdef_no_eoi",
		"classdef",
		"classdeflist",
		"classdeflistopt",
	};

	public interface Nonterminals extends Tokens {
		// non-terminals
		int classdef_no_eoi = 7;
		int classdef = 8;
		int classdeflist = 9;
		int classdeflistopt = 10;
	}

	// set(follow error)
	private static int[] afterErr = {
		3, 5
	};

	/**
	 * -3-n   Lookahead (state id)
	 * -2     Error
	 * -1     Shift
	 * 0..n   Reduce (rule index)
	 */
	protected static int tmAction(int state, int symbol) {
		int p;
		if (tmAction[state] < -2) {
			if (symbol == Tokens.Unavailable_) {
				return -3 - state;
			}
			for (p = -tmAction[state] - 3; tmLalr[p] >= 0; p += 2) {
				if (tmLalr[p] == symbol) {
					break;
				}
			}
			return tmLalr[p + 1];
		}
		return tmAction[state];
	}

	protected static int gotoState(int state, int symbol) {
		int min = tmGoto[symbol], max = tmGoto[symbol + 1];
		int i, e;

		while (min < max) {
			e = (min + max) >> 2 << 1;
			i = tmFromTo[e];
			if (i == state) {
				return tmFromTo[e+1];
			} else if (i < state) {
				min = e + 2;
			} else {
				max = e;
			}
		}
		return -1;
	}

	protected int tmHead;
	protected Span[] tmStack;
	protected Span tmNext;
	protected SampleALexer tmLexer;

	private Object parse(SampleALexer lexer, int initialState, int finalState, boolean noEoi) throws IOException, ParseException {

		tmLexer = lexer;
		tmStack = new Span[1024];
		tmHead = 0;
		int tmShiftsAfterError = 4;

		tmStack[0] = new Span();
		tmStack[0].state = initialState;
		tmNext = tmLexer.next();

		while (tmStack[tmHead].state != finalState) {
			int action = tmAction(tmStack[tmHead].state, tmNext == null ? Tokens.Unavailable_ : tmNext.symbol);
			if (action <= -3 && tmNext == null) {
				tmNext = tmLexer.next();
				action = tmAction(tmStack[tmHead].state, tmNext.symbol);
			}

			if (action >= 0) {
				reduce(action);
			} else if (action == -1) {
				shift(noEoi);
				tmShiftsAfterError++;
			}

			if (action == -2 || tmStack[tmHead].state == -1) {
				if (restore()) {
					if (tmShiftsAfterError >= 4) {
						reporter.error(MessageFormat.format("syntax error before line {0}, column {1}",
								tmLexer.getTokenLine(), tmNext.column), tmNext.line, tmNext.offset, tmNext.column, tmNext.endline, tmNext.endoffset, tmNext.endcolumn);
					}
					if (tmShiftsAfterError <= 1) {
						tmNext = tmLexer.next();
					}
					tmShiftsAfterError = 0;
					continue;
				}
				if (tmHead < 0) {
					tmHead = 0;
					tmStack[0] = new Span();
					tmStack[0].state = initialState;
				}
				break;
			}
		}

		if (tmStack[tmHead].state != finalState) {
			if (tmShiftsAfterError >= 4) {
				reporter.error(MessageFormat.format("syntax error before line {0}, column {1}",
								tmLexer.getTokenLine(), tmNext == null ? tmLexer.getColumn() : tmNext.column), tmNext == null ? tmLexer.getLine() : tmNext.line, tmNext == null ? tmLexer.getOffset() : tmNext.offset, tmNext == null ? tmLexer.getColumn() : tmNext.column, tmNext == null ? tmLexer.getLine() : tmNext.endline, tmNext == null ? tmLexer.getOffset() : tmNext.endoffset, tmNext == null ? tmLexer.getColumn() : tmNext.endcolumn);
			}
			throw new ParseException();
		}
		return tmStack[noEoi ? tmHead : tmHead - 1].value;
	}

	protected boolean restore() throws IOException {
		if (tmNext == null) {
			tmNext = tmLexer.next();
		}
		if (tmNext.symbol == 0) {
			return false;
		}
		while (tmHead >= 0 && gotoState(tmStack[tmHead].state, 6) == -1) {
			dispose(tmStack[tmHead]);
			tmStack[tmHead] = null;
			tmHead--;
		}
		if (tmHead >= 0) {
			tmStack[++tmHead] = new Span();
			tmStack[tmHead].symbol = 6;
			tmStack[tmHead].value = null;
			tmStack[tmHead].state = gotoState(tmStack[tmHead - 1].state, 6);
			tmStack[tmHead].line = tmNext.line;
			tmStack[tmHead].offset = tmNext.offset;
			tmStack[tmHead].column = tmNext.column;
			tmStack[tmHead].endline = tmNext.endline;
			tmStack[tmHead].endoffset = tmNext.endoffset;
			tmStack[tmHead].endcolumn = tmNext.endcolumn;
			return true;
		}
		return false;
	}

	protected void shift(boolean lazy) throws IOException {
		if (tmNext == null) {
			tmNext = tmLexer.next();
		}
		tmStack[++tmHead] = tmNext;
		tmStack[tmHead].state = gotoState(tmStack[tmHead - 1].state, tmNext.symbol);
		if (DEBUG_SYNTAX) {
			System.out.println(MessageFormat.format("shift: {0} ({1})", tmSymbolNames[tmNext.symbol], tmLexer.tokenText()));
		}
		if (tmStack[tmHead].state != -1 && tmNext.symbol != 0) {
			tmNext = lazy ? null : tmLexer.next();
		}
	}

	protected void reduce(int rule) {
		Span left = new Span();
		left.value = (tmRuleLen[rule] != 0) ? tmStack[tmHead + 1 - tmRuleLen[rule]].value : null;
		left.symbol = tmRuleSymbol[rule];
		left.state = 0;
		if (DEBUG_SYNTAX) {
			System.out.println("reduce to " + tmSymbolNames[tmRuleSymbol[rule]]);
		}
		Span startsym = (tmRuleLen[rule] != 0) ? tmStack[tmHead + 1 - tmRuleLen[rule]] : tmNext;
		left.line = startsym == null ? tmLexer.getLine() : startsym.line;
		left.column = startsym == null ? tmLexer.getColumn() : startsym.column;
		left.offset = startsym == null ? tmLexer.getOffset() : startsym.offset;
		left.endline = (tmRuleLen[rule] != 0) ? tmStack[tmHead].endline : tmNext == null ? tmLexer.getLine() : tmNext.line;
		left.endcolumn = (tmRuleLen[rule] != 0) ? tmStack[tmHead].endcolumn : tmNext == null ? tmLexer.getColumn() : tmNext.column;
		left.endoffset = (tmRuleLen[rule] != 0) ? tmStack[tmHead].endoffset : tmNext == null ? tmLexer.getOffset() : tmNext.offset;
		applyRule(left, rule, tmRuleLen[rule]);
		for (int e = tmRuleLen[rule]; e > 0; e--) {
			tmStack[tmHead--] = null;
		}
		tmStack[++tmHead] = left;
		tmStack[tmHead].state = gotoState(tmStack[tmHead - 1].state, left.symbol);
	}

	@SuppressWarnings("unchecked")
	protected void applyRule(Span tmLeft, int ruleIndex, int ruleLength) {
		switch (ruleIndex) {
			case 1:  // classdef : Lclass identifier '{' classdeflistopt '}'
				tmLeft.value = new AstClassdef(
						((String)tmStack[tmHead - 3].value) /* identifier */,
						((List<AstClassdeflistItem>)tmStack[tmHead - 1].value) /* classdeflist */,
						null /* input */, tmStack[tmHead - 4].line, tmStack[tmHead - 4].offset, tmStack[tmHead - 4].column, tmStack[tmHead].endline, tmStack[tmHead].endoffset, tmStack[tmHead].endcolumn);
				break;
			case 2:  // classdeflist : classdef
				tmLeft.value = new ArrayList();
				((List<AstClassdeflistItem>)tmLeft.value).add(new AstClassdeflistItem(
						((AstClassdef)tmStack[tmHead].value) /* classdef */,
						null /* input */, tmStack[tmHead].line, tmStack[tmHead].offset, tmStack[tmHead].column, tmStack[tmHead].endline, tmStack[tmHead].endoffset, tmStack[tmHead].endcolumn));
				break;
			case 3:  // classdeflist : classdeflist classdef
				((List<AstClassdeflistItem>)tmLeft.value).add(new AstClassdeflistItem(
						((AstClassdef)tmStack[tmHead].value) /* classdef */,
						null /* input */, tmStack[tmHead - 1].line, tmStack[tmHead - 1].offset, tmStack[tmHead - 1].column, tmStack[tmHead].endline, tmStack[tmHead].endoffset, tmStack[tmHead].endcolumn));
				break;
			case 4:  // classdeflist : error
				tmLeft.value = new ArrayList();
				((List<AstClassdeflistItem>)tmLeft.value).add(new AstClassdeflistItem(
						null /* classdef */,
						null /* input */, tmStack[tmHead].line, tmStack[tmHead].offset, tmStack[tmHead].column, tmStack[tmHead].endline, tmStack[tmHead].endoffset, tmStack[tmHead].endcolumn));
				break;
		}
	}

	/**
	 * disposes symbol dropped by error recovery mechanism
	 */
	protected void dispose(Span value) {
	}

	public IAstClassdefNoEoi parseClassdef_no_eoi(SampleALexer lexer) throws IOException, ParseException {
		return (IAstClassdefNoEoi) parse(lexer, 0, 12, true);
	}

	public AstClassdef parseClassdef(SampleALexer lexer) throws IOException, ParseException {
		return (AstClassdef) parse(lexer, 1, 14, false);
	}
}

package org.textmapper.json;

import java.io.IOException;
import java.text.MessageFormat;
import org.textmapper.json.JsonLexer.ErrorReporter;
import org.textmapper.json.JsonLexer.Span;
import org.textmapper.json.JsonLexer.Tokens;

public class JsonParser {

	public static class ParseException extends Exception {
		private static final long serialVersionUID = 1L;

		public ParseException() {
		}
	}

	private final ErrorReporter reporter;

	public JsonParser(ErrorReporter reporter) {
		this.reporter = reporter;
	}

	private static final boolean DEBUG_SYNTAX = false;
	private static final int[] tmAction = JsonLexer.unpack_int(28,
		"\uffff\uffff\uffff\uffff\uffff\uffff\6\0\7\0\1\0\2\0\3\0\0\0\4\0\5\0\11\0\uffff\uffff" +
		"\13\0\uffff\uffff\16\0\17\0\uffff\uffff\uffff\uffff\10\0\uffff\uffff\15\0\uffff\uffff" +
		"\12\0\14\0\20\0\uffff\uffff\ufffe\uffff");

	private static final int[] tmGoto = JsonLexer.unpack_int(21,
		"\0\0\2\0\12\0\16\0\26\0\32\0\34\0\40\0\40\0\54\0\64\0\74\0\104\0\114\0\116\0\126" +
		"\0\136\0\142\0\144\0\154\0\156\0");

	private static final int[] tmFromTo = JsonLexer.unpack_int(110,
		"\32\0\33\0\0\0\1\0\2\0\1\0\22\0\1\0\26\0\1\0\1\0\13\0\16\0\23\0\0\0\2\0\2\0\2\0\22" +
		"\0\2\0\26\0\2\0\2\0\17\0\21\0\25\0\14\0\22\0\16\0\24\0\21\0\26\0\0\0\3\0\1\0\14\0" +
		"\2\0\3\0\22\0\3\0\24\0\14\0\26\0\3\0\0\0\4\0\2\0\4\0\22\0\4\0\26\0\4\0\0\0\5\0\2" +
		"\0\5\0\22\0\5\0\26\0\5\0\0\0\6\0\2\0\6\0\22\0\6\0\26\0\6\0\0\0\7\0\2\0\7\0\22\0\7" +
		"\0\26\0\7\0\0\0\32\0\0\0\10\0\2\0\20\0\22\0\27\0\26\0\31\0\0\0\11\0\2\0\11\0\22\0" +
		"\11\0\26\0\11\0\1\0\15\0\24\0\30\0\1\0\16\0\0\0\12\0\2\0\12\0\22\0\12\0\26\0\12\0" +
		"\2\0\21\0");

	private static final int[] tmRuleLen = JsonLexer.unpack_int(17,
		"\1\0\1\0\1\0\1\0\1\0\1\0\1\0\1\0\3\0\2\0\3\0\1\0\3\0\3\0\2\0\1\0\3\0");

	private static final int[] tmRuleSymbol = JsonLexer.unpack_int(17,
		"\15\0\16\0\16\0\16\0\16\0\16\0\16\0\16\0\17\0\17\0\20\0\21\0\21\0\22\0\22\0\23\0" +
		"\23\0");

	protected static final String[] tmSymbolNames = new String[] {
		"eoi",
		"'{'",
		"'}'",
		"'['",
		"']'",
		"':'",
		"','",
		"space",
		"JSONString",
		"JSONNumber",
		"'null'",
		"'true'",
		"'false'",
		"JSONText",
		"JSONValue",
		"JSONObject",
		"JSONMember",
		"JSONMemberList",
		"JSONArray",
		"JSONElementList",
	};

	public interface Nonterminals extends Tokens {
		// non-terminals
		int JSONText = 13;
		int JSONValue = 14;
		int JSONObject = 15;
		int JSONMember = 16;
		int JSONMemberList = 17;
		int JSONArray = 18;
		int JSONElementList = 19;
	}

	/**
	 * -3-n   Lookahead (state id)
	 * -2     Error
	 * -1     Shift
	 * 0..n   Reduce (rule index)
	 */
	protected static int tmAction(int state, int symbol) {
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
	protected JsonLexer tmLexer;

	public Object parse(JsonLexer lexer) throws IOException, ParseException {

		tmLexer = lexer;
		tmStack = new Span[1024];
		tmHead = 0;

		tmStack[0] = new Span();
		tmStack[0].state = 0;
		tmNext = tmLexer.next();

		while (tmStack[tmHead].state != 27) {
			int action = tmAction(tmStack[tmHead].state, tmNext.symbol);

			if (action >= 0) {
				reduce(action);
			} else if (action == -1) {
				shift();
			}

			if (action == -2 || tmStack[tmHead].state == -1) {
				break;
			}
		}

		if (tmStack[tmHead].state != 27) {
			reporter.error(MessageFormat.format("syntax error before line {0}",
								tmLexer.getTokenLine()), tmNext.line, tmNext.offset, tmNext.endoffset);
			throw new ParseException();
		}
		return tmStack[tmHead - 1].value;
	}

	protected void shift() throws IOException {
		tmStack[++tmHead] = tmNext;
		tmStack[tmHead].state = gotoState(tmStack[tmHead - 1].state, tmNext.symbol);
		if (DEBUG_SYNTAX) {
			System.out.println(MessageFormat.format("shift: {0} ({1})", tmSymbolNames[tmNext.symbol], tmLexer.tokenText()));
		}
		if (tmStack[tmHead].state != -1 && tmNext.symbol != 0) {
			tmNext = tmLexer.next();
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
		left.line = startsym.line;
		left.offset = startsym.offset;
		left.endoffset = (tmRuleLen[rule] != 0) ? tmStack[tmHead].endoffset : tmNext.offset;
		applyRule(left, rule, tmRuleLen[rule]);
		for (int e = tmRuleLen[rule]; e > 0; e--) {
			tmStack[tmHead--] = null;
		}
		tmStack[++tmHead] = left;
		tmStack[tmHead].state = gotoState(tmStack[tmHead - 1].state, left.symbol);
	}

	@SuppressWarnings("unchecked")
	protected void applyRule(Span tmLeft, int ruleIndex, int ruleLength) {
	}
}

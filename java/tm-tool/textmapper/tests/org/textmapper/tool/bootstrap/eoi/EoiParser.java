package org.textmapper.tool.bootstrap.eoi;

import java.io.IOException;
import java.text.MessageFormat;
import org.textmapper.tool.bootstrap.eoi.EoiLexer.ErrorReporter;
import org.textmapper.tool.bootstrap.eoi.EoiLexer.Span;
import org.textmapper.tool.bootstrap.eoi.EoiLexer.Tokens;

public class EoiParser {

	public static class ParseException extends Exception {
		private static final long serialVersionUID = 1L;

		public ParseException() {
		}
	}

	private final ErrorReporter reporter;

	public EoiParser(ErrorReporter reporter) {
		this.reporter = reporter;
	}

	private static final boolean DEBUG_SYNTAX = false;
	private static final int[] tmAction = EoiLexer.unpack_int(17,
		"\uffff\uffff\1\0\ufffd\uffff\0\0\uffff\uffff\ufff7\uffff\uffff\uffff\uffff\uffff" +
		"\uffff\uffff\ufff1\uffff\5\0\uffff\uffff\2\0\uffff\uffff\4\0\uffff\uffff\ufffe\uffff");

	private static final int[] tmLalr = EoiLexer.unpack_int(22,
		"\1\0\uffff\uffff\3\0\7\0\uffff\uffff\ufffe\uffff\4\0\uffff\uffff\3\0\6\0\uffff\uffff" +
		"\ufffe\uffff\10\0\uffff\uffff\0\0\3\0\3\0\3\0\4\0\3\0\uffff\uffff\ufffe\uffff");

	private static final int[] tmGoto = EoiLexer.unpack_int(17,
		"\0\0\2\0\14\0\20\0\22\0\24\0\24\0\24\0\32\0\34\0\34\0\34\0\34\0\36\0\44\0\46\0\50" +
		"\0");

	private static final int[] tmFromTo = EoiLexer.unpack_int(40,
		"\17\0\20\0\0\0\1\0\2\0\4\0\7\0\1\0\10\0\13\0\15\0\1\0\4\0\7\0\13\0\15\0\6\0\11\0" +
		"\5\0\10\0\0\0\2\0\7\0\2\0\15\0\2\0\11\0\14\0\0\0\17\0\0\0\3\0\7\0\12\0\15\0\16\0" +
		"\2\0\5\0\2\0\6\0");

	private static final int[] tmRuleLen = EoiLexer.unpack_int(8,
		"\1\0\1\0\4\0\3\0\5\0\3\0\1\0\0\0");

	private static final int[] tmRuleSymbol = EoiLexer.unpack_int(8,
		"\14\0\15\0\15\0\15\0\16\0\16\0\17\0\17\0");

	protected static final String[] tmSymbolNames = new String[] {
		"eoi",
		"id",
		"':'",
		"';'",
		"','",
		"gotoc",
		"_skip",
		"'('",
		"')'",
		"_customEOI",
		"_retfromA",
		"_retfromB",
		"input",
		"expr",
		"list_of_ID_and_2_elements_Comma_separated",
		"list_of_ID_and_2_elements_Comma_separatedopt",
	};

	public interface Nonterminals extends Tokens {
		// non-terminals
		int input = 12;
		int expr = 13;
		int list_of_ID_and_2_elements_Comma_separated = 14;
		int list_of_ID_and_2_elements_Comma_separatedopt = 15;
	}

	/**
	 * -3-n   Lookahead (state id)
	 * -2     Error
	 * -1     Shift
	 * 0..n   Reduce (rule index)
	 */
	protected static int tmAction(int state, int symbol) {
		int p;
		if (tmAction[state] < -2) {
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
	protected EoiLexer tmLexer;

	public Object parse(EoiLexer lexer) throws IOException, ParseException {

		tmLexer = lexer;
		tmStack = new Span[1024];
		tmHead = 0;

		tmStack[0] = new Span();
		tmStack[0].state = 0;
		tmNext = tmLexer.next();

		while (tmStack[tmHead].state != 16) {
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

		if (tmStack[tmHead].state != 16) {
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

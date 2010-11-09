package org.textway.templates.types;

import java.io.IOException;
import java.io.Reader;
import java.text.MessageFormat;

public class TypesLexer {

	public static class LapgSymbol {
		public Object sym;
		public int lexem;
		public int state;
		public int line;
		public int offset;
		public int endoffset;
	}

	public interface Lexems {
		public static final int eoi = 0;
		public static final int identifier = 1;
		public static final int scon = 2;
		public static final int icon = 3;
		public static final int bcon = 4;
		public static final int _skip = 5;
		public static final int DOTDOT = 6;
		public static final int DOT = 7;
		public static final int MULT = 8;
		public static final int SEMICOLON = 9;
		public static final int COMMA = 10;
		public static final int COLON = 11;
		public static final int EQUAL = 12;
		public static final int LCURLY = 13;
		public static final int RCURLY = 14;
		public static final int LPAREN = 15;
		public static final int RPAREN = 16;
		public static final int LSQUARE = 17;
		public static final int RSQUARE = 18;
		public static final int Lclass = 19;
		public static final int Lextends = 20;
		public static final int Lnew = 21;
		public static final int Lint = 22;
		public static final int Lbool = 23;
		public static final int Lstring = 24;
		public static final int Lset = 25;
		public static final int Lchoice = 26;
	}
	
	public interface ErrorReporter {
		void error(int start, int end, int line, String s);
	}

	public static final int TOKEN_SIZE = 2048;

	private Reader stream;
	final private ErrorReporter reporter;

	final private char[] data = new char[2048];
	private int datalen, l;
	private char chr;

	private int group;

	final private StringBuilder token = new StringBuilder(TOKEN_SIZE);

	private int tokenLine = 1;
	private int currLine = 1;
	private int currOffset = 0;
	
	private String unescape(String s, int start, int end) {
		StringBuilder sb = new StringBuilder();
		end = Math.min(end, s.length());
		for(int i = start; i < end; i++) {
			char c = s.charAt(i);
			if(c == '\\') {
				if(++i == end) {
					break;
				}
				c = s.charAt(i);
				if(c == 'u' || c == 'x') {
					// FIXME process unicode
				} else if(c == 'n') {
					sb.append('\n');
				} else if(c == 'r') {
					sb.append('\r');
				} else if(c == 't') {
					sb.append('\t');
				} else {
					sb.append(c);
				}
			} else {
				sb.append(c);
			}
		} 
		return sb.toString();
	}

	public TypesLexer(Reader stream, ErrorReporter reporter) throws IOException {
		this.reporter = reporter;
		reset(stream);
	}

	public void reset(Reader stream) throws IOException {
		this.stream = stream;
		this.datalen = stream.read(data);
		this.l = 0;
		this.group = 0;
		chr = l < datalen ? data[l++] : 0;
	}

	public int getState() {
		return group;
	}

	public void setState(int state) {
		this.group = state;
	}

	public int getTokenLine() {
		return tokenLine;
	}

	public void setLine(int currLine) {
		this.currLine = currLine;
	}

	public void setOffset(int currOffset) {
		this.currOffset = currOffset;
	}

	public String current() {
		return token.toString();
	}

	private static final short lapg_char2no[] = {
		0, 1, 1, 1, 1, 1, 1, 1, 1, 40, 4, 1, 1, 40, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		40, 1, 5, 15, 1, 1, 1, 2, 24, 25, 17, 1, 19, 6, 16, 1,
		39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 20, 18, 1, 21, 1, 1,
		1, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
		38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 26, 3, 27, 1, 38,
		1, 12, 34, 28, 31, 10, 11, 36, 37, 33, 38, 38, 13, 38, 30, 35,
		38, 38, 8, 14, 7, 9, 38, 32, 29, 38, 38, 22, 1, 23, 1, 1,
	};

	private static final short[][] lapg_lexem = new short[][] {
		{ -2, -1, 1, -1, 2, 3, 4, 5, 6, 6, 7, 8, 6, 6, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 6, 24, 6, 6, 25, 26, 6, 6, 6, 6, 27, 2, },
		{ -1, 1, 28, 29, -1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, },
		{ -7, -7, -7, -7, 2, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, 2, },
		{ -1, 3, 3, 30, -1, 31, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, },
		{ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 27, -1, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 32, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 33, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 34, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 35, 6, 6, 36, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -7, 10, 10, 10, -7, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, },
		{ -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, 37, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, },
		{ -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, },
		{ -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, },
		{ -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, },
		{ -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, },
		{ -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, },
		{ -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, },
		{ -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, },
		{ -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, },
		{ -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, },
		{ -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, },
		{ -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 38, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 39, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 40, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 41, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 42, 6, 6, 6, 6, -3, },
		{ -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 27, -5, },
		{ -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, },
		{ -1, 1, 1, 1, -1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, },
		{ -1, 3, 3, 3, -1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, },
		{ -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 43, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 44, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 45, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 46, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 47, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 48, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 49, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 50, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 51, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 52, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 53, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 54, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 55, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 56, 6, 6, 6, 6, 6, 6, -3, },
		{ -27, -27, -27, -27, -27, -27, -27, 6, 6, 6, 6, 6, 6, 6, 6, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -27, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 57, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 58, 6, 6, 6, 6, 6, 6, -3, },
		{ -23, -23, -23, -23, -23, -23, -23, 6, 6, 6, 6, 6, 6, 6, 6, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -23, },
		{ -24, -24, -24, -24, -24, -24, -24, 6, 6, 6, 6, 6, 6, 6, 6, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -24, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 59, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -6, -6, -6, -6, -6, -6, -6, 6, 6, 6, 6, 6, 6, 6, 6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -6, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 60, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 53, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 61, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 62, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 63, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -25, -25, -25, -25, -25, -25, -25, 6, 6, 6, 6, 6, 6, 6, 6, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -25, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 64, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 65, 6, 6, 6, -3, },
		{ -21, -21, -21, -21, -21, -21, -21, 6, 6, 6, 6, 6, 6, 6, 6, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -21, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 66, 6, 6, 6, 6, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 67, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -3, },
		{ -26, -26, -26, -26, -26, -26, -26, 6, 6, 6, 6, 6, 6, 6, 6, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -26, },
		{ -28, -28, -28, -28, -28, -28, -28, 6, 6, 6, 6, 6, 6, 6, 6, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -28, },
		{ -22, -22, -22, -22, -22, -22, -22, 6, 6, 6, 6, 6, 6, 6, 6, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, -22, },

	};

	private static int mapCharacter(int chr) {
		if (chr >= 0 && chr < 128) {
			return lapg_char2no[chr];
		}
		return 1;
	}

	public LapgSymbol next() throws IOException {
		LapgSymbol lapg_n = new LapgSymbol();
		int state;

		do {
			lapg_n.offset = currOffset;
			tokenLine = lapg_n.line = currLine;
			if (token.length() > TOKEN_SIZE) {
				token.setLength(TOKEN_SIZE);
				token.trimToSize();
			}
			token.setLength(0);
			int tokenStart = l - 1;

			for (state = group; state >= 0;) {
				state = lapg_lexem[state][mapCharacter(chr)];
				if (state >= -1 && chr != 0) {
					currOffset++;
					if (chr == '\n') {
						currLine++;
					}
					if (l >= datalen) {
						token.append(data, tokenStart, l - tokenStart);
						datalen = stream.read(data);
						tokenStart = l = 0;
					}
					chr = l < datalen ? data[l++] : 0;
				}
			}
			lapg_n.endoffset = currOffset;

			if (state == -1) {
				if (chr == 0) {
					reporter.error(lapg_n.offset, lapg_n.endoffset, currLine, "Unexpected end of file reached");
					break;
				}
				reporter.error(lapg_n.offset, lapg_n.endoffset, currLine, MessageFormat.format("invalid lexem at line {0}: `{1}`, skipped", currLine, current()));
				lapg_n.lexem = -1;
				continue;
			}

			if (l - 1 > tokenStart) {
				token.append(data, tokenStart, l - 1 - tokenStart);
			}

			lapg_n.lexem = - state - 2;
			lapg_n.sym = null;

		} while (lapg_n.lexem == -1 || !createToken(lapg_n));
		return lapg_n;
	}

	protected boolean createToken(LapgSymbol lapg_n) {
		switch (lapg_n.lexem) {
			case 1:
				 lapg_n.sym = current(); break; 
			case 2:
				 lapg_n.sym = unescape(current(), 1, token.length()-1); break; 
			case 3:
				 lapg_n.sym = Integer.parseInt(current()); break; 
			case 4:
				 lapg_n.sym = current().equals("true"); break; 
			case 5:
				 return false; 
		}
		return true;
	}
}

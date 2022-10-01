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
package org.textmapper.tool.importer;

import java.io.IOException;
import java.text.MessageFormat;
import org.textmapper.tool.importer.BisonLexer.ErrorReporter;
import org.textmapper.tool.importer.BisonLexer.Span;
import org.textmapper.tool.importer.BisonLexer.Tokens;

public class BisonParser {

	public static class ParseException extends Exception {
		private static final long serialVersionUID = 1L;

		public ParseException() {
		}
	}

	private final ErrorReporter reporter;

	public BisonParser(ErrorReporter reporter) {
		this.reporter = reporter;
	}

	private static final boolean DEBUG_SYNTAX = false;
	private static final int[] tmAction = BisonLexer.unpack_int(141,
		"\4\0\uffff\uffff\uffff\uffff\37\0\uffff\uffff\uffff\uffff\uffff\uffff\54\0\55\0\70" +
		"\0\71\0\72\0\73\0\uffff\uffff\44\0\uffff\uffff\ufffd\uffff\16\0\uffff\uffff\uffff" +
		"\uffff\12\0\uffff\uffff\22\0\uffff\uffff\uffff\uffff\uffff\uffff\45\0\26\0\27\0\uffff" +
		"\uffff\uffff\uffff\uffff\uffff\uffff\uffff\uffff\uffff\34\0\uffff\uffff\35\0\36\0" +
		"\7\0\3\0\6\0\5\0\uffff\uffff\41\0\40\0\uffaf\uffff\uffa5\uffff\uff6b\uffff\uffff" +
		"\uffff\1\0\116\0\uff49\uffff\ufef3\uffff\105\0\ufe9d\uffff\62\0\ufe4b\uffff\uffff" +
		"\uffff\uffff\uffff\46\0\137\0\140\0\ufdf9\uffff\15\0\17\0\20\0\21\0\23\0\24\0\25" +
		"\0\30\0\11\0\ufda7\uffff\32\0\33\0\144\0\145\0\146\0\43\0\uffff\uffff\51\0\uffff" +
		"\uffff\74\0\uffff\uffff\uffff\uffff\122\0\0\0\117\0\ufd59\uffff\110\0\ufd05\uffff" +
		"\114\0\61\0\ufcb1\uffff\64\0\47\0\141\0\142\0\143\0\147\0\13\0\10\0\50\0\103\0\104" +
		"\0\102\0\ufc5f\uffff\53\0\101\0\100\0\ufc07\uffff\67\0\ufbb5\uffff\uffff\uffff\ufb61" +
		"\uffff\ufb29\uffff\106\0\112\0\63\0\52\0\66\0\76\0\135\0\uffff\uffff\uffff\uffff" +
		"\uffff\uffff\131\0\ufb03\uffff\130\0\121\0\ufac9\uffff\122\0\125\0\132\0\133\0\134" +
		"\0\127\0\126\0\ufa8f\uffff\uffff\uffff\ufffe\uffff");

	private static final int[] tmLalr = BisonLexer.unpack_int(1446,
		"\6\0\uffff\uffff\11\0\14\0\13\0\14\0\20\0\14\0\21\0\14\0\22\0\14\0\23\0\14\0\24\0" +
		"\14\0\25\0\14\0\26\0\14\0\27\0\14\0\30\0\14\0\34\0\14\0\35\0\14\0\36\0\14\0\37\0" +
		"\14\0\41\0\14\0\42\0\14\0\43\0\14\0\44\0\14\0\45\0\14\0\46\0\14\0\47\0\14\0\50\0" +
		"\14\0\51\0\14\0\52\0\14\0\53\0\14\0\54\0\14\0\55\0\14\0\56\0\14\0\57\0\14\0\60\0" +
		"\14\0\61\0\14\0\62\0\14\0\63\0\14\0\64\0\14\0\65\0\14\0\70\0\14\0\uffff\uffff\ufffe" +
		"\uffff\73\0\uffff\uffff\2\0\75\0\5\0\75\0\6\0\75\0\uffff\uffff\ufffe\uffff\14\0\uffff" +
		"\uffff\0\0\136\0\1\0\136\0\2\0\136\0\5\0\136\0\6\0\136\0\12\0\136\0\13\0\136\0\20" +
		"\0\136\0\21\0\136\0\22\0\136\0\23\0\136\0\24\0\136\0\25\0\136\0\26\0\136\0\27\0\136" +
		"\0\30\0\136\0\31\0\136\0\32\0\136\0\33\0\136\0\34\0\136\0\35\0\136\0\40\0\136\0\52" +
		"\0\136\0\61\0\136\0\63\0\136\0\66\0\136\0\67\0\136\0\uffff\uffff\ufffe\uffff\1\0" +
		"\uffff\uffff\20\0\uffff\uffff\21\0\uffff\uffff\22\0\uffff\uffff\23\0\uffff\uffff" +
		"\24\0\uffff\uffff\25\0\uffff\uffff\26\0\uffff\uffff\27\0\uffff\uffff\30\0\uffff\uffff" +
		"\34\0\uffff\uffff\35\0\uffff\uffff\52\0\uffff\uffff\61\0\uffff\uffff\63\0\uffff\uffff" +
		"\0\0\2\0\uffff\uffff\ufffe\uffff\4\0\uffff\uffff\6\0\uffff\uffff\2\0\111\0\5\0\111" +
		"\0\11\0\111\0\13\0\111\0\20\0\111\0\21\0\111\0\22\0\111\0\23\0\111\0\24\0\111\0\25" +
		"\0\111\0\26\0\111\0\27\0\111\0\30\0\111\0\34\0\111\0\35\0\111\0\36\0\111\0\37\0\111" +
		"\0\41\0\111\0\42\0\111\0\43\0\111\0\44\0\111\0\45\0\111\0\46\0\111\0\47\0\111\0\50" +
		"\0\111\0\51\0\111\0\52\0\111\0\53\0\111\0\54\0\111\0\55\0\111\0\56\0\111\0\57\0\111" +
		"\0\60\0\111\0\61\0\111\0\62\0\111\0\63\0\111\0\64\0\111\0\65\0\111\0\70\0\111\0\73" +
		"\0\111\0\uffff\uffff\ufffe\uffff\4\0\uffff\uffff\6\0\uffff\uffff\2\0\115\0\5\0\115" +
		"\0\11\0\115\0\13\0\115\0\20\0\115\0\21\0\115\0\22\0\115\0\23\0\115\0\24\0\115\0\25" +
		"\0\115\0\26\0\115\0\27\0\115\0\30\0\115\0\34\0\115\0\35\0\115\0\36\0\115\0\37\0\115" +
		"\0\41\0\115\0\42\0\115\0\43\0\115\0\44\0\115\0\45\0\115\0\46\0\115\0\47\0\115\0\50" +
		"\0\115\0\51\0\115\0\52\0\115\0\53\0\115\0\54\0\115\0\55\0\115\0\56\0\115\0\57\0\115" +
		"\0\60\0\115\0\61\0\115\0\62\0\115\0\63\0\115\0\64\0\115\0\65\0\115\0\70\0\115\0\73" +
		"\0\115\0\uffff\uffff\ufffe\uffff\2\0\uffff\uffff\5\0\uffff\uffff\73\0\uffff\uffff" +
		"\11\0\57\0\13\0\57\0\20\0\57\0\21\0\57\0\22\0\57\0\23\0\57\0\24\0\57\0\25\0\57\0" +
		"\26\0\57\0\27\0\57\0\30\0\57\0\34\0\57\0\35\0\57\0\36\0\57\0\37\0\57\0\41\0\57\0" +
		"\42\0\57\0\43\0\57\0\44\0\57\0\45\0\57\0\46\0\57\0\47\0\57\0\50\0\57\0\51\0\57\0" +
		"\52\0\57\0\53\0\57\0\54\0\57\0\55\0\57\0\56\0\57\0\57\0\57\0\60\0\57\0\61\0\57\0" +
		"\62\0\57\0\63\0\57\0\64\0\57\0\65\0\57\0\70\0\57\0\uffff\uffff\ufffe\uffff\2\0\uffff" +
		"\uffff\5\0\uffff\uffff\73\0\uffff\uffff\11\0\56\0\13\0\56\0\20\0\56\0\21\0\56\0\22" +
		"\0\56\0\23\0\56\0\24\0\56\0\25\0\56\0\26\0\56\0\27\0\56\0\30\0\56\0\34\0\56\0\35" +
		"\0\56\0\36\0\56\0\37\0\56\0\41\0\56\0\42\0\56\0\43\0\56\0\44\0\56\0\45\0\56\0\46" +
		"\0\56\0\47\0\56\0\50\0\56\0\51\0\56\0\52\0\56\0\53\0\56\0\54\0\56\0\55\0\56\0\56" +
		"\0\56\0\57\0\56\0\60\0\56\0\61\0\56\0\62\0\56\0\63\0\56\0\64\0\56\0\65\0\56\0\70" +
		"\0\56\0\uffff\uffff\ufffe\uffff\2\0\uffff\uffff\6\0\uffff\uffff\66\0\uffff\uffff" +
		"\11\0\150\0\13\0\150\0\20\0\150\0\21\0\150\0\22\0\150\0\23\0\150\0\24\0\150\0\25" +
		"\0\150\0\26\0\150\0\27\0\150\0\30\0\150\0\34\0\150\0\35\0\150\0\36\0\150\0\37\0\150" +
		"\0\41\0\150\0\42\0\150\0\43\0\150\0\44\0\150\0\45\0\150\0\46\0\150\0\47\0\150\0\50" +
		"\0\150\0\51\0\150\0\52\0\150\0\53\0\150\0\54\0\150\0\55\0\150\0\56\0\150\0\57\0\150" +
		"\0\60\0\150\0\61\0\150\0\62\0\150\0\63\0\150\0\64\0\150\0\65\0\150\0\70\0\150\0\uffff" +
		"\uffff\ufffe\uffff\66\0\uffff\uffff\11\0\31\0\13\0\31\0\20\0\31\0\21\0\31\0\22\0" +
		"\31\0\23\0\31\0\24\0\31\0\25\0\31\0\26\0\31\0\27\0\31\0\30\0\31\0\34\0\31\0\35\0" +
		"\31\0\36\0\31\0\37\0\31\0\41\0\31\0\42\0\31\0\43\0\31\0\44\0\31\0\45\0\31\0\46\0" +
		"\31\0\47\0\31\0\50\0\31\0\51\0\31\0\52\0\31\0\53\0\31\0\54\0\31\0\55\0\31\0\56\0" +
		"\31\0\57\0\31\0\60\0\31\0\61\0\31\0\62\0\31\0\63\0\31\0\64\0\31\0\65\0\31\0\70\0" +
		"\31\0\uffff\uffff\ufffe\uffff\6\0\uffff\uffff\2\0\107\0\5\0\107\0\11\0\107\0\13\0" +
		"\107\0\20\0\107\0\21\0\107\0\22\0\107\0\23\0\107\0\24\0\107\0\25\0\107\0\26\0\107" +
		"\0\27\0\107\0\30\0\107\0\34\0\107\0\35\0\107\0\36\0\107\0\37\0\107\0\41\0\107\0\42" +
		"\0\107\0\43\0\107\0\44\0\107\0\45\0\107\0\46\0\107\0\47\0\107\0\50\0\107\0\51\0\107" +
		"\0\52\0\107\0\53\0\107\0\54\0\107\0\55\0\107\0\56\0\107\0\57\0\107\0\60\0\107\0\61" +
		"\0\107\0\62\0\107\0\63\0\107\0\64\0\107\0\65\0\107\0\70\0\107\0\73\0\107\0\uffff" +
		"\uffff\ufffe\uffff\6\0\uffff\uffff\2\0\113\0\5\0\113\0\11\0\113\0\13\0\113\0\20\0" +
		"\113\0\21\0\113\0\22\0\113\0\23\0\113\0\24\0\113\0\25\0\113\0\26\0\113\0\27\0\113" +
		"\0\30\0\113\0\34\0\113\0\35\0\113\0\36\0\113\0\37\0\113\0\41\0\113\0\42\0\113\0\43" +
		"\0\113\0\44\0\113\0\45\0\113\0\46\0\113\0\47\0\113\0\50\0\113\0\51\0\113\0\52\0\113" +
		"\0\53\0\113\0\54\0\113\0\55\0\113\0\56\0\113\0\57\0\113\0\60\0\113\0\61\0\113\0\62" +
		"\0\113\0\63\0\113\0\64\0\113\0\65\0\113\0\70\0\113\0\73\0\113\0\uffff\uffff\ufffe" +
		"\uffff\2\0\uffff\uffff\5\0\uffff\uffff\6\0\uffff\uffff\11\0\60\0\13\0\60\0\20\0\60" +
		"\0\21\0\60\0\22\0\60\0\23\0\60\0\24\0\60\0\25\0\60\0\26\0\60\0\27\0\60\0\30\0\60" +
		"\0\34\0\60\0\35\0\60\0\36\0\60\0\37\0\60\0\41\0\60\0\42\0\60\0\43\0\60\0\44\0\60" +
		"\0\45\0\60\0\46\0\60\0\47\0\60\0\50\0\60\0\51\0\60\0\52\0\60\0\53\0\60\0\54\0\60" +
		"\0\55\0\60\0\56\0\60\0\57\0\60\0\60\0\60\0\61\0\60\0\62\0\60\0\63\0\60\0\64\0\60" +
		"\0\65\0\60\0\70\0\60\0\uffff\uffff\ufffe\uffff\2\0\uffff\uffff\5\0\uffff\uffff\6" +
		"\0\uffff\uffff\7\0\uffff\uffff\10\0\uffff\uffff\73\0\uffff\uffff\11\0\42\0\13\0\42" +
		"\0\20\0\42\0\21\0\42\0\22\0\42\0\23\0\42\0\24\0\42\0\25\0\42\0\26\0\42\0\27\0\42" +
		"\0\30\0\42\0\34\0\42\0\35\0\42\0\36\0\42\0\37\0\42\0\41\0\42\0\42\0\42\0\43\0\42" +
		"\0\44\0\42\0\45\0\42\0\46\0\42\0\47\0\42\0\50\0\42\0\51\0\42\0\52\0\42\0\53\0\42" +
		"\0\54\0\42\0\55\0\42\0\56\0\42\0\57\0\42\0\60\0\42\0\61\0\42\0\62\0\42\0\63\0\42" +
		"\0\64\0\42\0\65\0\42\0\70\0\42\0\uffff\uffff\ufffe\uffff\2\0\uffff\uffff\5\0\uffff" +
		"\uffff\6\0\uffff\uffff\11\0\65\0\13\0\65\0\20\0\65\0\21\0\65\0\22\0\65\0\23\0\65" +
		"\0\24\0\65\0\25\0\65\0\26\0\65\0\27\0\65\0\30\0\65\0\34\0\65\0\35\0\65\0\36\0\65" +
		"\0\37\0\65\0\41\0\65\0\42\0\65\0\43\0\65\0\44\0\65\0\45\0\65\0\46\0\65\0\47\0\65" +
		"\0\50\0\65\0\51\0\65\0\52\0\65\0\53\0\65\0\54\0\65\0\55\0\65\0\56\0\65\0\57\0\65" +
		"\0\60\0\65\0\61\0\65\0\62\0\65\0\63\0\65\0\64\0\65\0\65\0\65\0\70\0\65\0\uffff\uffff" +
		"\ufffe\uffff\4\0\uffff\uffff\2\0\77\0\5\0\77\0\6\0\77\0\11\0\77\0\13\0\77\0\20\0" +
		"\77\0\21\0\77\0\22\0\77\0\23\0\77\0\24\0\77\0\25\0\77\0\26\0\77\0\27\0\77\0\30\0" +
		"\77\0\34\0\77\0\35\0\77\0\36\0\77\0\37\0\77\0\41\0\77\0\42\0\77\0\43\0\77\0\44\0" +
		"\77\0\45\0\77\0\46\0\77\0\47\0\77\0\50\0\77\0\51\0\77\0\52\0\77\0\53\0\77\0\54\0" +
		"\77\0\55\0\77\0\56\0\77\0\57\0\77\0\60\0\77\0\61\0\77\0\62\0\77\0\63\0\77\0\64\0" +
		"\77\0\65\0\77\0\70\0\77\0\uffff\uffff\ufffe\uffff\2\0\uffff\uffff\5\0\uffff\uffff" +
		"\6\0\uffff\uffff\31\0\uffff\uffff\32\0\uffff\uffff\33\0\uffff\uffff\40\0\uffff\uffff" +
		"\66\0\uffff\uffff\67\0\uffff\uffff\0\0\123\0\1\0\123\0\12\0\123\0\13\0\123\0\20\0" +
		"\123\0\21\0\123\0\22\0\123\0\23\0\123\0\24\0\123\0\25\0\123\0\26\0\123\0\27\0\123" +
		"\0\30\0\123\0\34\0\123\0\35\0\123\0\52\0\123\0\61\0\123\0\63\0\123\0\uffff\uffff" +
		"\ufffe\uffff\12\0\uffff\uffff\13\0\uffff\uffff\0\0\120\0\1\0\120\0\20\0\120\0\21" +
		"\0\120\0\22\0\120\0\23\0\120\0\24\0\120\0\25\0\120\0\26\0\120\0\27\0\120\0\30\0\120" +
		"\0\34\0\120\0\35\0\120\0\52\0\120\0\61\0\120\0\63\0\120\0\uffff\uffff\ufffe\uffff" +
		"\14\0\uffff\uffff\0\0\136\0\1\0\136\0\2\0\136\0\5\0\136\0\6\0\136\0\12\0\136\0\13" +
		"\0\136\0\20\0\136\0\21\0\136\0\22\0\136\0\23\0\136\0\24\0\136\0\25\0\136\0\26\0\136" +
		"\0\27\0\136\0\30\0\136\0\31\0\136\0\32\0\136\0\33\0\136\0\34\0\136\0\35\0\136\0\40" +
		"\0\136\0\52\0\136\0\61\0\136\0\63\0\136\0\66\0\136\0\67\0\136\0\uffff\uffff\ufffe" +
		"\uffff\14\0\uffff\uffff\0\0\136\0\1\0\136\0\2\0\136\0\5\0\136\0\6\0\136\0\12\0\136" +
		"\0\13\0\136\0\20\0\136\0\21\0\136\0\22\0\136\0\23\0\136\0\24\0\136\0\25\0\136\0\26" +
		"\0\136\0\27\0\136\0\30\0\136\0\31\0\136\0\32\0\136\0\33\0\136\0\34\0\136\0\35\0\136" +
		"\0\40\0\136\0\52\0\136\0\61\0\136\0\63\0\136\0\66\0\136\0\67\0\136\0\uffff\uffff" +
		"\ufffe\uffff\2\0\uffff\uffff\5\0\uffff\uffff\6\0\uffff\uffff\31\0\uffff\uffff\32" +
		"\0\uffff\uffff\33\0\uffff\uffff\40\0\uffff\uffff\66\0\uffff\uffff\67\0\uffff\uffff" +
		"\0\0\124\0\1\0\124\0\12\0\124\0\13\0\124\0\20\0\124\0\21\0\124\0\22\0\124\0\23\0" +
		"\124\0\24\0\124\0\25\0\124\0\26\0\124\0\27\0\124\0\30\0\124\0\34\0\124\0\35\0\124" +
		"\0\52\0\124\0\61\0\124\0\63\0\124\0\uffff\uffff\ufffe\uffff");

	private static final int[] tmGoto = BisonLexer.unpack_int(99,
		"\0\0\2\0\6\0\54\0\54\0\70\0\124\0\202\0\206\0\212\0\214\0\216\0\224\0\232\0\234\0" +
		"\234\0\234\0\242\0\250\0\256\0\264\0\272\0\300\0\306\0\314\0\322\0\326\0\332\0\336" +
		"\0\344\0\352\0\354\0\356\0\362\0\364\0\366\0\370\0\372\0\374\0\376\0\u0100\0\u0102" +
		"\0\u0104\0\u010a\0\u010c\0\u010e\0\u0110\0\u0112\0\u0114\0\u0116\0\u011c\0\u011e" +
		"\0\u0124\0\u0126\0\u0128\0\u013e\0\u0142\0\u0144\0\u0144\0\u0144\0\u0156\0\u0156" +
		"\0\u0156\0\u0156\0\u0156\0\u0156\0\u0156\0\u0156\0\u0156\0\u0158\0\u015a\0\u015c" +
		"\0\u015e\0\u0160\0\u0162\0\u0168\0\u016a\0\u0170\0\u0176\0\u017a\0\u017c\0\u0182" +
		"\0\u0184\0\u018a\0\u018c\0\u0190\0\u0194\0\u0198\0\u01a0\0\u01a4\0\u01a8\0\u01ac" +
		"\0\u01ae\0\u01b2\0\u01b8\0\u01ba\0\u01bc\0\u01d0\0\u01d2\0");

	private static final int[] tmFromTo = BisonLexer.unpack_int(466,
		"\213\0\214\0\2\0\56\0\57\0\56\0\4\0\63\0\5\0\63\0\15\0\72\0\17\0\74\0\41\0\113\0" +
		"\43\0\117\0\66\0\63\0\70\0\63\0\71\0\113\0\76\0\140\0\121\0\113\0\123\0\113\0\124" +
		"\0\161\0\135\0\113\0\152\0\113\0\156\0\113\0\162\0\113\0\173\0\113\0\212\0\113\0" +
		"\22\0\100\0\23\0\101\0\63\0\130\0\64\0\132\0\160\0\171\0\174\0\206\0\4\0\64\0\5\0" +
		"\64\0\41\0\114\0\66\0\64\0\70\0\64\0\71\0\114\0\121\0\114\0\123\0\114\0\135\0\114" +
		"\0\152\0\114\0\156\0\114\0\162\0\114\0\173\0\114\0\212\0\114\0\17\0\75\0\20\0\77" +
		"\0\25\0\102\0\30\0\104\0\31\0\105\0\35\0\106\0\37\0\111\0\40\0\112\0\41\0\115\0\63" +
		"\0\131\0\64\0\133\0\71\0\115\0\76\0\141\0\121\0\115\0\123\0\115\0\130\0\164\0\132" +
		"\0\165\0\135\0\115\0\152\0\115\0\156\0\115\0\162\0\115\0\173\0\115\0\212\0\115\0" +
		"\121\0\147\0\152\0\147\0\121\0\150\0\152\0\150\0\1\0\2\0\163\0\203\0\1\0\3\0\60\0" +
		"\127\0\163\0\204\0\56\0\124\0\177\0\124\0\202\0\124\0\161\0\172\0\1\0\4\0\2\0\4\0" +
		"\57\0\4\0\1\0\5\0\2\0\5\0\57\0\5\0\1\0\6\0\2\0\6\0\57\0\6\0\1\0\7\0\2\0\7\0\57\0" +
		"\7\0\1\0\10\0\2\0\10\0\57\0\10\0\1\0\11\0\2\0\11\0\57\0\11\0\1\0\12\0\2\0\12\0\57" +
		"\0\12\0\1\0\13\0\2\0\13\0\57\0\13\0\1\0\14\0\2\0\14\0\57\0\14\0\162\0\173\0\212\0" +
		"\173\0\162\0\174\0\212\0\174\0\162\0\175\0\212\0\175\0\1\0\15\0\2\0\15\0\57\0\15" +
		"\0\1\0\16\0\2\0\16\0\57\0\16\0\1\0\17\0\1\0\20\0\162\0\176\0\212\0\176\0\1\0\21\0" +
		"\1\0\22\0\1\0\23\0\1\0\24\0\1\0\25\0\1\0\26\0\1\0\27\0\1\0\30\0\1\0\31\0\1\0\32\0" +
		"\2\0\32\0\57\0\32\0\1\0\33\0\1\0\34\0\1\0\35\0\1\0\36\0\1\0\37\0\1\0\40\0\1\0\41" +
		"\0\2\0\41\0\57\0\41\0\1\0\42\0\1\0\43\0\2\0\43\0\57\0\43\0\1\0\44\0\1\0\45\0\15\0" +
		"\73\0\27\0\103\0\36\0\107\0\43\0\120\0\52\0\121\0\72\0\137\0\76\0\142\0\110\0\145" +
		"\0\117\0\146\0\162\0\177\0\212\0\177\0\162\0\200\0\212\0\200\0\1\0\46\0\4\0\65\0" +
		"\5\0\65\0\6\0\71\0\55\0\122\0\66\0\65\0\70\0\65\0\121\0\151\0\152\0\151\0\175\0\207" +
		"\0\2\0\57\0\0\0\213\0\0\0\1\0\1\0\47\0\36\0\110\0\1\0\50\0\1\0\51\0\2\0\60\0\57\0" +
		"\60\0\121\0\152\0\1\0\52\0\2\0\52\0\57\0\52\0\1\0\53\0\2\0\53\0\57\0\53\0\4\0\66" +
		"\0\5\0\70\0\71\0\135\0\1\0\54\0\2\0\54\0\57\0\54\0\123\0\156\0\1\0\55\0\2\0\55\0" +
		"\57\0\55\0\55\0\123\0\123\0\157\0\156\0\170\0\121\0\153\0\152\0\167\0\121\0\154\0" +
		"\152\0\154\0\4\0\67\0\5\0\67\0\66\0\134\0\70\0\134\0\2\0\61\0\57\0\126\0\2\0\62\0" +
		"\57\0\62\0\125\0\162\0\203\0\212\0\125\0\163\0\162\0\201\0\212\0\201\0\56\0\125\0" +
		"\177\0\210\0\202\0\211\0\17\0\76\0\76\0\143\0\41\0\116\0\71\0\136\0\121\0\155\0\123" +
		"\0\160\0\135\0\166\0\152\0\155\0\156\0\160\0\162\0\202\0\173\0\205\0\212\0\202\0" +
		"\76\0\144\0");

	private static final int[] tmRuleLen = BisonLexer.unpack_int(105,
		"\2\0\1\0\3\0\2\0\0\0\1\0\1\0\1\0\2\0\1\0\1\0\3\0\1\0\2\0\1\0\2\0\2\0\2\0\1\0\2\0" +
		"\2\0\2\0\1\0\1\0\2\0\2\0\2\0\2\0\1\0\1\0\1\0\1\0\1\0\1\0\3\0\2\0\1\0\1\0\2\0\3\0" +
		"\3\0\2\0\2\0\1\0\1\0\1\0\2\0\2\0\3\0\2\0\1\0\2\0\1\0\3\0\2\0\1\0\1\0\1\0\1\0\1\0" +
		"\1\0\0\0\2\0\1\0\1\0\1\0\1\0\1\0\1\0\1\0\3\0\2\0\2\0\1\0\3\0\2\0\2\0\1\0\1\0\2\0" +
		"\3\0\2\0\0\0\1\0\3\0\2\0\2\0\2\0\1\0\1\0\2\0\2\0\2\0\3\0\0\0\1\0\1\0\1\0\1\0\1\0" +
		"\1\0\1\0\1\0\1\0\0\0");

	private static final int[] tmRuleSymbol = BisonLexer.unpack_int(105,
		"\104\0\104\0\105\0\106\0\106\0\107\0\107\0\107\0\110\0\110\0\111\0\111\0\111\0\111" +
		"\0\111\0\111\0\111\0\111\0\111\0\111\0\111\0\111\0\111\0\111\0\111\0\111\0\111\0" +
		"\111\0\111\0\111\0\111\0\111\0\112\0\112\0\112\0\112\0\112\0\112\0\112\0\112\0\112" +
		"\0\112\0\113\0\113\0\114\0\114\0\115\0\115\0\115\0\116\0\116\0\117\0\117\0\120\0" +
		"\121\0\121\0\122\0\122\0\122\0\122\0\123\0\123\0\124\0\124\0\125\0\125\0\126\0\126" +
		"\0\126\0\127\0\127\0\127\0\127\0\127\0\127\0\127\0\127\0\127\0\130\0\130\0\131\0" +
		"\132\0\132\0\133\0\133\0\133\0\134\0\134\0\134\0\134\0\134\0\134\0\134\0\135\0\135" +
		"\0\136\0\136\0\137\0\137\0\137\0\140\0\140\0\140\0\141\0\141\0");

	protected static final String[] tmSymbolNames = new String[] {
		"eoi",
		"ID_COLON",
		"ID",
		"skip",
		"INT",
		"CHAR",
		"STRING",
		"'<*>'",
		"'<>'",
		"'%%'",
		"'|'",
		"';'",
		"'['",
		"']'",
		"skip_comment",
		"skip_ml_comment",
		"'%token'",
		"'%nterm'",
		"'%type'",
		"'%destructor'",
		"'%printer'",
		"'%left'",
		"'%right'",
		"'%nonassoc'",
		"'%precedence'",
		"'%prec'",
		"'%dprec'",
		"'%merge'",
		"'%code'",
		"'%default-prec'",
		"'%define'",
		"'%defines'",
		"'%empty'",
		"'%error-verbose'",
		"'%expect'",
		"'%expect-rr'",
		"'%<flag>'",
		"'%file-prefix'",
		"'%glr-parser'",
		"'%initial-action'",
		"'%language'",
		"'%name-prefix'",
		"'%no-default-prec'",
		"'%no-lines'",
		"'%nondeterministic-parser'",
		"'%output'",
		"'%param'",
		"'%require'",
		"'%skeleton'",
		"'%start'",
		"'%token-table'",
		"'%union'",
		"'%verbose'",
		"'%yacc'",
		"'{...}'",
		"'%?{...}'",
		"'%{...%}'",
		"tag_any",
		"tag_inc_nesting",
		"TAG",
		"code_char",
		"code_string",
		"code_comment",
		"code_ml_comment",
		"code_any",
		"code_inc_nesting",
		"code_dec_nesting",
		"code_lessless",
		"grammar_part_list",
		"input",
		"prologue_declaration_optlist",
		"prologue_declaration",
		"LBRACEDOTDOTDOTRBRACE_list",
		"prologue_directive",
		"grammar_declaration",
		"symbol_or_tag_list",
		"code_props_type",
		"symbol_declaration",
		"symbol_def_list",
		"symbol_list",
		"prec_declaration",
		"symbol_prec_list",
		"prec_directive",
		"tag_op",
		"symbol_prec",
		"symbol_or_tag",
		"tag_nt",
		"symbol_def",
		"grammar_part",
		"nonterm_rules",
		"rhsPart_optlist",
		"rules",
		"rhsPart",
		"named_ref_op",
		"variable",
		"value",
		"symbol",
		"valueopt",
	};

	public interface Nonterminals extends Tokens {
		// non-terminals
		int grammar_part_list = 68;
		int input = 69;
		int prologue_declaration_optlist = 70;
		int prologue_declaration = 71;
		int LBRACEDOTDOTDOTRBRACE_list = 72;
		int prologue_directive = 73;
		int grammar_declaration = 74;
		int symbol_or_tag_list = 75;
		int code_props_type = 76;
		int symbol_declaration = 77;
		int symbol_def_list = 78;
		int symbol_list = 79;
		int prec_declaration = 80;
		int symbol_prec_list = 81;
		int prec_directive = 82;
		int tag_op = 83;
		int symbol_prec = 84;
		int symbol_or_tag = 85;
		int tag_nt = 86;
		int symbol_def = 87;
		int grammar_part = 88;
		int nonterm_rules = 89;
		int rhsPart_optlist = 90;
		int rules = 91;
		int rhsPart = 92;
		int named_ref_op = 93;
		int variable = 94;
		int value = 95;
		int symbol = 96;
		int valueopt = 97;
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
	protected BisonLexer tmLexer;

	public Object parse(BisonLexer lexer) throws IOException, ParseException {

		tmLexer = lexer;
		tmStack = new Span[1024];
		tmHead = 0;

		tmStack[0] = new Span();
		tmStack[0].state = 0;
		tmNext = tmLexer.next();

		while (tmStack[tmHead].state != 140) {
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

		if (tmStack[tmHead].state != 140) {
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

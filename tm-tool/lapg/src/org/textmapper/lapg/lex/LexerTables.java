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
package org.textmapper.lapg.lex;

import org.textmapper.lapg.api.LexerData;

/**
 * Representation of generated lexer tables.
 */
class LexerTables implements LexerData {

	private final int nchars;
	private final int[] char2no, groupset, change, backtracking;

	LexerTables(int nchars, int[] char2no, int[] groupset, int[] change, int[] backtracking) {
		this.nchars = nchars;
		this.char2no = char2no;
		this.groupset = groupset;
		this.change = change;
		this.backtracking = backtracking;
	}

	@Override
	public int getNchars() {
		return nchars;
	}

	@Override
	public int[] getChar2no() {
		return char2no;
	}

	@Override
	public int[] getGroupset() {
		return groupset;
	}

	@Override
	public int[] getChange() {
		return change;
	}

	@Override
	public int[] getBacktracking() {
		return backtracking;
	}
}

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
package org.textmapper.lapg.lalr;

import org.textmapper.lapg.api.LookaheadRule;
import org.textmapper.lapg.api.Marker;
import org.textmapper.lapg.api.ParserData;
import org.textmapper.lapg.api.Symbol;

/**
 * Representation of generated parser tables.
 */
class ParserTables implements ParserData {

	private Symbol[] sym;
	private int rules, nsyms, nterms, nstates;
	private int[] rleft, rlen;
	private int[] sym_goto, sym_fromto, action_table;
	private int[] action_index;
	private int[] final_states;
	private Marker[] markers;
	private LookaheadRule[] lookaheadRules;

	ParserTables(Symbol[] sym,
				 int rules, int nsyms, int nterms, int nstates,
				 int[] rleft, int[] rlen,
				 int[] sym_goto, int[] sym_fromto,
				 int[] action_table, int[] action_index, int[] final_states,
				 Marker[] markers, LookaheadRule[] lookaheadRules) {
		this.sym = sym;
		this.rules = rules;
		this.nsyms = nsyms;
		this.nterms = nterms;
		this.nstates = nstates;
		this.rleft = rleft;
		this.rlen = rlen;
		this.sym_goto = sym_goto;
		this.sym_fromto = sym_fromto;
		this.action_table = action_table;
		this.action_index = action_index;
		this.final_states = final_states;
		this.markers = markers;
		this.lookaheadRules = lookaheadRules;
	}

	@Override
	public int[] getRuleLength() {
		return rlen;
	}

	@Override
	public Symbol[] getSymbols() {
		return sym;
	}

	@Override
	public int getRules() {
		return rules;
	}

	@Override
	public int getNsyms() {
		return nsyms;
	}

	@Override
	public int getNterms() {
		return nterms;
	}

	@Override
	public int getStatesCount() {
		return nstates;
	}

	@Override
	public int[] getSymGoto() {
		return sym_goto;
	}

	@Override
	public int[] getSymFromTo() {
		return sym_fromto;
	}

	@Override
	public int[] getLalr() {
		return action_table;
	}

	@Override
	public int[] getAction() {
		return action_index;
	}

	@Override
	public int[] getFinalStates() {
		return final_states;
	}

	@Override
	public int[] getLeft() {
		return rleft;
	}

	@Override
	public Marker[] getMarkers() {
		return markers;
	}

	@Override
	public LookaheadRule[] getLookaheadRules() {
		return lookaheadRules;
	}

	private static int byteSize(int maxInt) {
		return maxInt < Short.MAX_VALUE ? 2 : 4;
	}

	@Override
	public int getByteSize() {
		int result = 0;
		if (sym_goto.length > 0) {
			int max = sym_goto[sym_goto.length - 1];
			result += sym_goto.length * byteSize(max);
			result += sym_fromto.length * byteSize(nstates);
		}
		result += (action_index.length + action_table.length) * 4;
		result += (rlen.length + rleft.length) * byteSize(nsyms);
		return result;
	}
}

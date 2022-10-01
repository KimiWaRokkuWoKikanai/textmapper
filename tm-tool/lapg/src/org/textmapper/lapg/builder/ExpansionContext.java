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
package org.textmapper.lapg.builder;

import org.textmapper.lapg.api.Terminal;
import org.textmapper.lapg.api.rule.RhsSet;

import java.util.HashMap;
import java.util.Map;

final class ExpansionContext {

	private Map<RhsSet, Terminal[]> resolvedSets = new HashMap<>();

	ExpansionContext() {
	}

	Terminal[] resolveSet(RhsSet set) {
		return resolvedSets.get(set);
	}

	void addSet(RhsSet set, Terminal[] value) {
		resolvedSets.put(set, value);
	}
}

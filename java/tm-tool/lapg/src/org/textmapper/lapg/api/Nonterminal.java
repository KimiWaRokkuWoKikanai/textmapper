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
package org.textmapper.lapg.api;

import org.textmapper.lapg.api.rule.RhsRoot;

/**
 * evgeny, 10/27/12
 */
public interface Nonterminal extends Symbol {
	/**
	 * User data key for template parameters (contains {@code List<TemplateParameter>})
	 */
	String UD_TEMPLATE_PARAMS = "templateParameters";

	RhsRoot getDefinition();

	Iterable<Rule> getRules();

	/**
	 * @return true for nonterminals that are templates and cannot be referenced without parameters.
	 */
	boolean isTemplate();

	Nonterminal getTemplate();

	/**
	 * @return true if it can derive an empty string
	 */
	boolean isNullable();
}

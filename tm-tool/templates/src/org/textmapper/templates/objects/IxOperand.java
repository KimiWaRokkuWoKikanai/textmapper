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
package org.textmapper.templates.objects;

import org.textmapper.templates.api.EvaluationException;

public interface IxOperand {

	Object plus(Object v) throws EvaluationException;
	Object minus(Object v) throws EvaluationException;

	Object multiply(Object v) throws EvaluationException;
	Object div(Object v) throws EvaluationException;
	Object mod(Object v) throws EvaluationException;

	int compareTo(Object v) throws EvaluationException;
	boolean equalsTo(Object v) throws EvaluationException;
}

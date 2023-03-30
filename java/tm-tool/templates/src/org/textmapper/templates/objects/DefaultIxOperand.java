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

public class DefaultIxOperand implements IxOperand, IxWrapper {

	private final Object myObject;

	public DefaultIxOperand(Object object) {
		this.myObject = object;
	}

	@Override
	public Object plus(Object v) throws EvaluationException {
		throw new EvaluationException("'+' is not supported for `" + getType() + "`");
	}

	@Override
	public Object minus(Object v) throws EvaluationException {
		throw new EvaluationException("'-' is not supported for `" + getType() + "`");
	}

	@Override
	public Object multiply(Object v) throws EvaluationException {
		throw new EvaluationException("'*' is not supported for `" + getType() + "`");
	}

	@Override
	public Object div(Object v) throws EvaluationException {
		throw new EvaluationException("'/' is not supported for `" + getType() + "`");
	}

	@Override
	public Object mod(Object v) throws EvaluationException {
		throw new EvaluationException("'%' is not supported for `" + getType() + "`");
	}

	@Override
	public int compareTo(Object v) throws EvaluationException {
		throw new EvaluationException("compare is not supported for `" + getType() + "`");
	}

	@Override
	public boolean equalsTo(Object v) throws EvaluationException {
		Object real = v instanceof IxWrapper ? ((IxWrapper) v).getObject() : v;
		return myObject == null ? real == null : myObject.equals(real);
	}

	protected String getType() {
		return "Object";
	}

	@Override
	public Object getObject() {
		return myObject;
	}
}

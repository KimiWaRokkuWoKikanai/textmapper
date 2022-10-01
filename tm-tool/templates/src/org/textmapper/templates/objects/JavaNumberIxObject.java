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

import java.util.Iterator;

public class JavaNumberIxObject implements IxWrapper, IxOperand, IxAdaptable {

	private final Number myNumber;

	public JavaNumberIxObject(Number myNumber) {
		this.myNumber = myNumber;
		if (!isLong(myNumber)) {
			throw new IllegalArgumentException("unsupported number: " + myNumber.toString());
		}
	}

	@Override
	public Object plus(Object v) throws EvaluationException {
		Number op = asNumber(v);
		if (isInt(myNumber) && isInt(op)) {
			return myNumber.intValue() + op.intValue();
		}
		if (isLong(myNumber) && isLong(op)) {
			return myNumber.longValue() + op.longValue();
		}
		throw new EvaluationException("unsupported operand of type `" + op.getClass() + "`");
	}

	@Override
	public Object minus(Object v) throws EvaluationException {
		Number op = asNumber(v);
		if (isInt(myNumber) && isInt(op)) {
			return myNumber.intValue() - op.intValue();
		}
		if (isLong(myNumber) && isLong(op)) {
			return myNumber.longValue() - op.longValue();
		}
		throw new EvaluationException("unsupported operand of type `" + op.getClass() + "`");
	}

	@Override
	public Object multiply(Object v) throws EvaluationException {
		Number op = asNumber(v);
		if (isInt(myNumber) && isInt(op)) {
			return myNumber.intValue() * op.intValue();
		}
		if (isLong(myNumber) && isLong(op)) {
			return myNumber.longValue() * op.longValue();
		}
		throw new EvaluationException("unsupported operand of type `" + op.getClass() + "`");
	}

	@Override
	public Object div(Object v) throws EvaluationException {
		Number op = asNumber(v);
		if (isInt(myNumber) && isInt(op)) {
			return myNumber.intValue() / op.intValue();
		}
		if (isLong(myNumber) && isLong(op)) {
			return myNumber.longValue() / op.longValue();
		}
		throw new EvaluationException("unsupported operand of type `" + op.getClass() + "`");
	}

	@Override
	public Object mod(Object v) throws EvaluationException {
		Number op = asNumber(v);
		if (isInt(myNumber) && isInt(op)) {
			return myNumber.intValue() % op.intValue();
		}
		if (isLong(myNumber) && isLong(op)) {
			return myNumber.longValue() % op.longValue();
		}
		throw new EvaluationException("unsupported operand of type `" + op.getClass() + "`");
	}

	@Override
	public int compareTo(Object v) throws EvaluationException {
		Number op = asNumber(v);
		if (isInt(myNumber) && isInt(op)) {
			return new Integer(myNumber.intValue()).compareTo(op.intValue());
		}
		if (isLong(myNumber) && isLong(op)) {
			return new Long(myNumber.longValue()).compareTo(op.longValue());
		}
		throw new EvaluationException("unsupported operand of type `" + op.getClass() + "`");
	}

	@Override
	public boolean equalsTo(Object v) throws EvaluationException {
		return compareTo(v) == 0;
	}

	private Number asNumber(Object v) throws EvaluationException {
		Object real = v instanceof IxWrapper ? ((IxWrapper) v).getObject() : v;
		if (real == null) {
			throw new EvaluationException("number is null");
		}
		if (real instanceof Number && isLong((Number) real)) {
			return (Number) real;
		}
		if (real instanceof String) {
			try {
				if (((String) real).length() < 10) {
					return Integer.parseInt((String) real);
				} else {
					return Long.parseLong((String) real);
				}
			} catch (NumberFormatException ex) {
				/* ignore */
			}
		}
		throw new EvaluationException("unsupported operand of type `" + real.getClass() + "`");
	}

	private static boolean isInt(Number object) {
		return object instanceof Byte || object instanceof Short || object instanceof Integer;
	}

	private static boolean isLong(Number object) {
		return object instanceof Byte || object instanceof Short || object instanceof Integer || object instanceof Long;
	}

	private static boolean isFloat(Object object) {
		return object instanceof Double || object instanceof Float;
	}

	@Override
	public Object getObject() {
		return myNumber;
	}

	@Override
	public Object castTo(String qualifiedName) throws EvaluationException {
		throw new EvaluationException("cannot cast a number");
	}

	@Override
	public String asString() throws EvaluationException {
		return myNumber.toString();
	}

	@Override
	public boolean asBoolean() {
		return myNumber.longValue() != 0;
	}

	@Override
	public Iterator asSequence() throws EvaluationException {
		throw new EvaluationException("cannot iterate over a number");
	}
}

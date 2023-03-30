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

import org.textmapper.lapg.api.DerivedSourceElement;
import org.textmapper.lapg.api.SourceElement;
import org.textmapper.lapg.api.ast.AstRawType;
import org.textmapper.lapg.api.ast.AstType;

class LiRawAstType implements AstRawType, DerivedSourceElement {

	private final String type;
	private final SourceElement origin;

	public LiRawAstType(String type, SourceElement origin) {
		if (type == null) {
			throw new NullPointerException("type");
		}
		this.type = type;
		this.origin = origin;
	}

	@Override
	public String getRawType() {
		return type;
	}

	@Override
	public SourceElement getOrigin() {
		return origin;
	}

	@Override
	public boolean isSubtypeOf(AstType another) {
		return equals(another) || another == AstType.ANY;
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;

		LiRawAstType that = (LiRawAstType) o;
		return type.equals(that.type);
	}

	@Override
	public int hashCode() {
		return type.hashCode();
	}

	@Override
	public String toString() {
		return type;
	}
}

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

import org.textmapper.lapg.api.SourceElement;
import org.textmapper.lapg.api.Terminal;
import org.textmapper.lapg.api.ast.AstType;
import org.textmapper.lapg.api.rule.*;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * evgeny, 12/5/12
 */
class LiRhsSequence extends LiRhsPart implements RhsSequence {

	private static final List<RhsCFPart[]> ONE = Collections.singletonList(RhsSymbol.EMPTY_LIST);

	private final String name;
	private final LiRhsPart[] parts;
	private AstType type;
	private RhsMapping mapping;
	private Terminal precedence;

	LiRhsSequence(String name, LiRhsPart[] parts, boolean isRewrite, SourceElement origin) {
		super(origin);
		this.name = name;
		this.parts = parts;
		register(isRewrite, parts);
	}

	@Override
	public String getName() {
		return name;
	}

	@Override
	public RhsPart[] getParts() {
		return parts;
	}

	@Override
	public AstType getType() {
		return type;
	}

	@Override
	public RhsMapping getMapping() {
		return mapping;
	}

	void map(AstType type, RhsMapping mapping) {
		this.type = type;
		this.mapping = mapping;
	}

	@Override
	List<RhsCFPart[]> expand(ExpansionContext context) {
		return expandList(parts, context);
	}

	@Override
	public boolean structurallyEquals(LiRhsPart o) {
		if (this == o) return true;
		if (o == null || getClass() != o.getClass()) return false;
		LiRhsSequence that = (LiRhsSequence) o;
		return structurallyEquals(parts, that.parts);
	}

	@Override
	public int structuralHashCode() {
		return structuralHashCode(parts);
	}

	static List<RhsCFPart[]> expandList(LiRhsPart[] list, ExpansionContext context) {
		boolean simplePartsOnly = true;
		for (RhsPart part : list) {
			if (!(part instanceof RhsSymbol)) {
				simplePartsOnly = false;
				break;
			}
		}
		if (simplePartsOnly) {
			RhsSymbol[] parts = new RhsSymbol[list.length];
			System.arraycopy(list, 0, parts, 0, parts.length);
			return Collections.singletonList(parts);

		} else {
			List<RhsCFPart[]> result = ONE;
			for (LiRhsPart part : list) {
				List<RhsCFPart[]> val = part.expand(context);
				result = cartesianProduct(result, val);
			}
			return result;
		}
	}

	private static List<RhsCFPart[]> cartesianProduct(List<RhsCFPart[]> left, List<RhsCFPart[]> right) {
		if (left == ONE) {
			return right;
		}
		List<RhsCFPart[]> result = new ArrayList<>(left.size() * right.size());
		for (RhsCFPart[] leftElement : left) {
			for (RhsCFPart[] rightElement : right) {
				RhsCFPart[] elem = new RhsCFPart[leftElement.length + rightElement.length];
				System.arraycopy(leftElement, 0, elem, 0, leftElement.length);
				System.arraycopy(rightElement, 0, elem, leftElement.length, rightElement.length);
				result.add(elem);
			}
		}
		return result;
	}

	@Override
	public Kind getKind() {
		return Kind.Sequence;
	}

	@Override
	protected void toString(StringBuilder sb) {
		if (parts.length == 1) {
			parts[0].toString(sb);
		} else {
			sb.append("(");
			toString(sb, parts, " ");
			sb.append(")");
		}
	}

	Terminal getPrecedence() {
		return precedence;
	}

	void setPrecedence(Terminal precedence) {
		// TODO check unused priorities
		this.precedence = precedence;
	}
}

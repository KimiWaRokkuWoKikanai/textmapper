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
package org.textmapper.lapg.eval;


import org.textmapper.lapg.eval.GenericParseContext.TextSource;

/**
 * Gryaznov Evgeny, 3/28/11
 */
public class GenericNode {

	private final TextSource source;
	private final int offset;
	private final int endoffset;
	private final GenericNode[] children;

	public GenericNode(TextSource source, int offset, int endoffset, GenericNode... children) {
		this.source = source;
		this.offset = offset;
		this.endoffset = endoffset;
		this.children = children != null && children.length > 0 ? children : null;
	}

	public String getLocation() {
		return source.getLocation(offset);
	}

	public int getLine() {
		return source.lineForOffset(offset);
	}

	public int getOffset() {
		return offset;
	}

	public int getEndOffset() {
		return endoffset;
	}

	public TextSource getInput() {
		return source;
	}

	@Override
	public String toString() {
		return source.getText(offset, endoffset);
	}

	public String toSignature() {
		StringBuilder sb = new StringBuilder();
		toString(sb);
		return sb.toString();
	}

	public void toString(StringBuilder sb) {
		sb.append('[');
		int offset = this.offset;
		if (children != null) {
			for (GenericNode node : children) {
				if (offset < node.offset) {
					sb.append(source.getText(offset, node.offset));
				}
				node.toString(sb);
				offset = node.endoffset;
			}
		}
		if (offset < endoffset) {
			sb.append(source.getText(offset, endoffset));
		}
		sb.append(']');
	}
}

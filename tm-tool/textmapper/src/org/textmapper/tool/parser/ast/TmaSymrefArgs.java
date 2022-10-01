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
package org.textmapper.tool.parser.ast;

import java.util.List;
import org.textmapper.tool.parser.TMTree.TextSource;

public class TmaSymrefArgs extends TmaNode {

	private final List<TmaArgument> argList;

	public TmaSymrefArgs(List<TmaArgument> argList, TextSource source, int line, int offset, int endoffset) {
		super(source, line, offset, endoffset);
		this.argList = argList;
	}

	public List<TmaArgument> getArgList() {
		return argList;
	}

	@Override
	public void accept(TmaVisitor v) {
		if (!v.visit(this)) {
			return;
		}
		if (argList != null) {
			for (TmaArgument it : argList) {
				it.accept(v);
			}
		}
	}
}

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
package org.textmapper.templates.types.ast;

import java.util.List;
import org.textmapper.templates.types.TypesTree.TextSource;

public class AstTypeDeclaration extends AstNode {

	private final String name;
	private final List<List<String>> _super;
	private final List<IAstMemberDeclaration> members;

	public AstTypeDeclaration(String name, List<List<String>> _super, List<IAstMemberDeclaration> members, TextSource source, int line, int offset, int endoffset) {
		super(source, line, offset, endoffset);
		this.name = name;
		this._super = _super;
		this.members = members;
	}

	public String getName() {
		return name;
	}

	public List<List<String>> getSuper() {
		return _super;
	}

	public List<IAstMemberDeclaration> getMembers() {
		return members;
	}

	@Override
	public void accept(AstVisitor v) {
		if (!v.visit(this)) {
			return;
		}
		if (members != null) {
			for (IAstMemberDeclaration it : members) {
				it.accept(v);
			}
		}
	}
}

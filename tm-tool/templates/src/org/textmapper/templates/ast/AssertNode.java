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
package org.textmapper.templates.ast;

import org.textmapper.templates.api.EvaluationContext;
import org.textmapper.templates.api.EvaluationException;
import org.textmapper.templates.api.IEvaluationStrategy;
import org.textmapper.templates.api.TemplatesStatus;
import org.textmapper.templates.ast.TemplatesTree.TextSource;

public class AssertNode extends Node {

	private final ExpressionNode expr;

	public AssertNode(ExpressionNode expr, TextSource source, int offset, int endoffset) {
		super(source, offset, endoffset);
		this.expr = expr;
	}

	@Override
	protected void emit(StringBuilder sb, EvaluationContext context, IEvaluationStrategy env) {
		try {
			if (!env.asAdaptable(env.evaluate(expr, context, true)).asBoolean()) {
				env.report(TemplatesStatus.KIND_ERROR, "Assertion `" + expr.toString()
						+ "` failed for " + env.getTitle(context.getThisObject()), this);
			}
		} catch (EvaluationException ex) {
			/* already handled, ignore */
		}
	}

	@Override
	public void toJavascript(StringBuilder sb) {
		sb.append("/* TODO assert(");
		expr.toJavascript(sb);
		sb.append(") */");
	}
}

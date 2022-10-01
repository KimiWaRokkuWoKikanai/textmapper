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

public abstract class TmaVisitor {

	protected boolean visit(TmaNontermTypeAST n) {
		return true;
	}

	protected boolean visit(TmaNontermTypeHint n) {
		return true;
	}

	protected boolean visit(TmaSetBinary n) {
		return true;
	}

	protected boolean visit(TmaInlineParameter n) {
		return true;
	}

	protected boolean visit(TmaPredicateBinary n) {
		return true;
	}

	protected boolean visit(TmaArray n) {
		return true;
	}

	protected boolean visit(TmaIdentifier n) {
		return true;
	}

	protected boolean visit(TmaLiteral n) {
		return true;
	}

	protected boolean visit(TmaPattern n) {
		return true;
	}

	protected boolean visit(TmaName n) {
		return true;
	}

	protected boolean visit(TmaCommand n) {
		return true;
	}

	protected boolean visit(TmaSyntaxProblem n) {
		return true;
	}

	protected boolean visit(TmaInput1 n) {
		return true;
	}

	protected boolean visit(TmaHeader n) {
		return true;
	}

	protected boolean visit(TmaImport n) {
		return true;
	}

	protected boolean visit(TmaOption n) {
		return true;
	}

	protected boolean visit(TmaSymref n) {
		return true;
	}

	protected boolean visit(TmaRawType n) {
		return true;
	}

	protected boolean visit(TmaNamedPattern n) {
		return true;
	}

	protected boolean visit(TmaStartConditionsScope n) {
		return true;
	}

	protected boolean visit(TmaStartConditions n) {
		return true;
	}

	protected boolean visit(TmaLexeme n) {
		return true;
	}

	protected boolean visit(TmaLexemeAttrs n) {
		return true;
	}

	protected boolean visit(TmaStateref n) {
		return true;
	}

	protected boolean visit(TmaLexerState n) {
		return true;
	}

	protected boolean visit(TmaNonterm n) {
		return true;
	}

	protected boolean visit(TmaInputref n) {
		return true;
	}

	protected boolean visit(TmaRule0 n) {
		return true;
	}

	protected boolean visit(TmaRhsSuffix n) {
		return true;
	}

	protected boolean visit(TmaReportClause n) {
		return true;
	}

	protected boolean visit(TmaReportAs n) {
		return true;
	}

	protected boolean visit(TmaRhsLookahead n) {
		return true;
	}

	protected boolean visit(TmaLookaheadPredicate n) {
		return true;
	}

	protected boolean visit(TmaRhsStateMarker n) {
		return true;
	}

	protected boolean visit(TmaAnnotations n) {
		return true;
	}

	protected boolean visit(TmaAnnotation n) {
		return true;
	}

	protected boolean visit(TmaNontermParams n) {
		return true;
	}

	protected boolean visit(TmaParamRef n) {
		return true;
	}

	protected boolean visit(TmaSymrefArgs n) {
		return true;
	}

	protected boolean visit(TmaArgument n) {
		return true;
	}

	protected boolean visit(TmaBracketsDirective n) {
		return true;
	}

	protected boolean visit(TmaStatesClause n) {
		return true;
	}

	protected boolean visit(TmaTemplateParam n) {
		return true;
	}

	protected boolean visit(TmaDirectivePrio n) {
		return true;
	}

	protected boolean visit(TmaDirectiveInput n) {
		return true;
	}

	protected boolean visit(TmaDirectiveInterface n) {
		return true;
	}

	protected boolean visit(TmaDirectiveAssert n) {
		return true;
	}

	protected boolean visit(TmaDirectiveSet n) {
		return true;
	}

	protected boolean visit(TmaDirectiveExpect n) {
		return true;
	}

	protected boolean visit(TmaDirectiveExpectRR n) {
		return true;
	}

	protected boolean visit(TmaRhsAnnotated n) {
		return true;
	}

	protected boolean visit(TmaRhsAssignment n) {
		return true;
	}

	protected boolean visit(TmaRhsQuantifier n) {
		return true;
	}

	protected boolean visit(TmaRhsCast n) {
		return true;
	}

	protected boolean visit(TmaRhsAsLiteral n) {
		return true;
	}

	protected boolean visit(TmaRhsSymbol n) {
		return true;
	}

	protected boolean visit(TmaRhsNested n) {
		return true;
	}

	protected boolean visit(TmaRhsList n) {
		return true;
	}

	protected boolean visit(TmaRhsIgnored n) {
		return true;
	}

	protected boolean visit(TmaRhsSet n) {
		return true;
	}

	protected boolean visit(TmaSetSymbol n) {
		return true;
	}

	protected boolean visit(TmaSetCompound n) {
		return true;
	}

	protected boolean visit(TmaSetComplement n) {
		return true;
	}

	protected boolean visit(TmaBoolPredicate n) {
		return true;
	}

	protected boolean visit(TmaComparePredicate n) {
		return true;
	}
}

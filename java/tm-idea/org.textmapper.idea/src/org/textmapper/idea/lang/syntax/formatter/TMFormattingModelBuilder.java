/**
 * Copyright 2010-2017 Evgeny Gryaznov
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see http://www.gnu.org/licenses/.
 */
package org.textmapper.idea.lang.syntax.formatter;

import com.intellij.formatting.*;
import com.intellij.lang.ASTNode;
import com.intellij.openapi.util.TextRange;
import com.intellij.psi.PsiElement;
import com.intellij.psi.PsiFile;
import com.intellij.psi.codeStyle.CodeStyleSettings;
import org.jetbrains.annotations.NotNull;
import org.textmapper.idea.lang.syntax.TMFileType;

import static org.textmapper.idea.lang.syntax.lexer.TMTokenTypes.*;
import static org.textmapper.idea.lang.syntax.parser.TextmapperElementTypes.*;

/**
 * evgeny, 8/14/12
 */
public class TMFormattingModelBuilder implements FormattingModelBuilder {
	@NotNull
	@Override
	public FormattingModel createModel(PsiElement element, CodeStyleSettings settings) {
		final TMFormattingBlock formattingBlock = new TMFormattingBlock(
				element.getNode(), null, Indent.getNoneIndent(), null, settings,
				createSpacingBuilder(settings));

		return FormattingModelProvider.createFormattingModelForPsiFile(
				element.getContainingFile(), formattingBlock, settings);
	}

	private static SpacingBuilder createSpacingBuilder(CodeStyleSettings settings) {
		return new SpacingBuilder(settings, TMFileType.TM_LANGUAGE)
				.between(COMMENT, OPTION).blankLines(1)

				.between(OPTION, OPTION).lineBreakInCode()
				.after(OPTION).blankLines(1)

				/* lexer */
				.before(START_CONDITIONS_SCOPE).blankLines(1)
				.after(START_CONDITIONS).spaces(1)
				.before(LEXEM).lineBreakInCode()
				.before(NAMED_PATTERN).lineBreakInCode()
				.beforeInside(OP_EQ, NAMED_PATTERN).spaces(1)
				.afterInside(OP_EQ, NAMED_PATTERN).spacing(1, 100, 0, true, 0)

				/* nonterm */
				.between(NONTERM, NONTERM).blankLines(1)
				.before(NONTERM).lineBreakInCode()
				.afterInside(RULE, NONTERM).lineBreakInCode()
				.after(OP_OR).spaces(1)
				.beforeInside(OP_SEMICOLON, NONTERM).lineBreakInCode()
				;
	}

	@Override
	public TextRange getRangeAffectingIndent(PsiFile file, int offset, ASTNode elementAtOffset) {
		return null;
	}
}

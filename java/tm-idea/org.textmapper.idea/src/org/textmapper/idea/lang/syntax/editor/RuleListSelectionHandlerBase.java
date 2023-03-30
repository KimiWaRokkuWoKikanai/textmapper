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
package org.textmapper.idea.lang.syntax.editor;

import com.intellij.codeInsight.editorActions.ExtendWordSelectionHandlerBase;
import com.intellij.openapi.editor.Editor;
import com.intellij.openapi.util.TextRange;
import com.intellij.openapi.util.text.LineTokenizer;
import com.intellij.psi.PsiElement;
import com.intellij.psi.PsiWhiteSpace;
import org.textmapper.idea.lang.syntax.lexer.TMTokenTypes;
import org.textmapper.idea.lang.syntax.lexer.TmToken;
import org.textmapper.idea.lang.syntax.psi.TmNonterm;
import org.textmapper.idea.lang.syntax.psi.TmRule;

import java.util.ArrayList;
import java.util.List;

/**
 * Gryaznov Evgeny, 9/21/12
 */
public class RuleListSelectionHandlerBase extends ExtendWordSelectionHandlerBase {

	@Override
	public boolean canSelect(PsiElement e) {
		return e instanceof TmRule ||
				e instanceof TmToken && ((TmToken) e).getTokenType() == TMTokenTypes.OP_OR;
	}

	@Override
	public List<TextRange> select(PsiElement e, CharSequence editorText, int cursorOffset, Editor editor) {
		List<TextRange> result = new ArrayList<>();

		PsiElement parent = e.getParent();

		if (!(parent instanceof TmNonterm)) {
			return result;
		}

		PsiElement startElement = e;
		PsiElement endElement = e;
		while (startElement.getPrevSibling() != null) {
			PsiElement sibling = startElement.getPrevSibling();

			if (sibling instanceof TmToken) {
				TmToken token = (TmToken) sibling;
				if (token.getNode().getElementType() == TMTokenTypes.OP_COLON) {
					break;
				}
			}

			if (sibling instanceof PsiWhiteSpace) {
				PsiWhiteSpace whiteSpace = (PsiWhiteSpace) sibling;

				String[] strings = LineTokenizer.tokenize(whiteSpace.getText().toCharArray(), false);
				if (strings.length > 2) {
					break;
				}
			}

			startElement = sibling;
		}

		while (startElement instanceof PsiWhiteSpace) {
			startElement = startElement.getNextSibling();
		}

		while (endElement.getNextSibling() != null) {
			PsiElement sibling = endElement.getNextSibling();

			if (sibling instanceof TmToken) {
				TmToken token = (TmToken) sibling;
				if (token.getTokenType() == TMTokenTypes.OP_SEMICOLON) {
					break;
				}
			}

			if (sibling instanceof PsiWhiteSpace) {
				PsiWhiteSpace whiteSpace = (PsiWhiteSpace) sibling;

				String[] strings = LineTokenizer.tokenize(whiteSpace.getText().toCharArray(), false);
				if (strings.length > 2) {
					break;
				}
			}

			endElement = sibling;
		}

		while (endElement instanceof PsiWhiteSpace) {
			endElement = endElement.getPrevSibling();
		}

		result.addAll(expandToWholeLine(editorText, new TextRange(startElement.getTextRange().getStartOffset(),
				endElement.getTextRange().getEndOffset())));

		return result;

	}
}

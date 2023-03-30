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
package org.textmapper.idea.lang.syntax;

import com.intellij.lang.Language;
import com.intellij.openapi.fileTypes.SingleLazyInstanceSyntaxHighlighterFactory;
import com.intellij.openapi.fileTypes.SyntaxHighlighter;
import com.intellij.openapi.fileTypes.SyntaxHighlighterFactory;
import org.jetbrains.annotations.NotNull;

public class TextmapperLanguage extends Language {

	public static final String ID = "Textmapper";

	public TextmapperLanguage() {
		super(ID);

		SyntaxHighlighterFactory.LANGUAGE_FACTORY.addExplicitExtension(this, new LapgHighlighterFactory());
	}

	private static class LapgHighlighterFactory extends SingleLazyInstanceSyntaxHighlighterFactory {
		@NotNull
		protected SyntaxHighlighter createHighlighter() {
			return new TMSyntaxHighlighter();
		}
	}
}

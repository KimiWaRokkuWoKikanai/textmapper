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
package org.textmapper.templates.storage;

import java.net.URI;

public class Resource {

	private final URI uri;
	private final String contents;
	private final int line;
	private final int offset;

	public Resource(URI uri, String contents) {
		this(uri, contents, 1, 0);
	}

	public Resource(URI uri, String contents, int line, int offset) {
		this.uri = uri;
		this.contents = contents;
		this.line = line;
		this.offset = offset;
	}

	public URI getUri() {
		return uri;
	}

	public String getContents() {
		return contents;
	}

	public int getInitialLine() {
		return line;
	}

	public int getInitialOffset() {
		return offset;
	}
}

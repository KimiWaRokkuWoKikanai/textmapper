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
package org.textmapper.tool.compiler;

import java.util.List;

public class RangeType {

	private String name;
	private List<String> flags;
	private String iface;

	public RangeType(String name, List<String> flags, String iface) {
		this.name = name;
		this.flags = flags;
		this.iface = iface;
	}

	public String getName() {
		return name;
	}

	public List<String> getFlags() {
		return flags;
	}

	public String getIface() {
		return iface;
	}
}

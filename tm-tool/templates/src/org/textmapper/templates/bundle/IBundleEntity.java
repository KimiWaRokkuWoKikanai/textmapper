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
package org.textmapper.templates.bundle;

import org.textmapper.templates.api.SourceElement;

public interface IBundleEntity extends SourceElement {

	int KIND_ANY = 0;
	int KIND_TEMPLATE = 1;
	int KIND_QUERY = 2;

	/**
	 * @return KIND_TEMPLATE or KIND_QUERY
	 */
	int getKind();

	/**
	 * @return member name
	 */
	String getName();

	/**
	 * @return qualified name of package
	 */
	String getPackage();

	/**
	 * @return signature to map overrides
	 */
	String getSignature();

	/**
	 * Returns overridden member
	 */
	IBundleEntity getBase();

	/**
	 * Internal: Used for binding.
	 */
	void setBase(IBundleEntity base);
}

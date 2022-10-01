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
package org.textmapper.lapg.util;

import java.util.Iterator;

public class ArrayIterable<T> implements Iterable<T> {
	private final T[] array;
	private final boolean reversed;

	public ArrayIterable(T[] array, boolean reversed) {
		this.array = array;
		this.reversed = reversed;
	}

	@Override
	public Iterator<T> iterator() {
		return new Iterator<T>() {
			int index = reversed ? array.length - 1 : 0;

			@Override
			public boolean hasNext() {
				return reversed ? index >= 0 : index < array.length;
			}

			@Override
			public T next() {
				return array[reversed ? index-- : index++];
			}

			@Override
			public void remove() {
				throw new UnsupportedOperationException();
			}
		};
	}

}

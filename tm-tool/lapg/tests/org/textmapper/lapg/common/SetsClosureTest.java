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
package org.textmapper.lapg.common;

import org.junit.Test;
import org.textmapper.lapg.common.SetsClosure;

import java.util.BitSet;

import static org.junit.Assert.*;

public class SetsClosureTest {

	@Test
	public void empty() throws Exception {
		SetsClosure closure = new SetsClosure();
		assertTrue(closure.compute());
	}

	@Test
	public void singleSet() throws Exception {
		SetsClosure closure = new SetsClosure();
		closure.addSet(new int[]{42, 84}, 1);
		closure.addDependencies(0, 0);
		assertTrue(closure.compute());

		assertArrayEquals(new int[]{42, 84}, closure.getSet(0));
	}

	@Test
	public void simpleUnion() throws Exception {
		SetsClosure closure = new SetsClosure();

		closure.addSet(new int[]{1, 2, 3}, 1);
		closure.addSet(new int[]{7, 8}, 1);
		closure.addSet(new int[]{13}, 1);

		closure.addDependencies(0, 1, 2);
		closure.addDependencies(2, 1);

		assertTrue(closure.compute());

		assertArrayEquals(new int[]{1, 2, 3, 7, 8, 13}, closure.getSet(0));
		assertArrayEquals(new int[]{7, 8}, closure.getSet(1));
		assertArrayEquals(new int[]{7, 8, 13}, closure.getSet(2));
	}

	@Test
	public void unionWithCycle() throws Exception {
		SetsClosure closure = new SetsClosure();

		closure.addSet(new int[]{78}, 1);
		closure.addSet(new int[]{1, 11, 12}, 1);  // 1: depends on 0 and 2
		closure.addSet(new int[]{7, 8}, 1);     // 2: depends on 3
		closure.addSet(new int[]{4, 5, 27}, 1);  // 3: depends on 1
		closure.addSet(new int[]{89}, 1);      // 4: depends on 3

		closure.addDependencies(1, 0, 2);
		closure.addDependencies(2, 3);
		closure.addDependencies(3, 1);
		closure.addDependencies(4, 3);

		assertTrue(closure.compute());

		assertArrayEquals(new int[]{78}, closure.getSet(0));
		assertArrayEquals(new int[]{1, 4, 5, 7, 8, 11, 12, 27, 78}, closure.getSet(1));
		assertArrayEquals(new int[]{1, 4, 5, 7, 8, 11, 12, 27, 78}, closure.getSet(2));
		assertArrayEquals(new int[]{1, 4, 5, 7, 8, 11, 12, 27, 78}, closure.getSet(3));
		assertArrayEquals(new int[]{1, 4, 5, 7, 8, 11, 12, 27, 78, 89}, closure.getSet(4));
	}

	@Test
	public void intersection() throws Exception {
		SetsClosure closure = new SetsClosure();

		closure.addSet(new int[]{1, 2, 3, 4, 5, 6, 9}, 1);
		closure.addSet(new int[]{0, 3, 5, 9, 12}, 1);

		closure.addSet(new int[]{3, 7, 8}, 1);     // 2: depends on 3
		closure.addSet(new int[]{9, 11, 12}, 1);  // 3: depends on 4
		closure.addSet(new int[]{89}, 1);       // 4: depends on 2

		closure.addDependencies(2, 3);
		closure.addDependencies(3, 4);
		closure.addDependencies(4, 2);

		closure.addIntersection(new int[]{0, 1}, 2);    // 0 & 1: 3, 5, 9
		closure.addIntersection(new int[]{1, 4}, 2);    // 3, 9, 12
		closure.addIntersection(new int[]{5, 6}, 2);

		closure.addIntersection(new int[]{1, closure.complement(7, 78)}, 2); // 8, 9: = 1 \ {3, 9} = 0 5 12
		closure.addIntersection(new int[]{9, closure.complement(0, 78)}, 2); // 10, 11: = 8 \ 0

		assertTrue(closure.compute());

		assertArrayEquals(new int[]{3, 7, 8, 9, 11, 12, 89}, closure.getSet(3));
		assertArrayEquals(new int[]{3, 5, 9}, closure.getSet(5));
		assertArrayEquals(new int[]{3, 9, 12}, closure.getSet(6));
		assertArrayEquals(new int[]{3, 9}, closure.getSet(7));

		assertArrayEquals(new int[]{0, 5, 12}, closure.getSet(9));
		assertArrayEquals(new int[]{0, 12}, closure.getSet(11));
	}

	@Test
	public void complement() throws Exception {
		SetsClosure closure = new SetsClosure();

		closure.addSet(new int[]{1, 2, 3, 4, 5, 6, 9}, 1);
		closure.addSet(new int[]{0, 3, 5, 9, 12}, 2);

		closure.addDependencies(1, closure.complement(0, 78));

		assertTrue(closure.compute());

		assertArrayEquals(new int[]{1, 2, 3, 4, 5, 6, 9}, closure.getSet(0));
		assertFalse(closure.isComplement(0));

		assertArrayEquals(new int[]{1, 2, 4, 6}, closure.getSet(1));
		assertTrue(closure.isComplement(1));

		BitSet b = new BitSet(13);
		closure.exportIntoBitset(0, 13, b);
		assertEquals("{1, 2, 3, 4, 5, 6, 9}", b.toString());

		b.clear();
		closure.exportIntoBitset(1, 9, b);
		assertEquals("{0, 3, 5, 7, 8}", b.toString());
	}

	@Test
	public void intersectionCycle() throws Exception {
		SetsClosure closure = new SetsClosure();

		closure.addSet(new int[]{1, 2, 3, 4, 5, 6, 9}, 1);
		closure.addSet(new int[]{0, 3, 5, 9, 12}, 2);

		closure.addIntersection(new int[]{0, 1}, 3);

		closure.addSet(new int[]{1, 48}, 2);
		closure.addDependencies(1, 2, 3);

		assertTrue(closure.compute());

		assertArrayEquals(new int[]{1, 2, 3, 4, 5, 6, 9}, closure.getSet(0));
		assertArrayEquals(new int[]{0, 1, 3, 5, 9, 12, 48}, closure.getSet(1));
		assertArrayEquals(new int[]{1, 3, 5, 9}, closure.getSet(2));
	}

	@Test
	public void complementError() throws Exception {
		SetsClosure closure = new SetsClosure();

		closure.addSet(new int[]{1, 2, 3, 4, 5, 6, 9}, 1);
		closure.addSet(new int[]{0, 3, 5, 9, 12}, 2);

		closure.addDependencies(0, 1);
		closure.addDependencies(1, closure.complement(0, 78));

		assertFalse(closure.compute());

		assertArrayEquals(new Object[]{78}, closure.getErrorNodes());
	}

	@Test
	public void testComplexCycle() throws Exception {
		SetsClosure closure = new SetsClosure();
		closure.addSet(new int[]{}, 1);
		closure.addSet(new int[]{}, 2);
		closure.addSet(new int[]{}, 3);
		closure.addSet(new int[]{0}, 4);
		closure.addSet(new int[]{}, 5);
		closure.addSet(new int[]{0}, 7);
		closure.complement(5, 8);
		closure.addIntersection(new int[]{3, 6}, 9);
		closure.addSet(new int[]{}, 10);
		closure.addSet(new int[]{}, 11);

		closure.addDependencies(0, 4);
		closure.addDependencies(4, 1);
		closure.addDependencies(1, 7);
		closure.addDependencies(2, 1);
		closure.addDependencies(3, 8, 9);
		closure.addDependencies(8, 2);
		closure.addDependencies(9, 2);
		assertTrue(closure.compute());

		assertArrayEquals(new int[]{}, closure.getSet(0));
		assertArrayEquals(new int[]{}, closure.getSet(1));
		assertArrayEquals(new int[]{}, closure.getSet(2));
		assertArrayEquals(new int[]{0}, closure.getSet(3));
		assertArrayEquals(new int[]{}, closure.getSet(4));
		assertArrayEquals(new int[]{0}, closure.getSet(5));
		assertArrayEquals(new int[]{0}, closure.getSet(6));
		assertTrue(closure.isComplement(6));
		assertArrayEquals(new int[]{}, closure.getSet(7));
		assertArrayEquals(new int[]{}, closure.getSet(8));
		assertArrayEquals(new int[]{}, closure.getSet(9));
	}
}

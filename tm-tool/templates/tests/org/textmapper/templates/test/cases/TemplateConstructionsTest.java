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
package org.textmapper.templates.test.cases;

import org.junit.Test;
import org.textmapper.templates.api.EvaluationContext;
import org.textmapper.templates.api.types.ITypesRegistry;
import org.textmapper.templates.bundle.DefaultTemplateLoader;
import org.textmapper.templates.bundle.IBundleLoader;
import org.textmapper.templates.bundle.StringTemplateLoader;
import org.textmapper.templates.bundle.TemplatesRegistry;
import org.textmapper.templates.eval.DefaultStaticMethods;
import org.textmapper.templates.eval.TemplatesFacade;
import org.textmapper.templates.objects.JavaIxFactory;
import org.textmapper.templates.storage.ClassResourceLoader;
import org.textmapper.templates.storage.Resource;
import org.textmapper.templates.storage.ResourceRegistry;
import org.textmapper.templates.types.TypesRegistry;

import java.net.URI;
import java.util.*;

import static org.junit.Assert.assertEquals;

public class TemplateConstructionsTest {

	private static final String TEMPLATES_LOCATION = "org/textmapper/templates/test/ltp";

	private static final String TEMPLATES_CHARSET = "utf8";

	// loop.ltp
	@Test
	public void testLoops() {
		Hashtable<String, String[]> h = new Hashtable<>();
		h.put("list", new String[]{"a", "b"});

		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		// test 1
		String q = env.executeTemplate("loop.loop1", new EvaluationContext(h), null, null);
		assertEquals("Hmm: \n\n0: a\n1: b\n\n", q);

		// test 2
		q = env.executeTemplate("loop.loop2", new EvaluationContext(h), null, null);
		assertEquals("\nHmm: \n\n0: a\n1: b\n\n", q);

		// test 3
		h.put("list", new String[]{});
		q = env.executeTemplate("loop.loop2", new EvaluationContext(h), null, null);
		assertEquals("\nHmm: \n\n\n", q);

		// test 4
		q = env.executeTemplate("loop.loop3", new EvaluationContext(h), null, null);
		assertEquals("2\n3\n4\n5\n6\n", q);

		// test 5
		q = env.executeTemplate("loop.loop4", new EvaluationContext(h), null, null);
		assertEquals("2, 3, 4, 5, 6", q);

		// test 6
		q = env.executeTemplate("loop.loop5", new EvaluationContext(h), null, null);
		assertEquals("2:1:6", q);
	}

	// eval.ltp
	@Test
	public void testEval() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		// test 1
		String q = env.executeTemplate("eval.eval1", null, null, null);
		assertEquals("w1 is bad\nw2 is bad\nt4 is bad\n", q);
	}

	// query.ltp
	@Test
	public void testQuery() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		// test 1
		String q = env.executeTemplate("query.a", new EvaluationContext(new Object()), null, null);
		assertEquals("\n123\n", q);
	}

	// dollar.ltp
	@Test
	public void testDollar() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		// test 1
		String q = env.executeTemplate("dollar.testdollar", null, null, null);
		assertEquals("$ is dollar\n", q);

		// test 2
		EvaluationContext context = new EvaluationContext(null);
		context.setVariable("$", "My");
		q = env.executeTemplate("dollar.testdollarvar", context, null, null);
		assertEquals("My is value, $ is dollar\n", q);

		// test 3
		q = env.executeTemplate("dollar.testdollarindex", null, null, null);
		assertEquals("ww-yt-7\n", q);

		// test 4
		collector.addErrors("dollar.ltp,12: Evaluation of `self[2]` failed for java.lang.Object[]: 2 is out of 0..1");
		q = env.executeTemplate("dollar.testdollarindexerr", null, null, null);
		collector.assertEmptyErrors();
		assertEquals("ww-yt-\n", q);

		// test 5
		q = env.executeTemplate("dollar.testsharp", null, null, null);
		assertEquals("\n88 - 67\n99 - 45\n77 - 54\n\n", q);
	}

	// filter.ltp
	@Test
	public void testMap() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		EvaluationContext context = new EvaluationContext(null);
		context.setVariable("util", new DefaultStaticMethods());

		// test 1
		String q = env.executeTemplate("filter.map1", context, null, null);
		assertEquals("[nbsss -> a3,a45 -> 943q,ano -> yes]\n", q);
	}

	@Test
	public void testCollect() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		EvaluationContext context = new EvaluationContext(null);
		context.setVariable("util", new DefaultStaticMethods());

		// test 1
		String q = env.executeTemplate("filter.collectorUnique", context, null, null);
		assertEquals("1A BB C D ", q);

		// test 2
		q = env.executeTemplate("filter.collectorStd", context, null, null);
		assertEquals("1A BB BB C D D C ", q);
	}

	@Test
	public void testSort() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		EvaluationContext context = new EvaluationContext(null);
		context.setVariable("util", new DefaultStaticMethods());

		// test 1
		String q = env.executeTemplate("filter.sorted1", context, null, null);
		assertEquals("1a -> yo4; a -> yo1; daa -> yo2; xb -> yo3; ", q);

		// test 2
		q = env.executeTemplate("filter.sorted2", context, null, null);
		assertEquals("a -> yo1; daa -> yo2; xb -> yo3; 1a -> yo4; ", q);
	}

	@Test
	public void testMax() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		EvaluationContext context = new EvaluationContext(null);
		context.setVariable("util", new DefaultStaticMethods());

		// test 1
		String q = env.executeTemplate("filter.max1", context, null, null);
		assertEquals("7 0 10", q);
	}

	@Test
	public void testGroupBy() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		EvaluationContext context = new EvaluationContext(null);
		context.setVariable("util", new DefaultStaticMethods());

		// test 1
		String q = env.executeTemplate("filter.grouped", context, null, null);
		assertEquals("->  1a:yo1 a:yo1\n" +
				"man:yo23\n" +
				"->  xb:yo2 daa:yo2\n" +
				"rtt:yo3\n", q);
	}

	// arithm.ltp
	@Test
	public void testArithm() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		// test 1
		String q = env.executeTemplate("arithm.arithm1", new EvaluationContext(null), null, null);
		assertEquals("\n2 = 2\n\n", q);

		// test 2
		q = env.executeTemplate("arithm.arithm2", new EvaluationContext(null), null, null);
		assertEquals("\n10 = 10\n\n", q);

		// test 3
		q = env.executeTemplate("arithm.arithm3", new EvaluationContext(null), null, null);
		assertEquals("\n-1 = -1\n\n", q);

		// test 4
		q = env.executeTemplate("arithm.arithm4", new EvaluationContext(null), null, null);
		assertEquals("2 = 2\n", q);

		// test 5
		q = env.executeTemplate("arithm.arithm5", new EvaluationContext(null), null, null);
		assertEquals("true false true true false -2\n", q);

		// test 6
		q = env.executeTemplate("arithm.arithm6", new EvaluationContext(null), null, null);
		assertEquals("uh: lite1\noh: okey\n", q);

		// assign
		q = env.executeTemplate("arithm.assign1", new EvaluationContext(null), null, null);
		assertEquals("30\n42", q);

		// instanceof
		q = env.executeTemplate("arithm.instanceof1", new EvaluationContext(null), null, null);
		assertEquals("true\ntrue\ntrue\n", q);
	}

	// assert.ltp
	@Test
	public void testAssert() {
		Hashtable<String, String[]> h = new Hashtable<>();
		h.put("list", new String[]{"w1", "w2"});
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		// test 1
		collector.addErrors("assert.ltp,4: Evaluation of `l` failed for java.util.Hashtable: null");
		collector.addErrors("assert.ltp,5: Assertion `list.length > 5` failed for java.util.Hashtable");
		collector.addErrors("assert.ltp,7: Assertion `list[1] == 'w4'` failed for java.util.Hashtable");
		env.executeTemplate("assert.assertit", new EvaluationContext(h), null, null);
		collector.assertEmptyErrors();
	}

	// TODO split call & format
	@Test
	public void testCall() {
		Hashtable<String, String[]> h = new Hashtable<>();
		h.put("list", new String[]{"a", "b"});

		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		EvaluationContext context = new EvaluationContext(h);
		context.setVariable("util", new DefaultStaticMethods());

		// test 1
		String q = env.executeTemplate("format.callTempl", context, null, null);
		assertEquals("\nstatic int a[] {\n	0,\n1,\n2,\n3\n};\n\n", q);

		// test 2
		q = env.executeTemplate("format.useformat", context, null, null);
		assertEquals("\nstatic int a[] {\n	1,2,aa,4,5,\n6,7,8,9,10,\n11,12,13,14,\n15,16,17,19,\n20,21,22,23,\n24,25\n};\n\n", q);

		// test 3
		q = env.executeTemplate("format.useCall2", context, null, null);
		assertEquals("Table is mine\n", q);

		// test 4
		q = env.executeTemplate("format.useCall3", context, null, null);
		assertEquals("site is mine\n", q);
	}

	@Test
	public void testOverrides() {
		Hashtable<String, String[]> h = new Hashtable<>();
		h.put("list", new String[]{"a", "b"});

		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector,
				"${template overrides.my2}go next my2(${call base})\n\n${end}"), collector);

		EvaluationContext context = new EvaluationContext(h);
		context.setVariable("util", new DefaultStaticMethods());

		// test 1
		String q = env.executeTemplate("overrides.my1", context, null, null);
		assertEquals("my1\n", q);

		// test 2
		q = env.executeTemplate("overrides.my2", context, null, null);
		assertEquals("go next my2(my2\n)\n\n", q);
	}

	@Test
	public void testOverrides2() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector,
				"${template overrides.my1(aa)}go next my1\n\n${end}"), collector);

		EvaluationContext context = new EvaluationContext(null);

		// test 1
		collector.addErrors("inline,1: Template `my1(aa)` is not compatible with base template `my1()`");
		collector.addErrors("inline,1: Wrong number of arguments used while calling `my1(aa)`: should be 1 instead of 0");
		String q = env.executeTemplate("overrides.my1", context, null, null);
		assertEquals("", q);
	}

	@Test
	public void testFile() {
		final Map<String, String> fileContent = new HashMap<>();
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector) {
			@Override
			public void createStream(String name, String contents) {
				fileContent.put(name, contents);
			}
		};
		EvaluationContext context = new EvaluationContext(new String[]{"aa", "bb"});

		// test 1
		String q = env.executeTemplate("file.file1", context, null, null);
		assertEquals("", q);
		assertEquals("Next\n", fileContent.get("aaaa.txt"));
		assertEquals(1, fileContent.size());
		fileContent.clear();

		// test 2
		q = env.executeTemplate("file.file2", context, null, null);
		assertEquals("", q);
		assertEquals(2, fileContent.size());
		assertEquals("Next\n", fileContent.get("aa.txt"));
		assertEquals("Next2\n", fileContent.get("bb.txt"));
	}

	// switch.ltp
	@Test
	public void testSwitch() {
		final Map<String, Object> this_ = new HashMap<>();
		this_.put("aa", 11);

		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		EvaluationContext context = new EvaluationContext(this_);

		// test 1
		String q = env.executeTemplate("switch.check1", context, null, null);
		assertEquals("Yo", q);

		// test 2
		this_.put("aa", "abcd");
		q = env.executeTemplate("switch.check1", context, null, null);
		assertEquals("Ye", q);

		// test 3
		this_.put("aa", 12);
		q = env.executeTemplate("switch.check1", context, null, null);
		assertEquals("No", q);
	}

	// types.ltp
	@Test
	public void testTypes() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		EvaluationContext context = new EvaluationContext(null);

		// test 1
		collector.addErrors("types.ltp,2: `types.Symbol` is not a subtype of `types.Symbol[]`");
		String q = env.executeTemplate("types.newClass", context, null, null);
		assertEquals("template process\n", q);
		collector.assertEmptyErrors();

		// test 2
		q = env.executeTemplate("types.newClassCorrect", context, null, null);
		assertEquals("template process\n", q);

		// test 3
		collector.addErrors("types.ltp,10: `types.Symbol` is not a subtype of `types.SubSymbol`");
		q = env.executeTemplate("types.newClassInvalidSubclassing", context, null, null);
		assertEquals("template process\n", q);
		collector.assertEmptyErrors();
	}

	// closures.ltp
	@Test
	public void testClosures() {
		TestProblemCollector collector = new TestProblemCollector();
		TemplatesFacade env = new TemplatesFacade(new JavaIxFactory(), createRegistry(collector), collector);

		EvaluationContext context = new EvaluationContext(null);
		context.setVariable("collector", new ObjectCollector());

		// test 1
		String q = env.executeTemplate("closures.test1", context, null, null);
		assertEquals("=22 and 42", q);

		// test 2
		q = env.executeTemplate("closures.test2", context, null, null);
		assertEquals("0: 9, 1: 15, 2: 21, 3: 27", q);

		// test 3
		q = env.executeTemplate("closures.loopTest2", context, null, null);
		assertEquals("Is: 24,42,72", q);
	}

	private TemplatesRegistry createRegistry(TestProblemCollector collector) {
		ResourceRegistry resources = new ResourceRegistry(new ClassResourceLoader(getClass().getClassLoader(), TEMPLATES_LOCATION, TEMPLATES_CHARSET));
		ITypesRegistry types = new TypesRegistry(resources, collector);
		return new TemplatesRegistry(collector, types, (IBundleLoader) new DefaultTemplateLoader(resources));
	}

	private TemplatesRegistry createRegistry(TestProblemCollector collector, String inlineCode) {
		ResourceRegistry resources = new ResourceRegistry(new ClassResourceLoader(getClass().getClassLoader(), TEMPLATES_LOCATION, TEMPLATES_CHARSET));
		ITypesRegistry types = new TypesRegistry(resources, collector);
		return new TemplatesRegistry(collector, types, new StringTemplateLoader(new Resource(URI.create("inline"), inlineCode)), new DefaultTemplateLoader(resources));
	}

	public static class ObjectCollector implements Iterable {
		private List<Object> inner = new ArrayList<>();

		public void add(Object o) {
			inner.add(o);
		}

		@Override
		public Iterator iterator() {
			return inner.iterator();
		}
	}
}

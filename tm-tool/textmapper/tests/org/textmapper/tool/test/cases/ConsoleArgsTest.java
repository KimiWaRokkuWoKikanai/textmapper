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
package org.textmapper.tool.test.cases;

import org.junit.Test;
import org.textmapper.tool.gen.TMOptions;
import org.textmapper.tool.test.CheckingErrorStream;

import java.io.IOException;
import java.io.OutputStream;
import java.io.PrintStream;
import java.io.UnsupportedEncodingException;

import static org.junit.Assert.*;

/**
 * Tests for {@link org.textmapper.tool.gen.TMOptions} command-line arguments parsing.
 */
public class ConsoleArgsTest {

	private static class FailingStream extends OutputStream {

		@Override
		public void write(int b) throws IOException {
			fail("write-protected stream");
		}
	}

	private CheckingErrorStream current;
	private final PrintStream originalErr = System.err;
	protected PrintStream failingStream = new PrintStream(new FailingStream());

	protected void expectError(String expect) {
		try {
			String nl = System.getProperty("line.separator");
			if (nl != null && !nl.equals("\n")) {
				expect = expect.replaceAll("\n", nl);
			}
			closeError();
			System.setErr(new PrintStream(current = new CheckingErrorStream(expect), true, "utf8"));
		} catch (UnsupportedEncodingException e) {
			fail("Exception: " + e.getMessage());
		}
	}

	protected void closeError() {
		if (current != null) {
			try {
				System.setErr(originalErr);
				current.close();
			} catch (IOException e) {
				fail("Exception: " + e.getMessage());
			}
			current = null;
		}
	}

	@Test
	public void testCheckNoArgs() {
		TMOptions lo = TMOptions.parseArguments(new String[0], failingStream);
		assertNotNull(lo);
		assertEquals(null, lo.getInput());
		assertNull(lo.getOutputDirectory());
		assertEquals(0, lo.getDebug());
		assertNull(lo.getTemplateName());
		assertEquals(0, lo.getAdditionalOptions().size());
		assertEquals(0, lo.getIncludeFolders().size());
		assertEquals(true, lo.isUseDefaultTemplates());
	}

	@Test
	public void testCheckDebug() {
		TMOptions lo = TMOptions.parseArguments("-e".split(" "), failingStream);
		assertNotNull(lo);
		assertEquals(null, lo.getInput());
		assertEquals(TMOptions.DEBUG_TABLES, lo.getDebug());
		lo = TMOptions.parseArguments("-d".split(" "), failingStream);
		assertEquals(TMOptions.DEBUG_AMBIG, lo.getDebug());
	}

	@Test
	public void testInput() {
		TMOptions lo = TMOptions.parseArguments("-e synt1".split(" "), failingStream);
		assertNotNull(lo);
		assertEquals("synt1", lo.getInput());
		assertEquals(TMOptions.DEBUG_TABLES, lo.getDebug());
	}

	@Test
	public void testInput2() {
		expectError("textmapper: should be only one input in arguments\n");
		TMOptions lo = TMOptions.parseArguments("synt2 synt1".split(" "), System.err);
		assertNull(lo);
		closeError();
	}

	@Test
	public void testTwiceArg() {
		expectError("textmapper: option cannot be used twice -e\n");
		TMOptions lo = TMOptions.parseArguments("-e -e".split(" "), System.err);
		assertNull(lo);
		closeError();

		expectError("textmapper: option cannot be used twice -x\n");
		lo = TMOptions.parseArguments("--no-default-templates -x".split(" "), System.err);
		assertNull(lo);
		closeError();

		expectError("textmapper: option cannot be used twice --no-default-templates\n");
		lo = TMOptions.parseArguments("-x --no-default-templates".split(" "), System.err);
		assertNull(lo);
		closeError();
	}

	@Test
	public void testError() {
		expectError("textmapper: no value for option -o\n");
		TMOptions lo = TMOptions.parseArguments("-o".split(" "), System.err);
		assertNull(lo);
		closeError();

		expectError("textmapper: no value for option --output\n");
		lo = TMOptions.parseArguments("--output".split(" "), System.err);
		assertNull(lo);
		closeError();

		expectError("textmapper: no value for option --output=\n");
		lo = TMOptions.parseArguments("--output=".split(" "), System.err);
		assertNull(lo);
		closeError();

		expectError("textmapper: invalid option --output1=we\n");
		lo = TMOptions.parseArguments("--output1=we".split(" "), System.err);
		assertNull(lo);
		closeError();

		expectError("textmapper: invalid option -q\n");
		lo = TMOptions.parseArguments("-q".split(" "), System.err);
		assertNull(lo);
		closeError();

		expectError("textmapper: key is used twice: a\n");
		lo = TMOptions.parseArguments("a=2 a=5".split(" "), System.err);
		assertNull(lo);
		closeError();
	}

	@Test
	public void testComplicated() {
		TMOptions lo = TMOptions.parseArguments("-e -x -o outputY -i folder1;folder2 -i folder3 -t java2 a=5 b=6 syntax.g".split(" "), failingStream);
		assertNotNull(lo);
		assertEquals("syntax.g", lo.getInput());
		assertEquals("outputY", lo.getOutputDirectory());
		assertEquals(TMOptions.DEBUG_TABLES, lo.getDebug());
		assertEquals("java2", lo.getTemplateName());
		assertEquals(3, lo.getIncludeFolders().size());
		assertEquals("folder1", lo.getIncludeFolders().get(0));
		assertEquals("folder2", lo.getIncludeFolders().get(1));
		assertEquals("folder3", lo.getIncludeFolders().get(2));
		assertEquals(2, lo.getAdditionalOptions().size());
		assertEquals("5", lo.getAdditionalOptions().get("a"));
		assertEquals("6", lo.getAdditionalOptions().get("b"));
		assertEquals(false, lo.isUseDefaultTemplates());
	}
}

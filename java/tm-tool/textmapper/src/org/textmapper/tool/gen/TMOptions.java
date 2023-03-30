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
package org.textmapper.tool.gen;

import java.io.PrintStream;
import java.util.*;

/**
 * Represents Textmapper options.
 */
public class TMOptions {

	public static final int DEBUG_AMBIG = 1;
	public static final int DEBUG_TABLES = 2;

	private int debug;

	private String input;
	private String outputDir;
	private String templateName;

	private final List<String> includeFolders;
	private final Map<String, String> templateOptions;

	private boolean useDefaultTemplates;

	public TMOptions() {
		this.debug = 0;
		this.input = null;
		this.outputDir = null;
		this.templateName = null;
		this.includeFolders = new LinkedList<>();
		this.useDefaultTemplates = true;
		this.templateOptions = new HashMap<>();
	}

	public String getInput() {
		return input;
	}

	public void setInput(String input) {
		this.input = input;
	}

	public String getOutputDirectory() {
		return outputDir;
	}

	public void setOutputDirectory(String outputDir) {
		this.outputDir = outputDir;
	}

	public String getTemplateName() {
		return templateName;
	}

	public void setTemplateName(String templateName) {
		this.templateName = templateName;
	}

	public int getDebug() {
		return debug;
	}

	public void setDebug(int debug) {
		this.debug = debug;
	}

	public boolean isUseDefaultTemplates() {
		return useDefaultTemplates;
	}

	public void setUseDefaultTemplates(boolean useDefaultTemplates) {
		this.useDefaultTemplates = useDefaultTemplates;
	}

	public Map<String, String> getAdditionalOptions() {
		return templateOptions;
	}

	public List<String> getIncludeFolders() {
		return includeFolders;
	}

	static final int HAS_VALUE = 32;
	static final int MULTI_VALUE = 64;

	static final int OPT_DEBUG = 1;
	static final int OPT_EXT_DEBUG = 2;
	static final int OPT_NO_DEF = 3;
	static final int OPT_OUTPUT = 4 | HAS_VALUE;
	static final int OPT_INCLUDE = 5 | HAS_VALUE | MULTI_VALUE;
	static final int OPT_TEMPLATE = 6 | HAS_VALUE;
	static final int OPT_INPUT = 7 | HAS_VALUE;

	public static final String HELP_OPTIONS =
		"  -d,  --debug                   debug info\n" +
		"  -e,  --extended-debug          extended debug info\n" +
		"  -x,  --no-default-templates    removes default templates from engine\n" +
		"  -o dir, --output=dir           target directory\n" +
		"  -i dir, --include=dir          adds directory (or semicolon separated directory list) " +
		"                                 to the textmapper.templates stack\n" +
		"  -t templateId, --template=id   use template for generation\n" +
		"  key=val                        any generation option\n";

	private static Map<String, Integer> buildOptionsHash() {
		Map<String, Integer> res = new HashMap<>();
		res.put("d", OPT_DEBUG);
		res.put("-debug", OPT_DEBUG);
		res.put("e", OPT_EXT_DEBUG);
		res.put("-extended-debug", OPT_EXT_DEBUG);
		res.put("x", OPT_NO_DEF);
		res.put("-no-default-templates", OPT_NO_DEF);
		res.put("o", OPT_OUTPUT);
		res.put("-output", OPT_OUTPUT);
		res.put("i", OPT_INCLUDE);
		res.put("-include", OPT_INCLUDE);
		res.put("t", OPT_TEMPLATE);
		res.put("-template", OPT_TEMPLATE);
		return res;
	}

	public static TMOptions parseArguments(String[] args, PrintStream errorStream) {
		TMOptions opts = new TMOptions();
		Map<String, Integer> optionsHash = buildOptionsHash();
		Set<Integer> usedOptions = new HashSet<>();

		for (int i = 0; i < args.length; i++) {
			int equalIndex = args[i].indexOf('=');
			if (args[i].length() > 1 && args[i].charAt(0) == '-') {
				String option = equalIndex >= 0 ? args[i].substring(1, equalIndex) : args[i].substring(1);
				int optionId = optionsHash.containsKey(option) ? optionsHash.get(option) : 0;
				boolean hasValue = (optionId & HAS_VALUE) != 0;
				if (optionId == 0 || equalIndex >= 0 && !hasValue) {
					errorStream.println("textmapper: invalid option " + args[i]);
					return null;
				}
				if (hasValue && (equalIndex < 0 && i + 1 == args.length || equalIndex + 1 == args[i].length())) {
					errorStream.println("textmapper: no value for option " + args[i]);
					return null;
				}
				if (usedOptions.contains(optionId)) {
					errorStream.println("textmapper: option cannot be used twice " + args[i]);
					return null;
				}
				if ((optionId & MULTI_VALUE) == 0) {
					usedOptions.add(optionId);
				}
				setOption(opts, optionId,
						hasValue ? (equalIndex >= 0 ? args[i].substring(equalIndex + 1) : args[++i]) : null);

			} else if (equalIndex >= 0) {
				String key = args[i].substring(0, equalIndex);
				String value = args[i].substring(equalIndex + 1);
				if (opts.getAdditionalOptions().containsKey(key)) {
					errorStream.println("textmapper: key is used twice: " + key);
					return null;
				}
				opts.getAdditionalOptions().put(key, value);

			} else {
				if (usedOptions.contains(OPT_INPUT)) {
					errorStream.println("textmapper: should be only one input in arguments");
					return null;
				}
				usedOptions.add(OPT_INPUT);
				opts.setInput(args[i]);
			}
		}
		return opts;
	}

	private static void setOption(TMOptions opts, int optionId, String value) {
		switch (optionId) {
			case OPT_DEBUG:
				opts.setDebug(TMOptions.DEBUG_AMBIG);
				break;
			case OPT_EXT_DEBUG:
				opts.setDebug(TMOptions.DEBUG_TABLES);
				break;
			case OPT_NO_DEF:
				opts.setUseDefaultTemplates(false);
				break;
			case OPT_INCLUDE:
				for (String s : value.split(";")) {
					if (s.trim().length() > 0) {
						opts.getIncludeFolders().add(s);
					}
				}
				break;
			case OPT_TEMPLATE:
				opts.setTemplateName(value);
				break;
			case OPT_OUTPUT:
				opts.setOutputDirectory(value);
				break;
		}
	}
}

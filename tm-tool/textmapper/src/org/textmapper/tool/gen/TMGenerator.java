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

import org.textmapper.lapg.LapgCore;
import org.textmapper.lapg.api.*;
import org.textmapper.lapg.api.ast.AstModel;
import org.textmapper.templates.api.EvaluationContext;
import org.textmapper.templates.api.TemplatesStatus;
import org.textmapper.templates.api.types.IClass;
import org.textmapper.templates.bundle.DefaultTemplateLoader;
import org.textmapper.templates.bundle.IBundleLoader;
import org.textmapper.templates.bundle.StringTemplateLoader;
import org.textmapper.templates.bundle.TemplatesRegistry;
import org.textmapper.templates.eval.TemplatesFacade;
import org.textmapper.templates.objects.IxFactory;
import org.textmapper.templates.storage.ClassResourceLoader;
import org.textmapper.templates.storage.IResourceLoader;
import org.textmapper.templates.storage.Resource;
import org.textmapper.templates.storage.ResourceRegistry;
import org.textmapper.templates.types.TiInstance;
import org.textmapper.templates.types.TypesRegistry;
import org.textmapper.tool.compiler.TMDataUtil;
import org.textmapper.tool.compiler.TMEventMapper;
import org.textmapper.tool.compiler.TMGrammar;
import org.textmapper.tool.compiler.TMMapper;
import org.textmapper.tool.parser.TMTree.TextSource;

import javax.script.Invocable;
import javax.script.ScriptEngine;
import javax.script.ScriptEngineManager;
import javax.script.ScriptException;
import java.io.File;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;

public final class TMGenerator {

	private final TMOptions options;
	private final ProcessingStatus status;
	private final ProcessingStrategy strategy;

	public TMGenerator(TMOptions options, ProcessingStatus status, ProcessingStrategy strategy) {
		this.options = options;
		this.status = status;
		this.strategy = strategy;
	}

	public boolean compileGrammar(TextSource input, boolean checkOnly) {
		try {
			TemplatesStatusAdapter templatesStatus = new TemplatesStatusAdapter(status);
			ResourceRegistry resources = createResourceRegistry();
			TypesRegistry types = new TypesRegistry(resources, templatesStatus);

			TMGrammar s = SyntaxUtil.parseSyntax(input, status, types);
			if (s == null || s.hasErrors()) {
				return false;
			}

			// prepare options
			Map<String, Object> genOptions = new HashMap<>(s.getOptions());
			for (Entry<String, String> entry : options.getAdditionalOptions().entrySet()) {

				// TODO parse value, check type
				genOptions.put(entry.getKey(), entry.getValue());
			}

			boolean genParser = Boolean.TRUE.equals(genOptions.get("genParser")) &&
					s.getGrammar().getRules() != null;

			// Language-specific processing.
			AstModel astModel = null;
			if (genParser) {
				boolean genast = Boolean.TRUE.equals(genOptions.get("genast"));
				boolean eventBased = Boolean.TRUE.equals(genOptions.get("eventBased"));

				if (genast || eventBased) {
					boolean hasAny = Boolean.TRUE.equals(genOptions.get("__hasAny"));
					if (genast) {
						astModel = new TMMapper(s.getGrammar(), status, hasAny).deriveAST();
					} else {
						new TMEventMapper(s.getGrammar(), genOptions, status).deriveTypes();
					}
				}
			}

			// Generate user actions.
			long start = System.currentTimeMillis();
			TemplatesRegistry registry = createTemplateRegistry(
					s.getTemplates(), resources, types, templatesStatus);
			EvaluationContext context = createEvaluationContext(types, s, astModel,
					genOptions, null, null);

			TemplatesFacade ruleEnv = new TemplatesFacade(
					new GrammarIxFactory(s, getTemplatePackage(s), context), registry);
			if (genParser) {
				for (Rule rule : s.getGrammar().getRules()) {
					String action = ruleEnv.executeTemplate(getTemplatePackage(s) + ".parserAction",
							context, new Object[]{rule}, null);
					if (action.isEmpty()) continue;
					TMDataUtil.putCode(rule, action);
				}
			}
			for (LexerRule rule : s.getGrammar().getLexerRules()) {
				String action = ruleEnv.executeTemplate(getTemplatePackage(s) + ".lexerAction",
						context, new Object[]{rule}, null);
				if (action.isEmpty()) continue;
				TMDataUtil.putCode(rule, action);
			}
			long userActionsTime = System.currentTimeMillis() - start;

			// Generate tables
			start = System.currentTimeMillis();
			ParserData r = null;
			if (genParser) {
				r = LapgCore.generateParser(s.getGrammar(), status);
				if (r == null) {
					return false;
				}
			}
			boolean allowBacktracking = !Boolean.TRUE.equals(genOptions.get("nonBacktracking"));
			LexerData l = LapgCore.generateLexer(s.getGrammar(), allowBacktracking, status);
			if (l == null) {
				return false;
			}

			if (checkOnly) return true;

			long generationTime = System.currentTimeMillis() - start;

			// Generate text
			start = System.currentTimeMillis();
			context = createEvaluationContext(types, s, astModel, genOptions, l, r);
			TemplatesFacade env = new TemplatesFacadeExt(
					new GrammarIxFactory(s, getTemplatePackage(s), context),
					registry, genOptions);
			env.executeTemplate(getTemplatePackage(s) + ".main", context, null, null);
			long textTime = userActionsTime + (System.currentTimeMillis() - start);
			StringBuilder sb = new StringBuilder();
			sb.append("lalr: ").append(generationTime / 1000.).append("s");
			sb.append(", text: ").append(textTime / 1000.).append("s");
			if (r != null) {
				sb.append(", parser: ").append(r.getStatesCount()).append(" states, ")
						.append(r.getByteSize() / 1024).append("KB");
			}
			status.report(ProcessingStatus.KIND_INFO, sb.toString());
			return true;
		} catch (Exception t) {
			String message = "lapg: internal error: " + t.getClass().getName();
			status.report(message, t);
			return false;
		}
	}

	private String getTemplatePackage(TMGrammar g) {
		String result = options.getTemplateName();
		if (result != null) {
			return result;
		}

		result = g.getTargetLanguage();
		if (result != null) {
			return result;
		}

		return "common";
	}

	private boolean checkOptions(TMGrammar s, TemplatesRegistry registry) {
		String templPackage = getTemplatePackage(s);
		IClass cl = registry.getTypesRegistry().getClass(templPackage + ".Options", null);

		// TODO

		return true;
	}

	private ResourceRegistry createResourceRegistry() {
		List<IResourceLoader> loaders = new ArrayList<>();
		for (String path : options.getIncludeFolders()) {
			IResourceLoader resourceLoader = strategy.createResourceLoader(path);
			if (resourceLoader != null) {
				loaders.add(resourceLoader);
			}
		}
		if (options.isUseDefaultTemplates()) {
			loaders.add(new ClassResourceLoader(getClass().getClassLoader(),
					"org/textmapper/tool/templates", "utf8"));
		}
		return new ResourceRegistry(loaders.toArray(new IResourceLoader[loaders.size()]));
	}

	private TemplatesRegistry createTemplateRegistry(TextSourceElement grammarTemplates,
													 ResourceRegistry resources,
													 TypesRegistry types,
													 TemplatesStatus templatesStatus) {
		List<IBundleLoader> loaders = new ArrayList<>();
		if (grammarTemplates != null) {
			File file = new File(grammarTemplates.getResourceName());
			loaders.add(new StringTemplateLoader(new Resource(file.toURI(),
					grammarTemplates.getText(), grammarTemplates.getLine(),
					grammarTemplates.getOffset())));
		}
		loaders.add(new DefaultTemplateLoader(resources));
		return new TemplatesRegistry(templatesStatus, types,
				loaders.toArray(new IBundleLoader[loaders.size()]));
	}

	private EvaluationContext createEvaluationContext(TypesRegistry types, TMGrammar s,
													  AstModel astModel,
													  Map<String, Object> genOptions, LexerData l,
													  ParserData r) {
		Map<String, Object> map = new HashMap<>();
		map.put("syntax", s);
		map.put("lex", l);
		map.put("parser", r);
		if (astModel != null) {
			map.put("ast", astModel);
		}

		String templPackage = getTemplatePackage(s);
		IClass optsClass = types.getClass(templPackage + ".Options", null);
		if (optsClass != null) {
			map.put("opts", new TiInstance(optsClass, genOptions));
		} else {
			map.put("opts", genOptions);
		}


		TiInstance context = new TiInstance(types.getClass("common.Context", null), map);
		EvaluationContext evaluationContext = new EvaluationContext(context);
		evaluationContext.setVariable("util", new TemplateStaticMethods());
		evaluationContext.setVariable("context", context);
		evaluationContext.setVariable("$", "tmLeft.value");
		return evaluationContext;
	}

	private final class TemplatesFacadeExt extends TemplatesFacade {
		Map<String, Object> options;

		private TemplatesFacadeExt(IxFactory factory, TemplatesRegistry registry,
								   Map<String, Object> options) {
			super(factory, registry);
			this.options = options;
		}

		@Override
		public void createStream(String name, String contents) {
			strategy.createFile(name, contents, options, status);
		}
	}

	public final class ScriptBuiltin {
		private final ScriptEngine engine;
		private final ResourceRegistry resources;
		private final Map<String, Object> options;

		private ScriptBuiltin(ScriptEngine engine, ResourceRegistry resources,
							  Map<String, Object> options) {
			this.engine = engine;
			this.resources = resources;
			this.options = options;
		}

		public void write(String name, String contents) {
			strategy.createFile(name, contents, options, status);
		}

		public void load(String name) throws ScriptException {
			Resource[] res = resources.loadResources(name, "js");
			if (res == null || res.length == 0) {
				throw new RuntimeException("No resource found: " + name);
			}
			engine.eval(res[0].getContents());
		}
	}
}

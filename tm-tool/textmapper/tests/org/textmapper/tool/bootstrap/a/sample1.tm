#  syntax: sample1 grammar

#  Copyright 2002-2022 Evgeny Gryaznov
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

language sample1(java);

prefix = "SampleA"
package = "org.textmapper.tool.bootstrap.a"
maxtoken = 2048
breaks = true
genast = true
gentree = true
genastdef = true
positions = "line,column,offset"
endpositions = "line,column,offset"
genCleanup = false
genCopyright = true

:: lexer

identifier {String}: /[a-zA-Z_][a-zA-Z_0-9]*/ -1 { $$ = tokenText(); }
_skip:          /[\n\t\r ]+/ (space)

Lclass: /class/
'{': /\{/
'}': /\}/

error:

:: parser

%input classdef_no_eoi no-eoi, classdef ;

classdef_no_eoi interface :
	classdef ;

classdef :
	Lclass identifier '{' classdeflistopt '}' ;

classdeflist :
	classdef
  | classdeflist classdef
  | error
;

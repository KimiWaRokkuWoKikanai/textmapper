// generated by Textmapper; DO NOT EDIT

package tm

import (
	"fmt"
)

type NodeType int

type Listener func(t NodeType, offset, endoffset int)

const (
	NoType NodeType = iota
	Identifier
	IntegerLiteral
	StringLiteral
	BooleanLiteral
	Pattern
	Command
	SyntaxProblem
	File          // Header imports=(Import)* options=(Option)* SyntaxProblem? lexer=LexerSection? parser=ParserSection?
	Header        // name=Identifier target=Identifier?
	LexerSection  // (LexerPart)+
	ParserSection // (GrammarPart)+
	Import        // alias=Identifier? path=StringLiteral
	Option        // key=Identifier value=Expression
	Symref        // name=Identifier args=SymrefArgs?
	RawType
	NamedPattern         // name=Identifier Pattern
	StartConditionsScope // StartConditions (LexerPart)+
	StartConditions      // (Stateref)*
	Lexeme               // StartConditions? name=Identifier RawType? Pattern? priority=IntegerLiteral? attrs=LexemeAttrs? Command?
	LexemeAttrs          // LexemeAttribute
	LexemeAttribute
	DirectiveBrackets   // opening=Symref closing=Symref
	InclusiveStartConds // states=(LexerState)+
	ExclusiveStartConds // states=(LexerState)+
	Stateref            // name=Identifier
	LexerState          // name=Identifier
	Nonterm             // Annotations? name=Identifier params=NontermParams? NontermType? ReportClause? (Rule0)+
	SubType             // reference=Symref
	InterfaceType
	ClassType // Implements?
	VoidType
	Implements // (Symref)+
	Assoc
	ParamModifier
	TemplateParam      // modifier=ParamModifier? ParamType name=Identifier ParamValue?
	DirectivePrio      // Assoc symbols=(Symref)+
	DirectiveInput     // inputRefs=(Inputref)+
	DirectiveInterface // ids=(Identifier)+
	Empty
	NonEmpty
	DirectiveAssert // Empty? NonEmpty? RhsSet
	DirectiveSet    // name=Identifier RhsSet
	NoEoi
	Inputref  // reference=Symref NoEoi?
	Rule      // Predicate? (RhsPart)* RhsSuffix? ReportClause?
	Predicate // PredicateExpression
	Name
	RhsSuffix    // Name Symref
	ReportClause // action=Identifier kind=Identifier? ReportAs?
	ReportAs     // Identifier
	RhsLookahead // predicates=(LookaheadPredicate)+
	Not
	LookaheadPredicate // Not? Symref
	StateMarker        // name=Identifier
	RhsAnnotated       // Annotations inner=RhsPart
	RhsAssignment      // id=Identifier inner=RhsPart
	RhsPlusAssignment  // id=Identifier inner=RhsPart
	RhsOptional        // inner=RhsPart
	RhsCast            // inner=RhsPart target=Symref
	RhsAsLiteral       // inner=RhsPart Literal
	ListSeparator      // separator_=(Symref)+
	RhsSymbol          // reference=Symref
	RhsNested          // (Rule0)+
	RhsPlusList        // ruleParts=(RhsPart)+ ListSeparator
	RhsStarList        // ruleParts=(RhsPart)+ ListSeparator
	RhsPlusQuantifier  // inner=RhsPart
	RhsStarQuantifier  // inner=RhsPart
	RhsIgnored         // (Rule0)+
	RhsSet             // expr=SetExpression
	SetSymbol          // operator=Identifier? symbol=Symref
	SetCompound        // inner=SetExpression
	SetComplement      // inner=SetExpression
	SetOr              // left=SetExpression right=SetExpression
	SetAnd             // left=SetExpression right=SetExpression
	Annotations        // (Annotation)+
	AnnotationImpl     // name=Identifier Expression?
	NontermParams      // list=(NontermParam)+
	InlineParameter    // param_type=Identifier name=Identifier ParamValue?
	ParamRef           // Identifier
	SymrefArgs         // arg_list=(Argument)*
	ArgumentVal        // name=ParamRef val=ParamValue?
	ArgumentTrue       // name=ParamRef
	ArgumentFalse      // name=ParamRef
	ParamType
	PredicateNot   // ParamRef
	PredicateEq    // ParamRef Literal
	PredicateNotEq // ParamRef Literal
	PredicateAnd   // left=PredicateExpression right=PredicateExpression
	PredicateOr    // left=PredicateExpression right=PredicateExpression
	Array          // (Expression)*
	InvalidToken
	MultilineComment
	Comment
	Templates
	NodeTypeMax
)

var nodeTypeStr = [...]string{
	"NONE",
	"Identifier",
	"IntegerLiteral",
	"StringLiteral",
	"BooleanLiteral",
	"Pattern",
	"Command",
	"SyntaxProblem",
	"File",
	"Header",
	"LexerSection",
	"ParserSection",
	"Import",
	"Option",
	"Symref",
	"RawType",
	"NamedPattern",
	"StartConditionsScope",
	"StartConditions",
	"Lexeme",
	"LexemeAttrs",
	"LexemeAttribute",
	"DirectiveBrackets",
	"InclusiveStartConds",
	"ExclusiveStartConds",
	"Stateref",
	"LexerState",
	"Nonterm",
	"SubType",
	"InterfaceType",
	"ClassType",
	"VoidType",
	"Implements",
	"Assoc",
	"ParamModifier",
	"TemplateParam",
	"DirectivePrio",
	"DirectiveInput",
	"DirectiveInterface",
	"Empty",
	"NonEmpty",
	"DirectiveAssert",
	"DirectiveSet",
	"NoEoi",
	"Inputref",
	"Rule",
	"Predicate",
	"Name",
	"RhsSuffix",
	"ReportClause",
	"ReportAs",
	"RhsLookahead",
	"Not",
	"LookaheadPredicate",
	"StateMarker",
	"RhsAnnotated",
	"RhsAssignment",
	"RhsPlusAssignment",
	"RhsOptional",
	"RhsCast",
	"RhsAsLiteral",
	"ListSeparator",
	"RhsSymbol",
	"RhsNested",
	"RhsPlusList",
	"RhsStarList",
	"RhsPlusQuantifier",
	"RhsStarQuantifier",
	"RhsIgnored",
	"RhsSet",
	"SetSymbol",
	"SetCompound",
	"SetComplement",
	"SetOr",
	"SetAnd",
	"Annotations",
	"AnnotationImpl",
	"NontermParams",
	"InlineParameter",
	"ParamRef",
	"SymrefArgs",
	"ArgumentVal",
	"ArgumentTrue",
	"ArgumentFalse",
	"ParamType",
	"PredicateNot",
	"PredicateEq",
	"PredicateNotEq",
	"PredicateAnd",
	"PredicateOr",
	"Array",
	"InvalidToken",
	"MultilineComment",
	"Comment",
	"Templates",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}

var Annotation = []NodeType{
	AnnotationImpl,
	SyntaxProblem,
}

var Argument = []NodeType{
	ArgumentFalse,
	ArgumentTrue,
	ArgumentVal,
}

var Expression = []NodeType{
	Array,
	BooleanLiteral,
	IntegerLiteral,
	StringLiteral,
	Symref,
	SyntaxProblem,
}

var GrammarPart = []NodeType{
	DirectiveAssert,
	DirectiveInput,
	DirectiveInterface,
	DirectivePrio,
	DirectiveSet,
	Nonterm,
	SyntaxProblem,
	TemplateParam,
}

var LexerPart = []NodeType{
	DirectiveBrackets,
	ExclusiveStartConds,
	InclusiveStartConds,
	Lexeme,
	NamedPattern,
	StartConditionsScope,
	SyntaxProblem,
}

var Literal = []NodeType{
	BooleanLiteral,
	IntegerLiteral,
	StringLiteral,
}

var NontermParam = []NodeType{
	InlineParameter,
	ParamRef,
}

var NontermType = []NodeType{
	ClassType,
	InterfaceType,
	RawType,
	SubType,
	VoidType,
}

var ParamValue = []NodeType{
	BooleanLiteral,
	IntegerLiteral,
	ParamRef,
	StringLiteral,
}

var PredicateExpression = []NodeType{
	ParamRef,
	PredicateAnd,
	PredicateEq,
	PredicateNot,
	PredicateNotEq,
	PredicateOr,
}

var RhsPart = []NodeType{
	Command,
	RhsAnnotated,
	RhsAsLiteral,
	RhsAssignment,
	RhsCast,
	RhsIgnored,
	RhsLookahead,
	RhsNested,
	RhsOptional,
	RhsPlusAssignment,
	RhsPlusList,
	RhsPlusQuantifier,
	RhsSet,
	RhsStarList,
	RhsStarQuantifier,
	RhsSymbol,
	StateMarker,
	SyntaxProblem,
}

var Rule0 = []NodeType{
	Rule,
	SyntaxProblem,
}

var SetExpression = []NodeType{
	SetAnd,
	SetComplement,
	SetCompound,
	SetOr,
	SetSymbol,
}

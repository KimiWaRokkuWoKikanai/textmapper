// generated by Textmapper; DO NOT EDIT

package js

import (
	"fmt"
)

type NodeType int

type Listener func(t NodeType, offset, endoffset int)

const (
	NoType NodeType = iota
	Accessor
	AdditiveExpr // left=Expr right=Expr
	Arguments    // TsTypeArguments? list=(Expr)*
	ArrayLiteral // list=(Expr)*
	ArrayPattern // list=(ElementPattern | Expr)* BindingRestElement?
	ArrowFunc    // NameIdent? TsTypeParameters? Parameters? TsTypeAnnotation? FnArrow Body? ConciseBody? SyntaxProblem?
	AssertClause // (AssertEntry)*
	AssertEntry  // AssertionKey
	AssertionKey
	AssignmentExpr // left=Expr AssignmentOperator? right=Expr
	AssignmentOperator
	AsyncArrowFunc            // NameIdent? TsTypeParameters? Parameters? TsTypeAnnotation? FnArrow Body? ConciseBody?
	AsyncFunc                 // NameIdent? TsTypeParameters? Parameters TsTypeAnnotation? Body?
	AsyncFuncExpr             // NameIdent? TsTypeParameters? Parameters TsTypeAnnotation? Body?
	AsyncGeneratorDeclaration // NameIdent? TsTypeParameters? Parameters TsTypeAnnotation? Body?
	AsyncGeneratorExpression  // NameIdent? TsTypeParameters? Parameters TsTypeAnnotation? Body?
	AsyncGeneratorMethod      // PropertyName TsTypeParameters? Parameters TsTypeAnnotation? Body?
	AsyncMethod               // PropertyName TsTypeParameters? Parameters TsTypeAnnotation? Body?
	Await
	AwaitExpr            // Expr
	BindingRestElement   // NameIdent
	BitwiseAND           // left=Expr right=Expr
	BitwiseOR            // left=Expr right=Expr
	BitwiseXOR           // left=Expr right=Expr
	Block                // (CaseClause)* list=(StmtListItem)*
	Body                 // list=(StmtListItem)*
	BreakStmt            // LabelIdent?
	CallExpr             // expr=Expr Arguments
	Case                 // Cond (StmtListItem)*
	Catch                // BindingPattern? NameIdent? TsTypeAnnotation? Block
	Class                // (Modifier)* NameIdent? TsTypeParameters? Extends? TsImplementsClause? ClassBody
	ClassBody            // (ClassElement)*
	ClassExpr            // (Modifier)* NameIdent? TsTypeParameters? Extends? TsImplementsClause? ClassBody
	CoalesceExpr         // left=Expr right=Expr
	CommaExpr            // left=Expr right=Expr
	ComputedPropertyName // Expr
	ConciseBody          // Expr
	Cond                 // Expr
	ConditionalExpr      // cond=Expr then=Expr else=Expr
	Const
	ContinueStmt // LabelIdent?
	DebuggerStmt
	DecoratorCall    // (ReferenceIdent)+ Arguments
	DecoratorExpr    // (ReferenceIdent)+
	Default          // (StmtListItem)*
	DefaultParameter // (Modifier)* BindingPattern? NameIdent? TsOptional? TsTypeAnnotation? Initializer?
	DoWhileStmt      // Stmt Expr
	ElementBinding   // BindingPattern Initializer?
	EmptyDecl
	EmptyStmt
	EqualityExpr       // left=Expr right=Expr
	ExponentiationExpr // left=Expr right=Expr
	ExportClause       // (ExportElement)*
	ExportDecl         // (Modifier)* TsTypeOnly? VarStmt? Decl? ExportClause? NameIdent? ModuleSpec? AssertClause?
	ExportDefault      // Expr? (Modifier)* Decl?
	ExportSpec         // TsTypeOnly? ReferenceIdent? NameIdent
	ExprStmt           // Expr
	Extends            // Expr? TsTypeReference?
	File               // (ModuleItem)*
	Finally            // Block
	FnArrow
	ForBinding       // BindingPattern? NameIdent?
	ForCondition     // Expr?
	ForFinalExpr     // Expr?
	ForInStmt        // var=Expr object=Expr Stmt
	ForInStmtWithVar // Const? Var? Let? ForBinding object=Expr Stmt
	ForOfStmt        // Await? var=Expr iterable=Expr Stmt
	ForOfStmtWithVar // Await? Const? Var? Let? ForBinding iterable=Expr Stmt
	ForStmt          // var=Expr? ForCondition ForFinalExpr Stmt
	ForStmtWithVar   // Const? Var? Let? (VarDecl)* (LexicalBinding)* ForCondition ForFinalExpr Stmt
	Func             // NameIdent? TsTypeParameters? Parameters TsTypeAnnotation? Body?
	FuncExpr         // NameIdent? TsTypeParameters? Parameters TsTypeAnnotation? Body?
	Generator        // NameIdent? TsTypeParameters? Parameters TsTypeAnnotation? Body?
	GeneratorExpr    // NameIdent? TsTypeParameters? Parameters TsTypeAnnotation? Body?
	GeneratorMethod  // PropertyName TsTypeParameters? Parameters TsTypeAnnotation? Body?
	Getter           // (Modifier)* PropertyName TsTypeAnnotation? Body?
	IdentExpr        // ReferenceIdent
	IfStmt           // Expr then=Stmt else=Stmt?
	ImportDecl       // TsTypeOnly? NameIdent? NameSpaceImport? NamedImports? ModuleSpec AssertClause?
	ImportSpec       // TsTypeOnly? ReferenceIdent? NameIdent
	InExpr           // left=Expr right=Expr
	IndexAccess      // expr=Expr index=Expr
	Initializer      // Expr
	InstanceOfExpr   // left=Expr right=Expr
	JSXAttributeName
	JSXClosingElement // JSXElementName
	JSXElement        // JSXOpeningElement? JSXSelfClosingElement? (JSXChild)* JSXClosingElement?
	JSXElementName
	JSXExpr     // Expr?
	JSXFragment // (JSXChild)*
	JSXLiteral
	JSXNormalAttribute    // JSXAttributeName JSXAttributeValue?
	JSXOpeningElement     // JSXElementName TsTypeArguments? (JSXAttribute)*
	JSXSelfClosingElement // JSXElementName TsTypeArguments? (JSXAttribute)*
	JSXSpreadAttribute    // Expr
	JSXSpreadExpr         // Expr?
	JSXText
	LabelIdent
	LabelledStmt // LabelIdent Func? Stmt?
	Let
	LexicalBinding // BindingPattern? NameIdent? TsExclToken? TsTypeAnnotation? Initializer?
	LexicalDecl    // Const? Let? (LexicalBinding)+
	Literal
	LiteralPropertyName // Literal? NameIdent?
	LogicalAND          // left=Expr right=Expr
	LogicalOR           // left=Expr right=Expr
	MemberMethod        // (Modifier)* MethodDefinition
	MemberVar           // (Modifier)* PropertyName TsTypeAnnotation? Initializer?
	Method              // PropertyName TsTypeParameters? Parameters TsTypeAnnotation? Body?
	ModuleSpec
	MultiplicativeExpr // left=Expr right=Expr
	NameIdent          // ReferenceIdent?
	NameSpaceImport    // NameIdent
	NamedImports       // (NamedImport)*
	NamedTupleMember   // TsType
	NewExpr            // expr=Expr Arguments?
	NewTarget
	NoElement
	NotExpr                // Expr
	ObjectLiteral          // (PropertyDefinition)*
	ObjectMethod           // (Modifier)* MethodDefinition
	ObjectPattern          // (PropertyPattern)* BindingRestElement?
	OptionalCallExpr       // expr=Expr Arguments
	OptionalIndexAccess    // expr=Expr index=Expr
	OptionalPropertyAccess // expr=Expr selector=ReferenceIdent
	OptionalTaggedTemplate // tag=Expr literal=TemplateLiteral
	Parameters             // (Parameter)*
	Parenthesized          // Expr
	PostDec                // Expr
	PostInc                // Expr
	PreDec                 // Expr
	PreInc                 // Expr
	Property               // (Modifier)* PropertyName value=Expr
	PropertyAccess         // expr=Expr? selector=ReferenceIdent
	PropertyBinding        // PropertyName ElementPattern
	ReferenceIdent
	Regexp
	RelationalExpr    // left=Expr right=Expr
	RestParameter     // BindingPattern? NameIdent? TsTypeAnnotation?
	ReturnStmt        // Expr?
	Setter            // (Modifier)* PropertyName Parameter Body?
	ShiftExpr         // left=Expr right=Expr
	ShorthandProperty // NameIdent
	SingleNameBinding // NameIdent Initializer?
	SpreadElement     // Expr
	SpreadProperty    // Expr
	Static
	StaticBlock     // Body
	SuperExpr       // ReferenceIdent
	SwitchStmt      // Expr Block
	SyntaxProblem   // ReferenceIdent? Initializer?
	TaggedTemplate  // tag=Expr literal=TemplateLiteral
	TemplateLiteral // template=(NoSubstitutionTemplate | TemplateHead | TemplateMiddle | TemplateTail)+ substitution=(Expr)*
	This
	ThisExpr  // This
	ThrowStmt // Expr
	TryStmt   // Block Catch? Finally?
	TsAbstract
	TsAmbientBinding     // NameIdent TsTypeAnnotation? Initializer?
	TsAmbientClass       // (Modifier)* NameIdent TsTypeParameters? Extends? TsImplementsClause? ClassBody
	TsAmbientEnum        // TsConst? NameIdent TsEnumBody
	TsAmbientExportDecl  // ExportClause
	TsAmbientFunc        // NameIdent TsTypeParameters? Parameters TsTypeAnnotation?
	TsAmbientGlobal      // (ModuleItem)*
	TsAmbientImportAlias // TsImportAliasDecl
	TsAmbientInterface   // (Modifier)* NameIdent TsTypeParameters? TsInterfaceExtends? TsObjectType
	TsAmbientModule      // Literal? (NameIdent)* (TsAmbientElement)*
	TsAmbientNamespace   // (NameIdent)+ (TsAmbientElement)*
	TsAmbientTypeAlias   // TsTypeAliasDecl
	TsAmbientVar         // Const? Let? Var? (TsAmbientBinding)+
	TsArrayType          // TsType
	TsAsConstExpr        // left=Expr TsConst
	TsAsExpr             // left=Expr TsType
	TsAssertsType        // ReferenceIdent? This? TsType?
	TsCallSignature      // TsTypeParameters? Parameters TsTypeAnnotation?
	TsCastExpr           // TsType Expr
	TsConditional        // check=TsType ext=TsType truet=TsType falset=TsType
	TsConst
	TsConstructSignature // (Modifier)* TsTypeParameters? Parameters TsTypeAnnotation?
	TsConstructorType    // TsAbstract? TsTypeParameters? Parameters FnArrow TsType
	TsDeclare
	TsDynamicImport // Arguments
	TsEnum          // TsConst? NameIdent TsEnumBody
	TsEnumBody      // (TsEnumMember)*
	TsEnumMember    // PropertyName Expr?
	TsExclToken
	TsExport
	TsExportAssignment    // Expr
	TsFuncType            // TsTypeParameters? Parameters FnArrow TsType
	TsImplementsClause    // (TsTypeReference)+
	TsImportAliasDecl     // NameIdent ref=(ReferenceIdent)+
	TsImportRequireDecl   // TsExport? NameIdent ModuleSpec
	TsImportType          // TsImportTypeStart TsTypeArguments?
	TsImportTypeStart     // TsImportTypeStart? TsTypeOf? ReferenceIdent? TsType?
	TsIndexMemberDecl     // TsIndexSignature
	TsIndexSignature      // (Modifier)* NameIdent TsType TsTypeAnnotation
	TsIndexedAccessType   // left=TsType index=TsType
	TsInterface           // NameIdent TsTypeParameters? TsInterfaceExtends? TsObjectType
	TsInterfaceExtends    // (TsTypeReference)+
	TsIntersectionType    // inner=(TsType)*
	TsKeyOfType           // TsType
	TsLiteralType         // TsTemplateLiteralType?
	TsMappedType          // NameIdent inType=TsType asType=TsType? TsTypeAnnotation
	TsMethodSignature     // (Modifier)* PropertyName TsTypeParameters? Parameters TsTypeAnnotation?
	TsNamespace           // (NameIdent)+ TsNamespaceBody
	TsNamespaceBody       // (ModuleItem)*
	TsNamespaceExportDecl // NameIdent
	TsNamespaceName       // ref=(ReferenceIdent)+
	TsNonNull             // expr=Expr
	TsNonNullableType     // TsType
	TsNullableType        // TsType
	TsObjectType          // (TsTypeMember)*
	TsOptional
	TsOverride
	TsParenthesizedType // TsType
	TsPredefinedType
	TsPrivate
	TsPropertySignature // (Modifier)* PropertyName TsTypeAnnotation?
	TsProtected
	TsPublic
	TsReadonly
	TsReadonlyType        // TsType
	TsRestType            // TsType
	TsSatisfiesExpr       // left=Expr TsType
	TsTemplateLiteralType // template=(NoSubstitutionTemplate | TemplateHead | TemplateMiddle | TemplateTail)+ substitution=(TsType)*
	TsThisParameter       // TsTypeAnnotation
	TsThisType
	TsTupleType      // (TupleMember)*
	TsTypeAliasDecl  // NameIdent TsTypeParameters? TsType
	TsTypeAnnotation // TsType
	TsTypeArguments  // (TsType)+
	TsTypeConstraint // TsType
	TsTypeName       // TsNamespaceName? ref=(ReferenceIdent)+
	TsTypeOf
	TsTypeOnly
	TsTypeParameter  // (TsVarianceModifier)* NameIdent TsTypeConstraint? TsType?
	TsTypeParameters // (TsTypeParameter)+
	TsTypePredicate  // paramref=ReferenceIdent TsType
	TsTypeQuery      // (ReferenceIdent)+
	TsTypeReference  // TsTypeName TsTypeArguments?
	TsTypeVar        // NameIdent
	TsUnionType      // inner=(TsType)*
	TsUniqueType     // TsType
	TsVarianceModifier
	UnaryExpr // Expr
	Var
	VarDecl   // BindingPattern? NameIdent? TsExclToken? TsTypeAnnotation? Initializer?
	VarStmt   // (VarDecl)+
	WhileStmt // Expr Stmt
	WithStmt  // Expr Stmt
	Yield     // Expr?
	MultiLineComment
	SingleLineComment
	InvalidToken
	NoSubstitutionTemplate
	TemplateHead
	TemplateMiddle
	TemplateTail
	InsertedSemicolon
	NodeTypeMax
)

var nodeTypeStr = [...]string{
	"NONE",
	"Accessor",
	"AdditiveExpr",
	"Arguments",
	"ArrayLiteral",
	"ArrayPattern",
	"ArrowFunc",
	"AssertClause",
	"AssertEntry",
	"AssertionKey",
	"AssignmentExpr",
	"AssignmentOperator",
	"AsyncArrowFunc",
	"AsyncFunc",
	"AsyncFuncExpr",
	"AsyncGeneratorDeclaration",
	"AsyncGeneratorExpression",
	"AsyncGeneratorMethod",
	"AsyncMethod",
	"Await",
	"AwaitExpr",
	"BindingRestElement",
	"BitwiseAND",
	"BitwiseOR",
	"BitwiseXOR",
	"Block",
	"Body",
	"BreakStmt",
	"CallExpr",
	"Case",
	"Catch",
	"Class",
	"ClassBody",
	"ClassExpr",
	"CoalesceExpr",
	"CommaExpr",
	"ComputedPropertyName",
	"ConciseBody",
	"Cond",
	"ConditionalExpr",
	"Const",
	"ContinueStmt",
	"DebuggerStmt",
	"DecoratorCall",
	"DecoratorExpr",
	"Default",
	"DefaultParameter",
	"DoWhileStmt",
	"ElementBinding",
	"EmptyDecl",
	"EmptyStmt",
	"EqualityExpr",
	"ExponentiationExpr",
	"ExportClause",
	"ExportDecl",
	"ExportDefault",
	"ExportSpec",
	"ExprStmt",
	"Extends",
	"File",
	"Finally",
	"FnArrow",
	"ForBinding",
	"ForCondition",
	"ForFinalExpr",
	"ForInStmt",
	"ForInStmtWithVar",
	"ForOfStmt",
	"ForOfStmtWithVar",
	"ForStmt",
	"ForStmtWithVar",
	"Func",
	"FuncExpr",
	"Generator",
	"GeneratorExpr",
	"GeneratorMethod",
	"Getter",
	"IdentExpr",
	"IfStmt",
	"ImportDecl",
	"ImportSpec",
	"InExpr",
	"IndexAccess",
	"Initializer",
	"InstanceOfExpr",
	"JSXAttributeName",
	"JSXClosingElement",
	"JSXElement",
	"JSXElementName",
	"JSXExpr",
	"JSXFragment",
	"JSXLiteral",
	"JSXNormalAttribute",
	"JSXOpeningElement",
	"JSXSelfClosingElement",
	"JSXSpreadAttribute",
	"JSXSpreadExpr",
	"JSXText",
	"LabelIdent",
	"LabelledStmt",
	"Let",
	"LexicalBinding",
	"LexicalDecl",
	"Literal",
	"LiteralPropertyName",
	"LogicalAND",
	"LogicalOR",
	"MemberMethod",
	"MemberVar",
	"Method",
	"ModuleSpec",
	"MultiplicativeExpr",
	"NameIdent",
	"NameSpaceImport",
	"NamedImports",
	"NamedTupleMember",
	"NewExpr",
	"NewTarget",
	"NoElement",
	"NotExpr",
	"ObjectLiteral",
	"ObjectMethod",
	"ObjectPattern",
	"OptionalCallExpr",
	"OptionalIndexAccess",
	"OptionalPropertyAccess",
	"OptionalTaggedTemplate",
	"Parameters",
	"Parenthesized",
	"PostDec",
	"PostInc",
	"PreDec",
	"PreInc",
	"Property",
	"PropertyAccess",
	"PropertyBinding",
	"ReferenceIdent",
	"Regexp",
	"RelationalExpr",
	"RestParameter",
	"ReturnStmt",
	"Setter",
	"ShiftExpr",
	"ShorthandProperty",
	"SingleNameBinding",
	"SpreadElement",
	"SpreadProperty",
	"Static",
	"StaticBlock",
	"SuperExpr",
	"SwitchStmt",
	"SyntaxProblem",
	"TaggedTemplate",
	"TemplateLiteral",
	"This",
	"ThisExpr",
	"ThrowStmt",
	"TryStmt",
	"TsAbstract",
	"TsAmbientBinding",
	"TsAmbientClass",
	"TsAmbientEnum",
	"TsAmbientExportDecl",
	"TsAmbientFunc",
	"TsAmbientGlobal",
	"TsAmbientImportAlias",
	"TsAmbientInterface",
	"TsAmbientModule",
	"TsAmbientNamespace",
	"TsAmbientTypeAlias",
	"TsAmbientVar",
	"TsArrayType",
	"TsAsConstExpr",
	"TsAsExpr",
	"TsAssertsType",
	"TsCallSignature",
	"TsCastExpr",
	"TsConditional",
	"TsConst",
	"TsConstructSignature",
	"TsConstructorType",
	"TsDeclare",
	"TsDynamicImport",
	"TsEnum",
	"TsEnumBody",
	"TsEnumMember",
	"TsExclToken",
	"TsExport",
	"TsExportAssignment",
	"TsFuncType",
	"TsImplementsClause",
	"TsImportAliasDecl",
	"TsImportRequireDecl",
	"TsImportType",
	"TsImportTypeStart",
	"TsIndexMemberDecl",
	"TsIndexSignature",
	"TsIndexedAccessType",
	"TsInterface",
	"TsInterfaceExtends",
	"TsIntersectionType",
	"TsKeyOfType",
	"TsLiteralType",
	"TsMappedType",
	"TsMethodSignature",
	"TsNamespace",
	"TsNamespaceBody",
	"TsNamespaceExportDecl",
	"TsNamespaceName",
	"TsNonNull",
	"TsNonNullableType",
	"TsNullableType",
	"TsObjectType",
	"TsOptional",
	"TsOverride",
	"TsParenthesizedType",
	"TsPredefinedType",
	"TsPrivate",
	"TsPropertySignature",
	"TsProtected",
	"TsPublic",
	"TsReadonly",
	"TsReadonlyType",
	"TsRestType",
	"TsSatisfiesExpr",
	"TsTemplateLiteralType",
	"TsThisParameter",
	"TsThisType",
	"TsTupleType",
	"TsTypeAliasDecl",
	"TsTypeAnnotation",
	"TsTypeArguments",
	"TsTypeConstraint",
	"TsTypeName",
	"TsTypeOf",
	"TsTypeOnly",
	"TsTypeParameter",
	"TsTypeParameters",
	"TsTypePredicate",
	"TsTypeQuery",
	"TsTypeReference",
	"TsTypeVar",
	"TsUnionType",
	"TsUniqueType",
	"TsVarianceModifier",
	"UnaryExpr",
	"Var",
	"VarDecl",
	"VarStmt",
	"WhileStmt",
	"WithStmt",
	"Yield",
	"MultiLineComment",
	"SingleLineComment",
	"InvalidToken",
	"NoSubstitutionTemplate",
	"TemplateHead",
	"TemplateMiddle",
	"TemplateTail",
	"InsertedSemicolon",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}

var BindingPattern = []NodeType{
	ArrayPattern,
	ObjectPattern,
}

var CaseClause = []NodeType{
	Case,
	Default,
}

var ClassElement = []NodeType{
	EmptyDecl,
	MemberMethod,
	MemberVar,
	StaticBlock,
	SyntaxProblem,
	TsIndexMemberDecl,
}

var Decl = []NodeType{
	AsyncFunc,
	AsyncGeneratorDeclaration,
	Class,
	ExportDefault,
	Func,
	Generator,
	ImportDecl,
	LexicalDecl,
	TsAmbientClass,
	TsAmbientEnum,
	TsAmbientExportDecl,
	TsAmbientFunc,
	TsAmbientGlobal,
	TsAmbientImportAlias,
	TsAmbientInterface,
	TsAmbientModule,
	TsAmbientNamespace,
	TsAmbientTypeAlias,
	TsAmbientVar,
	TsEnum,
	TsExportAssignment,
	TsImportAliasDecl,
	TsInterface,
	TsNamespace,
	TsTypeAliasDecl,
}

var Decorator = []NodeType{
	DecoratorCall,
	DecoratorExpr,
}

var ElementPattern = []NodeType{
	ElementBinding,
	NoElement,
	SingleNameBinding,
	SyntaxProblem,
}

var ExportElement = []NodeType{
	ExportSpec,
	SyntaxProblem,
}

var Expr = []NodeType{
	AdditiveExpr,
	ArrayLiteral,
	ArrowFunc,
	AssignmentExpr,
	AsyncArrowFunc,
	AsyncFuncExpr,
	AsyncGeneratorExpression,
	AwaitExpr,
	BitwiseAND,
	BitwiseOR,
	BitwiseXOR,
	CallExpr,
	ClassExpr,
	CoalesceExpr,
	CommaExpr,
	ConditionalExpr,
	EqualityExpr,
	ExponentiationExpr,
	FuncExpr,
	GeneratorExpr,
	IdentExpr,
	InExpr,
	IndexAccess,
	InstanceOfExpr,
	JSXElement,
	JSXFragment,
	Literal,
	LogicalAND,
	LogicalOR,
	MultiplicativeExpr,
	NewExpr,
	NewTarget,
	NoElement,
	NotExpr,
	ObjectLiteral,
	OptionalCallExpr,
	OptionalIndexAccess,
	OptionalPropertyAccess,
	OptionalTaggedTemplate,
	Parenthesized,
	PostDec,
	PostInc,
	PreDec,
	PreInc,
	PropertyAccess,
	Regexp,
	RelationalExpr,
	ShiftExpr,
	SpreadElement,
	SuperExpr,
	SyntaxProblem,
	TaggedTemplate,
	TemplateLiteral,
	ThisExpr,
	TsAsConstExpr,
	TsAsExpr,
	TsCastExpr,
	TsDynamicImport,
	TsNonNull,
	TsSatisfiesExpr,
	UnaryExpr,
	Yield,
}

var IterationStmt = []NodeType{
	DoWhileStmt,
	ForInStmt,
	ForInStmtWithVar,
	ForOfStmt,
	ForOfStmtWithVar,
	ForStmt,
	ForStmtWithVar,
	WhileStmt,
}

var JSXAttribute = []NodeType{
	JSXNormalAttribute,
	JSXSpreadAttribute,
}

var JSXAttributeValue = []NodeType{
	JSXElement,
	JSXExpr,
	JSXFragment,
	JSXLiteral,
}

var JSXChild = []NodeType{
	JSXElement,
	JSXExpr,
	JSXFragment,
	JSXSpreadExpr,
	JSXText,
}

var MethodDefinition = []NodeType{
	AsyncGeneratorMethod,
	AsyncMethod,
	GeneratorMethod,
	Getter,
	Method,
	Setter,
}

var Modifier = []NodeType{
	Accessor,
	DecoratorCall,
	DecoratorExpr,
	Static,
	TsAbstract,
	TsDeclare,
	TsOverride,
	TsPrivate,
	TsProtected,
	TsPublic,
	TsReadonly,
}

var ModuleItem = []NodeType{
	AsyncFunc,
	AsyncGeneratorDeclaration,
	Block,
	BreakStmt,
	Class,
	ContinueStmt,
	DebuggerStmt,
	DoWhileStmt,
	EmptyStmt,
	ExportDecl,
	ExportDefault,
	ExprStmt,
	ForInStmt,
	ForInStmtWithVar,
	ForOfStmt,
	ForOfStmtWithVar,
	ForStmt,
	ForStmtWithVar,
	Func,
	Generator,
	IfStmt,
	ImportDecl,
	LabelledStmt,
	LexicalDecl,
	ReturnStmt,
	SwitchStmt,
	SyntaxProblem,
	ThrowStmt,
	TryStmt,
	TsAmbientClass,
	TsAmbientEnum,
	TsAmbientExportDecl,
	TsAmbientFunc,
	TsAmbientGlobal,
	TsAmbientImportAlias,
	TsAmbientInterface,
	TsAmbientModule,
	TsAmbientNamespace,
	TsAmbientTypeAlias,
	TsAmbientVar,
	TsEnum,
	TsExportAssignment,
	TsImportAliasDecl,
	TsImportRequireDecl,
	TsInterface,
	TsNamespace,
	TsNamespaceExportDecl,
	TsTypeAliasDecl,
	VarStmt,
	WhileStmt,
	WithStmt,
}

var NamedImport = []NodeType{
	ImportSpec,
	SyntaxProblem,
}

var Parameter = []NodeType{
	DefaultParameter,
	RestParameter,
	SyntaxProblem,
	TsThisParameter,
}

var PropertyDefinition = []NodeType{
	ObjectMethod,
	Property,
	ShorthandProperty,
	SpreadProperty,
	SyntaxProblem,
}

var PropertyName = []NodeType{
	ComputedPropertyName,
	LiteralPropertyName,
}

var PropertyPattern = []NodeType{
	PropertyBinding,
	SingleNameBinding,
	SyntaxProblem,
}

var Stmt = []NodeType{
	Block,
	BreakStmt,
	ContinueStmt,
	DebuggerStmt,
	DoWhileStmt,
	EmptyStmt,
	ExprStmt,
	ForInStmt,
	ForInStmtWithVar,
	ForOfStmt,
	ForOfStmtWithVar,
	ForStmt,
	ForStmtWithVar,
	IfStmt,
	LabelledStmt,
	ReturnStmt,
	SwitchStmt,
	ThrowStmt,
	TryStmt,
	VarStmt,
	WhileStmt,
	WithStmt,
}

var StmtListItem = []NodeType{
	AsyncFunc,
	AsyncGeneratorDeclaration,
	Block,
	BreakStmt,
	Class,
	ContinueStmt,
	DebuggerStmt,
	DoWhileStmt,
	EmptyStmt,
	ExportDefault,
	ExprStmt,
	ForInStmt,
	ForInStmtWithVar,
	ForOfStmt,
	ForOfStmtWithVar,
	ForStmt,
	ForStmtWithVar,
	Func,
	Generator,
	IfStmt,
	ImportDecl,
	LabelledStmt,
	LexicalDecl,
	ReturnStmt,
	SwitchStmt,
	SyntaxProblem,
	ThrowStmt,
	TryStmt,
	TsAmbientClass,
	TsAmbientEnum,
	TsAmbientExportDecl,
	TsAmbientFunc,
	TsAmbientGlobal,
	TsAmbientImportAlias,
	TsAmbientInterface,
	TsAmbientModule,
	TsAmbientNamespace,
	TsAmbientTypeAlias,
	TsAmbientVar,
	TsEnum,
	TsExportAssignment,
	TsImportAliasDecl,
	TsInterface,
	TsNamespace,
	TsTypeAliasDecl,
	VarStmt,
	WhileStmt,
	WithStmt,
}

var TokenSet = []NodeType{
	NoSubstitutionTemplate,
	TemplateHead,
	TemplateMiddle,
	TemplateTail,
}

var TsAccessibilityModifier = []NodeType{
	TsPrivate,
	TsProtected,
	TsPublic,
}

var TsAmbientElement = []NodeType{
	ExportDefault,
	ImportDecl,
	TsAmbientClass,
	TsAmbientEnum,
	TsAmbientExportDecl,
	TsAmbientFunc,
	TsAmbientGlobal,
	TsAmbientImportAlias,
	TsAmbientInterface,
	TsAmbientModule,
	TsAmbientNamespace,
	TsAmbientTypeAlias,
	TsAmbientVar,
	TsExportAssignment,
}

var TsType = []NodeType{
	TsArrayType,
	TsAssertsType,
	TsConditional,
	TsConstructorType,
	TsFuncType,
	TsImportType,
	TsIndexedAccessType,
	TsIntersectionType,
	TsKeyOfType,
	TsLiteralType,
	TsMappedType,
	TsNonNullableType,
	TsNullableType,
	TsObjectType,
	TsParenthesizedType,
	TsPredefinedType,
	TsReadonlyType,
	TsRestType,
	TsThisType,
	TsTupleType,
	TsTypePredicate,
	TsTypeQuery,
	TsTypeReference,
	TsTypeVar,
	TsUnionType,
	TsUniqueType,
}

var TsTypeMember = []NodeType{
	Getter,
	Setter,
	SyntaxProblem,
	TsCallSignature,
	TsConstructSignature,
	TsIndexSignature,
	TsMethodSignature,
	TsPropertySignature,
}

var TupleMember = []NodeType{
	NamedTupleMember,
	TsArrayType,
	TsAssertsType,
	TsConditional,
	TsConstructorType,
	TsFuncType,
	TsImportType,
	TsIndexedAccessType,
	TsIntersectionType,
	TsKeyOfType,
	TsLiteralType,
	TsMappedType,
	TsNonNullableType,
	TsNullableType,
	TsObjectType,
	TsParenthesizedType,
	TsPredefinedType,
	TsReadonlyType,
	TsRestType,
	TsThisType,
	TsTupleType,
	TsTypePredicate,
	TsTypeQuery,
	TsTypeReference,
	TsTypeVar,
	TsUnionType,
	TsUniqueType,
}

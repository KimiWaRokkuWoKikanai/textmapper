package js

import "fmt"

type NodeType int

type Listener interface {
	Node(t NodeType, offset, endoffset int)
}

const (
	IdentifierName NodeType = iota + 1
	IdentifierReference
	BindingIdentifier
	LabelIdentifier
	ThisExpression
	RegularExpression
	ParenthesizedExpression
	Literal
	ArrayLiteral
	SpreadElement
	ObjectLiteral
	PropertyDefinition
	SyntaxError
	LiteralPropertyName
	ComputedPropertyName
	CoverInitializedName
	Initializer
	TemplateLiteral
	IndexAccess
	PropertyAccess
	TaggedTemplate
	NewExpression
	SuperExpression
	NewTarget
	CallExpression
	Arguments
	PostIncrementExpression
	PostDecrementExpression
	PreIncrementExpression
	PreDecrementExpression
	UnaryExpression
	AdditiveExpression
	ShiftExpression
	MultiplicativeExpression
	ExponentiationExpression
	RelationalExpression
	EqualityExpression
	BitwiseANDExpression
	BitwiseXORExpression
	BitwiseORExpression
	LogicalANDExpression
	LogicalORExpression
	ConditionalExpression
	AssignmentExpression
	AssignmentOperator
	Block
	LexicalDeclaration
	LexicalBinding
	VariableStatement
	VariableDeclaration
	ObjectBindingPattern
	ArrayBindingPattern
	BindingElisionElement
	BindingProperty
	BindingElement
	SingleNameBinding
	BindingRestElement
	EmptyStatement
	ExpressionStatement
	IfStatement
	DoWhileStatement
	WhileStatement
	ForStatement
	ForInStatement
	ForOfStatement
	ForBinding
	ContinueStatement
	BreakStatement
	ReturnStatement
	WithStatement
	SwitchStatement
	CaseBlock
	CaseClause
	DefaultClause
	LabelledStatement
	ThrowStatement
	TryStatement
	Catch
	Finally
	CatchParameter
	DebuggerStatement
	FunctionDeclaration
	FunctionExpression
	FormalParameters
	FunctionRestParameter
	FormalParameter
	ArrowFunction
	ArrowParameters
	ConciseBody
	MethodDefinition
	GeneratorMethod
	GeneratorDeclaration
	GeneratorExpression
	YieldExpression
	ClassDeclaration
	ClassExpression
	ClassHeritage
	ClassBody
	ClassElement
	Module
	ModuleItem
	ImportDeclaration
	NameSpaceImport
	NamedImports
	ImportSpecifier
	ModuleSpecifier
	ExportDeclaration
	ExportDefault
	ExportClause
	ExportSpecifier
	JSXElement
	JSXSelfClosingElement
	JSXOpeningElement
	JSXClosingElement
	JSXElementName
	JSXAttribute
	JSXSpreadAttribute
	JSXAttributeName
	JSXAttributeValue
	JSXText
	JSXChild
	InsertedSemicolon
	NodeTypeMax
)

var ruleNodeType = [...]NodeType{
	IdentifierName, // IdentifierName ::= Identifier
	IdentifierName, // IdentifierName ::= 'break'
	IdentifierName, // IdentifierName ::= 'do'
	IdentifierName, // IdentifierName ::= 'in'
	IdentifierName, // IdentifierName ::= 'typeof'
	IdentifierName, // IdentifierName ::= 'case'
	IdentifierName, // IdentifierName ::= 'else'
	IdentifierName, // IdentifierName ::= 'instanceof'
	IdentifierName, // IdentifierName ::= 'var'
	IdentifierName, // IdentifierName ::= 'catch'
	IdentifierName, // IdentifierName ::= 'export'
	IdentifierName, // IdentifierName ::= 'new'
	IdentifierName, // IdentifierName ::= 'void'
	IdentifierName, // IdentifierName ::= 'class'
	IdentifierName, // IdentifierName ::= 'extends'
	IdentifierName, // IdentifierName ::= 'return'
	IdentifierName, // IdentifierName ::= 'while'
	IdentifierName, // IdentifierName ::= 'const'
	IdentifierName, // IdentifierName ::= 'finally'
	IdentifierName, // IdentifierName ::= 'super'
	IdentifierName, // IdentifierName ::= 'with'
	IdentifierName, // IdentifierName ::= 'continue'
	IdentifierName, // IdentifierName ::= 'for'
	IdentifierName, // IdentifierName ::= 'switch'
	IdentifierName, // IdentifierName ::= 'yield'
	IdentifierName, // IdentifierName ::= 'debugger'
	IdentifierName, // IdentifierName ::= 'function'
	IdentifierName, // IdentifierName ::= 'this'
	IdentifierName, // IdentifierName ::= 'default'
	IdentifierName, // IdentifierName ::= 'if'
	IdentifierName, // IdentifierName ::= 'throw'
	IdentifierName, // IdentifierName ::= 'delete'
	IdentifierName, // IdentifierName ::= 'import'
	IdentifierName, // IdentifierName ::= 'try'
	IdentifierName, // IdentifierName ::= 'enum'
	IdentifierName, // IdentifierName ::= 'await'
	IdentifierName, // IdentifierName ::= 'null'
	IdentifierName, // IdentifierName ::= 'true'
	IdentifierName, // IdentifierName ::= 'false'
	IdentifierName, // IdentifierName ::= 'as'
	IdentifierName, // IdentifierName ::= 'from'
	IdentifierName, // IdentifierName ::= 'get'
	IdentifierName, // IdentifierName ::= 'let'
	IdentifierName, // IdentifierName ::= 'of'
	IdentifierName, // IdentifierName ::= 'set'
	IdentifierName, // IdentifierName ::= 'static'
	IdentifierName, // IdentifierName ::= 'target'
	IdentifierReference, // IdentifierReference ::= Identifier
	IdentifierReference, // IdentifierReference ::= 'yield'
	IdentifierReference, // IdentifierReference ::= 'let'
	IdentifierReference, // IdentifierReference ::= 'as'
	IdentifierReference, // IdentifierReference ::= 'from'
	IdentifierReference, // IdentifierReference ::= 'get'
	IdentifierReference, // IdentifierReference ::= 'of'
	IdentifierReference, // IdentifierReference ::= 'set'
	IdentifierReference, // IdentifierReference ::= 'static'
	IdentifierReference, // IdentifierReference ::= 'target'
	IdentifierReference, // IdentifierReference_NoLet ::= Identifier
	IdentifierReference, // IdentifierReference_NoLet ::= 'yield'
	IdentifierReference, // IdentifierReference_NoLet ::= 'as'
	IdentifierReference, // IdentifierReference_NoLet ::= 'from'
	IdentifierReference, // IdentifierReference_NoLet ::= 'get'
	IdentifierReference, // IdentifierReference_NoLet ::= 'of'
	IdentifierReference, // IdentifierReference_NoLet ::= 'set'
	IdentifierReference, // IdentifierReference_NoLet ::= 'static'
	IdentifierReference, // IdentifierReference_NoLet ::= 'target'
	IdentifierReference, // IdentifierReference_NoLet_Yield ::= Identifier
	IdentifierReference, // IdentifierReference_NoLet_Yield ::= 'as'
	IdentifierReference, // IdentifierReference_NoLet_Yield ::= 'from'
	IdentifierReference, // IdentifierReference_NoLet_Yield ::= 'get'
	IdentifierReference, // IdentifierReference_NoLet_Yield ::= 'of'
	IdentifierReference, // IdentifierReference_NoLet_Yield ::= 'set'
	IdentifierReference, // IdentifierReference_NoLet_Yield ::= 'static'
	IdentifierReference, // IdentifierReference_NoLet_Yield ::= 'target'
	IdentifierReference, // IdentifierReference_Yield ::= Identifier
	IdentifierReference, // IdentifierReference_Yield ::= 'let'
	IdentifierReference, // IdentifierReference_Yield ::= 'as'
	IdentifierReference, // IdentifierReference_Yield ::= 'from'
	IdentifierReference, // IdentifierReference_Yield ::= 'get'
	IdentifierReference, // IdentifierReference_Yield ::= 'of'
	IdentifierReference, // IdentifierReference_Yield ::= 'set'
	IdentifierReference, // IdentifierReference_Yield ::= 'static'
	IdentifierReference, // IdentifierReference_Yield ::= 'target'
	BindingIdentifier, // BindingIdentifier ::= Identifier
	BindingIdentifier, // BindingIdentifier ::= 'yield'
	BindingIdentifier, // BindingIdentifier ::= 'as'
	BindingIdentifier, // BindingIdentifier ::= 'from'
	BindingIdentifier, // BindingIdentifier ::= 'get'
	BindingIdentifier, // BindingIdentifier ::= 'let'
	BindingIdentifier, // BindingIdentifier ::= 'of'
	BindingIdentifier, // BindingIdentifier ::= 'set'
	BindingIdentifier, // BindingIdentifier ::= 'static'
	BindingIdentifier, // BindingIdentifier ::= 'target'
	BindingIdentifier, // BindingIdentifier_Yield ::= Identifier
	BindingIdentifier, // BindingIdentifier_Yield ::= 'as'
	BindingIdentifier, // BindingIdentifier_Yield ::= 'from'
	BindingIdentifier, // BindingIdentifier_Yield ::= 'get'
	BindingIdentifier, // BindingIdentifier_Yield ::= 'let'
	BindingIdentifier, // BindingIdentifier_Yield ::= 'of'
	BindingIdentifier, // BindingIdentifier_Yield ::= 'set'
	BindingIdentifier, // BindingIdentifier_Yield ::= 'static'
	BindingIdentifier, // BindingIdentifier_Yield ::= 'target'
	LabelIdentifier, // LabelIdentifier ::= Identifier
	LabelIdentifier, // LabelIdentifier ::= 'yield'
	LabelIdentifier, // LabelIdentifier ::= 'as'
	LabelIdentifier, // LabelIdentifier ::= 'from'
	LabelIdentifier, // LabelIdentifier ::= 'get'
	LabelIdentifier, // LabelIdentifier ::= 'let'
	LabelIdentifier, // LabelIdentifier ::= 'of'
	LabelIdentifier, // LabelIdentifier ::= 'set'
	LabelIdentifier, // LabelIdentifier ::= 'static'
	LabelIdentifier, // LabelIdentifier ::= 'target'
	LabelIdentifier, // LabelIdentifier_Yield ::= Identifier
	LabelIdentifier, // LabelIdentifier_Yield ::= 'as'
	LabelIdentifier, // LabelIdentifier_Yield ::= 'from'
	LabelIdentifier, // LabelIdentifier_Yield ::= 'get'
	LabelIdentifier, // LabelIdentifier_Yield ::= 'let'
	LabelIdentifier, // LabelIdentifier_Yield ::= 'of'
	LabelIdentifier, // LabelIdentifier_Yield ::= 'set'
	LabelIdentifier, // LabelIdentifier_Yield ::= 'static'
	LabelIdentifier, // LabelIdentifier_Yield ::= 'target'
	ThisExpression, // PrimaryExpression ::= 'this'
	0, // PrimaryExpression ::= IdentifierReference
	0, // PrimaryExpression ::= Literal
	0, // PrimaryExpression ::= ArrayLiteral
	0, // PrimaryExpression ::= ObjectLiteral
	0, // PrimaryExpression ::= FunctionExpression
	0, // PrimaryExpression ::= ClassExpression
	0, // PrimaryExpression ::= GeneratorExpression
	RegularExpression, // PrimaryExpression ::= RegularExpressionLiteral
	0, // PrimaryExpression ::= TemplateLiteral
	ParenthesizedExpression, // PrimaryExpression ::= CoverParenthesizedExpressionAndArrowParameterList
	0, // PrimaryExpression ::= JSXElement
	ThisExpression, // PrimaryExpression_NoFuncClass ::= 'this'
	0, // PrimaryExpression_NoFuncClass ::= IdentifierReference
	0, // PrimaryExpression_NoFuncClass ::= Literal
	0, // PrimaryExpression_NoFuncClass ::= ArrayLiteral
	0, // PrimaryExpression_NoFuncClass ::= ObjectLiteral
	RegularExpression, // PrimaryExpression_NoFuncClass ::= RegularExpressionLiteral
	0, // PrimaryExpression_NoFuncClass ::= TemplateLiteral
	ParenthesizedExpression, // PrimaryExpression_NoFuncClass ::= CoverParenthesizedExpressionAndArrowParameterList
	0, // PrimaryExpression_NoFuncClass ::= JSXElement
	ThisExpression, // PrimaryExpression_NoFuncClass_NoLet ::= 'this'
	0, // PrimaryExpression_NoFuncClass_NoLet ::= IdentifierReference_NoLet
	0, // PrimaryExpression_NoFuncClass_NoLet ::= Literal
	0, // PrimaryExpression_NoFuncClass_NoLet ::= ArrayLiteral
	0, // PrimaryExpression_NoFuncClass_NoLet ::= ObjectLiteral
	RegularExpression, // PrimaryExpression_NoFuncClass_NoLet ::= RegularExpressionLiteral
	0, // PrimaryExpression_NoFuncClass_NoLet ::= TemplateLiteral
	ParenthesizedExpression, // PrimaryExpression_NoFuncClass_NoLet ::= CoverParenthesizedExpressionAndArrowParameterList
	0, // PrimaryExpression_NoFuncClass_NoLet ::= JSXElement
	ThisExpression, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral ::= 'this'
	0, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral ::= IdentifierReference_NoLet
	0, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral ::= Literal
	0, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral ::= ArrayLiteral
	RegularExpression, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral ::= RegularExpressionLiteral
	0, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral ::= TemplateLiteral
	ParenthesizedExpression, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral ::= CoverParenthesizedExpressionAndArrowParameterList
	0, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral ::= JSXElement
	ThisExpression, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral_Yield ::= 'this'
	0, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral_Yield ::= IdentifierReference_NoLet_Yield
	0, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral_Yield ::= Literal
	0, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral_Yield ::= ArrayLiteral_Yield
	RegularExpression, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral_Yield ::= RegularExpressionLiteral
	0, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral_Yield ::= TemplateLiteral_Yield
	ParenthesizedExpression, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral_Yield ::= CoverParenthesizedExpressionAndArrowParameterList_Yield
	0, // PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral_Yield ::= JSXElement_Yield
	ThisExpression, // PrimaryExpression_NoFuncClass_NoObjLiteral ::= 'this'
	0, // PrimaryExpression_NoFuncClass_NoObjLiteral ::= IdentifierReference
	0, // PrimaryExpression_NoFuncClass_NoObjLiteral ::= Literal
	0, // PrimaryExpression_NoFuncClass_NoObjLiteral ::= ArrayLiteral
	RegularExpression, // PrimaryExpression_NoFuncClass_NoObjLiteral ::= RegularExpressionLiteral
	0, // PrimaryExpression_NoFuncClass_NoObjLiteral ::= TemplateLiteral
	ParenthesizedExpression, // PrimaryExpression_NoFuncClass_NoObjLiteral ::= CoverParenthesizedExpressionAndArrowParameterList
	0, // PrimaryExpression_NoFuncClass_NoObjLiteral ::= JSXElement
	ThisExpression, // PrimaryExpression_NoFuncClass_NoObjLiteral_Yield ::= 'this'
	0, // PrimaryExpression_NoFuncClass_NoObjLiteral_Yield ::= IdentifierReference_Yield
	0, // PrimaryExpression_NoFuncClass_NoObjLiteral_Yield ::= Literal
	0, // PrimaryExpression_NoFuncClass_NoObjLiteral_Yield ::= ArrayLiteral_Yield
	RegularExpression, // PrimaryExpression_NoFuncClass_NoObjLiteral_Yield ::= RegularExpressionLiteral
	0, // PrimaryExpression_NoFuncClass_NoObjLiteral_Yield ::= TemplateLiteral_Yield
	ParenthesizedExpression, // PrimaryExpression_NoFuncClass_NoObjLiteral_Yield ::= CoverParenthesizedExpressionAndArrowParameterList_Yield
	0, // PrimaryExpression_NoFuncClass_NoObjLiteral_Yield ::= JSXElement_Yield
	ThisExpression, // PrimaryExpression_NoLet ::= 'this'
	0, // PrimaryExpression_NoLet ::= IdentifierReference_NoLet
	0, // PrimaryExpression_NoLet ::= Literal
	0, // PrimaryExpression_NoLet ::= ArrayLiteral
	0, // PrimaryExpression_NoLet ::= ObjectLiteral
	0, // PrimaryExpression_NoLet ::= FunctionExpression
	0, // PrimaryExpression_NoLet ::= ClassExpression
	0, // PrimaryExpression_NoLet ::= GeneratorExpression
	RegularExpression, // PrimaryExpression_NoLet ::= RegularExpressionLiteral
	0, // PrimaryExpression_NoLet ::= TemplateLiteral
	ParenthesizedExpression, // PrimaryExpression_NoLet ::= CoverParenthesizedExpressionAndArrowParameterList
	0, // PrimaryExpression_NoLet ::= JSXElement
	ThisExpression, // PrimaryExpression_NoLet_NoObjLiteral ::= 'this'
	0, // PrimaryExpression_NoLet_NoObjLiteral ::= IdentifierReference_NoLet
	0, // PrimaryExpression_NoLet_NoObjLiteral ::= Literal
	0, // PrimaryExpression_NoLet_NoObjLiteral ::= ArrayLiteral
	0, // PrimaryExpression_NoLet_NoObjLiteral ::= FunctionExpression
	0, // PrimaryExpression_NoLet_NoObjLiteral ::= ClassExpression
	0, // PrimaryExpression_NoLet_NoObjLiteral ::= GeneratorExpression
	RegularExpression, // PrimaryExpression_NoLet_NoObjLiteral ::= RegularExpressionLiteral
	0, // PrimaryExpression_NoLet_NoObjLiteral ::= TemplateLiteral
	ParenthesizedExpression, // PrimaryExpression_NoLet_NoObjLiteral ::= CoverParenthesizedExpressionAndArrowParameterList
	0, // PrimaryExpression_NoLet_NoObjLiteral ::= JSXElement
	ThisExpression, // PrimaryExpression_NoLet_Yield ::= 'this'
	0, // PrimaryExpression_NoLet_Yield ::= IdentifierReference_NoLet_Yield
	0, // PrimaryExpression_NoLet_Yield ::= Literal
	0, // PrimaryExpression_NoLet_Yield ::= ArrayLiteral_Yield
	0, // PrimaryExpression_NoLet_Yield ::= ObjectLiteral_Yield
	0, // PrimaryExpression_NoLet_Yield ::= FunctionExpression
	0, // PrimaryExpression_NoLet_Yield ::= ClassExpression_Yield
	0, // PrimaryExpression_NoLet_Yield ::= GeneratorExpression
	RegularExpression, // PrimaryExpression_NoLet_Yield ::= RegularExpressionLiteral
	0, // PrimaryExpression_NoLet_Yield ::= TemplateLiteral_Yield
	ParenthesizedExpression, // PrimaryExpression_NoLet_Yield ::= CoverParenthesizedExpressionAndArrowParameterList_Yield
	0, // PrimaryExpression_NoLet_Yield ::= JSXElement_Yield
	ThisExpression, // PrimaryExpression_NoObjLiteral ::= 'this'
	0, // PrimaryExpression_NoObjLiteral ::= IdentifierReference
	0, // PrimaryExpression_NoObjLiteral ::= Literal
	0, // PrimaryExpression_NoObjLiteral ::= ArrayLiteral
	0, // PrimaryExpression_NoObjLiteral ::= FunctionExpression
	0, // PrimaryExpression_NoObjLiteral ::= ClassExpression
	0, // PrimaryExpression_NoObjLiteral ::= GeneratorExpression
	RegularExpression, // PrimaryExpression_NoObjLiteral ::= RegularExpressionLiteral
	0, // PrimaryExpression_NoObjLiteral ::= TemplateLiteral
	ParenthesizedExpression, // PrimaryExpression_NoObjLiteral ::= CoverParenthesizedExpressionAndArrowParameterList
	0, // PrimaryExpression_NoObjLiteral ::= JSXElement
	ThisExpression, // PrimaryExpression_Yield ::= 'this'
	0, // PrimaryExpression_Yield ::= IdentifierReference_Yield
	0, // PrimaryExpression_Yield ::= Literal
	0, // PrimaryExpression_Yield ::= ArrayLiteral_Yield
	0, // PrimaryExpression_Yield ::= ObjectLiteral_Yield
	0, // PrimaryExpression_Yield ::= FunctionExpression
	0, // PrimaryExpression_Yield ::= ClassExpression_Yield
	0, // PrimaryExpression_Yield ::= GeneratorExpression
	RegularExpression, // PrimaryExpression_Yield ::= RegularExpressionLiteral
	0, // PrimaryExpression_Yield ::= TemplateLiteral_Yield
	ParenthesizedExpression, // PrimaryExpression_Yield ::= CoverParenthesizedExpressionAndArrowParameterList_Yield
	0, // PrimaryExpression_Yield ::= JSXElement_Yield
	0, // CoverParenthesizedExpressionAndArrowParameterList ::= '(' Expression_In ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList ::= '(' ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList ::= '(' '...' BindingIdentifier ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList ::= '(' '...' BindingPattern ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList ::= '(' Expression_In ',' '...' BindingIdentifier ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList ::= '(' Expression_In ',' '...' BindingPattern ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList_Yield ::= '(' Expression_In_Yield ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList_Yield ::= '(' ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList_Yield ::= '(' '...' BindingIdentifier_Yield ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList_Yield ::= '(' '...' BindingPattern_Yield ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList_Yield ::= '(' Expression_In_Yield ',' '...' BindingIdentifier_Yield ')'
	0, // CoverParenthesizedExpressionAndArrowParameterList_Yield ::= '(' Expression_In_Yield ',' '...' BindingPattern_Yield ')'
	Literal, // Literal ::= 'null'
	Literal, // Literal ::= 'true'
	Literal, // Literal ::= 'false'
	Literal, // Literal ::= NumericLiteral
	Literal, // Literal ::= StringLiteral
	ArrayLiteral, // ArrayLiteral ::= '[' Elisionopt ']'
	ArrayLiteral, // ArrayLiteral ::= '[' ElementList ']'
	ArrayLiteral, // ArrayLiteral ::= '[' ElementList ',' Elisionopt ']'
	ArrayLiteral, // ArrayLiteral_Yield ::= '[' Elisionopt ']'
	ArrayLiteral, // ArrayLiteral_Yield ::= '[' ElementList_Yield ']'
	ArrayLiteral, // ArrayLiteral_Yield ::= '[' ElementList_Yield ',' Elisionopt ']'
	0, // ElementList ::= Elisionopt AssignmentExpression_In
	0, // ElementList ::= Elisionopt SpreadElement
	0, // ElementList ::= ElementList ',' Elisionopt AssignmentExpression_In
	0, // ElementList ::= ElementList ',' Elisionopt SpreadElement
	0, // ElementList_Yield ::= Elisionopt AssignmentExpression_In_Yield
	0, // ElementList_Yield ::= Elisionopt SpreadElement_Yield
	0, // ElementList_Yield ::= ElementList_Yield ',' Elisionopt AssignmentExpression_In_Yield
	0, // ElementList_Yield ::= ElementList_Yield ',' Elisionopt SpreadElement_Yield
	0, // Elision ::= ','
	0, // Elision ::= Elision ','
	SpreadElement, // SpreadElement ::= '...' AssignmentExpression_In
	SpreadElement, // SpreadElement_Yield ::= '...' AssignmentExpression_In_Yield
	ObjectLiteral, // ObjectLiteral ::= '{' '}'
	ObjectLiteral, // ObjectLiteral ::= '{' PropertyDefinitionList '}'
	ObjectLiteral, // ObjectLiteral ::= '{' PropertyDefinitionList ',' '}'
	ObjectLiteral, // ObjectLiteral_Yield ::= '{' '}'
	ObjectLiteral, // ObjectLiteral_Yield ::= '{' PropertyDefinitionList_Yield '}'
	ObjectLiteral, // ObjectLiteral_Yield ::= '{' PropertyDefinitionList_Yield ',' '}'
	0, // PropertyDefinitionList ::= PropertyDefinition
	0, // PropertyDefinitionList ::= PropertyDefinitionList ',' PropertyDefinition
	0, // PropertyDefinitionList_Yield ::= PropertyDefinition_Yield
	0, // PropertyDefinitionList_Yield ::= PropertyDefinitionList_Yield ',' PropertyDefinition_Yield
	PropertyDefinition, // PropertyDefinition ::= IdentifierReference
	SyntaxError, // PropertyDefinition ::= CoverInitializedName
	PropertyDefinition, // PropertyDefinition ::= PropertyName ':' AssignmentExpression_In
	0, // PropertyDefinition ::= MethodDefinition
	PropertyDefinition, // PropertyDefinition_Yield ::= IdentifierReference_Yield
	SyntaxError, // PropertyDefinition_Yield ::= CoverInitializedName_Yield
	PropertyDefinition, // PropertyDefinition_Yield ::= PropertyName_Yield ':' AssignmentExpression_In_Yield
	0, // PropertyDefinition_Yield ::= MethodDefinition_Yield
	0, // PropertyName ::= LiteralPropertyName
	0, // PropertyName ::= ComputedPropertyName
	0, // PropertyName_Yield ::= LiteralPropertyName
	0, // PropertyName_Yield ::= ComputedPropertyName_Yield
	LiteralPropertyName, // LiteralPropertyName ::= IdentifierName
	LiteralPropertyName, // LiteralPropertyName ::= StringLiteral
	LiteralPropertyName, // LiteralPropertyName ::= NumericLiteral
	ComputedPropertyName, // ComputedPropertyName ::= '[' AssignmentExpression_In ']'
	ComputedPropertyName, // ComputedPropertyName_Yield ::= '[' AssignmentExpression_In_Yield ']'
	CoverInitializedName, // CoverInitializedName ::= IdentifierReference Initializer_In
	CoverInitializedName, // CoverInitializedName_Yield ::= IdentifierReference_Yield Initializer_In_Yield
	Initializer, // Initializer ::= '=' AssignmentExpression
	Initializer, // Initializer_In ::= '=' AssignmentExpression_In
	Initializer, // Initializer_In_Yield ::= '=' AssignmentExpression_In_Yield
	Initializer, // Initializer_Yield ::= '=' AssignmentExpression_Yield
	TemplateLiteral, // TemplateLiteral ::= NoSubstitutionTemplate
	TemplateLiteral, // TemplateLiteral ::= TemplateHead Expression_In TemplateSpans
	TemplateLiteral, // TemplateLiteral_Yield ::= NoSubstitutionTemplate
	TemplateLiteral, // TemplateLiteral_Yield ::= TemplateHead Expression_In_Yield TemplateSpans_Yield
	0, // TemplateSpans ::= TemplateTail
	0, // TemplateSpans ::= TemplateMiddleList TemplateTail
	0, // TemplateSpans_Yield ::= TemplateTail
	0, // TemplateSpans_Yield ::= TemplateMiddleList_Yield TemplateTail
	0, // TemplateMiddleList ::= TemplateMiddle Expression_In
	0, // TemplateMiddleList ::= TemplateMiddleList TemplateMiddle Expression_In
	0, // TemplateMiddleList_Yield ::= TemplateMiddle Expression_In_Yield
	0, // TemplateMiddleList_Yield ::= TemplateMiddleList_Yield TemplateMiddle Expression_In_Yield
	0, // MemberExpression ::= PrimaryExpression
	IndexAccess, // MemberExpression ::= MemberExpression '[' Expression_In ']'
	PropertyAccess, // MemberExpression ::= MemberExpression '.' IdentifierName
	TaggedTemplate, // MemberExpression ::= MemberExpression TemplateLiteral
	0, // MemberExpression ::= SuperProperty
	0, // MemberExpression ::= MetaProperty
	NewExpression, // MemberExpression ::= 'new' MemberExpression Arguments
	0, // MemberExpression_NoFuncClass ::= PrimaryExpression_NoFuncClass
	IndexAccess, // MemberExpression_NoFuncClass ::= MemberExpression_NoFuncClass '[' Expression_In ']'
	PropertyAccess, // MemberExpression_NoFuncClass ::= MemberExpression_NoFuncClass '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoFuncClass ::= MemberExpression_NoFuncClass TemplateLiteral
	0, // MemberExpression_NoFuncClass ::= SuperProperty
	0, // MemberExpression_NoFuncClass ::= MetaProperty
	NewExpression, // MemberExpression_NoFuncClass ::= 'new' MemberExpression Arguments
	0, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= PrimaryExpression_NoFuncClass_NoObjLiteral
	IndexAccess, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral '[' Expression_In ']'
	PropertyAccess, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral TemplateLiteral
	0, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= SuperProperty
	0, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= MetaProperty
	NewExpression, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= 'new' MemberExpression Arguments
	0, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= PrimaryExpression_NoFuncClass_NoObjLiteral_Yield
	IndexAccess, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield TemplateLiteral_Yield
	0, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= SuperProperty_Yield
	0, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= MetaProperty
	NewExpression, // MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= 'new' MemberExpression_Yield Arguments_Yield
	0, // MemberExpression_NoLet ::= PrimaryExpression_NoLet
	IndexAccess, // MemberExpression_NoLet ::= MemberExpression_NoLet '[' Expression_In ']'
	PropertyAccess, // MemberExpression_NoLet ::= MemberExpression_NoLet '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLet ::= MemberExpression_NoLet TemplateLiteral
	0, // MemberExpression_NoLet ::= SuperProperty
	0, // MemberExpression_NoLet ::= MetaProperty
	NewExpression, // MemberExpression_NoLet ::= 'new' MemberExpression Arguments
	0, // MemberExpression_NoLet_Yield ::= PrimaryExpression_NoLet_Yield
	IndexAccess, // MemberExpression_NoLet_Yield ::= MemberExpression_NoLet_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // MemberExpression_NoLet_Yield ::= MemberExpression_NoLet_Yield '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLet_Yield ::= MemberExpression_NoLet_Yield TemplateLiteral_Yield
	0, // MemberExpression_NoLet_Yield ::= SuperProperty_Yield
	0, // MemberExpression_NoLet_Yield ::= MetaProperty
	NewExpression, // MemberExpression_NoLet_Yield ::= 'new' MemberExpression_Yield Arguments_Yield
	0, // MemberExpression_NoLetOnly ::= PrimaryExpression_NoLet
	IndexAccess, // MemberExpression_NoLetOnly ::= MemberExpression '[' Expression_In ']'
	PropertyAccess, // MemberExpression_NoLetOnly ::= MemberExpression '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLetOnly ::= MemberExpression TemplateLiteral
	0, // MemberExpression_NoLetOnly ::= SuperProperty
	0, // MemberExpression_NoLetOnly ::= MetaProperty
	NewExpression, // MemberExpression_NoLetOnly ::= 'new' MemberExpression Arguments
	0, // MemberExpression_NoLetOnly_NoFuncClass ::= PrimaryExpression_NoFuncClass_NoLet
	IndexAccess, // MemberExpression_NoLetOnly_NoFuncClass ::= MemberExpression_NoFuncClass '[' Expression_In ']'
	PropertyAccess, // MemberExpression_NoLetOnly_NoFuncClass ::= MemberExpression_NoFuncClass '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLetOnly_NoFuncClass ::= MemberExpression_NoFuncClass TemplateLiteral
	0, // MemberExpression_NoLetOnly_NoFuncClass ::= SuperProperty
	0, // MemberExpression_NoLetOnly_NoFuncClass ::= MetaProperty
	NewExpression, // MemberExpression_NoLetOnly_NoFuncClass ::= 'new' MemberExpression Arguments
	0, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral ::= PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral
	IndexAccess, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral ::= MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral '[' Expression_In ']'
	PropertyAccess, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral TemplateLiteral
	0, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral ::= SuperProperty
	0, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral ::= MetaProperty
	NewExpression, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral ::= 'new' MemberExpression Arguments
	0, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= PrimaryExpression_NoFuncClass_NoLet_NoObjLiteral_Yield
	IndexAccess, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield TemplateLiteral_Yield
	0, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= SuperProperty_Yield
	0, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= MetaProperty
	NewExpression, // MemberExpression_NoLetOnly_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= 'new' MemberExpression_Yield Arguments_Yield
	0, // MemberExpression_NoLetOnly_NoLet ::= PrimaryExpression_NoLet
	IndexAccess, // MemberExpression_NoLetOnly_NoLet ::= MemberExpression_NoLet '[' Expression_In ']'
	PropertyAccess, // MemberExpression_NoLetOnly_NoLet ::= MemberExpression_NoLet '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLetOnly_NoLet ::= MemberExpression_NoLet TemplateLiteral
	0, // MemberExpression_NoLetOnly_NoLet ::= SuperProperty
	0, // MemberExpression_NoLetOnly_NoLet ::= MetaProperty
	NewExpression, // MemberExpression_NoLetOnly_NoLet ::= 'new' MemberExpression Arguments
	0, // MemberExpression_NoLetOnly_NoLet_Yield ::= PrimaryExpression_NoLet_Yield
	IndexAccess, // MemberExpression_NoLetOnly_NoLet_Yield ::= MemberExpression_NoLet_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // MemberExpression_NoLetOnly_NoLet_Yield ::= MemberExpression_NoLet_Yield '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLetOnly_NoLet_Yield ::= MemberExpression_NoLet_Yield TemplateLiteral_Yield
	0, // MemberExpression_NoLetOnly_NoLet_Yield ::= SuperProperty_Yield
	0, // MemberExpression_NoLetOnly_NoLet_Yield ::= MetaProperty
	NewExpression, // MemberExpression_NoLetOnly_NoLet_Yield ::= 'new' MemberExpression_Yield Arguments_Yield
	0, // MemberExpression_NoLetOnly_NoObjLiteral ::= PrimaryExpression_NoLet_NoObjLiteral
	IndexAccess, // MemberExpression_NoLetOnly_NoObjLiteral ::= MemberExpression_NoObjLiteral '[' Expression_In ']'
	PropertyAccess, // MemberExpression_NoLetOnly_NoObjLiteral ::= MemberExpression_NoObjLiteral '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLetOnly_NoObjLiteral ::= MemberExpression_NoObjLiteral TemplateLiteral
	0, // MemberExpression_NoLetOnly_NoObjLiteral ::= SuperProperty
	0, // MemberExpression_NoLetOnly_NoObjLiteral ::= MetaProperty
	NewExpression, // MemberExpression_NoLetOnly_NoObjLiteral ::= 'new' MemberExpression Arguments
	IndexAccess, // MemberExpression_NoLetOnly_StartWithLet ::= MemberExpression_NoLetOnly_StartWithLet '[' Expression_In ']'
	PropertyAccess, // MemberExpression_NoLetOnly_StartWithLet ::= MemberExpression_StartWithLet '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLetOnly_StartWithLet ::= MemberExpression_StartWithLet TemplateLiteral
	IndexAccess, // MemberExpression_NoLetOnly_StartWithLet_Yield ::= MemberExpression_NoLetOnly_StartWithLet_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // MemberExpression_NoLetOnly_StartWithLet_Yield ::= MemberExpression_StartWithLet_Yield '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLetOnly_StartWithLet_Yield ::= MemberExpression_StartWithLet_Yield TemplateLiteral_Yield
	0, // MemberExpression_NoLetOnly_Yield ::= PrimaryExpression_NoLet_Yield
	IndexAccess, // MemberExpression_NoLetOnly_Yield ::= MemberExpression_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // MemberExpression_NoLetOnly_Yield ::= MemberExpression_Yield '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoLetOnly_Yield ::= MemberExpression_Yield TemplateLiteral_Yield
	0, // MemberExpression_NoLetOnly_Yield ::= SuperProperty_Yield
	0, // MemberExpression_NoLetOnly_Yield ::= MetaProperty
	NewExpression, // MemberExpression_NoLetOnly_Yield ::= 'new' MemberExpression_Yield Arguments_Yield
	0, // MemberExpression_NoObjLiteral ::= PrimaryExpression_NoObjLiteral
	IndexAccess, // MemberExpression_NoObjLiteral ::= MemberExpression_NoObjLiteral '[' Expression_In ']'
	PropertyAccess, // MemberExpression_NoObjLiteral ::= MemberExpression_NoObjLiteral '.' IdentifierName
	TaggedTemplate, // MemberExpression_NoObjLiteral ::= MemberExpression_NoObjLiteral TemplateLiteral
	0, // MemberExpression_NoObjLiteral ::= SuperProperty
	0, // MemberExpression_NoObjLiteral ::= MetaProperty
	NewExpression, // MemberExpression_NoObjLiteral ::= 'new' MemberExpression Arguments
	IdentifierReference, // MemberExpression_StartWithLet ::= 'let'
	IndexAccess, // MemberExpression_StartWithLet ::= MemberExpression_NoLetOnly_StartWithLet '[' Expression_In ']'
	PropertyAccess, // MemberExpression_StartWithLet ::= MemberExpression_StartWithLet '.' IdentifierName
	TaggedTemplate, // MemberExpression_StartWithLet ::= MemberExpression_StartWithLet TemplateLiteral
	IdentifierReference, // MemberExpression_StartWithLet_Yield ::= 'let'
	IndexAccess, // MemberExpression_StartWithLet_Yield ::= MemberExpression_NoLetOnly_StartWithLet_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // MemberExpression_StartWithLet_Yield ::= MemberExpression_StartWithLet_Yield '.' IdentifierName
	TaggedTemplate, // MemberExpression_StartWithLet_Yield ::= MemberExpression_StartWithLet_Yield TemplateLiteral_Yield
	0, // MemberExpression_Yield ::= PrimaryExpression_Yield
	IndexAccess, // MemberExpression_Yield ::= MemberExpression_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // MemberExpression_Yield ::= MemberExpression_Yield '.' IdentifierName
	TaggedTemplate, // MemberExpression_Yield ::= MemberExpression_Yield TemplateLiteral_Yield
	0, // MemberExpression_Yield ::= SuperProperty_Yield
	0, // MemberExpression_Yield ::= MetaProperty
	NewExpression, // MemberExpression_Yield ::= 'new' MemberExpression_Yield Arguments_Yield
	SuperExpression, // SuperExpression ::= 'super'
	IndexAccess, // SuperProperty ::= SuperExpression '[' Expression_In ']'
	PropertyAccess, // SuperProperty ::= SuperExpression '.' IdentifierName
	IndexAccess, // SuperProperty_Yield ::= SuperExpression '[' Expression_In_Yield ']'
	PropertyAccess, // SuperProperty_Yield ::= SuperExpression '.' IdentifierName
	0, // MetaProperty ::= NewTarget
	NewTarget, // NewTarget ::= 'new' '.' 'target'
	0, // NewExpression ::= MemberExpression
	NewExpression, // NewExpression ::= 'new' NewExpression
	0, // NewExpression_NoFuncClass ::= MemberExpression_NoFuncClass
	NewExpression, // NewExpression_NoFuncClass ::= 'new' NewExpression
	0, // NewExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral
	NewExpression, // NewExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= 'new' NewExpression
	0, // NewExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield
	NewExpression, // NewExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= 'new' NewExpression_Yield
	0, // NewExpression_NoLet ::= MemberExpression_NoLet
	NewExpression, // NewExpression_NoLet ::= 'new' NewExpression
	0, // NewExpression_NoLet_Yield ::= MemberExpression_NoLet_Yield
	NewExpression, // NewExpression_NoLet_Yield ::= 'new' NewExpression_Yield
	0, // NewExpression_NoObjLiteral ::= MemberExpression_NoObjLiteral
	NewExpression, // NewExpression_NoObjLiteral ::= 'new' NewExpression
	0, // NewExpression_StartWithLet ::= MemberExpression_StartWithLet
	0, // NewExpression_StartWithLet_Yield ::= MemberExpression_StartWithLet_Yield
	0, // NewExpression_Yield ::= MemberExpression_Yield
	NewExpression, // NewExpression_Yield ::= 'new' NewExpression_Yield
	CallExpression, // CallExpression ::= MemberExpression Arguments
	CallExpression, // CallExpression ::= SuperCall
	CallExpression, // CallExpression ::= CallExpression Arguments
	IndexAccess, // CallExpression ::= CallExpression '[' Expression_In ']'
	PropertyAccess, // CallExpression ::= CallExpression '.' IdentifierName
	TaggedTemplate, // CallExpression ::= CallExpression TemplateLiteral
	CallExpression, // CallExpression_NoFuncClass ::= MemberExpression_NoFuncClass Arguments
	CallExpression, // CallExpression_NoFuncClass ::= SuperCall
	CallExpression, // CallExpression_NoFuncClass ::= CallExpression_NoFuncClass Arguments
	IndexAccess, // CallExpression_NoFuncClass ::= CallExpression_NoFuncClass '[' Expression_In ']'
	PropertyAccess, // CallExpression_NoFuncClass ::= CallExpression_NoFuncClass '.' IdentifierName
	TaggedTemplate, // CallExpression_NoFuncClass ::= CallExpression_NoFuncClass TemplateLiteral
	CallExpression, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral Arguments
	CallExpression, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= SuperCall
	CallExpression, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= CallExpression_NoFuncClass_NoLetSq_NoObjLiteral Arguments
	IndexAccess, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= CallExpression_NoFuncClass_NoLetSq_NoObjLiteral '[' Expression_In ']'
	PropertyAccess, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= CallExpression_NoFuncClass_NoLetSq_NoObjLiteral '.' IdentifierName
	TaggedTemplate, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= CallExpression_NoFuncClass_NoLetSq_NoObjLiteral TemplateLiteral
	CallExpression, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= MemberExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield Arguments_Yield
	CallExpression, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= SuperCall_Yield
	CallExpression, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield Arguments_Yield
	IndexAccess, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '.' IdentifierName
	TaggedTemplate, // CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield TemplateLiteral_Yield
	CallExpression, // CallExpression_NoLet ::= MemberExpression_NoLet Arguments
	CallExpression, // CallExpression_NoLet ::= SuperCall
	CallExpression, // CallExpression_NoLet ::= CallExpression_NoLet Arguments
	IndexAccess, // CallExpression_NoLet ::= CallExpression_NoLet '[' Expression_In ']'
	PropertyAccess, // CallExpression_NoLet ::= CallExpression_NoLet '.' IdentifierName
	TaggedTemplate, // CallExpression_NoLet ::= CallExpression_NoLet TemplateLiteral
	CallExpression, // CallExpression_NoLet_Yield ::= MemberExpression_NoLet_Yield Arguments_Yield
	CallExpression, // CallExpression_NoLet_Yield ::= SuperCall_Yield
	CallExpression, // CallExpression_NoLet_Yield ::= CallExpression_NoLet_Yield Arguments_Yield
	IndexAccess, // CallExpression_NoLet_Yield ::= CallExpression_NoLet_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // CallExpression_NoLet_Yield ::= CallExpression_NoLet_Yield '.' IdentifierName
	TaggedTemplate, // CallExpression_NoLet_Yield ::= CallExpression_NoLet_Yield TemplateLiteral_Yield
	CallExpression, // CallExpression_NoObjLiteral ::= MemberExpression_NoObjLiteral Arguments
	CallExpression, // CallExpression_NoObjLiteral ::= SuperCall
	CallExpression, // CallExpression_NoObjLiteral ::= CallExpression_NoObjLiteral Arguments
	IndexAccess, // CallExpression_NoObjLiteral ::= CallExpression_NoObjLiteral '[' Expression_In ']'
	PropertyAccess, // CallExpression_NoObjLiteral ::= CallExpression_NoObjLiteral '.' IdentifierName
	TaggedTemplate, // CallExpression_NoObjLiteral ::= CallExpression_NoObjLiteral TemplateLiteral
	CallExpression, // CallExpression_StartWithLet ::= MemberExpression_StartWithLet Arguments
	CallExpression, // CallExpression_StartWithLet ::= CallExpression_StartWithLet Arguments
	IndexAccess, // CallExpression_StartWithLet ::= CallExpression_StartWithLet '[' Expression_In ']'
	PropertyAccess, // CallExpression_StartWithLet ::= CallExpression_StartWithLet '.' IdentifierName
	TaggedTemplate, // CallExpression_StartWithLet ::= CallExpression_StartWithLet TemplateLiteral
	CallExpression, // CallExpression_StartWithLet_Yield ::= MemberExpression_StartWithLet_Yield Arguments_Yield
	CallExpression, // CallExpression_StartWithLet_Yield ::= CallExpression_StartWithLet_Yield Arguments_Yield
	IndexAccess, // CallExpression_StartWithLet_Yield ::= CallExpression_StartWithLet_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // CallExpression_StartWithLet_Yield ::= CallExpression_StartWithLet_Yield '.' IdentifierName
	TaggedTemplate, // CallExpression_StartWithLet_Yield ::= CallExpression_StartWithLet_Yield TemplateLiteral_Yield
	CallExpression, // CallExpression_Yield ::= MemberExpression_Yield Arguments_Yield
	CallExpression, // CallExpression_Yield ::= SuperCall_Yield
	CallExpression, // CallExpression_Yield ::= CallExpression_Yield Arguments_Yield
	IndexAccess, // CallExpression_Yield ::= CallExpression_Yield '[' Expression_In_Yield ']'
	PropertyAccess, // CallExpression_Yield ::= CallExpression_Yield '.' IdentifierName
	TaggedTemplate, // CallExpression_Yield ::= CallExpression_Yield TemplateLiteral_Yield
	0, // SuperCall ::= SuperExpression Arguments
	0, // SuperCall_Yield ::= SuperExpression Arguments_Yield
	Arguments, // Arguments ::= '(' ')'
	Arguments, // Arguments ::= '(' ArgumentList ')'
	Arguments, // Arguments_Yield ::= '(' ')'
	Arguments, // Arguments_Yield ::= '(' ArgumentList_Yield ')'
	0, // ArgumentList ::= AssignmentExpression_In
	0, // ArgumentList ::= '...' AssignmentExpression_In
	0, // ArgumentList ::= ArgumentList ',' AssignmentExpression_In
	0, // ArgumentList ::= ArgumentList ',' '...' AssignmentExpression_In
	0, // ArgumentList_Yield ::= AssignmentExpression_In_Yield
	0, // ArgumentList_Yield ::= '...' AssignmentExpression_In_Yield
	0, // ArgumentList_Yield ::= ArgumentList_Yield ',' AssignmentExpression_In_Yield
	0, // ArgumentList_Yield ::= ArgumentList_Yield ',' '...' AssignmentExpression_In_Yield
	0, // LeftHandSideExpression ::= NewExpression
	0, // LeftHandSideExpression ::= CallExpression
	0, // LeftHandSideExpression_NoFuncClass ::= NewExpression_NoFuncClass
	0, // LeftHandSideExpression_NoFuncClass ::= CallExpression_NoFuncClass
	0, // LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= NewExpression_NoFuncClass_NoLetSq_NoObjLiteral
	0, // LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= CallExpression_NoFuncClass_NoLetSq_NoObjLiteral
	0, // LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= NewExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield
	0, // LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= CallExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield
	0, // LeftHandSideExpression_NoLet ::= NewExpression_NoLet
	0, // LeftHandSideExpression_NoLet ::= CallExpression_NoLet
	0, // LeftHandSideExpression_NoLet_Yield ::= NewExpression_NoLet_Yield
	0, // LeftHandSideExpression_NoLet_Yield ::= CallExpression_NoLet_Yield
	0, // LeftHandSideExpression_NoObjLiteral ::= NewExpression_NoObjLiteral
	0, // LeftHandSideExpression_NoObjLiteral ::= CallExpression_NoObjLiteral
	0, // LeftHandSideExpression_StartWithLet ::= NewExpression_StartWithLet
	0, // LeftHandSideExpression_StartWithLet ::= CallExpression_StartWithLet
	0, // LeftHandSideExpression_StartWithLet_Yield ::= NewExpression_StartWithLet_Yield
	0, // LeftHandSideExpression_StartWithLet_Yield ::= CallExpression_StartWithLet_Yield
	0, // LeftHandSideExpression_Yield ::= NewExpression_Yield
	0, // LeftHandSideExpression_Yield ::= CallExpression_Yield
	0, // UpdateExpression ::= LeftHandSideExpression
	PostIncrementExpression, // UpdateExpression ::= LeftHandSideExpression .noLineBreak '++'
	PostDecrementExpression, // UpdateExpression ::= LeftHandSideExpression .noLineBreak '--'
	PreIncrementExpression, // UpdateExpression ::= '++' UnaryExpression
	PreDecrementExpression, // UpdateExpression ::= '--' UnaryExpression
	0, // UpdateExpression_NoFuncClass ::= LeftHandSideExpression_NoFuncClass
	PostIncrementExpression, // UpdateExpression_NoFuncClass ::= LeftHandSideExpression_NoFuncClass .noLineBreak '++'
	PostDecrementExpression, // UpdateExpression_NoFuncClass ::= LeftHandSideExpression_NoFuncClass .noLineBreak '--'
	PreIncrementExpression, // UpdateExpression_NoFuncClass ::= '++' UnaryExpression
	PreDecrementExpression, // UpdateExpression_NoFuncClass ::= '--' UnaryExpression
	0, // UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral
	PostIncrementExpression, // UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral .noLineBreak '++'
	PostDecrementExpression, // UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral .noLineBreak '--'
	PreIncrementExpression, // UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= '++' UnaryExpression
	PreDecrementExpression, // UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= '--' UnaryExpression
	0, // UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield
	PostIncrementExpression, // UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield .noLineBreak '++'
	PostDecrementExpression, // UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield .noLineBreak '--'
	PreIncrementExpression, // UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= '++' UnaryExpression_Yield
	PreDecrementExpression, // UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= '--' UnaryExpression_Yield
	0, // UpdateExpression_NoLet ::= LeftHandSideExpression_NoLet
	PostIncrementExpression, // UpdateExpression_NoLet ::= LeftHandSideExpression_NoLet .noLineBreak '++'
	PostDecrementExpression, // UpdateExpression_NoLet ::= LeftHandSideExpression_NoLet .noLineBreak '--'
	PreIncrementExpression, // UpdateExpression_NoLet ::= '++' UnaryExpression
	PreDecrementExpression, // UpdateExpression_NoLet ::= '--' UnaryExpression
	0, // UpdateExpression_NoLet_Yield ::= LeftHandSideExpression_NoLet_Yield
	PostIncrementExpression, // UpdateExpression_NoLet_Yield ::= LeftHandSideExpression_NoLet_Yield .noLineBreak '++'
	PostDecrementExpression, // UpdateExpression_NoLet_Yield ::= LeftHandSideExpression_NoLet_Yield .noLineBreak '--'
	PreIncrementExpression, // UpdateExpression_NoLet_Yield ::= '++' UnaryExpression_Yield
	PreDecrementExpression, // UpdateExpression_NoLet_Yield ::= '--' UnaryExpression_Yield
	0, // UpdateExpression_NoObjLiteral ::= LeftHandSideExpression_NoObjLiteral
	PostIncrementExpression, // UpdateExpression_NoObjLiteral ::= LeftHandSideExpression_NoObjLiteral .noLineBreak '++'
	PostDecrementExpression, // UpdateExpression_NoObjLiteral ::= LeftHandSideExpression_NoObjLiteral .noLineBreak '--'
	PreIncrementExpression, // UpdateExpression_NoObjLiteral ::= '++' UnaryExpression
	PreDecrementExpression, // UpdateExpression_NoObjLiteral ::= '--' UnaryExpression
	0, // UpdateExpression_StartWithLet ::= LeftHandSideExpression_StartWithLet
	PostIncrementExpression, // UpdateExpression_StartWithLet ::= LeftHandSideExpression_StartWithLet .noLineBreak '++'
	PostDecrementExpression, // UpdateExpression_StartWithLet ::= LeftHandSideExpression_StartWithLet .noLineBreak '--'
	0, // UpdateExpression_StartWithLet_Yield ::= LeftHandSideExpression_StartWithLet_Yield
	PostIncrementExpression, // UpdateExpression_StartWithLet_Yield ::= LeftHandSideExpression_StartWithLet_Yield .noLineBreak '++'
	PostDecrementExpression, // UpdateExpression_StartWithLet_Yield ::= LeftHandSideExpression_StartWithLet_Yield .noLineBreak '--'
	0, // UpdateExpression_Yield ::= LeftHandSideExpression_Yield
	PostIncrementExpression, // UpdateExpression_Yield ::= LeftHandSideExpression_Yield .noLineBreak '++'
	PostDecrementExpression, // UpdateExpression_Yield ::= LeftHandSideExpression_Yield .noLineBreak '--'
	PreIncrementExpression, // UpdateExpression_Yield ::= '++' UnaryExpression_Yield
	PreDecrementExpression, // UpdateExpression_Yield ::= '--' UnaryExpression_Yield
	0, // UnaryExpression ::= UpdateExpression
	UnaryExpression, // UnaryExpression ::= 'delete' UnaryExpression
	UnaryExpression, // UnaryExpression ::= 'void' UnaryExpression
	UnaryExpression, // UnaryExpression ::= 'typeof' UnaryExpression
	UnaryExpression, // UnaryExpression ::= '+' UnaryExpression
	UnaryExpression, // UnaryExpression ::= '-' UnaryExpression
	UnaryExpression, // UnaryExpression ::= '~' UnaryExpression
	UnaryExpression, // UnaryExpression ::= '!' UnaryExpression
	0, // UnaryExpression_NoFuncClass ::= UpdateExpression_NoFuncClass
	UnaryExpression, // UnaryExpression_NoFuncClass ::= 'delete' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass ::= 'void' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass ::= 'typeof' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass ::= '+' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass ::= '-' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass ::= '~' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass ::= '!' UnaryExpression
	0, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= 'delete' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= 'void' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= 'typeof' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= '+' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= '-' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= '~' UnaryExpression
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= '!' UnaryExpression
	0, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= 'delete' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= 'void' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= 'typeof' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= '+' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= '-' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= '~' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= '!' UnaryExpression_Yield
	0, // UnaryExpression_NoLet ::= UpdateExpression_NoLet
	UnaryExpression, // UnaryExpression_NoLet ::= 'delete' UnaryExpression
	UnaryExpression, // UnaryExpression_NoLet ::= 'void' UnaryExpression
	UnaryExpression, // UnaryExpression_NoLet ::= 'typeof' UnaryExpression
	UnaryExpression, // UnaryExpression_NoLet ::= '+' UnaryExpression
	UnaryExpression, // UnaryExpression_NoLet ::= '-' UnaryExpression
	UnaryExpression, // UnaryExpression_NoLet ::= '~' UnaryExpression
	UnaryExpression, // UnaryExpression_NoLet ::= '!' UnaryExpression
	0, // UnaryExpression_NoLet_Yield ::= UpdateExpression_NoLet_Yield
	UnaryExpression, // UnaryExpression_NoLet_Yield ::= 'delete' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoLet_Yield ::= 'void' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoLet_Yield ::= 'typeof' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoLet_Yield ::= '+' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoLet_Yield ::= '-' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoLet_Yield ::= '~' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_NoLet_Yield ::= '!' UnaryExpression_Yield
	0, // UnaryExpression_NoObjLiteral ::= UpdateExpression_NoObjLiteral
	UnaryExpression, // UnaryExpression_NoObjLiteral ::= 'delete' UnaryExpression
	UnaryExpression, // UnaryExpression_NoObjLiteral ::= 'void' UnaryExpression
	UnaryExpression, // UnaryExpression_NoObjLiteral ::= 'typeof' UnaryExpression
	UnaryExpression, // UnaryExpression_NoObjLiteral ::= '+' UnaryExpression
	UnaryExpression, // UnaryExpression_NoObjLiteral ::= '-' UnaryExpression
	UnaryExpression, // UnaryExpression_NoObjLiteral ::= '~' UnaryExpression
	UnaryExpression, // UnaryExpression_NoObjLiteral ::= '!' UnaryExpression
	0, // UnaryExpression_StartWithLet ::= UpdateExpression_StartWithLet
	0, // UnaryExpression_StartWithLet_Yield ::= UpdateExpression_StartWithLet_Yield
	0, // UnaryExpression_Yield ::= UpdateExpression_Yield
	UnaryExpression, // UnaryExpression_Yield ::= 'delete' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_Yield ::= 'void' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_Yield ::= 'typeof' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_Yield ::= '+' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_Yield ::= '-' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_Yield ::= '~' UnaryExpression_Yield
	UnaryExpression, // UnaryExpression_Yield ::= '!' UnaryExpression_Yield
	0, // ArithmeticExpression ::= UnaryExpression
	AdditiveExpression, // ArithmeticExpression ::= ArithmeticExpression '+' ArithmeticExpression
	AdditiveExpression, // ArithmeticExpression ::= ArithmeticExpression '-' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression ::= ArithmeticExpression '<<' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression ::= ArithmeticExpression '>>' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression ::= ArithmeticExpression '>>>' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression ::= ArithmeticExpression '*' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression ::= ArithmeticExpression '/' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression ::= ArithmeticExpression '%' ArithmeticExpression
	ExponentiationExpression, // ArithmeticExpression ::= UpdateExpression '**' ArithmeticExpression
	0, // ArithmeticExpression_NoFuncClass ::= UnaryExpression_NoFuncClass
	AdditiveExpression, // ArithmeticExpression_NoFuncClass ::= ArithmeticExpression_NoFuncClass '+' ArithmeticExpression
	AdditiveExpression, // ArithmeticExpression_NoFuncClass ::= ArithmeticExpression_NoFuncClass '-' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoFuncClass ::= ArithmeticExpression_NoFuncClass '<<' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoFuncClass ::= ArithmeticExpression_NoFuncClass '>>' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoFuncClass ::= ArithmeticExpression_NoFuncClass '>>>' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoFuncClass ::= ArithmeticExpression_NoFuncClass '*' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoFuncClass ::= ArithmeticExpression_NoFuncClass '/' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoFuncClass ::= ArithmeticExpression_NoFuncClass '%' ArithmeticExpression
	ExponentiationExpression, // ArithmeticExpression_NoFuncClass ::= UpdateExpression_NoFuncClass '**' ArithmeticExpression
	0, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral
	AdditiveExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral '+' ArithmeticExpression
	AdditiveExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral '-' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral '<<' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral '>>' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral '>>>' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral '*' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral '/' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral '%' ArithmeticExpression
	ExponentiationExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral ::= UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral '**' ArithmeticExpression
	0, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= UnaryExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield
	AdditiveExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '+' ArithmeticExpression_Yield
	AdditiveExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '-' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '<<' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '>>' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '>>>' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '*' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '/' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '%' ArithmeticExpression_Yield
	ExponentiationExpression, // ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= UpdateExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '**' ArithmeticExpression_Yield
	0, // ArithmeticExpression_NoLet ::= UnaryExpression_NoLet
	AdditiveExpression, // ArithmeticExpression_NoLet ::= ArithmeticExpression_NoLet '+' ArithmeticExpression
	AdditiveExpression, // ArithmeticExpression_NoLet ::= ArithmeticExpression_NoLet '-' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoLet ::= ArithmeticExpression_NoLet '<<' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoLet ::= ArithmeticExpression_NoLet '>>' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoLet ::= ArithmeticExpression_NoLet '>>>' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoLet ::= ArithmeticExpression_NoLet '*' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoLet ::= ArithmeticExpression_NoLet '/' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoLet ::= ArithmeticExpression_NoLet '%' ArithmeticExpression
	ExponentiationExpression, // ArithmeticExpression_NoLet ::= UpdateExpression_NoLet '**' ArithmeticExpression
	0, // ArithmeticExpression_NoLet_Yield ::= UnaryExpression_NoLet_Yield
	AdditiveExpression, // ArithmeticExpression_NoLet_Yield ::= ArithmeticExpression_NoLet_Yield '+' ArithmeticExpression_Yield
	AdditiveExpression, // ArithmeticExpression_NoLet_Yield ::= ArithmeticExpression_NoLet_Yield '-' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_NoLet_Yield ::= ArithmeticExpression_NoLet_Yield '<<' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_NoLet_Yield ::= ArithmeticExpression_NoLet_Yield '>>' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_NoLet_Yield ::= ArithmeticExpression_NoLet_Yield '>>>' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_NoLet_Yield ::= ArithmeticExpression_NoLet_Yield '*' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_NoLet_Yield ::= ArithmeticExpression_NoLet_Yield '/' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_NoLet_Yield ::= ArithmeticExpression_NoLet_Yield '%' ArithmeticExpression_Yield
	ExponentiationExpression, // ArithmeticExpression_NoLet_Yield ::= UpdateExpression_NoLet_Yield '**' ArithmeticExpression_Yield
	0, // ArithmeticExpression_NoObjLiteral ::= UnaryExpression_NoObjLiteral
	AdditiveExpression, // ArithmeticExpression_NoObjLiteral ::= ArithmeticExpression_NoObjLiteral '+' ArithmeticExpression
	AdditiveExpression, // ArithmeticExpression_NoObjLiteral ::= ArithmeticExpression_NoObjLiteral '-' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoObjLiteral ::= ArithmeticExpression_NoObjLiteral '<<' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoObjLiteral ::= ArithmeticExpression_NoObjLiteral '>>' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_NoObjLiteral ::= ArithmeticExpression_NoObjLiteral '>>>' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoObjLiteral ::= ArithmeticExpression_NoObjLiteral '*' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoObjLiteral ::= ArithmeticExpression_NoObjLiteral '/' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_NoObjLiteral ::= ArithmeticExpression_NoObjLiteral '%' ArithmeticExpression
	ExponentiationExpression, // ArithmeticExpression_NoObjLiteral ::= UpdateExpression_NoObjLiteral '**' ArithmeticExpression
	0, // ArithmeticExpression_StartWithLet ::= UnaryExpression_StartWithLet
	AdditiveExpression, // ArithmeticExpression_StartWithLet ::= ArithmeticExpression_StartWithLet '+' ArithmeticExpression
	AdditiveExpression, // ArithmeticExpression_StartWithLet ::= ArithmeticExpression_StartWithLet '-' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_StartWithLet ::= ArithmeticExpression_StartWithLet '<<' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_StartWithLet ::= ArithmeticExpression_StartWithLet '>>' ArithmeticExpression
	ShiftExpression, // ArithmeticExpression_StartWithLet ::= ArithmeticExpression_StartWithLet '>>>' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_StartWithLet ::= ArithmeticExpression_StartWithLet '*' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_StartWithLet ::= ArithmeticExpression_StartWithLet '/' ArithmeticExpression
	MultiplicativeExpression, // ArithmeticExpression_StartWithLet ::= ArithmeticExpression_StartWithLet '%' ArithmeticExpression
	ExponentiationExpression, // ArithmeticExpression_StartWithLet ::= UpdateExpression_StartWithLet '**' ArithmeticExpression
	0, // ArithmeticExpression_StartWithLet_Yield ::= UnaryExpression_StartWithLet_Yield
	AdditiveExpression, // ArithmeticExpression_StartWithLet_Yield ::= ArithmeticExpression_StartWithLet_Yield '+' ArithmeticExpression_Yield
	AdditiveExpression, // ArithmeticExpression_StartWithLet_Yield ::= ArithmeticExpression_StartWithLet_Yield '-' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_StartWithLet_Yield ::= ArithmeticExpression_StartWithLet_Yield '<<' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_StartWithLet_Yield ::= ArithmeticExpression_StartWithLet_Yield '>>' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_StartWithLet_Yield ::= ArithmeticExpression_StartWithLet_Yield '>>>' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_StartWithLet_Yield ::= ArithmeticExpression_StartWithLet_Yield '*' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_StartWithLet_Yield ::= ArithmeticExpression_StartWithLet_Yield '/' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_StartWithLet_Yield ::= ArithmeticExpression_StartWithLet_Yield '%' ArithmeticExpression_Yield
	ExponentiationExpression, // ArithmeticExpression_StartWithLet_Yield ::= UpdateExpression_StartWithLet_Yield '**' ArithmeticExpression_Yield
	0, // ArithmeticExpression_Yield ::= UnaryExpression_Yield
	AdditiveExpression, // ArithmeticExpression_Yield ::= ArithmeticExpression_Yield '+' ArithmeticExpression_Yield
	AdditiveExpression, // ArithmeticExpression_Yield ::= ArithmeticExpression_Yield '-' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_Yield ::= ArithmeticExpression_Yield '<<' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_Yield ::= ArithmeticExpression_Yield '>>' ArithmeticExpression_Yield
	ShiftExpression, // ArithmeticExpression_Yield ::= ArithmeticExpression_Yield '>>>' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_Yield ::= ArithmeticExpression_Yield '*' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_Yield ::= ArithmeticExpression_Yield '/' ArithmeticExpression_Yield
	MultiplicativeExpression, // ArithmeticExpression_Yield ::= ArithmeticExpression_Yield '%' ArithmeticExpression_Yield
	ExponentiationExpression, // ArithmeticExpression_Yield ::= UpdateExpression_Yield '**' ArithmeticExpression_Yield
	0, // BinaryExpression ::= ArithmeticExpression
	RelationalExpression, // BinaryExpression ::= BinaryExpression '<' BinaryExpression
	RelationalExpression, // BinaryExpression ::= BinaryExpression '>' BinaryExpression
	RelationalExpression, // BinaryExpression ::= BinaryExpression '<=' BinaryExpression
	RelationalExpression, // BinaryExpression ::= BinaryExpression '>=' BinaryExpression
	RelationalExpression, // BinaryExpression ::= BinaryExpression 'instanceof' BinaryExpression
	EqualityExpression, // BinaryExpression ::= BinaryExpression '==' BinaryExpression
	EqualityExpression, // BinaryExpression ::= BinaryExpression '!=' BinaryExpression
	EqualityExpression, // BinaryExpression ::= BinaryExpression '===' BinaryExpression
	EqualityExpression, // BinaryExpression ::= BinaryExpression '!==' BinaryExpression
	BitwiseANDExpression, // BinaryExpression ::= BinaryExpression '&' BinaryExpression
	BitwiseXORExpression, // BinaryExpression ::= BinaryExpression '^' BinaryExpression
	BitwiseORExpression, // BinaryExpression ::= BinaryExpression '|' BinaryExpression
	LogicalANDExpression, // BinaryExpression ::= BinaryExpression '&&' BinaryExpression
	LogicalORExpression, // BinaryExpression ::= BinaryExpression '||' BinaryExpression
	0, // BinaryExpression_In ::= ArithmeticExpression
	RelationalExpression, // BinaryExpression_In ::= BinaryExpression_In '<' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In ::= BinaryExpression_In '>' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In ::= BinaryExpression_In '<=' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In ::= BinaryExpression_In '>=' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In ::= BinaryExpression_In 'instanceof' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In ::= BinaryExpression_In 'in' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In ::= BinaryExpression_In '==' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In ::= BinaryExpression_In '!=' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In ::= BinaryExpression_In '===' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In ::= BinaryExpression_In '!==' BinaryExpression_In
	BitwiseANDExpression, // BinaryExpression_In ::= BinaryExpression_In '&' BinaryExpression_In
	BitwiseXORExpression, // BinaryExpression_In ::= BinaryExpression_In '^' BinaryExpression_In
	BitwiseORExpression, // BinaryExpression_In ::= BinaryExpression_In '|' BinaryExpression_In
	LogicalANDExpression, // BinaryExpression_In ::= BinaryExpression_In '&&' BinaryExpression_In
	LogicalORExpression, // BinaryExpression_In ::= BinaryExpression_In '||' BinaryExpression_In
	0, // BinaryExpression_In_NoFuncClass ::= ArithmeticExpression_NoFuncClass
	RelationalExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '<' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '>' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '<=' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '>=' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass 'instanceof' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass 'in' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '==' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '!=' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '===' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '!==' BinaryExpression_In
	BitwiseANDExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '&' BinaryExpression_In
	BitwiseXORExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '^' BinaryExpression_In
	BitwiseORExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '|' BinaryExpression_In
	LogicalANDExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '&&' BinaryExpression_In
	LogicalORExpression, // BinaryExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '||' BinaryExpression_In
	0, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '<' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '>' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '<=' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '>=' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral 'instanceof' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral 'in' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '==' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '!=' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '===' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '!==' BinaryExpression_In
	BitwiseANDExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '&' BinaryExpression_In
	BitwiseXORExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '^' BinaryExpression_In
	BitwiseORExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '|' BinaryExpression_In
	LogicalANDExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '&&' BinaryExpression_In
	LogicalORExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '||' BinaryExpression_In
	0, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ArithmeticExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '<' BinaryExpression_In_Yield
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '>' BinaryExpression_In_Yield
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '<=' BinaryExpression_In_Yield
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '>=' BinaryExpression_In_Yield
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield 'instanceof' BinaryExpression_In_Yield
	RelationalExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield 'in' BinaryExpression_In_Yield
	EqualityExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '==' BinaryExpression_In_Yield
	EqualityExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '!=' BinaryExpression_In_Yield
	EqualityExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '===' BinaryExpression_In_Yield
	EqualityExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '!==' BinaryExpression_In_Yield
	BitwiseANDExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '&' BinaryExpression_In_Yield
	BitwiseXORExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '^' BinaryExpression_In_Yield
	BitwiseORExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '|' BinaryExpression_In_Yield
	LogicalANDExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '&&' BinaryExpression_In_Yield
	LogicalORExpression, // BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '||' BinaryExpression_In_Yield
	0, // BinaryExpression_In_NoObjLiteral ::= ArithmeticExpression_NoObjLiteral
	RelationalExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '<' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '>' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '<=' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '>=' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral 'instanceof' BinaryExpression_In
	RelationalExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral 'in' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '==' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '!=' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '===' BinaryExpression_In
	EqualityExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '!==' BinaryExpression_In
	BitwiseANDExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '&' BinaryExpression_In
	BitwiseXORExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '^' BinaryExpression_In
	BitwiseORExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '|' BinaryExpression_In
	LogicalANDExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '&&' BinaryExpression_In
	LogicalORExpression, // BinaryExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '||' BinaryExpression_In
	0, // BinaryExpression_In_Yield ::= ArithmeticExpression_Yield
	RelationalExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '<' BinaryExpression_In_Yield
	RelationalExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '>' BinaryExpression_In_Yield
	RelationalExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '<=' BinaryExpression_In_Yield
	RelationalExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '>=' BinaryExpression_In_Yield
	RelationalExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield 'instanceof' BinaryExpression_In_Yield
	RelationalExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield 'in' BinaryExpression_In_Yield
	EqualityExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '==' BinaryExpression_In_Yield
	EqualityExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '!=' BinaryExpression_In_Yield
	EqualityExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '===' BinaryExpression_In_Yield
	EqualityExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '!==' BinaryExpression_In_Yield
	BitwiseANDExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '&' BinaryExpression_In_Yield
	BitwiseXORExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '^' BinaryExpression_In_Yield
	BitwiseORExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '|' BinaryExpression_In_Yield
	LogicalANDExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '&&' BinaryExpression_In_Yield
	LogicalORExpression, // BinaryExpression_In_Yield ::= BinaryExpression_In_Yield '||' BinaryExpression_In_Yield
	0, // BinaryExpression_NoLet ::= ArithmeticExpression_NoLet
	RelationalExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '<' BinaryExpression
	RelationalExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '>' BinaryExpression
	RelationalExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '<=' BinaryExpression
	RelationalExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '>=' BinaryExpression
	RelationalExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet 'instanceof' BinaryExpression
	EqualityExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '==' BinaryExpression
	EqualityExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '!=' BinaryExpression
	EqualityExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '===' BinaryExpression
	EqualityExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '!==' BinaryExpression
	BitwiseANDExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '&' BinaryExpression
	BitwiseXORExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '^' BinaryExpression
	BitwiseORExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '|' BinaryExpression
	LogicalANDExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '&&' BinaryExpression
	LogicalORExpression, // BinaryExpression_NoLet ::= BinaryExpression_NoLet '||' BinaryExpression
	0, // BinaryExpression_NoLet_Yield ::= ArithmeticExpression_NoLet_Yield
	RelationalExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '<' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '>' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '<=' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '>=' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield 'instanceof' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '==' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '!=' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '===' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '!==' BinaryExpression_Yield
	BitwiseANDExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '&' BinaryExpression_Yield
	BitwiseXORExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '^' BinaryExpression_Yield
	BitwiseORExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '|' BinaryExpression_Yield
	LogicalANDExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '&&' BinaryExpression_Yield
	LogicalORExpression, // BinaryExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '||' BinaryExpression_Yield
	0, // BinaryExpression_NoObjLiteral ::= ArithmeticExpression_NoObjLiteral
	RelationalExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '<' BinaryExpression
	RelationalExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '>' BinaryExpression
	RelationalExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '<=' BinaryExpression
	RelationalExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '>=' BinaryExpression
	RelationalExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral 'instanceof' BinaryExpression
	EqualityExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '==' BinaryExpression
	EqualityExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '!=' BinaryExpression
	EqualityExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '===' BinaryExpression
	EqualityExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '!==' BinaryExpression
	BitwiseANDExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '&' BinaryExpression
	BitwiseXORExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '^' BinaryExpression
	BitwiseORExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '|' BinaryExpression
	LogicalANDExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '&&' BinaryExpression
	LogicalORExpression, // BinaryExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '||' BinaryExpression
	0, // BinaryExpression_StartWithLet ::= ArithmeticExpression_StartWithLet
	RelationalExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '<' BinaryExpression
	RelationalExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '>' BinaryExpression
	RelationalExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '<=' BinaryExpression
	RelationalExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '>=' BinaryExpression
	RelationalExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet 'instanceof' BinaryExpression
	EqualityExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '==' BinaryExpression
	EqualityExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '!=' BinaryExpression
	EqualityExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '===' BinaryExpression
	EqualityExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '!==' BinaryExpression
	BitwiseANDExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '&' BinaryExpression
	BitwiseXORExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '^' BinaryExpression
	BitwiseORExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '|' BinaryExpression
	LogicalANDExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '&&' BinaryExpression
	LogicalORExpression, // BinaryExpression_StartWithLet ::= BinaryExpression_StartWithLet '||' BinaryExpression
	0, // BinaryExpression_StartWithLet_Yield ::= ArithmeticExpression_StartWithLet_Yield
	RelationalExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '<' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '>' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '<=' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '>=' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield 'instanceof' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '==' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '!=' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '===' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '!==' BinaryExpression_Yield
	BitwiseANDExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '&' BinaryExpression_Yield
	BitwiseXORExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '^' BinaryExpression_Yield
	BitwiseORExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '|' BinaryExpression_Yield
	LogicalANDExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '&&' BinaryExpression_Yield
	LogicalORExpression, // BinaryExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '||' BinaryExpression_Yield
	0, // BinaryExpression_Yield ::= ArithmeticExpression_Yield
	RelationalExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '<' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '>' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '<=' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '>=' BinaryExpression_Yield
	RelationalExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield 'instanceof' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '==' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '!=' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '===' BinaryExpression_Yield
	EqualityExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '!==' BinaryExpression_Yield
	BitwiseANDExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '&' BinaryExpression_Yield
	BitwiseXORExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '^' BinaryExpression_Yield
	BitwiseORExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '|' BinaryExpression_Yield
	LogicalANDExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '&&' BinaryExpression_Yield
	LogicalORExpression, // BinaryExpression_Yield ::= BinaryExpression_Yield '||' BinaryExpression_Yield
	0, // ConditionalExpression ::= BinaryExpression
	ConditionalExpression, // ConditionalExpression ::= BinaryExpression '?' AssignmentExpression_In ':' AssignmentExpression
	0, // ConditionalExpression_In ::= BinaryExpression_In
	ConditionalExpression, // ConditionalExpression_In ::= BinaryExpression_In '?' AssignmentExpression_In ':' AssignmentExpression_In
	0, // ConditionalExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass
	ConditionalExpression, // ConditionalExpression_In_NoFuncClass ::= BinaryExpression_In_NoFuncClass '?' AssignmentExpression_In ':' AssignmentExpression_In
	0, // ConditionalExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral
	ConditionalExpression, // ConditionalExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral '?' AssignmentExpression_In ':' AssignmentExpression_In
	0, // ConditionalExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield
	ConditionalExpression, // ConditionalExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= BinaryExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield '?' AssignmentExpression_In_Yield ':' AssignmentExpression_In_Yield
	0, // ConditionalExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral
	ConditionalExpression, // ConditionalExpression_In_NoObjLiteral ::= BinaryExpression_In_NoObjLiteral '?' AssignmentExpression_In ':' AssignmentExpression_In
	0, // ConditionalExpression_In_Yield ::= BinaryExpression_In_Yield
	ConditionalExpression, // ConditionalExpression_In_Yield ::= BinaryExpression_In_Yield '?' AssignmentExpression_In_Yield ':' AssignmentExpression_In_Yield
	0, // ConditionalExpression_NoLet ::= BinaryExpression_NoLet
	ConditionalExpression, // ConditionalExpression_NoLet ::= BinaryExpression_NoLet '?' AssignmentExpression_In ':' AssignmentExpression
	0, // ConditionalExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield
	ConditionalExpression, // ConditionalExpression_NoLet_Yield ::= BinaryExpression_NoLet_Yield '?' AssignmentExpression_In_Yield ':' AssignmentExpression_Yield
	0, // ConditionalExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral
	ConditionalExpression, // ConditionalExpression_NoObjLiteral ::= BinaryExpression_NoObjLiteral '?' AssignmentExpression_In ':' AssignmentExpression
	0, // ConditionalExpression_StartWithLet ::= BinaryExpression_StartWithLet
	ConditionalExpression, // ConditionalExpression_StartWithLet ::= BinaryExpression_StartWithLet '?' AssignmentExpression_In ':' AssignmentExpression
	0, // ConditionalExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield
	ConditionalExpression, // ConditionalExpression_StartWithLet_Yield ::= BinaryExpression_StartWithLet_Yield '?' AssignmentExpression_In_Yield ':' AssignmentExpression_Yield
	0, // ConditionalExpression_Yield ::= BinaryExpression_Yield
	ConditionalExpression, // ConditionalExpression_Yield ::= BinaryExpression_Yield '?' AssignmentExpression_In_Yield ':' AssignmentExpression_Yield
	0, // AssignmentExpression ::= ConditionalExpression
	0, // AssignmentExpression ::= ArrowFunction
	AssignmentExpression, // AssignmentExpression ::= LeftHandSideExpression '=' AssignmentExpression
	AssignmentExpression, // AssignmentExpression ::= LeftHandSideExpression AssignmentOperator AssignmentExpression
	0, // AssignmentExpression_In ::= ConditionalExpression_In
	0, // AssignmentExpression_In ::= ArrowFunction_In
	AssignmentExpression, // AssignmentExpression_In ::= LeftHandSideExpression '=' AssignmentExpression_In
	AssignmentExpression, // AssignmentExpression_In ::= LeftHandSideExpression AssignmentOperator AssignmentExpression_In
	0, // AssignmentExpression_In_NoFuncClass ::= ConditionalExpression_In_NoFuncClass
	0, // AssignmentExpression_In_NoFuncClass ::= ArrowFunction_In
	AssignmentExpression, // AssignmentExpression_In_NoFuncClass ::= LeftHandSideExpression_NoFuncClass '=' AssignmentExpression_In
	AssignmentExpression, // AssignmentExpression_In_NoFuncClass ::= LeftHandSideExpression_NoFuncClass AssignmentOperator AssignmentExpression_In
	0, // AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= ConditionalExpression_In_NoFuncClass_NoLetSq_NoObjLiteral
	0, // AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= ArrowFunction_In
	AssignmentExpression, // AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral '=' AssignmentExpression_In
	AssignmentExpression, // AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral AssignmentOperator AssignmentExpression_In
	0, // AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ConditionalExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield
	0, // AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= YieldExpression_In
	0, // AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= ArrowFunction_In_Yield
	AssignmentExpression, // AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield '=' AssignmentExpression_In_Yield
	AssignmentExpression, // AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= LeftHandSideExpression_NoFuncClass_NoLetSq_NoObjLiteral_Yield AssignmentOperator AssignmentExpression_In_Yield
	0, // AssignmentExpression_In_NoObjLiteral ::= ConditionalExpression_In_NoObjLiteral
	0, // AssignmentExpression_In_NoObjLiteral ::= ArrowFunction_In
	AssignmentExpression, // AssignmentExpression_In_NoObjLiteral ::= LeftHandSideExpression_NoObjLiteral '=' AssignmentExpression_In
	AssignmentExpression, // AssignmentExpression_In_NoObjLiteral ::= LeftHandSideExpression_NoObjLiteral AssignmentOperator AssignmentExpression_In
	0, // AssignmentExpression_In_Yield ::= ConditionalExpression_In_Yield
	0, // AssignmentExpression_In_Yield ::= YieldExpression_In
	0, // AssignmentExpression_In_Yield ::= ArrowFunction_In_Yield
	AssignmentExpression, // AssignmentExpression_In_Yield ::= LeftHandSideExpression_Yield '=' AssignmentExpression_In_Yield
	AssignmentExpression, // AssignmentExpression_In_Yield ::= LeftHandSideExpression_Yield AssignmentOperator AssignmentExpression_In_Yield
	0, // AssignmentExpression_NoLet ::= ConditionalExpression_NoLet
	0, // AssignmentExpression_NoLet ::= ArrowFunction
	AssignmentExpression, // AssignmentExpression_NoLet ::= LeftHandSideExpression_NoLet '=' AssignmentExpression
	AssignmentExpression, // AssignmentExpression_NoLet ::= LeftHandSideExpression_NoLet AssignmentOperator AssignmentExpression
	0, // AssignmentExpression_NoLet_Yield ::= ConditionalExpression_NoLet_Yield
	0, // AssignmentExpression_NoLet_Yield ::= YieldExpression
	0, // AssignmentExpression_NoLet_Yield ::= ArrowFunction_Yield
	AssignmentExpression, // AssignmentExpression_NoLet_Yield ::= LeftHandSideExpression_NoLet_Yield '=' AssignmentExpression_Yield
	AssignmentExpression, // AssignmentExpression_NoLet_Yield ::= LeftHandSideExpression_NoLet_Yield AssignmentOperator AssignmentExpression_Yield
	0, // AssignmentExpression_NoObjLiteral ::= ConditionalExpression_NoObjLiteral
	0, // AssignmentExpression_NoObjLiteral ::= ArrowFunction
	AssignmentExpression, // AssignmentExpression_NoObjLiteral ::= LeftHandSideExpression_NoObjLiteral '=' AssignmentExpression
	AssignmentExpression, // AssignmentExpression_NoObjLiteral ::= LeftHandSideExpression_NoObjLiteral AssignmentOperator AssignmentExpression
	0, // AssignmentExpression_StartWithLet ::= ConditionalExpression_StartWithLet
	AssignmentExpression, // AssignmentExpression_StartWithLet ::= LeftHandSideExpression_StartWithLet '=' AssignmentExpression
	AssignmentExpression, // AssignmentExpression_StartWithLet ::= LeftHandSideExpression_StartWithLet AssignmentOperator AssignmentExpression
	0, // AssignmentExpression_StartWithLet_Yield ::= ConditionalExpression_StartWithLet_Yield
	AssignmentExpression, // AssignmentExpression_StartWithLet_Yield ::= LeftHandSideExpression_StartWithLet_Yield '=' AssignmentExpression_Yield
	AssignmentExpression, // AssignmentExpression_StartWithLet_Yield ::= LeftHandSideExpression_StartWithLet_Yield AssignmentOperator AssignmentExpression_Yield
	0, // AssignmentExpression_Yield ::= ConditionalExpression_Yield
	0, // AssignmentExpression_Yield ::= YieldExpression
	0, // AssignmentExpression_Yield ::= ArrowFunction_Yield
	AssignmentExpression, // AssignmentExpression_Yield ::= LeftHandSideExpression_Yield '=' AssignmentExpression_Yield
	AssignmentExpression, // AssignmentExpression_Yield ::= LeftHandSideExpression_Yield AssignmentOperator AssignmentExpression_Yield
	AssignmentOperator, // AssignmentOperator ::= '*='
	AssignmentOperator, // AssignmentOperator ::= '/='
	AssignmentOperator, // AssignmentOperator ::= '%='
	AssignmentOperator, // AssignmentOperator ::= '+='
	AssignmentOperator, // AssignmentOperator ::= '-='
	AssignmentOperator, // AssignmentOperator ::= '<<='
	AssignmentOperator, // AssignmentOperator ::= '>>='
	AssignmentOperator, // AssignmentOperator ::= '>>>='
	AssignmentOperator, // AssignmentOperator ::= '&='
	AssignmentOperator, // AssignmentOperator ::= '^='
	AssignmentOperator, // AssignmentOperator ::= '|='
	AssignmentOperator, // AssignmentOperator ::= '**='
	0, // Expression_In ::= AssignmentExpression_In
	0, // Expression_In ::= Expression_In ',' AssignmentExpression_In
	0, // Expression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral
	0, // Expression_In_NoFuncClass_NoLetSq_NoObjLiteral ::= Expression_In_NoFuncClass_NoLetSq_NoObjLiteral ',' AssignmentExpression_In
	0, // Expression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= AssignmentExpression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield
	0, // Expression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ::= Expression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ',' AssignmentExpression_In_Yield
	0, // Expression_In_Yield ::= AssignmentExpression_In_Yield
	0, // Expression_In_Yield ::= Expression_In_Yield ',' AssignmentExpression_In_Yield
	0, // Expression_NoLet ::= AssignmentExpression_NoLet
	0, // Expression_NoLet ::= Expression_NoLet ',' AssignmentExpression
	0, // Expression_NoLet_Yield ::= AssignmentExpression_NoLet_Yield
	0, // Expression_NoLet_Yield ::= Expression_NoLet_Yield ',' AssignmentExpression_Yield
	0, // Expression_StartWithLet ::= AssignmentExpression_StartWithLet
	0, // Expression_StartWithLet ::= Expression_StartWithLet ',' AssignmentExpression
	0, // Expression_StartWithLet_Yield ::= AssignmentExpression_StartWithLet_Yield
	0, // Expression_StartWithLet_Yield ::= Expression_StartWithLet_Yield ',' AssignmentExpression_Yield
	0, // Statement ::= BlockStatement
	0, // Statement ::= VariableStatement
	0, // Statement ::= EmptyStatement
	0, // Statement ::= ExpressionStatement
	0, // Statement ::= IfStatement
	0, // Statement ::= BreakableStatement
	0, // Statement ::= ContinueStatement
	0, // Statement ::= BreakStatement
	0, // Statement ::= WithStatement
	0, // Statement ::= LabelledStatement
	0, // Statement ::= ThrowStatement
	0, // Statement ::= TryStatement
	0, // Statement ::= DebuggerStatement
	0, // Statement_Return ::= BlockStatement_Return
	0, // Statement_Return ::= VariableStatement
	0, // Statement_Return ::= EmptyStatement
	0, // Statement_Return ::= ExpressionStatement
	0, // Statement_Return ::= IfStatement_Return
	0, // Statement_Return ::= BreakableStatement_Return
	0, // Statement_Return ::= ContinueStatement
	0, // Statement_Return ::= BreakStatement
	0, // Statement_Return ::= ReturnStatement
	0, // Statement_Return ::= WithStatement_Return
	0, // Statement_Return ::= LabelledStatement_Return
	0, // Statement_Return ::= ThrowStatement
	0, // Statement_Return ::= TryStatement_Return
	0, // Statement_Return ::= DebuggerStatement
	0, // Statement_Return_Yield ::= BlockStatement_Return_Yield
	0, // Statement_Return_Yield ::= VariableStatement_Yield
	0, // Statement_Return_Yield ::= EmptyStatement
	0, // Statement_Return_Yield ::= ExpressionStatement_Yield
	0, // Statement_Return_Yield ::= IfStatement_Return_Yield
	0, // Statement_Return_Yield ::= BreakableStatement_Return_Yield
	0, // Statement_Return_Yield ::= ContinueStatement_Yield
	0, // Statement_Return_Yield ::= BreakStatement_Yield
	0, // Statement_Return_Yield ::= ReturnStatement_Yield
	0, // Statement_Return_Yield ::= WithStatement_Return_Yield
	0, // Statement_Return_Yield ::= LabelledStatement_Return_Yield
	0, // Statement_Return_Yield ::= ThrowStatement_Yield
	0, // Statement_Return_Yield ::= TryStatement_Return_Yield
	0, // Statement_Return_Yield ::= DebuggerStatement
	0, // Declaration ::= HoistableDeclaration
	0, // Declaration ::= ClassDeclaration
	0, // Declaration ::= LexicalDeclaration_In
	0, // Declaration_Yield ::= HoistableDeclaration_Yield
	0, // Declaration_Yield ::= ClassDeclaration_Yield
	0, // Declaration_Yield ::= LexicalDeclaration_In_Yield
	0, // HoistableDeclaration ::= FunctionDeclaration
	0, // HoistableDeclaration ::= GeneratorDeclaration
	0, // HoistableDeclaration_Default ::= FunctionDeclaration_Default
	0, // HoistableDeclaration_Default ::= GeneratorDeclaration_Default
	0, // HoistableDeclaration_Yield ::= FunctionDeclaration_Yield
	0, // HoistableDeclaration_Yield ::= GeneratorDeclaration_Yield
	0, // BreakableStatement ::= IterationStatement
	0, // BreakableStatement ::= SwitchStatement
	0, // BreakableStatement_Return ::= IterationStatement_Return
	0, // BreakableStatement_Return ::= SwitchStatement_Return
	0, // BreakableStatement_Return_Yield ::= IterationStatement_Return_Yield
	0, // BreakableStatement_Return_Yield ::= SwitchStatement_Return_Yield
	0, // BlockStatement ::= Block
	0, // BlockStatement_Return ::= Block_Return
	0, // BlockStatement_Return_Yield ::= Block_Return_Yield
	Block, // Block ::= '{' StatementList '}'
	Block, // Block ::= '{' '}'
	Block, // Block_Return ::= '{' StatementList_Return '}'
	Block, // Block_Return ::= '{' '}'
	Block, // Block_Return_Yield ::= '{' StatementList_Return_Yield '}'
	Block, // Block_Return_Yield ::= '{' '}'
	0, // StatementList ::= StatementListItem
	0, // StatementList ::= StatementList StatementListItem
	0, // StatementList_Return ::= StatementListItem_Return
	0, // StatementList_Return ::= StatementList_Return StatementListItem_Return
	0, // StatementList_Return_Yield ::= StatementListItem_Return_Yield
	0, // StatementList_Return_Yield ::= StatementList_Return_Yield StatementListItem_Return_Yield
	0, // StatementListItem ::= Statement
	0, // StatementListItem ::= Declaration
	0, // StatementListItem_Return ::= Statement_Return
	0, // StatementListItem_Return ::= Declaration
	0, // StatementListItem_Return_Yield ::= Statement_Return_Yield
	0, // StatementListItem_Return_Yield ::= Declaration_Yield
	LexicalDeclaration, // LexicalDeclaration_In ::= LetOrConst BindingList_In ';'
	LexicalDeclaration, // LexicalDeclaration_In_Yield ::= LetOrConst BindingList_In_Yield ';'
	0, // LetOrConst ::= 'let'
	0, // LetOrConst ::= 'const'
	0, // BindingList ::= LexicalBinding
	0, // BindingList ::= BindingList ',' LexicalBinding
	0, // BindingList_In ::= LexicalBinding_In
	0, // BindingList_In ::= BindingList_In ',' LexicalBinding_In
	0, // BindingList_In_Yield ::= LexicalBinding_In_Yield
	0, // BindingList_In_Yield ::= BindingList_In_Yield ',' LexicalBinding_In_Yield
	0, // BindingList_Yield ::= LexicalBinding_Yield
	0, // BindingList_Yield ::= BindingList_Yield ',' LexicalBinding_Yield
	LexicalBinding, // LexicalBinding ::= BindingIdentifier Initializeropt
	LexicalBinding, // LexicalBinding ::= BindingPattern Initializer
	LexicalBinding, // LexicalBinding_In ::= BindingIdentifier Initializeropt_In
	LexicalBinding, // LexicalBinding_In ::= BindingPattern Initializer_In
	LexicalBinding, // LexicalBinding_In_Yield ::= BindingIdentifier_Yield Initializeropt_In_Yield
	LexicalBinding, // LexicalBinding_In_Yield ::= BindingPattern_Yield Initializer_In_Yield
	LexicalBinding, // LexicalBinding_Yield ::= BindingIdentifier_Yield Initializeropt_Yield
	LexicalBinding, // LexicalBinding_Yield ::= BindingPattern_Yield Initializer_Yield
	VariableStatement, // VariableStatement ::= 'var' VariableDeclarationList_In ';'
	VariableStatement, // VariableStatement_Yield ::= 'var' VariableDeclarationList_In_Yield ';'
	0, // VariableDeclarationList ::= VariableDeclaration
	0, // VariableDeclarationList ::= VariableDeclarationList ',' VariableDeclaration
	0, // VariableDeclarationList_In ::= VariableDeclaration_In
	0, // VariableDeclarationList_In ::= VariableDeclarationList_In ',' VariableDeclaration_In
	0, // VariableDeclarationList_In_Yield ::= VariableDeclaration_In_Yield
	0, // VariableDeclarationList_In_Yield ::= VariableDeclarationList_In_Yield ',' VariableDeclaration_In_Yield
	0, // VariableDeclarationList_Yield ::= VariableDeclaration_Yield
	0, // VariableDeclarationList_Yield ::= VariableDeclarationList_Yield ',' VariableDeclaration_Yield
	VariableDeclaration, // VariableDeclaration ::= BindingIdentifier Initializeropt
	VariableDeclaration, // VariableDeclaration ::= BindingPattern Initializer
	VariableDeclaration, // VariableDeclaration_In ::= BindingIdentifier Initializeropt_In
	VariableDeclaration, // VariableDeclaration_In ::= BindingPattern Initializer_In
	VariableDeclaration, // VariableDeclaration_In_Yield ::= BindingIdentifier_Yield Initializeropt_In_Yield
	VariableDeclaration, // VariableDeclaration_In_Yield ::= BindingPattern_Yield Initializer_In_Yield
	VariableDeclaration, // VariableDeclaration_Yield ::= BindingIdentifier_Yield Initializeropt_Yield
	VariableDeclaration, // VariableDeclaration_Yield ::= BindingPattern_Yield Initializer_Yield
	0, // BindingPattern ::= ObjectBindingPattern
	0, // BindingPattern ::= ArrayBindingPattern
	0, // BindingPattern_Yield ::= ObjectBindingPattern_Yield
	0, // BindingPattern_Yield ::= ArrayBindingPattern_Yield
	ObjectBindingPattern, // ObjectBindingPattern ::= '{' '}'
	ObjectBindingPattern, // ObjectBindingPattern ::= '{' BindingPropertyList '}'
	ObjectBindingPattern, // ObjectBindingPattern ::= '{' BindingPropertyList ',' '}'
	ObjectBindingPattern, // ObjectBindingPattern_Yield ::= '{' '}'
	ObjectBindingPattern, // ObjectBindingPattern_Yield ::= '{' BindingPropertyList_Yield '}'
	ObjectBindingPattern, // ObjectBindingPattern_Yield ::= '{' BindingPropertyList_Yield ',' '}'
	ArrayBindingPattern, // ArrayBindingPattern ::= '[' Elisionopt BindingRestElementopt ']'
	ArrayBindingPattern, // ArrayBindingPattern ::= '[' BindingElementList ']'
	ArrayBindingPattern, // ArrayBindingPattern ::= '[' BindingElementList ',' Elisionopt BindingRestElementopt ']'
	ArrayBindingPattern, // ArrayBindingPattern_Yield ::= '[' Elisionopt BindingRestElementopt_Yield ']'
	ArrayBindingPattern, // ArrayBindingPattern_Yield ::= '[' BindingElementList_Yield ']'
	ArrayBindingPattern, // ArrayBindingPattern_Yield ::= '[' BindingElementList_Yield ',' Elisionopt BindingRestElementopt_Yield ']'
	0, // BindingPropertyList ::= BindingProperty
	0, // BindingPropertyList ::= BindingPropertyList ',' BindingProperty
	0, // BindingPropertyList_Yield ::= BindingProperty_Yield
	0, // BindingPropertyList_Yield ::= BindingPropertyList_Yield ',' BindingProperty_Yield
	0, // BindingElementList ::= BindingElisionElement
	0, // BindingElementList ::= BindingElementList ',' BindingElisionElement
	0, // BindingElementList_Yield ::= BindingElisionElement_Yield
	0, // BindingElementList_Yield ::= BindingElementList_Yield ',' BindingElisionElement_Yield
	BindingElisionElement, // BindingElisionElement ::= Elisionopt BindingElement
	BindingElisionElement, // BindingElisionElement_Yield ::= Elisionopt BindingElement_Yield
	BindingProperty, // BindingProperty ::= SingleNameBinding
	BindingProperty, // BindingProperty ::= PropertyName ':' BindingElement
	BindingProperty, // BindingProperty_Yield ::= SingleNameBinding_Yield
	BindingProperty, // BindingProperty_Yield ::= PropertyName_Yield ':' BindingElement_Yield
	BindingElement, // BindingElement ::= SingleNameBinding
	BindingElement, // BindingElement ::= BindingPattern Initializeropt_In
	BindingElement, // BindingElement_Yield ::= SingleNameBinding_Yield
	BindingElement, // BindingElement_Yield ::= BindingPattern_Yield Initializeropt_In_Yield
	SingleNameBinding, // SingleNameBinding ::= BindingIdentifier Initializeropt_In
	SingleNameBinding, // SingleNameBinding_Yield ::= BindingIdentifier_Yield Initializeropt_In_Yield
	BindingRestElement, // BindingRestElement ::= '...' BindingIdentifier
	BindingRestElement, // BindingRestElement_Yield ::= '...' BindingIdentifier_Yield
	EmptyStatement, // EmptyStatement ::= ';' .emptyStatement
	ExpressionStatement, // ExpressionStatement ::= Expression_In_NoFuncClass_NoLetSq_NoObjLiteral ';'
	ExpressionStatement, // ExpressionStatement_Yield ::= Expression_In_NoFuncClass_NoLetSq_NoObjLiteral_Yield ';'
	IfStatement, // IfStatement ::= 'if' '(' Expression_In ')' Statement 'else' Statement
	IfStatement, // IfStatement ::= 'if' '(' Expression_In ')' Statement %prec 'else'
	IfStatement, // IfStatement_Return ::= 'if' '(' Expression_In ')' Statement_Return 'else' Statement_Return
	IfStatement, // IfStatement_Return ::= 'if' '(' Expression_In ')' Statement_Return %prec 'else'
	IfStatement, // IfStatement_Return_Yield ::= 'if' '(' Expression_In_Yield ')' Statement_Return_Yield 'else' Statement_Return_Yield
	IfStatement, // IfStatement_Return_Yield ::= 'if' '(' Expression_In_Yield ')' Statement_Return_Yield %prec 'else'
	DoWhileStatement, // IterationStatement ::= 'do' Statement 'while' '(' Expression_In ')' ';' .doWhile
	WhileStatement, // IterationStatement ::= 'while' '(' Expression_In ')' Statement
	ForStatement, // IterationStatement ::= 'for' '(' Expressionopt_NoLet ';' .forSC Expressionopt_In ';' .forSC Expressionopt_In ')' Statement
	ForStatement, // IterationStatement ::= 'for' '(' Expression_StartWithLet ';' .forSC Expressionopt_In ';' .forSC Expressionopt_In ')' Statement
	ForStatement, // IterationStatement ::= 'for' '(' 'var' VariableDeclarationList ';' .forSC Expressionopt_In ';' .forSC Expressionopt_In ')' Statement
	ForStatement, // IterationStatement ::= 'for' '(' LetOrConst BindingList ';' .forSC Expressionopt_In ';' .forSC Expressionopt_In ')' Statement
	ForInStatement, // IterationStatement ::= 'for' '(' LeftHandSideExpression_NoLet 'in' Expression_In ')' Statement
	ForInStatement, // IterationStatement ::= 'for' '(' LeftHandSideExpression_StartWithLet 'in' Expression_In ')' Statement
	ForInStatement, // IterationStatement ::= 'for' '(' 'var' ForBinding 'in' Expression_In ')' Statement
	ForInStatement, // IterationStatement ::= 'for' '(' ForDeclaration 'in' Expression_In ')' Statement
	ForOfStatement, // IterationStatement ::= 'for' '(' LeftHandSideExpression_NoLet 'of' AssignmentExpression_In ')' Statement
	ForOfStatement, // IterationStatement ::= 'for' '(' 'var' ForBinding 'of' AssignmentExpression_In ')' Statement
	ForOfStatement, // IterationStatement ::= 'for' '(' ForDeclaration 'of' AssignmentExpression_In ')' Statement
	DoWhileStatement, // IterationStatement_Return ::= 'do' Statement_Return 'while' '(' Expression_In ')' ';' .doWhile
	WhileStatement, // IterationStatement_Return ::= 'while' '(' Expression_In ')' Statement_Return
	ForStatement, // IterationStatement_Return ::= 'for' '(' Expressionopt_NoLet ';' .forSC Expressionopt_In ';' .forSC Expressionopt_In ')' Statement_Return
	ForStatement, // IterationStatement_Return ::= 'for' '(' Expression_StartWithLet ';' .forSC Expressionopt_In ';' .forSC Expressionopt_In ')' Statement_Return
	ForStatement, // IterationStatement_Return ::= 'for' '(' 'var' VariableDeclarationList ';' .forSC Expressionopt_In ';' .forSC Expressionopt_In ')' Statement_Return
	ForStatement, // IterationStatement_Return ::= 'for' '(' LetOrConst BindingList ';' .forSC Expressionopt_In ';' .forSC Expressionopt_In ')' Statement_Return
	ForInStatement, // IterationStatement_Return ::= 'for' '(' LeftHandSideExpression_NoLet 'in' Expression_In ')' Statement_Return
	ForInStatement, // IterationStatement_Return ::= 'for' '(' LeftHandSideExpression_StartWithLet 'in' Expression_In ')' Statement_Return
	ForInStatement, // IterationStatement_Return ::= 'for' '(' 'var' ForBinding 'in' Expression_In ')' Statement_Return
	ForInStatement, // IterationStatement_Return ::= 'for' '(' ForDeclaration 'in' Expression_In ')' Statement_Return
	ForOfStatement, // IterationStatement_Return ::= 'for' '(' LeftHandSideExpression_NoLet 'of' AssignmentExpression_In ')' Statement_Return
	ForOfStatement, // IterationStatement_Return ::= 'for' '(' 'var' ForBinding 'of' AssignmentExpression_In ')' Statement_Return
	ForOfStatement, // IterationStatement_Return ::= 'for' '(' ForDeclaration 'of' AssignmentExpression_In ')' Statement_Return
	DoWhileStatement, // IterationStatement_Return_Yield ::= 'do' Statement_Return_Yield 'while' '(' Expression_In_Yield ')' ';' .doWhile
	WhileStatement, // IterationStatement_Return_Yield ::= 'while' '(' Expression_In_Yield ')' Statement_Return_Yield
	ForStatement, // IterationStatement_Return_Yield ::= 'for' '(' Expressionopt_NoLet_Yield ';' .forSC Expressionopt_In_Yield ';' .forSC Expressionopt_In_Yield ')' Statement_Return_Yield
	ForStatement, // IterationStatement_Return_Yield ::= 'for' '(' Expression_StartWithLet_Yield ';' .forSC Expressionopt_In_Yield ';' .forSC Expressionopt_In_Yield ')' Statement_Return_Yield
	ForStatement, // IterationStatement_Return_Yield ::= 'for' '(' 'var' VariableDeclarationList_Yield ';' .forSC Expressionopt_In_Yield ';' .forSC Expressionopt_In_Yield ')' Statement_Return_Yield
	ForStatement, // IterationStatement_Return_Yield ::= 'for' '(' LetOrConst BindingList_Yield ';' .forSC Expressionopt_In_Yield ';' .forSC Expressionopt_In_Yield ')' Statement_Return_Yield
	ForInStatement, // IterationStatement_Return_Yield ::= 'for' '(' LeftHandSideExpression_NoLet_Yield 'in' Expression_In_Yield ')' Statement_Return_Yield
	ForInStatement, // IterationStatement_Return_Yield ::= 'for' '(' LeftHandSideExpression_StartWithLet_Yield 'in' Expression_In_Yield ')' Statement_Return_Yield
	ForInStatement, // IterationStatement_Return_Yield ::= 'for' '(' 'var' ForBinding_Yield 'in' Expression_In_Yield ')' Statement_Return_Yield
	ForInStatement, // IterationStatement_Return_Yield ::= 'for' '(' ForDeclaration_Yield 'in' Expression_In_Yield ')' Statement_Return_Yield
	ForOfStatement, // IterationStatement_Return_Yield ::= 'for' '(' LeftHandSideExpression_NoLet_Yield 'of' AssignmentExpression_In_Yield ')' Statement_Return_Yield
	ForOfStatement, // IterationStatement_Return_Yield ::= 'for' '(' 'var' ForBinding_Yield 'of' AssignmentExpression_In_Yield ')' Statement_Return_Yield
	ForOfStatement, // IterationStatement_Return_Yield ::= 'for' '(' ForDeclaration_Yield 'of' AssignmentExpression_In_Yield ')' Statement_Return_Yield
	0, // ForDeclaration ::= LetOrConst ForBinding
	0, // ForDeclaration_Yield ::= LetOrConst ForBinding_Yield
	ForBinding, // ForBinding ::= BindingIdentifier
	ForBinding, // ForBinding ::= BindingPattern
	ForBinding, // ForBinding_Yield ::= BindingIdentifier_Yield
	ForBinding, // ForBinding_Yield ::= BindingPattern_Yield
	ContinueStatement, // ContinueStatement ::= 'continue' ';'
	ContinueStatement, // ContinueStatement ::= 'continue' .noLineBreak LabelIdentifier ';'
	ContinueStatement, // ContinueStatement_Yield ::= 'continue' ';'
	ContinueStatement, // ContinueStatement_Yield ::= 'continue' .noLineBreak LabelIdentifier_Yield ';'
	BreakStatement, // BreakStatement ::= 'break' ';'
	BreakStatement, // BreakStatement ::= 'break' .noLineBreak LabelIdentifier ';'
	BreakStatement, // BreakStatement_Yield ::= 'break' ';'
	BreakStatement, // BreakStatement_Yield ::= 'break' .noLineBreak LabelIdentifier_Yield ';'
	ReturnStatement, // ReturnStatement ::= 'return' ';'
	ReturnStatement, // ReturnStatement ::= 'return' .noLineBreak Expression_In ';'
	ReturnStatement, // ReturnStatement_Yield ::= 'return' ';'
	ReturnStatement, // ReturnStatement_Yield ::= 'return' .noLineBreak Expression_In_Yield ';'
	WithStatement, // WithStatement ::= 'with' '(' Expression_In ')' Statement
	WithStatement, // WithStatement_Return ::= 'with' '(' Expression_In ')' Statement_Return
	WithStatement, // WithStatement_Return_Yield ::= 'with' '(' Expression_In_Yield ')' Statement_Return_Yield
	SwitchStatement, // SwitchStatement ::= 'switch' '(' Expression_In ')' CaseBlock
	SwitchStatement, // SwitchStatement_Return ::= 'switch' '(' Expression_In ')' CaseBlock_Return
	SwitchStatement, // SwitchStatement_Return_Yield ::= 'switch' '(' Expression_In_Yield ')' CaseBlock_Return_Yield
	CaseBlock, // CaseBlock ::= '{' CaseClausesopt '}'
	CaseBlock, // CaseBlock ::= '{' CaseClausesopt DefaultClause CaseClausesopt '}'
	CaseBlock, // CaseBlock_Return ::= '{' CaseClausesopt_Return '}'
	CaseBlock, // CaseBlock_Return ::= '{' CaseClausesopt_Return DefaultClause_Return CaseClausesopt_Return '}'
	CaseBlock, // CaseBlock_Return_Yield ::= '{' CaseClausesopt_Return_Yield '}'
	CaseBlock, // CaseBlock_Return_Yield ::= '{' CaseClausesopt_Return_Yield DefaultClause_Return_Yield CaseClausesopt_Return_Yield '}'
	0, // CaseClauses ::= CaseClause
	0, // CaseClauses ::= CaseClauses CaseClause
	0, // CaseClauses_Return ::= CaseClause_Return
	0, // CaseClauses_Return ::= CaseClauses_Return CaseClause_Return
	0, // CaseClauses_Return_Yield ::= CaseClause_Return_Yield
	0, // CaseClauses_Return_Yield ::= CaseClauses_Return_Yield CaseClause_Return_Yield
	CaseClause, // CaseClause ::= 'case' Expression_In ':' StatementList
	CaseClause, // CaseClause ::= 'case' Expression_In ':'
	CaseClause, // CaseClause_Return ::= 'case' Expression_In ':' StatementList_Return
	CaseClause, // CaseClause_Return ::= 'case' Expression_In ':'
	CaseClause, // CaseClause_Return_Yield ::= 'case' Expression_In_Yield ':' StatementList_Return_Yield
	CaseClause, // CaseClause_Return_Yield ::= 'case' Expression_In_Yield ':'
	DefaultClause, // DefaultClause ::= 'default' ':' StatementList
	DefaultClause, // DefaultClause ::= 'default' ':'
	DefaultClause, // DefaultClause_Return ::= 'default' ':' StatementList_Return
	DefaultClause, // DefaultClause_Return ::= 'default' ':'
	DefaultClause, // DefaultClause_Return_Yield ::= 'default' ':' StatementList_Return_Yield
	DefaultClause, // DefaultClause_Return_Yield ::= 'default' ':'
	LabelledStatement, // LabelledStatement ::= Identifier ':' LabelledItem
	LabelledStatement, // LabelledStatement ::= 'yield' ':' LabelledItem
	LabelledStatement, // LabelledStatement_Return ::= Identifier ':' LabelledItem_Return
	LabelledStatement, // LabelledStatement_Return ::= 'yield' ':' LabelledItem_Return
	LabelledStatement, // LabelledStatement_Return_Yield ::= Identifier ':' LabelledItem_Return_Yield
	0, // LabelledItem ::= Statement
	0, // LabelledItem ::= FunctionDeclaration
	0, // LabelledItem_Return ::= Statement_Return
	0, // LabelledItem_Return ::= FunctionDeclaration
	0, // LabelledItem_Return_Yield ::= Statement_Return_Yield
	0, // LabelledItem_Return_Yield ::= FunctionDeclaration_Yield
	ThrowStatement, // ThrowStatement ::= 'throw' .noLineBreak Expression_In ';'
	ThrowStatement, // ThrowStatement_Yield ::= 'throw' .noLineBreak Expression_In_Yield ';'
	TryStatement, // TryStatement ::= 'try' Block Catch
	TryStatement, // TryStatement ::= 'try' Block Catch Finally
	TryStatement, // TryStatement ::= 'try' Block Finally
	TryStatement, // TryStatement_Return ::= 'try' Block_Return Catch_Return
	TryStatement, // TryStatement_Return ::= 'try' Block_Return Catch_Return Finally_Return
	TryStatement, // TryStatement_Return ::= 'try' Block_Return Finally_Return
	TryStatement, // TryStatement_Return_Yield ::= 'try' Block_Return_Yield Catch_Return_Yield
	TryStatement, // TryStatement_Return_Yield ::= 'try' Block_Return_Yield Catch_Return_Yield Finally_Return_Yield
	TryStatement, // TryStatement_Return_Yield ::= 'try' Block_Return_Yield Finally_Return_Yield
	Catch, // Catch ::= 'catch' '(' CatchParameter ')' Block
	Catch, // Catch_Return ::= 'catch' '(' CatchParameter ')' Block_Return
	Catch, // Catch_Return_Yield ::= 'catch' '(' CatchParameter_Yield ')' Block_Return_Yield
	Finally, // Finally ::= 'finally' Block
	Finally, // Finally_Return ::= 'finally' Block_Return
	Finally, // Finally_Return_Yield ::= 'finally' Block_Return_Yield
	CatchParameter, // CatchParameter ::= BindingIdentifier
	CatchParameter, // CatchParameter ::= BindingPattern
	CatchParameter, // CatchParameter_Yield ::= BindingIdentifier_Yield
	CatchParameter, // CatchParameter_Yield ::= BindingPattern_Yield
	DebuggerStatement, // DebuggerStatement ::= 'debugger' ';'
	FunctionDeclaration, // FunctionDeclaration ::= 'function' BindingIdentifier '(' FormalParameters ')' '{' FunctionBody '}'
	FunctionDeclaration, // FunctionDeclaration_Default ::= 'function' BindingIdentifier '(' FormalParameters ')' '{' FunctionBody '}'
	FunctionDeclaration, // FunctionDeclaration_Default ::= 'function' '(' FormalParameters ')' '{' FunctionBody '}'
	FunctionDeclaration, // FunctionDeclaration_Yield ::= 'function' BindingIdentifier_Yield '(' FormalParameters ')' '{' FunctionBody '}'
	FunctionExpression, // FunctionExpression ::= 'function' BindingIdentifier '(' FormalParameters ')' '{' FunctionBody '}'
	FunctionExpression, // FunctionExpression ::= 'function' '(' FormalParameters ')' '{' FunctionBody '}'
	0, // StrictFormalParameters ::= FormalParameters
	0, // StrictFormalParameters_Yield ::= FormalParameters_Yield
	FormalParameters, // FormalParameters ::= FormalParameterList
	FormalParameters, // FormalParameters ::=
	FormalParameters, // FormalParameters_Yield ::= FormalParameterList_Yield
	FormalParameters, // FormalParameters_Yield ::=
	0, // FormalParameterList ::= FunctionRestParameter
	0, // FormalParameterList ::= FormalsList
	0, // FormalParameterList ::= FormalsList ',' FunctionRestParameter
	0, // FormalParameterList_Yield ::= FunctionRestParameter_Yield
	0, // FormalParameterList_Yield ::= FormalsList_Yield
	0, // FormalParameterList_Yield ::= FormalsList_Yield ',' FunctionRestParameter_Yield
	0, // FormalsList ::= FormalParameter
	0, // FormalsList ::= FormalsList ',' FormalParameter
	0, // FormalsList_Yield ::= FormalParameter_Yield
	0, // FormalsList_Yield ::= FormalsList_Yield ',' FormalParameter_Yield
	FunctionRestParameter, // FunctionRestParameter ::= BindingRestElement
	FunctionRestParameter, // FunctionRestParameter_Yield ::= BindingRestElement_Yield
	FormalParameter, // FormalParameter ::= BindingElement
	FormalParameter, // FormalParameter_Yield ::= BindingElement_Yield
	0, // FunctionBody ::= StatementList_Return
	0, // FunctionBody ::=
	0, // FunctionBody_Yield ::= StatementList_Return_Yield
	0, // FunctionBody_Yield ::=
	ArrowFunction, // ArrowFunction ::= ArrowParameters .noLineBreak '=>' ConciseBody
	ArrowFunction, // ArrowFunction_In ::= ArrowParameters .noLineBreak '=>' ConciseBody_In
	ArrowFunction, // ArrowFunction_In_Yield ::= ArrowParameters_Yield .noLineBreak '=>' ConciseBody_In
	ArrowFunction, // ArrowFunction_Yield ::= ArrowParameters_Yield .noLineBreak '=>' ConciseBody
	ArrowParameters, // ArrowParameters ::= BindingIdentifier
	ArrowParameters, // ArrowParameters ::= CoverParenthesizedExpressionAndArrowParameterList
	ArrowParameters, // ArrowParameters_Yield ::= BindingIdentifier_Yield
	ArrowParameters, // ArrowParameters_Yield ::= CoverParenthesizedExpressionAndArrowParameterList_Yield
	ConciseBody, // ConciseBody ::= AssignmentExpression_NoObjLiteral
	ConciseBody, // ConciseBody ::= '{' FunctionBody '}'
	ConciseBody, // ConciseBody_In ::= AssignmentExpression_In_NoObjLiteral
	ConciseBody, // ConciseBody_In ::= '{' FunctionBody '}'
	MethodDefinition, // MethodDefinition ::= PropertyName '(' StrictFormalParameters ')' '{' FunctionBody '}'
	0, // MethodDefinition ::= GeneratorMethod
	MethodDefinition, // MethodDefinition ::= 'get' PropertyName '(' ')' '{' FunctionBody '}'
	MethodDefinition, // MethodDefinition ::= 'set' PropertyName '(' PropertySetParameterList ')' '{' FunctionBody '}'
	MethodDefinition, // MethodDefinition_Yield ::= PropertyName_Yield '(' StrictFormalParameters_Yield ')' '{' FunctionBody_Yield '}'
	0, // MethodDefinition_Yield ::= GeneratorMethod_Yield
	MethodDefinition, // MethodDefinition_Yield ::= 'get' PropertyName_Yield '(' ')' '{' FunctionBody_Yield '}'
	MethodDefinition, // MethodDefinition_Yield ::= 'set' PropertyName_Yield '(' PropertySetParameterList ')' '{' FunctionBody_Yield '}'
	0, // PropertySetParameterList ::= FormalParameter
	GeneratorMethod, // GeneratorMethod ::= '*' PropertyName '(' StrictFormalParameters_Yield ')' '{' GeneratorBody '}'
	GeneratorMethod, // GeneratorMethod_Yield ::= '*' PropertyName_Yield '(' StrictFormalParameters_Yield ')' '{' GeneratorBody '}'
	GeneratorDeclaration, // GeneratorDeclaration ::= 'function' '*' BindingIdentifier '(' FormalParameters_Yield ')' '{' GeneratorBody '}'
	GeneratorDeclaration, // GeneratorDeclaration_Default ::= 'function' '*' BindingIdentifier '(' FormalParameters_Yield ')' '{' GeneratorBody '}'
	GeneratorDeclaration, // GeneratorDeclaration_Default ::= 'function' '*' '(' FormalParameters_Yield ')' '{' GeneratorBody '}'
	GeneratorDeclaration, // GeneratorDeclaration_Yield ::= 'function' '*' BindingIdentifier_Yield '(' FormalParameters_Yield ')' '{' GeneratorBody '}'
	GeneratorExpression, // GeneratorExpression ::= 'function' '*' BindingIdentifier_Yield '(' FormalParameters_Yield ')' '{' GeneratorBody '}'
	GeneratorExpression, // GeneratorExpression ::= 'function' '*' '(' FormalParameters_Yield ')' '{' GeneratorBody '}'
	0, // GeneratorBody ::= FunctionBody_Yield
	YieldExpression, // YieldExpression ::= 'yield'
	YieldExpression, // YieldExpression ::= 'yield' .afterYield .noLineBreak AssignmentExpression_Yield
	YieldExpression, // YieldExpression ::= 'yield' .afterYield .noLineBreak '*' AssignmentExpression_Yield
	YieldExpression, // YieldExpression_In ::= 'yield'
	YieldExpression, // YieldExpression_In ::= 'yield' .afterYield .noLineBreak AssignmentExpression_In_Yield
	YieldExpression, // YieldExpression_In ::= 'yield' .afterYield .noLineBreak '*' AssignmentExpression_In_Yield
	ClassDeclaration, // ClassDeclaration ::= 'class' BindingIdentifier ClassTail
	ClassDeclaration, // ClassDeclaration_Default ::= 'class' BindingIdentifier ClassTail
	ClassDeclaration, // ClassDeclaration_Default ::= 'class' ClassTail
	ClassDeclaration, // ClassDeclaration_Yield ::= 'class' BindingIdentifier_Yield ClassTail_Yield
	ClassExpression, // ClassExpression ::= 'class' BindingIdentifier ClassTail
	ClassExpression, // ClassExpression ::= 'class' ClassTail
	ClassExpression, // ClassExpression_Yield ::= 'class' BindingIdentifier_Yield ClassTail_Yield
	ClassExpression, // ClassExpression_Yield ::= 'class' ClassTail_Yield
	0, // ClassTail ::= ClassHeritage ClassBody
	0, // ClassTail ::= ClassBody
	0, // ClassTail_Yield ::= ClassHeritage_Yield ClassBody_Yield
	0, // ClassTail_Yield ::= ClassBody_Yield
	ClassHeritage, // ClassHeritage ::= 'extends' LeftHandSideExpression
	ClassHeritage, // ClassHeritage_Yield ::= 'extends' LeftHandSideExpression_Yield
	ClassBody, // ClassBody ::= '{' ClassElementList '}'
	ClassBody, // ClassBody ::= '{' '}'
	ClassBody, // ClassBody_Yield ::= '{' ClassElementList_Yield '}'
	ClassBody, // ClassBody_Yield ::= '{' '}'
	0, // ClassElementList ::= ClassElement
	0, // ClassElementList ::= ClassElementList ClassElement
	0, // ClassElementList_Yield ::= ClassElement_Yield
	0, // ClassElementList_Yield ::= ClassElementList_Yield ClassElement_Yield
	ClassElement, // ClassElement ::= MethodDefinition
	ClassElement, // ClassElement ::= 'static' MethodDefinition
	ClassElement, // ClassElement ::= ';'
	ClassElement, // ClassElement_Yield ::= MethodDefinition_Yield
	ClassElement, // ClassElement_Yield ::= 'static' MethodDefinition_Yield
	ClassElement, // ClassElement_Yield ::= ';'
	Module, // Module ::= ModuleBodyopt
	0, // ModuleBody ::= ModuleItemList
	0, // ModuleItemList ::= ModuleItem
	0, // ModuleItemList ::= ModuleItemList ModuleItem
	ModuleItem, // ModuleItem ::= ImportDeclaration
	ModuleItem, // ModuleItem ::= ExportDeclaration
	ModuleItem, // ModuleItem ::= StatementListItem
	ImportDeclaration, // ImportDeclaration ::= 'import' ImportClause FromClause ';'
	ImportDeclaration, // ImportDeclaration ::= 'import' ModuleSpecifier ';'
	0, // ImportClause ::= ImportedDefaultBinding
	0, // ImportClause ::= NameSpaceImport
	0, // ImportClause ::= NamedImports
	0, // ImportClause ::= ImportedDefaultBinding ',' NameSpaceImport
	0, // ImportClause ::= ImportedDefaultBinding ',' NamedImports
	0, // ImportedDefaultBinding ::= ImportedBinding
	NameSpaceImport, // NameSpaceImport ::= '*' 'as' ImportedBinding
	NamedImports, // NamedImports ::= '{' '}'
	NamedImports, // NamedImports ::= '{' ImportsList '}'
	NamedImports, // NamedImports ::= '{' ImportsList ',' '}'
	0, // FromClause ::= 'from' ModuleSpecifier
	0, // ImportsList ::= ImportSpecifier
	0, // ImportsList ::= ImportsList ',' ImportSpecifier
	ImportSpecifier, // ImportSpecifier ::= ImportedBinding
	ImportSpecifier, // ImportSpecifier ::= IdentifierName 'as' ImportedBinding
	ModuleSpecifier, // ModuleSpecifier ::= StringLiteral
	0, // ImportedBinding ::= BindingIdentifier
	ExportDeclaration, // ExportDeclaration ::= 'export' '*' FromClause ';'
	ExportDeclaration, // ExportDeclaration ::= 'export' ExportClause FromClause ';'
	ExportDeclaration, // ExportDeclaration ::= 'export' ExportClause ';'
	ExportDeclaration, // ExportDeclaration ::= 'export' VariableStatement
	ExportDeclaration, // ExportDeclaration ::= 'export' Declaration
	ExportDefault, // ExportDeclaration ::= 'export' 'default' HoistableDeclaration_Default
	ExportDefault, // ExportDeclaration ::= 'export' 'default' ClassDeclaration_Default
	ExportDefault, // ExportDeclaration ::= 'export' 'default' AssignmentExpression_In_NoFuncClass ';'
	ExportClause, // ExportClause ::= '{' '}'
	ExportClause, // ExportClause ::= '{' ExportsList '}'
	ExportClause, // ExportClause ::= '{' ExportsList ',' '}'
	0, // ExportsList ::= ExportSpecifier
	0, // ExportsList ::= ExportsList ',' ExportSpecifier
	ExportSpecifier, // ExportSpecifier ::= IdentifierName
	ExportSpecifier, // ExportSpecifier ::= IdentifierName 'as' IdentifierName
	0, // JSXChild_optlist ::= JSXChild_optlist JSXChild
	0, // JSXChild_optlist ::=
	0, // JSXChild_Yield_optlist ::= JSXChild_Yield_optlist JSXChild_Yield
	0, // JSXChild_Yield_optlist ::=
	JSXElement, // JSXElement ::= JSXSelfClosingElement
	JSXElement, // JSXElement ::= JSXOpeningElement JSXChild_optlist JSXClosingElement
	JSXElement, // JSXElement_Yield ::= JSXSelfClosingElement_Yield
	JSXElement, // JSXElement_Yield ::= JSXOpeningElement_Yield JSXChild_Yield_optlist JSXClosingElement
	0, // JSXAttribute_optlist ::= JSXAttribute_optlist JSXAttribute
	0, // JSXAttribute_optlist ::=
	0, // JSXAttribute_Yield_optlist ::= JSXAttribute_Yield_optlist JSXAttribute_Yield
	0, // JSXAttribute_Yield_optlist ::=
	JSXSelfClosingElement, // JSXSelfClosingElement ::= '<' JSXElementName JSXAttribute_optlist '/' '>'
	JSXSelfClosingElement, // JSXSelfClosingElement_Yield ::= '<' JSXElementName JSXAttribute_Yield_optlist '/' '>'
	JSXOpeningElement, // JSXOpeningElement ::= '<' JSXElementName JSXAttribute_optlist '>'
	JSXOpeningElement, // JSXOpeningElement_Yield ::= '<' JSXElementName JSXAttribute_Yield_optlist '>'
	JSXClosingElement, // JSXClosingElement ::= '<' '/' JSXElementName '>'
	JSXElementName, // JSXElementName ::= jsxIdentifier
	JSXElementName, // JSXElementName ::= jsxIdentifier ':' jsxIdentifier
	JSXElementName, // JSXElementName ::= JSXMemberExpression
	0, // JSXMemberExpression ::= jsxIdentifier '.' jsxIdentifier
	0, // JSXMemberExpression ::= JSXMemberExpression '.' jsxIdentifier
	JSXAttribute, // JSXAttribute ::= JSXAttributeName '=' JSXAttributeValue
	JSXSpreadAttribute, // JSXAttribute ::= '{' '...' AssignmentExpression_In '}'
	JSXAttribute, // JSXAttribute_Yield ::= JSXAttributeName '=' JSXAttributeValue_Yield
	JSXSpreadAttribute, // JSXAttribute_Yield ::= '{' '...' AssignmentExpression_In_Yield '}'
	JSXAttributeName, // JSXAttributeName ::= jsxIdentifier
	JSXAttributeName, // JSXAttributeName ::= jsxIdentifier ':' jsxIdentifier
	JSXAttributeValue, // JSXAttributeValue ::= jsxStringLiteral
	JSXAttributeValue, // JSXAttributeValue ::= '{' AssignmentExpression_In '}'
	JSXAttributeValue, // JSXAttributeValue ::= JSXElement
	JSXAttributeValue, // JSXAttributeValue_Yield ::= jsxStringLiteral
	JSXAttributeValue, // JSXAttributeValue_Yield ::= '{' AssignmentExpression_In_Yield '}'
	JSXAttributeValue, // JSXAttributeValue_Yield ::= JSXElement_Yield
	JSXText, // JSXChild ::= jsxText
	JSXChild, // JSXChild ::= JSXElement
	JSXChild, // JSXChild ::= '{' AssignmentExpressionopt_In '}'
	JSXText, // JSXChild_Yield ::= jsxText
	JSXChild, // JSXChild_Yield ::= JSXElement_Yield
	JSXChild, // JSXChild_Yield ::= '{' AssignmentExpressionopt_In_Yield '}'
	0, // Elisionopt ::= Elision
	0, // Elisionopt ::=
	0, // Initializeropt ::= Initializer
	0, // Initializeropt ::=
	0, // Initializeropt_In ::= Initializer_In
	0, // Initializeropt_In ::=
	0, // Initializeropt_In_Yield ::= Initializer_In_Yield
	0, // Initializeropt_In_Yield ::=
	0, // Initializeropt_Yield ::= Initializer_Yield
	0, // Initializeropt_Yield ::=
	0, // BindingRestElementopt ::= BindingRestElement
	0, // BindingRestElementopt ::=
	0, // BindingRestElementopt_Yield ::= BindingRestElement_Yield
	0, // BindingRestElementopt_Yield ::=
	0, // Expressionopt_In ::= Expression_In
	0, // Expressionopt_In ::=
	0, // Expressionopt_In_Yield ::= Expression_In_Yield
	0, // Expressionopt_In_Yield ::=
	0, // Expressionopt_NoLet ::= Expression_NoLet
	0, // Expressionopt_NoLet ::=
	0, // Expressionopt_NoLet_Yield ::= Expression_NoLet_Yield
	0, // Expressionopt_NoLet_Yield ::=
	0, // CaseClausesopt ::= CaseClauses
	0, // CaseClausesopt ::=
	0, // CaseClausesopt_Return ::= CaseClauses_Return
	0, // CaseClausesopt_Return ::=
	0, // CaseClausesopt_Return_Yield ::= CaseClauses_Return_Yield
	0, // CaseClausesopt_Return_Yield ::=
	0, // ModuleBodyopt ::= ModuleBody
	0, // ModuleBodyopt ::=
	0, // AssignmentExpressionopt_In ::= AssignmentExpression_In
	0, // AssignmentExpressionopt_In ::=
	0, // AssignmentExpressionopt_In_Yield ::= AssignmentExpression_In_Yield
	0, // AssignmentExpressionopt_In_Yield ::=
}

var nodeTypeStr = [...]string{
	"NONE",
	"IdentifierName",
	"IdentifierReference",
	"BindingIdentifier",
	"LabelIdentifier",
	"ThisExpression",
	"RegularExpression",
	"ParenthesizedExpression",
	"Literal",
	"ArrayLiteral",
	"SpreadElement",
	"ObjectLiteral",
	"PropertyDefinition",
	"SyntaxError",
	"LiteralPropertyName",
	"ComputedPropertyName",
	"CoverInitializedName",
	"Initializer",
	"TemplateLiteral",
	"IndexAccess",
	"PropertyAccess",
	"TaggedTemplate",
	"NewExpression",
	"SuperExpression",
	"NewTarget",
	"CallExpression",
	"Arguments",
	"PostIncrementExpression",
	"PostDecrementExpression",
	"PreIncrementExpression",
	"PreDecrementExpression",
	"UnaryExpression",
	"AdditiveExpression",
	"ShiftExpression",
	"MultiplicativeExpression",
	"ExponentiationExpression",
	"RelationalExpression",
	"EqualityExpression",
	"BitwiseANDExpression",
	"BitwiseXORExpression",
	"BitwiseORExpression",
	"LogicalANDExpression",
	"LogicalORExpression",
	"ConditionalExpression",
	"AssignmentExpression",
	"AssignmentOperator",
	"Block",
	"LexicalDeclaration",
	"LexicalBinding",
	"VariableStatement",
	"VariableDeclaration",
	"ObjectBindingPattern",
	"ArrayBindingPattern",
	"BindingElisionElement",
	"BindingProperty",
	"BindingElement",
	"SingleNameBinding",
	"BindingRestElement",
	"EmptyStatement",
	"ExpressionStatement",
	"IfStatement",
	"DoWhileStatement",
	"WhileStatement",
	"ForStatement",
	"ForInStatement",
	"ForOfStatement",
	"ForBinding",
	"ContinueStatement",
	"BreakStatement",
	"ReturnStatement",
	"WithStatement",
	"SwitchStatement",
	"CaseBlock",
	"CaseClause",
	"DefaultClause",
	"LabelledStatement",
	"ThrowStatement",
	"TryStatement",
	"Catch",
	"Finally",
	"CatchParameter",
	"DebuggerStatement",
	"FunctionDeclaration",
	"FunctionExpression",
	"FormalParameters",
	"FunctionRestParameter",
	"FormalParameter",
	"ArrowFunction",
	"ArrowParameters",
	"ConciseBody",
	"MethodDefinition",
	"GeneratorMethod",
	"GeneratorDeclaration",
	"GeneratorExpression",
	"YieldExpression",
	"ClassDeclaration",
	"ClassExpression",
	"ClassHeritage",
	"ClassBody",
	"ClassElement",
	"Module",
	"ModuleItem",
	"ImportDeclaration",
	"NameSpaceImport",
	"NamedImports",
	"ImportSpecifier",
	"ModuleSpecifier",
	"ExportDeclaration",
	"ExportDefault",
	"ExportClause",
	"ExportSpecifier",
	"JSXElement",
	"JSXSelfClosingElement",
	"JSXOpeningElement",
	"JSXClosingElement",
	"JSXElementName",
	"JSXAttribute",
	"JSXSpreadAttribute",
	"JSXAttributeName",
	"JSXAttributeValue",
	"JSXText",
	"JSXChild",
	"InsertedSemicolon",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}

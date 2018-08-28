package og

import (
	"github.com/alecthomas/participle"
)

type INI struct {
	Pack     string    `"package" @Ident`
	ProgBody *ProgBody `[ @@ ]`
}

type ProgBody struct {
	Imports  []string    `[ "import"  "{" { @String } "}" ]`
	TopLevel []*TopLevel `{ @@ }`
}

type TopLevel struct {
	Struct *Struct `( @@`
	Func   *Func   `| @@ )`
}

type Struct struct {
	Name   string         `"struct" @Ident "{"`
	Fields []*StructField `{ @@ } "}"`
}

type StructField struct {
	Name string  `@Ident`
	Type string  `@Ident`
	Tag  *string `[ @String ]`
}

type Func struct {
	Name       string  `[ @Ident ]`
	Args       []*Arg  `[ "(" { @@ } ")" ]`
	ReturnType *Type   `[ ":" @@ ]`
	Body       []*Stmt `"-" ">" [ ( @@ ) | ("{" { @@ } "}") ]`
}

type Type struct {
	Array []string `{ @("[" "]") | @"*" }`
	Type  string   `@Ident`
}

type Arg struct {
	Name string `@Ident`
	Type *Type  `@@ [","]`
}

type Stmt struct {
	If             *If             `@@`
	For            *For            `| @@`
	Return         *Value          `| ("return" @@)`
	GoRoutine      *GoRoutine      `| @@`
	IdentOrVarDecl *IdentOrVarDecl `| @@`
	// Value             *Value             `| @@`
}

type GoRoutine struct {
	Func  *Func  `( "go" @@ )`
	Value *Value `| ( "go" @@ )`
}

type For struct {
	Iterator string  `"for" @Ident`
	Value    string  `[ "," @Ident ]`
	Source   string  `"in" @Ident "{"`
	Body     []*Stmt `{ @@ } "}"`
}

type If struct {
	Predicat *Predicat `"if" @@ "{"`
	Body     []*Stmt   `{ @@ } "}"`
	ElseIf   *ElseIf   `[ "else" @@ ]`
}

type ElseIf struct {
	If   *If   `( @@`
	Else *Else `| @@ )`
}

type Else struct {
	Body []*Stmt `"{" { @@ } "}"`
}

type Predicat struct {
	First    *Value            `@@`
	Operator *PredicatOperator `@@`
	Second   *Value            `@@`
}

type PredicatOperator struct {
	Eq  string `@(("=" "=") | "is")`
	Neq string `| @(("!" "=") | "isnt")`
	Gt  string `| @">"`
	Gte string `| @(">" "=")`
	Lt  string `| @"<"`
	Lte string `| @("<" "=")`
}

type IdentOrVarDecl struct {
	Ident   *NestedProperty `@@`
	VarDecl *VarDecl        `[ @@ ]`
}

type ArrAccess struct {
	Value *Value `"[" @@ "]"`
}

type VarDecl struct {
	Value *Value `"=" @@`
}

type NestedProperty struct {
	// Ref                 *string                `[ @"*" | @"&" ]`
	Ident               string                 `@Ident`
	ArrAccessOrFuncCall []*ArrAccessOrFuncCall `[ [ { @@ } ]`
	Nested              *NestedProperty        `[ "." @@ ] ]`
	// StructInst          *StructInst            `[ @@ ]`
	// Increment           *Increment             `[ @@ ]`
}

type ArrAccessOrFuncCall struct {
	ArrAccess *ArrAccess `( @@`
	FuncCall  *FuncCall  `| @@ )`
}

type FuncCall struct {
	Args []*Value `"(" { @@ [","] } ")"`
}

type Number struct {
	Float *float64 `( @Float`
	Int   *int64   `| @Int )`
}

type ParenthesisValue struct {
	Open  string     `@"("`
	Value *Operation `@@`
	Close string     `@")"`
}
type Value struct {
	Bool           *bool           `( @"true" | "false")`
	Operation      *Operation      `| @@`
	String         *string         `| @String`
	NestedProperty *NestedProperty `| @@`
	ArrDecl        *ArrDecl        `| @@`
}

type StructInst struct {
	Open  string `@"{"`
	Ident string `@Ident ":"`
	Value *Value `@@`
	Close string `@"}"`
}

type Operation struct {
	// ParenthesisValue *ParenthesisValue `@@ |`
	First  *Number    `( @@`
	Op     *Operator  `[ @@`
	Nested *Operation `@@ ] )`
}

type Operator struct {
	Plus  *string `( @"+"`
	Less  *string `| @"-"`
	Times *string `| @"*"`
	Div   *string `| @"/"`
	Mod   *string `| @"%" )`
}

type Increment struct {
	Inc *string `( @("+" "+")`
	Dec *string `| @("-" "-") )`
}

type ArrDecl struct {
	Type   string   `"[" "]" @Ident "{"`
	Values []*Value `{ @@ [ "," ] } "}"`
}

func Build(str string) (*INI, error) {
	parser, err := participle.Build(&INI{}, participle.UseLookahead())

	if err != nil {
		return &INI{}, err
	}

	ast := &INI{}

	err = parser.ParseString(str, ast)

	if err != nil {
		return &INI{}, err
	}

	return ast, nil
}

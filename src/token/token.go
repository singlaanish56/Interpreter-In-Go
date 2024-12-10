package token

const (
	//keyword
	FUNCTION="fn"
	LET="let"

	//literals
	VARIABLE="VAR"
	STRING="STR"
	NUMBER="INT"
	TRUE="T"
	FALSE="F"
	NULL="N"

	//brackets
	OROUNDBR="("
	CROUNDBR=")"
	OCURLYBR="{"
	CCURLYBR="}"
	OSQAUREBR="["
	CSQUAREBR="]"

	//signs
	SEMICOLON=";"
	COLON=":"
	EQUALTO="="
	UNDERSCORE="_"
	PLUS="+"
	COMMA=","

	//illegal
	INV="INVALID"
	EOF="EOF"
)

var KeywordMap = map[string]TokenType{
	"fn":FUNCTION,
	"let":LET,
	"true":TRUE,
	"false":FALSE,
	"null":NULL,
	"var":VARIABLE,
}

type TokenType string



type Token struct {
	Type          TokenType
	Identifier    string
	StartPosition int
	EndPosition   int
}

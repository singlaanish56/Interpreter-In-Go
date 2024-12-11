package token

const (
	//keyword
	FUNCTION="fn"
	LET="let"
	IF="if"
	ELSE="else"
	RETURN="return"

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
	OANGLEDBR="<"
	CANGLEDBR=">"
	
	//signs
	SEMICOLON=";"
	COLON=":"
	EQUALTO="="
	UNDERSCORE="_"
	PLUS="+"
	COMMA=","
	DOUBLEEQUALTO="=="
	EXCLAMATION="!"
	EXCLAMATIONEQUALTO="!="
	MINUS="-"
	DIVIDE="/"
	MULTIPLY="*"

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
	"if":IF,
	"else":ELSE,
	"return":RETURN,
}

type TokenType string



type Token struct {
	Type          TokenType
	Identifier    string
	StartPosition int
	EndPosition   int
}

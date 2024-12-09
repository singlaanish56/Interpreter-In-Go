package token

const (
	//keyword
	FUNCTION="fn"
	LET="let"

	//literals
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


	//illegal
	INV="INVALID"
	EOF="EOF"
)

var keywordMap = map[string]TokenType{
	"fn":FUNCTION,
	"let":LET,
}

type TokenType string



type Token struct {
	Type          TokenType
	Identifier    string
	startPosition int
	endPosition   int
}

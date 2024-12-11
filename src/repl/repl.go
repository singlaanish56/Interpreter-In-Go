package repl

import (
	"bufio"
	"fmt"

	"io"
	"github.com/singlaanish56/Interpreter-In-Go/token"
	"github.com/singlaanish56/Interpreter-In-Go/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader,out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned{
			return
		}

		input := scanner.Text()
		lexer := lexer.New(input)
		
		for tk:=lexer.GetToken(); tk.Type!=token.EOF;  tk=lexer.GetToken(){
			fmt.Printf("%+v\n", tk)
		}
	}
}

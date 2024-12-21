package repl

import (
	"bufio"
	"fmt"

	"io"

	"github.com/singlaanish56/Interpreter-In-Go/lexer"

	"github.com/singlaanish56/Interpreter-In-Go/parser"
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
		parser := parser.New(lexer)	
		
		program := parser.ParseProgram()
		if len(parser.Errors()) !=0{
			printParserErrors(out, parser.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out,"\n")
	}
}


func printParserErrors(out io.Writer, parserErrors []error){
	io.WriteString(out,"ran into these parser errors:\n")
	for _, err := range parserErrors{
		io.WriteString(out, "\t"+err.Error()+"\n")
	}
}

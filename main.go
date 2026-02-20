package main

import (
	"fmt"
	"parser/parsing"
	"github.com/antlr4-go/antlr/v4"
)

type MyCalculatorListener struct {
	parsing.BaseCalculatorListener // Embed the base listener
	stack                          []int
}

// Called when exiting an addSub rule
func (l *MyCalculatorListener) ExitAddSub(ctx *parsing.AddSubContext) {
	if (len(l.stack) < 2) {
		l.stack = append(l.stack, 0, 0)
	}
	right := l.stack[len(l.stack)-1]
	left := l.stack[len(l.stack)-2]
	l.stack = l.stack[:len(l.stack)-2]

	op := ctx.Op.GetText()
	switch op {
		case "+":
			l.stack = append(l.stack, left+right)
		default:
			fmt.Printf("Hello", op)
			panic("Hello")
			l.stack = append(l.stack, left-right)
	}
}


// EnterTermOnly is called when entering the termOnly production.
func (l *MyCalculatorListener) EnterTermOnly(c *parsing.TermOnlyContext) {

}

// EnterAddSub is called when entering the addSub production.
func (l *MyCalculatorListener) EnterAddSub(c *parsing.AddSubContext) {

}

// EnterFactorOnly is called when entering the factorOnly production.
func (l *MyCalculatorListener) EnterFactorOnly(c *parsing.FactorOnlyContext) {

}

// EnterMulDiv is called when entering the mulDiv production.
func (l *MyCalculatorListener) EnterMulDiv(c *parsing.MulDivContext) {

}

// EnterNumber is called when entering the number production.
func (l *MyCalculatorListener) EnterNumber(c *parsing.NumberContext) {

}

// EnterParentheses is called when entering the parentheses production.
func (l *MyCalculatorListener) EnterParentheses(c *parsing.ParenthesesContext) {

}

// ExitTermOnly is called when exiting the termOnly production.
func (l *MyCalculatorListener) ExitTermOnly(c *parsing.TermOnlyContext) {

}

// ExitFactorOnly is called when exiting the factorOnly production.
func (l *MyCalculatorListener) ExitFactorOnly(c *parsing.FactorOnlyContext) {

}

// ExitNumber is called when exiting the number production.
func (l *MyCalculatorListener) ExitNumber(c *parsing.NumberContext) {

}

// ExitParentheses is called when exiting the parentheses production.
func (l *MyCalculatorListener) ExitParentheses(c *parsing.ParenthesesContext) {

}

// Called when exiting an addSub rule
func (l *MyCalculatorListener) ExitMultSub(ctx *parsing.MulDivContext) {
	if (len(l.stack) < 2) {
		l.stack = append(l.stack, 0, 0)
	}
	right := l.stack[len(l.stack)-1]
	left := l.stack[len(l.stack)-2]
	l.stack = l.stack[:len(l.stack)-2]

	op := ctx.Op.GetText()
	switch op {
		case "+":
			l.stack = append(l.stack, left+right)
		default:
			fmt.Printf("Hello", op)
			panic("Hello")
			l.stack = append(l.stack, left-right)
	}
}

// Implement other Exit methods for different rules...

func (l *MyCalculatorListener) GetResult() int {
	if len(l.stack) > 0 {
		return l.stack[0]
	}
	return 0
}

func main() {
	// Create input stream
	input := antlr.NewInputStream("3 + 4 * 2")

	// Create lexer
	lexer := parsing.NewCalculatorLexer(input)

	// Create token stream
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create parser
	parser := parsing.NewCalculatorParser(tokens)

	// Parse the input
	tree := parser.Expression()

	// Print the parse tree (for debugging)
	fmt.Println(tree.ToStringTree([]string{}, parser))

	// Process the parse tree with a listener
	listener := &MyCalculatorListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	fmt.Printf("Result: %d\n", listener.GetResult())
}

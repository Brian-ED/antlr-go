package main

import (
	"fmt"
	"strconv"
	"strings"
	"parser/parsing"
	"github.com/antlr4-go/antlr/v4"
)

type MyCalculatorListener struct {
	parsing.BaseCalculatorListener // Embed the base listener
	stack                          []int
}

func StackMinLen(l *MyCalculatorListener, length int) {
	if (len(l.stack) < length) {
		println("Stack length: ", len(l.stack), "; Expected:", length)
		panic("Unexpected stack length")
	}
}

// Called when exiting an addSub rule
func (l *MyCalculatorListener) ExitAddSub(ctx *parsing.AddSubContext) {
	StackMinLen(l, 2)
	right := l.stack[len(l.stack)-1]
	left := l.stack[len(l.stack)-2]
	l.stack = l.stack[:len(l.stack)-2]

	op := ctx.Op.GetText()
	switch op {
		case "+":
			l.stack = append(l.stack, left+right)
		default:
			fmt.Printf("Op: ", op)
			panic("Unexpected operator")
			// l.stack = append(l.stack, left-right) // Old thing
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
	v, err := strconv.ParseInt(c.GetText(), 10, 32)
	if err != nil {
		println("token to parse:", c.GetText())
		panic("ParseInt errored while parsing an integer literal")
	}
	v2 := int(v);
	l.stack = append(l.stack, v2)
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
func (l *MyCalculatorListener) ExitMulDiv(ctx *parsing.MulDivContext) {
	StackMinLen(l, 2)
	right := l.stack[len(l.stack)-1]
	left := l.stack[len(l.stack)-2]
	l.stack = l.stack[:len(l.stack)-2]

	op := ctx.Op.GetText()
	switch op {
		case "*":
			l.stack = append(l.stack, left*right)
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

func Accept() {}
func ParseTree() {}

// PrettyPrintTree recursively prints an ANTLR parse tree with indentation.
// parser is needed to map rule indexes to rule names.
func PrettyPrintTree(node antlr.ParseTree, parser antlr.Parser, depth int) {
	indent := strings.Repeat("  ", depth) // 2 spaces per level

	switch n := node.(type) {
		case antlr.TerminalNode:
			// Terminal nodes: just print the token text
			text := strings.TrimSpace(n.GetText())
			if text != "" {
				fmt.Println(indent + text)
			}
		case antlr.RuleNode:
			// Non-terminal nodes: print the rule name
			ruleIndex := n.GetRuleContext().GetRuleIndex()
			ruleName := parser.GetRuleNames()[ruleIndex]
			fmt.Println(indent + ruleName)

			// Recurse into children
			for i := 0; i < n.GetChildCount(); i++ {
				switch v := n.GetChild(i).(type) {
					case antlr.ParseTree: PrettyPrintTree(v, parser, depth+1)
					default: panic("Unknown type")
				}
			}
		default: panic("Unhandled node")
	}
}

func main() {
	// Create input stream
	input := antlr.NewInputStream("3 + 5 * 2")

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

	PrettyPrintTree(tree, parser, 0)

	fmt.Printf("Result: %d\n", listener.GetResult())
}

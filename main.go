package main

import (
	"fmt"
	"parser/parsing"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// PYTHON: AUTOPLACE BEGIN
func (l *MyCalculatorListener) EnterTermOnly(ctx *parsing.TermOnlyContext)       {}
func (l *MyCalculatorListener) ExitTermOnly(ctx *parsing.TermOnlyContext)        {}
func (l *MyCalculatorListener) EnterParentheses(ctx *parsing.ParenthesesContext) {}
func (l *MyCalculatorListener) ExitParentheses(ctx *parsing.ParenthesesContext)  {}
func (l *MyCalculatorListener) EnterValueName(ctx *parsing.ValueNameContext)     {}
func (l *MyCalculatorListener) ExitValueName(ctx *parsing.ValueNameContext)      {}
func (l *MyCalculatorListener) EnterFuncBody(ctx *parsing.FuncBodyContext)       {}
func (l *MyCalculatorListener) ExitFuncBody(ctx *parsing.FuncBodyContext)        {}

// PYTHON: AUTOPLACE END

type runtimeType int

const (
	runtimeFunction runtimeType = iota
	runtimeInt
)

type MyCalculatorListener struct {
	parsing.BaseCalculatorListener // Embed the base listener
	stack                          []int
}

func StackMinLen(l *MyCalculatorListener, length int) {
	if len(l.stack) < length {
		println("Stack length: ", len(l.stack), "; Expected:", length)
		panic("Unexpected stack length")
	}
}

func printStack(x []int) {
	for _, j := range x {
		println(j)
	}
}

func (l *MyCalculatorListener) ExitCall(ctx *parsing.CallContext) {
	left := l.stack[0]
	right := l.stack[1:]

	op := strings.Join([]string{ctx.Op.GetText(), strconv.Itoa(len(right))}, "")
	switch op {
	case "*1":
		l.stack = []int{left*right[0]}
	case "+1":
		l.stack = []int{left+right[0]}
	case "+0":
		l.stack = []int{left+1}
	case "-0":
		l.stack = []int{-left}
	case "M1":
		l.stack = []int{left*5+right[0]}
	default:
		fmt.Println("Operator", op)
		panic("Dyadic call failed")
	}
}

// EnterNumber is called when entering the number production.
func (l *MyCalculatorListener) EnterNumber(c *parsing.NumberContext) {
	v, err := strconv.ParseInt(c.GetText(), 10, 32)
	if err != nil {
		println("token to parse:", c.GetText())
		panic("ParseInt errored while parsing an integer literal")
	}

	v2 := int(v)
	l.stack = append(l.stack, v2)
}

func (l *MyCalculatorListener) GetResult() int {
	if len(l.stack) > 0 {
		return l.stack[0]
	} else {
		panic("No result")
	}
}

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
			case antlr.ParseTree:
				PrettyPrintTree(v, parser, depth+1)
			default:
				panic("Unknown type")
			}
		}
	default:
		panic("Unhandled node")
	}
}

func main() {
	input := antlr.NewInputStream("3 + 5 M 5 * 2 + 2 + 1 - + + +")
	lexer := parsing.NewCalculatorLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parsing.NewCalculatorParser(tokens)
	tree := parser.Expression()
	fmt.Println(tree.ToStringTree([]string{}, parser))

	listener := &MyCalculatorListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	PrettyPrintTree(tree, parser, 0)

	fmt.Printf("Result: %d\n", listener.GetResult())
}

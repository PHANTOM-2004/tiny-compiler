package main

import (
	"fmt"
	"log"
	"os"
	"tj-compiler/cmd"

	"github.com/antlr4-go/antlr/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage:\n   %v <source_file>\n", os.Args[0])
		os.Exit(-1)
	}

	filename := os.Args[1]
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		log.Fatal(err)
	}

	lexer, lexerErrMsg := cmd.NewLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	fmt.Println("Token Stream:")
	// 获取所有 Token 并打印位置信息
	tokens.Fill() // 填充 Token 流
	for _, token := range tokens.GetAllTokens() {
		tokenType := token.GetTokenType() // Token 类型（数值）
		if tokenType == antlr.TokenEOF {
			break
		}

		// 获取 Token 信息
		line := token.GetLine()     // 行号（从 1 开始）
		column := token.GetColumn() // 列号（从 0 开始）
		text := token.GetText()     // 文本内容
		tokenName := lexer.SymbolicNames[tokenType]

		// 打印信息
		fmt.Printf("token(%2d:%2d) type=%7s, text='%s'\n",
			line, column, tokenName, text)
	}

	parser, parserErrMsg := cmd.NewParser(tokens)
	tree := parser.Prog()
	if len(lexerErrMsg()) != 0 || len(parserErrMsg()) != 0 {
		os.Exit(-1)
	}
	logrus.Info("basic check passed")

	fmt.Println("Parsing Tree:")
	fmt.Println(tree.ToStringTree(nil, parser))
}

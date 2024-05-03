package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countUniqueOperations(expression string) int {
	var stack []string
	uniqueExpressions := make(map[string]bool)
	expressionMap := make(map[string]string)
	operationCount := 0
	exprID := 0

	for i := 0; i < len(expression); i++ {
		ch := expression[i]
		switch {
		case ch == '(' || ch == '#' || ch == '$' || ch == '@':
			stack = append(stack, string(ch))
		case ch >= 'a' && ch <= 'z':
			stack = append(stack, string(ch))
		case ch == ')':
			subExpr := ""
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				subExpr = stack[len(stack)-1] + subExpr
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			if len(subExpr) > 0 {
				if _, exists := uniqueExpressions[subExpr]; !exists {
					uniqueExpressions[subExpr] = true
					operationCount++
					expressionMap[subExpr] = "expr" + strconv.Itoa(exprID)
					exprID++
				}
				stack = append(stack, expressionMap[subExpr])
			}
		}
	}

	return operationCount
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression := scanner.Text()

	result := countUniqueOperations(expression)
	fmt.Println(result)
}

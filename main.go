package main

import (
	"context"
	"fmt"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Finance Calucator",
		"1.0.0",
		server.WithResourceCapabilities(true, true), // Resource の機能で使われるオプションなのでToolの公開のみであれば不要そう
		server.WithLogging(),
	)

	// 四則計算ツールのインターフェース登録
	calculatorTool := mcp.NewTool("financial_calculator",
		mcp.WithDescription("Performs financial calculations."),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description(`The calculations to perform (Sinking Fund Factor, Capital Recovery Factor)
			- Sinking Fund Factor : This factor is used to calculate the annual deposit amount required to reach a specific future target sum over a certain period at a fixed interest rate.
			- Capital Recovery Factor : This factor is used to determine the amount needed to recover a certain capital investment over a specific period
			`),
			mcp.Enum("Sinking Fund Factor", "Capital Recovery Factor"),
		),
		mcp.WithNumber("r",
			mcp.Required(),
			mcp.Description("interest rate(0.0 ~ 1.0)"),
		),
		mcp.WithNumber("n",
			mcp.Required(),
			mcp.Description("number of periods(years)"),
		),
		mcp.WithNumber("amount",
			mcp.Required(),
			mcp.Description("amount"),
		),
	)

	// Add tool handler
	s.AddTool(calculatorTool, financeHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func financeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	op, ok := request.GetArguments()["operation"].(string)
	if !ok {
		return nil, fmt.Errorf("missing operation")
	}

	r, ok := request.GetArguments()["r"].(float64)
	if !ok {
		return nil, fmt.Errorf("missing r")
	}

	if r > 1 {
		r /= 10.0
	}

	n, ok := request.GetArguments()["n"].(float64)
	if !ok {
		return nil, fmt.Errorf("missing n")
	}

	amount, ok := request.GetArguments()["amount"].(float64)
	if !ok {
		return nil, fmt.Errorf("missing amount")
	}

	var rate, result float64
	switch op {
	case "Sinking Fund Factor":
		rate = r / (math.Pow(1+r, n) - 1)
		result = amount * rate
	case "Capital Recovery Factor":
		rate = (r * math.Pow(1+r, n)) / (math.Pow(1+r, n) - 1)
		result = amount * rate
	}

	return mcp.NewToolResultText(fmt.Sprintf("rate = %f n = %f amount = %f result rate = %f result is %d", r, n, amount, rate, int(result))), nil

}

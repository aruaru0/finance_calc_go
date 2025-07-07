package main

import (
	"context"
	"fmt"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Define constants for financial calculation types
const (
	FutureValueFactor           = "Future Value Factor"
	PresentValueFactor          = "Present Value Factor"
	FutureValueOfAnnuityFactor  = "Future Value of Annuity Factor"
	PresentValueOfAnnuityFactor = "Present Value of Annuity Factor"
	SinkingFundFactor           = "Sinking Fund Factor"
	CapitalRecoveryFactor       = "Capital Recovery Factor"
)

// Define a slice of available financial calculation operations
var financialOperations = []string{
	FutureValueFactor,
	PresentValueFactor,
	FutureValueOfAnnuityFactor,
	PresentValueOfAnnuityFactor,
	SinkingFundFactor,
	CapitalRecoveryFactor,
}

// main is the entry point for the Finance Calculator MCP server.
//
// The Finance Calculator performs the following financial calculations:
// 1. Future Value Factor: Calculates the future value of a present sum after a period of compound interest.
// 2. Present Value Factor: Calculates the present value of a future sum.
// 3. Future Value of Annuity Factor: Calculates the future value of a series of equal payments (annuity).
// 4. Present Value of Annuity Factor: Calculates the present value of a series of equal payments (annuity).
// 5. Sinking Fund Factor: Calculates the annual deposit required to reach a specific future sum.
// 6. Capital Recovery Factor: Calculates the payment required to recover an initial investment.
//
// The server is started with standard input/output.
func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Finance Calculator",
		"1.0.0",
		server.WithResourceCapabilities(true, true), // Options used for Resource features
		server.WithLogging(),
	)

	// Define the interface for the financial calculator tool
	calculatorTool := mcp.NewTool("financial_calculator",
		mcp.WithDescription("Performs financial calculations."),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description(`The financial calculation to perform. Select one of the following:
- Future Value Factor: Calculates the future value of a present sum after a period of compound interest.
- Present Value Factor: Calculates the present value of a future sum.
- Future Value of Annuity Factor: Calculates the future value of a series of equal payments (annuity).
- Present Value of Annuity Factor: Calculates the present value of a series of equal payments (annuity).
- Sinking Fund Factor: Calculates the annual deposit required to reach a specific future sum.
- Capital Recovery Factor: Calculates the payment required to recover an initial investment.`),
			mcp.Enum(financialOperations...),
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

	// Add a tool handler
	s.AddTool(calculatorTool, financeHandler)

	// Start the server with standard I/O
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

// calculateFinancialRate calculates various financial rates based on the given operation.
// Parameters:
//   - op: The type of financial calculation to perform. Supported operations include:
//     "Future Value Factor", "Present Value Factor", "Future Value of Annuity Factor",
//     "Present Value of Annuity Factor", "Sinking Fund Factor", and "Capital Recovery Factor".
//   - r: The interest rate expressed as a decimal (e.g., 0.05 for 5%).
//   - n: The number of periods (e.g., years) over which the calculation is performed.
//
// Returns:
// - The calculated rate as a float64 based on the specified operation.
// - An error if the provided operation is unknown.
func calculateFinancialRate(op string, r, n float64) (float64, error) {
	var rate float64
	// Branch calculation based on operation type
	switch op {
	case FutureValueFactor:
		rate = math.Pow(1+r, n)
	case PresentValueFactor:
		rate = math.Pow(1+r, -n)
	case FutureValueOfAnnuityFactor:
		rate = (math.Pow(1+r, n) - 1) / r
	case PresentValueOfAnnuityFactor:
		rate = (1 - math.Pow(1+r, -n)) / r
	case SinkingFundFactor:
		rate = r / (math.Pow(1+r, n) - 1)
	case CapitalRecoveryFactor:
		rate = (r * math.Pow(1+r, n)) / (math.Pow(1+r, n) - 1)
	default:
		return 0, fmt.Errorf("unknown operation: %s", op)
	}
	return rate, nil
}

// financeHandler is the handler for the financial tool.
// It takes four parameters:
// - operation: The type of financial calculation to perform.
// - r: The interest rate expressed as a decimal (e.g., 0.05 for 5%).
// - n: The number of periods (e.g., years) over which the calculation is performed.
// - amount: The amount on which the calculation is based.
// It returns the calculated financial rate and the result as an integer.
func financeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Get operation from arguments
	op, ok := request.GetArguments()["operation"].(string)
	if !ok {
		return nil, fmt.Errorf("missing operation")
	}

	// Get interest rate from arguments
	r, ok := request.GetArguments()["r"].(float64)
	if !ok {
		return nil, fmt.Errorf("missing r")
	}

	// If the interest rate is greater than 1, divide by 10
	if r > 1 {
		r /= 10.0
	}

	// Get number of periods from arguments
	n, ok := request.GetArguments()["n"].(float64)
	if !ok {
		return nil, fmt.Errorf("missing n")
	}

	// Get amount from arguments
	amount, ok := request.GetArguments()["amount"].(float64)
	if !ok {
		return nil, fmt.Errorf("missing amount")
	}

	// Calculate financial rate
	rate, err := calculateFinancialRate(op, r, n)
	if err != nil {
		return nil, err
	}

	// Calculate result
	result := amount * rate

	// Format and return the calculation result
	return mcp.NewToolResultText(fmt.Sprintf("rate = %f n = %f amount = %f result rate = %f result is %d", r, n, amount, rate, int(result))), nil
}

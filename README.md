# finance_calc_go

This is an MCP server that provides a financial calculator tool.

## Tool: `financial_calculator`

Performs financial calculations.

### Parameters:

- `operation` (string, required):
  The calculations to perform:
  - `Future Value Factor`: Calculates the future value of a present sum after a period of compound interest.
  - `Present Value Factor`: Calculates the present value of a future sum.
  - `Future Value of Annuity Factor`: Calculates the future value of a series of equal payments (annuity).
  - `Present Value of Annuity Factor`: Calculates the present value of a series of equal payments (annuity).
  - `Sinking Fund Factor`: Calculates the annual deposit required to reach a specific future sum.
  - `Capital Recovery Factor`: Calculates the payment required to recover an initial investment.
- `r` (number, required): Interest rate (0.0 ~ 1.0).
- `n` (number, required): Number of periods (years).
- `amount` (number, required): The amount for calculation.
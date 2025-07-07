# finance_calc_go

This is an MCP server that provides a financial calculator tool.

## Tool: `financial_calculator`

Performs financial calculations.

### Parameters:

- `operation` (string, required):
  The calculations to perform:
  - `Sinking Fund Factor`: This factor is used to calculate the annual deposit amount required to reach a specific future target sum over a certain period at a fixed interest rate.
  - `Capital Recovery Factor`: This factor is used to determine the amount needed to recover a certain capital investment over a specific period.
- `r` (number, required): Interest rate (0.0 ~ 1.0).
- `n` (number, required): Number of periods (years).
- `amount` (number, required): The amount for calculation.


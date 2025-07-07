# finance_calc_go

This is an MCP server that provides a financial calculator tool.

## Tool: `financial_calculator`

Performs financial calculations.

### Parameters:

- `operation` (string, required):
  The financial operation to perform. Must be one of:
  - "FVIF": Future value of a present lump sum
  - "PVIF": Present value of a future lump sum
  - "FVAIF": Future value of periodic savings
  - "PVAIF": Present value of periodic payments
  - "SFF": Required periodic savings to reach a future goal
  - "CRF": Fixed periodic payment to repay a loan or deplete a fund
- `r` (number, required): The annual interest rate as a decimal (e.g., 0.05 for 5%)
- `n` (number, required): The number of periods (typically years) for the calculation.
- `amount` (number, required):
  The monetary value used in the calculation:
  - For "FVIF" and "PVIF": a single lump-sum amount
  - For "FVAIF", "PVAIF", "SFF", and "CRF": a periodic (e.g. annual) amount
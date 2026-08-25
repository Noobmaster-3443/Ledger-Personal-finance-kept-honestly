# Ledger — Wallet System V6

Personal finance tracker in Thai Baht (฿), extended from V5 with a two-wallet asset system.

## Wallet System

- **Bank Wallet** — main spending/deposit wallet.
- **Savings Wallet** — separate savings wallet.
- Income and expense transactions can be assigned to a wallet.
- **Transfer** moves money between wallets without creating income or expense.
- Each wallet shows its own balance.
- **Total Assets = Bank + Savings**.
- Existing V5 transactions are automatically assigned to Bank Wallet on first run, so old data is preserved.
- Wallet opening balances can be edited from the Wallets tab.
- Transaction list can be filtered by wallet.
- CSV export includes wallet and transfer information.

## Balance formula

`Wallet Balance = Opening Balance + Income - Expense + Transfer In - Transfer Out`

Transfers therefore change the location of money but do not change Total Assets.

## Storage

The app continues to use browser `localStorage`, matching V5. Wallets and transfers are stored under their own keys. The Go binary only serves the embedded HTML.

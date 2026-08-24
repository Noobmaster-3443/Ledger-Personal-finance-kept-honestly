# Ledger

A tiny, single-binary personal finance tracker for tracking income and expenses in Thai Baht (฿). No install, no account, no cloud — it's a self-contained desktop app that runs a local web server and opens your browser to a clean, ledger-styled UI. All data stays on your machine in `localStorage`.

## Features

- **Add & track transactions** — log income and expenses with category, date, and notes
- **Monthly budget tracking** — set a monthly spending limit and get a progress bar with an alert once you cross 80%
- **Filters** — filter the transaction list by month, type (income/expense), and category
- **Spending by category** — daily breakdown plus a category pie chart (via Chart.js) for the bigger picture
- **CSV export** — export all transactions to CSV for backup or spreadsheet analysis
- **Light & dark mode** — including a "Liquid Glass" frosted, translucent dark theme
- **Fully local** — everything is stored in the browser's `localStorage`; no server, no accounts, no data leaving your machine

## Getting started

### Option 1: Run the prebuilt binary (Windows)

Download `Ledger.exe` and double-click it. A local server starts on `127.0.0.1:37841` and your default browser opens automatically to the app. Close the terminal window (or kill the process) to stop it.

> **Note:** Prebuilt `.exe` binaries are not tracked in this repo — see [Releases](../../releases) for downloads, or build from source below.

### Option 2: Build from source

Requires [Go](https://go.dev/dl/) 1.21+.

```bash
git clone https://github.com/Noobmaster-3443/ledger.git
cd ledger
go run .
```

This starts the app and opens it in your browser at `http://127.0.0.1:37841/`.

### Building a Windows executable

```bash
build-windows.bat
```

This cross-compiles a windowed (no console) `Ledger.exe` for `windows/amd64` using:

```bat
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-H windowsgui" -o Ledger.exe .
```

To build for your current platform instead, just run `go build .`.

## How it works

The app is a single Go binary that embeds `index.html` (a self-contained HTML/CSS/JS app styled with Tailwind CSS and charted with Chart.js) using `go:embed`. On launch it:

1. Starts a local HTTP server on `127.0.0.1:37841`
2. Serves the embedded UI at `/`
3. Opens your system's default browser to that address

All transaction data, budget settings, and preferences are stored client-side in `localStorage` — the Go server only serves static HTML, it never reads or writes your data.

## Project structure

```
.
├── main.go              # Go server: embeds and serves index.html, opens the browser
├── index.html            # The entire app UI, logic, and styling
├── build-windows.bat      # Cross-compile script for a windowed Windows .exe
└── package.json           # Placeholder / metadata (no Node runtime dependency)
```

## Tech stack

- **Backend:** Go (standard library only — `net/http`, `embed`, `os/exec`)
- **Frontend:** Vanilla HTML/CSS/JS, [Tailwind CSS](https://tailwindcss.com/) (Play CDN), [Chart.js](https://www.chartjs.org/)
- **Storage:** Browser `localStorage`

## Roadmap ideas

- Swap `localStorage` for a real backend/database
- Multi-currency support
- Recurring transactions
- Import from CSV/bank statements

## License

[MIT](LICENSE)

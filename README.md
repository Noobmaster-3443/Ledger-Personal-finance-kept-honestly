# 💎 Ledger — Personal Finance Tracker (v6.3 Pro)

<div align="center">

![Ledger Banner](https://img.shields.io/badge/Ledger-v6.3%20Pro-6366F1?style=for-the-badge&logo=wallet)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)
![Offline First](https://img.shields.io/badge/Offline--First-100%25-10B981?style=for-the-badge)
![Languages](https://img.shields.io/badge/Language-TH%20%7C%20EN-F59E0B?style=for-the-badge)

**A sleek, privacy-first, luxury fintech personal finance tracker in Thai Baht (฿) with persistent wallet separation, interactive analytics, and multi-language support.**

[Features](#-key-features) • [Wallet Architecture](#-wallet-system--formula) • [Quick Start](#-quick-start) • [Tech Stack](#-tech-stack) • [Data Privacy](#-data-privacy--backup)

</div>

---

## ✨ Key Features

### 💳 1. Dual-Wallet & Asset Separation System
* **Bank Wallet (🏦)**: Main spending and everyday cash flow account.
* **Savings Wallet (💰)**: Dedicated savings account kept separate from day-to-day spending.
* **Inter-Wallet Transfers (↔)**: Move money smoothly between wallets without creating artificial income or expense.
* **Dedicated Deposit Flow (`+ Deposit`)**: Deposit funds directly into a specific wallet with category and note tagging.
* **Initial Balance Configuration (`⚙`)**: Set or adjust opening balances via clean dedicated modals.
* **Luxury Titanium Card Aesthetics**: Styled like modern metallic cards with contactless NFC wave badge `📡` and reflection sheens.

### 🎨 2. Custom Category Personalization
* Create, customize, and delete your own expense & income categories.
* Quick color presets (Palette picker) for visual clarity.
* Seamless real-time sync with category filter dropdowns, transaction rows, and analytics charts.

### 🇹🇭 / 🇬🇧 3. Bilingual Support (i18n)
* Instant **Thai (TH)** and **English (EN)** language switcher in the header.
* Fully localized tabs, forms, buttons, alerts, status messages, and date formatting.
* Automatically remembers your preferred language in `localStorage`.

### 📊 4. Interactive Analytics & Visual Breakdown
* Dynamic Doughnut Chart powered by Chart.js for monthly category breakdowns.
* Color-coded category legends with total spending amounts and percentage contributions.

### 📅 5. Daily Overview & 7-Day Expense Tracker
* Daily net cash flow overview (`Income - Expense = Net`).
* Interactive 7-day mini calendar buttons to review recent daily spending at a glance.
* Category-by-category daily progress bars.

### 🛡️ 6. Emergency Fund Runway Calculator
* Automatically calculates your average monthly living expense.
* Interactive 3-month (lean) or 6-month (comfortable) coverage window slider.
* Real-time progress bar showing funding status and the remaining gap needed.

### ⚡ 7. Omnibar Quick Add
* Modern command-bar style input for rapid transaction logging.
* Smart parsing: type `"Lunch 60"` to instantly prefill Title and Amount.
* Keyboard shortcut enabled (`Enter ↵`).

### 🌓 8. Cyber Luxe Dual Themes
* **Light Mode (Warm Cashmere Slate)**: Soft, eyes-friendly matte slate (`#F4F6F9`) with crisp typography that prevents eye fatigue.
* **Dark Mode (Deep OLED Obsidian)**: Deep black-blue background (`#07090E`) with layered frosted glass panels and ambient nebula glows.

### 🔤 9. Modern Typography
* **Thai Typography**: `Prompt` (Google Fonts) — modern geometric loopless sans-serif widely used in modern fintech apps.
* **English Typography**: `Plus Jakarta Sans` & `Inter`.
* **Financial Figures**: `JetBrains Mono` for precise tabular numeric alignment.

### 🚀 10. Quality of Life (QoL) Suite
* **📋 Duplicate Transaction**: 1-click copy of any transaction to easily log recurring expenses (e.g. daily coffee, meals) with today's date.
* **🧠 Smart Memory**: Automatically remembers and pre-selects your last-used wallet, category, and type for seamless new entries.
* **⌨️ Global Keyboard Shortcuts**:
  * `N` or `+`: Open New Transaction modal from any tab.
  * `Ctrl + Enter` / `Cmd + Enter`: Submit & save active form.
  * `1` – `6`: Quick jump between tabs (Dashboard, Transactions, Wallets, Daily, Analytics, Emergency).
  * `Esc`: Close open modal or deselect rows.
* **📊 Filtered Live Summary**: Real-time summary bar displaying item count, total income, total expense, and net cash flow of filtered search results.
* **⚡ Quick Date Presets**: Instant 1-click date filters (`All Time`, `Today`, `Yesterday`, `This Week`, `This Month`, `Last Month`).
* **☑️ Bulk Selection & Batch Delete**: Select multiple rows with master checkbox and delete simultaneously in one confirmed batch.
* **🧮 Inline Math Calculator & Quick Adders**: Type arithmetic expressions directly into Amount inputs (`50+20+45`, `120*3`) with 1-tap quick increment pills (`+10`, `+50`, `+100`, `+500`, `+1k`).
* **🤖 Smart Auto-Category by Title**: Intelligently infers and auto-selects categories and transaction types based on keywords in your item title (e.g. food, gas, bills, shopping).
* **🔄 Multi-Criteria Sorting**: Sort records by Date (Newest/Oldest) or by Amount (Highest/Lowest) to pinpoint largest financial outflows in one click.

---

## 📐 Wallet System & Formula

Total Assets represent the true combined net worth across all your physical and digital accounts:

$$\text{Total Assets} = \text{Bank Wallet Balance} + \text{Savings Wallet Balance}$$

### Individual Wallet Balance Formula:
$$\text{Wallet Balance} = \text{Opening Balance} + \text{Income} - \text{Expense} + \text{Transfer In} - \text{Transfer Out}$$

> **Note**: Transferring money between wallets changes the distribution of your money across accounts, but **never changes your Total Assets** and never affects monthly income/expense totals.

---

## 🚀 Quick Start

### Method 1: Desktop Application (Windows)
1. Download or clone this repository.
2. Double-click **`Ledger-WalletSystem-V6-Persistent.exe`** to launch the standalone desktop app.

### Method 2: Web Browser (Any OS / Mobile / Tablet)
1. Open **`index.html`** in any modern web browser (Google Chrome, Microsoft Edge, Brave, Safari, Firefox).
2. No internet server or installation required — runs completely client-side!

---

## 💾 Data Privacy & Backup

* **100% Client-Side & Offline**: All transactions, transfers, categories, and settings are saved securely in your browser's `localStorage`.
* **Zero Telemetry**: No servers, no tracking, and no external data collection.
* **JSON Backup & Restore**: Export a single `.json` snapshot of your complete ledger anytime, or restore from a previous backup in 1-click.
* **CSV Export**: Export clean `.csv` spreadsheets of all transactions and transfers for Microsoft Excel, Google Sheets, or Apple Numbers.

---

## 🛠️ Tech Stack

* **Core**: Vanilla JavaScript (ES6+), HTML5, CSS3 Variables
* **Styling**: Tailwind CSS (JIT CDN) + Custom Cyber Luxe Design System
* **Typography**: Google Fonts (`Prompt`, `Plus Jakarta Sans`, `Inter`, `JetBrains Mono`)
* **Visualization**: Chart.js (UMD)
* **Desktop Wrapper**: C# / .NET WebView2 Launcher

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).


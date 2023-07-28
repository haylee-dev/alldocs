package func_ebay

func SetCurrencySymbol(CurrencyId string, GlobalId string) string {
	currency_symbol := "?"
	switch {
	case CurrencyId == "AUD":
		currency_symbol = "$"
	case CurrencyId == "CAD":
		currency_symbol = "£"
	case CurrencyId == "CHF":
		currency_symbol = "₣"
	case CurrencyId == "CNY":
		currency_symbol = "¥"
	case CurrencyId == "EUR":
		currency_symbol = "€"
	case CurrencyId == "GBP":
		currency_symbol = "£"
	case CurrencyId == "HKD":
		currency_symbol = "$"
	case CurrencyId == "INR":
		currency_symbol = "₹"
	case CurrencyId == "MYR":
		currency_symbol = "RM"
	case CurrencyId == "PHP":
		currency_symbol = "₱"
	case CurrencyId == "PLN":
		currency_symbol = "zł"
	case CurrencyId == "SEK":
		currency_symbol = "kr"
	case CurrencyId == "SGD":
		currency_symbol = "$"
	case CurrencyId == "TWD":
		currency_symbol = "$"
	case CurrencyId == "USD":
		currency_symbol = "$"
	case GlobalId == "EBAY-GB":
		currency_symbol = "£"
	}
	return currency_symbol
}

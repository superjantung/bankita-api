package util

const (
	USD = "USD"
	IDR = "IDR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, IDR:
		return true
	}
	return false
}

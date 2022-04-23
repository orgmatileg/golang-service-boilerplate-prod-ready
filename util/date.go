package util

import "golang_service/enum"

func ConvertDayOfTheLocalMonth(month int) string {
	var (
		result string
	)

	switch month {
	case int(enum.January):
		return "Januari"
	case int(enum.February):
		return "Februari"
	case int(enum.March):
		return "Maret"
	case int(enum.April):
		return "April"
	case int(enum.May):
		return "Mei"
	case int(enum.June):
		return "Juni"
	case int(enum.July):
		return "Juli"
	case int(enum.August):
		return "Agustus"
	case int(enum.September):
		return "September"
	case int(enum.October):
		return "Oktober"
	case int(enum.November):
		return "November"
	case int(enum.December):
		return "Desember"
	}

	return result
}

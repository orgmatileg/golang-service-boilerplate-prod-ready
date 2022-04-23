package enum

type monthOfTheYear int

const (
	MonthOfTheYearJanuary   monthOfTheYear = 1
	MonthOfTheYearFebruary  monthOfTheYear = 2
	MonthOfTheYearMarch     monthOfTheYear = 3
	MonthOfTheYearApril     monthOfTheYear = 4
	MonthOfTheYearMay       monthOfTheYear = 5
	MonthOfTheYearJune      monthOfTheYear = 6
	MonthOfTheYearJuly      monthOfTheYear = 7
	MonthOfTheYearAugust    monthOfTheYear = 8
	MonthOfTheYearSeptember monthOfTheYear = 9
	MonthOfTheYearOctober   monthOfTheYear = 10
	MonthOfTheYearNovember  monthOfTheYear = 11
	MonthOfTheYearDecember  monthOfTheYear = 12
)

func (c monthOfTheYear) GetMonthName() string {
	switch c {
	case MonthOfTheYearJanuary:
		return "Januari"
	case MonthOfTheYearFebruary:
		return "Februari"
	case MonthOfTheYearMarch:
		return "Maret"
	case MonthOfTheYearApril:
		return "April"
	case MonthOfTheYearMay:
		return "Mei"
	case MonthOfTheYearJune:
		return "Juni"
	case MonthOfTheYearJuly:
		return "Juli"
	case MonthOfTheYearAugust:
		return "Agustus"
	case MonthOfTheYearSeptember:
		return "September"
	case MonthOfTheYearOctober:
		return "Oktober"
	case MonthOfTheYearNovember:
		return "November"
	case MonthOfTheYearDecember:
		return "Desember"
	default:
		return "unknown"
	}
}

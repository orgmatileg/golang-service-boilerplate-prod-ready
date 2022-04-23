package enum

type Gender string

const (
	GenderMale   Gender = "M"
	GenderFemale Gender = "F"
)

func (t Gender) String() string {
	switch t {
	case GenderMale:
		return "Laki-laki"
	case GenderFemale:
		return "Perempuan"
	default:
		return "Unknown"
	}
}

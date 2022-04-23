package enum

type UserStatus int

const (
	UserStatusNotVerified int = 0
	UserStatusVerified    int = 1
	UserStatusBlocked     int = 2
)

func (u UserStatus) String() string {
	switch u {
	case 0:
		return "Unregistered"
	case 1:
		return "Registered"
	case 2:
		return "Blocked"
	default:
		return "Unknown"
	}
}

package exception

var (
	AUTH_EMPTY_AUTHORIZATION   = Code("AUTH4001")
	AUTH_NOT_VALID             = Code("AUTH4002")
	AUTH_EXPIRED               = Code("AUTH4003")
	AUTH_REFRESH_TOKEN_EXPIRED = Code("AUTH4004")
	AUTH_OTP_STILL_VALID       = Code("AUTH4005")
	AUTH_PIN_NOT_MATCH         = Code("AUTH4006")
	AUTH_PIN_NOT_SET           = Code("AUTH4007")
)

func init() {
	except[AUTH_EMPTY_AUTHORIZATION] = Exception{
		ErrorCode:  string(AUTH_EMPTY_AUTHORIZATION),
		Message:    "Authorization Header is empty",
		StatusCode: 401,
	}
	except[AUTH_NOT_VALID] = Exception{
		ErrorCode:  string(AUTH_NOT_VALID),
		Message:    "Authorization Header is not valid",
		StatusCode: 401,
	}
	except[AUTH_EXPIRED] = Exception{
		ErrorCode:  string(AUTH_EXPIRED),
		Message:    "Authorization expired",
		StatusCode: 401,
	}
	except[AUTH_REFRESH_TOKEN_EXPIRED] = Exception{
		ErrorCode:  string(AUTH_REFRESH_TOKEN_EXPIRED),
		Message:    "Sepertinya Kamu sudah lama tidak aktif, silahkan login kembali ya demi keamanan",
		StatusCode: 401,
	}
	except[AUTH_OTP_STILL_VALID] = Exception{
		ErrorCode:  string(AUTH_OTP_STILL_VALID),
		Message:    "OTP sebelum nya masih valid, Silahkan gunakan OTP tersebut",
		StatusCode: 400,
	}
	except[AUTH_PIN_NOT_MATCH] = Exception{
		ErrorCode:  string(AUTH_PIN_NOT_MATCH),
		Message:    "PIN tidak benar",
		StatusCode: 400,
	}
	except[AUTH_PIN_NOT_SET] = Exception{
		ErrorCode:  string(AUTH_PIN_NOT_SET),
		Message:    "PIN belum diatur",
		StatusCode: 400,
	}
}

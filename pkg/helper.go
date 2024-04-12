package pkg

import (
	"database/sql"
	"math/rand"
)

func NullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}

	return ""
}

func GenerateOTP() int {

	return rand.Intn(900000) + 100000
}

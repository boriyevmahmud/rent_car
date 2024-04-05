package check

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}

func ValidateEmail(email string) (bool, error) {
	boolean := strings.Contains(email, "@gmail.com")
	if !boolean {
		err := errors.New("error in validate email")
		return boolean, err
	}
	return boolean, nil
}

func ValidatePhone(phone string) bool {
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	if !strings.HasPrefix(phone, "998") {
		return false
	}

	return len(phone) == 12
}


func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[^a-zA-Z0-9\s]`).MatchString(password)

	if !hasLetter || !hasDigit || !hasSpecial {
		return errors.New("password must contain letters, number and symbol")
	}

	return nil
}
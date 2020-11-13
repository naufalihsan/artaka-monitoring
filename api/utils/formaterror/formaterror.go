package formaterror

import (
	"strings"
)

// FormatError
func FormatError(errString string) map[string]string {
	var errorMessages = make(map[string]string)

	if strings.Contains(errString, "username") {
		errorMessages["Taken_username"] = "Username Already Taken"
		return errorMessages
	}

	if strings.Contains(errString, "email") {
		errorMessages["Taken_email"] = "Email Already Taken"
		return errorMessages
	}

	if strings.Contains(errString, "phone") {
		errorMessages["Taken_title"] = "Phone Already Taken"
		return errorMessages
	}

	if strings.Contains(errString, "hashedPassword") {
		errorMessages["Incorrect_password"] = "Incorrect Password"
		return errorMessages
	}

	if strings.Contains(errString, "record not found") {
		errorMessages["No_record"] = "User Belum Terdaftar"
		return errorMessages
	}

	if len(errorMessages) > 0 {
		return errorMessages
	}

	errorMessages["Incorrect_details"] = "Incorrect Details"
	return errorMessages
}

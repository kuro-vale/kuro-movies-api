package tools

import "strings"

func FormatErr(err string) string {
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Email") {
		return "Email is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "email") && strings.Contains(err, "Email") {
		return "Invalid email format"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Password") {
		return "Password is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Email") {
		return "Email is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "min") && strings.Contains(err, "Password") {
		return "Password must have more than 3 characters"
	}
	if strings.Contains(err, "SQLSTATE") && strings.Contains(err, "unique") && strings.Contains(err, "email") {
		return "Email is already used"
	}
	if strings.Contains(err, "SQLSTATE") && strings.Contains(err, "varying(255)") {
		return "Fields must be under 255 characters"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Name") {
		return "Name is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Age") {
		return "Age is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "required") && strings.Contains(err, "Gender") {
		return "Gender is required"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "max") && strings.Contains(err, "Name") {
		return "Name must be under 50 characters"
	}
	if strings.Contains(err, "tag") && strings.Contains(err, "lte") && strings.Contains(err, "Age") {
		return "Age must be less than 90"
	}
	if strings.Contains(err, "SQ;STATE") && strings.Contains(err, "gender") {
		return "Gender must be male, female or x"
	}
	if strings.Contains(err, "cannot unmarshal") {
		return "Invalid field type"
	}
	return err
}

package validators

import "regexp"

// IsEmailValid validates an email address.
//
// This function checks if the given email address adheres to a standard email format
// using a regular expression. It also validates the email's length to ensure it falls
// within acceptable bounds.
//
// Parameters:
// - email: The email address to validate.
//
// Returns:
// - A boolean indicating whether the email address is valid (true) or invalid (false).
func IsEmailValid(email string) bool {
	// Regular expression to validate an email address format.
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{1,64}@[a-zA-Z0-9]" +
		"(?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9]" +
		"(?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	// Validate email length and check if it matches the regex pattern.
	if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email) {
		return false // Invalid email
	}

	return true // Valid email
}

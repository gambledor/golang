// Package palindrome provides palidrome function library
package palindrome

// IsPalindrome return true if the string is a palindrome.
// This implementation operates recursively by noting that all
// strings of length 0 or 1 are palindrome (the simple case)
// and that longer strings are palindrome only if their first
// and last character match and the remaining substring is a
// palindrome.
func IsPalindrome(str string) bool {
	strLength := len(str)
	if strLength <= 1 {
		return true
	} else {
		return str[0] == str[strLength-1] &&
			IsPalindrome(str[1:strLength-1])
	}
}

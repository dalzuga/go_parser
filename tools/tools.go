package tools

// SubStringInString
func SubStringInString(sub string, str string) bool {
	stringLength := len(str)
	subStringLength := len(sub)
	for i := 0; i < stringLength; i++ {
		if stringLength-i < subStringLength {
			return false
		}
		if sub == str[i:i+subStringLength] {
			return true
		}
	}
	return false
}

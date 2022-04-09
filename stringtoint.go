package myFunctions

// The algorithm for myAtoi(string s) is as follows:

// Read in and ignore any leading whitespace.
// Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
// Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
// Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
// If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
// Return the integer as the final result.

func myAtoi(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			s = s[i+1:]
			i = -1
		} else {
			break
		}
	}
	if len(s) == 0 {
		return 0
	}

	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	num := 0
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-48)
		} else {
			if num == 0 {
				return 0
			} else {
				break
			}
		}
		if num/10 > 214748364 || num/10 == 214748364 && s[i] > '7' {
			if sign == 1 {
				return 2147483647
			} else {
				return -2147483648
			}
		}
	}

	return num * sign
}

package secret

func Handshake(code uint) []string {
	var strings []string

	if (code | 0b00001) == code {
		strings = append(strings, "wink")
	}

	if (code | 0b00010) == code {
		strings = append(strings, "double blink")
	}

	if (code | 0b00100) == code {
		strings = append(strings, "close your eyes")
	}

	if (code | 0b01000) == code {
		strings = append(strings, "jump")
	}

	if (code | 0b10000) == code {
		for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
			strings[i], strings[j] = strings[j], strings[i]
		}
	}

	return strings
}

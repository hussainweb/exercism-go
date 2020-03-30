package secret

// Handshake translates the secret handshake code.
func Handshake(code uint) []string {
	var operations []string

	if code&1 != 0 {
		operations = append(operations, "wink")
	}
	if code&2 != 0 {
		operations = append(operations, "double blink")
	}
	if code&4 != 0 {
		operations = append(operations, "close your eyes")
	}
	if code&8 != 0 {
		operations = append(operations, "jump")
	}

	if code&16 != 0 {
		for i := 0; i < len(operations)/2; i++ {
			j := len(operations) - i - 1
			operations[i], operations[j] = operations[j], operations[i]
		}
	}

	return operations
}

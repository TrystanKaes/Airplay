package parameters

import "fmt"

func Marshal(tp *TextParameters) string {
	encoded := ""

	for _, key := range tp.Keys() {
		value := tp.Get(key)
		var line string

		if value == "" {
			line = key + "\n"
		} else {
			line = fmt.Sprintf("%s: %s\n", key, value)
		}

		encoded += line
	}

	return encoded
}

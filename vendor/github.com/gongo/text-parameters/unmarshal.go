package parameters

import (
	"bufio"
	"io"
	"regexp"
)

var parameterLinePattern = regexp.MustCompile("^([!-9;-~]+)(?:[ \t]*:[ \t]*(.+))?$")

func Unmarshal(r io.Reader) (*TextParameters, error) {
	params := &TextParameters{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		results := parameterLinePattern.FindStringSubmatch(line)
		if results == nil {
			return nil, &DecodeFormatError{line}
		}

		params.Set(results[1], results[2])
	}

	return params, nil
}

go-text-parameters
==================

[![GoDoc](http://godoc.org/github.com/gongo/text-parameters?status.svg)](http://godoc.org/github.com/gongo/text-parameters)
[![Build Status](https://travis-ci.org/gongo/text-parameters.svg)](https://travis-ci.org/gongo/text-parameters)
[![Coverage Status](https://coveralls.io/repos/gongo/text-parameters/badge.png)](https://coveralls.io/r/gongo/text-parameters)

Encoding and decoding of text/parameters written by Go

## Description

**text/parameters** is consists of either a list of parameters or a list of parameters and associated values.
Each entry of the list is a single line of text, and parameters are separated from values by a colon.

This package provides a read and write function of `text/parameters` format text from your Go programs.

### Spec

- [RFC 2326 - Real Time Streaming Protocol (RTSP)](http://tools.ietf.org/html/rfc2326)
- [draft-ietf-mmusic-rfc2326bis-40 - Real Time Streaming Protocol 2.0 (RTSP)](http://tools.ietf.org/html/draft-ietf-mmusic-rfc2326bis-40#page-297)

## Usage

### Decoder

```go
package main

import (
	"fmt"
	"strings"

	"github.com/gongo/text-parameters"
)

func main() {
	body := "Name: Wataru MIYAGUNI\nlogin: gongo\nAge: 30\n"

	u := struct {
		Name    string
		LoginId string `parameters:"login"`
		Age     int
	}{}

	decoder := parameters.NewDecorder(strings.NewReader(body))
	decoder.Decode(&u)

	fmt.Println("u.Name    =", u.Name)
	fmt.Println("u.LoginId =", u.LoginId)
	fmt.Println("u.Age     =", u.Age)
}
```

Output:

```
u.Name    = Wataru MIYAGUNI
u.LoginId = gongo
u.Age     = 30
```

### Encoder

```go
package main

import (
	"bytes"
	"fmt"

	"github.com/gongo/text-parameters"
)

func main() {
	var body bytes.Buffer

	u := struct {
		Name    string
		LoginId string `parameters:"login_id"`
		Rate    float64
	}{
		Name:    "Wataru MIYAGUNI",
		LoginId: "gongo",
		Rate:    0.923,
	}

	encoder := parameters.NewEncoder(&body)
	encoder.Encode(&u)

	fmt.Println(body.String())
}
```

Output:

```
Name: Wataru MIYAGUNI
Rate: 0.923
login_id: gongo
```

## Install

```
$ go get github.com/gongo/text-parameters
```

## LICENSE

[MIT License](./LICENSE.txt)

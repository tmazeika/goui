package main

import (
	. "github.com/tmazeika/goui/goui"
)

func main() {
	CreateApp(Div(A{"id", "my-id"})(
		Div(A{"class", "my-class"})(
			Str("Hello, world!"))))
}

package goui

import (
	"fmt"
	"strings"
)

type Node interface {
	toHtml() string
}

type A struct{
	Key string
	Value string
}

func (k A) String() string {
	// TODO
	return fmt.Sprintf("%s=%q", k.Key, k.Value)
}

type strNode string

func (s strNode) toHtml() string {
	return string(s)
}

func Str(s string) Node {
	return strNode(s)
}

func Div(attributes ...A) func(children ...Node) Node {
	attrStr := make([]string, len(attributes))
	for i, kv := range attributes {
		attrStr[i] = kv.String()
	}
	return func(children ...Node) Node {
		cStr := make([]string, len(children))
		for i, c := range children {
			cStr[i] = c.toHtml()
		}
		return Str(fmt.Sprintf("<div %s>%s</div>", strings.Join(attrStr, " "), strings.Join(cStr, "")))
	}
}

func CreateApp(node Node) {
	fmt.Println(node.toHtml())
}

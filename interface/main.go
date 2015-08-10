package main

import (
	"fmt"
)

type trimmedString string

func (t trimmedString) trim() trimmedString {
	return t[:3]
}

func main2() {
	var t trimmedString = "abcdef"
	fmt.Println(t.trim())
}

type accessor interface {
	getText() string
	setText(string)
}

type document struct {
	text string
}

func (d *document) getText() string {
	return d.text
}

func (d *document) setText(text string) {
	d.text = text
}

func main3() {
	var doc *document = &document{}
	doc.setText("document")
	fmt.Println(doc.getText())

	var acsr accessor = &document{}
	acsr.setText("document")
	fmt.Println(acsr.getText())
}

func main() {
	fmt.Println("-------------")
	main2()
	fmt.Println("-------------")
	main3()
}

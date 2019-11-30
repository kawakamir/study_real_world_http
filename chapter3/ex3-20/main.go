package main

import (
	"fmt"
	"golang.org/x/net/idna"
)

func main() {
	src := "握力王"
	ascii, err := idna.ToASCII(src)
	unicode, err := idna.ToUnicode(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s → ascii: %s or unicode: %s\n", src, ascii, unicode)
}

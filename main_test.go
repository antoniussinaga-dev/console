package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println(FontStyle("RED Background Yellow Padding 1").Red().Bold().Padding(1).BackgroundColor("yellow"))
}

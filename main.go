package main

import (
	"fmt"
	"io"
)

func main() {
	sth := (*io.Writer)(nil)
	fmt.Println(sth)
}

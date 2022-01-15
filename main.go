package main

import (
	"fmt"

	"github.com/gocraft/dbr/v2"
)

func main() {
	line := dbr.NewNullString("hellow wordl")
	fmt.Printf("%v", line)
}

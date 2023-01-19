package main

import (
	"fmt"

	"github.com/souliot/gotorch/version"
)

func main() {
	ver := version.String()
	fmt.Print(ver)
}

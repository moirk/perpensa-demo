package main

import (
	"fmt"

	"github.com/containers/podman/v5/version"
)

func main() {
	fmt.Println("podman library version:", version.Version.String())
}

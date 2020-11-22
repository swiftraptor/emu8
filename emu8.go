package main

import (
	"fmt"

	"github.com/swiftraptor/emu8/soc"
)

func main() {
	chip8 := soc.NewCpu()
	chip8.Cycle()
	fmt.Println("Initialised")
}

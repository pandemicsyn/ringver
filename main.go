package main

import (
	"fmt"
	"os"

	"github.com/gholt/ring"
)

func main() {
	rfiles := os.Args[1:]
	for _, v := range rfiles {
		r, _, err := ring.RingOrBuilder(v)
		if err != nil {
			panic(err)
		}
		if r != nil {
			fmt.Println(r.Version())
		} else {
			panic("Not a ring file")
		}
	}
}

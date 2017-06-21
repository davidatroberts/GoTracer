package main

import (
	"flag"
	"fmt"
)

func main() {
	wp := flag.Int("width", 400, "width of the image")
	hp := flag.Int("height", 200, "height of the image")
	flag.Parse()

	fmt.Printf("width: %d, height: %d\n", *wp, *hp)
}

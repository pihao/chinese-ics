package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pihao/ics/internal/hko"
)

var Version = "?"

func main() {
	v := flag.Bool("v", false, "show version.")
	c := flag.String("c", "", "calendar: hko.")
	flag.Parse()

	if *v {
		fmt.Println(Version)
		return
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	switch *c {
	case "hko":
		hko.GenSolarTerms()
	default:
		log.Fatal("calendar error")
	}
}

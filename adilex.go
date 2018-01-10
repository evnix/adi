package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var ip = flag.String("i", "", "Input file path")
	flag.Parse()
	if *ip == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	dat, err := ioutil.ReadFile(*ip)
	check(err)
	fmt.Print(string(dat))
}

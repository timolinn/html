package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"

	// html "github.com/timolinn/htmlparser"
	"golang.org/x/net/html"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	contents, err := ioutil.ReadFile("test.html")
	if err != nil {
		panic(err)
	}
	parsed := html.Parse(contents)
	fmt.Printf("%+v\n", parsed)
}

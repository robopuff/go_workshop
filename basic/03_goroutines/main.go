package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mytheresa/go-workshop/internal"
)

func main() {
	words := flag.String("words", "", "words separated by comma")
	flag.Parse()

	if *words == "" {
		fmt.Println("empty list provided")
	}

	reader := internal.NewEntriesReader()
	printer := internal.NewEntriesPrinter()

	for _, word := range strings.Split(*words, ",") {
		entries, err := reader.Read(word)
		if err != nil {
			fmt.Errorf(err.Error())
			continue
		}
		printer.Print(entries)
	}
}

package internal

import (
	"fmt"
	"strings"
)

type EntriesPrinter interface {
	Print(entries Entries)
}

type printer struct{}

func (p printer) Print(entries Entries) {
	for _, entry := range entries {
		fmt.Println(entry.Word)
		fmt.Println("-------")
		fmt.Println("Phonetics")
		for _, p := range entry.Phonetics {
			fmt.Printf("  %s\n", p.Text)
		}
		fmt.Println("-------")
		fmt.Println("Meanings")
		for _, m := range entry.Meanings {
			fmt.Printf(" Used as %s\n", m.PartOfSpeech)
			for _, d := range m.Definitions {
				fmt.Printf(" Definition:\n  %s\n", d.Definition)
				fmt.Printf(" Example:\n  %s\n", d.Example)
				if len(d.Synonyms) > 0 {
					fmt.Printf(" Synonyms:\n  %s", strings.Join(d.Synonyms, "\n  "))
				}
				fmt.Println()
			}
		}
		fmt.Println()
	}
}

func NewEntriesPrinter() EntriesPrinter {
	return &printer{}
}

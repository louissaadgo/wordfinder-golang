package main

import (
	"fmt"
	"os"
	"strings"
)

const core = `Lorem Ipsum is simply dummy text of the printing and typesetting industry.
Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown
printer took a galley of type and scrambled it to make a type specimen book. It has survived not only
five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was
popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more
recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

func main() {
	words := strings.Fields(strings.ToLower(core))
	query := os.Args[1:]
	for _, q := range query {
		q = strings.ToLower(q)
		for i, w := range words {
			if q == w {
				fmt.Printf("\"%v\": %q\n", i+1, w)
			}
		}
	}
}

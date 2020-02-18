package main

import "github.com/minskylab/gim"

func main() {
	g, err := gim.NewDefaultGim()
	if err != nil {
		panic(err)
	}

	if err := g.Run(); err != nil {
		panic(err)
	}
}

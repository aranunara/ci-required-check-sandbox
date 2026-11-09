package main

import "os"

func main() {
	if err := os.Remove("tmp.txt"); err != nil {
		panic(err)
	}
	if err := cleanup(); err != nil {
		panic(err)
	}
}

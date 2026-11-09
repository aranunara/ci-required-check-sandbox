package main

import "os"

func cleanup() {
	os.Remove("other.txt")
}

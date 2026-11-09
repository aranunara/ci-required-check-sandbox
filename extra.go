package main

import "os"

func cleanup() error {
	return os.Remove("other.txt")
}

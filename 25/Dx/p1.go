package main

func Part1(filename string) {
	f := ReadFile(filename)
	defer f.Close()
}

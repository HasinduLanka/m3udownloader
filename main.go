package main

var wsroot string = "workspace/"

func main() {

	NoConsole = true

	// RunTests()
	// return

	RunM3U()
}

func RunTests() {
	dir := wsroot + "output/"
	MakeDir(dir)

}

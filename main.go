package main

import (
	"os"
)

var wsroot string = "workspace/"

func main() {

	RunTests()
	return

	println()
	println("M3U Downloader")
	println("By github.com/HasinduLanka")
	println()

	NoConsole = true

	// filename := "playlist.m3u8"
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = Prompt("Enter file name : ")
		if len(filename) == 0 {
			filename = wsroot + "playlist.m3u8"
		}
	}

	println("Reading File " + filename)
	// dat, err := ioutil.ReadFile("/tmp/dat")
}

func RunTests() {
	// DownloadToFile(wsroot+"example.html", "https://example.com")
	s, err := LoadURIToString("https://example.com")
	CheckError(err)
	println(s)

	println("\n\n")

	s, err = LoadURIToString(wsroot + "playlist.m3u8")
	CheckError(err)
	println(s)

}

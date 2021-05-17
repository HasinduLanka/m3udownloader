package main

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RunM3U() {
	println()
	println("M3U Downloader")
	println("By github.com/HasinduLanka")
	println()

	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = Prompt("Enter file name or URL : ")
		if len(filename) == 0 {
			filename = wsroot + "playlist.m3u8"
		}
	}

	println("Reading File " + filename)
	file, lurierr := LoadURIToString(filename)
	CheckError(lurierr)

	keyMethod, keyURI := ParseKey(file)
	var key []byte

	if len(keyMethod) == 0 {
		r := PromptOptions("Key info not found", map[string]string{"p": "I'll provide key data", "n": "Do not use encryption"})
		switch r {
		case "p":
			keyMethod, key = PromptKeyData()
		case "n":
			keyMethod = ""
			key = nil
			keyURI = ""
		}

	} else {
		r := PromptOptions("Key found Method:"+keyMethod+" URI:"+keyURI, map[string]string{"r": "Read this key (Default)", "p": "I'll provide key data", "n": "Do not use encryption"})
		switch r {
		case "r":
			keyMethod, key = GetKeyData(keyURI)
			if len(keyMethod) == 0 {
				println("Reading Key from " + keyURI + " failed")
				keyMethod, key = PromptKeyData()
			}
		case "p":
			keyMethod, key = PromptKeyData()
		case "n":
			keyMethod = ""
			key = nil
			keyURI = ""
		}
	}

	UseAES := (len(keyMethod) != 0)

	if UseAES && !TestEncryptionKey(key) {
		panic(errors.New("encryption key failed the simple test"))
	} else {
		println("Key is looking good. Method:" + keyMethod)
	}

	// tstf, _ := LoadURI(wsroot + "chunk0.ts")
	// decrf := DecryptAES(key, tstf)
	// WriteFile(wsroot+"dchunk0.ts", decrf)

	var outFileName string

	if NoConsole {
		outputDir := wsroot + "output/"
		MakeDir(outputDir)
		outFileName = outputDir + "out.mkv"
	} else {
		outFileName = Prompt("Enter output file name : ")
	}

	os.Rename(outFileName, outFileName+".old")

	DecrypAndMerge(strings.Split(file, "\n"), outFileName, key)

}

func DecrypAndMerge(URIList []string, outFileName string, key []byte) error {
	for _, uri := range URIList {

		if strings.HasPrefix(uri, "#") {
			continue
		}

		println("Loading " + uri)
		IB, err := LoadURI(uri)
		if err != nil {
			return err
		}

		OB := DecryptAES(key, IB)
		AppendFile(outFileName, OB)

	}

	return nil

}

func GetKeyData(uri string) (string, []byte) {
	B, err := LoadURI(uri)
	if err != nil {
		return "", nil
	} else {
		if (len(B) % 8) == 0 {
			return "AES-" + strconv.Itoa(len(B)*8), B
		} else {
			println("Key is not in the correct length. Keys should be 16, 24, 32 long but this is " + strconv.Itoa(len(B)))
			return "", nil
		}
	}
}

func PromptKeyData() (string, []byte) {
	var r string

	if NoConsole {
		r = wsroot + "video.key"
	} else {
		r = Prompt("Enter key file name or URL : ")
	}

	M, B := GetKeyData(r)

	if len(M) == 0 {
		println("Sorry, I can't get that")
		return PromptKeyData()
	} else {
		println("Key loaded from " + r)
		return M, B
	}

}

var regex_key *regexp.Regexp = regexp.MustCompile("#EXT-X-KEY:METHOD=(.*),URI=(.*)")

// Returns : (Method string, Key string)
func ParseKey(c string) (string, string) {
	re := regex_key.FindStringSubmatch(c)
	if re == nil || len(re) != 3 {
		return "", ""
	}

	return re[1], re[2]

}

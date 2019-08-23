package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mr-tron/base58"
)

var helpFlags = map[string]bool{
	"-h":     true,
	"-help":  true,
	"--help": true,
}

const helpStr = `Encodes input as base58
Usage:
	-h		show help
	-d		decode from input
	-df [filename]	decode from file
`

func main() {
	if len(os.Args) == 2 && helpFlags[os.Args[1]] {
		_, _ = fmt.Fprint(os.Stderr, helpStr)
		os.Exit(2)
	}

	var source source = stdin{}
	if len(os.Args) == 3 && os.Args[1] == "-df" {
		source = file{os.Args[2]}
	}
	bs, err := source.read()
	if err != nil {
		fail(fmt.Errorf("could not read from %s: %v", source.name(), err))
	}

	if len(os.Args) == 1 {
		encode(bs)
		return
	}

	if os.Args[1] == "-d" || os.Args[1] == "-df" {
		decode(bs)
		return
	}

	fail(fmt.Errorf("unknown arguments: %v", os.Args[1:]))
}

type source interface {
	name() string
	read() ([]byte, error)
}

type stdin struct{}

func (s stdin) name() string {
	return "STDIN"
}

func (s stdin) read() ([]byte, error) {
	return ioutil.ReadAll(os.Stdin)
}

type file struct {
	fname string
}

func (f file) name() string {
	return f.fname
}

func (f file) read() ([]byte, error) {
	return ioutil.ReadFile(f.fname)
}

func encode(raw []byte) {
	_, err := io.WriteString(os.Stdout, base58.Encode(raw))
	if err != nil {
		fail(fmt.Errorf("could not write to STDOUT: %v", err))
	}
}

func decode(enc []byte) {
	str := strings.ReplaceAll(strings.ReplaceAll(string(enc), "\n", ""), "\r", "")

	raw, err := base58.Decode(str)
	if err != nil {
		fail(fmt.Errorf("could not decode from base58: %v", err))
	}
	_, err = os.Stdout.Write(raw)
	if err != nil {
		fail(fmt.Errorf("could not write to STDOUT: %v", err))
	}
}

func fail(err error) {
	_, _ = fmt.Fprint(os.Stderr, err)
	os.Exit(2)
}

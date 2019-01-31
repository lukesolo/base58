package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/mr-tron/base58"
)

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fail(fmt.Errorf("could not read from STDIN: %v", err))
	}

	if len(os.Args) == 1 {
		encode(bs)
		return
	}

	if os.Args[1] == "-d" {
		decode(bs)
		return
	}

	fail(fmt.Errorf("unknown arguments: %v", os.Args[1:]))
}

func encode(raw []byte) {
	_, err := io.WriteString(os.Stdout, base58.Encode(raw))
	if err != nil {
		fail(fmt.Errorf("could not write to STDOUT: %v", err))
	}
}

func decode(enc []byte) {
	raw, err := base58.Decode(string(enc))
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

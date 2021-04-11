package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ReadWholeFile(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}

func ReadFileLineByLine(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	ln := 0
	for s.Scan() {
		s := fmt.Sprintf("line %d: %s", ln, s.Text())
		fmt.Println(s)
		ln++
	}
}

func ReadFileChunk(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	buf := make([]byte, 16)
	r := bufio.NewReader(f)
	for {
		n, err := r.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatalln(err)
			}
			break
		}

		fmt.Println(string(buf[0:n]))
	}
}

func main() {
	ReadFileChunk("someFile.txt")
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const BufferSize = 1000
	file, err := os.Open("500kb.txt")
	check(err)

	defer file.Close()

	buffer := make([]byte, BufferSize)

	var overalTime float64 = 0

	for i := 0; i < 100; i++ {
		start := time.Now()
		for {
			bytesread, err := file.Read(buffer)

			if err != nil {
				if err != io.EOF {
					check(err)
				}
				break
			}

			if strings.Contains(string(buffer[:bytesread]), "RTCPeerConnection") {
				fmt.Println("true")
			}

		}
		end := time.Now()

		elapsed := end.Sub(start)

		overalTime += elapsed.Seconds()
		_, err = file.Seek(0, 0)
		check(err)
	}

	log.Print(overalTime / 100)
}

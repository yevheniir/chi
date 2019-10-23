package src

import (
	"bufio"
	"log"
	"os"
)

func ScanAndSend(path string, send Sender, gen msgGenerator) int {
	file, err := os.Open(path)
	count := 0
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++

		if count%10000 == 0 {
			log.Printf("Sended: %d lines", count)
		}

		err = send(gen(scanner.Text()))

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return count
}

package src

import (
	"bufio"
	"log"
	"os"
)

func ScanAndSend(path string, send Sender) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err = send(parseMetric(scanner.Text()))

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

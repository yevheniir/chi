package src

import (
	"fmt"
	"net"
)

func Send(s string) {
	fmt.Println(s)
}

func GetSender(c []net.Conn) func(string) error {
	currentConn := 0

	return func(s string) error {
		// defer c.Close()
		fmt.Printf("Send: %s", s)
		_, err := fmt.Fprintf(c[currentConn], s)

		if currentConn == len(c)-1 {
			currentConn = 0
		} else {
			currentConn++
		}

		return err

	}

}

type Sender = func(string) error

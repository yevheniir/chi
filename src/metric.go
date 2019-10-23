package src

import (
	"fmt"
	"strings"
)

// func GenMessageTCP(s string) string {
// 	return "107ddd2e-0dbf-458c-8bbc-5019bdcfd501." + strconv.Itoa(num) + ".worker." + name + " " + strconv.Itoa(j) + "\n"
// }

func parseTCP(metric string) string {
	return parseUDP(metric) + "\n"
}

func parseUDP(metric string) string {
	x := strings.Split(metric, "[")[2]
	x = strings.TrimSuffix(x, "]]")

	args := strings.Split(x, ",")
	fmt.Printf("metric: %s\n", args)

	return strings.ReplaceAll(args[0], "\"", "") + "." + strings.ReplaceAll(strings.ReplaceAll(args[1], "\"", ""), " ", "") + args[2] + strings.ReplaceAll(args[3], "\n", "")
}

func GetMsgGenerator(protocol string) msgGenerator {
	switch protocol {
	case "tcp":
		return parseTCP
	case "udp":
		return parseUDP
	default:
		return parseTCP
	}
}

type msgGenerator = func(string) string

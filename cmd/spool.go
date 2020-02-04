/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/spf13/cobra"
	"github.com/yevheniir/chi/src"
)

var Path string

// spoolCmd represents the spool command
var spoolCmd = &cobra.Command{
	Use:   "spool",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		startTime := time.Now()
		protocol, _ := cmd.Flags().GetString("protocol")
		address, _ := cmd.Flags().GetString("address")
		conn, _ := cmd.Flags().GetString("connections")

		var conns []net.Conn
		for range conn {
			c, err := net.Dial(protocol, address)
			fmt.Printf("Connected")

			if err != nil {
				fmt.Printf("Oh nooo: %v\n", err)
				return
			}

			conns = append(conns, c)
		}

		sender := src.GetSender(conns)
		genMessage := src.GetMsgGenerator(protocol)

		count := src.ScanAndSend(args[0], sender, genMessage)

		elapsed := time.Since(startTime)
		time := float32(float32(elapsed) / float32(time.Second))
		log.Printf("Time spended: %s", elapsed)
		log.Printf("Metric puts/s: %f", float32(count)/time)

		for _, c := range conns {
			c.Close()
		}

	},
}

func init() {
	rootCmd.AddCommand(spoolCmd)
	spoolCmd.PersistentFlags().StringP("path", "p", "/spool.spool", "Path to spool file")
	spoolCmd.MarkFlagRequired("path")

	spoolCmd.Flags().StringP("protocol", "r", "tcp", "Protocol")

	spoolCmd.Flags().StringP("address", "a", "localhost:2003", "Network path")

	spoolCmd.Flags().StringP("connections", "c", "1", "How many connections")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spoolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spoolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

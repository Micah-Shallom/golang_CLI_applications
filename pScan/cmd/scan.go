/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Micah-Shallom/modules/scan"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run a port scan on the hosts",
	RunE: func(cmd *cobra.Command, args []string) error {
		// hostsFile, err := cmd.Flags().GetString("hosts-file")
		// if err != nil {
		// 	return err
		// }
		ports, err := cmd.Flags().GetIntSlice("ports")
		if err != nil {
			return err
		}
		portRange, err := cmd.Flags().GetString("portRange")
		if err != nil {
			return err
		}
		filter, err := cmd.Flags().GetString("filter")
		if err != nil {
			return err
		}
		// portRange := viper.GetString("portRange")
		// ports := viper.GetIntSlice("ports") //didnt use viper as it erases the default port setup
		hostsFile := viper.GetString("hosts-file")
		return scanAction(os.Stdout, hostsFile, ports, portRange,filter)
	},
}

func scanAction(out io.Writer, hostsFile string, ports []int, portRange string, filter string) error {
	hl := &scan.HostLists{}
	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	//disable ability to pass both ports and portRange
	if (len(ports) > 3) && (portRange != "") {
		flagErr := errors.New("error: Specify either ports or portRange and not both")
		return flagErr
	}

	//if portRange is provided loop through it and populate ports
	if portRange != "" {
		portStr := strings.Split(portRange, "-")
		start, err := strconv.Atoi(portStr[0])
		if err != nil {
			fmt.Println("Error converting start:", err)
			return err
		}
		end, err := strconv.Atoi(portStr[1])
		if err != nil {
			fmt.Println("Error converting end:", err)
			return err
		}
		if (start >= 1 && end <= 65535) && (end > start) {
			ports = []int{}
			for i := start; i < end; i++ {
				ports = append(ports, i)
			}
		} else {
			flagErr := errors.New("error: port range should be between 1-65535 | upper port number must be greater than lower port number")
			return flagErr
		}
	}

	results := scan.Run(hl, ports)
	return printResults(out, results, filter)
}

func printResults(out io.Writer, results []scan.Results, filter string) error {
	message := ""
	for _, result := range results {
		message += fmt.Sprintf("%s:", result.Host)

		if result.NotFound {
			message += fmt.Sprint("Host not found\n\n")
			continue
		}
		message += fmt.Sprintln()
		//looping through tcp ports
		for _, port := range result.PortStates {
			if port.TCPOpen.String() == filter {
				message += fmt.Sprintf("\t%d TCP: %s\n", port.Port, port.TCPOpen )
			}
			// continue
		}
		message += fmt.Sprintln()
		
		//looping through udp ports
		for _, port := range result.PortStates {
			if port.TCPOpen.String() == filter {
				message += fmt.Sprintf("\t%d UDP: %s\n", port.Port, port.UDPOpen)
			}
		}
		message += fmt.Sprintln()
	}
	_, err := fmt.Fprint(out, message)
	return err
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().IntSliceP("ports", "p", []int{22, 80, 443}, "ports to scan") //sets default scan command
	scanCmd.Flags().StringP("portRange", "r", "80-82", "range of ports to scan")
	//add filter for open/closed ports
	scanCmd.Flags().StringP("filter","t","both","filter open or closed ports(both, opened, closed)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

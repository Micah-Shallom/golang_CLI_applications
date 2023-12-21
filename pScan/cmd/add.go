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
	"fmt"
	"io"
	"os"

	"github.com/Micah-Shallom/modules/scan"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add <host1>...<hostn>",
	Aliases: []string{"a"},
	Short:   "add new host to list",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}
		return addAction(os.Stdout, hostsFile, args)
	},
	SilenceUsage: true,
}

func addAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostLists{}

	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	for _, host := range args {
		if err := hl.Add(host); err != nil {
			return err
		}
		fmt.Fprintln(out, "Added host:", host)
	}
	return hl.Save(hostsFile)
}


func init() {
	hostsCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"strings"

	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/patchbookGo/parser"
	"github.com/morgulbrut/patchbookGo/targets"
	"github.com/morgulbrut/patchbookGo/utils"
	"github.com/spf13/cobra"
)

// dotCmd represents the dot command
var dotCmd = &cobra.Command{
	Use:   "dot filename [output filename]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		Dot(args)
	},
}

func init() {
	generateCmd.AddCommand(dotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Dot generates a dot file and returns its name
func Dot(args []string) string {

	if len(args) < 1 || len(args) > 2 {
		colorlog.Fatal("Wrong number of arguments")
	}
	infn := args[0]
	var outfn string
	if len(args) == 1 {
		outfn = strings.Split(infn, ".")[0] + ".dot"
	} else {
		outfn = args[1]
	}
	patch := parser.File(utils.ReadFile(infn))
	dot := targets.Graphviz(patch)
	colorlog.Debug("Writing dot")
	err := ioutil.WriteFile(outfn, []byte(dot), 0644)
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	return outfn
}

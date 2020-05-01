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
	"os"
	"os/exec"
	"strings"

	"github.com/morgulbrut/colorlog"
	"github.com/spf13/cobra"
)

// pngCmd represents the png command
var pngCmd = &cobra.Command{
	Use:   "png",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		png(args)
	},
}

func init() {
	generateCmd.AddCommand(pngCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pngCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pngCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func png(args []string) {
	dotFn := Dot(args)
	pngFn := strings.Split(dotFn, ".")[0] + ".png"
	colorlog.Debug("Running dot")
	cmd := exec.Command("dot", "-Tpng", dotFn)
	png, err := cmd.Output()
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	colorlog.Debug("Writing png")
	err = ioutil.WriteFile(pngFn, []byte(png), 0644)
	if err != nil {
		colorlog.Fatal(err.Error())
	}
	os.Remove(dotFn)
}

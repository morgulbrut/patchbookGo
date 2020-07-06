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
package main

import "github.com/morgulbrut/patchbookGo/cmd"

func main() {
	cmd.Execute()
}

// func main() {
// 	flag.StringVar(&f, "f", "", "open file")
// 	flag.Parse()

// 	buf, err := ioutil.ReadFile(f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	s := string(buf)

// 	re := regexp.MustCompile(`\n(\s*)\|`)
// 	s = re.ReplaceAllString(s, " |")

// 	re2 := regexp.MustCompile(`\:(\s*)\|`)
// 	s = re2.ReplaceAllString(s, ":")

// 	patch := parser.File(strings.Split(s, "\n"))

// 	// fmt.Println(patch)

// 	dot := targets.Graphviz(patch)
// 	// fmt.Println(targets.Graphviz(patch))

// 	err = ioutil.WriteFile("patch.txt", []byte(dot), 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

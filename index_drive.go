package main

import (
	"flag"
	"fmt"
	"github.com/PeterHickman/toolbox"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
)

var crc string

type data struct {
	Name  string
	Error string
	Size  int64
	Mode  string
	Crc   string
}

func init() {
	var c = flag.String("crc", "", "Calculate the md5 or sha256 of a file")

	flag.Parse()

	crc = strings.ToLower(*c)

	if crc != "" && crc != "md5" && crc != "sha256" {
		fmt.Printf("Unknown crc method [%s]. Only md5 and sha256 implemented\n", crc)
		os.Exit(7)
	}

	if len(flag.Args()) != 1 {
		println("Enter a directory to scan")
		os.Exit(9)
	}
}

func main() {
	root := flag.Arg(0)

	output := []data{}

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			error_value := "none"
			if err != nil {
				error_value = err.Error()
			}

			crcvalue := "none"
			if !info.IsDir() {
				if crc == "md5" {
					crcvalue = toolbox.CalculateMD5(path)
				} else if crc == "sha256" {
					crcvalue = toolbox.CalculateSHA256(path)
				}
			}

			output = append(output, data{path, error_value, info.Size(), fmt.Sprint(info.Mode()), crcvalue})

			return nil
		})

	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	d, err := yaml.Marshal(output)
	if err != nil {
		fmt.Println(err)
		os.Exit(7)
	}
	fmt.Println(string(d))
}

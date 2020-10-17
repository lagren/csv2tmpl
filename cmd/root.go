package cmd

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var inputFilename string
var templateString string
var templateFile string
var headers string
var headerNames []string
var headerRow bool

var rootCmd = &cobra.Command{
	Use: "csv2tmpl",
	Run: func(cmd *cobra.Command, args []string) {
		if len(headers) > 0 {
			headerNames = strings.Split(headers, ",")
		}

		tmpl := getTemplate()

		tmpl.Name()

		f, err := os.Open(inputFilename)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		c := csv.NewReader(f)

		if headerRow {
			headerNames, err = c.Read()
			if err != nil {
				panic(err)
			}
		}

		for {
			r, err := c.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}

			data := make(map[string]string)

			if len(headerNames) > 0 {
				for idx, value := range r {
					data[headerNames[idx]] = value
				}
			} else {
				for idx, _ := range r {
					data[fmt.Sprintf("col%d", idx)] = r[idx]
				}
			}

			//fmt.Println(r)
			//fmt.Println(data)

			out := bytes.NewBuffer(nil)

			err = tmpl.Execute(out, data)
			fmt.Println(out.String())
			if err != nil {
				panic(err)
			}

		}

	},
}

func getTemplate() *template.Template {
	var err error
	tmpl := template.New("tmpl")

	tmpl = tmpl.Funcs(map[string]interface{}{
		"upper": toUpper,
		"lower": toLower,
		"md5":   toMd5,
	})

	if templateString != "" {
		tmpl, err = tmpl.Parse(templateString)
	} else if templateFile != "" {
		tmpl, err = tmpl.ParseFiles(templateFile)
	} else {
		panic("missing template or template-file")
	}

	if err != nil {
		panic(err)
	}

	tmpl.Funcs(map[string]interface{}{})

	return tmpl
}

func init() {
	rootCmd.Flags().StringVarP(&inputFilename, "input", "i", "", "input CSV file")

	rootCmd.Flags().StringVarP(&templateString, "template", "t", "", "template string to use")
	rootCmd.Flags().StringVar(&templateFile, "template-file", "", "template file to use")

	rootCmd.Flags().StringVar(&headers, "headers", "", "header names (comma separated list)")
	rootCmd.Flags().BoolVar(&headerRow, "header-row", false, "use first row as header names")

	rootCmd.MarkFlagRequired("input")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

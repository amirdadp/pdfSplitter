/*
Copyright © 2023 Amir H. Dadpour me@adadpour.site
*/
package cmd

import (
	"github.com/amirdadp/go-ghostscript/ghostscript"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var (
	outputFile string
	deleteFlag bool
	rootCmd    = &cobra.Command{
		Use:   "pdfSplitter",
		Short: "A simple PDF Splitter, use -h, --help for usage info",
		Long: `pdfSplitter; Provide at least two arguments
The first one should be the relative address of the PDF you want to split followed by any number of numbers as the page separators.`,

		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]
			startPage := 1
			args = args[1:]
			fpTemplate := "-dFirstPage="
			lpTemplate := "-dLastPage="
			outTemplate := "-sOutputFile="

			gsArgs := []string{
				"gs",
				"-sDEVICE=pdfwrite",
				"-dNOPAUSE",
				"-dBATCH",
				"-dSAFER",
				"-dFirstPage=",  //#5
				"-dLastPage=",   //#6
				"-sOutputFile=", //#7
				fileName,
			}
			for counter, endPage := range args {
				rev, err := ghostscript.GetRevision()
				if err != nil {
					panic(err)
				}
				log.Printf("Revision: %+v\n", rev)
				gs, err := ghostscript.NewInstance()
				if err != nil {
					panic(err)
				}
				startPageStr := strconv.Itoa(startPage)

				gsArgs[5] = fpTemplate + startPageStr
				gsArgs[6] = lpTemplate + endPage
				gsArgs[7] = outTemplate + "\"" + strings.Replace(outputFile, "{}", strconv.Itoa(counter+1), 1) + "\""
				if err := gs.Init(gsArgs); err != nil {
					panic(err)
				}

				tmp, errr := strconv.Atoi(endPage)
				if errr != nil {
					panic(errr)
				}
				startPage = tmp + 1

				func() {
					gs.Exit()
					gs.Destroy()
				}()

			}

			if deleteFlag {
				err := os.Remove(fileName)
				if err != nil {
					panic(err)
				}
			}

		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "Part{}.pdf", "The name of the output files, must contain '{}' as a placeholder for the counter.")
	rootCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", false, "Include if you want the original file to be deleted.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// sub commands

}

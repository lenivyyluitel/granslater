package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	granslater "gitlab.com/lenivyyluitel/granslater/pkg"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "granslater",
	Short: "Translate text using libretranslate",
	Run:   translate,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func translate(cmd *cobra.Command, args []string) {

	var fullArgument string

	sourceFlag, err := cmd.Flags().GetString("source")
	if err != nil {
		fmt.Println(err)
	}
	outputFlag, _ := cmd.Flags().GetString("output")
	fileFlag, _ := cmd.Flags().GetString("file")
	debugFlag, _ := cmd.Flags().GetBool("debug")

	// opens the file if fileFlag (--f)
	if fileFlag != "" {
		file, err := os.Open(fileFlag)

		if err != nil {
			log.Panic("Error: ", err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fullArgument = scanner.Text()
		}

	} else {
		for i := range args {
			fullArgument += args[i] + " "
		}
	}

	values := granslater.Translate(fullArgument, sourceFlag, outputFlag, debugFlag)
	fmt.Println(values)
}

func init() {
	rootCmd.Flags().StringP("source", "s", "en", "Source language")
	rootCmd.Flags().StringP("output", "o", "ru", "Output language")
	rootCmd.Flags().StringP("file", "f", "", "Translate a text file")
	rootCmd.Flags().BoolP("debug", "d", false, "Toggle debug log")
}

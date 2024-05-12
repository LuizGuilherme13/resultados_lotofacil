package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	log.Println("Initializing...")
}

var rootCmd = &cobra.Command{
	Use: "loto",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

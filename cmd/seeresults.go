package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/LuizGuilherme13/resultados_lotofacil/models"
	"github.com/adhocore/chin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(SeeResultsCmd)

	SeeResultsCmd.Flags().BoolP("latest", "l", false, "Latest draw result.")
}

var SeeResultsCmd = &cobra.Command{
	Use:   "see",
	Short: "Displays the results in the terminal.",
	Run:   SeeResults,
}

func SeeResults(cmd *cobra.Command, args []string) {
	if _, err := os.Stat("./results.json"); errors.Is(err, os.ErrNotExist) {
		log.Println("The file with the results does not exist or was not found.")
		fmt.Println("Run 'loto get' to fetch and generate the file.")
		os.Exit(1)
	}

	s := chin.New()
	go func() {
		fmt.Print("Searching... ")
		s.Start()
	}()

	content, err := os.ReadFile("./results.json")
	if err != nil {
		log.Fatalln(err)
	}

	results := []models.DailyResult{}
	err = json.Unmarshal(content, &results)
	if err != nil {
		log.Fatalln(err)
	}

	s.Stop()
	fmt.Printf("Successful! Quantity found: %d\n", len(results))

	for _, result := range results {
		fmt.Printf("Date: %s - Dozens: %s\n", result.Data, result.Dezenas)
	}
}

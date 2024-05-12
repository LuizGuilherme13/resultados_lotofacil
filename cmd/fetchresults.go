package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LuizGuilherme13/resultados_lotofacil/models"
	"github.com/adhocore/chin"
	"github.com/spf13/cobra"
)

const baseURL = "https://loteriascaixa-api.herokuapp.com/api/lotofacil"

func init() {
	rootCmd.AddCommand(FetchResultsCmd)

	initialDefault := time.Now().Add(-168 * time.Hour).Format("02/01/2006")
	finalDefault := time.Now().Format("02/01/2006")

	FetchResultsCmd.Flags().BoolP("latest", "l", false, "Latest draw result.")
	FetchResultsCmd.Flags().StringP("initial", "i", initialDefault, "Initial date.")
	FetchResultsCmd.Flags().StringP("final", "f", finalDefault, "Final date.")
}

var FetchResultsCmd = &cobra.Command{
	Use:   "get",
	Short: "Searches the results and generates the file.",
	Run:   fetchResults,
}

func fetchResults(cmd *cobra.Command, args []string) {
	s := chin.New()
	go func() {
		fmt.Print("Searching... ")
		s.Start()
	}()

	body, err := httpGetResults(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	lotoResults := []models.LotoResult{}
	err = json.Unmarshal(body, &lotoResults)
	if err != nil {
		log.Fatalln(err)
	}

	results, err := filterByPeriod(cmd, lotoResults)
	if err != nil {
		log.Fatalln(err)
	}

	if err := os.WriteFile("results.json", results, 0644); err != nil {
		log.Fatalln(err)
	}

	s.Stop()
	fmt.Println("Successful!")
	fmt.Println("Run 'loto see' to see results.")
}

func httpGetResults(cmd *cobra.Command) ([]byte, error) {
	latest, err := cmd.Flags().GetBool("latest")
	if err != nil {
		return nil, err
	}

	url := baseURL
	if latest {
		url += "/latest"
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func filterByPeriod(cmd *cobra.Command, results []models.LotoResult) ([]byte, error) {
	slicedResults := []models.DailyResult{}

	initalDate, err := cmd.Flags().GetString("initial")
	if err != nil {
		log.Fatalln(err)
	}

	finalDate, err := cmd.Flags().GetString("final")
	if err != nil {
		log.Fatalln(err)
	}

	initalDateTime, err := time.Parse("02/01/2006", initalDate)
	if err != nil {
		log.Fatalln(err)
	}

	finalDateTime, err := time.Parse("02/01/2006", finalDate)
	if err != nil {
		log.Fatalln(err)
	}

	for _, res := range results {
		date, err := time.Parse("02/01/2006", res.Data)
		if err != nil {
			return nil, err
		}

		if !date.Before(initalDateTime) && !date.After(finalDateTime) {
			slicedResults = append(slicedResults, models.DailyResult{
				Concurso: res.Concurso,
				Data:     res.Data,
				Dezenas:  res.Dezenas,
			})
		}
	}

	data, err := json.Marshal(slicedResults)
	if err != nil {
		return nil, err
	}

	return data, nil
}

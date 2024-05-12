package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/LuizGuilherme13/resultados_lotofacil/models"
	"github.com/adhocore/chin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(GenerateBetCmd)

	GenerateBetCmd.Flags().Int64P("quantity", "q", 1, "amount of bets generated.")
}

var GenerateBetCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a random bet.",
	Long:  "Generates a random bet with the 15 most drawn numbers based on the search results.",
	Run:   GenerateBet,
}

type biggestDozens struct {
	ten   string
	count int
}

func GenerateBet(cmd *cobra.Command, args []string) {
	if _, err := os.Stat("./results.json"); errors.Is(err, os.ErrNotExist) {
		log.Println("The file with the results does not exist or was not found.")
		fmt.Println("Run 'loto get' to fetch and generate the file.")
		os.Exit(1)
	}

	s := chin.New()
	go func() {
		fmt.Print("Generating... ")
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

	ranking := map[string]int{}
	for _, res := range results {
		for _, dez := range res.Dezenas {
			ranking[dez]++
		}
	}

	dozens := []biggestDozens{}
	for ten, count := range ranking {
		dozens = append(dozens, biggestDozens{ten: ten, count: count})
	}

	sort.Slice(dozens, func(i, j int) bool {
		return dozens[i].count > dozens[j].count
	})

	minDrawns := float64(len(results) / 2)

	fmt.Println()

	tens := []string{}
	for _, v := range dozens {
		if float64(v.count) >= minDrawns {
			tens = append(tens, v.ten)
		}
	}

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	qtndBet, err := cmd.Flags().GetInt64("quantity")
	if err != nil {
		log.Fatalln(err)
	}

	newBet := []string{}
	for i := range qtndBet {
		for j := 0; j < 15; j++ {
			pos := r.Intn(len(tens[0:16]))
			newBet = append(newBet, tens[pos])
		}
		fmt.Printf("%d°: %v\n", i+1, newBet)
		newBet = []string{}
	}

	s.Stop()
	fmt.Printf("Successful! Passing score: %f\n", minDrawns)
}

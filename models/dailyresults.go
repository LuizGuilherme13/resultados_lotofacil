package models

type DailyResult struct {
	Concurso int      `json:"concurso"`
	Data     string   `json:"data"`
	Dezenas  []string `json:"dezenas"`
}

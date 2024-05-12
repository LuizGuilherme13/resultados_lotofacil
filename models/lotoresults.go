package models

// LotoResult ...
type LotoResult struct {
	Loteria             string        `json:"loteria"`
	Concurso            int           `json:"concurso"`
	Data                string        `json:"data"`
	Local               string        `json:"local"`
	DezenasOrdemSorteio []string      `json:"dezenasOrdemSorteio"`
	Dezenas             []string      `json:"dezenas"`
	Trevos              []interface{} `json:"trevos"`
	TimeCoracao         interface{}   `json:"timeCoracao"`
	MesSorte            interface{}   `json:"mesSorte"`
	Premiacoes          []struct {
		Descricao   string  `json:"descricao"`
		Faixa       int     `json:"faixa"`
		Ganhadores  int     `json:"ganhadores"`
		ValorPremio float64 `json:"valorPremio"`
	} `json:"premiacoes"`
	EstadosPremiados    []interface{} `json:"estadosPremiados"`
	Observacao          string        `json:"observacao"`
	Acumulou            bool          `json:"acumulou"`
	ProximoConcurso     int           `json:"proximoConcurso"`
	DataProximoConcurso string        `json:"dataProximoConcurso"`
	LocalGanhadores     []struct {
		Ganhadores     int    `json:"ganhadores"`
		Municipio      string `json:"municipio"`
		NomeFatansiaUL string `json:"nomeFatansiaUL"`
		Serie          string `json:"serie"`
		Posicao        int    `json:"posicao"`
		Uf             string `json:"uf"`
	} `json:"localGanhadores"`
	ValorArrecadado                float64 `json:"valorArrecadado"`
	ValorAcumuladoConcurso05       float64 `json:"valorAcumuladoConcurso_0_5"`
	ValorAcumuladoConcursoEspecial float64 `json:"valorAcumuladoConcursoEspecial"`
	ValorAcumuladoProximoConcurso  float64 `json:"valorAcumuladoProximoConcurso"`
	ValorEstimadoProximoConcurso   float64 `json:"valorEstimadoProximoConcurso"`
}

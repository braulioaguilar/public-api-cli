package presenter

// Data struct
type Data struct {
	API         string `json:"api"`
	Description string `json:"description"`
	Auth        string `json:"auth"`
	HTTPS       bool   `json:"https"`
	Cors        string `json:"cors"`
	Link        string `json:"link"`
	Category    string `json:"category"`
}

// Entry struct
type Entry struct {
	Count   int    `json:"count"`
	Entries []Data `json:"entries"`
}

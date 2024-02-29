package models

type Intervention struct {
	Site         string           `json:"site"`
	Descriptions []string         `json:"descriptions"`
	Materials    []string         `json:"materials"`
	Workers      []map[string]int `json:"workers"`
	Notes        string           `json:"notes"`
	Timestamp    string           `json:"timestamp"`
}

package models

type Intervention struct {
	Guid         string           `json:"guid"`
	Site         string           `json:"site"`
	Descriptions []string         `json:"descriptions"`
	Materials    []string         `json:"materials"`
	Workers      []map[string]int `json:"workers"`
	Notes        string           `json:"notes"`
	Timestamp    string           `json:"timestamp"`
}

// Flatten takes a slice of maps and returns a single map
// effectively flattening the input
func Flatten(mapToFlatten []map[string]int) map[string]int {
	flattened := make(map[string]int)
	for _, item := range mapToFlatten {
		for name, hours := range item {
			flattened[name] = hours
		}
	}
	return flattened
}

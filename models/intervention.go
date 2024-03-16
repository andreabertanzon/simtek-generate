package models

type Intervention struct {
	Guid         string               `json:"guid"`
	Site         string               `json:"site"`
	Descriptions []string             `json:"descriptions"`
	Materials    []string             `json:"materials"`
	Workers      []map[string]float32 `json:"workers"`
	Notes        string               `json:"notes"`
	Timestamp    string               `json:"timestamp"`
}

// Flatten takes a slice of maps and returns a single map
// effectively flattening the input
func Flatten[T comparable](mapToFlatten []map[string]T) map[string]T {
	flattened := make(map[string]T)
	for _, item := range mapToFlatten {
		for name, value := range item {
			flattened[name] = value
		}
	}
	return flattened
}

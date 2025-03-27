package graph

import "fmt"

func parseInt32(s string) (*int32, error) {
	var i int32
	_, err := fmt.Sscanf(s, "%d", &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
func parseFloat64(s string) (*float64, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

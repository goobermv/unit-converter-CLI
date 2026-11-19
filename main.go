package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lengthMap = map[string]float64{
	"m":  1.0,
	"km": 1000.0,
	"cm": 0.01,
	"mm": 0.001,
	"mi": 1609.344,
	"ft": 0.3048,
	"in": 0.0254,
}

func convertLength(value float64, from, to string) (float64, error) {
	fromFactor, okFrom := lengthMap[strings.ToLower(from)]
	toFactor, okTo := lengthMap[strings.ToLower(to)]

	if !okFrom || !okTo {
		return 0, fmt.Errorf("incorrect unit: %s or %s", from, to)
	}

	metre := value * fromFactor
	return metre / toFactor, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <value> <unit> <unit>")
		fmt.Println("Example: go run main.go 5 km m")
	}

	val, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid number: ", os.Args[1])
	}

	result, err := convertLength(val, os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("%.4f %s = %.4f %s\n", val, os.Args[2], result, os.Args[3])

}

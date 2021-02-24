package read

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFloats(fileName string) ([]float64, error) {

	var result []float64
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var aux float64
	for scanner.Scan() {
		aux, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		result = append(result, aux)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return result, nil
}

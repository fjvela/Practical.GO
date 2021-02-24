package read

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFloats(fileName string) ([3]float64, error) {

	var result [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
			return result, err
		}
		i++
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	if scanner.Err() != nil {
		return result, scanner.Err()
	}

	return result, nil
}

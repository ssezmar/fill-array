package main

import "log"

const (
	rowsCount = 5
	colsCount = 5
	randomLimit = 100
)


func main() {
	randomizer := NewRandomizer()

	randomIntegers, err := randomizer.getUniquerandomIntegers(GetUniqueIntegresInput{
		SizeX: rowsCount, 
		SizeY: colsCount,
		RandomLimit: randomLimit,
	})
	if err != nil {
		log.Fatal(err)
	}

	array, err := NewIntArray(rowsCount, colsCount)
	if err != nil {
		log.Fatal(err)
	}

	if err := array.Fill(randomIntegers); err != nil {
		log.Fatal(err)
	}

	if err:= Print(array); err != nil {
		log.Fatal(err)
	}

}

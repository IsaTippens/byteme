package main

import (
	"fmt"
	"os"
	"time"

	"github.com/segmentio/encoding/json"
)

/*
	34: "
	44: ,
	58: :
	123: {
	125: }
*/

type Response struct {
	U         float64 `json:"u"`
	Code      string  `json:"s"`
	Bid       string  `json:"b"`
	BidVolume float64 `json:"B"`
	//BidVolume float64 `json:"B"`
	Ask       string  `json:"a"`
	AskVolume float64 `json:"A"`
	//AskVolume float64 `json:"A"`
}

var TRIALS = 1000000

func main() {
	data, err := os.ReadFile("ticker.txt")
	if err != nil {
		panic(err)
	}

	jt := JsonTest(&data)
	bt, r := ByteSliceTest(&data)

	fmt.Printf("Number of trials: %+v\n", TRIALS)
	fmt.Printf("Json time: %+v\n", jt)
	fmt.Printf("Byte slice time: %+v\n", bt)
	for i := range r {
		fmt.Printf("%+v\n", string(r[i]))
	}

}

func JsonTest(data *[]byte) time.Duration {
	start := time.Now()
	for i := 0; i < TRIALS; i++ {
		result := JsonUnmarshal(data)
		var _ string
		_ = result.Bid
		_ = result.BidVolume
		_ = result.Ask
		_ = result.AskVolume
	}

	end := time.Since(start)
	return end
}

func JsonUnmarshal(data *[]byte) Response {
	var result Response
	json.Unmarshal(*data, &result)
	return result
}

func ByteSliceTest(data *[]byte) (time.Duration, [][]byte) {
	var result [][]byte
	start := time.Now()
	for i := 0; i < TRIALS; i++ {
		result = SliceByteMap(data)
	}
	return time.Since(start), result
}

func SliceByteMap(data *[]byte) [][]byte {
	d := *data
	ticker := make([][]byte, 5)
	start := 0
	end := 0
	extracting := false
	count := 0
	//RANGE 48-57 "0-9" or 46 "."
	for i := range d {
		if (d[i] >= 48 && d[i] <= 57) || (d[i] == 46) {
			if !extracting {
				extracting = true
				start = i
			}
		} else {
			if extracting {
				extracting = false
				end = i
				ticker[count] = d[start:end]
				count++
			}
		}

	}
	return ticker
}

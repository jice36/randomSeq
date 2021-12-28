package main

import (
	"fmt"
	"math/rand"
	"os"
	"randomNumber"
	"time"
)



func main() {
	str, test := timeTest("test gen key 100mb " )
	var i int
	var t randomNumber.Test
	var rng randomNumber.ISAAC
	file, err := os.Create("1mb.txt")
	if err != nil{
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	for i < 6250000{
		seq := rng.GenSeq()
		if flag, err := t.CheckSeq(seq); flag == true && err == nil{
			file.Write(seq)
		}

		i++
	}

	defer func() {
		testing(str, test)
	}()
}

func timeTest(test string) (string, time.Time) {
	return test, time.Now()
}

func testing(test string, start time.Time) {
	fmt.Println( test, time.Since(start).Seconds())
}
func randSlice(s []byte) {
	i := 0
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 255
	for i < len(s) {
		s[i] = byte(rand.Intn(max-min+1) + min)
		i++
	}
}



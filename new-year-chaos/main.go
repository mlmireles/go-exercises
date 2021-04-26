package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	bribes := int32(0)
	chaotic := false

	length := int32(len(q))

	// Start from front. Array positions match people position
	for i := int32(0); i < length; i++ {
		// if people and array positions match, there is no bribe
		if q[i]-(i+1) > 2 {
			chaotic = true
		}

		j := int32(q[i] - 2)
		if j < 0 {
			j = 0
		}
		for ; j < i; j++ {
			if q[j] > q[i] {
				bribes++
			}
		}

		/*
		   // get the number of bribes for the current person
		   b := q[i] - (i + 1)

		   // If is bigger than zero the person bribed
		   if b > 0 {
		       bribes += b
		   }

		   // If more than two bribes return
		   if b > 2 {
		       fmt.Println("Too chaotic")
		       return
		   }

		   // If just move one position continue
		   if b >= -1 {
		       continue
		   }

		   // The person move more than one position on back
		   // Loop for check if the persons in front move further
		   var adjust int32

		   // Init loop position
		   n := i + b + 1
		   // Continue if position is close to end
		   if n >= length - 2 {
		       continue
		   }

		   //fmt.Printf("i: %d, b: %d \n", i, b)
		   for ; n < length; n++ {
		       adjust += q[n] - (n + 1)
		       //fmt.Printf("n: %d, q[n]: %d, ad: %d \n", n, q[n], adjust)
		   }
		   if adjust < 0 {
		       bribes -= adjust
		   }
		*/

	}

	if chaotic {
		fmt.Println("Too chaotic")
		return
	}

	fmt.Println(bribes)
	//fmt.Println("---")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "errors"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
    fmt.Println(arr)
    // n - (i + 1)
    swaps := int32(0)
    weights, err := getArrayWeights(arr)
    if  err != nil {
        return swaps
    }
    
    //fmt.Printf("weights: %d \n", weights)
    
    min, max := getMinAndMaxPositions(weights)
    //fmt.Printf("min: %d, max: %d\n", min, max)
    
    swap(arr, min, max)
    swaps++
    
    //fmt.Printf("%d: %d \n", swaps, arr)
    //fmt.Println("-----")
    return swaps + minimumSwaps(arr)
}

func getArrayWeights(arr []int32) ([]int32, error) {
    var weights = make([]int32, len(arr), len(arr))
    copy(weights, arr)
    
    w := false
    for i := 0; i < len(arr); i++ {
        n := arr[i]
        s := n - (int32(i) + 1)
        weights[i] = s
        if  s != 0 {
            w = true
        }
    }
    
    if  !w {
        return nil, errors.New("All in position")
    }
    
    return weights, nil
}

func getMinAndMaxPositions(arr []int32) (min int, max int) {
    mi := arr[0]
    ma := arr[0]
    for i, value := range arr {
        if value < mi {
            min = i
            mi = value
        }
        if value > ma {
            max = i
            ma = value
        }
    }
    
    return min, max
}

func swap(arr []int32, a, b int) {
    x := arr[a]
    z := arr[b]
    
    arr[a] = z
    arr[b] = x
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    res := minimumSwaps(arr)

    fmt.Fprintf(writer, "%d\n", res)

    writer.Flush()
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

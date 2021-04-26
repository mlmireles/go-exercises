package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
    fmt.Println("----")
    
    as1 := strings.Split(s1, "")
    as2 := strings.Split(s2, "")
    
    sort.Strings(as1)
    fmt.Printf("s1: %s\n", as1)
    //sort.Strings(as2)
    //fmt.Printf("s2: %s\n", as2)
    
    //l := len(as2)
    for i, v := range as1 {
        if i != 0 && v == as1[i - 1] {
            continue
        }
        
        for _, b := range as2 {
            if v == b {
                return "YES"
            }
        }
        
        /*n := sort.SearchStrings(as2, v)
        fmt.Printf("v: %s; n: %d\n", v, n)
        if n != l {
            return "YES"
        }*/
    }
    
    return "NO"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        s1 := readLine(reader)

        s2 := readLine(reader)

        result := twoStrings(s1, s2)

        fmt.Fprintf(writer, "%s\n", result)
    }

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

package main

import (
    "fmt"
)

func main() {
    m := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
    fmt.Println(m)

    fmt.Println(m["name"])

    v := m["names"]
    fmt.Println(v)

    ma := map[string]int{"name":1, "animal":1, "color":1, "location":1}
    va := ma["names"]
    fmt.Println(va)
    if va == 0 {
        fmt.Println("Nil value")
    }
}

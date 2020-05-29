package main 

import (
    "fmt"
    "io/ioutil"
)

func main() {
    fmt.Print("Please enter Unix filename:-")
    var filename string
    fmt.Scanln(&filename)

    data, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

}

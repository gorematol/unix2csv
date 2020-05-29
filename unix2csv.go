package main 

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
)

func main() {

    // Specify Unix filename	
    fmt.Print("Please enter Unix filename:-")
    var filename string
    fmt.Scanln(&filename)

    // specify unix delimeter
    fmt.Print("Please enter file delimeter:-")
    var d string
    fmt.Scanln(&d)
    var csvd string = ","

    // Open unix file
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }

    // Read file line by then
    lineoutput := bufio.NewScanner(file)
    for lineoutput.Scan() {
	if strings.Contains(lineoutput.Text(), d) {
            newstr := strings.Replace(lineoutput.Text(), d,  csvd, -1)
            fmt.Println(newstr)
        }
    }

    file.Close()

}

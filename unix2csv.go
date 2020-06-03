package main 

import (
    "fmt"
    "log"
    "os"
//    "bufio"
    "io/ioutil"
    "encoding/json"
//    "reflect"
)


func main() {

    // var csvd string = ","
    type File struct {
        Input string 
	Delimeter string
	Output string  
    }

    type Unixfiles struct  {
	Files []File 
    }

    // Open json file
    jsonfile, err := openFile("unixmeta.json")
    if err != nil {
	log.Fatal(err)
        os.Exit(1) 
    }

    // Read json file into a byte array
    b, err := ioutil.ReadAll(jsonfile)
    if err != nil {
	log.Fatal(err)
        os.Exit(1) 
    }

 //   metadata := Unixfiles{}
 //   json.Unmarshal(b, &metadata)
    processData(Unixfiles{}, b)
    jsonfile.Close()
}

func openFile(filename string) (*os.File, error) {
        f, err := os.Open(filename)
        if err != nil {
		return nil, err 
        }
        return f, nil
}

func processData(metadata interface{}, b []byte) {

      json.Unmarshal(b, &metadata)
      m := metadata.(map[string]interface{}) 
      fmt.Println(m) 

      for _,value := range m {
	      for _,file := range value.([]interface{}) {
		      input, err := file["Input"]
		      fmt.Println(input)
	      }
      }
      /* 
      file, err := openFile(value.Input)
      if err != nil {
	log.Fatal(err)
        os.Exit(1) 
      }

      lineoutput := bufio.NewScanner(file)
      for lineoutput.Scan() {
//	if strings.Contains(lineoutput.Text(), d) {
//            newstr := strings.Replace(lineoutput.Text(), d,  csvd, -1)
//            fmt.Println(newstr)
        fmt.Println(lineoutput.Text())
//        }
      }
      file.Close()
    }
      */

}    


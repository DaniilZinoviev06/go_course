package main

import (
  "fmt"
  "flag" // для флага --flag
  "io/ioutil" // для четния из файла
)

func main() {
  file_flag := flag.String("flag", "", "Flag for JSON")
  flag.Parse()

  if *file_flag == "" {
    fmt.Println("--flag not defined")
  } else {
    fmt.Println("--flag defined")
  }

  file_data, err := ioutil.ReadFile(*file_flag)
  if err != nil {
    fmt.Println("Error when trying to read data from file")
  } else {
    fmt.Println("File name: ", *file_flag)
    fmt.Println(string(file_data))
  }
}

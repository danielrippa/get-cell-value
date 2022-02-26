package main

import (
  "os"
  "fmt"
  "github.com/xuri/excelize/v2"
)

func main() {

  args := os.Args[1:]
  if len(args) < 1 { return }

  f, err := excelize.OpenFile(args[0])
  if err != nil {
    fmt.Println(err)
    return
  }

  value, err := f.GetCellValue(args[1], args[2])
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(value)

}
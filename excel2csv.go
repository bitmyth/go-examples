package main

import (
    "flag"
    "fmt"
    "github.com/tealeg/xlsx"
    "os"
    "strings"
)

func main() {
    var excelFileName string
    var glue string
    var outputFile string

    flag.StringVar(&excelFileName, "f", "", "excel file path")
    flag.StringVar(&glue, "g", ",", "glue")
    flag.StringVar(&outputFile, "o", "tmp.csv", "output file path")

    flag.Parse()

    f := openFileForAppend(outputFile)
    defer f.Close()

    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        fmt.Printf("open failed: %s\n", err)
    }
    for _, sheet := range xlFile.Sheets {
        fmt.Printf("Sheet Name: %s\n", sheet.Name)
        i := 0
        for _, row := range sheet.Rows {
            //fmt.Printf("%d\n", len(row.Cells))
            i++

            cells := make([]string, 0, len(row.Cells))

            for _, cell := range row.Cells {
                cells = append(cells, cell.String())
            }
            csv := strings.Join(cells, glue)
            csv = strings.ReplaceAll(csv, "\n", "")

            if _, err := f.WriteString(csv + "\n"); err != nil {
                panic(err)
            }
        }

        fmt.Printf("Sheet name: %s rows %d:", sheet.Name, i)

    }
}

func appendFile(outputFile string, data string) error {

    f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    if _, err := f.WriteString(data); err != nil {
        return err
    }

    return nil
}
func openFileForAppend(outputFile string) *os.File {

    f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        return nil
    }

    return f
}

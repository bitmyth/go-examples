package main

import (
    "fmt"
    "flag"
    "github.com/tealeg/xlsx"
)

func main() {
    var excelFileName string

    flag.StringVar(&excelFileName, "file", "", "excel file path")

    flag.Parse()

    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        fmt.Printf("open failed: %s\n", err)
    }
    for _, sheet := range xlFile.Sheets {
        fmt.Printf("Sheet Name: %s\n", sheet.Name)
		i:=0
        for _, _ = range sheet.Rows {
            //fmt.Printf("%d\n", len(row.Cells))
			i++
		}

		fmt.Printf("Sheet name: %s rows %d:", sheet.Name, i)

        //for _, row := range sheet.Rows {
		//	i++
        //    for _, cell := range row.Cells {
        //        text := cell.String()
        //        fmt.Printf("%s\n", text)
        //    }
        //}
    }
}

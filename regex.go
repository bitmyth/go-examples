package main

import "fmt"
import "regexp"
import "io/ioutil"
import "math"
import "strings"

// match url by regex

func main(){
	var r, _ = regexp.Compile(`((http://|https://)?(www\.)?[^ "';<>\n>]+?\.(com|cn|net|org)[^ "';<>\n>]+?)[ "'>]`)
	data, _ := ioutil.ReadFile("test.html")
	matches := r.FindAllString(string(data), math.MaxInt64)
	count :=0

	if matches != nil {
		for _, value := range matches {
			url := strings.Trim(value, "//\"'")
			fmt.Println(url)
			count++
		}
	}

	println(count)
}

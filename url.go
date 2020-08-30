package main

import (
	"flag"
	"log"
	"net/url"
)

// Usage
// $ go run url.go -url /www.baidu.com?2=b

func main() {
	var u string
	flag.StringVar(&u, "url", "www.baidu.com/a/b/c?2=b", "input url to parse")
	flag.Parse()
	println(u)
	parsed, e := url.Parse(u)
	if e != nil {
		log.Fatal(e)
		return
	}
	log.Println(parsed.Scheme)
	log.Println(parsed.Host)
	log.Println(parsed.Path)
	log.Println(parsed.RawPath)
	log.Println(parsed.RawQuery)
}

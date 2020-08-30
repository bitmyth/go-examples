package hello

import (
	"rsc.io/quote/v3"
//	quoteV3 "rsc.io/quote/v3"
)

func Hello() string{
	//return "hello,world."
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}

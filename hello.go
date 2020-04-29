package hello

import (
	"rsc.io/quote/v3"
)

func Person() string {
	return "Bonjour SBE !?"
}

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}

// Learn Go module with rsc.io quote package.
package hello

import (
	quoteV3 "rsc.io/quote/v3"
)

// Hello returns a greeting.
func Hello() string {
	return quoteV3.HelloV3()
}

// Concurrency returns a Go proverb about concurrency.
func Proverb() string {
	return quoteV3.Concurrency()
}

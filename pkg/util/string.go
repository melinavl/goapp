package util

import (
	"fmt"
	"math/rand"
)

var randx = rand.NewSource(42)

// RandString returns a random hex of length n.
func RandString(n int) string {
	const letterBytes = "ABCDEF0123456789"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, randx.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randx.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// RandHex returns a random hex of length n.
func RandHex(n int) string {

	rng := rand.New(randx)

	// Define the minimum and maximum values for the random integer
	minValue := 0x1000000000
	maxValue := 0xFFFFFFFFFF

	// Generate a random integer between minValue and maxValue (inclusive)
	randomInt := rng.Intn(maxValue-minValue+1) + minValue

	hexString := fmt.Sprintf("%X", randomInt)

	return hexString
}

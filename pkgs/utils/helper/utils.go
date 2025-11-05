package helper

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateAccountNumber() string {
	prefix := "180"

	max := big.NewInt(100000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		// fallback if random fails (rare)
		return prefix + "00000"
	}

	accountNumber := fmt.Sprintf("%s%05d", prefix, n.Int64())
	return accountNumber
}

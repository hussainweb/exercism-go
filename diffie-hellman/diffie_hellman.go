package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// Not secure. But satisfies the conditions of the test.
// Use crypto.rand instead.
var randomGenerator = rand.New(rand.NewSource(time.Now().Unix()))

// PrivateKey generates a private key
func PrivateKey(p *big.Int) *big.Int {
	var key big.Int
	one := big.NewInt(1)
	// We can't return 1 or less.
	for key.Cmp(one) <= 0 {
		key.Rand(randomGenerator, p)
	}
	return &key
}

// PublicKey generates a public key for a private key
func PublicKey(private, p *big.Int, g int64) *big.Int {
	pubkey := big.NewInt(g)
	pubkey.Exp(pubkey, private, p)
	return pubkey
}

// NewPair generates a new pair
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	key := PrivateKey(p)
	return key, PublicKey(key, p, g)
}

// SecretKey calculates a secret key
func SecretKey(private1, public2, p *big.Int) *big.Int {
	var key big.Int
	key.Exp(public2, private1, p)
	return &key
}

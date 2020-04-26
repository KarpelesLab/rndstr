package rndstr

import (
	"io"
	"math/rand"
)

// Simple is a simple generator using a given len and range.
// Example use: rndstr.Simple(10, rndstr.Alnum)
func Simple(l int, rng string) string {
	buf := make([]byte, l)
	rlen := len(rng)

	for i := range buf {
		buf[i] = rng[rand.Intn(rlen)]
	}
	return string(buf)
}

// SimpleReader can be used with a random source, such as crypto/rand
//
// This is fairly slow and will consume at least 4 bytes of random input per
// generated character. This is however secure and follows best practices in
// random number generation.
//
// Usage: rndstr.SimpleReader(10, rndstr.Alnum, rand.Reader)
func SimpleReader(l int, rng string, src io.Reader) (string, error) {
	buf := make([]byte, l)
	rlen := int32(len(rng))

	for i := range buf {
		v, err := readint31n(src, rlen)
		if err != nil {
			return "", err
		}
		buf[i] = rng[v]
	}
	return string(buf), nil
}

// SimpleFastReader can be used with a random source, such as crypto/rand
//
// Unsafe for secure operations but fast.
//
// Usage: rndstr.SimpleFastReader(10, rndstr.Alnum, rand.Reader)
func SimpleFastReader(l int, rng string, src io.Reader) (string, error) {
	buf := make([]byte, l)
	_, err := io.ReadFull(src, buf)
	if err != nil {
		return "", err
	}
	rlen := len(rng)

	for i := range buf {
		buf[i] = rng[int(buf[i])%rlen]
	}
	return string(buf), nil
}

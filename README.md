[![GoDoc](https://godoc.org/github.com/KarpelesLab/rndstr?status.svg)](https://godoc.org/github.com/KarpelesLab/rndstr)

# Simple random string generator

This allows generating random strings very easily.

```go
	fmt.Printf("Random string: %s\n", rndstr.Simple(16, rndstr.Alnum))

	// or to generate a secure string, using crypto/rand
	secureStr := rndstr.SimpleReader(16, rndstr.Alnum, rand.Reader)
```

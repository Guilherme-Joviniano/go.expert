package main

// To use this package you need to authenticate in github account
// update the .netrc (token or SSH) file and the GOPRIVATE ENV Variable 

import (
	"github.com/gjovs/secret"
)

func main() {
	secret.Secret()
}

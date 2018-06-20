package main

import (
	"fmt"
	"github.com/r2ishiguro/vrf/go/vrf_ed25519"
	"golang.org/x/crypto/ed25519"
)

func main() {
	const message = "message"
	pk, sk, err := ed25519.GenerateKey(nil)
	if err != nil {
		fmt.Println("errors")
	}
	c := vrf_ed25519.ECVRF_hash_to_curve([]byte(message), pk)
	fmt.Println("Hash:")
	fmt.Println(c)
	pi, err := vrf_ed25519.ECVRF_prove(pk, sk, []byte(message))
	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println()
		fmt.Println("Proof:")
		fmt.Println(pi)
	}
	res, err := vrf_ed25519.ECVRF_verify(pk, pi, []byte(message))
	if err != nil {
		fmt.Println("errors")
	} else {
		fmt.Println()
		fmt.Println("Result:")
		fmt.Println(res)
	}
}

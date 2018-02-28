package main

import (
	"fmt"
	"os"

	"github.com/blaskovicz/go-nest"
)

// get an access_token given client id, secret, and a code
// Set NEST_CODE, NEST_CLIENT_ID and NEST_CLIENT_SECRET env vars before running this
// See https://developers.nest.com/documentation/cloud/how-to-auth
func main() {
	c := nest.New()
	at, err := c.CreateAccessToken(os.Getenv("NEST_CODE"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created token: %#v\n", at)
}

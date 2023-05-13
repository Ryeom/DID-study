package main

import (
	"fmt"
	"ssikr/core"
)

func main() {
	did, _ := core.NewDID("test", "qwerty")
	fmt.Printf("[did]%s", did.String())

}

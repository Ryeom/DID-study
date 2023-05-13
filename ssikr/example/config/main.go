package main

import (
	"fmt"
	"ssikr/config"
)

func main()  {

	fmt.Println(config.SystemConfig.RegistrarAddr)
	fmt.Println(config.SystemConfig.ResolverAddr)
}
// 금융권은 512써랑


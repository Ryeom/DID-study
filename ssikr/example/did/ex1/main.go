package main

import "fmt"

//did는 단순 식별자로 W3C에서 제시하는 표준형식을 따른다 (표준화 되어있지는 않고 표준화 과정중에 있다)
// https://w3.org/TR/did-core/
func main()  {
	method := "ssikr"
	specificIdentitier := "abcdeeeeeeeeeeeeeeeeeeeee"
	did := fmt.Sprintf("did:%s:%s",method,specificIdentitier)
	fmt.Printf("[DID]%s\n",did)
}

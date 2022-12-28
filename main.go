package main

import (
	"fmt"
	"go-study/something"
	"rsc.io/quote"
)

/*
 * 프로젝트의 Compile이 필요한 경우 main.go 만들기.
 *  - Java의 main(String[] args) 함수와 동일한 기능을 수행한다.
 */
func main() {
	fmt.Println(quote.Hello())

	// - Function 을 대문자로 작성하는 경우 export 할 수 있는 함수가 된다.
	something.SayHello()
	// something.sayBye()
}

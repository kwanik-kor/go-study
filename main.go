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

	// 1. Function 을 대문자로 작성하는 경우 export 할 수 있는 함수가 된다.
	something.SayHello()
	// something.sayBye()

	// 2. 변수와 상수
	// 2.1 상수 - 변수명 옆에 타입을 지정해줄 수 있음
	const name string = "kwan-ik"
	fmt.Println(name)

	// 2.2 변수 -
	nameVar := "kwan-ik"
	//  == var nameVar string = "kwan-ik"
	nameVar = "gani"
	fmt.Println(nameVar)
}

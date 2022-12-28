package main

import (
	"fmt"
	"rsc.io/quote"
)

/*
 * 프로젝트의 Compile이 필요한 경우 main.go 만들기.
 *  - Java의 main(String[] args) 함수와 동일한 기능을 수행한다.
 */
func main() {
	fmt.Println(quote.Go())
}

package main

import "log"

func div(firstNum, secondNum int) {
	if secondNum == 0 {
		log.Fatal("can not devided by 0!")
	}
	result := firstNum / secondNum
	log.Println("Result is", result)
}
func main() {
	firstNum := 10
	secondNum := 0
	div(firstNum, secondNum)
}

package main

import "fmt"

func sum(a int, b int) int {
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("_______________________")
	s := a + b
	return s
}

func sumString(s1 string, s2 string) string {
	s := s1 + s2

	return s
}
func main() {
	number := sum(11, 2)
	fmt.Printf("Вы получили число: %d\n", number)
	fmt.Println("_______________________")

	text1 := "Привет "
	text2 := "Мир!"

	if number < 10 {
		fmt.Printf("Вы ввели корректный диапозон чисел < 10, число: %d\n", number)
		text3 := sumString(text1, text2)
		fmt.Println(text3)
	} else {
		fmt.Println("Вы не попали в дапозон числе < 10")
	}
}

package stringsutil

import "fmt"

func ReadStrings() (str1, str2 string) {
	fmt.Print("Enter the first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&str2)
	return str1, str2
}

func AppendString(str1, str2 string) string {
	appendedString := str1 + " " + str2
	fmt.Println("Appended String:", appendedString)
	return appendedString
}

func PrintStringNTimes(appendedString string) {
	var x int
	fmt.Println("Enter an integer:")
	fmt.Scanln(&x)
	for i := 0; i < x; i++ {
		fmt.Println(appendedString)
	}
}
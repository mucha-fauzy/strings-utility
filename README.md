# Strings Utility

## Functions
This program have 3 function to:
  * Read 2 input string and print the value
```
func ReadStrings() (str1, str2 string) {
	fmt.Print("Enter the first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&str2)
	return str1, str2
}
```
  * Append the string and print the value
```
func AppendString(str1, str2 string) string {
	appendedString := str1 + " " + str2
	fmt.Println("Appended String:", appendedString)
	return appendedString
}
```
  * Read 1 input integer x and print the appended string x times
```
func PrintStringNTimes(appendedString string) {
	var x int
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&x)
	for i := 0; i < x; i++ {
		fmt.Println(appendedString)
	}
}
```
## How to Use
```
package main

import (
	"git/pkg/stringsutil"
)

func main() {
	str1, str2 := stringsutil.ReadStrings()

	appendedString := stringsutil.AppendString(str1, str2)

	stringsutil.PrintStringNTimes(appendedString)
}

```

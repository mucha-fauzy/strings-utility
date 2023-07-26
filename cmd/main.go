package main

import (
	"git/pkg/stringsutil"
)

func main() {
	str1, str2 := stringsutil.ReadStrings()
	appendedString := stringsutil.AppendString(str1, str2)
	stringsutil.PrintStringNTimes(appendedString)
}
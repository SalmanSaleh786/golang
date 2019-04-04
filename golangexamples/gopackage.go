package golangexamples

import (
	"github.com/ehteshamz/greetings"
)

//ConcatSlice function which concat slices
func ConcatSlice(sliceToConcate []byte) string {
	//concatanatedSlice := make([]byte, 0, 30)
	newString := ""
	for _, v := range sliceToConcate {
		newString += string(v)
		//concatanatedSlice = append(concatanatedSlice, v)
	}
	return newString
}

//Encrypt function which encrypts slice
func Encrypt(sliceToEncrypt []byte, ceaserCount int) string {
	for i := range sliceToEncrypt {
		sliceToEncrypt[i] += byte(ceaserCount)
	}
	return string(sliceToEncrypt)
}

//EZGreetings func
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}

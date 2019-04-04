package greetings

import (
	"github.com/ehteshamz/greetings"
)

func ConcatSlice(sliceToConcat []byte) string {

	str := ""
	for _, i := range sliceToConcat {
		//fmt.Printf(string(i))
		str += string(i)
		str += "-"

	}
	return str
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i := 0; i < len(sliceToEncrypt); i++ {
		sliceToEncrypt[i] += byte(ceaserCount)
		//println(string(sliceToEncrypt[i]))
	}

}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}

// func main() {
// 	//var Slice []byte
// 	Slice := make([]byte, 9)
// 	Slice[0] = 'J'
// 	Slice[1] = 'a'
// 	Slice[2] = 'i'
// 	Slice[3] = 'K'
// 	Slice[4] = 'i'
// 	Slice[5] = 's'
// 	Slice[6] = 'h'
// 	Slice[7] = 'a'
// 	Slice[8] = 'n'
//
// 	//Q1
// 	fmt.Println(ConcatSlice(Slice))
//
// 	//Q2
// 	Encrypt(Slice, 1)
// 	fmt.Println(string(Slice))
//
// 	//Q3
// 	fmt.Printf(EZGreetings("Jai Kishan"))
// }

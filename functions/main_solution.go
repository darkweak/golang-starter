package functions

import "fmt"

func returnTwoStrings() (string, string) {
	return "My first string", "My second string"
}

func returnThreeMixed() (string, int, bool) {
	return "Hello", 43, false
}

type AbstractTypeInterface interface {
	stringToByte(s string) []byte
	byteToString(b []byte) string
}

type FirstType struct{}
type SecondType struct{}

func (f *FirstType) stringToByte(s string) []byte {
	return []byte(fmt.Sprintf("%s %s", "My first : ", s))
}

func (f *FirstType) byteToString(b []byte) string {
	return fmt.Sprintf("%s %s", "First : ", string(b))
}

func (f *SecondType) stringToByte(s string) []byte {
	return []byte(fmt.Sprintf("%s %s", "My second : ", s))
}

func (f *SecondType) byteToString(b []byte) string {
	return fmt.Sprintf("%s %s", "Second : ", string(b))
}

func Functions_Solution() {
	abstract := make([]AbstractTypeInterface, 2)
	abstract[0] = &FirstType{}
	abstract[1] = &SecondType{}

	for _, i := range abstract {
		fmt.Println(i.stringToByte("Hello world"))
	}
}

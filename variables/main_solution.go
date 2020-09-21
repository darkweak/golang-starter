package variables

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const MY_CONSTANT = "This is my constant"

func GetConstantFromOutside() string {
	return MY_CONSTANT
}

func myFunctionReturnAVariable() string {
	myDynamicVarA := "something"
	myDynamicVarA = "something else"
	myDynamicVarA = "another thing"
	return myDynamicVarA
}

func myFunctionReturnAStringArray() []string {
	const constName = "John"
	var varName2 = "Patrick"
	// Cannot do another assignation for constName
	// constName = ""
	varName2 = "We can reassign a var"
	return []string{constName, varName2}
}

type talk struct {
	length int64
	title  string
}

func myFunctionReturnTheNonTypedArgument(nonTypedItem interface{}) interface{} {
	return nonTypedItem
}

type CustomField struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CustomStructureJsonable struct {
	FirstField  string      `json:"first_field"`
	SecondField string      `json:"second_field"`
	ThirdField  int         `json:"thirdField"`
	FourthField CustomField `json:"fourthField"`
}

func parseJsonFile() CustomStructureJsonable {
	jsonFile, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	var custom CustomStructureJsonable
	if err := json.Unmarshal(jsonFile, &custom); err != nil {
		log.Fatal(err)
	}
	return custom
}

func Variables_Solution() {
	fmt.Println("GetConstantFromOutside : ", GetConstantFromOutside())
	fmt.Println("myFunctionReturnAVariable : ", myFunctionReturnAVariable())
	fmt.Println("myFunctionReturnAStringArray : ", myFunctionReturnAStringArray())
	fmt.Println("myFunctionReturnTheNonTypedArgument as int : ", myFunctionReturnTheNonTypedArgument(1))
	fmt.Println("myFunctionReturnTheNonTypedArgument as string : ", myFunctionReturnTheNonTypedArgument("string is a string"))
	myTalk := talk{60, "Discover golang"}
	fmt.Println("myFunctionReturnTheNonTypedArgument as object : ", myFunctionReturnTheNonTypedArgument(myTalk))
	fmt.Println("Get talk length", myTalk.length)
	fmt.Println("Get talk title", myTalk.title)
	fmt.Println("Get custom item from file", parseJsonFile())
}

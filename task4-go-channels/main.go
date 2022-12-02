package main
import (
	"fmt"
	"regexp"
)

func removeNumerals(s []string, data chan string) {
	for _,str:= range s {
		re := regexp.MustCompile("[0-9]")
		data <- re.ReplaceAllString(str,"")
	}
	close(data)
}

func main() {
	numeralString:=[]string{"gopher123", "alpha99beta", "1cita2del3"}
	numeralFreeString:=[]string{}
	data:=make(chan string)
	go removeNumerals(numeralString,data)
	for s:= range data {
		numeralFreeString=append(numeralFreeString,s)
	}
	fmt.Println("Before Processing",numeralString)
	fmt.Println("After Processing",numeralFreeString)
}

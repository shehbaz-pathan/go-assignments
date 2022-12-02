package main
import (
	"fmt"
	"regexp"
	"sync"
)

func removeNumerals(s []string,wg *sync.WaitGroup) {
	for _,str:= range s {
		re := regexp.MustCompile("[0-9]")
		fmt.Printf("%v ",re.ReplaceAllString(str,""))
	}
	wg.Done()
}

func main() {
	numeralString:=[]string{"gopher123", "alpha99beta", "1cita2del3"}
	var wg sync.WaitGroup
	wg.Add(1)
	go removeNumerals(numeralString,&wg)
	wg.Wait()
}

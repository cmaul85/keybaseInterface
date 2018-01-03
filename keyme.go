package main

import (
    "fmt"
	"external"
)

func ApplePrint(s string){
    fmt.Println(s + " Apple")
}

func main(){
	ApplePrint("Green")
	external.Banana("Yellow")
}

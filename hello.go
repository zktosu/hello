package main

import(
	"fmt"
)
const englishPrefix = "Hello "
const trPrefix = "Merhaba "
const frPrefix = "Bonjour "

func getPrefix(lang string)(prefix string){
	switch lang{
	case "turkish":
		prefix = trPrefix
	case "french":
		prefix = frPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func Hello(name,lang string) string{
	if name == ""{
		name = "World"
	}
	return getPrefix(lang) + name + "!"
}


func main(){
	fmt.Println(Hello("Sebastian","french"))
}

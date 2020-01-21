package main

import(
	"fmt"
)
const englishPrefix = "Hello "
const trPrefix = "Merhaba "

func Hello(name,lang string) string{
	if name == ""{
		name = "World"
	}
	if lang == "turkish"{
		return trPrefix + name +"!"
	}
	if lang == "english"{
		return englishPrefix + name + "!"
	}
	return ""
}


func main(){
	fmt.Println(Hello("world","turkish"))
}

package main

import(
	"fmt"
)
const englishPrefix = "Hello "
const trPrefix = "Merhaba "
const frPrefix = "Bonjour "
func Hello(name,lang string) string{
	prefix := englishPrefix
	if name == ""{
		name = "World"
	}
	if lang == "turkish"{
		prefix = trPrefix
	}
	if lang == "french" {
		prefix = frPrefix
	}
	return prefix + name + "!"
}


func main(){
	fmt.Println(Hello("Sebastian","french"))
}

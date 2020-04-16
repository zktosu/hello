package greeter


func greet(name string) string {
	if name == "" {
		return "Hello"
	}else{
		return "Hello, "+name+"!"
	}
}

func Greet() string {
	return greet("")
}
func GreetName(name string) string{
	return greet(name)
}


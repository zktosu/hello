package iteration

const repeatCounter = 5

func Repeat(text string)(result string){
	for i:=0;i<repeatCounter;i++ {
		result +=text
	}
	return 
}

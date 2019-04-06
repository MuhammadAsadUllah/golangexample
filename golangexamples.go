package main

import "fmt"
func ConcatSLice(sliceToConcat []byte) string {
	var myStr string
	for i := 0; i < len(sliceToConcat); i++ {
		var ndt string
		ndt = string(sliceToConcat[i])+"-"
		myStr += ndt
		var length int
		length = len(sliceToConcat)
		if (length-1 != i){
			}else{
				myStr += string(sliceToConcat[i])
			}
		}
		return myStr
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int){
	funslice := sliceToEncrypt
	ceaserCount = (ceaserCount%26 + 26) % 26
	for i := 0; i < len(sliceToEncrypt); i++ {
			myVar := sliceToEncrypt[i]
			var count int
			switch {
			case 'A' <= myVar && myVar <= 'Z':
					count = 'A'
			case 'a' <= myVar && myVar <= 'z':
					count = 'a'
			default:
					funslice[i] = myVar
					continue
			}
			var tempv int
			tempv = (ceaserCount+(int(myVar)-count))%26
			sliceToEncrypt[i] = byte(count + tempv)
	}
	return
}
func EZGreetings(name string){
	//fmt.Printf(greetings.PrintGreetings(name))
}
func main(){
    // Your code here!


    b := []byte("ASAD")

    //b =  ConcatSLice(b)


		fmt.Println(ConcatSLice(b))

}

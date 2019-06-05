package main

import(
	"strconv"
	"fmt"
	"os"
	"bufio"
)


func getInt(x string) (int64,error){
	return strconv.ParseInt(x, 10, 64)
}

func main(){
	
	for {
        	consolereader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter your input")
        	input,_ ,err := consolereader.ReadLine() // this will prompt the user for input
		input_str := string(input)
		fmt.Println("Input is ",input_str)
        	if err != nil {
             		fmt.Println(err)
             		os.Exit(1)
        	} else {
			if(input_str == "STOP"){
				fmt.Println("STOP received. Exiting the program")
				os.Exit(0)
			} else {
				a, b := getInt(input_str)
				if b != nil{
					fmt.Println("Please enter a valid integer")
				} else {
					fmt.Printf("You entered : %v\n", a)
				}
			}
		}
    	}

}

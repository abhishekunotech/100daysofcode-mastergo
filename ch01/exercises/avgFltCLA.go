package main

import(
	"strconv"
	"os"
	"fmt"
)


func checkNumeric(x string) bool{
	_, err := strconv.ParseFloat(x, 64)
	if err != nil{
		return false
	} else {
		return true
	}
}

func main(){
	if len(os.Args) == 1 {
                fmt.Println("Please give one or more floats.")
                os.Exit(1)
        }

        arguments := os.Args
        //var err error = errors.New("An error")
        //k := 1
        var sum float64
	sum = 0
	//var n float64
        /*for err != nil {
               if k >= len(arguments) {
                        fmt.Println("None of the arguments is a float!")
                        return
                }
                n, err = strconv.ParseFloat(arguments[k], 64)
		fmt.Printf("n is : %v\n", n)
                if err == nil{
			sum += n
		}
		k++
        }*/

	cont := 0
	for _,y := range(arguments){
		if(checkNumeric(y)){
			temp,_ :=strconv.ParseFloat(y, 64)
			sum += temp
			cont++
		}
	}
	if(cont > 0){
		fmt.Printf("Average is %v\n",sum/float64(cont))
	}else{
		fmt.Println("None of the arguments was numeric")
	}
}

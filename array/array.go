package main

import "fmt"

func main() {
	var arrayOfAnu = [4]string{"anu1", "anu2", "anu3", "anu4"} // ini kita ngasi tau secara explisit length array

	for i := 0; i < len(arrayOfAnu); i++ {
		fmt.Println(i, arrayOfAnu[i])
	}
	fmt.Println(arrayOfAnu[2])


	var arrayOfAnuLainnya = [...]string{"anu1", "anu2", "anu3", "anu4", "anu5"} // ini kita ga ngasi tau secara explisit length array
	for index, option_ := range arrayOfAnuLainnya {
		fmt.Println(index, option_)
	}
}
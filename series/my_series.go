package series

import "fmt"

func GetFibonacciSerie(n int) []int {
	ret := []int{1,1}
	for i:=2; i<n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

func square(n int) int {
	return n*n
}


func init(){
	fmt.Println("init1")
}

func init(){
	fmt.Println("init1")
}

func init(){
	fmt.Println("init1")
}
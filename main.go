package main
import "fmt"

func main(){
arr:=[]int{1,2,3,3,2,1,4,5,5,6,7}
countfrequency(arr)
}

func countfrequency(num []int){

	countmap :=make(map[int]int)

	for _,v := range num{
		countmap[v]++
	}

	for k,v := range countmap{
		if v>1{
			fmt.Printf("%d appears %d times",k,v)
		}
	}


}
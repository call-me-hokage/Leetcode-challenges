package main


import "fmt"


func main(){
	number_list := []int{2,7,11,15}
	target := 9
	two_number(number_list[:],target)

}


func two_number(number_list_array [] int,target int){
	for i,num := range number_list_array{
		missing := target - num
		for j,n := range number_list_array{
			if n == missing && j!=i{
				fmt.Println(i,j)
			}
		}

	}

}
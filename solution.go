package main
import "fmt"

func reverseList(list []int) []int{
	for i,j :=0, len(list)-1; i<j; {
		list[i],list[j] = list[j],list[i]
		i++
		j--
	}
	return list
}

func addition(l1, l2 []int)[]int{
	maxlength := len(l1)
	if len(l2) >  len(l1){
		maxlength = len(l2)
	}
	var result []int
	carry := 0
	for i :=0; i < maxlength; i++ {
		digit1 := 0
		digit2 := 0
		if i < len(l1){
			digit1 = l1[i]
		}
		if i < len(l2){
			digit2 = l2[i]
		}
		tmp := digit1 + digit2 + carry
		carry = tmp/10
		digit := tmp%10
		result = append(result,digit)
	}
	if carry > 0 {
		result = append(result,carry)
	}
	fmt.Println(result)
	return result
}

func main(){
	l1_list := [] int{9,9,9,9,9,9,9,1,2,3,4,5,6,7,8,9,0}
	l2_list := [] int{9,9,9,9,9}
	reversedList1 := reverseList(l1_list)
	reversedList2 := reverseList(l2_list)
	addition(reversedList1,reversedList2)
}

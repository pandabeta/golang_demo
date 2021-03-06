package main
import "fmt"

func select_sort(a []int){
	len:=len(a)
	for i:=0;i<len-1;i++{
		k:=i
		j:=i+1
		for ;j<len;j++{
			if a[j]<a[k]{k=j}
		}
		if k!=i{
			a[i],a[k]=a[k],a[i]
		
		}
		print_array(a)
		fmt.Printf("\n")
	}
}

func quicksort(a []int,left int,right int){
	if left>right{return}
	i:=left
	j:=right
	pivot:=a[left]
	pivotindex:=left

	for i<j{
		for ;i<j&&pivot<=a[j];j--{}
		if i<j{
			a[pivotindex]=a[j]
			pivotindex=j
			}
		for ;i<j&&a[i]<=pivot;i++{}
		if i<j{
			a[pivotindex]=a[i]
			pivotindex=i
			}
	}
	a[pivotindex]=pivot

	quicksort(a,left,i-1)
	quicksort(a,i+1,right)
}
		

func print_array(a []int){
	for i:=0;i<len(a)-1;i++{
		fmt.Printf("%d, ",a[i])
	}
	fmt.Print(a[len(a)-1])
}

func main(){
	a:=[]int{9, 8, 5, 9, 4, 3, 6, 6,3,2,65,4,233,6,2}
	print_array(a)
	fmt.Printf("\n")
	quicksort(a,0,len(a)-1)
	print_array(a)
}

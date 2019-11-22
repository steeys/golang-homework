package main
  
import "fmt"

func main() {
    fmt.Println(Multiply(0.0013,2.12))
    fmt.Println(Fibonacci1(8))
    fmt.Println(Fibonacci2(9))

    var a = [] float64 {5, 4, 9, 7, 0}
    
    bubble_sort(a)
    fmt.Println(a)
    var b = [] int {1, 2, 3, 4, 1, 2, 2, 3, 2}
    fmt.Println(unique_count(b))

}

  
func Multiply (a , b float64) float64 {
    var sum float64 = 0

    for i:=b; i>0; i=i-0.0001{
    	sum = sum + a
    }
    return sum/10000
}

func Fibonacci1 (position int) int {
	var f1,f2, fsum int = 0, 1, 0
	if position == 1 {
		fsum = f1
	}
	if position == 2{
		fsum = f2
	}
	for i:=2; i<position; i++{
		fsum = f1 + f2
		f1 = f2
		f2 = fsum
	}
	return fsum
}

func Fibonacci2 (position int) int{
	if position == 1{
		return 0
	}
	if position == 2{
		return 1
	}else{
		return Fibonacci2(position - 1) + Fibonacci2(position - 2)
	}
}

func bubble_sort (a[] float64){
	var x float64
	for i:=0; i<len(a); i++{
		for j:=len(a)-1; j>i; j--{
			if a[j-1] > a[j]{
				x = a[j-1]
				a[j-1] = a[j]
				a[j] = x	
			}
		}
	}
}

func bubble_sort2 (a[] int){
	var x int
	for i:=0; i<len(a); i++{
		for j:=len(a)-1; j>i; j--{
			if a[j-1] > a[j]{
				x = a[j-1]
				a[j-1] = a[j]
				a[j] = x	
			}
		}
	}
}

func unique_count (a[] int) int{
	var sum_uniq int = 1
	bubble_sort2(a)
	for i:=1; i<len(a); i++{
		if a[i-1] != a[i]{
			sum_uniq++
		}
	}
	return sum_uniq

}

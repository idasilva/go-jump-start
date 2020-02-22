package principles

import "fmt"

var (
 array = [10]int{0,1,2,3,4,5,6,7,8,9}
)

func LoopForWithRange(){
	for i := range array{
		fmt.Println(i)
	}
	fmt.Println("loop for with range")
}

func LoopForTraditional(){
	for i := 0; i<=9; i++{
		fmt.Println(array[i])

	}
	fmt.Println("loop for traditional")

}

func LoopForWithBreak(){
	for i := 0; i<=9; i++{
		if i == 4 {
			break
		}
		fmt.Println(array[i])
	}
	fmt.Println("loop for with break")
}

func LoopForWithContinue(){
	for i := 0; i<=9; i++{
	    if i == 4 {
	        break
	    }
	    fmt.Println(array[i])
	}
	fmt.Println("loop for with break")


}

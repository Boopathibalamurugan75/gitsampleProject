package main

import (
	"fmt"
	"strconv"
)

func main() {
	arrayDeclaration()
	// matchTargetValue()
	// SubsequentPosition()
	// sorting()
}

func arrayDeclaration() {
	// check := false
	// values := [5]int{}
	// fmt.Println("Enter any ", len(values), " integer values : ")
	// for i := 0; i < len(values); i++ {
	// 	_, err := fmt.Scanln(&values[i])
	// 	if err != nil {
	// 		fmt.Println("Please enter integer value ")
	// 		var buffer string
	// 		fmt.Scanln(&buffer)
	// 		check = false
	// 		break
	// 	}
	// }
	// if check {
	// 	integerValues(values)
	// }

	stringValue := [5]string{}
	fmt.Println("Enter any ", len(stringValue), " String values : ")
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	for i := 0; i < len(stringValue); i++ {
		fmt.Scanln(&stringValue[i])
		_, err := strconv.Atoi(stringValue[i])
		if err == nil {
			fmt.Println("Please enter string value ")
			return
		}
	}
	stringValues(stringValue)
}

// trigger := false
// var (
// 	stringValue [5]int
// 	trigger bool=false
// )
// for i := 0; i < len(stringValue); i++ {
// 	fmt.Scanln(&stringValue[i])
// 	for _, input := range stringValue[i] {
// 		if input > 'a' || input == '\n' || input == ' ' {
// 			trigger = true
// 		} else {
// 			trigger = false
// 			fmt.Println("Please Enter only String value...")
// 			return
// 		}
// 	}
// }

func integerValues(value [5]int) {
	fmt.Println("\n The values is : ", value, "\n Length of the Array :", len(value), "\n capacity of the Array : ", cap(value))
}
func stringValues(stringValues [5]string) {
	fmt.Println("\n The values is : ", stringValues, "\n Length of the Array :", len(stringValues), "\n capacity of the Array : ", cap(stringValues))
}

func matchTargetValue() {
	var (
		a, countValue, targetValue int
		result                     [5]int
	)
	fmt.Println("Enter ", len(result), " values : ")
	for i := 0; i < len(result); i++ {
		_, err := fmt.Scanln(&result[i])
		if err != nil {
			fmt.Println("Please enter integer value ")
			return
		}
	}
	fmt.Println("Your array is: ", result)
	fmt.Printf("Provide the target value : ")
	fmt.Scanln(&targetValue)
	for i := 0; i < len(result); i++ {
		if result[i] == targetValue {
			for outer_Loop := 0; outer_Loop < len(result); outer_Loop++ {
				for inner_Loop := outer_Loop + 1; inner_Loop < len(result); inner_Loop++ {
					a = result[outer_Loop] + result[inner_Loop]
					if targetValue == a {
						countValue += 1
						if countValue == 1 {
							fmt.Printf("Shortest Path : %d,%d-%d,%d\n", outer_Loop, inner_Loop, result[outer_Loop], result[inner_Loop])
						} else {
							fmt.Printf("other possibilities : \n%d,%d-%d,%d\n", outer_Loop, inner_Loop, result[outer_Loop], result[inner_Loop])
						}
					}
				}
			}
			return
		}
	}
	if true {
		fmt.Println("Your target value doesnot exist in that array")
	}
}

func SubsequentPosition() {
	var (
		values [8]int
		first  int
	)
	fmt.Println("Enter any number : ")
	_, err := fmt.Scanln(&first)
	if err != nil {
		fmt.Println("Please Enter Integer Value")
		return
	}
	values[0] = first
	for j := 0; j < len(values)-1; j++ {
		values[j+1] = values[j] * 2
		fmt.Printf("array_nums[%d]=%d\n", j, values[j])
	}
}

func sorting() {
	sorting := [5]int{}
	fmt.Println("Enter any", len(sorting), "numbers : ")
	for i := 0; i < len(sorting); i++ {
		_, err := fmt.Scanln(&sorting[i])
		if err != nil {
			fmt.Println("Please Enter integer value ")
			var buffer string
			fmt.Scanln(&buffer)
			return
		}
	}
	fmt.Printf("Before Sorting Array :%d\n", sorting)
	acending := acending(sorting)
	fmt.Println("After Ascending order : ", acending)
	descending := descending(sorting)
	fmt.Println("After Descending order : ", descending)
}

func acending(sorting [5]int) (ascending [5]int) {
	for i := 1; i < len(sorting); i++ {
		if sorting[i] < sorting[i-1] {
			sorting[i], sorting[i-1] = sorting[i-1], sorting[i]
			i = 0
		}
	}
	return
}

func descending(sorting [5]int) (descending [5]int) {
	for i := 1; i < len(sorting); i++ {
		if sorting[i] > sorting[i-1] {
			sorting[i], sorting[i-1] = sorting[i-1], sorting[i]
			i = 0
		}
	}
	return
}

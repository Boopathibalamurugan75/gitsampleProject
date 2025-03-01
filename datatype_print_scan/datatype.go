package main

import (
	"fmt"
	"os"
)

func main() {
	scanfPrintf()
	// structConcept()
	// arrayConcept()
	// formatSpecifier()
	// inputGetter()
	// defaultValues()
}

func scanfPrintf() {
	fmt.Println("--------------------------------------")
	var name string
	var day, month, year int
	var mobile int
	fmt.Printf("Enter your name : ")
	fmt.Scanln(&name)
	fmt.Printf("Enter your birthday : ")
	fmt.Scan(&day, &month, &year)
	fmt.Printf("Enter your Mobile number : ")
	fmt.Scanf("%d", &mobile)
	fmt.Printf("%s\n", name)
	fmt.Printf("%d / %d / %d \n", day, month, year)
	fmt.Println("mobile unit", mobile)
}

type details struct {
	personName string
	personAge  int
}

func structConcept() {
	fmt.Println("--------------------------------------")
	personDetail := details{personName: "Raj"}
	fmt.Printf("Print Struct values : %v\n", personDetail)
	fmt.Printf("Print Struct key value : %+v\n", personDetail)
	fmt.Printf("Print Struct Syntax Format %#v\n", personDetail)
}

func arrayConcept() {
	fmt.Println("--------------------------------------")
	d := []int{1, 2, 3, 4}
	fmt.Printf("Print Array values : %v\n", d)
	fmt.Printf("Print Array key value : %+v\n", d)
	fmt.Printf("Print Array Syntax Format : %#v\n", d)
	fmt.Printf("Print type : %T\n", d)
}

func formatSpecifier() {
	fmt.Println("--------------------------------------")
	stringName := "boopathi"
	fmt.Printf("Strind :  %s\n", stringName)
	fmt.Printf("Integer : %d\n", 5)
	fmt.Printf("Float : %f\n", 6.0)
	fmt.Printf("Float .2 :%.2f\n", 5.0)
	fmt.Printf("Boolean : %t\n", true)
}

func inputGetter() {
	fmt.Println("--------------------------------------")
	var x, y, b int
	b = 10
	var name string = "Boopathiraj"
	sprint := fmt.Sprint("This is morning ", name)
	fmt.Println(sprint)
	fmt.Printf("%T\n", sprint)
	fmt.Printf("Before convert the number %d is integer", b)
	sprintf := fmt.Sprintf(" %d", &b)

	fmt.Printf("\nNow the number is converted to %d is %T\n", b, sprintf)
	fmt.Printf("%T\n", sprintf)

	fmt.Sscanf("10 20", "%d %d", &x, &y)
	fmt.Println(x, y)

	fmt.Fprint(os.Stdout, "Hello world")
}

func defaultValues() {
	fmt.Println("\n--------------------------------------")
	var intDefaultValue int
	fmt.Println("Integer DefaultValue : ", intDefaultValue)

	var floatDefaultValue float32
	fmt.Println("Float DefaultValue : ", floatDefaultValue)

	var unitDefaultValue uint
	fmt.Println("Unit DefaultValue : ", unitDefaultValue)
	var unitMinusDefaultValue uint = 0
	fmt.Println("Unit Minusvalue : ", unitMinusDefaultValue-3)

	var runeDefaultValue rune
	fmt.Println("Rune DefaultValue : ", runeDefaultValue)

	var byteDefaultValue byte
	fmt.Println("Byte DefaultValue : ", byteDefaultValue)
	var byteMinusDefaultValue byte = 0
	fmt.Println("Byte Minus Value : ", byteMinusDefaultValue-3)

	var characterByte byte = 'b'
	fmt.Println("Print character Byte : ", characterByte)

	var arr = [5]int{}
	fmt.Println("Array Default value : ", arr)

	var complexDefaultVale complex64
	fmt.Println("Complex Default value : ", complexDefaultVale)
}

// data tpyes in Go
// cd go/src/day1
// go run datatype.go
// output:
package main

import "fmt"

func main() {

	// Built-in Basic Types:
	// ประกาศตัวแปรแบบระบุชนิดข้อมูล
	var name string = "Hello, Go!"
	var age int = 30
	var height float64 = 5.9
	var isEmployed bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Employed:", isEmployed)

	// แบบย่อ
	name2 := "Hello, Go!"
	age2 := 30
	height2 := 5.9
	isEmployed2 := true

	fmt.Println("Name:", name2)
	fmt.Println("Age:", age2)
	fmt.Println("Height:", height2)
	fmt.Println("Employed:", isEmployed2)

	// Composite Types:
	// slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// การเข้าถึงสมาชิกใน slice
	fmt.Println("First Number:", numbers[0])
	fmt.Println("Second Number:", numbers[1])
	fmt.Println("Third Number:", numbers[2])
	fmt.Println("Fourth Number:", numbers[3])
	fmt.Println("Fifth Number:", numbers[4])

	// การเพิ่มสมาชิกใน slice
	numbers = append(numbers, 6)
	fmt.Println("Updated Numbers:", numbers)

	// map
	person := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Eve":   28,
	}
	fmt.Println("Person:", person)
	fmt.Println("Alice's Age:", person["Alice"])

	// การเพิ่มสมาชิกใน map
	person["Charlie"] = 35
	fmt.Println("Updated Person:", person)

	// การลบสมาชิกใน map
	delete(person, "Bob")
	fmt.Println("Updated Person:", person)

	// struct
	type Product struct {
		Name   string
		ItemID int
		Price  float64
	}

	newProduct := Product{
		Name:   "Air purifier",
		ItemID: 101,
		Price:  199.99,
	}

	fmt.Println("Product Name:", newProduct.Name)
	fmt.Println("Product ItemID:", newProduct.ItemID)
	fmt.Println("Product Price:", newProduct.Price)
}

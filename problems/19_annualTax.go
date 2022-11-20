package main

import "fmt"

func main() {

	var salary int

	fmt.Println("Please enter your annual salary")
	fmt.Scan(&salary)
	taxCalc(salary)

}

func taxCalc(salary int) {

	switch true {
	case salary < 250000:
		fmt.Println("No Tax")
	case salary > 249999 && salary < 500000:
		tax := (salary * 5) / 100
		fmt.Println(tax)
	case salary > 499999 && salary < 1000000:
		tax := (salary * 20) / 100
		fmt.Println(tax)
	case salary > 999999 && salary < 5000000:
		tax := (salary * 30) / 100
		fmt.Println(tax)
	}
}

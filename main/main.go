package main

import (
	"fmt"
	"math"
)

func main() {

	// fmt.Println(second())

	// third()

	//forth()
	//circle()
	//centimrter()
	//celcius()
	//fahrenheit()
	//day()
	//power()
	//squreroot()
	//traingle_angle()
	//traingle_area()
	//equi_traingle()
	//subject()
	interest()
	//compound()
}

func second() int {
	var a, b int

	fmt.Println("Enter First Number")
	fmt.Scanln(&a)
	fmt.Println("Enter second Number")
	fmt.Scanln(&b)

	return a + b

}

func third() {
	var a, b int

	fmt.Println("Enter First Number")
	fmt.Scanln(&a)
	fmt.Println("Enter second Number")
	fmt.Scanln(&b)

	fmt.Println("Addition :", a+b)
	fmt.Println("Subtraction :", a-b)
	fmt.Println("Multiplicatin :", a*b)
	fmt.Println("div :", a/b)
}

func forth() {
	var a, b int

	fmt.Println("Enter length")
	fmt.Scanln(&a)
	fmt.Println("Enter breath")
	fmt.Scanln(&b)

	fmt.Println("perimeter :", 2*(a+b))

	fmt.Println("area :", a*b)

}

func circle() {
	var a float32
	var pie float32 = 3.14
	fmt.Println("Enter radius")
	fmt.Scanln(&a)

	fmt.Println("Dimeter :", 2*(a))

	fmt.Println("perimeter :", 2*pie*a)
	fmt.Println("area :", pie*a*a)

}

func centimrter() {
	var a float32

	fmt.Println("Length in Centirmeter")
	fmt.Scanln(&a)

	fmt.Println("Length in Meter :", a/100)

	fmt.Println("Length in KM :", a/100/1000)
}

func celcius() {
	var a float32
	fmt.Println("Temp in celcius ")
	fmt.Scanln(&a)
	fmt.Println("Temp in Fahrenheit :", (a*1.8)+32)

}

func fahrenheit() {
	var a float32
	fmt.Println("Temp in Fahrenheit")
	fmt.Scanln(&a)
	fmt.Println("Temp in celcius  :", (a-32)/1.8)

}

func day() {
	var a float32
	fmt.Println("Enter Days")
	fmt.Println("Eter Days")
	fmt.Scanln(&a)
	fmt.Println("Years  :", a/35)
	fmt.Println("Years  :", a/7)

}

func power() {
	var a, b float64
	fmt.Println("Eter X")
	fmt.Scanln(&a)
	fmt.Println("nter Y")
	fmt.Scanln(&b)
	fmt.Println("X^Y  :", math.Pow(a, b))

}

func squreroot() {
	var a float64
	fmt.Println("Enter X")
	fmt.Scanln(&a)

	fmt.Println("Squareroot  :", math.Sqrt(a))

}

func traingle_angle() {
	var a, b float64
	fmt.Println("Angle One")
	fmt.Scanln(&a)
	fmt.Println("Angle Two")
	fmt.Scanln(&b)
	fmt.Println("Third angle  :", 180-a-b)

}

func traingle_area() {
	var a, b float64
	fmt.Println("Traingle Base")
	fmt.Scanln(&a)
	fmt.Println("Traingle Height")
	fmt.Scanln(&b)
	fmt.Println("Traingle Area  :", (a*b)/2)

}

func equi_traingle() {
	var a float64
	fmt.Println("Traingle Length")
	fmt.Scanln(&a)

	fmt.Println("Traingle Area  :", (math.Sqrt(3)/4)*a*a)

}

func subject() {
	var a, b, c, d, e float64
	fmt.Println("subject 1")
	fmt.Scanln(&a)
	fmt.Println("subject 2")
	fmt.Scanln(&b)
	fmt.Println("subject 3")
	fmt.Scanln(&c)
	fmt.Println("subject 4")
	fmt.Scanln(&d)
	fmt.Println("subject 5")
	fmt.Scanln(&e)

	fmt.Println("Total  :", a+b+c+d+e)
	fmt.Println("AVG  :", (a+b+c+d+e)/5)
	fmt.Println("Percentage  :", ((a+b+c+d+e)/500)*100)
}

func interest() {
	var a, b, c float64
	fmt.Println("Principle Amount")
	fmt.Scanln(&a)
	fmt.Println("Time")
	fmt.Scanln(&b)
	fmt.Println("Rate of interest")
	fmt.Scanln(&c)

	fmt.Println("Simple interest  :", (a*b*c)/100)
	//fmt.Println("Compound interest  :", (a*b*c)/100)

}


func compound()  {
	var Pamount, InterestRate, timePeriod, ciFuture, compoundI float64

    fmt.Print("Enter the Total or Principal Amount = ")
    fmt.Scanln(&Pamount)

    fmt.Print("Enter the Interest Rate = ")
    fmt.Scanln(&InterestRate)

    fmt.Print("Enter the Total number of Years = ")
    fmt.Scanln(&timePeriod)

    ciFuture = Pamount * (math.Pow((1 + InterestRate/100), timePeriod))
    compoundI = ciFuture - Pamount

    fmt.Println("\nThe Compound Interest  = ", compoundI)
    fmt.Println("The Future Compound Interest  = ", ciFuture)
}
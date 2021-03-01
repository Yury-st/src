package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	var myDate Date
	myDate.SetYear(2019)

	fmt.Println(myDate)
}

func (d *Date) SetYear(year int) {
	d.Year = year
}

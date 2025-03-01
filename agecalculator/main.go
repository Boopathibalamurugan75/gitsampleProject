package main

import (
	"fmt"
	"time"
)

func main() {
	var dateOfBirth, motherDateOfBirth string
	fmt.Print("Enter your Date of birth  (DD-MM-YYYY) : ")
	fmt.Scanln(&dateOfBirth)
	fmt.Print("Enter your Mother Date of birth  (DD-MM-YYYY) : ")
	fmt.Scanln(&motherDateOfBirth)

	myBirthDate, err := time.Parse("02-01-2006", dateOfBirth)
	motherBirthDate, err := time.Parse("02-01-2006", motherDateOfBirth)
	birthMonth := int(myBirthDate.Month())
	birthDay := myBirthDate.Day()
	if err != nil {
		fmt.Println("Error when parsing date:", err)
		return
	}
	//Find my age
	age := calculateAge(myBirthDate)
	fmt.Printf("My Age is : %d years\n", age)
	//Find Mother age
	motherAge := motherAgeCalc(motherBirthDate, age)
	fmt.Printf("My mother Age when I was born : %d years\n", motherAge)
	//Find Days left my Birthday
	daysLeft := DaysUntilBirthday(birthMonth, birthDay)
	fmt.Printf("Days left for my Birthday : %d days\n", daysLeft)
}

func calculateAge(birthDate time.Time) int {
	today := time.Now()
	age := today.Year() - birthDate.Year()
	if today.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}

func motherAgeCalc(motherBirthDate time.Time, ageOfSon int) int {
	today := time.Now()
	age := today.Year() - motherBirthDate.Year()
	if today.YearDay() < motherBirthDate.YearDay() {
		age--
	}
	return age - ageOfSon
}

func DaysUntilBirthday(birthMonth, birthDay int) int {
	today := time.Now()
	birthday := time.Date(today.Year(), time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.Local)
	daysLeft := birthday.YearDay() - today.YearDay()
	if daysLeft < 0 {
		daysLeft = (365 - today.YearDay()) + birthday.YearDay()
		if time.Date(today.Year(), 12, 31, 0, 0, 0, 0, time.Local).YearDay() == 366 {
			daysLeft++
		}
	}
	return daysLeft
}

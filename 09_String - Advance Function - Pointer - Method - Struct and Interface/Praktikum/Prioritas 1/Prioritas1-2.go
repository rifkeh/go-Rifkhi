package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

type StudentSlice []Student

func (s StudentSlice) AverageScore() float64 {
	totalScore := 0
	for _, i := range s {
		totalScore += i.Score
	}
	return float64(totalScore / len(s))
}

func main() {
	person := make(StudentSlice, 5)
	person[0].Name = "Rizky"
	person[1].Name = "Kobar"
	person[2].Name = "Ismail"
	person[3].Name = "Umam"
	person[4].Name = "Yopan"
	person[0].Score = 80
	person[1].Score = 75
	person[2].Score = 70
	person[3].Score = 75
	person[4].Score = 60
	fmt.Println(person.AverageScore())
}

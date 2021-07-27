package main

import "fmt"

// practical use of first class fun.
// In this program we pass fun as argument to the fun and filters student grade

type student struct {
	name       string
	country    string
	grade      string
	age        int
	bloodGroup string
}

// return list of student having filter fun...
func getStudent(s []student, filter func(s student) bool) []student {

	var filterStudents []student

	for _, v := range s {
		if filter(v) == true {
			filterStudents = append(filterStudents, v)
		}
	}

	return filterStudents

}

func main() {

	s1 := student{name: "aakash", country: "India", grade: "A", age: 18, bloodGroup: "0+"}
	s2 := student{name: "aakash12", country: "India", grade: "A", age: 20, bloodGroup: "0-"}
	s3 := student{name: "Ram", country: "India", grade: "B", age: 21, bloodGroup: "0+"}

	f := func(s student) bool {

		if s.grade == "B" {
			return true
		} else {
			return false
		}
	}

	stList := []student{s1, s2, s3, {name: "Inline", country: "Usa", grade: "B", age: 18, bloodGroup: ""}}
	fmt.Println("List of Grade B students:: ", getStudent(stList, f))

	// second way to call get A grade student

	stA := getStudent(stList, func(s student) bool {

		if s.grade == "A" {
			return true
		} else {
			return false
		}
	})

	fmt.Println("List of Grade A students::: ", stA)

	stBlood := getStudent(stList, func(s student) bool {
		if s.bloodGroup != "" {
			return true
		} else {
			return false
		}
	})

	fmt.Println("List of student having blood group present::::", stBlood)

	stAge := getStudent(stList, func(s student) bool {
		if s.age > 18 {
			return true
		}else {
			return false
		}
	})

	fmt.Println("List of student whose age is 18+ :: ", stAge)

}

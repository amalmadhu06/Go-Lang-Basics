// we will calculate total score of a student using struct method here

package main

import "fmt"

type scoreCard struct {    						//struct is created
	maths   int
	science int
	english int
}

func (student *scoreCard) totalScore() int {	//struct method to calculate total score
	totalScore := student.maths + student.english + student.science
	fmt.Println(totalScore)
	return totalScore
	
}


func main() {

	var amal scoreCard = scoreCard{        //created an object with type as scoreCard
		maths:   60,
		science: 70,
		english: 57,
	}
	amalPointer := &amal					//assigned pointer of amal to amalPointer

	amalPointer.totalScore()				//calling struct method on the pointer
}
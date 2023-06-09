package main

import (
	"fmt"
)

/**
* 1. Consistent formatting
* 2. Comments
* 3. Variable and function naming
* 4. Return early
 */

type Student struct {
	ClassID int
	Name string
	Scores []float64
	Average float64
	Grade string
}

type Class struct {
	ID int
	Name string
	HighestAverage float64
}

func main() {

	classes := []*Class{
		{
		ID: 1,
		Name: "Class 1",
		HighestAverage: 0,
		},
		{
		ID: 2,
		Name: "Class 2",
		HighestAverage: 0,
		},
	}

	students := []*Student{
		{
		ClassID: 1,
		Name: "Budi",
		Scores: []float64{80, 90, 88},
		Average: 0,
		},
		{
		ClassID: 2,
		Name: "Tono",
		Scores: []float64{70, 77, 75},
		Average: 0,
		},
		{
		ClassID: 1,
		Name: "Alexa",
		Scores: []float64{60, 67, 80},
		Average: 0,
		},
		{
			ClassID: 2,
			Name: "Tina",
			Scores: []float64{90, 91, 86},
			Average: 0,
			},
		{
			ClassID: 3,
			Name: "",
			Scores: []float64{123},
			Average: 0,
		},
	}

	for _, s := range students {
			t:=float64(0)

			for _, sc := range s.Scores {
				t+=float64(sc)
			}

			s.Average= t / float64(len(s.Scores))
			g := ""

			if s.Average > 80 {
				g = "A"
			} else if s.Average > 65 {
				g = "B"
			} else if s.Average > 50 {
				g = "C"
			} else if s.Average > 40 {
				g = "D"
			} else {
				g = "E"
			}

			s.Grade = g
			className := ""

			for _, c := range classes {
				if c.ID == s.ClassID {
					if c.HighestAverage < s.Average {
						c.HighestAverage = s.Average
					}
					className = c.Name
				}
			}

			fmt.Printf("Student %s [%s] - Grade %s (%v)\n", s.Name, className, s.Grade, s.Average)
	}

	for _, c := range classes {
		fmt.Printf("Class %s - Highest Average %v\n", c.Name, c.HighestAverage)
	}
}
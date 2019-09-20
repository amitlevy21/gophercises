package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type quizSpec struct {
	score             int
	totalNumQuestions int
}

func main() {
	var problemFile string
	flag.StringVar(&problemFile, "problem_file", "problems.csv", "A csv file that contains problems with format of question,answer")
	flag.Parse()

	f, err := os.Open(problemFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	spec := quiz(r)
	fmt.Printf("You scored %d/%d\n", spec.score, spec.totalNumQuestions)
}

func quiz(r *csv.Reader) (spec quizSpec) {
	for {
		problem, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		question := problem[0]
		solution := problem[1]
		fmt.Printf("Problem #%d: %s = ", spec.totalNumQuestions+1, question)
		userAnswer := getUserAnswer(question)
		spec.score += getAnswerScore(solution, userAnswer)
		spec.totalNumQuestions++
	}

	return spec
}

func getUserAnswer(question string) string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func getAnswerScore(solution, userAnswer string) int {
	if userAnswer == solution {
		return 1
	}
	return 0
}

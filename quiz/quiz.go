package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type quizSpec struct {
	score             int
	totalNumQuestions int
	currentQuestion   int
}

func main() {
	var problemFile string
	var limit int
	getCommandFlags(&problemFile, &limit)

	f, err := os.Open(problemFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	problems := getProblems(f)
	serveQuiz(problems, limit)
}

func getProblems(f *os.File) [][]string {
	r := csv.NewReader(f)
	problems, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return problems
}

func getCommandFlags(problemFile *string, limit *int) {
	flag.StringVar(problemFile, "problem_file", "problems.csv", "A csv file that contains problems with format of question,answer")
	flag.IntVar(limit, "limit", 30, "time to solve the quiz in seconds")
	flag.Parse()
}

func serveQuiz(problems [][]string, limit int) {
	spec := quizSpec{totalNumQuestions: len(problems)}
	time.AfterFunc(time.Second*time.Duration(limit), func() { endQuiz(&spec) })

	for _, problem := range problems {
		serveProblem(problem, &spec)
	}
	endQuiz(&spec)
}

func endQuiz(spec *quizSpec) {
	fmt.Printf("\nYou scored %d/%d\n", spec.score, spec.totalNumQuestions)
	os.Exit(0)
}

func serveProblem(problem []string, spec *quizSpec) {
	question := problem[0]
	solution := problem[1]
	spec.currentQuestion++
	fmt.Printf("Problem #%d: %s = ", spec.currentQuestion, question)
	userAnswer := getUserAnswer(question)
	spec.score += getAnswerScore(solution, userAnswer)
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

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	chosenYear := ""
	chosenQuestion := ""
	err := errors.New("")
	for {
		chosenYear, err = retrieveYearFromCLI()
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	for {
		chosenQuestion, err = retrieveQuestionNumberFromCLI()
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	fmt.Println(fmt.Sprintf("You have selected year: %v, question: %v", chosenYear, chosenQuestion))
}

func retrieveYearFromCLI() (string, error) {
	// read input using go std lib bufio, getting input from cmdline
	reader := bufio.NewReader(os.Stdin)

	// get the appropriate year from the user
	possibleYears := []string{"2015", "2016", "2017", "2018", "2019", "2020", "2021", "2022", "2023"}
	fmt.Println("Please select a year : ")
	selectedYear, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("Could not read from command line. Please retry with valid input.")
	}

	selectedYear = strings.TrimSuffix(selectedYear, "\n")

	if !slices.Contains(possibleYears, selectedYear) {
		return "", errors.New(fmt.Sprintf("Not a valid year. Please retry from the following options: %v", possibleYears))
	}

	return selectedYear, nil
}

func retrieveQuestionNumberFromCLI() (string, error) {
	// read input using go std lib bufio, getting input from cmdline
	reader := bufio.NewReader(os.Stdin)

	// get the appropriate question number from the user
	fmt.Println("Please select a question : ")
	selectedQuestionNumber, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("Could not read from command line. Please retry with valid input.")
	}

	selectedQuestionNumber = strings.TrimSuffix(selectedQuestionNumber, "\n")
	selectedQuestionNumberAsInt, err := strconv.Atoi(selectedQuestionNumber)

	if selectedQuestionNumberAsInt < 1 || selectedQuestionNumberAsInt > 25 {
		return "", errors.New(fmt.Sprintf("Not a valid Question Number. Please pick a number between 1 and 25, inclusively."))
	}

	return selectedQuestionNumber, nil
}

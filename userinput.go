package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	chosenYear := ""
	err := errors.New("")
	for {
		chosenYear, err = retrieveValidYear()
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	fmt.Println(fmt.Sprintf("You have selected year: %v", chosenYear))
}

func retrieveValidYear() (string, error) {
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

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {
	//reads all the problems from .csv file. open the file
	if fObj, err := os.Open(fileName); err == nil {
		//we will create a new reader
		csvR := csv.NewReader(fObj)
		//it will nead to read the file
		if cLines, err := csvR.ReadAll(); err == nil {
			//call the parseProblem function
			return parseProblem(cLines), nil
		} else {
			return nil, fmt.Errorf("error while reading data in csv"+"format from %s file; %s", fileName, err.Error())
		}

	} else {
		return nil, fmt.Errorf("error in opening %s file; %s", fileName, err)
	}
}

//  read the file, call the parseProblem

func main() {
	//get the name of the file
	fName := flag.String("f", "quiz.csv", "path of csv file")
	//set the duration of timer
	timer := flag.Int("t", 30, "timer for quiz")
	flag.Parse()
	//pull the problems from the file by problemPuller function
	problems, err := problemPuller(*fName)

	//handle the error
	if err != nil {
		exit(fmt.Sprintf("something went wrong:%s", err.Error()))
	}
	//create variable to count the correct answers
	correctAns := 0
	//initialize the timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
	//loop through the problems, print the questions, accept the answers
problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("Problem %d: %s= ", i+1, p.q)

		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.a {
				correctAns++
			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}
	}
	//calculate and print out the result.
	fmt.Printf("Your Result is %d out of %d\n", correctAns, len(problems))
	fmt.Printf("Press enter to exit")
	<-ansC

}

func parseProblem(lines [][]string) []problem {
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

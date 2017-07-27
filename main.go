package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"robot-simulator/robot"
	"strings"
)

func main() {
	//Handle all the available flags
	fileName := flag.String("f", "", "A file name to be executed.")
	tableWidth := flag.Int("width", 5, "Table width has to be positive number")
	tableLength := flag.Int("length", 5, "Table length has to be positive number")
	flag.Parse()

	//initialise table & robot
	rbot, err := robot.NewRobot(*tableWidth, *tableLength)
	if err != nil {
		panic(err)
	}

	//initiate with a file commands
	if *fileName != "" {
		fileCommands, err := ioutil.ReadFile(*fileName)
		check(err)

		commands := strings.Split(string(fileCommands), "\n")
		for _, command := range commands {
			fmt.Printf("Performing %s --> ", command)
			msg, err := rbot.Perform(command)
			check(err)
			fmt.Println(msg)
		}
	}

	//Command Center
	fmt.Printf("Hi, Table(W:%dxL:%d) has been prepared and I'm at your command:\n", *tableWidth, *tableLength)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg, err := rbot.Perform(scanner.Text())
		check(err)
		fmt.Println(msg)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type gameCase struct {
  step    int
  command string
  answer  string
}

func main(){
  initGame()
}

func initGame(){
  scanner := bufio.NewScanner(os.Stdin)

  for {
    command, err := getCommand("getCommand:", scanner)
  
    if err != nil {
      fmt.Println(err)
      continue
    }

    // gaming navigation logic below
    fmt.Println("Play", command)
  }  
}

func getCommand(message string, scanner *bufio.Scanner ) (string, error) {
  fmt.Println(message)
  scanner.Scan()
  scanner.Text()
  vals := strings.Fields(scanner.Text())
  
  if len(vals) == 1 {
    return vals[0], nil
  }

  return "", errors.New("WTF are you typing?")
}

func handleCommand(command string) string {
  return command
}
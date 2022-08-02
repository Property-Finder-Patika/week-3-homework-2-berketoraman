/*
Implement a number-guessing game in which the computer computes a four digit number as a secret number
and a player tries to guess that number correctly.
Player would enter her guess and the computer would produce a feedback on the positions of the digits.
Four-digit number can't start with 0 and have repeating digits.
Let's say the computer computes 2658 as a secret number to be guessed by the player.
When player enters her guess such as 1234, the computer would display -1 meaning that only one digit of 1234 exist in the secret number
and its position is wrong.
When the player enters 5678 similarly the computer displays +2-1.
And the game goes on until the player correctly guess the secret number and the computer displays +4.
The game also keeps track of all guesses entered by the players so far and lists them
when it displays its feedback to the player so that the player can compute her next guess correctly.
*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)
func main() {

	rand.Seed(time.Now().UnixNano())
	min, max := 1000, 9999
	secretNumber := rand.Intn(max-min) + min 
	fmt.Println("Secret number is generated ")
	
	var guessingNumber int
	attempt := 0
	for {
		fmt.Print("Guess a number ")
		_, err := fmt.Scanf("%d", &guessingNumber)
		fmt.Printf("Generated secret number : %d\n", secretNumber)
		if err != nil {
			return
		}
		fmt.Println("Guessing number: ", guessingNumber)
       
	//Checks the range
		if guessingNumber < 1000 || guessingNumber > 9999 {
			fmt.Println("Please enter a number between 1000 and 9999.")
		} else {
			attempt++
			if secretNumber == guessingNumber {
				fmt.Printf("Correct after %d attempt!: +4", attempt)
				break
			}

			compareTheNumbers(secretNumber, guessingNumber)
		}

	}

}

func compareTheNumbers(secretNumber int, guessingNumber int) {
	secretStr := strconv.Itoa(secretNumber)
	guessingStr := strconv.Itoa(guessingNumber)

	var secretNumDigits, guessingNumDigits []string

	//Creates a list of secret numbers
	for _, s := range secretStr {
		secretNumDigits = append(secretNumDigits, string(s))
	}

	//Creates a list of guessing numbers
	for _, g := range guessingStr {
		guessingNumDigits = append(guessingNumDigits, string(g))
	}

	correctPosition := 0
	wrongPosition := 0
	//check whether each digit of the guessing number is in the secret number.
	for i, v := range guessingNumDigits {
		if contains(secretNumDigits, v) {
			if secretNumDigits[i] == v {
				correctPosition++
			} else {
				wrongPosition++
			}

		}
	}

	//Feedback
	result := ""
	if correctPosition > 0 {
		result = "+" + strconv.Itoa(correctPosition)
	}
	if wrongPosition > 0 {
		result = result + "-" + strconv.Itoa(wrongPosition)
	}

	fmt.Printf("Not correct, try again: %s\n", result)

}

func contains(slice []string, elem string) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var board [3][3]string
const PLAYER = "X"
const COMPUTER  = "O"
//function  to reset the board
func resetBoard(){
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
		
	}
}
//function to print the board
func printBoard(){
	fmt.Printf(" %s | %s | %s ", board[0][0], board[0][1], board[0][2]);
	fmt.Printf("\n---|---|---\n");
	fmt.Printf(" %s | %s | %s ", board[1][0], board[1][1], board[1][2]);
	fmt.Printf("\n---|---|---\n");
	fmt.Printf(" %s | %s | %s ", board[2][0], board[2][1], board[2][2]);
	fmt.Printf("\n");
}
//func to check for free spaces
func checkFreeSpaces() int{
	freeSpaces := 9
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] != " " {
				freeSpaces--
			}
		}
		
	}
	return freeSpaces
}
//function to show how the player moves 
func playerMove(){
	var x int
	var y int
	for ok := true; ok; ok = (board[x][y] != " ") {
		fmt.Println("Enter row #(1-3):")
		fmt.Scanln(&x)
		x--
		fmt.Println("Enter column #(1-3):")
		fmt.Scanln(&y)
		y--
		if board[x][y] != " " {
			fmt.Println("Invalid Move!")
		}else{
			board[x][y] = PLAYER
			break
		}
	}
	
}
//function for random movement of the computer
func computerMove(){
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var x int
	var y int
	if checkFreeSpaces() >0 {
		for ok := true; ok; ok = (board[x][y] != " "){
			x = r1.Intn(3)
			y = r1.Intn(3)
		}
		board[x][y] = COMPUTER
	}else{
		printWinner(" ")
	}
	
}
//function to check winner
func checkWinner() string{
	 //check rows
	 for i := 0; i < 3; i++{
		if (board[i][0] == board[i][1] && board[i][0] == board[i][2]){
		   return board[i][0]
		}
	 }
	 //check columns
	 for i := 0; i < 3; i++{
		if(board[0][i] == board[1][i] && board[0][i] == board[2][i]){
		   return board[0][i]
		}
	 }
	 //check diagonals
	 if(board[0][0] == board[1][1] && board[0][0] == board[2][2]){
		return board[0][0]
	 }
	 if(board[0][2] == board[1][1] && board[0][2] == board[2][0]){
		return board[0][2]
	 }
  
	 return " "
}
//function to print the winner 
func printWinner(winner string){
	if winner == PLAYER{
		fmt.Println("YOU WIN!")
	}else if winner == COMPUTER{
		fmt.Println("YOU LOOSE!")
	}else{
		fmt.Println("IT'S A TIE!")
	}
}

func main(){
	var winner string
	var response string

	for ok := true; ok; ok = (response == "Y"){
		winner = " "
		response = " "
		resetBoard()
		for (winner == " " && checkFreeSpaces() != 0){
			printBoard()
			playerMove()
			winner = checkWinner()
			if winner != " " || checkFreeSpaces() != 0 {
				break
			}
			computerMove()
			if winner != " " || checkFreeSpaces() != 0 {
				break
			}
		
		}
		printBoard()
		printWinner(winner)
		fmt.Println("Would you like to play again? (Y/N)")
		fmt.Scanln(&response)
		response = strings.ToUpper(response)
	}
	fmt.Println("Thanks for Playing!")

}

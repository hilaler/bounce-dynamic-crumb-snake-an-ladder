/*	Name : Hilal Ramadhan Utomo
	SID : 1301194236
	XX. Bounce dynamic crumb snake and ladder */

package main

import "fmt"
import "math/rand"
import "time"

const N = 101

type component struct {
	head, tail, bottom, top, position, crumb, bounce int
}

var Box [N] component

func main() {
	var play, playerName string
	var playerPosition, dice, tempPos, points, snakeAndladder int

	snakeAndladder = 0
	playerPosition = 0
	points = 0
	playerName = welcomeMessage()
	checkingSnakeandLadder(&snakeAndladder)
	fmt.Printf("Press %q to roll the dice and type %q to end the game\n", "Enter", "QUIT")
	fmt.Scanln(&play)
	for play != "QUIT" && play != "quit" {
		if playerPosition < 100 {
			dice = rollDice()
			fmt.Printf("%s dice is %d\n", playerName, dice)
			playerPosition = playerPosition + dice
			checkabove100(&playerPosition, tempPos)
			Box[playerPosition].position = playerPosition
			Crumb(playerPosition)
			checkingBoxv2(&playerPosition, &points, &playerName)
			playerPosition = checkingBox(playerPosition, playerName)
			fmt.Printf("%s position is at %d with %d point\n", playerName, playerPosition, points)
			if playerPosition == 100 {
				fmt.Printf("You Win!\n%s has %d point\n", playerName, points)
			} else {
				fmt.Scanln(&play)
			}
		} else {
			play = "QUIT"
		}
	}
	endMessage()
}

func checkabove100(playerPosition *int, tempPos int) {
	if *playerPosition > 100 {
		tempPos = *playerPosition - 100
		*playerPosition = 100 - tempPos
	}
}

func checkingSnakeandLadder(snakeAndladder *int) {
	if *snakeAndladder == 0 {
		firstThingToDo()
		*snakeAndladder =+ 1
	}
}

func rollDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func firstThingToDo() {
	fmt.Println("Snakes :")
	randomingSnakes()
	fmt.Println("Ladders :")
	randomingLadders()
}

func randomingLadders() {
	var boxindex1 = 0
	var boxindex2 = 0
	i := 0
	rand.Seed(time.Now().UnixNano())
	for i < 7 {
		boxindex1 = (rand.Intn(80)) + 10
		boxindex2 = (rand.Intn(80)) + 10
		if Box[boxindex1].top == 0 && Box[boxindex1].bottom == 0 && Box[boxindex2].top == 0 && Box[boxindex2].bottom == 0 && Box[boxindex1].head == 0 && Box[boxindex1].tail == 0 && Box[boxindex2].head == 0 && Box[boxindex2].tail == 0 {
			if boxindex1 > boxindex2 {
				for (boxindex2 > (boxindex1 / 10) * 10) {
					boxindex2 = rand.Intn(80) + 10
				}
				Box[boxindex2].top = boxindex1
				Box[boxindex2].bottom = boxindex2 
				fmt.Println("bottom ladder is at :" , Box[boxindex2].bottom, "And top ladder is at :" ,Box[boxindex2].top)
				i++
			} else if boxindex2 > boxindex1 {
				for (boxindex1 > (boxindex2 / 10) * 10) {
					boxindex1 = rand.Intn(80) + 10
				}
				Box[boxindex1].top = boxindex2
				Box[boxindex1].bottom = boxindex1
				fmt.Println("bottom ladder is at :" ,Box[boxindex1].bottom, "And top ladder is at :", Box[boxindex1].top)
				i++
			} else {
				boxindex1 = (rand.Intn(80)) + 10
				boxindex2 = (rand.Intn(80)) + 10
			}
		}
	}
}

func randomingSnakes() {
	var boxindex1 = 0
	var boxindex2 = 0
	i := 0
	rand.Seed(time.Now().UnixNano())
	for i < 10 { 
		boxindex1 = (rand.Intn(64)) + 26
		boxindex2 = (rand.Intn(64)) + 26
		if Box[boxindex1].top == 0 && Box[boxindex1].bottom == 0 && Box[boxindex2].top == 0 && Box[boxindex2].bottom == 0 && Box[boxindex1].head == 0 && Box[boxindex1].tail == 0 && Box[boxindex2].head == 0 && Box[boxindex2].tail == 0 {
			if boxindex1 > boxindex2 {
				for (boxindex2 > (boxindex1 / 10) * 10) {
					boxindex2 = rand.Intn(65) + 26
				}
				Box[boxindex1].head = boxindex1
				Box[boxindex1].tail = boxindex2
				fmt.Println("head snake is at :" , Box[boxindex1].head, "tail snake is at :", Box[boxindex1].tail)
				i++
			} else if boxindex2 > boxindex1 {
				for (boxindex1 > (boxindex2 / 10) * 10) {
					boxindex1 = rand.Intn(65) + 26
				}
				Box[boxindex2].head = boxindex2
				Box[boxindex2].tail = boxindex1
				fmt.Println("head snake is at :" , Box[boxindex2].head, "tail snake is at :" ,Box[boxindex2].tail)
				i++
			} else {
				boxindex1 = (rand.Intn(65)) + 25
				boxindex2 = (rand.Intn(65)) + 25
			}
		}
	}
}

func checkingBox(playerPosition int, playerName string) int {
	if Box[playerPosition].position == Box[playerPosition].bottom {
		fmt.Printf("%s hit a ladder at %d %s go to %d\n", playerName, playerPosition, playerName, Box[playerPosition].top)
		playerPosition = Box[playerPosition].top
	} else if Box[playerPosition].position == Box[playerPosition].head {
		fmt.Printf("%s hit a head of snake at %d %s go to %d\n", playerName, playerPosition, playerName, Box[playerPosition].tail)
		playerPosition = Box[playerPosition].tail
	}
	return playerPosition
}

func checkingBoxv2(playerPosition, points *int, playerName *string){
	var temp int
	if Box[*playerPosition].position == Box[*playerPosition].crumb {
		temp = rand.Intn(10) + 1
		rand.Seed(time.Now().UnixNano())
		fmt.Printf("%s lands on crumb , %s get %d points\n", *playerName, *playerName, temp)
		*points = *points + temp
	} 
}

func welcomeMessage() string {
	var playerName string
	fmt.Println("Welcome snake and ladder game")
	fmt.Println("Write your name:")
	fmt.Scan(&playerName)
	return playerName
}

func endMessage() {
	fmt.Println("Thank you for playing this game")
}

func Crumb(playerPosition int) {
	var boxindex = 0
	rand.Seed(time.Now().UnixNano())
	boxindex = rand.Intn(100) + 1
	if boxindex <= 20 && boxindex > 10 {
		Box[playerPosition].crumb = playerPosition
	}
}
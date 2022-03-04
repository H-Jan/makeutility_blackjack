# makeutility_blackjack
A demonstration of Go skills through the creation of a BlackJack dealer

[![Go Report Card](https://goreportcard.com/badge/github.com/H-Jan/makeutility_blackjack)](https://goreportcard.com/report/github.com/H-Jan/makeutility_blackjack)


# Black Jack 

> The following is an implementation of a *Black Jack* program which takes in a user's response and plays a game of black jack with them. 

<h2> Tech Stack </h2>
This program is written in Golang and employs several critical libraries, including:

- fmt
- math/rand
- strconv
- strings
- time
 
<h2> Deployment </h2>
In order to deploy this project, simply download the code to your local I.D.E. and

```
$go build main.go

```

This will build our project

```
$go run main.go
```

This will run the black jack program while asking questions and returning the appropriate response to the user input and play the game!

<h2> How to Use </h2>

Simply answer the following questions with **Y**, **N**, or the appropriate response
Answer: 
- Hello, Enter your name
- Are you ready to play Black Jack?
- Would you like to hit?


<img width="564" alt="Black_Jack" src="https://user-images.githubusercontent.com/69823896/156737015-b8d87156-ddbf-4c96-893b-69016b63596d.png">

Once the score is reached, the game will return a message either of Congratulations or Failure compared to the dealer's hand, as well as the total tally score across all games

package main

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
)

func main() {
	for {
		displayCheckerboard()
		filledPlayerBox(getUserInput())

		if checkVictory() {
			displayCheckerboard()

			logrus.WithField("player-name", getPlayerName()).Infoln("vous avez gagné la partie !")

			return
		} else if checkDraw() {
			logrus.Infoln("Match nul !")

			return
		}

		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}
	}
}

func displayCheckerboard() {
	for i, v := range gameArea {
		fmt.Printf(" %s ", v)

		// each 3 elem, new line
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

func getUserInput() int {
	var (
		ok        bool
		boxNumber = 0
		err       error
	)

	for !ok {
		logrus.Infof("%s: Veuillez entrer un nombre entre 1 et %d\n", getPlayerName(), area)

		var userInput string
		if _, err := fmt.Scanf("%v\n", &userInput); err != nil {
			logrus.Infof("une erreur est apparue lors de la saisie: %v\n", err)
		}

		boxNumber, err = strconv.Atoi(userInput)
		if err != nil {
			logrus.Infof("une erreur est apparue lors de la saisie: %v\n", err)
		}

		if boxNumber < 1 || boxNumber > area {
			logrus.Infof("Merci de rentrer un chiffre entre 1 et %d\n", area)
		} else if gameArea[boxNumber-1] == player1 || gameArea[boxNumber-1] == player2 {
			logrus.Infoln("Case déjà prise !")
		}

		ok = true
	}

	return boxNumber - 1
}

func getPlayerName() string {
	if currentPlayer == player1 {
		return "Joueur 1 "
	}

	return "Joueur 2 "
}

func filledPlayerBox(b int) {
	if currentPlayer == player1 {
		gameArea[b] = player1
		return
	}

	gameArea[b] = player2
}

func checkVictory() bool {
	// on crée un damier pour visualiser les combinaison de la partie actuelle
	var currentCheckerboard [area]bool

	for index, v := range gameArea {
		if currentPlayer == player1 && v == player1 {
			currentCheckerboard[index] = true
		} else if currentPlayer == player2 && v == player2 {
			currentCheckerboard[index] = true
		}
	}

	comparaison := 0
	for _, victory := range possibleVictories {
		for i, v := range currentCheckerboard {
			// si c'est à true dans le même index de tableauPartieEnCours et tableauVictoire alors on incrémente la comparaison
			if v && v == victory[i] {
				comparaison++
				if comparaison == 3 {
					return true
				}
			}
		}

		comparaison = 0
	}

	return false
}

func checkDraw() bool {
	compteur := 0
	for _, v := range gameArea {
		// si la case est prise par un joueur
		if v == player1 || v == player2 {
			compteur++
		}
	}

	return (compteur == len(gameArea))
}

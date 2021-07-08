package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// On défini les constantes globales :
const (
	// taille de l'aire de jeu
	area = 9
	// symbole des joueurs
	symboleJ1 = "X"
	symboleJ2 = "O"
)

// On défini les variables globales
var (
	// aire de jeu
	tableauPlateauMorpion = [area]string{
		"1", "2", "3",
		"4", "5", "6",
		"7", "8", "9"}
	// le joueur 1 commence
	joueur1 = true
)

// fonction proncipale au lancement
func main() {
	play()
}

// fonction qui lance et gère la partie
func play() {
	var numeroDeLaCase int
	for true {
		affichageDuDamier()                         // On affiche le damier
		numeroDeLaCase = inputUtilisateurVerifiee() // on récup l'input entré pas l'utilisateur
		remplirCaseJouee(numeroDeLaCase)            // on rempli la case jouee par le symbole selon le joueur

		if partieGagnee() { // après le coup on vérifie si le joueur à gagner
			affichageDuDamier()
			fmt.Println(nomDuJoueur(), "vous avez gagné la partie !")
			os.Exit(0) // on quitte
		} else if partieExaequo() { // si pas gagné on vérifie si match nul
			fmt.Println("Match nul !")
			os.Exit(0) // on quitte
		}
		joueur1 = !joueur1 // changement de joueur si partie contiue
	}

}

// fonction qui affiche le damier selon cases remplies ou non
func affichageDuDamier() {
	for i := 0; i < len(tableauPlateauMorpion); i++ {
		fmt.Print(" ", tableauPlateauMorpion[i], " ")
		if (i+1)%3 == 0 { // tous les 3 éléments on revient à la ligne
			fmt.Println()
		}
	}
}

// fonction qui vérifie l'input entrée par l'utilisateur
func inputUtilisateurVerifiee() int {
	var (
		valeurValide   = false // booleen qui permet de savoir quand l'input sera valide
		numeroDeLaCase = 0
		err            error
		scanner        = bufio.NewScanner(os.Stdin) //permet de lire une input entrée par l'utilisateur
	)

	for valeurValide == false { // tant que la valeur n'est pas valide
		fmt.Print(nomDuJoueur(), "Veuillez entrer un nombre entre 1 et ", area, " :")
		scanner.Scan()
		numeroDeLaCase, err = strconv.Atoi(scanner.Text()) // A CHECK
		if err != nil {
			fmt.Println("Merci de rentrer uniquement un chiffre !")
		} else if numeroDeLaCase < 1 || numeroDeLaCase > area {
			fmt.Println("Merci de rentrer un chiffre entre 1 et ", area, " !")
			// on check si la case est deja prise ou pas
		} else if tableauPlateauMorpion[numeroDeLaCase-1] == symboleJ1 || tableauPlateauMorpion[numeroDeLaCase-1] == symboleJ2 {
			fmt.Println("Case déjà prise !")
		} else {
			valeurValide = true
		}
	}
	return numeroDeLaCase - 1
}

// fonction qui détermine le nom du joueur selon a qui c'est de jouer
func nomDuJoueur() string {
	if joueur1 == true {
		return "Joueur 1 "
	} else {
		return "Joueur 2 "
	}
}

// fonction qui remplie la case choisie par le symbole du joueur en cours
func remplirCaseJouee(numeroDeLaCase int) {
	if joueur1 == true {
		tableauPlateauMorpion[numeroDeLaCase] = symboleJ1
	} else {
		tableauPlateauMorpion[numeroDeLaCase] = symboleJ2
	}
}

// fonction qui détermine si on gagne la partie selon tous les schéma de victoire possible
func partieGagnee() bool {

	// tableau double dimension des différentes configurations de victoire
	tableauVictoirePossible := [][area]bool{
		{
			true, true, true,
			false, false, false,
			false, false, false},

		{
			false, false, true,
			false, false, true,
			false, false, true},
		{
			false, false, false,
			false, false, false,
			true, true, true},
		{
			true, false, false,
			true, false, false,
			true, false, false},
		{
			true, false, false,
			false, true, false,
			false, false, true},
		{
			false, false, true,
			false, true, false,
			true, false, false},
		{
			false, true, false,
			false, true, false,
			false, true, false},
		{
			false, false, false,
			true, true, true,
			false, false, false}}

	// on crée un damier pour visualiser les combinaison de la partie actuelle
	var tableauPartieEnCours [area]bool

	for index, valeur := range tableauPlateauMorpion {
		if joueur1 && valeur == symboleJ1 { // si c'est le J1 qui joue et que la case a son symbole
			tableauPartieEnCours[index] = true
		} else if !joueur1 && valeur == symboleJ2 { // et inversement
			tableauPartieEnCours[index] = true
		}
	}

	comparaison := 0 // on compte le nombre de true dans les meme cases entre tableauPartieEnCours et TableauVictoirePossible
	for _, tableauVictoire := range tableauVictoirePossible {
		for i := 0; i < len(tableauPartieEnCours); i++ {
			if tableauPartieEnCours[i] == true && tableauPartieEnCours[i] == tableauVictoire[i] { // si c'est à true dans le même index de tableauPartieEnCours et tableauVictoire alors on incrémente la comparaison
				comparaison++
				if comparaison == 3 { // si comparaison true alors le joueur a gagné
					return true
				}
			}
		}
		comparaison = 0 // comparaison remis a 0 pour vérifier une autre combinaison
	}
	return false
}

// fonction qui vérifie si match nul
func partieExaequo() bool {
	compteur := 0

	for _, valeur := range tableauPlateauMorpion {
		if valeur == symboleJ1 || valeur == symboleJ2 { // si la case est prise par un joueur
			compteur++ // on incrémente le compteur
		}
	}

	return (compteur == len(tableauPlateauMorpion))
}

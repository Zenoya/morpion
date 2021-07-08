package main

const (
	area    = 9
	player1 = "X"
	player2 = "O"
)

var (
	gameArea = [area]string{
		"1", "2", "3",
		"4", "5", "6",
		"7", "8", "9"}
	currentPlayer     = player1
	possibleVictories = [][area]bool{
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
)

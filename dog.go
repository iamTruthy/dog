package dog

// Var bingo sets the age of dog (bingo) to a constant which is an int 7
var bingo int = 7

// Years takes in a parameter of type int and returns an int
// Years returns an int multipied by the variable bingo which is set to 7
func Years(a int) int {

	return (a * bingo)
}
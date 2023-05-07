package techpalace

import "strings"

const welcome string = "Welcome to the Tech Palace, "

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return welcome + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return repeatStar(numStarsPerLine) + "\n" + welcomeMsg + "\n" + repeatStar(numStarsPerLine)
}

func repeatStar(numStars int) string {
	line := ""
	for i := 0; i < numStars; i++ {
		line = line + "*"
	}
	return line
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimRight(strings.TrimLeft(strings.ReplaceAll(strings.ReplaceAll(oldMsg, "*", ""), "\n", ""), " "), " ")
}

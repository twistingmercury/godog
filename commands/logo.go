package commands

import "fmt"

func Logo() {

	logo := " _____           _               _\n" +
		"|  __ \\         | |             | |\n" +
		"| |  \\/ ___   __| | ___   __ _  | |\n" +
		"| | __ / _ \\ / _` |/ _ \\ / _` | | |\n" +
		"| |_\\ \\ (_) | (_| | (_) | (_| | | |\n" +
		"\\____/ \\___/ \\__,_|\\___/ \\__, | |_|\n" +
		"                          __/ /  _ \n" +
		"                         |___/  |_| \n\n" +
		"                Who's a good boy!?!"
	fmt.Println(logo)
}

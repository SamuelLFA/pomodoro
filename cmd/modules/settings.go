package modules

import "fmt"

func GetUserSettings(minutesToWork int, minutesToRest int) (int, int) {
	var input int = 0

	fmt.Print("\nEnter the number of minutes to work: ")
	for input <= 0 {
		fmt.Scanf("%d\r", &input)
	}
	minutesToWork = input
	input = 0

	fmt.Print("\nEnter the number of minutes to rest: ")
	for input <= 0 {
		fmt.Scanf("%d\r", &input)
	}
	minutesToRest = input
	input = 0
	return minutesToWork, minutesToRest
}

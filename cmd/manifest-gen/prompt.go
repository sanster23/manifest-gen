package main

// import (
// 	"fmt"

// 	"github.com/manifoldco/promptui"
// )

// func main() {
// 	promptDay := promptui.Select{
// 		Label: "Select Day",
// 		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
// 			"Saturday", "Sunday"},
// 	}

// 	_, day, err := promptDay.Run()

// 	if err != nil {
// 		fmt.Printf("Prompt day failed %v\n", err)
// 		return
// 	}
// 	promptMonth := promptui.Select{
// 		Label: "Select Month",
// 		Items: []string{"Jan", "Feb", "Mar"},
// 	}

// 	_, month, err := promptMonth.Run()

// 	if err != nil {
// 		fmt.Printf("Prompt month failed %v\n", err)
// 		return
// 	}
// 	fmt.Printf("You choose \n \tday = %q \n \t month = %q", day, month)
// }

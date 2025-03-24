// package main

// import (
// 	"fmt"
// 	"os"
// )

// // first assignment of the exercise

// // func getPassedArgs(minArgs int) []string {
// // 	if len(os.Args) < minArgs {
// // 		fmt.Printf("At least %v arguments are needed\n", minArgs)
// // 		os.Exit(1)
// // 	}
// // 	var args []string
// // 	for i := 1; i < len(os.Args); i++ {
// // 		args = append(args, os.Args[i])
// // 	}

// // 	return args
// // }

// // Second assignement of the exercise

// func getPassedArgs() []string {
// 	var args []string
// 	for i := 1; i < len(os.Args); i++ {
// 		args = append(args, os.Args[i])
// 	}

// 	return args
// }

// // func findLongestString(args []string) string {
// // 	var longest = args[0]
// // 	for i := 0; i < len(args); i++ {
// // 		if len(longest) < len(args[i]) {
// // 			longest = args[i]
// // 		}
// // 	}
// // 	return longest
// // }

// func addLocales(extraLocals []string) []string {
// 	var locals []string

// 	locals = append(locals, "en_US", "fr_FR")
// 	locals = append(locals, extraLocals...)

// 	return locals
// }

// func main() {

// 	// Main: first assignment
// 	// if longest := findLongestString(getPassedArgs(3)); len(longest) > 0 {
// 	// 	fmt.Printf("the longs word passed is %v\n", longest)
// 	// } else {
// 	// 	fmt.Println("There was an error")
// 	// 	os.Exit(1)
// 	// }

// 	locals := addLocales(getPassedArgs())
// 	fmt.Printf("the arguments that are used in the command: %v\n", locals)
// }

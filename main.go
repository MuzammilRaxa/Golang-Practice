package main

import "fmt"

func main() {
	var shortGolangCourse = "Watch Go crash course"
	var fullGolangCourse = "Watch Go full course"
	var listItems = []string{shortGolangCourse, fullGolangCourse}
	fmt.Println("List of my todos:")

	for index, task := range listItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}

}

package main

import "fmt"

var arr = [10]int{1,2,3,14,15,6,7,8,19,0}

func main() {
	display();
	printLn(13);

	int location = find(15);

	if location != -1 {
		fmt.Printf("Location was found at %d", location+1);
	} else {
		fmt.Printf("Location not found");
	}
}

func printLn(count int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("=");
	}
	fmt.Println("");
}

func find(data int) {
	var comparisons int = 0;
	var index int = -1;
	for i := 0; i < 10; i++ {
		comparisons++;

		if data == arr[i] {
			index = i;
			break;
		}
	}
	fmt.Printf("Total comparisons are %d", comparisons);
	return index;
}

func display() {
	for i := 0; i < 10; i++ {
		fmt.Println("%d", arr[i]);
	}
}

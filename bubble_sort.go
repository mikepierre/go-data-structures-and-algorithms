package main

import "fmt"

var list = [10]int{10,1,2,3,14,5,6,17,8,9}

func main(){
fmt.Println("Bubble Sort In Go");
display();
fmt.Println("");
bubbleSort();
fmt.Println("");
display();
}

func display(){
	for i :=0; i < 10; i++ {
		fmt.Printf("value: %d\n", i);
	}
}

func bubbleSort(){
	var swapped bool = false;

	for i := 0; i < 10-1; i++{
		swapped = false;
		for j := 0; j < 10-1-i; j++ {
			fmt.Printf(" Items compared: [%d, %d] \n", list[j],list[j+1]);

			if list[j] > list[j+1] {
				temp := list[j];
				list[j] = list[j+1];
				list[j+1] = temp;

				swapped = true;
				fmt.Printf(" swapped [%d, %d]\n", list[j],list[j+1]);
			} else {
				fmt.Printf("Not swapped\n");
			}
		}
		if !swapped {
			break;
		}
		fmt.Println("Iteration %d", i+1);
		display();
	}
}

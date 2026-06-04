package main
import "fmt"

func main(){

	// iteration using for loop
	for i :=1; i<=5; i++{
		fmt.Println("Itrateion number:", i)
	}

	// sum of first n natural numbers
	var sum, x int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&x)
	for n:=1; n<= x; n++{
		sum +=n
	}
	fmt.Println("The sum of first %d natural numbers is %d",x ,sum)

	// infinite loop
	// for {
	// 	fmt.Println("This is an infinite loop. Press Ctrl+C to stop it.")
	// }

	//  loop with break statement
	for i :=1; i<= 10; i++ {
		fmt.Println("Iteration number:", i)
		if i==5{
			fmt.Println("\n\n\n\nBreaking the loop at iteration", i)
			break
		}
	}

	// loop with continue statement
	for i:=1; i<=10; i++{
		if i%2==0 {
			// fmt.Println("\n\n\nSkipping even number:", i)
			continue
		}
		fmt.Println("Iteration number:", i)
		
	// nested loops calculating multiplication table
	fmt.Print("enter two numbers: ")
	var num1, num2 int
	fmt.Scanln(&num1, &num2)

	for i:=1; i<=num1; i++{
		for j:=1; j<=num2; j++{
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// print numbers vertically

	for i:=1; i<=10; i++{
		fmt.Println(i)
	}

	// print even numbers from 1 to 20
	for i:=1; i<=20; i++{
		if i%2==0 {
			fmt.Println(i)
		}
	}

	// print numbers lable with 5x i
	num5:= 5
	for i:=1; i<=10; i++{
		fmt.Printf("%d x %d = %d\n",num5, i, num5*i)
	}

	// calculate the factorial of a number
	var n, factorial int 
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)
	factorial = 1
	for i:=1; i<=n; i++{
	factorial *= i
	}
	fmt.Printf("The factorial of %d is %d\n", n, factorial)	

	// Student Score Analyzer
	var n int

	fmt.Print("Enter the number of students: ")
	fmt.Scanln(&n)
	for i:=1; i<=n; i++ {
		var std string
		fmt.Print("Enter student name: ")
		fmt.Scanln(&std)

		var score1, score2, score3 float64
		fmt.Print("Enter score for subject 1: ")
		fmt.Scanln(&score1)
		fmt.Print("Enter score for subject 2: ")
		fmt.Scanln(&score2)
		fmt.Print("Enter score for subject 3: ")
		fmt.Scanln(&score3)

		var highest, lowest float64
		totalScore := score1 + score2 + score3
		avg := totalScore / 3

		if score1 > score2 {
			if score1 > score3 {
				highest = score1
			} else {
				highest = score3
			}
		} else {
			if score2 > score3 {
				highest = score2
			} else {
				highest = score3
			}
		}
		if score1 < score2 {
			if score1 < score3 {
				lowest = score1
			} else {
				lowest = score3
			}
		} else {
			if score2 < score3 {
				lowest = score2
			} else {
				lowest = score3
			}
		}
		fmt.Printf("Student: %s\n", std)
		fmt.Printf("Total Score: %.2f\n", totalScore)
		fmt.Printf("Average Score: %.2f\n", avg)
		fmt.Printf("Highest Score: %.2f\n", highest)
		fmt.Printf("Lowest Score: %.2f\n", lowest)
	}
}
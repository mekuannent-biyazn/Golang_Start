package main
import "fmt"

func main() {
	numbers := []int {}
	numbers = append(numbers, 0,1,2,3,4,5,6,7,8,9,10)

	number1 :=make([]int ,len(numbers) ,cap(numbers)*2)
	copy(number1 ,numbers)

	fmt.Println("the copy of numbers are: ",number1)
	fmt.Println("the original number is: ",numbers)

	// nil slice
	var num []int
	fmt.Println(num)
	if (num == nil){
		fmt.Println("Number is nil!")
		}	

		
	numbers := []int{0,1,2,3,4,5,6,7,8,9}

	for i, number := range numbers{
		fmt.Printf("Slice item %d is %d\n",i,number)
	}

	n :=[]int{1,2,3}
	for i:= range n{
		fmt.Printf("Array item %d is value of: %d\n",i, n[i])
	}

	fmt.Println("Maps!")

	// var myMap map[string]string
	// myMap = make(map[string]string)
	// fmt.Println("My map is: ",myMap) 

	countryCapitalMap := make(map[string]string)

	countryCapitalMap["Ethiopia"]="Addis abeba"
	countryCapitalMap["Erthra"]="semra"
	countryCapitalMap["Dijjibuti"]="djjibuti"

	for capital :=range countryCapitalMap{
		fmt.Println("the capital of ",capital," is ",countryCapitalMap[capital])
	}
	countryCap, ok := countryCapitalMap["Ethiopia"]
	if (ok){
		fmt.Println("The capital of: ",countryCapitalMap[countryCap]," is", countryCap)
	}else{
		fmt.Println("Does not exist!")
	}

	// delete entry from the map!
	delete(countryCapitalMap,"Dijjibuti")
	fmt.Println("the entry is deleted success fully!")

	for countryd := range countryCapitalMap{
		fmt.Println("the updated country capital of ",countryd," is ",countryCapitalMap[countryd])
	}
}
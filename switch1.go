package main
import "fmt"


func getprice(){
	fmt.Println("Gotv packages!!\n")
	fmt.Print("choose a package Basic, Student,Premium,Throne:\n")
	var subscription string
	fmt.Scan(&subscription)



	switch subscription {
	case "basic":
		fmt.Println("1000")
	case "student":
		fmt.Println("2000")
    case "premium":
		fmt.Println("3000")
	case "throne":
		fmt.Println("4000")
	}

}

func main(){

	getprice()
}
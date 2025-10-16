package main 
import "fmt"

/*unc pickscore(){

	,Biology,Physics,Chemistry,Geography float64

	
	fmt.Print("Enter your score in biology: \n")
	fmt.Scan(&Biology)
	fmt.Print("Enter your score in physics: \n")
	fmt.Scan(&Physics)
	fmt.Print("Enter your score in chemistry: \n")
	fmt.Scan(&Chemistry)
	fmt.Print("Enter your score in Geography: \n")
	fmt.Scan(&Geography)
}
	*/

func mathgrade(){
	var Mathematics float64
	fmt.Print("Enter your score in mathematics: \n")
	fmt.Scan(&Mathematics)

	if Mathematics>=70 && Mathematics <=100 {
		fmt.Printf("Mathematics %v Grade A\n",Mathematics)
	} else if Mathematics >=60 && Mathematics<=69 {
		fmt.Printf("Mathematics %v Grade B\n",Mathematics)
	} else if Mathematics >=50 && Mathematics<=59{
		fmt.Printf("Mathematics %v Grade C\n",Mathematics)
	} else if Mathematics >= 40 && Mathematics <=49 {
		fmt.Printf("Mathematics %v Grade D\n",Mathematics)
	}else if Mathematics <40 && Mathematics >0 {
		fmt.Printf("Mathematics %v Grade E\n",Mathematics)
	} else {
		fmt.Println("Enter marks between 1 and 100\n")
	}

	var Biology float64
	fmt.Print("Enter your score in Biology: \n")
	fmt.Scan(&Biology)

	if Biology>=70 && Biology <=100 {
		marks:=Biology
		Bgrade:="A"
		
	} else if Biology >=60 && Biology<=69 {
		fmt.Printf("Biology %v Grade B\n",Biology)
	} else if Biology >=50 && Biology<=59{
		fmt.Printf("Biology %v Grade C\n",Biology)
	} else if Biology >= 40 && Biology <=49 {
		fmt.Printf("Biology %v Grade D\n",Biology)
	}else if Biology <40 && Biology >0 {
		fmt.Printf("Biology %v Grade E\n",Biology)
	} else {
		fmt.Println("Enter marks between 1 and 100\n")
	}

	fmt.Printf("marks: %v grade: %v",marks,Bgrade)
}
			




func main (){
	fmt.Println("welcome to student grade manager\n")
	fmt.Println("Grades\n")
	mathgrade()

}
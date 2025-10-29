package main
import "fmt"

func getgrade(midterm float64,endterm float64) (float64,string){
	averagegrade:=(midterm+endterm)/2
	var gradeletter string

	if averagegrade>=90{
		gradeletter="A"
	}else if averagegrade>=80{
		gradeletter="B"
	}else if averagegrade >=70{
		gradeletter="C"
	}else if averagegrade >=60{
		gradeletter="D"
	}else {
		gradeletter="Fail"
	}

	return averagegrade,gradeletter

}

func main (){
	var midtermmarks,endtermmarks float64
	fmt.Print("Enter midterm marks : \n")
	fmt.Scan(&midtermmarks)

	fmt.Print("Enter endterm marks : \n")
	fmt.Scan(&endtermmarks)

	averagegrade,gradeletter:=getgrade(midtermmarks,endtermmarks)

	fmt.Printf("Avaragegrade: %v\nGrade:%v\n",averagegrade,gradeletter)

}
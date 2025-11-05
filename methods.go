package main 
import "fmt"

type student struct{
	name string
	age int
	city string
	grades []int
}

func(s student) getavarege()float64{
	sum:=0

	for _ ,v:= range s.grades{
		sum+=v
	}
	mean:=sum/len(s.grades)

	return float64(mean)
}

func (s student) greetings(){
	fmt.Println("hello, ",s.name)
}

func (m student) getmax()int{
	max:=0
	for _,v:= range m.grades{
		if max<v{
			max=v
		}
	}
	 return max
}

func main () {
	s1:=student{"kimani",22,"nairobi",[]int{5,66,77}}
	//avarage:=s1.getavarege()

	s1.greetings()
	max1:=s1.getmax()
	av1:=s1.getavarege()
	fmt.Printf("max grade:%v average:%v",max1,av1)

}




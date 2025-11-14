package main
import ("fmt")

type userinfo struct{
	fname string
	lname string
	email string
	numtickets int
}

type userslist []userinfo

func (pointer *userslist) add(fname,lname,email string, numtickets int){
	newuser =: userinfo{
		fname:fname,
		lname:lname,
		email:email
		numtickets:	numtickets,
	}
	*pointer:=append(*pointer,newuser)
}

func main (){
 var person1 userslist

 
}


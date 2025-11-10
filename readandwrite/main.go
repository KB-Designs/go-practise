package main

import (
	"fmt"
	"os"
)

func writefile(){
	newfile,err:=os.Create("example.txt")
	if err !=nil{
		fmt.Println("error creating file")
		return
	}
	
	defer newfile.Close()

	_,err = newfile.WriteString("hello from go world.")
	if err != nil{
		fmt.Println ("Error in creating the file")
		return
	}
	
}                          

func readfile(){
	data,err := os.ReadFile("example.txt")
	if err != nil{
		fmt.Println("Error in reading file.")
		return
	}	
	fmt.Println(string(data))
}

func main (){
	writefile()
	readfile()
}

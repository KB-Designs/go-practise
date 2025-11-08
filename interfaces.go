package main
import "fmt"

type shape interface{
	peri()float64
	area()float64
}

type cirlce struct{
	radius float64
}

type rectangle struct{
	length float64
	width float64
}

func(c cirlce) peri()float64{
	return 3.14*c.radius*2
}

func (c cirlce) area()float64{
	return 3.14*c.radius*c.radius
}

func (r rectangle) peri()float64{
	return (r.length+r.width)*2
}

func (r rectangle) area()float64{
	return r.length*r.width
}

func main(){
	shapes:=[]shape{
		cirlce{radius:21},
		rectangle{length:12, width:10},
	}

	for _,v:= range shapes{
		
		fmt.Println("area",v.area())
		fmt.Println("perimeter", v.peri())
		fmt.Printf("...rectangle\n")
	}

}

/*
Your code structure is great — you just need to:

Fix small typos,

Fill the interface slice, and

Loop through it to print each shape’s perimeter and area.

Explanation

shape is an interface requiring two methods:

peri() float64
area() float64


Both circle and rectangle implement these methods, so they automatically satisfy shape.

In main(), we create a slice of type []shape — it can hold any type that implements the shape interface.

The for loop calls the methods polymorphically — Go decides which version (circle’s or rectangle’s) to call at runtime.
*/


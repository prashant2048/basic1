package main
import ("fmt" )
type st struct {
x int 
y int
}
func main(){
	var z int
	fmt.Scanln(&z)
	s:=st{1,2}
	fmt.Println(s)
	fmt.Println(z)
	
	fmt.Println("changed")
}

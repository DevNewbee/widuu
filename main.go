package main
import "fmt"
type User struct{
	Member string
	Age int
}
func main(){
	a:=User{Member:"widuu",Age:12}
	a.test()
	fmt.Println(a)
}
func (u *User) test(){
	fmt.Println(u.Age)
}

package main
import ("fmt" 
"net"
"time"
)
func main() {	
	t:=time.Now()
	fmt.Println(t.Location())
   	ipr,_:=net.LookupIp("facebook.com")
	for _, ip:=range ipr {
		fmt.Println(ip)
	fmt.Println("done..\n")
	}
}

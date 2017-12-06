package main

import ("fmt"
	"net"
	)
func main() {
	packet := []byte{27,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	//ip:=(ipResolver("time.nist.gov"))
	ip:= "129.6.15.28:123"
	fmt.Println(ip)
	conn,em1 := net.Dial("udp",ip)
	errorCheck(em1)
	fmt.Println("Writing Packet")
	bOut, _ := conn.Write(packet)
	response:= make([]byte,64)
	fmt.Println(bOut, "bytes sent")
	fmt.Println("Reading Response")
	conn.Read(response)
	fmt.Println(response)
}

func errorCheck(e error){
	if e != nil{
		panic(e)
	}
}
func ipResolver(name string) []net.IP{
	ip, e1 := net.LookupIP(name)
	errorCheck(e1)
	return ip
}
func packetParser(ntpResponse []byte){

}
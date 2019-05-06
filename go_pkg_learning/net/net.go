package net

import (
	"fmt"
	"net"
	"os"
	"time"
)

func parseIp()  {
	name := "172.19.120.142"

	ip := net.ParseIP(name)

	mask := ip.DefaultMask()

	network := ip.Mask(mask)

	fmt.Fprintf(os.Stdout, "network: %s\n", network.String()) // 192.168.1.0

	// ip: 		192.168.1.97
	// mask:	255.255.255.0
	// network:	192.168.1.0


}

/**/
/*根据域名获取IP*/
func parseIpAddress()  {

	domain := "www.google.com"
	ipAddr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "%s IP: %s Network: %s Zone: %s\n", ipAddr.String(), ipAddr.IP, ipAddr.Network(), ipAddr.Zone)

	// 百度，虽然只有一个域名，但实际上，他对应电信，网通，联通等又有多个IP地址
	ns, err := net.LookupHost(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	for _, n := range ns {
		fmt.Fprintf(os.Stdout, "%s\n", n) // 115.239.210.26    115.239.210.27 这2个地址打开都是百度
	}

}
/*返回该域名的string类型所有地址*/
func lookUpHost()  {
	address,err:=net.LookupHost("www.google.com")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(address)

}
/*返回该域名的IP类型所有地址*/
func lookUpIp()  {
	address,err:=net.LookupIP("www.google.com")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(address)

}

func lookUpPort() {
	port,err := net.LookupPort("tcp","http")
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(port)

}

/*网络类型可以是以下几种

- "tcp"
- "tcp4" (仅限IPv4)
- "tcp6" (仅限IPv6)
- "udp"
- "udp4" (仅限IPv4)
- "udp6" (仅限IPv6)
- "ip"
- "ip4" (仅限IPv4)
- "ip6" (仅限IPv6)
- "unix"
- "unixpacket"*/
func dail()  {
	//conn,err:=net.Dial("tcp","www.163.com:80")
	conn,err:=net.Dial("tcp","www.google.com:80")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("conn=",conn.LocalAddr())
}
func dailIp()  {
	raddr, _ := net.ResolveIPAddr("tcp", "www.google.com")
	ipconn, err := net.DialIP("ip:1", nil, raddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ipconn:%#v", ipconn) //ipconn:&net.IPConn{fd:(*net.netFD)(0xf840069000)}


}
func dailTimeOut()  {
	timeout:=time.Duration(5*time.Second)
	conn,err:=net.DialTimeout("tcp","www.google.com:80" ,
		timeout)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("conn=",conn)

}
func interfaceAddrs()  {
	address,err:=net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("address=",address)


}
func listenIp() {
	ipaddr,_ := net.ResolveIPAddr("tcp","www.baidu.com")
	ipconn,err := net.ListenIP("ip:1",ipaddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ipconn:%#v",ipconn) //返回 ipconn:&net.IPConn{fd:(*net.netFD)(0xf840069000)}

}
func main() {

	hw,err:=net.ParseMAC("38:f9:d3:2e:28:4f")
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("HW=",hw)

	ip:=net.ParseIP("172.19.120.142")
	fmt.Println("ip=",ip)
	ip = net.ParseIP("2404:6800:4005:80d::2004")
	fmt.Println("ip=",ip)
	parseIp()
	parseIpAddress()
	lookUpHost()
	lookUpIp()
	lookUpPort()
	dail()
	dailIp()
	dailTimeOut()
	interfaceAddrs()
	listenIp()

}



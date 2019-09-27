//  go run nmap.go --port 9100

package main

import (
	"flag"
	"math"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var printerList []string = make([]string, 0)

func main() {
	networks := GetHostNet()

	port := flag.Int("port", 22, "Port number to be scanned")
	flag.Parse()

	for _, network := range networks {
		ScanNetwork(network, *port)
	}

	for _, host := range printerList {
		println(strings.Repeat("=", 50))
		println("ONLINE ", host)
	}
}

func GetHostNetByShell() []string {
	cmd := exec.Command("ip", "addr")
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Fprintf(os.Stdout, "Result: %s", buf)
	r, _ := regexp.Compile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.\d{1,3}/\d{1,2}`)
	//fmt.Println(r.FindAllString(string(buf),-1))

	temp := r.FindAllString(string(buf), -1)
	networks := make([]string, 0)
	for _, net := range temp {
		networks = append(networks, net)
	}
	return networks
}

func GetHostNet() []string {
	ips := []string{"127.0.0.1/24"}

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.String())
			}
		}
	}
	return ips
}

func ip2long(ip net.IP) uint {
	number := 0
	for i := 0; i < len(ip); i++ {
		number = number << 8
		number += int(ip[i])
	}
	return uint(number)
}
func long2ip(ip uint) net.IP {
	parts := make([]byte, net.IPv4len)
	for i := 0; ip > 0; i++ {
		b := ip & 0xFF
		parts[i] = byte(b)
		ip = ip >> 8
	}
	out := net.IPv4(parts[3], parts[2], parts[1], parts[0])
	return out
}

// 扫描
func ScanNetwork(network string, port int) {
	ip, inet, err := net.ParseCIDR(network)
	if err != nil {
		return
	}

	waitGroup := sync.WaitGroup{}

	mutex := sync.Mutex{}

	println(ip.Mask(inet.Mask).String())
	lowerBound := ip2long(ip.Mask(inet.Mask))

	ones, bits := inet.Mask.Size()
	space := int(math.Exp2(float64(bits-ones))) - 2 // Minus all 0s  and all 1s address
    if inet.IP.IsLoopback(){
        space = 1
    }

	for index := 0; index < space; index++ {
		ip := long2ip(lowerBound + uint(index))
		fmt.Printf("scanning %s \r", ip)
		fmt.Printf(strings.Repeat(" ", 50) + "\r")

		waitGroup.Add(1)

		go func() {
			if IsPortOpen(ip, port, time.Second*1) {
				mutex.Lock()
				printerList = append(printerList, ip.String())
				mutex.Unlock()
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

// 检查端口是否开放
func IsPortOpen(ip net.IP, port int, timeout time.Duration) bool {
	remoteAddr := ip.String() + ":" + strconv.Itoa(port)
	conn, _ := net.DialTimeout("tcp", remoteAddr, timeout)

	if conn == nil {
		return false
	}

	defer conn.Close()

	return true
}

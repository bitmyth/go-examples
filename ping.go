package main

import (
    "golang.org/x/net/icmp"
    "golang.org/x/net/ipv4"
    "net"
    "os"
    "time"
)

func SimplePing(ip net.IP) error {
    c, err := net.DialIP("ip4:icmp", nil, &net.IPAddr{IP: ip})
	if err != nil {
		println(err.Error())
		return err
	}
    defer c.Close()
    if err != nil {
        return err
    }

    m := icmp.Message{
        Type: ipv4.ICMPTypeEcho, Code: 0,
        Body: &icmp.Echo{
            ID: os.Getpid() & 0xffff, Seq: 1,
            Data: []byte(""),
        },
    }

    b, err := m.Marshal(nil)
    if err != nil {
        return err
    }
    if _, err := c.Write(b); err != nil {
        return err
    }

    reply := make([]byte, 1500)

    if err = c.SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil {
        return err
    }
    _, _, err = c.ReadFrom(reply)
    if err != nil {
        return err
    }
    return nil

}

func main(){
	err:=SimplePing(net.ParseIP("192.1.1.1"))
	if err != nil {
		println(err.Error())
	}else{
		println("reachable")
	}
}

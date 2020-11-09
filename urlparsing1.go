package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	p := fmt.Println
	s := "Mysql://admin:password@serverhost.com:8081/server/page1?key=value&key2=value2#X"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	p(u.Scheme)
	p(u.User)
	p(u.User.Username())
	pass, _ := u.User.Password()
	p(pass)
	p(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	p(host)
	p(port)
	p(u.Path)
	p(u.Fragment)
	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["key2"][0])
}

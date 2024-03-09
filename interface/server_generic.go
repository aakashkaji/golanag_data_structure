package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	Id            string
	Port          int
	Ip            string
	MaxConnection int
	Tls           bool
}

type Server struct {
	Ops Opts
}

func withTls(otps *Opts) {
	otps.Tls = true
}

func witPort(s int) OptFunc {
	return func(opts *Opts) {
		opts.Port = s
	}
}

func withMaxConnection(s int) OptFunc {
	return func(o *Opts) {
		o.MaxConnection = s
	}
}

func setDefault() Opts {
	return Opts{Id: "1234", Port: 80, Ip: "192.168.20.203", MaxConnection: 10, Tls: false}

}

func GetServer(otpfun ...OptFunc) *Server {

	o := setDefault()

	for _, fns := range otpfun {
		//
		fns(&o)
	}

	return &Server{
		Ops: o,
	}
}

func main() {

	fmt.Println(GetServer(withTls, witPort(433), withMaxConnection(99)))

}

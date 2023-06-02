package main

type LogicProvider struct {

}

func (lp LogicProvider) Process(data string) string {
	return ""
}

type Login interface {
	Process(data string) string
}

type Client struct {
	L Login
}

func (c Client) Program() {
	c.L.Process("data")
}

func demo() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
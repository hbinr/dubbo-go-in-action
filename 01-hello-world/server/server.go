package server

import (
	"fmt"

	"github.com/apache/dubbo-go/common/logger"
)

type PeopleReq struct {
	ID int64 `json:"id"`
}

type PeopleRes struct {
	Name string `json:"name"`
}

type HelloWorldProvider struct{}

func (h *HelloWorldProvider) SayHi(msg string) (string, error) {
	logger.Info("server: SayHi method,arg is ", msg)
	return fmt.Sprintf("Hello %s", msg), nil
}

func (h *HelloWorldProvider) Who(req *PeopleReq) (res *PeopleRes, err error) {
	logger.Info("server: Who method,arg is ", req)
	res = new(PeopleRes)
	res.Name = "Bob"
	if req.ID == 1 {
		res.Name = "Tom"
	}
	return
}

func (u *HelloWorldProvider) Reference() string {
	return "HelloWorldProvider"
}

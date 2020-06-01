package main

import (
	"github.com/noaway/tutorial/internal/svc"
	"github.com/noaway/tutorial/proc"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := svc.Run(new(proc.Proc), nil); err != nil {
		logrus.Error(err)
	}
}

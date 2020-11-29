package gtime

import "github.com/gogf/gf/os/gcron"

var (
	TimeTask = gcron.New()
)

func Start()  {
	TimeTask.Add("@every 10m",synctags,"synctags")
}


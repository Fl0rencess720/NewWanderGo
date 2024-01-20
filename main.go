package main

import (
	conf "WanderGo/configs"
	rt "WanderGo/router"
)

func main() {
	conf.InitLogging()
	conf.ConnectToDb()
	rt.Start()
}

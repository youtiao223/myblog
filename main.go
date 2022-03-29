package main

import (
	"myBlog/config"
	"myBlog/model"
	"myBlog/routers"
)

func init() {
	config.Init()
	model.Init()
	routers.Init()
}
func main() {

}

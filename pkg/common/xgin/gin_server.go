package xgin

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	Engine *gin.Engine
}

func NewGinServer() *GinServer {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	// 1、使用 Recovery 中间件
	engine.Use(gin.Recovery())
	// 2、跨域
	//engine.Use(middleware.Cors())
	return &GinServer{engine}
}

func (s *GinServer) Run(port int) {
	addr := ":" + strconv.Itoa(port)
	err := s.Engine.Run(addr)
	if err != nil {
		fmt.Println("GinServer Start Failed.", err.Error())
	}
}

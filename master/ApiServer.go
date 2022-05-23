package master

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

// ApiServer 任务的HTTP接口
type ApiServer struct {
	httpServer *http.Server
}

var (
	G_apiServer *ApiServer
)

func handleJobSave(w http.ResponseWriter, r *http.Request) {

}

// InitApiServer 初始化服务
func InitApiServer() error {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		err        error
		httpServer *http.Server
	)

	//配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)

	// 启动TCP
	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(G_config.ApiPort)); err != nil {
		return nil
	}

	//创建一个HTTP服务
	httpServer = &http.Server{
		ReadTimeout:  time.Duration(G_config.ApiReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}

	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	go httpServer.Serve(listener)
	return nil
}

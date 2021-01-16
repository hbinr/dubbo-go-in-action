package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/apache/dubbo-go/config"
	"github.com/apache/dubbo-go/protocol/rest/server/server_impl"
	"github.com/emicklei/go-restful/v3"

	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/rest"
	_ "github.com/apache/dubbo-go/registry/nacos"
	_ "github.com/apache/dubbo-go/registry/protocol"

	"study.dubbogo/01-hello-world/server"
)

// Environment variables must be set first,such as:
// 	export CONF_PROVIDER_FILE_PATH="../conf/server.yml"
//	export APP_LOG_CONF_FILE="../conf/log.yml"
func main() {
	// 1.set provider,register service
	config.SetProviderService(new(server.HelloWorldProvider))

	// 2.add RESTfule server
	server_impl.AddGoRestfulServerFilter(func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		chain.ProcessFilter(request, response)
	})

	// 3.init config
	config.Load()

	// 4. init signal (Optional)
	initSignal()
}

var (
	survivalTimeout = int(3e9)
)

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
		// reload()
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("test exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider test exit now...")
			return
		}
	}
}

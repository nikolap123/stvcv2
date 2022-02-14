package network

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"time"
)

var server *socketio.Server
var liveHosts []string
var allHosts []string

func GetServer() *socketio.Server {

	server = socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	return server
}

func sendHosts() {

	server.BroadcastToNamespace("/", "alive", liveHosts)
}

func updateHosts(H []string) {

	for i, _ := range H {
		go ping(H[i])
	}
}

func ping(ip string) {

}

func Start() {
	allHosts = []string{
		"192.168.50.21",
		"192.168.50.106",
		"192.168.50.199",
	}

	quit := make(chan struct{})

	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				updateHosts(allHosts)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ticker.C:
				sendHosts()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

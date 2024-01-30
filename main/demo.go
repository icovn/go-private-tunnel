package main

import (
	"container/list"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"time"

	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/log"
	"github.com/Allenxuxu/toolkit/sync/atomic"
)

const clientsKey = "demo_push_message_key"

var connections = list.New()

type example struct {
	Count atomic.Int64
}

func (s *example) OnConnect(c *gev.Connection) {
	s.Count.Add(1)
	e := connections.PushBack(c)
	c.Set(clientsKey, e)
	//log.Println(" OnConnect ： ", c.PeerAddr())
}
func (s *example) OnMessage(c *gev.Connection, ctx interface{}, data []byte) (out interface{}) {
	//log.Println("OnMessage")
	out = data
	return
}

func (s *example) OnClose(c *gev.Connection) {
	s.Count.Add(-1)
	//log.Println("OnClose")
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	message := request.URL.Path[1:]
	var next *list.Element
	for e := connections.Front(); e != nil; e = next {
		next = e.Next()

		c := e.Value.(*gev.Connection)
		_ = c.Send([]byte(message))
	}
	fmt.Fprintf(writer, "Hi there, I love %s!", request.URL.Path[1:])
}

func main() {
	handler := new(example)

	go func() {
		http.HandleFunc("/", rootHandler)
		if err := http.ListenAndServe(":6060", nil); err != nil {
			panic(err)
		}
	}()

	var port int
	var loops int

	flag.IntVar(&port, "port", 1833, "server port")
	flag.IntVar(&loops, "loops", -1, "num loops")
	flag.Parse()

	s, err := gev.NewServer(handler,
		gev.Network("tcp"),
		gev.Address(":"+strconv.Itoa(port)),
		gev.NumLoops(loops),
		gev.MetricsServer("", ":9091"),
	)
	if err != nil {
		panic(err)
	}

	s.RunEvery(time.Second*2, func() {
		log.Info("connections :", handler.Count.Get())
	})

	s.Start()
}

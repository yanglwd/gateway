package gateway

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/viper"
	"github.com/yanglwd/gateway/service/internal/connection"
)

type GatewayService struct {
	network    string
	address    string
	serverId   uint16
	sig        chan os.Signal
	ln         net.Listener
	connSerial uint32
	conns      sync.Map
}

func (g *GatewayService) Init() error {
	network, ok := viper.Get("net").(string)
	if !ok {
		return errors.New("invalid network type")
	}
	addr, ok := viper.Get("addr").(string)
	if !ok {
		return errors.New("invalid listen addr")
	}
	id, ok := viper.Get("id").(uint16)
	if !ok {
		return errors.New("invalid server id")
	}

	ln, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	g.network = network
	g.address = addr
	g.serverId = id
	g.ln = ln

	g.sig = make(chan os.Signal, 8)
	signal.Notify(g.sig, syscall.SIGABRT|syscall.SIGKILL|syscall.SIGINT)
	return nil
}

func (g *GatewayService) Serve() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

loop:
	for {
		select {
		case s := <-g.sig:
			fmt.Println(s.String())
		default:
		}

		conn, err := g.ln.Accept()
		if err != nil {
			break loop
		}

		gc := &connection.GatewayConn{}
		if err := gc.Init(conn); err != nil {
			conn.Close()
		} else {
			g.connSerial++
			g.conns.Store(g.connSerial, gc)
		}
	}

	g.conns.Range(func(_, value interface{}) bool {
		if conn, ok := value.(connection.GatewayConn); ok {
			conn.Exit()
		}
		return true
	})
}

func (g *GatewayService) Exit() {
	g.sig <- os.Kill
}

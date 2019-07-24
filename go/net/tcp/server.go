package tcp

import (
	"context"
	"github.com/searKing/golang/go/sync/atomic"
	time2 "github.com/searKing/golang/go/time"
	"github.com/searKing/golang/go/util/object"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type Handler interface {
	OnOpenHandler
	OnMsgReadHandler
	OnMsgHandleHandler
	OnCloseHandler
	OnErrorHandler
}

func NewServerFunc(
	onOpen OnOpenHandler,
	onMsgRead OnMsgReadHandler,
	onMsgHandle OnMsgHandleHandler,
	onClose OnCloseHandler,
	onError OnErrorHandler) *Server {
	return &Server{
		onOpenHandler:      object.RequireNonNullElse(onOpen, NopOnOpenHandler).(OnOpenHandler),
		onMsgReadHandler:   object.RequireNonNullElse(onMsgRead, NopOnMsgReadHandler).(OnMsgReadHandler),
		onMsgHandleHandler: object.RequireNonNullElse(onMsgHandle, NopOnMsgHandleHandler).(OnMsgHandleHandler),
		onCloseHandler:     object.RequireNonNullElse(onClose, NopOnCloseHandler).(OnCloseHandler),
		onErrorHandler:     object.RequireNonNullElse(onError, NopOnErrorHandler).(OnErrorHandler),
	}
}
func NewServer(h Handler) *Server {
	return NewServerFunc(h, h, h, h, h)
}

type Server struct {
	Addr               string // TCP address to listen on, ":tcp" if empty
	onOpenHandler      OnOpenHandler
	onMsgReadHandler   OnMsgReadHandler
	onMsgHandleHandler OnMsgHandleHandler
	onCloseHandler     OnCloseHandler
	onErrorHandler     OnErrorHandler

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	MaxBytes     int

	ErrorLog *log.Logger

	mu         sync.Mutex
	listeners  map[*net.Listener]struct{}
	activeConn map[*conn]struct{}
	doneChan   chan struct{}
	onShutdown []func()

	// server state
	inShutdown atomic.Bool

	// ConnState specifies an optional callback function that is
	// called when a client connection changes state. See the
	// ConnState type and associated constants for details.
	ConnState func(net.Conn, ConnState)
}

func (srv *Server) CheckError(w io.Writer, r io.Reader, err error) error {
	if err == nil {
		return nil
	}
	return srv.onErrorHandler.OnError(w, r, err)
}

func (srv *Server) ListenAndServe() error {
	if srv.shuttingDown() {
		return srv.CheckError(nil, nil, ErrServerClosed)
	}
	addr := srv.Addr
	if addr == "" {
		addr = ":tcp"
	}
	ln, err := net.Listen("tcp", addr)
	if srv.CheckError(nil, nil, err) != nil {
		return err
	}
	return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}

func (srv *Server) Serve(l net.Listener) error {
	l = &onceCloseListener{Listener: l}
	defer l.Close()

	var tempDelay = time2.NewDefaultDelay() // how long to sleep on accept failure
	ctx := context.WithValue(context.Background(), ServerContextKey, srv)
	for {
		rw, e := l.Accept()
		if e != nil {
			// return if server is cancaled, means normally close
			select {
			case <-srv.getDoneChan():
				return ErrServerClosed
			default:
			}
			// retry if it's recoverable
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				tempDelay.Update()
				srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay.Duration())
				continue
			}
			// return otherwise
			return srv.CheckError(nil, nil, e)
		}
		tempDelay.Reset()

		// takeover the connect
		c := srv.newConn(rw)
		// Handle websocket On
		err := srv.onOpenHandler.OnOpen(c.rwc)
		if err = srv.CheckError(c.w, c.r, err); err != nil {
			c.close()
			return err
		}
		c.setState(c.rwc, StateNew) // before Serve can return
		go c.serve(ctx)
	}
}

func (s *Server) trackConn(c *conn, add bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.activeConn == nil {
		s.activeConn = make(map[*conn]struct{})
	}
	if add {
		s.activeConn[c] = struct{}{}
	} else {
		delete(s.activeConn, c)
	}
}

// Create new connection from rwc.
func (srv *Server) newConn(rwc net.Conn) *conn {
	c := &conn{
		server: srv,
		rwc:    rwc,
	}
	return c
}

func (s *Server) logf(format string, args ...interface{}) {
	if s.ErrorLog != nil {
		s.ErrorLog.Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func ListenAndServe(addr string, readMsg OnMsgReadHandler, handleMsg OnMsgHandleHandler) error {
	server := &Server{Addr: addr, onMsgReadHandler: readMsg, onMsgHandleHandler: handleMsg}
	return server.ListenAndServe()
}

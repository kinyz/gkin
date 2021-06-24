package stream

import (
	"context"
	"errors"
	"gkin/connect"
	"gkin/pb"
	"gkin/utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

func NewStream() *Stream {
	return &Stream{connMgr: connect.NewManager()}
}

const streamKey = "dwigoqhdoiq(U)(J()_"

type Stream struct {
	addr    string
	srv     *grpc.Server
	connMgr connect.Manager
}

func (s *Stream) start() {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		panic("网络异常" + err.Error())
	}
	s.srv = grpc.NewServer()
	pb.RegisterStreamServer(s.srv, s)
	log.Println("stream启动成功.......")
	log.Println(s.addr)
	err = s.srv.Serve(ln)
	if err != nil {
		panic("网络启动异常" + err.Error())
	}

}
func (s *Stream) Serve(addr string) {
	s.addr = addr
	s.start()
}
func (s *Stream) Stop() {
	s.srv.Stop()
}
func (s *Stream) RequestConnect(ctx context.Context, conn *pb.Connection) (*pb.Connection, error) {

	if conn.GetKey() != streamKey {
		return nil, errors.New("验证key错误")
	}
	conn.Token = utils.NewToken()
	s.connMgr.AddOrUpData(conn)
	return conn, nil
}

func (s *Stream) WatchStream(req *pb.RequestListenTopic, server pb.Stream_WatchStreamServer) error {
	if !s.connMgr.Oauth(req.GetConn().GetClientId(), req.GetConn().GetToken()) {
		return errors.New("token验证失败")
	}
	wg := sync.WaitGroup{}
	wg.Add(1)

	server.Context().Done()
	log.Println(" 我到这了")
	wg.Wait()
	return nil
}

func (s *Stream) SendStream(server pb.Stream_SendStreamServer) error {
	panic("implement me")
}

func (s *Stream) oauth(conn *pb.Connection) {

}

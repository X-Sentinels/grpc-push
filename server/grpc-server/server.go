package gs

import (
	"log"
	"net"

	pb "github.com/X-Sentinels/grpc-push/protos"

	"github.com/X-Sentinels/grpc-push/server/g"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	name string
	Strm map[string]pb.PushNotif_RegisterServer
}

func newServer() *Server {
	return &Server{name: "pushnotifServer", Strm: make(map[string]pb.PushNotif_RegisterServer)}
}

func inArray(str string, array []string) bool {
	for _, s := range array {
		if s == str {
			return true
		}
	}
	return false
}

func (s *Server) Register(m *pb.RegistrationRequest, stream pb.PushNotif_RegisterServer) error {
	clientName := m.GetClientName()
	log.Printf("Received a Client Regist %s", clientName)
	s.Strm[clientName] = stream
	clients := g.Config().AliveClients
	if !inArray(clientName, clients) {
		g.Config().AliveClients = append(clients, clientName)
	}
	log.Printf("Client %s will now recieve streams", m.GetClientName())
	for {
	}
}

func (s *Server) pushUpdates() {
	for {
		message := <-g.NotifMessage
		log.Printf("Send Message To Client: %s", message.ClientName)
		for k, v := range s.Strm {
			if k == message.ClientName {
				if err := v.Send(&pb.RegistrationResponse{Notice: message.Notification}); err != nil {
					log.Fatalf("Send failed %v", err)
				}
			}
		}
	}
}

func Start() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	myServer := newServer()
	pb.RegisterPushNotifServer(s, myServer)

	go myServer.pushUpdates()

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

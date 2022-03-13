package haberdasherserver

import (
    "context"
    "math/rand"

    "github.com/twitchtv/twirp"
    pb "gitlab.com/yukung/hello-twirp/rpc/haberdasher"
)

type Server struct {}

func (s *Server) MakeHat(ctx context.Context, size *pb.Size) (hat *pb.Hat, err error) {
    if size.Inches <= 0 {
        return nil, twirp.InvalidArgumentError("Inches", "I can't make a hat that small!")
    }
    return &pb.Hat{
        Inches: size.Inches,
        Color: []string{"white", "black", "brown", "red", "blue"}[rand.Intn(5)],
        Name: []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(4)],
    }, nil
}

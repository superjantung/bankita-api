package gapi

import (
	"fmt"

	db "github.com/superjantung/bankita-api/db/sqlc"
	"github.com/superjantung/bankita-api/pb"
	"github.com/superjantung/bankita-api/token"
	"github.com/superjantung/bankita-api/util"
)

type Server struct {
	pb.UnimplementedBankitaServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

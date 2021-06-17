package service

import (
	"context"
	"log"

	"github.com/404th/grpce/protos/protos"
	"github.com/404th/grpce/service/grpc_client"
	"github.com/404th/grpce/storage"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthService struct {
	strg storage.StorageI
	client *grpc_client.GrpcClient
}

func NewAuthService(db *sqlx.DB, client *grpc_client.GrpcClient) *AuthService {
	return &AuthService{
		strg: storage.NewStoragePg(db),
		client: client,
	}
}

func (srv *AuthService) Create(ctx context.Context, req *protos.CreateRequest ) (*protos.CreateResponse, error) {
	var (
		id string
		createdUser protos.CreateResponse
	)

	id = uuid.New().String()

	created_id, is_created, err := srv.strg.Auth().Create(id, req.GetUsername(), req.GetAge(), req.GetIsAdult())
	if err != nil {
		log.Printf("error while creating user: %v", err)
		return nil, err
	}

	createdUser.Id = created_id
	createdUser.IsCreated = is_created

	return &createdUser, nil
}

func (srv *AuthService) Get(ctx context.Context, req *protos.WithID) (*protos.User, error) {

	usr := &protos.User{}

	user, err := srv.strg.Auth().Get( req.GetId() )
	if err != nil {
		log.Printf("error while getting user: %v", err)
		return nil, err
	}

	usr.Age = user.Age
	usr.Id = user.ID
	usr.Username = user.Username
	usr.IsAdult = user.IsAdult

	return usr, nil
}

func (srv *AuthService) Delete(ctx context.Context, req *protos.WithID) (*protos.DeleteResponse, error) {
	id, is_deleted, err := srv.strg.Auth().Delete(req.GetId())
	if err != nil {
		log.Printf("error while deleting user: %v", err)
		return nil, err
	}

	return &protos.DeleteResponse{
		Id: id,
		IsDeleted: is_deleted,
	}, nil
} 

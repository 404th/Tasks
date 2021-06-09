package repo

import (
	"github.com/404th/contactlistgrpc/postgres"
	"github.com/404th/contactlistgrpc/service/grpcclient"
)

type StorageI interface {
	Contact() ContactStorageI
}

type storagePg struct {
	contactRepo ContactStorageI
}

func NewStoragePg(client *grpcclient.GrpcClient) StorageI {
	return &storagePg{
		contactRepo: repo.New
	}
}

func (s storagePg) Contact() ContactStorageI {
	return s.contactRepo
}

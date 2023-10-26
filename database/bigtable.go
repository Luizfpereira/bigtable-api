package database

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/bigtable"
)

var admInstance *bigtable.AdminClient
var clientInstance *bigtable.Client
var lock = &sync.Mutex{}

func GetAdminClientSingleton(ctx context.Context) (*bigtable.AdminClient, error) {
	if admInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if admInstance == nil {
			log.Println("Creating bigtable admin instance")
			var err error
			admInstance, err = bigtable.NewAdminClient(ctx, os.Getenv("PROJECT_ID"), os.Getenv("_INSTANCE_ID"))
			if err != nil {
				return nil, err
			}
		} else {
			log.Println("Admin instance connected!")
		}
	} else {
		log.Println("Admin instance connected!")
	}
	return admInstance, nil
}

func GetClientSingleton(ctx context.Context) (*bigtable.Client, error) {
	if clientInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if clientInstance == nil {
			log.Println("Creating bigtable client instance")
			var err error
			clientInstance, err = bigtable.NewClient(ctx, os.Getenv("PROJECT_ID"), os.Getenv("_INSTANCE_ID"))
			if err != nil {
				return nil, err
			}
		} else {
			log.Println("Client instance connected!")
		}
	} else {
		log.Println("Client instance connected!")
	}
	return clientInstance, nil
}

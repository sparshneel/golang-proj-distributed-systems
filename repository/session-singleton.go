package repository

import (
	"github.com/gocql/gocql"

	"sync"

	"golang-proj-distributed-systems/helpers"
	)

var once sync.Once

// type global
var sessionInstance *gocql.Session
var err error

func  GetSessionInstance() *gocql.Session {

	once.Do(func() {
		sessionInstance, err = getClusterInstance().CreateSession()
		if err != nil {
			panic("error connecting session to cassandra cluster " + helpers.Cluster)
		}
	})

	return sessionInstance
}


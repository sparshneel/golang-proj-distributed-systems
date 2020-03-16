package repository

import (
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"

	"sync"

	"golang-proj-distributed-systems/helpers"
	)

var once sync.Once

// type global
var sessionInstance *gocql.Session
var err error

func  GetSessionInstance() *gocql.Session {

	once.Do(func() {
		clusterInstance := gocql.NewCluster(helpers.Cluster)
		clusterInstance.Keyspace = helpers.Keyspace
		sessionInstance, err = clusterInstance.CreateSession()
		if err != nil {
			logrus.Error("error connecting session to cassandra cluster " + helpers.Cluster)
			panic("error connecting session to cassandra cluster " + helpers.Cluster)
		}
	})

	return sessionInstance
}


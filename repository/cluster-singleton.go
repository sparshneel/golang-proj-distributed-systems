package repository

import (
	"github.com/gocql/gocql"

	"golang-proj-distributed-systems/helpers"

	"sync"
)

var clusterOnce sync.Once

// type global
var clusterInstance *gocql.ClusterConfig

func getClusterInstance() *gocql.ClusterConfig{

	clusterOnce.Do(func() { // <-- atomic, does not allow repeating
		clusterInstance := gocql.NewCluster(helpers.Cluster)
		clusterInstance.Keyspace = helpers.Keyspace
	})
	return clusterInstance
}

package fees

import (
	"context"
	"fmt"

	encore "encore.dev"
	"encore.dev/storage/cache"
	"encore.dev/storage/sqldb"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/dsha256/feesapi/entity"
	"github.com/dsha256/feesapi/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

var (
	EnvName          = encore.Meta().Environment.Name
	ServiceTaskQueue = EnvName + "-fees"
)

var BillingDB = sqldb.NewDatabase("fees", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

var CacheCluster = cache.NewCluster("fees", cache.ClusterConfig{
	EvictionPolicy: cache.NoEviction,
})

type TemporalWorkflowsIDsCacheKey struct {
	UUID string
}

var TemporalWorkflowsIDsCache = cache.NewStringKeyspace[TemporalWorkflowsIDsCacheKey](CacheCluster, cache.KeyspaceConfig{
	KeyPattern:    "workflows/:UUID",
	DefaultExpiry: nil,
})

//encore:service
type Fees struct {
	client          client.Client
	worker          worker.Worker
	entity          *entity.Client
	workflowIDCache *cache.StringKeyspace[TemporalWorkflowsIDsCacheKey]
}

func initFees() (*Fees, error) {
	fees := &Fees{}
	fees.workflowIDCache = TemporalWorkflowsIDsCache

	dbDriver := entsql.OpenDB(dialect.Postgres, BillingDB.Stdlib())
	entClient := entity.NewClient(entity.Driver(dbDriver))
	fees.entity = entClient

	c, err := client.Dial(client.Options{})
	if err != nil {
		return nil, fmt.Errorf("create temporal client: %v", err)
	}
	fees.client = c

	w := worker.New(c, ServiceTaskQueue, worker.Options{})

	w.RegisterWorkflow(workflow.BillWorkflow)
	w.RegisterActivity(&workflow.Activities{
		Entity: entClient,
	})

	err = w.Start()
	if err != nil {
		c.Close()
		return nil, fmt.Errorf("start temporal worker: %v", err)
	}
	fees.worker = w

	return fees, nil
}

func (f *Fees) Shutdown(force context.Context) {
	f.client.Close()
	f.worker.Stop()
}

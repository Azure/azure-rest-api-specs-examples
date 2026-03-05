package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-01-01-preview/ServersRestartWithFailover.json
func ExampleServersClient_BeginRestart_restartPostgreSqlDatabaseEngineInAServerWithAForcedFailoverToStandbyServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginRestart(ctx, "exampleresourcegroup", "exampleserver", &armpostgresqlflexibleservers.ServersClientBeginRestartOptions{
		Parameters: &armpostgresqlflexibleservers.RestartParameter{
			FailoverMode:        to.Ptr(armpostgresqlflexibleservers.FailoverModeForcedFailover),
			RestartWithFailover: to.Ptr(true),
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

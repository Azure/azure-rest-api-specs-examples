package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPool.json
func ExampleSQLPoolsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolsClient("01234567-89ab-4def-0123-456789abcdef", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"ExampleResourceGroup",
		"ExampleWorkspace",
		"ExampleSqlPool",
		armsynapse.SQLPool{
			Location: to.Ptr("Southeast Asia"),
			Tags:     map[string]*string{},
			Properties: &armsynapse.SQLPoolResourceProperties{
				Collation:             to.Ptr(""),
				CreateMode:            to.Ptr(armsynapse.CreateMode("")),
				MaxSizeBytes:          to.Ptr[int64](0),
				RecoverableDatabaseID: to.Ptr(""),
				SourceDatabaseID:      to.Ptr(""),
				StorageAccountType:    to.Ptr(armsynapse.StorageAccountTypeLRS),
			},
			SKU: &armsynapse.SKU{
				Name: to.Ptr(""),
				Tier: to.Ptr(""),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

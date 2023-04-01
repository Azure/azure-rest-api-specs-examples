package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateSqlPool.json
func ExampleSQLPoolsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLPoolsClient().BeginUpdate(ctx, "ExampleResourceGroup", "ExampleWorkspace", "ExampleSqlPool", armsynapse.SQLPoolPatchInfo{
		Location: to.Ptr("West US 2"),
		Properties: &armsynapse.SQLPoolResourceProperties{
			Collation:          to.Ptr(""),
			MaxSizeBytes:       to.Ptr[int64](0),
			RestorePointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t }()),
		},
		SKU: &armsynapse.SKU{
			Name: to.Ptr(""),
			Tier: to.Ptr(""),
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLPool = armsynapse.SQLPool{
	// 	Name: to.Ptr("ExampleSqlPool"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools"),
	// 	ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspaces/sqlPools/ExampleSqlPool"),
	// 	Location: to.Ptr("West US 2"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsynapse.SQLPoolResourceProperties{
	// 		Collation: to.Ptr(""),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t}()),
	// 		MaxSizeBytes: to.Ptr[int64](0),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RecoverableDatabaseID: to.Ptr(""),
	// 		RestorePointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t}()),
	// 		Status: to.Ptr("Paused"),
	// 		StorageAccountType: to.Ptr(armsynapse.StorageAccountTypeGRS),
	// 	},
	// 	SKU: &armsynapse.SKU{
	// 		Name: to.Ptr(""),
	// 		Tier: to.Ptr(""),
	// 	},
	// }
}

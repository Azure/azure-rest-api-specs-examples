Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv1.0.0/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationMigrationItems_Resync.json
func ExampleReplicationMigrationItemsClient_BeginResync() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("migrationvault",
		"resourcegroup1",
		"cb53d0c3-bd59-4721-89bc-06916a9147ef", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginResync(ctx,
		"vmwarefabric1",
		"vmwareContainer1",
		"virtualmachine1",
		armrecoveryservicessiterecovery.ResyncInput{
			Properties: &armrecoveryservicessiterecovery.ResyncInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtResyncInput{
					InstanceType: to.Ptr("VMwareCbt"),
					SkipCbtReset: to.Ptr("true"),
				},
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
```

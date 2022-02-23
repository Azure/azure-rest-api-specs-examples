Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv0.2.1/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationMigrationItems_Create.json
func ExampleReplicationMigrationItemsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationMigrationItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<migration-item-name>",
		armrecoveryservicessiterecovery.EnableMigrationInput{
			Properties: &armrecoveryservicessiterecovery.EnableMigrationInputProperties{
				PolicyID: to.StringPtr("<policy-id>"),
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.VMwareCbtEnableMigrationInput{
					InstanceType:            to.StringPtr("<instance-type>"),
					DataMoverRunAsAccountID: to.StringPtr("<data-mover-run-as-account-id>"),
					DisksToInclude: []*armrecoveryservicessiterecovery.VMwareCbtDiskInput{
						{
							DiskID:                         to.StringPtr("<disk-id>"),
							IsOSDisk:                       to.StringPtr("<is-osdisk>"),
							LogStorageAccountID:            to.StringPtr("<log-storage-account-id>"),
							LogStorageAccountSasSecretName: to.StringPtr("<log-storage-account-sas-secret-name>"),
						}},
					SnapshotRunAsAccountID: to.StringPtr("<snapshot-run-as-account-id>"),
					TargetNetworkID:        to.StringPtr("<target-network-id>"),
					TargetResourceGroupID:  to.StringPtr("<target-resource-group-id>"),
					VmwareMachineID:        to.StringPtr("<vmware-machine-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ReplicationMigrationItemsClientCreateResult)
}
```

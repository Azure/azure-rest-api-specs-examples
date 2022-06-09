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

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationProtectedItems_AddDisks.json
func ExampleReplicationProtectedItemsClient_BeginAddDisks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginAddDisks(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<replicated-protected-item-name>",
		armrecoveryservicessiterecovery.AddDisksInput{
			Properties: &armrecoveryservicessiterecovery.AddDisksInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AAddDisksInput{
					AddDisksProviderSpecificInput: armrecoveryservicessiterecovery.AddDisksProviderSpecificInput{
						InstanceType: to.StringPtr("<instance-type>"),
					},
					VMDisks: []*armrecoveryservicessiterecovery.A2AVMDiskInputDetails{
						{
							DiskURI:                             to.StringPtr("<disk-uri>"),
							PrimaryStagingAzureStorageAccountID: to.StringPtr("<primary-staging-azure-storage-account-id>"),
							RecoveryAzureStorageAccountID:       to.StringPtr("<recovery-azure-storage-account-id>"),
						}},
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
	log.Printf("ReplicationProtectedItem.ID: %s\n", *res.ID)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv0.1.0/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

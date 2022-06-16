package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationProtectedItems_RemoveDisks.json
func ExampleReplicationProtectedItemsClient_BeginRemoveDisks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginRemoveDisks(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<replicated-protected-item-name>",
		armrecoveryservicessiterecovery.RemoveDisksInput{
			Properties: &armrecoveryservicessiterecovery.RemoveDisksInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ARemoveDisksInput{
					RemoveDisksProviderSpecificInput: armrecoveryservicessiterecovery.RemoveDisksProviderSpecificInput{
						InstanceType: to.StringPtr("<instance-type>"),
					},
					VMDisksUris: []*string{
						to.StringPtr("https://vmstorage.blob.core.windows.net/vhds/datadisk1.vhd")},
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

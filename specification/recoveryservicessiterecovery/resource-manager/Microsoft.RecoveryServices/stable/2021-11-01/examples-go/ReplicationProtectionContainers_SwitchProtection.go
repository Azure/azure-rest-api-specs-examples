package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationProtectionContainers_SwitchProtection.json
func ExampleReplicationProtectionContainersClient_BeginSwitchProtection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectionContainersClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginSwitchProtection(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		armrecoveryservicessiterecovery.SwitchProtectionInput{
			Properties: &armrecoveryservicessiterecovery.SwitchProtectionInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ASwitchProtectionInput{
					InstanceType: to.StringPtr("<instance-type>"),
				},
				ReplicationProtectedItemName: to.StringPtr("<replication-protected-item-name>"),
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
	log.Printf("Response result: %#v\n", res.ReplicationProtectionContainersClientSwitchProtectionResult)
}

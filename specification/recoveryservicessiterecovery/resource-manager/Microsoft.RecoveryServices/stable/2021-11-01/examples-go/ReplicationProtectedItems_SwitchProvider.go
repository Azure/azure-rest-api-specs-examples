package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationProtectedItems_SwitchProvider.json
func ExampleReplicationProtectedItemsClient_BeginSwitchProvider() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginSwitchProvider(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		"<replicated-protected-item-name>",
		armrecoveryservicessiterecovery.SwitchProviderInput{
			Properties: &armrecoveryservicessiterecovery.SwitchProviderInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.InMageAzureV2SwitchProviderInput{
					InstanceType:      to.StringPtr("<instance-type>"),
					TargetApplianceID: to.StringPtr("<target-appliance-id>"),
					TargetFabricID:    to.StringPtr("<target-fabric-id>"),
					TargetVaultID:     to.StringPtr("<target-vault-id>"),
				},
				TargetInstanceType: to.StringPtr("<target-instance-type>"),
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
	log.Printf("Response result: %#v\n", res.ReplicationProtectedItemsClientSwitchProviderResult)
}

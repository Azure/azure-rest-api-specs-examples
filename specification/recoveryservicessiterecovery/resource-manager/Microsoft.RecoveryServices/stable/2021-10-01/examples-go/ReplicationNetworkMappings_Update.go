package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationNetworkMappings_Update.json
func ExampleReplicationNetworkMappingsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationNetworkMappingsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<fabric-name>",
		"<network-name>",
		"<network-mapping-name>",
		armrecoveryservicessiterecovery.UpdateNetworkMappingInput{
			Properties: &armrecoveryservicessiterecovery.UpdateNetworkMappingInputProperties{
				FabricSpecificDetails: &armrecoveryservicessiterecovery.VmmToAzureUpdateNetworkMappingInput{
					FabricSpecificUpdateNetworkMappingInput: armrecoveryservicessiterecovery.FabricSpecificUpdateNetworkMappingInput{
						InstanceType: to.StringPtr("<instance-type>"),
					},
				},
				RecoveryFabricName: to.StringPtr("<recovery-fabric-name>"),
				RecoveryNetworkID:  to.StringPtr("<recovery-network-id>"),
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
	log.Printf("NetworkMapping.ID: %s\n", *res.ID)
}

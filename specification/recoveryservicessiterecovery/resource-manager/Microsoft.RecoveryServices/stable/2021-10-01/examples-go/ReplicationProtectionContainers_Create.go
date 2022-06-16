package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationProtectionContainers_Create.json
func ExampleReplicationProtectionContainersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectionContainersClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<fabric-name>",
		"<protection-container-name>",
		armrecoveryservicessiterecovery.CreateProtectionContainerInput{
			Properties: &armrecoveryservicessiterecovery.CreateProtectionContainerInputProperties{
				ProviderSpecificInput: []armrecoveryservicessiterecovery.ReplicationProviderSpecificContainerCreationInputClassification{
					&armrecoveryservicessiterecovery.ReplicationProviderSpecificContainerCreationInput{
						InstanceType: to.StringPtr("<instance-type>"),
					}},
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
	log.Printf("ProtectionContainer.ID: %s\n", *res.ID)
}

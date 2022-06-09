```go
package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectionIntents_Create.json
func ExampleReplicationProtectionIntentsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient("vault1",
		"resourceGroupPS1",
		"509099b2-9d2c-4636-b43e-bd5cafb6be69", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"vm1",
		armrecoveryservicessiterecovery.CreateProtectionIntentInput{
			Properties: &armrecoveryservicessiterecovery.CreateProtectionIntentProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ACreateProtectionIntentInput{
					InstanceType:             to.Ptr("A2A"),
					FabricObjectID:           to.Ptr("/subscriptions/509099b2-9d2c-4636-b43e-bd5cafb6be69/resourceGroups/removeOne/providers/Microsoft.Compute/virtualMachines/vmPpgAv5"),
					PrimaryLocation:          to.Ptr("eastUs2"),
					RecoveryAvailabilityType: to.Ptr(armrecoveryservicessiterecovery.A2ARecoveryAvailabilityTypeSingle),
					RecoveryLocation:         to.Ptr("westus2"),
					RecoveryResourceGroupID:  to.Ptr("/subscriptions/509099b2-9d2c-4636-b43e-bd5cafb6be69/resourceGroups/removeOne-asr"),
					RecoverySubscriptionID:   to.Ptr("ed5bcdf6-d61e-47bd-8ea9-f2bd379a2640"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv1.0.0/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

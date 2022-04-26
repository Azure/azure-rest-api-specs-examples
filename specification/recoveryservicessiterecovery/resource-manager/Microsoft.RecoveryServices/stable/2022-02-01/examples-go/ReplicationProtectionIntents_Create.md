Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv0.4.0/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<intent-object-name>",
		armrecoveryservicessiterecovery.CreateProtectionIntentInput{
			Properties: &armrecoveryservicessiterecovery.CreateProtectionIntentProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ACreateProtectionIntentInput{
					InstanceType:             to.Ptr("<instance-type>"),
					FabricObjectID:           to.Ptr("<fabric-object-id>"),
					PrimaryLocation:          to.Ptr("<primary-location>"),
					RecoveryAvailabilityType: to.Ptr(armrecoveryservicessiterecovery.A2ARecoveryAvailabilityTypeSingle),
					RecoveryLocation:         to.Ptr("<recovery-location>"),
					RecoveryResourceGroupID:  to.Ptr("<recovery-resource-group-id>"),
					RecoverySubscriptionID:   to.Ptr("<recovery-subscription-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

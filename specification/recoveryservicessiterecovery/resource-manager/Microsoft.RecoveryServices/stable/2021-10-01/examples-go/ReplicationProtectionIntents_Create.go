package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationProtectionIntents_Create.json
func ExampleReplicationProtectionIntentsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<intent-object-name>",
		armrecoveryservicessiterecovery.CreateProtectionIntentInput{
			Properties: &armrecoveryservicessiterecovery.CreateProtectionIntentProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ACreateProtectionIntentInput{
					CreateProtectionIntentProviderSpecificDetails: armrecoveryservicessiterecovery.CreateProtectionIntentProviderSpecificDetails{
						InstanceType: to.StringPtr("<instance-type>"),
					},
					FabricObjectID:           to.StringPtr("<fabric-object-id>"),
					PrimaryLocation:          to.StringPtr("<primary-location>"),
					RecoveryAvailabilityType: armrecoveryservicessiterecovery.A2ARecoveryAvailabilityTypeSingle.ToPtr(),
					RecoveryLocation:         to.StringPtr("<recovery-location>"),
					RecoveryResourceGroupID:  to.StringPtr("<recovery-resource-group-id>"),
					RecoverySubscriptionID:   to.StringPtr("<recovery-subscription-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ReplicationProtectionIntent.ID: %s\n", *res.ID)
}

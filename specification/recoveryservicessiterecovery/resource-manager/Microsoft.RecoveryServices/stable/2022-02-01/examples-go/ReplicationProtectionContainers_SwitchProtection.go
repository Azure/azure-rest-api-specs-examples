package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectionContainers_SwitchProtection.json
func ExampleReplicationProtectionContainersClient_BeginSwitchProtection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectionContainersClient("priyanponeboxvault",
		"priyanprg",
		"42195872-7e70-4f8a-837f-84b28ecbb78b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSwitchProtection(ctx,
		"CentralUSCanSite",
		"CentralUSCancloud",
		armrecoveryservicessiterecovery.SwitchProtectionInput{
			Properties: &armrecoveryservicessiterecovery.SwitchProtectionInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ASwitchProtectionInput{
					InstanceType: to.Ptr("A2A"),
				},
				ReplicationProtectedItemName: to.Ptr("a2aSwapOsVm"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

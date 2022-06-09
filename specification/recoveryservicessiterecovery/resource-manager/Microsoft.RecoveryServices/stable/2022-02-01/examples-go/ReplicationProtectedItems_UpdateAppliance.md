```go
package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectedItems_UpdateAppliance.json
func ExampleReplicationProtectedItemsClient_BeginUpdateAppliance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectedItemsClient("Ayan-0106-SA-Vault",
		"Ayan-0106-SA-RG",
		"7c943c1b-5122-4097-90c8-861411bdd574", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateAppliance(ctx,
		"Ayan-0106-SA-Vaultreplicationfabric",
		"Ayan-0106-SA-Vaultreplicationcontainer",
		"idclab-vcen67_50158124-8857-3e08-0893-2ddf8ebb8c1f",
		armrecoveryservicessiterecovery.UpdateApplianceForReplicationProtectedItemInput{
			Properties: &armrecoveryservicessiterecovery.UpdateApplianceForReplicationProtectedItemInputProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.InMageRcmUpdateApplianceForReplicationProtectedItemInput{
					InstanceType:   to.Ptr("InMageRcm"),
					RunAsAccountID: to.Ptr(""),
				},
				TargetApplianceID: to.Ptr(""),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv1.0.0/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

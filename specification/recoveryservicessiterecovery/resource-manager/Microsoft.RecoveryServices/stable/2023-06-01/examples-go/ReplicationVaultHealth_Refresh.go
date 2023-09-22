package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationVaultHealth_Refresh.json
func ExampleReplicationVaultHealthClient_BeginRefresh() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationVaultHealthClient().BeginRefresh(ctx, "vault1", "resourceGroupPS1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VaultHealthDetails = armrecoveryservicessiterecovery.VaultHealthDetails{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationVaultHealth"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationVaultHealth/Default"),
	// 	Properties: &armrecoveryservicessiterecovery.VaultHealthProperties{
	// 		FabricsHealth: &armrecoveryservicessiterecovery.ResourceHealthSummary{
	// 			Issues: []*armrecoveryservicessiterecovery.HealthErrorSummary{
	// 			},
	// 			ResourceCount: to.Ptr[int32](1),
	// 		},
	// 		ProtectedItemsHealth: &armrecoveryservicessiterecovery.ResourceHealthSummary{
	// 			Issues: []*armrecoveryservicessiterecovery.HealthErrorSummary{
	// 			},
	// 			ResourceCount: to.Ptr[int32](2),
	// 		},
	// 		VaultErrors: []*armrecoveryservicessiterecovery.HealthError{
	// 		},
	// 	},
	// }
}

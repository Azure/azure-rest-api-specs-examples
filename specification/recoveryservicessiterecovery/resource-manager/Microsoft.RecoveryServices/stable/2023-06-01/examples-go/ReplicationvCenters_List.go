package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationvCenters_List.json
func ExampleReplicationvCentersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationvCentersClient().NewListPager("MadhaviVault", "MadhaviVRG", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VCenterCollection = armrecoveryservicessiterecovery.VCenterCollection{
		// 	Value: []*armrecoveryservicessiterecovery.VCenter{
		// 		{
		// 			Name: to.Ptr("esx-78"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationvCenters"),
		// 			ID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/MadhaviVRG/providers/Microsoft.RecoveryServices/vaults/MadhaviVault/replicationFabrics/239f778f368e34f78216d81f030725cdf2033174b47879b9f2eeede06fdd9c4d/replicationvCenters/esx-78"),
		// 			Properties: &armrecoveryservicessiterecovery.VCenterProperties{
		// 				DiscoveryStatus: to.Ptr("Pending"),
		// 				FabricArmResourceName: to.Ptr("239f778f368e34f78216d81f030725cdf2033174b47879b9f2eeede06fdd9c4d"),
		// 				FriendlyName: to.Ptr("esx-78"),
		// 				InternalID: to.Ptr("inmtest78"),
		// 				IPAddress: to.Ptr("inmtest78"),
		// 				Port: to.Ptr("443"),
		// 				ProcessServerID: to.Ptr("5A720CAB-39CB-F445-BD1662B0B33164B5"),
		// 				RunAsAccountID: to.Ptr("2"),
		// 			},
		// 	}},
		// }
	}
}

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationProtectionContainers_ListByReplicationFabrics.json
func ExampleReplicationProtectionContainersClient_NewListByReplicationFabricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationProtectionContainersClient().NewListByReplicationFabricsPager("vault1", "resourceGroupPS1", "cloud1", nil)
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
		// page.ProtectionContainerCollection = armrecoveryservicessiterecovery.ProtectionContainerCollection{
		// 	Value: []*armrecoveryservicessiterecovery.ProtectionContainer{
		// 		{
		// 			Name: to.Ptr("cloud_6d224fc6-f326-5d35-96de-fbf51efb3179"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179"),
		// 			Properties: &armrecoveryservicessiterecovery.ProtectionContainerProperties{
		// 				FabricFriendlyName: to.Ptr("cloud1"),
		// 				FabricType: to.Ptr("HyperVSite"),
		// 				FriendlyName: to.Ptr("cloud1"),
		// 				PairingStatus: to.Ptr("NotPaired"),
		// 				ProtectedItemCount: to.Ptr[int32](0),
		// 				Role: to.Ptr(""),
		// 			},
		// 	}},
		// }
	}
}

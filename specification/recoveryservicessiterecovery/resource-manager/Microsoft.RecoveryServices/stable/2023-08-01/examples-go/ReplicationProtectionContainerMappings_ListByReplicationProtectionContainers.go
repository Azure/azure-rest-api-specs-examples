package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationProtectionContainerMappings_ListByReplicationProtectionContainers.json
func ExampleReplicationProtectionContainerMappingsClient_NewListByReplicationProtectionContainersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationProtectionContainerMappingsClient().NewListByReplicationProtectionContainersPager("vault1", "resourceGroupPS1", "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", nil)
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
		// page.ProtectionContainerMappingCollection = armrecoveryservicessiterecovery.ProtectionContainerMappingCollection{
		// 	Value: []*armrecoveryservicessiterecovery.ProtectionContainerMapping{
		// 		{
		// 			Name: to.Ptr("cloud1protectionprofile1"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionContainerMappings"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectionContainerMappings/cloud1protectionprofile1"),
		// 			Properties: &armrecoveryservicessiterecovery.ProtectionContainerMappingProperties{
		// 				Health: to.Ptr("Normal"),
		// 				HealthErrorDetails: []*armrecoveryservicessiterecovery.HealthError{
		// 				},
		// 				PolicyFriendlyName: to.Ptr("protectionprofile1"),
		// 				PolicyID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionprofile1"),
		// 				SourceFabricFriendlyName: to.Ptr("cloud1"),
		// 				SourceProtectionContainerFriendlyName: to.Ptr("cloud1"),
		// 				State: to.Ptr("Paired"),
		// 				TargetFabricFriendlyName: to.Ptr("Microsoft Azure"),
		// 				TargetProtectionContainerFriendlyName: to.Ptr("Microsoft Azure"),
		// 				TargetProtectionContainerID: to.Ptr("Microsoft Azure"),
		// 			},
		// 	}},
		// }
	}
}

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationStorageClassificationMappings_Get.json
func ExampleReplicationStorageClassificationMappingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationStorageClassificationMappingsClient().Get(ctx, "vault1", "resourceGroupPS1", "2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0", "8891569e-aaef-4a46-a4a0-78c14f2d7b09", "testStorageMapping", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageClassificationMapping = armrecoveryservicessiterecovery.StorageClassificationMapping{
	// 	Name: to.Ptr("testStorageMapping"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationStorageClassifications/replicationStorageClassificationMappings"),
	// 	ID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0/replicationStorageClassifications/8891569e-aaef-4a46-a4a0-78c14f2d7b09/replicationStorageClassificationMappings/testStorageMapping"),
	// 	Properties: &armrecoveryservicessiterecovery.StorageClassificationMappingProperties{
	// 		TargetStorageClassificationID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0/replicationStorageClassifications/8891569e-aaef-4a46-a4a0-78c14f2d7b09"),
	// 	},
	// }
}

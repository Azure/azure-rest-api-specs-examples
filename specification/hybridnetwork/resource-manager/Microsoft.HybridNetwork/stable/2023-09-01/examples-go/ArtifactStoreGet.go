package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ArtifactStoreGet.json
func ExampleArtifactStoresClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArtifactStoresClient().Get(ctx, "rg", "TestPublisher", "TestArtifactStoreName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArtifactStore = armhybridnetwork.ArtifactStore{
	// 	Name: to.Ptr("TestArtifactStore"),
	// 	Type: to.Ptr("microsoft.hybridnetwork/publishers/artifactStores"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armhybridnetwork.ArtifactStorePropertiesFormat{
	// 		ManagedResourceGroupConfiguration: &armhybridnetwork.ArtifactStorePropertiesFormatManagedResourceGroupConfiguration{
	// 			Name: to.Ptr("testRg"),
	// 			Location: to.Ptr("eastus"),
	// 		},
	// 		ReplicationStrategy: to.Ptr(armhybridnetwork.ArtifactReplicationStrategySingleReplication),
	// 		StorageResourceID: to.Ptr("TestResourceId"),
	// 		StoreType: to.Ptr(armhybridnetwork.ArtifactStoreTypeAzureContainerRegistry),
	// 	},
	// }
}

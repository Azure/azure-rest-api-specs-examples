package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dc4c1eaef16e0bc8b1e96c3d1e014deb96259b35/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/RegistryGetPrivateLinkResource.json
func ExampleRegistriesClient_GetPrivateLinkResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegistriesClient().GetPrivateLinkResource(ctx, "myResourceGroup", "myRegistry", "registry", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armcontainerregistry.PrivateLinkResource{
	// 	Name: to.Ptr("registry"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/privateLinkResources/registry"),
	// 	Properties: &armcontainerregistry.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("registry"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("registry"),
	// 			to.Ptr("registry_data_myregion")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.azurecr.io")},
	// 			},
	// 		}
}

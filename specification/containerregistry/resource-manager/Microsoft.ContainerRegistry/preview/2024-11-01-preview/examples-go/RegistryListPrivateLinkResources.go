package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/300ff7c27c481d7074af06cd95a152aaea80ed2b/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2024-11-01-preview/examples/RegistryListPrivateLinkResources.json
func ExampleRegistriesClient_NewListPrivateLinkResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegistriesClient().NewListPrivateLinkResourcesPager("myResourceGroup", "myRegistry", nil)
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
		// page.PrivateLinkResourceListResult = armcontainerregistry.PrivateLinkResourceListResult{
		// 	Value: []*armcontainerregistry.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("registry"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/privateLinkResources/registry"),
		// 			Properties: &armcontainerregistry.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("registry"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("registry"),
		// 					to.Ptr("registry_data_myregion")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.azurecr.io")},
		// 					},
		// 			}},
		// 		}
	}
}

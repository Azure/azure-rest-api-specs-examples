package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PrivateLinkResources_ListByResource.json
func ExamplePrivateLinkResourcesClient_NewListByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbiprivatelinks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByResourcePager("resourceGroup", "azureResourceName", nil)
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
		// page.PrivateLinkResourcesListResult = armpowerbiprivatelinks.PrivateLinkResourcesListResult{
		// 	Value: []*armpowerbiprivatelinks.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("tenant"),
		// 			Type: to.Ptr("Microsoft.PowerBI/{resourceType}/privateLinkResources"),
		// 			ID: to.Ptr("subscriptions/a0020869-4d28-422a-89f4-c2413130d73c/resourceGroups/resourceGroup/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateLinkResources/tenant"),
		// 			Properties: &armpowerbiprivatelinks.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("tenant"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("tenant"),
		// 					to.Ptr("capacity:3df897a4f10b49e9bddb0e9cf062adba")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.powerbi.com")},
		// 					},
		// 			}},
		// 		}
	}
}

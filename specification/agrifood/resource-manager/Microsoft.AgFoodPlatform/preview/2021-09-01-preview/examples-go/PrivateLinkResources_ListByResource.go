package armagrifood_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a7af6049f4b4743ef3b649f3852bcc7bd9a43ee0/specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/PrivateLinkResources_ListByResource.json
func ExamplePrivateLinkResourcesClient_NewListByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armagrifood.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByResourcePager("examples-rg", "examples-farmbeatsResourceName", nil)
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
		// page.PrivateLinkResourceListResult = armagrifood.PrivateLinkResourceListResult{
		// 	Value: []*armagrifood.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("farmbeats"),
		// 			Type: to.Ptr("Microsoft.AgFoodPlatform/farmBeats/privateLinkResources"),
		// 			ID: to.Ptr("subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.AgFoodPlatform/farmBeats/examples-farmbeatsResourceName/privateLinkResources/farmbeats"),
		// 			Properties: &armagrifood.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("farmbeats"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("farmbeats")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.farmbeats.azure.net")},
		// 					},
		// 			}},
		// 		}
	}
}

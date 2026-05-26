package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2026-01-01/PrivateLinkResources/PrivateLinkResources_ListByPrivateLink.json
func ExamplePrivateLinkResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListPager("rg", armsecurity.PrivateLinkParameters{PrivateLinkName: "pls"}, nil)
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
		// page = armsecurity.PrivateLinkResourcesClientListResponse{
		// 	PrivateLinkGroupResourceListResult: armsecurity.PrivateLinkGroupResourceListResult{
		// 		Value: []*armsecurity.PrivateLinkGroupResource{
		// 			{
		// 				Name: to.Ptr("containers"),
		// 				Type: to.Ptr("Microsoft.Security/privateLinks/privateLinkResources"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg/providers/Microsoft.Security/privateLinks/pls/privateLinkResources/containers"),
		// 				Properties: &armsecurity.PrivateLinkResourceProperties{
		// 					GroupID: to.Ptr("containers"),
		// 					RequiredMembers: []*string{
		// 						to.Ptr("api"),
		// 						to.Ptr("data-eastus"),
		// 					},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.cloud.defender.microsoft.com"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

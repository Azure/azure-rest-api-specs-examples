package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkVirtualApplianceSiteList.json
func ExampleVirtualApplianceSitesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualApplianceSitesClient().NewListPager("rg1", "nva", nil)
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
		// page = armnetwork.VirtualApplianceSitesClientListResponse{
		// 	VirtualApplianceSiteListResult: armnetwork.VirtualApplianceSiteListResult{
		// 		Value: []*armnetwork.VirtualApplianceSite{
		// 			{
		// 				Name: to.Ptr("site1"),
		// 				Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva/virtualApplianceSites/site1"),
		// 				Properties: &armnetwork.VirtualApplianceSiteProperties{
		// 					AddressPrefix: to.Ptr("192.168.1.0/24"),
		// 					O365Policy: &armnetwork.Office365PolicyProperties{
		// 						BreakOutCategories: &armnetwork.BreakOutCategoryPolicies{
		// 							Default: to.Ptr(true),
		// 							Allow: to.Ptr(true),
		// 							Optimize: to.Ptr(true),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

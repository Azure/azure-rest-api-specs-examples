package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkVirtualApplianceSitePut.json
func ExampleVirtualApplianceSitesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualApplianceSitesClient().BeginCreateOrUpdate(ctx, "rg1", "nva", "site1", armnetwork.VirtualApplianceSite{
		Properties: &armnetwork.VirtualApplianceSiteProperties{
			AddressPrefix: to.Ptr("192.168.1.0/24"),
			O365Policy: &armnetwork.Office365PolicyProperties{
				BreakOutCategories: &armnetwork.BreakOutCategoryPolicies{
					Default:  to.Ptr(true),
					Allow:    to.Ptr(true),
					Optimize: to.Ptr(true),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualApplianceSite = armnetwork.VirtualApplianceSite{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva/virtualApplianceSites/site1"),
	// 	Name: to.Ptr("site1"),
	// 	Etag: to.Ptr("W/\"72090554-7e3b-43f2-80ad-99a9020dcb11\""),
	// 	Properties: &armnetwork.VirtualApplianceSiteProperties{
	// 		AddressPrefix: to.Ptr("192.168.1.0/24"),
	// 		O365Policy: &armnetwork.Office365PolicyProperties{
	// 			BreakOutCategories: &armnetwork.BreakOutCategoryPolicies{
	// 				Default: to.Ptr(true),
	// 				Allow: to.Ptr(true),
	// 				Optimize: to.Ptr(true),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

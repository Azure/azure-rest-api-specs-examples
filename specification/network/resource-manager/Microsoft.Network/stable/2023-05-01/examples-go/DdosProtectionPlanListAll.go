package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/DdosProtectionPlanListAll.json
func ExampleDdosProtectionPlansClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDdosProtectionPlansClient().NewListPager(nil)
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
		// page.DdosProtectionPlanListResult = armnetwork.DdosProtectionPlanListResult{
		// 	Value: []*armnetwork.DdosProtectionPlan{
		// 		{
		// 			Name: to.Ptr("plan1"),
		// 			Type: to.Ptr("Microsoft.Network/ddosProtectionPlans"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ddosProtectionPlans/plan1"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.DdosProtectionPlanPropertiesFormat{
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddresses: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-pip"),
		// 				}},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				VirtualNetworks: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet1"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("plan2"),
		// 			Type: to.Ptr("Microsoft.Network/ddosProtectionPlans"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ddosProtectionPlans/plan2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.DdosProtectionPlanPropertiesFormat{
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddresses: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-pip2"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-pip3"),
		// 				}},
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				VirtualNetworks: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet2"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet3"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

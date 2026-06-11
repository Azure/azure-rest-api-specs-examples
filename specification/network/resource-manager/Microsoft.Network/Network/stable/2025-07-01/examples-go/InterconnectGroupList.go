package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/InterconnectGroupList.json
func ExampleInterconnectGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInterconnectGroupsClient().NewListPager("rg1", nil)
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
		// page = armnetwork.InterconnectGroupsClientListResponse{
		// 	InterconnectGroupListResult: armnetwork.InterconnectGroupListResult{
		// 		Value: []*armnetwork.InterconnectGroup{
		// 			{
		// 				Name: to.Ptr("test-ig"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/interconnectGroups/test-ig"),
		// 				Type: to.Ptr("Microsoft.Network/interconnectGroups"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetwork.InterconnectGroupPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Scope: to.Ptr(armnetwork.InterconnectGroupScopeInfiniBand),
		// 					SubgroupProfile: &armnetwork.SubgroupProfile{
		// 						VMSize: to.Ptr("Standard_ND128isr_NDR_GB200_v6"),
		// 						Scope: to.Ptr(armnetwork.SubgroupProfileScopeVerticalConnect),
		// 						Size: to.Ptr[int32](18),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

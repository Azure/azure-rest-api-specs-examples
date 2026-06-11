package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/InterconnectGroupGet.json
func ExampleInterconnectGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInterconnectGroupsClient().Get(ctx, "rg1", "test-ig", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.InterconnectGroupsClientGetResponse{
	// 	InterconnectGroup: armnetwork.InterconnectGroup{
	// 		Name: to.Ptr("test-ig"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/interconnectGroups/test-ig"),
	// 		Type: to.Ptr("Microsoft.Network/interconnectGroups"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetwork.InterconnectGroupPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Scope: to.Ptr(armnetwork.InterconnectGroupScopeInfiniBand),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			SubgroupProfile: &armnetwork.SubgroupProfile{
	// 				VMSize: to.Ptr("Standard_ND128isr_NDR_GB200_v6"),
	// 				Scope: to.Ptr(armnetwork.SubgroupProfileScopeVerticalConnect),
	// 				Size: to.Ptr[int32](18),
	// 			},
	// 			Subgroups: []*armnetwork.Subgroup{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/interconnectGroups/test-ig/subgroups/subgroup0"),
	// 					Name: to.Ptr("subgroup0"),
	// 					Properties: &armnetwork.SubgroupProperties{
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						InternalSubgroupID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 						InterconnectBlock: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/interconnectBlocks/test-block"),
	// 						},
	// 						VirtualMachines: []*armnetwork.SubResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

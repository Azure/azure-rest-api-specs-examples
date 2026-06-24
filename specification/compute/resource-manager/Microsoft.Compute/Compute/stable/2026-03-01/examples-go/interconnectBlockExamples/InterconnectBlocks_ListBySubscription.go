package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/interconnectBlockExamples/InterconnectBlocks_ListBySubscription.json
func ExampleInterconnectBlocksClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInterconnectBlocksClient().NewListBySubscriptionPager(nil)
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
		// page = armcompute.InterconnectBlocksClientListBySubscriptionResponse{
		// 	InterconnectBlockListResult: armcompute.InterconnectBlockListResult{
		// 		Value: []*armcompute.InterconnectBlock{
		// 			{
		// 				Name: to.Ptr("myInterconnectBlock1"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup1/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock1"),
		// 				Type: to.Ptr("Microsoft.Compute/interconnectBlocks"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"department": to.Ptr("HR"),
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Standard_ND128isr_GB300_v6"),
		// 					Capacity: to.Ptr[int64](18),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1"),
		// 				},
		// 				Properties: &armcompute.InterconnectBlockProperties{
		// 					InterconnectGroup: &armcompute.APIEntityReference{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup1/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup1"),
		// 					},
		// 					InterconnectBlockID: to.Ptr("{GUID}"),
		// 					ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-01T01:02:38.3138469+00:00"); return t}()),
		// 					VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup1/providers/Microsoft.Compute/virtualMachines/myVM1"),
		// 						},
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup1/providers/Microsoft.Compute/virtualMachines/myVM2"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-01T01:02:38.3138469+00:00"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myInterconnectBlock2"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup2/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock2"),
		// 				Type: to.Ptr("Microsoft.Compute/interconnectBlocks"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"department": to.Ptr("HR"),
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Standard_ND128isr_GB300_v6"),
		// 					Capacity: to.Ptr[int64](18),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1"),
		// 				},
		// 				Properties: &armcompute.InterconnectBlockProperties{
		// 					InterconnectGroup: &armcompute.APIEntityReference{
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup2/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup2"),
		// 					},
		// 					InterconnectBlockID: to.Ptr("{GUID}"),
		// 					ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-01T01:02:38.3138469+00:00"); return t}()),
		// 					VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup2/providers/Microsoft.Compute/virtualMachines/myVM3"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-01T01:02:38.3138469+00:00"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

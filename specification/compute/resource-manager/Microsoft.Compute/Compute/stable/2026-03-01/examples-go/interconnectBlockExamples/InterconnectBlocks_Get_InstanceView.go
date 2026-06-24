package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/interconnectBlockExamples/InterconnectBlocks_Get_InstanceView.json
func ExampleInterconnectBlocksClient_Get_getAnInterconnectBlockWithInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInterconnectBlocksClient().Get(ctx, "myResourceGroup", "myInterconnectBlock", &armcompute.InterconnectBlocksClientGetOptions{
		Expand: to.Ptr(armcompute.InterconnectBlockExpandTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.InterconnectBlocksClientGetResponse{
	// 	InterconnectBlock: armcompute.InterconnectBlock{
	// 		Name: to.Ptr("myInterconnectBlock"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock"),
	// 		Type: to.Ptr("Microsoft.Compute/interconnectBlocks"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("HR"),
	// 		},
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("Standard_ND128isr_GB300_v6"),
	// 			Capacity: to.Ptr[int64](18),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 		Properties: &armcompute.InterconnectBlockProperties{
	// 			InterconnectGroup: &armcompute.APIEntityReference{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup"),
	// 			},
	// 			InterconnectBlockID: to.Ptr("{GUID}"),
	// 			ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-01T01:02:38.3138469+00:00"); return t}()),
	// 			VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM1"),
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM2"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			InstanceView: &armcompute.InterconnectBlockInstanceView{
	// 				CurrentCapacity: to.Ptr[int32](18),
	// 				Statuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("ProvisioningState/succeeded"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					},
	// 				},
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-01T01:02:38.3138469+00:00"); return t}()),
	// 		},
	// 	},
	// }
}

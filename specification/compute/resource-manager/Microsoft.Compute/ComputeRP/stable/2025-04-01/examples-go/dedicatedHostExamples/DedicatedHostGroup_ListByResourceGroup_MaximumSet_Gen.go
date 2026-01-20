package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/dedicatedHostExamples/DedicatedHostGroup_ListByResourceGroup_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_NewListByResourceGroupPager_dedicatedHostGroupListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostGroupsClient().NewListByResourceGroupPager("rgcompute", nil)
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
		// page.DedicatedHostGroupListResult = armcompute.DedicatedHostGroupListResult{
		// 	Value: []*armcompute.DedicatedHostGroup{
		// 		{
		// 			Name: to.Ptr("myDedicatedHostGroup"),
		// 			Type: to.Ptr("aaaa"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcompute.DedicatedHostGroupProperties{
		// 				Hosts: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("aaaa"),
		// 				}},
		// 				InstanceView: &armcompute.DedicatedHostGroupInstanceView{
		// 					Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
		// 						{
		// 							AssetID: to.Ptr("aaaa"),
		// 							AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
		// 								AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
		// 									{
		// 										Count: to.Ptr[float64](26),
		// 										VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 								}},
		// 							},
		// 							Statuses: []*armcompute.InstanceViewStatus{
		// 								{
		// 									Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 									DisplayStatus: to.Ptr("aaaaaa"),
		// 									Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 									Message: to.Ptr("a"),
		// 									Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 							}},
		// 							Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 					}},
		// 				},
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				SupportAutomaticPlacement: to.Ptr(true),
		// 			},
		// 			Zones: []*string{
		// 				to.Ptr("1")},
		// 		}},
		// 	}
	}
}

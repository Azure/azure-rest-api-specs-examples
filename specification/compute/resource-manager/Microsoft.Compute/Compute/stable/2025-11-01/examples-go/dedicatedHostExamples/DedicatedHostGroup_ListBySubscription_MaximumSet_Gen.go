package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/dedicatedHostExamples/DedicatedHostGroup_ListBySubscription_MaximumSet_Gen.json
func ExampleDedicatedHostGroupsClient_NewListBySubscriptionPager_dedicatedHostGroupListBySubscriptionMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostGroupsClient().NewListBySubscriptionPager(nil)
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
		// page = armcompute.DedicatedHostGroupsClientListBySubscriptionResponse{
		// 	DedicatedHostGroupListResult: armcompute.DedicatedHostGroupListResult{
		// 		Value: []*armcompute.DedicatedHostGroup{
		// 			{
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1"),
		// 				},
		// 				Properties: &armcompute.DedicatedHostGroupProperties{
		// 					PlatformFaultDomainCount: to.Ptr[int32](3),
		// 					SupportAutomaticPlacement: to.Ptr(true),
		// 					Hosts: []*armcompute.SubResourceReadOnly{
		// 						{
		// 							ID: to.Ptr("aaaa"),
		// 						},
		// 					},
		// 					InstanceView: &armcompute.DedicatedHostGroupInstanceView{
		// 						Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
		// 							{
		// 								Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 								AssetID: to.Ptr("aaaa"),
		// 								AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
		// 									AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
		// 										{
		// 											VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 											Count: to.Ptr[float64](26),
		// 										},
		// 									},
		// 								},
		// 								Statuses: []*armcompute.InstanceViewStatus{
		// 									{
		// 										Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										DisplayStatus: to.Ptr("aaaaaa"),
		// 										Message: to.Ptr("a"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
		// 				Name: to.Ptr("myDedicatedHostGroup"),
		// 				Type: to.Ptr("aaaa"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("a://example.com/aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 	},
		// }
	}
}

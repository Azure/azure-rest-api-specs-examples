package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/dedicatedHostExamples/DedicatedHost_ListByHostGroup_MaximumSet_Gen.json
func ExampleDedicatedHostsClient_NewListByHostGroupPager_dedicatedHostListByHostGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedHostsClient().NewListByHostGroupPager("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
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
		// page.DedicatedHostListResult = armcompute.DedicatedHostListResult{
		// 	Value: []*armcompute.DedicatedHost{
		// 		{
		// 			Name: to.Ptr("myDedicatedHost"),
		// 			Type: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myDedicatedHost"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcompute.DedicatedHostProperties{
		// 				AutoReplaceOnFailure: to.Ptr(true),
		// 				HostID: to.Ptr("{GUID}"),
		// 				InstanceView: &armcompute.DedicatedHostInstanceView{
		// 					AssetID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 					AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
		// 						AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
		// 							{
		// 								Count: to.Ptr[float64](26),
		// 								VMSize: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 						}},
		// 					},
		// 					Statuses: []*armcompute.InstanceViewStatus{
		// 						{
		// 							Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 							DisplayStatus: to.Ptr("aaaaaa"),
		// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 							Message: to.Ptr("a"),
		// 							Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 					}},
		// 				},
		// 				LicenseType: to.Ptr(armcompute.DedicatedHostLicenseTypesWindowsServerHybrid),
		// 				PlatformFaultDomain: to.Ptr[int32](1),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.526Z"); return t}()),
		// 				VirtualMachines: []*armcompute.SubResourceReadOnly{
		// 					{
		// 						ID: to.Ptr("aaaa"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("DSv3-Type1"),
		// 				Capacity: to.Ptr[int64](7),
		// 				Tier: to.Ptr("aaa"),
		// 			},
		// 	}},
		// }
	}
}

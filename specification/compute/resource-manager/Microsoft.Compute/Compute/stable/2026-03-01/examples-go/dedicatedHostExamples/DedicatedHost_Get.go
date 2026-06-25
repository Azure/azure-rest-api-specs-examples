package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/dedicatedHostExamples/DedicatedHost_Get.json
func ExampleDedicatedHostsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostsClient().Get(ctx, "myResourceGroup", "myDedicatedHostGroup", "myHost", &armcompute.DedicatedHostsClientGetOptions{
		Expand: to.Ptr(armcompute.InstanceViewTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DedicatedHostsClientGetResponse{
	// 	DedicatedHost: armcompute.DedicatedHost{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup/hosts/myHost"),
	// 		Properties: &armcompute.DedicatedHostProperties{
	// 			PlatformFaultDomain: to.Ptr[int32](1),
	// 			AutoReplaceOnFailure: to.Ptr(true),
	// 			HostID: to.Ptr("{GUID}"),
	// 			ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-27T01:02:38.3138469+00:00"); return t}()),
	// 			VirtualMachines: []*armcompute.SubResourceReadOnly{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			InstanceView: &armcompute.DedicatedHostInstanceView{
	// 				AssetID: to.Ptr("eb3f58b8-b4e8-4882-b69f-301a01812407"),
	// 				AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
	// 					AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
	// 						{
	// 							VMSize: to.Ptr("Standard_A1"),
	// 							Count: to.Ptr[float64](10),
	// 						},
	// 					},
	// 				},
	// 				Statuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("ProvisioningState/succeeded"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 					},
	// 					{
	// 						Code: to.Ptr("HealthState/available"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("Host available"),
	// 					},
	// 				},
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-27T01:02:38.3138469+00:00"); return t}()),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("HR"),
	// 		},
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("DSv3-Type1"),
	// 		},
	// 		Name: to.Ptr("myHost"),
	// 	},
	// }
}

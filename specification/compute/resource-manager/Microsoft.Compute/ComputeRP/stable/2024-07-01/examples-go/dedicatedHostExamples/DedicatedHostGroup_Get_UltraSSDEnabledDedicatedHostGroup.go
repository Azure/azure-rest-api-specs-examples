package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/dedicatedHostExamples/DedicatedHostGroup_Get_UltraSSDEnabledDedicatedHostGroup.json
func ExampleDedicatedHostGroupsClient_Get_createAnUltraSsdEnabledDedicatedHostGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostGroupsClient().Get(ctx, "myResourceGroup", "myDedicatedHostGroup", &armcompute.DedicatedHostGroupsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedHostGroup = armcompute.DedicatedHostGroup{
	// 	Name: to.Ptr("myDedicatedHostGroup"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"{tagName}": to.Ptr("{tagValue}"),
	// 	},
	// 	Properties: &armcompute.DedicatedHostGroupProperties{
	// 		AdditionalCapabilities: &armcompute.DedicatedHostGroupPropertiesAdditionalCapabilities{
	// 			UltraSSDEnabled: to.Ptr(true),
	// 		},
	// 		Hosts: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/myDedicatedHostGroup/myHostGroup/Hosts/myHost"),
	// 		}},
	// 		InstanceView: &armcompute.DedicatedHostGroupInstanceView{
	// 			Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
	// 				{
	// 					AssetID: to.Ptr("eb3f58b8-b4e8-4882-b69f-301a01812407"),
	// 					AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
	// 						AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
	// 							{
	// 								Count: to.Ptr[float64](10),
	// 								VMSize: to.Ptr("Standard_A1"),
	// 						}},
	// 					},
	// 					Statuses: []*armcompute.InstanceViewStatus{
	// 						{
	// 							Code: to.Ptr("ProvisioningState/succeeded"),
	// 							DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						},
	// 						{
	// 							Code: to.Ptr("HealthState/available"),
	// 							DisplayStatus: to.Ptr("Host available"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					}},
	// 					Name: to.Ptr("myHost"),
	// 			}},
	// 		},
	// 		PlatformFaultDomainCount: to.Ptr[int32](3),
	// 		SupportAutomaticPlacement: to.Ptr(true),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("3")},
	// 	}
}

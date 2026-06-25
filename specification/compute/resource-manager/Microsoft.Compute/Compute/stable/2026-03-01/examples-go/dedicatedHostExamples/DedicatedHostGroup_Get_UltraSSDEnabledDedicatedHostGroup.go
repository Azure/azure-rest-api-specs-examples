package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/dedicatedHostExamples/DedicatedHostGroup_Get_UltraSSDEnabledDedicatedHostGroup.json
func ExampleDedicatedHostGroupsClient_Get_createAnUltraSsdEnabledDedicatedHostGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedHostGroupsClient().Get(ctx, "myResourceGroup", "myDedicatedHostGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DedicatedHostGroupsClientGetResponse{
	// 	DedicatedHostGroup: armcompute.DedicatedHostGroup{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/HostGroups/myDedicatedHostGroup"),
	// 		Properties: &armcompute.DedicatedHostGroupProperties{
	// 			PlatformFaultDomainCount: to.Ptr[int32](3),
	// 			Hosts: []*armcompute.SubResourceReadOnly{
	// 				{
	// 					ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/myDedicatedHostGroup/myHostGroup/Hosts/myHost"),
	// 				},
	// 			},
	// 			SupportAutomaticPlacement: to.Ptr(true),
	// 			AdditionalCapabilities: &armcompute.DedicatedHostGroupPropertiesAdditionalCapabilities{
	// 				UltraSSDEnabled: to.Ptr(true),
	// 			},
	// 			InstanceView: &armcompute.DedicatedHostGroupInstanceView{
	// 				Hosts: []*armcompute.DedicatedHostInstanceViewWithName{
	// 					{
	// 						Name: to.Ptr("myHost"),
	// 						AssetID: to.Ptr("eb3f58b8-b4e8-4882-b69f-301a01812407"),
	// 						AvailableCapacity: &armcompute.DedicatedHostAvailableCapacity{
	// 							AllocatableVMs: []*armcompute.DedicatedHostAllocatableVM{
	// 								{
	// 									VMSize: to.Ptr("Standard_A1"),
	// 									Count: to.Ptr[float64](10),
	// 								},
	// 							},
	// 						},
	// 						Statuses: []*armcompute.InstanceViewStatus{
	// 							{
	// 								Code: to.Ptr("ProvisioningState/succeeded"),
	// 								Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 								DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 							},
	// 							{
	// 								Code: to.Ptr("HealthState/available"),
	// 								Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 								DisplayStatus: to.Ptr("Host available"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"{tagName}": to.Ptr("{tagValue}"),
	// 		},
	// 		Name: to.Ptr("myDedicatedHostGroup"),
	// 		Zones: []*string{
	// 			to.Ptr("3"),
	// 		},
	// 	},
	// }
}

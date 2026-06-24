package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/availabilitySetExamples/AvailabilitySet_Get_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_Get_availabilitySetGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilitySetsClient().Get(ctx, "rgcompute", "aaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.AvailabilitySetsClientGetResponse{
	// 	AvailabilitySet: armcompute.AvailabilitySet{
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armcompute.AvailabilitySetProperties{
	// 			PlatformFaultDomainCount: to.Ptr[int32](2),
	// 			PlatformUpdateDomainCount: to.Ptr[int32](20),
	// 			VirtualMachineScaleSetMigrationInfo: &armcompute.VirtualMachineScaleSetMigrationInfo{
	// 				DefaultVirtualMachineScaleSetInfo: &armcompute.DefaultVirtualMachineScaleSetInfo{
	// 					DefaultVirtualMachineScaleSet: &armcompute.SubResource{
	// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}"),
	// 					},
	// 					ConstrainedMaximumCapacity: to.Ptr(false),
	// 				},
	// 				MigrateToVirtualMachineScaleSet: &armcompute.SubResource{
	// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/{vmss-name}"),
	// 				},
	// 			},
	// 			VirtualMachines: []*armcompute.SubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 				},
	// 			},
	// 			ProximityPlacementGroup: &armcompute.SubResource{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 			},
	// 			ScheduledEventsPolicy: &armcompute.ScheduledEventsPolicy{
	// 				ScheduledEventsAdditionalPublishingTargets: &armcompute.ScheduledEventsAdditionalPublishingTargets{
	// 					EventGridAndResourceGraph: &armcompute.EventGridAndResourceGraph{
	// 						Enable: to.Ptr(true),
	// 						ScheduledEventsAPIVersion: to.Ptr("2020-07-01"),
	// 					},
	// 				},
	// 				UserInitiatedRedeploy: &armcompute.UserInitiatedRedeploy{
	// 					AutomaticallyApprove: to.Ptr(true),
	// 				},
	// 				UserInitiatedReboot: &armcompute.UserInitiatedReboot{
	// 					AutomaticallyApprove: to.Ptr(true),
	// 				},
	// 				AllInstancesDown: &armcompute.AllInstancesDown{
	// 					AutomaticallyApprove: to.Ptr(true),
	// 				},
	// 			},
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					DisplayStatus: to.Ptr("aaaaaa"),
	// 					Message: to.Ptr("a"),
	// 					Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 				},
	// 			},
	// 		},
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("Classic"),
	// 			Tier: to.Ptr("aaaaaaaaaaaaaa"),
	// 			Capacity: to.Ptr[int64](29),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet"),
	// 		Name: to.Ptr("myAvailabilitySet"),
	// 		Type: to.Ptr("Microsoft.Compute/availabilitySets"),
	// 		Tags: map[string]*string{
	// 			"key9626": to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 			"key2505": to.Ptr("aa"),
	// 		},
	// 	},
	// }
}

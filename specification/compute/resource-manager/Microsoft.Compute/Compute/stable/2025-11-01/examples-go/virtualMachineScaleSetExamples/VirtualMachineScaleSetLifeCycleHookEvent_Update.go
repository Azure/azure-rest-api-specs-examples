package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetLifeCycleHookEvent_Update.json
func ExampleVirtualMachineScaleSetLifeCycleHookEventsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("2167b012-c9f9-4b04-83b2-0ff304e7d51d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetLifeCycleHookEventsClient().Update(ctx, "RG01", "VMSS01", "445c0a08-cfc5-4ef6-bb89-fe77c5178628", armcompute.VMScaleSetLifecycleHookEventUpdate{
		Properties: &armcompute.VMScaleSetLifecycleHookEventProperties{
			WaitUntil: to.Ptr("2025-05-08T11:17:55.6844555+00:00"),
			TargetResources: []*armcompute.VMScaleSetLifecycleHookEventTargetResource{
				{
					Resource: &armcompute.APIEntityReference{
						ID: to.Ptr("/subscriptions/2167b012-c9f9-4b04-83b2-0ff304e7d51d/resourceGroups/RG01/providers/Microsoft.Compute/virtualMachineScaleSets/VMSS01/virtualMachines/2"),
					},
					ActionState: to.Ptr(armcompute.LifecycleHookActionStateApproved),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineScaleSetLifeCycleHookEventsClientUpdateResponse{
	// 	VMScaleSetLifecycleHookEvent: &armcompute.VMScaleSetLifecycleHookEvent{
	// 		Name: to.Ptr("445c0a08-cfc5-4ef6-bb89-fe77c5178628"),
	// 		ID: to.Ptr("/subscriptions/445c0a08-cfc5-4ef6-bb89-fe77c5178628/resourceGroups/RG01/providers/Microsoft.Compute/virtualMachineScaleSets/VMSS01/lifecycleHookEvents/445c0a08-cfc5-4ef6-bb89-fe77c5178628"),
	// 		Properties: &armcompute.VMScaleSetLifecycleHookEventProperties{
	// 			Type: to.Ptr(armcompute.VMScaleSetLifecycleHookEventTypeUpgradeAutoOSRollingBatchStarting),
	// 			WaitUntil: to.Ptr("2025-05-08T11:17:55.6844555+00:00"),
	// 			MaxWaitUntil: to.Ptr("2025-05-08T13:17:55.6844555+00:00"),
	// 			TimeCreated: to.Ptr("2025-05-08T07:17:55.6844555+00:00"),
	// 			DefaultAction: to.Ptr(armcompute.LifecycleHookActionApprove),
	// 			TargetResources: []*armcompute.VMScaleSetLifecycleHookEventTargetResource{
	// 				{
	// 					Resource: &armcompute.APIEntityReference{
	// 						ID: to.Ptr("/subscriptions/2167b012-c9f9-4b04-83b2-0ff304e7d51d/resourceGroups/RG01/providers/Microsoft.Compute/virtualMachineScaleSets/VMSS01/virtualMachines/2"),
	// 					},
	// 					ActionState: to.Ptr(armcompute.LifecycleHookActionStateApproved),
	// 				},
	// 				{
	// 					Resource: &armcompute.APIEntityReference{
	// 						ID: to.Ptr("/subscriptions/2167b012-c9f9-4b04-83b2-0ff304e7d51d/resourceGroups/RG01/providers/Microsoft.Compute/virtualMachineScaleSets/VMSS01/virtualMachines/3"),
	// 					},
	// 					ActionState: to.Ptr(armcompute.LifecycleHookActionStateWaiting),
	// 				},
	// 			},
	// 			State: to.Ptr(armcompute.VMScaleSetLifecycleHookEventStateActive),
	// 		},
	// 	},
	// }
}

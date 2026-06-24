package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetLifeCycleHookEvent_Get.json
func ExampleVirtualMachineScaleSetLifeCycleHookEventsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("2167b012-c9f9-4b04-83b2-0ff304e7d51d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetLifeCycleHookEventsClient().Get(ctx, "RG01", "VMSS01", "2e2e3046-f85f-4966-8fd2-5fd7bf6ea717", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineScaleSetLifeCycleHookEventsClientGetResponse{
	// 	VMScaleSetLifecycleHookEvent: armcompute.VMScaleSetLifecycleHookEvent{
	// 		Name: to.Ptr("2e2e3046-f85f-4966-8fd2-5fd7bf6ea717"),
	// 		ID: to.Ptr("/subscriptions/2167b012-c9f9-4b04-83b2-0ff304e7d51d/resourceGroups/RG01/providers/Microsoft.Compute/virtualMachineScaleSets/VMSS01/lifecycleHookEvents/2e2e3046-f85f-4966-8fd2-5fd7bf6ea717"),
	// 		Properties: &armcompute.VMScaleSetLifecycleHookEventProperties{
	// 			Type: to.Ptr(armcompute.VMScaleSetLifecycleHookEventTypeUpgradeAutoOSRollingBatchStarting),
	// 			WaitUntil: to.Ptr("2025-05-08T09:17:55.6844555+00:00"),
	// 			MaxWaitUntil: to.Ptr("2025-05-08T12:17:55.6844555+00:00"),
	// 			TimeCreated: to.Ptr("2025-05-08T06:17:55.6844555+00:00"),
	// 			DefaultAction: to.Ptr(armcompute.LifecycleHookActionApprove),
	// 			TargetResources: []*armcompute.VMScaleSetLifecycleHookEventTargetResource{
	// 				{
	// 					Resource: &armcompute.APIEntityReference{
	// 						ID: to.Ptr("/subscriptions/2167b012-c9f9-4b04-83b2-0ff304e7d51d/resourceGroups/RG01/providers/Microsoft.Compute/virtualMachineScaleSets/VMSS01/virtualMachines/0"),
	// 					},
	// 					ActionState: to.Ptr(armcompute.LifecycleHookActionStateApproved),
	// 				},
	// 				{
	// 					Resource: &armcompute.APIEntityReference{
	// 						ID: to.Ptr("/subscriptions/2167b012-c9f9-4b04-83b2-0ff304e7d51d/resourceGroups/RG01/providers/Microsoft.Compute/virtualMachineScaleSets/VMSS01/virtualMachines/1"),
	// 					},
	// 					ActionState: to.Ptr(armcompute.LifecycleHookActionStateApproved),
	// 				},
	// 			},
	// 			State: to.Ptr(armcompute.VMScaleSetLifecycleHookEventStateCompleted),
	// 		},
	// 	},
	// }
}

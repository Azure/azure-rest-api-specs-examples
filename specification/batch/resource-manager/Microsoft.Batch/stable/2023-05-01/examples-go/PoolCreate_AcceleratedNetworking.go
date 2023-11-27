package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolCreate_AcceleratedNetworking.json
func ExamplePoolClient_Create_createPoolAcceleratedNetworking() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolClient().Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					ImageReference: &armbatch.ImageReference{
						Offer:     to.Ptr("WindowsServer"),
						Publisher: to.Ptr("MicrosoftWindowsServer"),
						SKU:       to.Ptr("2016-datacenter-smalldisk"),
						Version:   to.Ptr("latest"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
				},
			},
			NetworkConfiguration: &armbatch.NetworkConfiguration{
				EnableAcceleratedNetworking: to.Ptr(true),
				SubnetID:                    to.Ptr("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
			},
			ScaleSettings: &armbatch.ScaleSettings{
				FixedScale: &armbatch.FixedScaleSettings{
					TargetDedicatedNodes:   to.Ptr[int32](1),
					TargetLowPriorityNodes: to.Ptr[int32](0),
				},
			},
			VMSize: to.Ptr("STANDARD_D1_V2"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Pool = armbatch.Pool{
	// 	Name: to.Ptr("testpool"),
	// 	Type: to.Ptr("Microsoft.Batch/batchAccounts/pools"),
	// 	Etag: to.Ptr("W/\"0x8DB46CB72A227E2\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool"),
	// 	Properties: &armbatch.PoolProperties{
	// 		AllocationState: to.Ptr(armbatch.AllocationStateSteady),
	// 		AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-27T03:00:34.064Z"); return t}()),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-27T02:59:41.859Z"); return t}()),
	// 		CurrentDedicatedNodes: to.Ptr[int32](1),
	// 		CurrentLowPriorityNodes: to.Ptr[int32](0),
	// 		CurrentNodeCommunicationMode: to.Ptr(armbatch.NodeCommunicationModeClassic),
	// 		DeploymentConfiguration: &armbatch.DeploymentConfiguration{
	// 			VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
	// 				ImageReference: &armbatch.ImageReference{
	// 					Offer: to.Ptr("WindowsServer"),
	// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 					SKU: to.Ptr("2016-datacenter-smalldisk"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
	// 			},
	// 		},
	// 		InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateDisabled),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-27T02:59:41.859Z"); return t}()),
	// 		NetworkConfiguration: &armbatch.NetworkConfiguration{
	// 			DynamicVNetAssignmentScope: to.Ptr(armbatch.DynamicVNetAssignmentScopeNone),
	// 			EnableAcceleratedNetworking: to.Ptr(true),
	// 			SubnetID: to.Ptr("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
	// 		},
	// 		ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
	// 		ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-27T02:59:41.859Z"); return t}()),
	// 		ResizeOperationStatus: &armbatch.ResizeOperationStatus{
	// 			NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionRequeue),
	// 			ResizeTimeout: to.Ptr("PT15M"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-27T02:59:41.859Z"); return t}()),
	// 			TargetDedicatedNodes: to.Ptr[int32](1),
	// 		},
	// 		ScaleSettings: &armbatch.ScaleSettings{
	// 			FixedScale: &armbatch.FixedScaleSettings{
	// 				ResizeTimeout: to.Ptr("PT15M"),
	// 				TargetDedicatedNodes: to.Ptr[int32](1),
	// 				TargetLowPriorityNodes: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
	// 			NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypeSpread),
	// 		},
	// 		TaskSlotsPerNode: to.Ptr[int32](1),
	// 		VMSize: to.Ptr("STANDARD_D1_V2"),
	// 	},
	// }
}

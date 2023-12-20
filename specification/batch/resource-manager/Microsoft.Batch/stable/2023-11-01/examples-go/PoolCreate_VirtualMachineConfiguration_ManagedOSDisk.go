package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolCreate_VirtualMachineConfiguration_ManagedOSDisk.json
func ExamplePoolClient_Create_createPoolVirtualMachineConfigurationOsDisk() {
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
						Offer:     to.Ptr("windowsserver"),
						Publisher: to.Ptr("microsoftwindowsserver"),
						SKU:       to.Ptr("2022-datacenter-smalldisk"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
					OSDisk: &armbatch.OSDisk{
						Caching:    to.Ptr(armbatch.CachingTypeReadWrite),
						DiskSizeGB: to.Ptr[int32](100),
						ManagedDisk: &armbatch.ManagedDisk{
							StorageAccountType: to.Ptr(armbatch.StorageAccountTypeStandardSSDLRS),
						},
						WriteAcceleratorEnabled: to.Ptr(false),
					},
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				FixedScale: &armbatch.FixedScaleSettings{
					TargetDedicatedNodes:   to.Ptr[int32](1),
					TargetLowPriorityNodes: to.Ptr[int32](0),
				},
			},
			VMSize: to.Ptr("Standard_d2s_v3"),
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
	// 	Etag: to.Ptr("W/\"0x8DB51F396992B8D\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool"),
	// 	Properties: &armbatch.PoolProperties{
	// 		AllocationState: to.Ptr(armbatch.AllocationStateResizing),
	// 		AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-24T02:12:27.752Z"); return t}()),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-24T02:12:27.752Z"); return t}()),
	// 		CurrentDedicatedNodes: to.Ptr[int32](0),
	// 		CurrentLowPriorityNodes: to.Ptr[int32](0),
	// 		DeploymentConfiguration: &armbatch.DeploymentConfiguration{
	// 			VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
	// 				ImageReference: &armbatch.ImageReference{
	// 					Offer: to.Ptr("windowsserver"),
	// 					Publisher: to.Ptr("microsoftwindowsserver"),
	// 					SKU: to.Ptr("2022-datacenter-smalldisk"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
	// 				OSDisk: &armbatch.OSDisk{
	// 					Caching: to.Ptr(armbatch.CachingTypeReadWrite),
	// 					DiskSizeGB: to.Ptr[int32](100),
	// 					ManagedDisk: &armbatch.ManagedDisk{
	// 						StorageAccountType: to.Ptr(armbatch.StorageAccountTypeStandardSSDLRS),
	// 					},
	// 					WriteAcceleratorEnabled: to.Ptr(false),
	// 				},
	// 			},
	// 		},
	// 		InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateDisabled),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-24T02:12:27.752Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
	// 		ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-24T02:12:27.752Z"); return t}()),
	// 		ResizeOperationStatus: &armbatch.ResizeOperationStatus{
	// 			NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionRequeue),
	// 			ResizeTimeout: to.Ptr("PT15M"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-24T02:12:27.752Z"); return t}()),
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
	// 		VMSize: to.Ptr("STANDARD_D2_V3"),
	// 	},
	// }
}

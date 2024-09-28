package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PoolCreate_UpgradePolicy.json
func ExamplePoolClient_Create_createPoolUpgradePolicy() {
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
						SKU:       to.Ptr("2019-datacenter-smalldisk"),
						Version:   to.Ptr("latest"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
					NodePlacementConfiguration: &armbatch.NodePlacementConfiguration{
						Policy: to.Ptr(armbatch.NodePlacementPolicyTypeZonal),
					},
					WindowsConfiguration: &armbatch.WindowsConfiguration{
						EnableAutomaticUpdates: to.Ptr(false),
					},
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				FixedScale: &armbatch.FixedScaleSettings{
					TargetDedicatedNodes:   to.Ptr[int32](2),
					TargetLowPriorityNodes: to.Ptr[int32](0),
				},
			},
			UpgradePolicy: &armbatch.UpgradePolicy{
				AutomaticOSUpgradePolicy: &armbatch.AutomaticOSUpgradePolicy{
					DisableAutomaticRollback: to.Ptr(true),
					EnableAutomaticOSUpgrade: to.Ptr(true),
					OSRollingUpgradeDeferral: to.Ptr(true),
					UseRollingUpgradePolicy:  to.Ptr(true),
				},
				Mode: to.Ptr(armbatch.UpgradeModeAutomatic),
				RollingUpgradePolicy: &armbatch.RollingUpgradePolicy{
					EnableCrossZoneUpgrade:                to.Ptr(true),
					MaxBatchInstancePercent:               to.Ptr[int32](20),
					MaxUnhealthyInstancePercent:           to.Ptr[int32](20),
					MaxUnhealthyUpgradedInstancePercent:   to.Ptr[int32](20),
					PauseTimeBetweenBatches:               to.Ptr("PT0S"),
					PrioritizeUnhealthyInstances:          to.Ptr(false),
					RollbackFailedInstancesOnPolicyBreach: to.Ptr(false),
				},
			},
			VMSize: to.Ptr("Standard_d4s_v3"),
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
	// 	Etag: to.Ptr("W/\"0x8DB51E64D3C3B69\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool"),
	// 	Properties: &armbatch.PoolProperties{
	// 		AllocationState: to.Ptr(armbatch.AllocationStateResizing),
	// 		AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T06:16:44.237Z"); return t}()),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T06:16:44.237Z"); return t}()),
	// 		CurrentDedicatedNodes: to.Ptr[int32](0),
	// 		CurrentLowPriorityNodes: to.Ptr[int32](0),
	// 		DeploymentConfiguration: &armbatch.DeploymentConfiguration{
	// 			VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
	// 				ImageReference: &armbatch.ImageReference{
	// 					Offer: to.Ptr("WindowsServer"),
	// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 					SKU: to.Ptr("2019-datacenter-smalldisk"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
	// 				NodePlacementConfiguration: &armbatch.NodePlacementConfiguration{
	// 					Policy: to.Ptr(armbatch.NodePlacementPolicyTypeZonal),
	// 				},
	// 				WindowsConfiguration: &armbatch.WindowsConfiguration{
	// 					EnableAutomaticUpdates: to.Ptr(false),
	// 				},
	// 			},
	// 		},
	// 		InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateDisabled),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T06:16:44.237Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
	// 		ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T06:16:44.237Z"); return t}()),
	// 		ResizeOperationStatus: &armbatch.ResizeOperationStatus{
	// 			NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionRequeue),
	// 			ResizeTimeout: to.Ptr("PT15M"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T06:16:44.237Z"); return t}()),
	// 			TargetDedicatedNodes: to.Ptr[int32](2),
	// 		},
	// 		ScaleSettings: &armbatch.ScaleSettings{
	// 			FixedScale: &armbatch.FixedScaleSettings{
	// 				ResizeTimeout: to.Ptr("PT15M"),
	// 				TargetDedicatedNodes: to.Ptr[int32](2),
	// 				TargetLowPriorityNodes: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
	// 			NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypeSpread),
	// 		},
	// 		TaskSlotsPerNode: to.Ptr[int32](1),
	// 		UpgradePolicy: &armbatch.UpgradePolicy{
	// 			AutomaticOSUpgradePolicy: &armbatch.AutomaticOSUpgradePolicy{
	// 				DisableAutomaticRollback: to.Ptr(true),
	// 				EnableAutomaticOSUpgrade: to.Ptr(true),
	// 				OSRollingUpgradeDeferral: to.Ptr(true),
	// 				UseRollingUpgradePolicy: to.Ptr(true),
	// 			},
	// 			Mode: to.Ptr(armbatch.UpgradeModeAutomatic),
	// 			RollingUpgradePolicy: &armbatch.RollingUpgradePolicy{
	// 				EnableCrossZoneUpgrade: to.Ptr(true),
	// 				MaxBatchInstancePercent: to.Ptr[int32](20),
	// 				MaxUnhealthyInstancePercent: to.Ptr[int32](20),
	// 				MaxUnhealthyUpgradedInstancePercent: to.Ptr[int32](20),
	// 				PauseTimeBetweenBatches: to.Ptr("PT0S"),
	// 				PrioritizeUnhealthyInstances: to.Ptr(false),
	// 				RollbackFailedInstancesOnPolicyBreach: to.Ptr(false),
	// 			},
	// 		},
	// 		VMSize: to.Ptr("STANDARD_D4S_V3"),
	// 	},
	// }
}

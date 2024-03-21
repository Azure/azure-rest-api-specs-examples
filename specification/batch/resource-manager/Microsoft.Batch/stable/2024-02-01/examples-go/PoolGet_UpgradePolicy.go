package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/PoolGet_UpgradePolicy.json
func ExamplePoolClient_Get_getPoolUpgradePolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolClient().Get(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", nil)
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
	// 		AllocationState: to.Ptr(armbatch.AllocationStateSteady),
	// 		AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T06:10:31.467Z"); return t}()),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T06:09:38.317Z"); return t}()),
	// 		CurrentDedicatedNodes: to.Ptr[int32](2),
	// 		CurrentLowPriorityNodes: to.Ptr[int32](0),
	// 		CurrentNodeCommunicationMode: to.Ptr(armbatch.NodeCommunicationModeClassic),
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
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T06:09:38.317Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
	// 		ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T06:09:38.317Z"); return t}()),
	// 		ResizeOperationStatus: &armbatch.ResizeOperationStatus{
	// 			NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionRequeue),
	// 			ResizeTimeout: to.Ptr("PT15M"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-11T06:09:38.317Z"); return t}()),
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

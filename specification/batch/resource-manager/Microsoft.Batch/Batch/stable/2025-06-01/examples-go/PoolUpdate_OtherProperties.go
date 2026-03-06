package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/PoolUpdate_OtherProperties.json
func ExamplePoolClient_Update_updatePoolOtherProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolClient().Update(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			ApplicationPackages: []*armbatch.ApplicationPackageReference{
				{
					ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
				},
				{
					ID:      to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_5678"),
					Version: to.Ptr("1.0"),
				},
			},
			Metadata: []*armbatch.MetadataItem{
				{
					Name:  to.Ptr("key1"),
					Value: to.Ptr("value1"),
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
	// res = armbatch.PoolClientUpdateResponse{
	// 	ETag: "W/\"0x8D4EDFEBFADF4AB\""
	// 	Pool: &armbatch.Pool{
	// 		Name: to.Ptr("testpool"),
	// 		Type: to.Ptr("Microsoft.Batch/batchAccounts/pools"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool"),
	// 		Properties: &armbatch.PoolProperties{
	// 			AllocationState: to.Ptr(armbatch.AllocationStateResizing),
	// 			AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T10:22:55.9407275Z"); return t}()),
	// 			ApplicationPackages: []*armbatch.ApplicationPackageReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_5678"),
	// 					Version: to.Ptr("1.0"),
	// 				},
	// 			},
	// 			AutoScaleRun: &armbatch.AutoScaleRun{
	// 				EvaluationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T10:22:55.9407275Z"); return t}()),
	// 				Results: to.Ptr("$TargetDedicatedNodes=34;NodeDeallocationOption=requeue"),
	// 			},
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.9407275Z"); return t}()),
	// 			CurrentDedicatedNodes: to.Ptr[int32](12),
	// 			CurrentLowPriorityNodes: to.Ptr[int32](0),
	// 			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
	// 				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
	// 					ImageReference: &armbatch.ImageReference{
	// 						Offer: to.Ptr("WindowsServer"),
	// 						Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 						SKU: to.Ptr("2016-Datacenter-SmallDisk"),
	// 						Version: to.Ptr("latest"),
	// 					},
	// 					NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
	// 				},
	// 			},
	// 			InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateDisabled),
	// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T10:22:55.9407275Z"); return t}()),
	// 			Metadata: []*armbatch.MetadataItem{
	// 				{
	// 					Name: to.Ptr("key1"),
	// 					Value: to.Ptr("value1"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
	// 			ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.9407275Z"); return t}()),
	// 			ResizeOperationStatus: &armbatch.ResizeOperationStatus{
	// 				NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionTaskCompletion),
	// 				ResizeTimeout: to.Ptr("PT8M"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T10:22:55.9407275Z"); return t}()),
	// 				TargetDedicatedNodes: to.Ptr[int32](8),
	// 			},
	// 			ScaleSettings: &armbatch.ScaleSettings{
	// 				FixedScale: &armbatch.FixedScaleSettings{
	// 					NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionTaskCompletion),
	// 					ResizeTimeout: to.Ptr("PT8M"),
	// 					TargetDedicatedNodes: to.Ptr[int32](1),
	// 					TargetLowPriorityNodes: to.Ptr[int32](0),
	// 				},
	// 			},
	// 			TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
	// 				NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypeSpread),
	// 			},
	// 			TaskSlotsPerNode: to.Ptr[int32](1),
	// 			VMSize: to.Ptr("STANDARD_D4"),
	// 		},
	// 	},
	// }
}

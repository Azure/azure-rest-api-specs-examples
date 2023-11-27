package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolUpdate_OtherProperties.json
func ExamplePoolClient_Update_updatePoolOtherProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolClient().Update(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			ApplicationPackages: []*armbatch.ApplicationPackageReference{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
				},
				{
					ID:      to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_5678"),
					Version: to.Ptr("1.0"),
				}},
			Certificates: []*armbatch.CertificateReference{
				{
					ID:            to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567"),
					StoreLocation: to.Ptr(armbatch.CertificateStoreLocationLocalMachine),
					StoreName:     to.Ptr("MY"),
				}},
			Metadata: []*armbatch.MetadataItem{
				{
					Name:  to.Ptr("key1"),
					Value: to.Ptr("value1"),
				}},
			TargetNodeCommunicationMode: to.Ptr(armbatch.NodeCommunicationModeSimplified),
		},
	}, &armbatch.PoolClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Pool = armbatch.Pool{
	// 	Name: to.Ptr("testpool"),
	// 	Type: to.Ptr("Microsoft.Batch/batchAccounts/pools"),
	// 	Etag: to.Ptr("W/\"0x8D4EDFEBFADF4AB\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool"),
	// 	Properties: &armbatch.PoolProperties{
	// 		AllocationState: to.Ptr(armbatch.AllocationStateResizing),
	// 		AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T10:22:55.940Z"); return t}()),
	// 		ApplicationPackages: []*armbatch.ApplicationPackageReference{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_5678"),
	// 				Version: to.Ptr("1.0"),
	// 		}},
	// 		AutoScaleRun: &armbatch.AutoScaleRun{
	// 			EvaluationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T10:22:55.940Z"); return t}()),
	// 			Results: to.Ptr("$TargetDedicatedNodes=34;NodeDeallocationOption=requeue"),
	// 		},
	// 		Certificates: []*armbatch.CertificateReference{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567"),
	// 				StoreLocation: to.Ptr(armbatch.CertificateStoreLocationLocalMachine),
	// 				StoreName: to.Ptr("MY"),
	// 				Visibility: []*armbatch.CertificateVisibility{
	// 					to.Ptr(armbatch.CertificateVisibilityStartTask),
	// 					to.Ptr(armbatch.CertificateVisibilityTask),
	// 					to.Ptr(armbatch.CertificateVisibilityRemoteUser)},
	// 			}},
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
	// 			CurrentDedicatedNodes: to.Ptr[int32](12),
	// 			CurrentLowPriorityNodes: to.Ptr[int32](0),
	// 			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
	// 				CloudServiceConfiguration: &armbatch.CloudServiceConfiguration{
	// 					OSFamily: to.Ptr("5"),
	// 					OSVersion: to.Ptr("*"),
	// 				},
	// 			},
	// 			InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateDisabled),
	// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T10:22:55.940Z"); return t}()),
	// 			Metadata: []*armbatch.MetadataItem{
	// 				{
	// 					Name: to.Ptr("key1"),
	// 					Value: to.Ptr("value1"),
	// 			}},
	// 			ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
	// 			ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
	// 			ResizeOperationStatus: &armbatch.ResizeOperationStatus{
	// 				NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionTaskCompletion),
	// 				ResizeTimeout: to.Ptr("PT8M"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-29T10:22:55.940Z"); return t}()),
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
	// 			TargetNodeCommunicationMode: to.Ptr(armbatch.NodeCommunicationModeSimplified),
	// 			TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
	// 				NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypeSpread),
	// 			},
	// 			TaskSlotsPerNode: to.Ptr[int32](1),
	// 			VMSize: to.Ptr("STANDARD_D4"),
	// 		},
	// 	}
}

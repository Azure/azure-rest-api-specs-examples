package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/PoolGet_VirtualMachineConfiguration_Extensions.json
func ExamplePoolClient_Get_getPoolVirtualMachineConfigurationExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
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
	// res = armbatch.PoolClientGetResponse{
	// 	ETag: "W/\"0x8D4EDFEBFADF4AB\""
	// 	Pool: &armbatch.Pool{
	// 		Name: to.Ptr("testpool"),
	// 		Type: to.Ptr("Microsoft.Batch/batchAccounts/pools"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool"),
	// 		Properties: &armbatch.PoolProperties{
	// 			AllocationState: to.Ptr(armbatch.AllocationStateResizing),
	// 			AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.9407275Z"); return t}()),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.9407275Z"); return t}()),
	// 			CurrentDedicatedNodes: to.Ptr[int32](0),
	// 			CurrentLowPriorityNodes: to.Ptr[int32](0),
	// 			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
	// 				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
	// 					ImageReference: &armbatch.ImageReference{
	// 						Offer: to.Ptr("0001-com-ubuntu-server-focal"),
	// 						Publisher: to.Ptr("Canonical"),
	// 						SKU: to.Ptr("20_04-lts"),
	// 					},
	// 					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 20.04"),
	// 					Extensions: []*armbatch.VMExtension{
	// 						{
	// 							Name: to.Ptr("batchextension1"),
	// 							Type: to.Ptr("KeyVaultForLinux"),
	// 							AutoUpgradeMinorVersion: to.Ptr(true),
	// 							EnableAutomaticUpgrade: to.Ptr(true),
	// 							Publisher: to.Ptr("Microsoft.Azure.KeyVault"),
	// 							Settings: map[string]any{
	// 								"authenticationSettingsKey": "authenticationSettingsValue",
	// 								"secretsManagementSettingsKey": "secretsManagementSettingsValue",
	// 							},
	// 							TypeHandlerVersion: to.Ptr("2.0"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateEnabled),
	// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.9407275Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
	// 			ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.9407275Z"); return t}()),
	// 			ResizeOperationStatus: &armbatch.ResizeOperationStatus{
	// 				NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionTaskCompletion),
	// 				ResizeTimeout: to.Ptr("PT8M"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.9407275Z"); return t}()),
	// 				TargetDedicatedNodes: to.Ptr[int32](6),
	// 				TargetLowPriorityNodes: to.Ptr[int32](28),
	// 			},
	// 			ScaleSettings: &armbatch.ScaleSettings{
	// 				FixedScale: &armbatch.FixedScaleSettings{
	// 					ResizeTimeout: to.Ptr("PT8M"),
	// 					TargetDedicatedNodes: to.Ptr[int32](6),
	// 					TargetLowPriorityNodes: to.Ptr[int32](28),
	// 				},
	// 			},
	// 			TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
	// 				NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypePack),
	// 			},
	// 			TaskSlotsPerNode: to.Ptr[int32](13),
	// 			VMSize: to.Ptr("STANDARD_D4"),
	// 		},
	// 	},
	// }
}

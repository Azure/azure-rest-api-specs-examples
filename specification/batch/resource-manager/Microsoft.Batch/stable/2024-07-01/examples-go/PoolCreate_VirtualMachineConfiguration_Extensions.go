package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PoolCreate_VirtualMachineConfiguration_Extensions.json
func ExamplePoolClient_Create_createPoolVirtualMachineConfigurationExtensions() {
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
						Offer:     to.Ptr("0001-com-ubuntu-server-focal"),
						Publisher: to.Ptr("Canonical"),
						SKU:       to.Ptr("20_04-lts"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 20.04"),
					Extensions: []*armbatch.VMExtension{
						{
							Name:                    to.Ptr("batchextension1"),
							Type:                    to.Ptr("KeyVaultForLinux"),
							AutoUpgradeMinorVersion: to.Ptr(true),
							EnableAutomaticUpgrade:  to.Ptr(true),
							Publisher:               to.Ptr("Microsoft.Azure.KeyVault"),
							Settings: map[string]any{
								"authenticationSettingsKey":    "authenticationSettingsValue",
								"secretsManagementSettingsKey": "secretsManagementSettingsValue",
							},
							TypeHandlerVersion: to.Ptr("2.0"),
						}},
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				AutoScale: &armbatch.AutoScaleSettings{
					EvaluationInterval: to.Ptr("PT5M"),
					Formula:            to.Ptr("$TargetDedicatedNodes=1"),
				},
			},
			TargetNodeCommunicationMode: to.Ptr(armbatch.NodeCommunicationModeDefault),
			VMSize:                      to.Ptr("STANDARD_D4"),
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
	// 	Etag: to.Ptr("W/\"0x8D4EDFEBFADF4AB\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool"),
	// 	Properties: &armbatch.PoolProperties{
	// 		AllocationState: to.Ptr(armbatch.AllocationStateResizing),
	// 		AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
	// 		CurrentDedicatedNodes: to.Ptr[int32](0),
	// 		CurrentLowPriorityNodes: to.Ptr[int32](0),
	// 		DeploymentConfiguration: &armbatch.DeploymentConfiguration{
	// 			VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
	// 				ImageReference: &armbatch.ImageReference{
	// 					Offer: to.Ptr("0001-com-ubuntu-server-focal"),
	// 					Publisher: to.Ptr("Canonical"),
	// 					SKU: to.Ptr("20_04-lts"),
	// 				},
	// 				NodeAgentSKUID: to.Ptr("batch.node.ubuntu 20.04"),
	// 				Extensions: []*armbatch.VMExtension{
	// 					{
	// 						Name: to.Ptr("batchextension1"),
	// 						Type: to.Ptr("KeyVaultForLinux"),
	// 						AutoUpgradeMinorVersion: to.Ptr(true),
	// 						EnableAutomaticUpgrade: to.Ptr(true),
	// 						Publisher: to.Ptr("Microsoft.Azure.KeyVault"),
	// 						Settings: map[string]any{
	// 							"authenticationSettingsKey": "authenticationSettingsValue",
	// 							"secretsManagementSettingsKey": "secretsManagementSettingsValue",
	// 						},
	// 						TypeHandlerVersion: to.Ptr("2.0"),
	// 				}},
	// 			},
	// 		},
	// 		InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateDisabled),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
	// 		ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
	// 		ScaleSettings: &armbatch.ScaleSettings{
	// 			AutoScale: &armbatch.AutoScaleSettings{
	// 				EvaluationInterval: to.Ptr("PT5M"),
	// 				Formula: to.Ptr("$TargetDedicatedNodes=1"),
	// 			},
	// 		},
	// 		TargetNodeCommunicationMode: to.Ptr(armbatch.NodeCommunicationModeDefault),
	// 		TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
	// 			NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypeSpread),
	// 		},
	// 		TaskSlotsPerNode: to.Ptr[int32](1),
	// 		VMSize: to.Ptr("STANDARD_D4"),
	// 	},
	// }
}

package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_VirtualMachineConfiguration_Extensions.json
func ExamplePoolClient_Create_createPoolVirtualMachineConfigurationExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
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
							Type:                    to.Ptr("SecurityMonitoringForLinux"),
							AutoUpgradeMinorVersion: to.Ptr(true),
							ProtectedSettings: map[string]interface{}{
								"protectedSettingsKey": "protectedSettingsValue",
							},
							Publisher: to.Ptr("Microsoft.Azure.Security.Monitoring"),
							Settings: map[string]interface{}{
								"settingsKey": "settingsValue",
							},
							TypeHandlerVersion: to.Ptr("1.0"),
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
	// TODO: use response item
	_ = res
}

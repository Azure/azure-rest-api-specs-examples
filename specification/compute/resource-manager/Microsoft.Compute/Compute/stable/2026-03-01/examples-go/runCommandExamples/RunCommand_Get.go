package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/runCommandExamples/RunCommand_Get.json
func ExampleVirtualMachineRunCommandsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("24fb23e3-6ba3-41f0-9b6e-e41131d5d61e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineRunCommandsClient().Get(ctx, "SoutheastAsia", "RunPowerShellScript", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineRunCommandsClientGetResponse{
	// 	RunCommandDocument: armcompute.RunCommandDocument{
	// 		Script: []*string{
	// 			to.Ptr("param("),
	// 			to.Ptr("    [string]$arg1,"),
	// 			to.Ptr("    [string]$arg2"),
	// 			to.Ptr(")"),
	// 			to.Ptr("Write-Host This is a sample script with parameters $arg1 $arg2"),
	// 		},
	// 		Parameters: []*armcompute.RunCommandParameterDefinition{
	// 			{
	// 				Name: to.Ptr("arg1"),
	// 				Type: to.Ptr("string"),
	// 				DefaultValue: to.Ptr("value1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("arg2"),
	// 				Type: to.Ptr("string"),
	// 				DefaultValue: to.Ptr("value2"),
	// 			},
	// 		},
	// 		Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
	// 		ID: to.Ptr("RunPowerShellScript"),
	// 		OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 		Label: to.Ptr("Executes a PowerShell script"),
	// 		Description: to.Ptr("Custom multiline PowerShell script should be defined in script property. Optional parameters can be set in parameters property."),
	// 	},
	// }
}

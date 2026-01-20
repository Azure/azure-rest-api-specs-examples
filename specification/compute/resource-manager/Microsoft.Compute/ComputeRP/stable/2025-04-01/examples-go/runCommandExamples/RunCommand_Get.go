package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/runCommandExamples/RunCommand_Get.json
func ExampleVirtualMachineRunCommandsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.RunCommandDocument = armcompute.RunCommandDocument{
	// 	Description: to.Ptr("Custom multiline PowerShell script should be defined in script property. Optional parameters can be set in parameters property."),
	// 	Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
	// 	ID: to.Ptr("RunPowerShellScript"),
	// 	Label: to.Ptr("Executes a PowerShell script"),
	// 	OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 	Parameters: []*armcompute.RunCommandParameterDefinition{
	// 		{
	// 			Name: to.Ptr("arg1"),
	// 			Type: to.Ptr("string"),
	// 			DefaultValue: to.Ptr("value1"),
	// 		},
	// 		{
	// 			Name: to.Ptr("arg2"),
	// 			Type: to.Ptr("string"),
	// 			DefaultValue: to.Ptr("value2"),
	// 	}},
	// 	Script: []*string{
	// 		to.Ptr("param("),
	// 		to.Ptr("    [string]$arg1,"),
	// 		to.Ptr("    [string]$arg2"),
	// 		to.Ptr(")"),
	// 		to.Ptr("Write-Host This is a sample script with parameters $arg1 $arg2")},
	// 	}
}

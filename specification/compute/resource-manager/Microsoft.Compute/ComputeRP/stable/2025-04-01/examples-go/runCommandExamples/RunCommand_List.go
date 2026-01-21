package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/runCommandExamples/RunCommand_List.json
func ExampleVirtualMachineRunCommandsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineRunCommandsClient().NewListPager("SoutheastAsia", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.RunCommandListResult = armcompute.RunCommandListResult{
		// 	Value: []*armcompute.RunCommandDocumentBase{
		// 		{
		// 			Description: to.Ptr("Configure the machine to enable remote PowerShell."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("EnableRemotePS"),
		// 			Label: to.Ptr("Enable remote PowerShell"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Shows detailed information for the IP address, subnet mask and default gateway for each adapter bound to TCP/IP."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("IPConfig"),
		// 			Label: to.Ptr("List IP configuration"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Custom multiline PowerShell script should be defined in script property. Optional parameters can be set in parameters property."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("RunPowerShellScript"),
		// 			Label: to.Ptr("Executes a PowerShell script"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Custom multiline shell script should be defined in script property. Optional parameters can be set in parameters property."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("RunShellScript"),
		// 			Label: to.Ptr("Executes a Linux shell script"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
		// 		},
		// 		{
		// 			Description: to.Ptr("Get the configuration of all network interfaces."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("ifconfig"),
		// 			Label: to.Ptr("List network configuration"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
		// 		},
		// 		{
		// 			Description: to.Ptr("Checks if the local Administrator account is disabled, and if so enables it."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("EnableAdminAccount"),
		// 			Label: to.Ptr("Enable administrator account"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Reset built-in Administrator account password."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("ResetAccountPassword"),
		// 			Label: to.Ptr("Reset built-in Administrator account password"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Checks registry settings and domain policy settings. Suggests policy actions if machine is part of a domain or modifies the settings to default values."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("RDPSettings"),
		// 			Label: to.Ptr("Verify RDP Listener Settings"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Sets the default or user specified port number for Remote Desktop connections. Enables firewall rule for inbound access to the port."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("SetRDPPort"),
		// 			Label: to.Ptr("Set Remote Desktop port"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 		},
		// 		{
		// 			Description: to.Ptr("Removes the SSL certificate tied to the RDP listener and restores the RDP listerner security to default. Use this script if you see any issues with the certificate."),
		// 			Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 			ID: to.Ptr("ResetRDPCert"),
		// 			Label: to.Ptr("Restore RDP Authentication mode to defaults"),
		// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 	}},
		// }
	}
}

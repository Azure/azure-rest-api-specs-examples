package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/runCommandExamples/RunCommand_List.json
func ExampleVirtualMachineRunCommandsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("subid", cred, nil)
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
		// page = armcompute.VirtualMachineRunCommandsClientListResponse{
		// 	RunCommandListResult: armcompute.RunCommandListResult{
		// 		Value: []*armcompute.RunCommandDocumentBase{
		// 			{
		// 				Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 				ID: to.Ptr("EnableRemotePS"),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				Label: to.Ptr("Enable remote PowerShell"),
		// 				Description: to.Ptr("Configure the machine to enable remote PowerShell."),
		// 			},
		// 			{
		// 				Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 				ID: to.Ptr("IPConfig"),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				Label: to.Ptr("List IP configuration"),
		// 				Description: to.Ptr("Shows detailed information for the IP address, subnet mask and default gateway for each adapter bound to TCP/IP."),
		// 			},
		// 			{
		// 				Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 				ID: to.Ptr("RunPowerShellScript"),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				Label: to.Ptr("Executes a PowerShell script"),
		// 				Description: to.Ptr("Custom multiline PowerShell script should be defined in script property. Optional parameters can be set in parameters property."),
		// 			},
		// 			{
		// 				Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 				ID: to.Ptr("RunShellScript"),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
		// 				Label: to.Ptr("Executes a Linux shell script"),
		// 				Description: to.Ptr("Custom multiline shell script should be defined in script property. Optional parameters can be set in parameters property."),
		// 			},
		// 			{
		// 				Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 				ID: to.Ptr("ifconfig"),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
		// 				Label: to.Ptr("List network configuration"),
		// 				Description: to.Ptr("Get the configuration of all network interfaces."),
		// 			},
		// 			{
		// 				Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 				ID: to.Ptr("EnableAdminAccount"),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				Label: to.Ptr("Enable administrator account"),
		// 				Description: to.Ptr("Checks if the local Administrator account is disabled, and if so enables it."),
		// 			},
		// 			{
		// 				Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 				ID: to.Ptr("ResetAccountPassword"),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				Label: to.Ptr("Reset built-in Administrator account password"),
		// 				Description: to.Ptr("Reset built-in Administrator account password."),
		// 			},
		// 			{
		// 				Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 				ID: to.Ptr("RDPSettings"),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				Label: to.Ptr("Verify RDP Listener Settings"),
		// 				Description: to.Ptr("Checks registry settings and domain policy settings. Suggests policy actions if machine is part of a domain or modifies the settings to default values."),
		// 			},
		// 			{
		// 				Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 				ID: to.Ptr("SetRDPPort"),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				Label: to.Ptr("Set Remote Desktop port"),
		// 				Description: to.Ptr("Sets the default or user specified port number for Remote Desktop connections. Enables firewall rule for inbound access to the port."),
		// 			},
		// 			{
		// 				Schema: to.Ptr("http://schema.management.azure.com/schemas/2016-11-17/runcommands.json"),
		// 				ID: to.Ptr("ResetRDPCert"),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				Label: to.Ptr("Restore RDP Authentication mode to defaults"),
		// 				Description: to.Ptr("Removes the SSL certificate tied to the RDP listener and restores the RDP listerner security to default. Use this script if you see any issues with the certificate."),
		// 			},
		// 		},
		// 	},
		// }
	}
}

package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2024-02-01/examples/CreateImageTemplateWindows.json
func ExampleVirtualMachineImageTemplatesClient_BeginCreateOrUpdate_createAnImageTemplateForWindows() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvirtualmachineimagebuilder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineImageTemplatesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myImageTemplate", armvirtualmachineimagebuilder.ImageTemplate{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"imagetemplate_tag1": to.Ptr("IT_T1"),
			"imagetemplate_tag2": to.Ptr("IT_T2"),
		},
		Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
			Type: to.Ptr(armvirtualmachineimagebuilder.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armvirtualmachineimagebuilder.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity_1": {},
			},
		},
		Properties: &armvirtualmachineimagebuilder.ImageTemplateProperties{
			Customize: []armvirtualmachineimagebuilder.ImageTemplateCustomizerClassification{
				&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
					Name: to.Ptr("PowerShell (inline) Customizer Example"),
					Type: to.Ptr("PowerShell"),
					Inline: []*string{
						to.Ptr("Powershell command-1"),
						to.Ptr("Powershell command-2"),
						to.Ptr("Powershell command-3")},
				},
				&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
					Name: to.Ptr("PowerShell (inline) Customizer Elevated user Example"),
					Type: to.Ptr("PowerShell"),
					Inline: []*string{
						to.Ptr("Powershell command-1"),
						to.Ptr("Powershell command-2"),
						to.Ptr("Powershell command-3")},
					RunElevated: to.Ptr(true),
				},
				&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
					Name: to.Ptr("PowerShell (inline) Customizer Elevated Local System user Example"),
					Type: to.Ptr("PowerShell"),
					Inline: []*string{
						to.Ptr("Powershell command-1"),
						to.Ptr("Powershell command-2"),
						to.Ptr("Powershell command-3")},
					RunAsSystem: to.Ptr(true),
					RunElevated: to.Ptr(true),
				},
				&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
					Name:      to.Ptr("PowerShell (script) Customizer Example"),
					Type:      to.Ptr("PowerShell"),
					ScriptURI: to.Ptr("https://example.com/path/to/script.ps1"),
					ValidExitCodes: []*int32{
						to.Ptr[int32](0),
						to.Ptr[int32](1)},
				},
				&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
					Name:        to.Ptr("PowerShell (script) Customizer Elevated Local System user Example"),
					Type:        to.Ptr("PowerShell"),
					RunElevated: to.Ptr(true),
					ScriptURI:   to.Ptr("https://example.com/path/to/script.ps1"),
					ValidExitCodes: []*int32{
						to.Ptr[int32](0),
						to.Ptr[int32](1)},
				},
				&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
					Name:        to.Ptr("PowerShell (script) Customizer Elevated Local System user Example"),
					Type:        to.Ptr("PowerShell"),
					RunAsSystem: to.Ptr(true),
					RunElevated: to.Ptr(true),
					ScriptURI:   to.Ptr("https://example.com/path/to/script.ps1"),
					ValidExitCodes: []*int32{
						to.Ptr[int32](0),
						to.Ptr[int32](1)},
				},
				&armvirtualmachineimagebuilder.ImageTemplateRestartCustomizer{
					Name:                to.Ptr("Restart Customizer Example"),
					Type:                to.Ptr("WindowsRestart"),
					RestartCheckCommand: to.Ptr("powershell -command \"& {Write-Output 'restarted.'}\""),
					RestartCommand:      to.Ptr("shutdown /f /r /t 0 /c \"packer restart\""),
					RestartTimeout:      to.Ptr("10m"),
				},
				&armvirtualmachineimagebuilder.ImageTemplateWindowsUpdateCustomizer{
					Name: to.Ptr("Windows Update Customizer Example"),
					Type: to.Ptr("WindowsUpdate"),
					Filters: []*string{
						to.Ptr("$_.BrowseOnly")},
					SearchCriteria: to.Ptr("BrowseOnly=0 and IsInstalled=0"),
					UpdateLimit:    to.Ptr[int32](100),
				}},
			Distribute: []armvirtualmachineimagebuilder.ImageTemplateDistributorClassification{
				&armvirtualmachineimagebuilder.ImageTemplateManagedImageDistributor{
					Type: to.Ptr("ManagedImage"),
					ArtifactTags: map[string]*string{
						"tagName": to.Ptr("value"),
					},
					RunOutputName: to.Ptr("image_it_pir_1"),
					ImageID:       to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/image_it_1"),
					Location:      to.Ptr("1_location"),
				}},
			Source: &armvirtualmachineimagebuilder.ImageTemplateManagedImageSource{
				Type:    to.Ptr("ManagedImage"),
				ImageID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/images/source_image"),
			},
			VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
				OSDiskSizeGB: to.Ptr[int32](64),
				VMSize:       to.Ptr("Standard_D2s_v3"),
				VnetConfig: &armvirtualmachineimagebuilder.VirtualNetworkConfig{
					SubnetID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet_name/subnets/subnet_name"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ImageTemplate = armvirtualmachineimagebuilder.ImageTemplate{
	// 	Name: to.Ptr("myImageTemplate"),
	// 	Type: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.VirtualMachineImages/imageTemplates/myImageTemplate"),
	// 	Location: to.Ptr("westus"),
	// 	Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
	// 		Type: to.Ptr(armvirtualmachineimagebuilder.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armvirtualmachineimagebuilder.UserAssignedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity_1": &armvirtualmachineimagebuilder.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armvirtualmachineimagebuilder.ImageTemplateProperties{
	// 		Customize: []armvirtualmachineimagebuilder.ImageTemplateCustomizerClassification{
	// 			&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
	// 				Name: to.Ptr("PowerShell (inline) Customizer Example"),
	// 				Type: to.Ptr("PowerShell"),
	// 				Inline: []*string{
	// 					to.Ptr("Powershell command-1"),
	// 					to.Ptr("Powershell command-2"),
	// 					to.Ptr("Powershell command-3")},
	// 					RunAsSystem: to.Ptr(false),
	// 					RunElevated: to.Ptr(false),
	// 				},
	// 				&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
	// 					Name: to.Ptr("PowerShell (inline) Customizer Elevated user Example"),
	// 					Type: to.Ptr("PowerShell"),
	// 					Inline: []*string{
	// 						to.Ptr("Powershell command-1"),
	// 						to.Ptr("Powershell command-2"),
	// 						to.Ptr("Powershell command-3")},
	// 						RunAsSystem: to.Ptr(false),
	// 						RunElevated: to.Ptr(true),
	// 					},
	// 					&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
	// 						Name: to.Ptr("PowerShell (inline) Customizer Elevated Local System user Example"),
	// 						Type: to.Ptr("PowerShell"),
	// 						Inline: []*string{
	// 							to.Ptr("Powershell command-1"),
	// 							to.Ptr("Powershell command-2"),
	// 							to.Ptr("Powershell command-3")},
	// 							RunAsSystem: to.Ptr(true),
	// 							RunElevated: to.Ptr(true),
	// 						},
	// 						&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
	// 							Name: to.Ptr("PowerShell (script) Customizer Example"),
	// 							Type: to.Ptr("PowerShell"),
	// 							RunAsSystem: to.Ptr(false),
	// 							RunElevated: to.Ptr(false),
	// 							ScriptURI: to.Ptr("https://example.com/path/to/script.ps1"),
	// 							ValidExitCodes: []*int32{
	// 								to.Ptr[int32](0),
	// 								to.Ptr[int32](1)},
	// 							},
	// 							&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
	// 								Name: to.Ptr("PowerShell (script) Customizer Elevated Local System user Example"),
	// 								Type: to.Ptr("PowerShell"),
	// 								RunAsSystem: to.Ptr(false),
	// 								RunElevated: to.Ptr(true),
	// 								ScriptURI: to.Ptr("https://example.com/path/to/script.ps1"),
	// 								ValidExitCodes: []*int32{
	// 									to.Ptr[int32](0),
	// 									to.Ptr[int32](1)},
	// 								},
	// 								&armvirtualmachineimagebuilder.ImageTemplatePowerShellCustomizer{
	// 									Name: to.Ptr("PowerShell (script) Customizer Elevated Local System user Example"),
	// 									Type: to.Ptr("PowerShell"),
	// 									RunAsSystem: to.Ptr(true),
	// 									RunElevated: to.Ptr(true),
	// 									ScriptURI: to.Ptr("https://example.com/path/to/script.ps1"),
	// 									ValidExitCodes: []*int32{
	// 										to.Ptr[int32](0),
	// 										to.Ptr[int32](1)},
	// 									},
	// 									&armvirtualmachineimagebuilder.ImageTemplateRestartCustomizer{
	// 										Name: to.Ptr("Restart Customizer Example"),
	// 										Type: to.Ptr("WindowsRestart"),
	// 										RestartCheckCommand: to.Ptr("powershell -command \"& {Write-Output 'restarted.'}\""),
	// 										RestartCommand: to.Ptr("shutdown /f /r /t 0 /c \"packer restart\""),
	// 										RestartTimeout: to.Ptr("10m"),
	// 									},
	// 									&armvirtualmachineimagebuilder.ImageTemplateWindowsUpdateCustomizer{
	// 										Name: to.Ptr("Windows Update Customizer Example"),
	// 										Type: to.Ptr("WindowsUpdate"),
	// 										Filters: []*string{
	// 											to.Ptr("$_.BrowseOnly")},
	// 											SearchCriteria: to.Ptr("BrowseOnly=0 and IsInstalled=0"),
	// 											UpdateLimit: to.Ptr[int32](100),
	// 									}},
	// 									Distribute: []armvirtualmachineimagebuilder.ImageTemplateDistributorClassification{
	// 										&armvirtualmachineimagebuilder.ImageTemplateManagedImageDistributor{
	// 											Type: to.Ptr("ManagedImage"),
	// 											ArtifactTags: map[string]*string{
	// 												"tagName": to.Ptr("value"),
	// 											},
	// 											RunOutputName: to.Ptr("image_it_pir_1"),
	// 											ImageID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/image_it_1"),
	// 											Location: to.Ptr("1_location"),
	// 									}},
	// 									Source: &armvirtualmachineimagebuilder.ImageTemplateManagedImageSource{
	// 										Type: to.Ptr("ManagedImage"),
	// 										ImageID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/images/source_image"),
	// 									},
	// 									VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
	// 										OSDiskSizeGB: to.Ptr[int32](64),
	// 										VMSize: to.Ptr("Standard_D2s_v3"),
	// 									},
	// 								},
	// 							}
}

package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/machine/Machines_ListByResourceGroup.json
func ExampleMachinesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMachinesClient().NewListByResourceGroupPager("myResourceGroup", &armhybridcompute.MachinesClientListByResourceGroupOptions{Expand: nil})
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
		// page.MachineListResult = armhybridcompute.MachineListResult{
		// 	Value: []*armhybridcompute.Machine{
		// 		{
		// 			Name: to.Ptr("myMachine"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/machines"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/machines/myMachine"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Identity: &armhybridcompute.Identity{
		// 				Type: to.Ptr("SystemAssigned"),
		// 				PrincipalID: to.Ptr("f7a068cc-b0b8-46e8-a203-22f301a62a8f"),
		// 				TenantID: to.Ptr("c4098cc-91b8-46c2-a205-d82ab1a62a8f"),
		// 			},
		// 			Properties: &armhybridcompute.MachineProperties{
		// 				AgentConfiguration: &armhybridcompute.AgentConfiguration{
		// 					ConfigMode: to.Ptr(armhybridcompute.AgentConfigurationModeFull),
		// 					ExtensionsEnabled: to.Ptr("true"),
		// 					GuestConfigurationEnabled: to.Ptr("true"),
		// 					IncomingConnectionsPorts: []*string{
		// 						to.Ptr("22"),
		// 						to.Ptr("23")},
		// 						ProxyBypass: []*string{
		// 							to.Ptr("proxy1"),
		// 							to.Ptr("proxy2")},
		// 							ProxyURL: to.Ptr("https://test.test"),
		// 						},
		// 						ClientPublicKey: to.Ptr("string"),
		// 						DetectedProperties: map[string]*string{
		// 							"cloudprovider": to.Ptr("N/A"),
		// 							"manufacturer": to.Ptr("Microsoft Corporation"),
		// 							"model": to.Ptr("Virtual Machine"),
		// 						},
		// 						LicenseProfile: &armhybridcompute.LicenseProfileMachineInstanceView{
		// 							EsuProfile: &armhybridcompute.LicenseProfileMachineInstanceViewEsuProperties{
		// 								EsuKeys: []*armhybridcompute.EsuKey{
		// 									{
		// 										LicenseStatus: to.Ptr("licenseStatus1"),
		// 										SKU: to.Ptr("skuNumber1"),
		// 									},
		// 									{
		// 										LicenseStatus: to.Ptr("licenseStatus2"),
		// 										SKU: to.Ptr("skuNumber2"),
		// 								}},
		// 								EsuEligibility: to.Ptr(armhybridcompute.EsuEligibilityIneligible),
		// 								EsuKeyState: to.Ptr(armhybridcompute.EsuKeyStateInactive),
		// 								ServerType: to.Ptr(armhybridcompute.EsuServerTypeStandard),
		// 								LicenseAssignmentState: to.Ptr(armhybridcompute.LicenseAssignmentStateAssigned),
		// 							},
		// 						},
		// 						LocationData: &armhybridcompute.LocationData{
		// 							Name: to.Ptr("Redmond"),
		// 						},
		// 						MssqlDiscovered: to.Ptr("false"),
		// 						OSProfile: &armhybridcompute.OSProfile{
		// 							LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
		// 								PatchSettings: &armhybridcompute.PatchSettings{
		// 								},
		// 							},
		// 							WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
		// 								PatchSettings: &armhybridcompute.PatchSettings{
		// 								},
		// 							},
		// 						},
		// 						PrivateLinkScopeResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						VMID: to.Ptr("b7a098cc-b0b8-46e8-a205-62f301a62a8f"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("myMachine2"),
		// 					Type: to.Ptr("Microsoft.HybridCompute/machines"),
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/machines/myMachine2"),
		// 					Location: to.Ptr("westus2"),
		// 					Identity: &armhybridcompute.Identity{
		// 						Type: to.Ptr("SystemAssigned"),
		// 						PrincipalID: to.Ptr("e7a068cc-b0b8-46e8-a203-22f301a62a8f"),
		// 						TenantID: to.Ptr("c4098cc-91b8-46c2-a205-d82ab1a62a8f"),
		// 					},
		// 					Properties: &armhybridcompute.MachineProperties{
		// 						AgentConfiguration: &armhybridcompute.AgentConfiguration{
		// 							ConfigMode: to.Ptr(armhybridcompute.AgentConfigurationModeFull),
		// 							ExtensionsEnabled: to.Ptr("true"),
		// 							GuestConfigurationEnabled: to.Ptr("true"),
		// 							IncomingConnectionsPorts: []*string{
		// 								to.Ptr("22"),
		// 								to.Ptr("23")},
		// 								ProxyBypass: []*string{
		// 									to.Ptr("proxy1"),
		// 									to.Ptr("proxy2")},
		// 									ProxyURL: to.Ptr("https://test.test"),
		// 								},
		// 								ClientPublicKey: to.Ptr("string"),
		// 								DetectedProperties: map[string]*string{
		// 									"cloudprovider": to.Ptr("N/A"),
		// 									"manufacturer": to.Ptr("Microsoft Corporation"),
		// 									"model": to.Ptr("Surfacebook"),
		// 								},
		// 								LicenseProfile: &armhybridcompute.LicenseProfileMachineInstanceView{
		// 									EsuProfile: &armhybridcompute.LicenseProfileMachineInstanceViewEsuProperties{
		// 										EsuKeys: []*armhybridcompute.EsuKey{
		// 											{
		// 												LicenseStatus: to.Ptr("licenseStatus1"),
		// 												SKU: to.Ptr("skuNumber1"),
		// 											},
		// 											{
		// 												LicenseStatus: to.Ptr("licenseStatus2"),
		// 												SKU: to.Ptr("skuNumber2"),
		// 										}},
		// 										EsuEligibility: to.Ptr(armhybridcompute.EsuEligibilityIneligible),
		// 										EsuKeyState: to.Ptr(armhybridcompute.EsuKeyStateInactive),
		// 										ServerType: to.Ptr(armhybridcompute.EsuServerTypeStandard),
		// 										LicenseAssignmentState: to.Ptr(armhybridcompute.LicenseAssignmentStateAssigned),
		// 									},
		// 								},
		// 								LocationData: &armhybridcompute.LocationData{
		// 									Name: to.Ptr("Redmond"),
		// 								},
		// 								MssqlDiscovered: to.Ptr("true"),
		// 								OSProfile: &armhybridcompute.OSProfile{
		// 									LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
		// 										PatchSettings: &armhybridcompute.PatchSettings{
		// 										},
		// 									},
		// 									WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
		// 										PatchSettings: &armhybridcompute.PatchSettings{
		// 										},
		// 									},
		// 								},
		// 								ParentClusterResourceID: to.Ptr("{AzureStackHCIResourceId}"),
		// 								PrivateLinkScopeResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
		// 								ProvisioningState: to.Ptr("Succeeded"),
		// 								VMID: to.Ptr("a4a098cc-b0b8-46e8-a205-62f301a62a8f"),
		// 							},
		// 					}},
		// 				}
	}
}

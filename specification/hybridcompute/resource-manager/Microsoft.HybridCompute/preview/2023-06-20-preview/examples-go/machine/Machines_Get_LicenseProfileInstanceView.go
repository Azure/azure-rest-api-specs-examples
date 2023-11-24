package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/machine/Machines_Get_LicenseProfileInstanceView.json
func ExampleMachinesClient_Get_getMachineWithLicenseProfileInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMachinesClient().Get(ctx, "myResourceGroup", "myMachine", &armhybridcompute.MachinesClientGetOptions{Expand: to.Ptr(armhybridcompute.InstanceViewTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Machine = armhybridcompute.Machine{
	// 	Name: to.Ptr("myMachine"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/machines"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/machines/myMachine"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Identity: &armhybridcompute.Identity{
	// 		Type: to.Ptr("SystemAssigned"),
	// 		PrincipalID: to.Ptr("string"),
	// 		TenantID: to.Ptr("string"),
	// 	},
	// 	Properties: &armhybridcompute.MachineProperties{
	// 		AgentConfiguration: &armhybridcompute.AgentConfiguration{
	// 			ConfigMode: to.Ptr(armhybridcompute.AgentConfigurationModeFull),
	// 			ExtensionsEnabled: to.Ptr("true"),
	// 			GuestConfigurationEnabled: to.Ptr("true"),
	// 			IncomingConnectionsPorts: []*string{
	// 				to.Ptr("22"),
	// 				to.Ptr("23")},
	// 				ProxyBypass: []*string{
	// 					to.Ptr("proxy1"),
	// 					to.Ptr("proxy2")},
	// 					ProxyURL: to.Ptr("https://test.test"),
	// 				},
	// 				ClientPublicKey: to.Ptr("string"),
	// 				DetectedProperties: map[string]*string{
	// 					"cloudprovider": to.Ptr("N/A"),
	// 					"manufacturer": to.Ptr("Microsoft Corporation"),
	// 					"model": to.Ptr("Virtual Machine"),
	// 				},
	// 				LicenseProfile: &armhybridcompute.LicenseProfileMachineInstanceView{
	// 					EsuProfile: &armhybridcompute.LicenseProfileMachineInstanceViewEsuProperties{
	// 						EsuKeys: []*armhybridcompute.EsuKey{
	// 							{
	// 								LicenseStatus: to.Ptr("licenseStatus1"),
	// 								SKU: to.Ptr("skuNumber1"),
	// 							},
	// 							{
	// 								LicenseStatus: to.Ptr("licenseStatus2"),
	// 								SKU: to.Ptr("skuNumber2"),
	// 						}},
	// 						EsuEligibility: to.Ptr(armhybridcompute.EsuEligibilityIneligible),
	// 						EsuKeyState: to.Ptr(armhybridcompute.EsuKeyStateInactive),
	// 						ServerType: to.Ptr(armhybridcompute.EsuServerTypeStandard),
	// 						AssignedLicense: &armhybridcompute.License{
	// 							Name: to.Ptr("{licenseName}"),
	// 							Type: to.Ptr("Microsoft.HybridCompute/licenses"),
	// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Licenses/{licenseName}"),
	// 							Location: to.Ptr("eastus2euap"),
	// 							Properties: &armhybridcompute.LicenseProperties{
	// 								LicenseDetails: &armhybridcompute.LicenseDetails{
	// 									Type: to.Ptr(armhybridcompute.LicenseCoreTypePCore),
	// 									AssignedLicenses: to.Ptr[int32](8),
	// 									Edition: to.Ptr(armhybridcompute.LicenseEditionDatacenter),
	// 									ImmutableID: to.Ptr("<generated Guid>"),
	// 									Processors: to.Ptr[int32](6),
	// 									State: to.Ptr(armhybridcompute.LicenseStateActivated),
	// 									Target: to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
	// 								},
	// 								LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
	// 								TenantID: to.Ptr("{tenandId}"),
	// 							},
	// 						},
	// 						LicenseAssignmentState: to.Ptr(armhybridcompute.LicenseAssignmentStateAssigned),
	// 					},
	// 				},
	// 				LocationData: &armhybridcompute.LocationData{
	// 					Name: to.Ptr("Redmond"),
	// 					City: to.Ptr("redmond"),
	// 					CountryOrRegion: to.Ptr("usa"),
	// 				},
	// 				MssqlDiscovered: to.Ptr("false"),
	// 				NetworkProfile: &armhybridcompute.NetworkProfile{
	// 					NetworkInterfaces: []*armhybridcompute.NetworkInterface{
	// 						{
	// 							IPAddresses: []*armhybridcompute.IPAddress{
	// 								{
	// 									Address: to.Ptr("192.168.12.345"),
	// 									IPAddressVersion: to.Ptr("IPv4"),
	// 									Subnet: &armhybridcompute.Subnet{
	// 										AddressPrefix: to.Ptr("192.168.12.0/24"),
	// 									},
	// 							}},
	// 						},
	// 						{
	// 							IPAddresses: []*armhybridcompute.IPAddress{
	// 								{
	// 									Address: to.Ptr("1001:0:34aa:5000:1234:aaaa:bbbb:cccc"),
	// 									IPAddressVersion: to.Ptr("IPv6"),
	// 									Subnet: &armhybridcompute.Subnet{
	// 										AddressPrefix: to.Ptr("1001:0:34aa:5000::/64"),
	// 									},
	// 							}},
	// 					}},
	// 				},
	// 				OSProfile: &armhybridcompute.OSProfile{
	// 					LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
	// 						PatchSettings: &armhybridcompute.PatchSettings{
	// 						},
	// 					},
	// 					WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
	// 						PatchSettings: &armhybridcompute.PatchSettings{
	// 						},
	// 					},
	// 				},
	// 				ParentClusterResourceID: to.Ptr("{AzureStackHCIResourceId}"),
	// 				PrivateLinkScopeResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				ServiceStatuses: &armhybridcompute.ServiceStatuses{
	// 					ExtensionService: &armhybridcompute.ServiceStatus{
	// 						StartupType: to.Ptr("Automatic"),
	// 						Status: to.Ptr("Running"),
	// 					},
	// 					GuestConfigurationService: &armhybridcompute.ServiceStatus{
	// 						StartupType: to.Ptr("Automatic"),
	// 						Status: to.Ptr("Running"),
	// 					},
	// 				},
	// 				VMID: to.Ptr("b7a098cc-b0b8-46e8-a205-62f301a62a8f"),
	// 			},
	// 			Resources: []*armhybridcompute.MachineExtension{
	// 			},
	// 		}
}

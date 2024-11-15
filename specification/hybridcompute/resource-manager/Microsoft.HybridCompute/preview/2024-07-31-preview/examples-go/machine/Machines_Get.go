package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/63d530d0def1c624f5d42d39170ff4ac196522e2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/machine/Machines_Get.json
func ExampleMachinesClient_Get_getMachine() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMachinesClient().Get(ctx, "myResourceGroup", "myMachine", &armhybridcompute.MachinesClientGetOptions{Expand: nil})
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
	// 				FirmwareProfile: &armhybridcompute.FirmwareProfile{
	// 					Type: to.Ptr("BIOS"),
	// 					SerialNumber: to.Ptr("007f0232-1c2e-4978-8604-ea44e7a5f5a0"),
	// 				},
	// 				HardwareProfile: &armhybridcompute.HardwareProfile{
	// 					NumberOfCPUSockets: to.Ptr[int32](2),
	// 					Processors: []*armhybridcompute.Processor{
	// 						{
	// 							Name: to.Ptr("Intel(R) Core(TM) i7-10610U CPU @ 1.80GHz"),
	// 							NumberOfCores: to.Ptr[int32](4),
	// 					}},
	// 					TotalPhysicalMemoryInBytes: to.Ptr[int64](34359738368),
	// 				},
	// 				LicenseProfile: &armhybridcompute.LicenseProfileMachineInstanceView{
	// 					EsuProfile: &armhybridcompute.LicenseProfileMachineInstanceViewEsuProperties{
	// 						EsuKeys: []*armhybridcompute.EsuKey{
	// 							{
	// 								LicenseStatus: to.Ptr[int32](1),
	// 								SKU: to.Ptr("skuNumber1"),
	// 							},
	// 							{
	// 								LicenseStatus: to.Ptr[int32](1),
	// 								SKU: to.Ptr("skuNumber2"),
	// 						}},
	// 						EsuEligibility: to.Ptr(armhybridcompute.EsuEligibilityIneligible),
	// 						EsuKeyState: to.Ptr(armhybridcompute.EsuKeyStateInactive),
	// 						ServerType: to.Ptr(armhybridcompute.EsuServerTypeStandard),
	// 						LicenseAssignmentState: to.Ptr(armhybridcompute.LicenseAssignmentStateAssigned),
	// 					},
	// 					LicenseChannel: to.Ptr("PSG"),
	// 					LicenseStatus: to.Ptr(armhybridcompute.LicenseStatusLicensed),
	// 					ProductProfile: &armhybridcompute.LicenseProfileArmProductProfileProperties{
	// 						BillingEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 						BillingStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 						DisenrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 						EnrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 						ProductFeatures: []*armhybridcompute.ProductFeature{
	// 							{
	// 								Name: to.Ptr("Hotpatch"),
	// 								BillingEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 								BillingStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 								DisenrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 								EnrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 								SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabled),
	// 						}},
	// 						ProductType: to.Ptr(armhybridcompute.LicenseProfileProductTypeWindowsServer),
	// 						SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabled),
	// 					},
	// 					SoftwareAssurance: &armhybridcompute.LicenseProfileMachineInstanceViewSoftwareAssurance{
	// 						SoftwareAssuranceCustomer: to.Ptr(true),
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
	// 							Name: to.Ptr("Wi-Fi"),
	// 							ID: to.Ptr("8"),
	// 							IPAddresses: []*armhybridcompute.IPAddress{
	// 								{
	// 									Address: to.Ptr("192.168.12.345"),
	// 									IPAddressVersion: to.Ptr("IPv4"),
	// 									Subnet: &armhybridcompute.Subnet{
	// 										AddressPrefix: to.Ptr("192.168.12.0/24"),
	// 									},
	// 							}},
	// 							MacAddress: to.Ptr("3c:49:6e:13:0e:73"),
	// 						},
	// 						{
	// 							Name: to.Ptr("Ethernet"),
	// 							ID: to.Ptr("23"),
	// 							IPAddresses: []*armhybridcompute.IPAddress{
	// 								{
	// 									Address: to.Ptr("1001:0:34aa:5000:1234:aaaa:bbbb:cccc"),
	// 									IPAddressVersion: to.Ptr("IPv6"),
	// 									Subnet: &armhybridcompute.Subnet{
	// 										AddressPrefix: to.Ptr("1001:0:34aa:5000::/64"),
	// 									},
	// 							}},
	// 							MacAddress: to.Ptr("10:15:5c:52:f9:b8"),
	// 					}},
	// 				},
	// 				OSEdition: to.Ptr("Standard"),
	// 				OSProfile: &armhybridcompute.OSProfile{
	// 					LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
	// 						PatchSettings: &armhybridcompute.PatchSettings{
	// 						},
	// 					},
	// 					WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
	// 						PatchSettings: &armhybridcompute.PatchSettings{
	// 							EnableHotpatching: to.Ptr(true),
	// 							Status: &armhybridcompute.PatchSettingsStatus{
	// 								HotpatchEnablementStatus: to.Ptr(armhybridcompute.HotpatchEnablementStatusEnabled),
	// 							},
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
	// 				StorageProfile: &armhybridcompute.StorageProfile{
	// 					Disks: []*armhybridcompute.Disk{
	// 						{
	// 							Name: to.Ptr("Windows"),
	// 							Path: to.Ptr("C:/"),
	// 							DiskType: to.Ptr("Fixed"),
	// 							GeneratedID: to.Ptr("94318602-6e46-4eaa-997e-0e528afe3d17"),
	// 							ID: to.Ptr("2"),
	// 							MaxSizeInBytes: to.Ptr[int64](1022870155264),
	// 							UsedSpaceInBytes: to.Ptr[int64](435501297664),
	// 					}},
	// 				},
	// 				VMID: to.Ptr("b7a098cc-b0b8-46e8-a205-62f301a62a8f"),
	// 			},
	// 			Resources: []*armhybridcompute.MachineExtension{
	// 			},
	// 		}
}

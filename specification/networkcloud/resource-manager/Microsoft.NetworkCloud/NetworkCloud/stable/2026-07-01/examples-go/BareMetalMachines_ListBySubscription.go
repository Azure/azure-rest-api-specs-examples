package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud/v2"
)

// Generated from example definition: 2026-07-01/BareMetalMachines_ListBySubscription.json
func ExampleBareMetalMachinesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("123e4567-e89b-12d3-a456-426655440000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBareMetalMachinesClient().NewListBySubscriptionPager(nil)
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
		// page = armnetworkcloud.BareMetalMachinesClientListBySubscriptionResponse{
		// 	BareMetalMachineList: armnetworkcloud.BareMetalMachineList{
		// 		NextLink: to.Ptr("https://fully.qualified.hyperlink"),
		// 		Value: []*armnetworkcloud.BareMetalMachine{
		// 			{
		// 				ExtendedLocation: &armnetworkcloud.ExtendedLocation{
		// 					Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
		// 					Type: to.Ptr("CustomLocation"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 				Location: to.Ptr("location"),
		// 				Name: to.Ptr("bareMetalMachineName"),
		// 				Properties: &armnetworkcloud.BareMetalMachineProperties{
		// 					ActionStates: []*armnetworkcloud.ActionState{
		// 						{
		// 							ActionType: to.Ptr("Microsoft.NetworkCloud/bareMetalMachines/reimage/action"),
		// 							CorrelationID: to.Ptr("a45a00bb-3b02-42d5-baaf-033497574e97"),
		// 							EndTime: to.Ptr("2023-04-29T12:00:00Z"),
		// 							Message: to.Ptr("Action completed successfully"),
		// 							StartTime: to.Ptr("2023-04-29T11:00:00Z"),
		// 							Status: to.Ptr(armnetworkcloud.ActionStateStatusCompleted),
		// 							StepStates: []*armnetworkcloud.StepState{
		// 								{
		// 									EndTime: to.Ptr("2023-04-29T11:30:00Z"),
		// 									Message: to.Ptr("BareMetalMachine was validated as ready for reimage"),
		// 									StartTime: to.Ptr("2023-04-29T11:00:00Z"),
		// 									Status: to.Ptr(armnetworkcloud.StepStateStatusCompleted),
		// 									StepName: to.Ptr("ValidateImageState"),
		// 								},
		// 								{
		// 									EndTime: to.Ptr("2023-04-29T11:45:00Z"),
		// 									Message: to.Ptr("BareMetalMachine was reimaged successfully"),
		// 									StartTime: to.Ptr("2023-04-29T11:30:00Z"),
		// 									Status: to.Ptr(armnetworkcloud.StepStateStatusCompleted),
		// 									StepName: to.Ptr("Reimage"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					AssociatedResourceIDs: []*string{
		// 						to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/virtualMachines/virtualMachineName"),
		// 					},
		// 					BmcConnectionString: to.Ptr("redfish+https://10.10.10.16/redfish/v1/Systems/System.Embedded.1"),
		// 					BmcCredentials: &armnetworkcloud.AdministrativeCredentials{
		// 						Username: to.Ptr("bmcuser"),
		// 					},
		// 					BmcIPv4Address: to.Ptr("10.10.10.16"),
		// 					BmcIPv6Address: to.Ptr("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
		// 					BmcMacAddress: to.Ptr("00:00:4f:00:57:00"),
		// 					BootMacAddress: to.Ptr("00:00:4e:00:58:af"),
		// 					CaCertificate: &armnetworkcloud.CertificateInfo{
		// 						Hash: to.Ptr("dea698309efd2830a1d440a807650d9aa6d954b3243ab8cb556ac98c1f3faa60"),
		// 						Value: to.Ptr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWgAwIBAgIJAL4a5b1d8f2wM...A0GCSqGSIb3DQEBCwUAMEUxCzAJB==\n-----END CERTIFICATE-----"),
		// 					},
		// 					ClusterID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
		// 					CordonStatus: to.Ptr(armnetworkcloud.BareMetalMachineCordonStatusUncordoned),
		// 					DetailedStatus: to.Ptr(armnetworkcloud.BareMetalMachineDetailedStatusAvailable),
		// 					DetailedStatusMessage: to.Ptr("DetailedStatusMessage"),
		// 					HardwareInventory: &armnetworkcloud.HardwareInventory{
		// 						AdditionalHostInformation: to.Ptr("Machine specific information..."),
		// 						Interfaces: []*armnetworkcloud.HardwareInventoryNetworkInterface{
		// 							{
		// 								LinkStatus: to.Ptr("Up"),
		// 								MacAddress: to.Ptr("2C:54:91:88:C9:E3"),
		// 								Name: to.Ptr("networkInterfaceName"),
		// 								NetworkInterfaceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkDevices/networkDeviceName/networkInterfaces/networkInterfaceName"),
		// 							},
		// 						},
		// 					},
		// 					HardwareValidationStatus: &armnetworkcloud.HardwareValidationStatus{
		// 						LastValidationTime: to.Ptr(time.Date(2022, time.September, 30, 13, 27, 3, 8000000, time.UTC)),
		// 						Result: to.Ptr(armnetworkcloud.BareMetalMachineHardwareValidationResultPass),
		// 					},
		// 					KubernetesNodeName: to.Ptr("node01"),
		// 					KubernetesVersion: to.Ptr("1.28.3"),
		// 					MachineClusterVersion: to.Ptr("3.8.7"),
		// 					MachineDetails: to.Ptr("User-provided machine details."),
		// 					MachineName: to.Ptr("r01c001"),
		// 					MachineRoles: []*string{
		// 						to.Ptr("platform.afo-nc.microsoft.com/management-plane=true"),
		// 					},
		// 					MachineSKUID: to.Ptr("684E-3B16-399E"),
		// 					MonitoringConfigurationStatus: &armnetworkcloud.BareMetalMachineMonitoringConfigurationStatus{
		// 						LogLevel: to.Ptr(armnetworkcloud.BareMetalMachineMonitoringConfigurationStatusLogLevelDefault),
		// 						MetricsLevel: to.Ptr(armnetworkcloud.BareMetalMachineMonitoringConfigurationStatusMetricsLevelDefault),
		// 					},
		// 					OamIPv4Address: to.Ptr("192.0.2.1"),
		// 					OamIPv6Address: to.Ptr("0:0:0:0:0:FFFF:7F00:0001"),
		// 					OSImage: to.Ptr("v20220805e"),
		// 					PowerState: to.Ptr(armnetworkcloud.BareMetalMachinePowerStateOn),
		// 					ProvisioningState: to.Ptr(armnetworkcloud.BareMetalMachineProvisioningStateSucceeded),
		// 					RackID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"),
		// 					RackSlot: to.Ptr[int64](1),
		// 					ReadyState: to.Ptr(armnetworkcloud.BareMetalMachineReadyStateTrue),
		// 					RuntimeProtectionStatus: &armnetworkcloud.RuntimeProtectionStatus{
		// 						AgentHealthStatus: to.Ptr(armnetworkcloud.RuntimeProtectionAgentHealthStatusUnhealthy),
		// 						AgentHealthStatusIssues: []*string{
		// 							to.Ptr("Issue1"),
		// 							to.Ptr("Issue2"),
		// 						},
		// 						AgentLicenseStatus: to.Ptr(armnetworkcloud.RuntimeProtectionAgentLicenseStatusLicensed),
		// 						DefinitionUpdateMode: to.Ptr(armnetworkcloud.RuntimeProtectionDefinitionUpdateModeNone),
		// 						DefinitionsLastUpdated: to.Ptr(time.Date(2023, time.September, 28, 13, 27, 3, 8000000, time.UTC)),
		// 						DefinitionsVersion: to.Ptr("1.2.3"),
		// 						EnforcementLevel: to.Ptr(armnetworkcloud.RuntimeProtectionEnforcementLevelOnDemand),
		// 						ScanCompletedTime: to.Ptr(time.Date(2023, time.September, 30, 13, 27, 23, 103000000, time.UTC)),
		// 						ScanScheduledTime: to.Ptr(time.Date(2023, time.October, 1, 13, 0, 0, 0, time.UTC)),
		// 						ScanStartedTime: to.Ptr(time.Date(2023, time.September, 30, 13, 0, 3, 8000000, time.UTC)),
		// 					},
		// 					SecretRotationStatus: []*armnetworkcloud.SecretRotationStatus{
		// 						{
		// 							ExpirePeriodDays: to.Ptr[int64](90),
		// 							LastRotationTime: to.Ptr(time.Date(2023, time.September, 30, 13, 27, 23, 103000000, time.UTC)),
		// 							RotationPeriodDays: to.Ptr[int64](60),
		// 							SecretArchiveReference: &armnetworkcloud.SecretArchiveReference{
		// 								KeyVaultID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.KeyVault/vaults/keyVaultName"),
		// 								SecretName: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff-resource-group-cluster-1679871-console-credential-manager-ffffffff"),
		// 								SecretVersion: to.Ptr("02ab6c1f9c0f4982b0632b0d5d74a33b"),
		// 							},
		// 							SecretType: to.Ptr("Bare Metal Machine Identity - console"),
		// 						},
		// 					},
		// 					SerialNumber: to.Ptr("BM1219XXX"),
		// 					ServiceTag: to.Ptr("ST1219XXX"),
		// 				},
		// 				SystemData: &armnetworkcloud.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2021, time.January, 22, 13, 27, 3, 8000000, time.UTC)),
		// 					CreatedBy: to.Ptr("identityA"),
		// 					CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(time.Date(2021, time.January, 22, 13, 29, 3, 1000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("identityB"),
		// 					LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("myvalue1"),
		// 					"key2": to.Ptr("myvalue2"),
		// 				},
		// 				Type: to.Ptr("Microsoft.NetworkCloud/bareMetalMachines"),
		// 			},
		// 		},
		// 	},
		// }
	}
}

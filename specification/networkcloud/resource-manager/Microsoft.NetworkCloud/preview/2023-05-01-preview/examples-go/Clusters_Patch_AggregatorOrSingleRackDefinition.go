package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/Clusters_Patch_AggregatorOrSingleRackDefinition.json
func ExampleClustersClient_BeginUpdate_patchClusterAggregatorOrSingleRackDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdate(ctx, "resourceGroupName", "clusterName", armnetworkcloud.ClusterPatchParameters{
		Properties: &armnetworkcloud.ClusterPatchProperties{
			AggregatorOrSingleRackDefinition: &armnetworkcloud.RackDefinition{
				BareMetalMachineConfigurationData: []*armnetworkcloud.BareMetalMachineConfigurationData{
					{
						BmcCredentials: &armnetworkcloud.AdministrativeCredentials{
							Password: to.Ptr("{password}"),
							Username: to.Ptr("username"),
						},
						BmcMacAddress:  to.Ptr("AA:BB:CC:DD:EE:FF"),
						BootMacAddress: to.Ptr("00:BB:CC:DD:EE:FF"),
						MachineDetails: to.Ptr("extraDetails"),
						MachineName:    to.Ptr("bmmName1"),
						RackSlot:       to.Ptr[int64](1),
						SerialNumber:   to.Ptr("BM1219XXX"),
					},
					{
						BmcCredentials: &armnetworkcloud.AdministrativeCredentials{
							Password: to.Ptr("{password}"),
							Username: to.Ptr("username"),
						},
						BmcMacAddress:  to.Ptr("AA:BB:CC:DD:EE:00"),
						BootMacAddress: to.Ptr("00:BB:CC:DD:EE:00"),
						MachineDetails: to.Ptr("extraDetails"),
						MachineName:    to.Ptr("bmmName2"),
						RackSlot:       to.Ptr[int64](2),
						SerialNumber:   to.Ptr("BM1219YYY"),
					}},
				NetworkRackID:    to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName"),
				RackLocation:     to.Ptr("Foo Datacenter, Floor 3, Aisle 9, Rack 2"),
				RackSerialNumber: to.Ptr("newSerialNumber"),
				RackSKUID:        to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"),
				StorageApplianceConfigurationData: []*armnetworkcloud.StorageApplianceConfigurationData{
					{
						AdminCredentials: &armnetworkcloud.AdministrativeCredentials{
							Password: to.Ptr("{password}"),
							Username: to.Ptr("username"),
						},
						RackSlot:             to.Ptr[int64](1),
						SerialNumber:         to.Ptr("BM1219XXX"),
						StorageApplianceName: to.Ptr("vmName"),
					}},
			},
			ComputeDeploymentThreshold: &armnetworkcloud.ValidationThreshold{
				Type:     to.Ptr(armnetworkcloud.ValidationThresholdTypePercentSuccess),
				Grouping: to.Ptr(armnetworkcloud.ValidationThresholdGroupingPerCluster),
				Value:    to.Ptr[int64](90),
			},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
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
	// res.Cluster = armnetworkcloud.Cluster{
	// 	Name: to.Ptr("clusterName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/clusters"),
	// 	ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterManagerExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.ClusterProperties{
	// 		AggregatorOrSingleRackDefinition: &armnetworkcloud.RackDefinition{
	// 			BareMetalMachineConfigurationData: []*armnetworkcloud.BareMetalMachineConfigurationData{
	// 				{
	// 					BmcConnectionString: to.Ptr("bmcConnectionString"),
	// 					BmcCredentials: &armnetworkcloud.AdministrativeCredentials{
	// 						Username: to.Ptr("username"),
	// 					},
	// 					BmcMacAddress: to.Ptr("AA:BB:CC:DD:EE:FF"),
	// 					BootMacAddress: to.Ptr("00:BB:CC:DD:EE:FF"),
	// 					MachineDetails: to.Ptr("extraDetails"),
	// 					MachineName: to.Ptr("bmmName1"),
	// 					RackSlot: to.Ptr[int64](1),
	// 					SerialNumber: to.Ptr("BM1219XXX"),
	// 				},
	// 				{
	// 					BmcConnectionString: to.Ptr("bmcConnectionString"),
	// 					BmcCredentials: &armnetworkcloud.AdministrativeCredentials{
	// 						Username: to.Ptr("username"),
	// 					},
	// 					BmcMacAddress: to.Ptr("AA:BB:CC:DD:EE:00"),
	// 					BootMacAddress: to.Ptr("00:BB:CC:DD:EE:00"),
	// 					MachineDetails: to.Ptr("extraDetails"),
	// 					MachineName: to.Ptr("bmmName2"),
	// 					RackSlot: to.Ptr[int64](2),
	// 					SerialNumber: to.Ptr("BM1219YYY"),
	// 			}},
	// 			NetworkRackID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName"),
	// 			RackLocation: to.Ptr("Foo Datacenter, Floor 3, Aisle 9, Rack 2"),
	// 			RackSerialNumber: to.Ptr("newSerialNumber"),
	// 			RackSKUID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"),
	// 			StorageApplianceConfigurationData: []*armnetworkcloud.StorageApplianceConfigurationData{
	// 				{
	// 					AdminCredentials: &armnetworkcloud.AdministrativeCredentials{
	// 						Username: to.Ptr("username"),
	// 					},
	// 					RackSlot: to.Ptr[int64](1),
	// 					SerialNumber: to.Ptr("BM1219XXX"),
	// 					StorageApplianceName: to.Ptr("vmName"),
	// 			}},
	// 		},
	// 		AnalyticsWorkspaceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/microsoft.operationalInsights/workspaces/logAnalyticsWorkspaceName"),
	// 		AvailableUpgradeVersions: []*armnetworkcloud.ClusterAvailableUpgradeVersion{
	// 			{
	// 				ControlImpact: to.Ptr(armnetworkcloud.ControlImpactFalse),
	// 				ExpectedDuration: to.Ptr("0:0:30"),
	// 				ImpactDescription: to.Ptr("can be done in place"),
	// 				SupportExpiryDate: to.Ptr("2025-01-01"),
	// 				TargetClusterVersion: to.Ptr("1.0.2"),
	// 				WorkloadImpact: to.Ptr(armnetworkcloud.WorkloadImpactFalse),
	// 		}},
	// 		ClusterCapacity: &armnetworkcloud.ClusterCapacity{
	// 			AvailableApplianceStorageGB: to.Ptr[int64](3),
	// 			AvailableCoreCount: to.Ptr[int64](10),
	// 			AvailableHostStorageGB: to.Ptr[int64](20),
	// 			AvailableMemoryGB: to.Ptr[int64](20),
	// 			TotalApplianceStorageGB: to.Ptr[int64](10),
	// 			TotalCoreCount: to.Ptr[int64](10),
	// 			TotalHostStorageGB: to.Ptr[int64](10),
	// 			TotalMemoryGB: to.Ptr[int64](10),
	// 		},
	// 		ClusterConnectionStatus: to.Ptr(armnetworkcloud.ClusterConnectionStatusConnected),
	// 		ClusterExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 			Type: to.Ptr("CustomLocation"),
	// 		},
	// 		ClusterLocation: to.Ptr("Foo Street, 3rd Floor, row 9"),
	// 		ClusterManagerConnectionStatus: to.Ptr(armnetworkcloud.ClusterManagerConnectionStatusConnected),
	// 		ClusterManagerID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusterManagers/clusterManagerName"),
	// 		ClusterServicePrincipal: &armnetworkcloud.ServicePrincipalInformation{
	// 			ApplicationID: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 			PrincipalID: to.Ptr("00000008-0004-0004-0004-000000000012"),
	// 			TenantID: to.Ptr("80000000-4000-4000-4000-120000000000"),
	// 		},
	// 		ClusterType: to.Ptr(armnetworkcloud.ClusterTypeSingleRack),
	// 		ClusterVersion: to.Ptr("1.0.0"),
	// 		ComputeDeploymentThreshold: &armnetworkcloud.ValidationThreshold{
	// 			Type: to.Ptr(armnetworkcloud.ValidationThresholdTypePercentSuccess),
	// 			Grouping: to.Ptr(armnetworkcloud.ValidationThresholdGroupingPerCluster),
	// 			Value: to.Ptr[int64](90),
	// 		},
	// 		ComputeRackDefinitions: []*armnetworkcloud.RackDefinition{
	// 			{
	// 				BareMetalMachineConfigurationData: []*armnetworkcloud.BareMetalMachineConfigurationData{
	// 					{
	// 						BmcConnectionString: to.Ptr("bmcConnectionString"),
	// 						BmcCredentials: &armnetworkcloud.AdministrativeCredentials{
	// 							Username: to.Ptr("username"),
	// 						},
	// 						BmcMacAddress: to.Ptr("AA:BB:CC:DD:EE:FF"),
	// 						BootMacAddress: to.Ptr("00:BB:CC:DD:EE:FF"),
	// 						MachineDetails: to.Ptr("extraDetails"),
	// 						MachineName: to.Ptr("bmmName1"),
	// 						RackSlot: to.Ptr[int64](1),
	// 						SerialNumber: to.Ptr("BM1219XXX"),
	// 					},
	// 					{
	// 						BmcConnectionString: to.Ptr("bmcConnectionString"),
	// 						BmcCredentials: &armnetworkcloud.AdministrativeCredentials{
	// 							Username: to.Ptr("username"),
	// 						},
	// 						BmcMacAddress: to.Ptr("AA:BB:CC:DD:EE:00"),
	// 						BootMacAddress: to.Ptr("00:BB:CC:DD:EE:00"),
	// 						MachineDetails: to.Ptr("extraDetails"),
	// 						MachineName: to.Ptr("bmmName2"),
	// 						RackSlot: to.Ptr[int64](2),
	// 						SerialNumber: to.Ptr("BM1219YYY"),
	// 				}},
	// 				NetworkRackID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName"),
	// 				RackLocation: to.Ptr("Foo Datacenter, Floor 3, Aisle 9, Rack 2"),
	// 				RackSerialNumber: to.Ptr("AA1234"),
	// 				RackSKUID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"),
	// 				StorageApplianceConfigurationData: []*armnetworkcloud.StorageApplianceConfigurationData{
	// 					{
	// 						AdminCredentials: &armnetworkcloud.AdministrativeCredentials{
	// 							Username: to.Ptr("username"),
	// 						},
	// 						RackSlot: to.Ptr[int64](1),
	// 						SerialNumber: to.Ptr("BM1219XXX"),
	// 						StorageApplianceName: to.Ptr("vmName"),
	// 				}},
	// 		}},
	// 		DetailedStatus: to.Ptr(armnetworkcloud.ClusterDetailedStatusRunning),
	// 		DetailedStatusMessage: to.Ptr("Cluster is running and healthy"),
	// 		ManagedResourceGroupConfiguration: &armnetworkcloud.ManagedResourceGroupConfiguration{
	// 			Name: to.Ptr("my-managed-rg"),
	// 			Location: to.Ptr("East US"),
	// 		},
	// 		ManualActionCount: to.Ptr[int64](0),
	// 		NetworkFabricID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabrics/fabricName"),
	// 		ProvisioningState: to.Ptr(armnetworkcloud.ClusterProvisioningStateSucceeded),
	// 		SupportExpiryDate: to.Ptr("2023-04-29"),
	// 		WorkloadResourceIDs: []*string{
	// 			to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName")},
	// 		},
	// 	}
}

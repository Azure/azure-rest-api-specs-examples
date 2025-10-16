package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
)

// Generated from example definition: 2025-09-01/CloudVmClusters_AddVms_MaximumSet_Gen.json
func ExampleCloudVMClustersClient_BeginAddVMs_addVMSToVMClusterGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudVMClustersClient().BeginAddVMs(ctx, "rgopenapi", "cloudvmcluster1", armoracledatabase.AddRemoveDbNode{
		DbServers: []*string{
			to.Ptr("ocid1..aaaa"),
			to.Ptr("ocid1..aaaaaa"),
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
	// res = armoracledatabase.CloudVMClustersClientAddVMsResponse{
	// 	CloudVMCluster: &armoracledatabase.CloudVMCluster{
	// 		Properties: &armoracledatabase.CloudVMClusterProperties{
	// 			DataStorageSizeInTbs: to.Ptr[float64](1000),
	// 			DbNodeStorageSizeInGbs: to.Ptr[int32](1000),
	// 			MemorySizeInGbs: to.Ptr[int32](1000),
	// 			TimeZone: to.Ptr("UTC"),
	// 			Hostname: to.Ptr("hostname1"),
	// 			Domain: to.Ptr("domain1"),
	// 			CPUCoreCount: to.Ptr[int32](2),
	// 			OcpuCount: to.Ptr[float32](3),
	// 			ClusterName: to.Ptr("cluster1"),
	// 			DataStoragePercentage: to.Ptr[int32](100),
	// 			IsLocalBackupEnabled: to.Ptr(true),
	// 			CloudExadataInfrastructureID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
	// 			IsSparseDiskgroupEnabled: to.Ptr(true),
	// 			SSHPublicKeys: []*string{
	// 				to.Ptr("ssh-key 1"),
	// 			},
	// 			NsgCidrs: []*armoracledatabase.NsgCidr{
	// 				{
	// 					Source: to.Ptr("10.0.0.0/16"),
	// 					DestinationPortRange: &armoracledatabase.PortRange{
	// 						Min: to.Ptr[int32](1520),
	// 						Max: to.Ptr[int32](1522),
	// 					},
	// 				},
	// 				{
	// 					Source: to.Ptr("10.10.0.0/24"),
	// 					DestinationPortRange: &armoracledatabase.PortRange{
	// 						Min: to.Ptr[int32](9434),
	// 						Max: to.Ptr[int32](11996),
	// 					},
	// 				},
	// 			},
	// 			LicenseModel: to.Ptr(armoracledatabase.LicenseModelLicenseIncluded),
	// 			ScanListenerPortTCP: to.Ptr[int32](1050),
	// 			ScanListenerPortTCPSSL: to.Ptr[int32](1025),
	// 			VnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 			GiVersion: to.Ptr("19.0.0.0"),
	// 			SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 			BackupSubnetCidr: to.Ptr("172.17.5.0/24"),
	// 			DataCollectionOptions: &armoracledatabase.DataCollectionOptions{
	// 				IsDiagnosticsEventsEnabled: to.Ptr(true),
	// 				IsHealthMonitoringEnabled: to.Ptr(true),
	// 				IsIncidentLogsEnabled: to.Ptr(true),
	// 			},
	// 			DisplayName: to.Ptr("cluster 1"),
	// 			DbServers: []*string{
	// 				to.Ptr("ocid1..aaaa"),
	// 			},
	// 			Ocid: to.Ptr("ocid1..aaaa"),
	// 			ListenerPort: to.Ptr[int64](1050),
	// 			NodeCount: to.Ptr[int32](100),
	// 			StorageSizeInGbs: to.Ptr[int32](1000),
	// 			FileSystemConfigurationDetails: []*armoracledatabase.FileSystemConfigurationDetails{
	// 				{
	// 					MountPoint: to.Ptr("gukfhjlmkqfqdgb"),
	// 					FileSystemSizeGb: to.Ptr[int32](20),
	// 				},
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-22T02:18:35.683Z"); return t}()),
	// 			LifecycleDetails: to.Ptr("success"),
	// 			ZoneID: to.Ptr("ocid1..aaaa"),
	// 			SystemVersion: to.Ptr("v1"),
	// 			DiskRedundancy: to.Ptr(armoracledatabase.DiskRedundancyHigh),
	// 			ScanIPIDs: []*string{
	// 				to.Ptr("10.0.0.1"),
	// 			},
	// 			VipIDs: []*string{
	// 				to.Ptr("10.0.1.3"),
	// 			},
	// 			ScanDNSName: to.Ptr("dbdns1"),
	// 			ScanDNSRecordID: to.Ptr("scandns1"),
	// 			Shape: to.Ptr("EXADATA.X9M"),
	// 			ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 			LifecycleState: to.Ptr(armoracledatabase.CloudVMClusterLifecycleStateProvisioning),
	// 			OciURL: to.Ptr("https://fake"),
	// 			NsgURL: to.Ptr("https://microsoft.com/a"),
	// 			IormConfigCache: &armoracledatabase.ExadataIormConfig{
	// 				DbPlans: []*armoracledatabase.DbIormConfig{
	// 					{
	// 						DbName: to.Ptr("db1"),
	// 						FlashCacheLimit: to.Ptr("none"),
	// 						Share: to.Ptr[int32](32),
	// 					},
	// 				},
	// 				LifecycleDetails: to.Ptr("Disabled"),
	// 				LifecycleState: to.Ptr(armoracledatabase.IormLifecycleStateDisabled),
	// 				Objective: to.Ptr(armoracledatabase.ObjectiveLowLatency),
	// 			},
	// 			LastUpdateHistoryEntryID: to.Ptr("none"),
	// 			CompartmentID: to.Ptr("ocid1..aaaaaa"),
	// 			SubnetOcid: to.Ptr("ocid1..aaaaaa"),
	// 			ComputeModel: to.Ptr(armoracledatabase.ComputeModelECPU),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudVmClusters/cluster1"),
	// 		Name: to.Ptr("cluster1"),
	// 		Type: to.Ptr("Oracle.Database/cloudVmClusters"),
	// 		SystemData: &armoracledatabase.SystemData{
	// 			CreatedBy: to.Ptr("sqehacivpuim"),
	// 			CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("axrqfdkqylvjv"),
	// 			LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
	// 		},
	// 	},
	// }
}

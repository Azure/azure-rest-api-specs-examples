package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: 2025-03-01/vmClusters_addVms.json
func ExampleCloudVMClustersClient_BeginAddVMs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudVMClustersClient().BeginAddVMs(ctx, "rg000", "cluster1", armoracledatabase.AddRemoveDbNode{
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
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudVmClusters/cluster1"),
	// 		Type: to.Ptr("Oracle.Database/cloudVmClusters"),
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 			"tagK1": to.Ptr("tagV1"),
	// 		},
	// 		Properties: &armoracledatabase.CloudVMClusterProperties{
	// 			Ocid: to.Ptr("ocid1..aaaa"),
	// 			ListenerPort: to.Ptr[int64](1050),
	// 			NodeCount: to.Ptr[int32](100),
	// 			StorageSizeInGbs: to.Ptr[int32](1000),
	// 			DataStorageSizeInTbs: to.Ptr[float64](10),
	// 			DbNodeStorageSizeInGbs: to.Ptr[int32](100),
	// 			MemorySizeInGbs: to.Ptr[int32](1000),
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-22T02:18:35.683Z"); return t}()),
	// 			LifecycleDetails: to.Ptr("success"),
	// 			TimeZone: to.Ptr("UTC"),
	// 			ZoneID: to.Ptr("ocid1..aaaa"),
	// 			Hostname: to.Ptr("hostname1"),
	// 			Domain: to.Ptr("domain1"),
	// 			CPUCoreCount: to.Ptr[int32](10),
	// 			OcpuCount: to.Ptr[float32](100),
	// 			ClusterName: to.Ptr("cluster1"),
	// 			DataStoragePercentage: to.Ptr[int32](80),
	// 			IsLocalBackupEnabled: to.Ptr(false),
	// 			CloudExadataInfrastructureID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
	// 			IsSparseDiskgroupEnabled: to.Ptr(false),
	// 			SSHPublicKeys: []*string{
	// 				to.Ptr("ssh-key 1"),
	// 			},
	// 			LicenseModel: to.Ptr(armoracledatabase.LicenseModelLicenseIncluded),
	// 			ScanListenerPortTCP: to.Ptr[int32](1050),
	// 			ScanListenerPortTCPSSL: to.Ptr[int32](1025),
	// 			VnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 			GiVersion: to.Ptr("19.0.0.0"),
	// 			SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
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
	// 			OciURL: to.Ptr("https://fake"),
	// 			DataCollectionOptions: &armoracledatabase.DataCollectionOptions{
	// 				IsDiagnosticsEventsEnabled: to.Ptr(false),
	// 				IsHealthMonitoringEnabled: to.Ptr(false),
	// 				IsIncidentLogsEnabled: to.Ptr(false),
	// 			},
	// 			DisplayName: to.Ptr("cluster 1"),
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
	// 			DbServers: []*string{
	// 				to.Ptr("ocid1..aaaa"),
	// 			},
	// 			CompartmentID: to.Ptr("ocid1..aaaaaa"),
	// 			SubnetOcid: to.Ptr("ocid1..aaaaaa"),
	// 		},
	// 	},
	// }
}

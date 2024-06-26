package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ee6d9fd7687d4b67117c5a167c191a7e7e70b53/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/vmClusters_listByResourceGroup.json
func ExampleCloudVMClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudVMClustersClient().NewListByResourceGroupPager("rg000", nil)
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
		// page.CloudVMClusterListResult = armoracledatabase.CloudVMClusterListResult{
		// 	Value: []*armoracledatabase.CloudVMCluster{
		// 		{
		// 			Type: to.Ptr("Oracle.Database/cloudVmClusters"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudVmClusters/cluster1"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"tagK1": to.Ptr("tagV1"),
		// 			},
		// 			Properties: &armoracledatabase.CloudVMClusterProperties{
		// 				BackupSubnetCidr: to.Ptr("172.17.5.0/24"),
		// 				CloudExadataInfrastructureID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
		// 				ClusterName: to.Ptr("cluster1"),
		// 				CompartmentID: to.Ptr("ocid1..aaaaaa"),
		// 				CPUCoreCount: to.Ptr[int32](10),
		// 				DataCollectionOptions: &armoracledatabase.DataCollectionOptions{
		// 					IsDiagnosticsEventsEnabled: to.Ptr(false),
		// 					IsHealthMonitoringEnabled: to.Ptr(false),
		// 					IsIncidentLogsEnabled: to.Ptr(false),
		// 				},
		// 				DataStoragePercentage: to.Ptr[int32](80),
		// 				DataStorageSizeInTbs: to.Ptr[float64](10),
		// 				DbNodeStorageSizeInGbs: to.Ptr[int32](100),
		// 				DbServers: []*string{
		// 					to.Ptr("ocid1..aaaa")},
		// 					DiskRedundancy: to.Ptr(armoracledatabase.DiskRedundancyHigh),
		// 					DisplayName: to.Ptr("cluster 1"),
		// 					Domain: to.Ptr("domain1"),
		// 					GiVersion: to.Ptr("19.0.0.0"),
		// 					Hostname: to.Ptr("hostname1"),
		// 					IormConfigCache: &armoracledatabase.ExadataIormConfig{
		// 						DbPlans: []*armoracledatabase.DbIormConfig{
		// 							{
		// 								DbName: to.Ptr("db1"),
		// 								FlashCacheLimit: to.Ptr("none"),
		// 								Share: to.Ptr[int32](32),
		// 						}},
		// 						LifecycleDetails: to.Ptr("Disabled"),
		// 						LifecycleState: to.Ptr(armoracledatabase.IormLifecycleStateDisabled),
		// 						Objective: to.Ptr(armoracledatabase.ObjectiveLowLatency),
		// 					},
		// 					IsLocalBackupEnabled: to.Ptr(false),
		// 					IsSparseDiskgroupEnabled: to.Ptr(false),
		// 					LastUpdateHistoryEntryID: to.Ptr("none"),
		// 					LicenseModel: to.Ptr(armoracledatabase.LicenseModelLicenseIncluded),
		// 					LifecycleDetails: to.Ptr("success"),
		// 					ListenerPort: to.Ptr[int64](1050),
		// 					MemorySizeInGbs: to.Ptr[int32](1000),
		// 					NodeCount: to.Ptr[int32](100),
		// 					OciURL: to.Ptr("https://fake"),
		// 					Ocid: to.Ptr("ocid1..aaaa"),
		// 					OcpuCount: to.Ptr[float32](100),
		// 					ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 					ScanDNSName: to.Ptr("dbdns1"),
		// 					ScanDNSRecordID: to.Ptr("scandns1"),
		// 					ScanIPIDs: []*string{
		// 						to.Ptr("10.0.0.1")},
		// 						ScanListenerPortTCP: to.Ptr[int32](1050),
		// 						ScanListenerPortTCPSSL: to.Ptr[int32](1025),
		// 						Shape: to.Ptr("EXADATA.X9M"),
		// 						SSHPublicKeys: []*string{
		// 							to.Ptr("ssh-key 1")},
		// 							StorageSizeInGbs: to.Ptr[int32](1000),
		// 							SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
		// 							SubnetOcid: to.Ptr("ocid1..aaaaaa"),
		// 							SystemVersion: to.Ptr("v1"),
		// 							TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-22T02:18:35.683Z"); return t}()),
		// 							TimeZone: to.Ptr("UTC"),
		// 							VipIDs: []*string{
		// 								to.Ptr("10.0.1.3")},
		// 								VnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
		// 								ZoneID: to.Ptr("ocid1..aaaa"),
		// 							},
		// 					}},
		// 				}
	}
}

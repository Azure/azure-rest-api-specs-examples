package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: 2025-03-01/ExadbVmClusters_RemoveVms_MaximumSet_Gen.json
func ExampleExadbVMClustersClient_BeginRemoveVMs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExadbVMClustersClient().BeginRemoveVMs(ctx, "rgopenapi", "vmClusterName", armoracledatabase.RemoveVirtualMachineFromExadbVMClusterDetails{
		DbNodes: []*armoracledatabase.DbNodeDetails{
			{
				DbNodeID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Oracle.Database/exadbVmClusters/vmCluster/dbNodes/dbNodeName"),
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
	// res = armoracledatabase.ExadbVMClustersClientRemoveVMsResponse{
	// 	ExadbVMCluster: &armoracledatabase.ExadbVMCluster{
	// 		Properties: &armoracledatabase.ExadbVMClusterProperties{
	// 			Ocid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 			ClusterName: to.Ptr("hssswln"),
	// 			BackupSubnetCidr: to.Ptr("fr"),
	// 			NsgURL: to.Ptr("https://microsoft.com/a"),
	// 			ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 			LifecycleState: to.Ptr(armoracledatabase.ExadbVMClusterLifecycleStateProvisioning),
	// 			VnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 			SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 			DataCollectionOptions: &armoracledatabase.DataCollectionOptions{
	// 				IsDiagnosticsEventsEnabled: to.Ptr(true),
	// 				IsHealthMonitoringEnabled: to.Ptr(true),
	// 				IsIncidentLogsEnabled: to.Ptr(true),
	// 			},
	// 			DisplayName: to.Ptr("dyhkncxzzgcwlfkfygvqlvqbxekvkbhjprmuxtaukewjnfrrnfqgrqsjfgsayrezkspqplrduuomuvtpkqurcoqjdnadlvedgfngglcgafbxtcewxlgckvehhqgfwkokbjoqftunjsgfjwbdaftxoytsphydwogtlnfxxuzzjvqcrin"),
	// 			Domain: to.Ptr("vzfak"),
	// 			EnabledEcpuCount: to.Ptr[int32](0),
	// 			ExascaleDbStorageVaultID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/exascaleDbStorageVaults/storageVaultName"),
	// 			GridImageOcid: to.Ptr("ocid1.dbpatch.oc1..aaaaa3klq"),
	// 			GridImageType: to.Ptr(armoracledatabase.GridImageTypeReleaseUpdate),
	// 			GiVersion: to.Ptr("mwuazccgbujljfeqjpttmwqmppliba"),
	// 			Hostname: to.Ptr("tdn"),
	// 			LicenseModel: to.Ptr(armoracledatabase.LicenseModelLicenseIncluded),
	// 			MemorySizeInGbs: to.Ptr[int32](6),
	// 			NodeCount: to.Ptr[int32](8),
	// 			NsgCidrs: []*armoracledatabase.NsgCidr{
	// 				{
	// 					Source: to.Ptr("10.0.0.0/16"),
	// 					DestinationPortRange: &armoracledatabase.PortRange{
	// 						Min: to.Ptr[int32](1520),
	// 						Max: to.Ptr[int32](1522),
	// 					},
	// 				},
	// 			},
	// 			PrivateZoneOcid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 			ScanListenerPortTCP: to.Ptr[int32](9),
	// 			ScanListenerPortTCPSSL: to.Ptr[int32](18),
	// 			ListenerPort: to.Ptr[int32](20),
	// 			Shape: to.Ptr("evenjblxxfsztdxlcr"),
	// 			SSHPublicKeys: []*string{
	// 				to.Ptr("wiseakllkohemfmnoyteg"),
	// 			},
	// 			SystemVersion: to.Ptr("udwsykstqawwjrvzebfyyccnjzhxtijvywlgnrkqxnboibyppistnbvqqsnsxmjnzeumgnatmarzfnjhglmtrknszthrxtwewwqjbkwytbhtjbgondtktkpvjmxdcuyjupr"),
	// 			TimeZone: to.Ptr("ayuqnpfkasdzmxlfnkjzrenyogbvcysdbkstjhyjdgxygitlykihtwdiktilukyplgxovaonrqcqaviyaqgyrqfrklqvqyobnxgmzpqbgjhdtjdquhyqnvhavnqpfupaqhdedgdvfomeueeuwjjfiqiyxspybpyj"),
	// 			TotalEcpuCount: to.Ptr[int32](56),
	// 			VMFileSystemStorage: &armoracledatabase.ExadbVMClusterStorageDetails{
	// 				TotalSizeInGbs: to.Ptr[int32](10),
	// 			},
	// 			LifecycleDetails: to.Ptr("tl"),
	// 			ScanDNSName: to.Ptr("ya"),
	// 			ScanIPIDs: []*string{
	// 				to.Ptr("hdjuwrxrykhbuhtehlwm"),
	// 			},
	// 			ScanDNSRecordID: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 			SnapshotFileSystemStorage: &armoracledatabase.ExadbVMClusterStorageDetails{
	// 				TotalSizeInGbs: to.Ptr[int32](10),
	// 			},
	// 			TotalFileSystemStorage: &armoracledatabase.ExadbVMClusterStorageDetails{
	// 				TotalSizeInGbs: to.Ptr[int32](10),
	// 			},
	// 			VipIDs: []*string{
	// 				to.Ptr("fleeucwejtkaxnquq"),
	// 			},
	// 			ZoneOcid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 			OciURL: to.Ptr("https://microsoft.com/a"),
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
	// 			BackupSubnetOcid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 			SubnetOcid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("iupjvcfsbvfhrebwf"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key9683": to.Ptr("umnfswbmhxzsomcluslxhnxxika"),
	// 		},
	// 		Location: to.Ptr("xojp"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Oracle.Database/exadbVmClusters/vmCluster"),
	// 		Name: to.Ptr("joekjadzkdyqh"),
	// 		Type: to.Ptr("ugekszbnqazqsofdniamlclolkcsi"),
	// 		SystemData: &armoracledatabase.SystemData{
	// 			CreatedBy: to.Ptr("ilrpjodjmvzhybazxipoplnql"),
	// 			CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("lhjbxchqkaia"),
	// 			LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
	// 		},
	// 	},
	// }
}

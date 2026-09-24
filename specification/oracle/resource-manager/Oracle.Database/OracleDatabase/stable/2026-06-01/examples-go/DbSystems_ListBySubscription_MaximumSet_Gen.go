package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v3"
)

// Generated from example definition: 2026-06-01/DbSystems_ListBySubscription_MaximumSet_Gen.json
func ExampleDbSystemsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDbSystemsClient().NewListBySubscriptionPager(nil)
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
		// page = armoracledatabase.DbSystemsClientListBySubscriptionResponse{
		// 	DbSystemListResult: armoracledatabase.DbSystemListResult{
		// 		Value: []*armoracledatabase.DbSystem{
		// 			{
		// 				Properties: &armoracledatabase.DbSystemProperties{
		// 					DatabaseEdition: to.Ptr(armoracledatabase.DbSystemDatabaseEditionTypeStandardEdition),
		// 					DbVersion: to.Ptr("hodeobcqsjy"),
		// 					Source: to.Ptr(armoracledatabase.DbSystemSourceTypeNone),
		// 					ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 					OciURL: to.Ptr("https://microsoft.com/a"),
		// 					ResourceAnchorID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resourceAnchors/anchor1"),
		// 					NetworkAnchorID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/networkAnchors/networkanchor1"),
		// 					ClusterName: to.Ptr("itawczrpsu"),
		// 					DisplayName: to.Ptr("resource1"),
		// 					DataStorageSizeInGbs: to.Ptr[int32](4),
		// 					DbSystemOptions: &armoracledatabase.DbSystemOptions{
		// 						StorageManagement: to.Ptr(armoracledatabase.StorageManagementTypeLVM),
		// 					},
		// 					DiskRedundancy: to.Ptr(armoracledatabase.DiskRedundancyTypeHigh),
		// 					DomainV2: to.Ptr("example"),
		// 					GridImageOcid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					Hostname: to.Ptr("resource1"),
		// 					Ocid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					LicenseModelV2: to.Ptr(armoracledatabase.LicenseModelLicenseIncluded),
		// 					LifecycleDetails: to.Ptr("example"),
		// 					LifecycleState: to.Ptr(armoracledatabase.DbSystemLifecycleStateProvisioning),
		// 					ListenerPort: to.Ptr[int32](17),
		// 					MemorySizeInGbs: to.Ptr[int32](9),
		// 					NodeCount: to.Ptr[int32](10),
		// 					ScanDNSName: to.Ptr("resource1"),
		// 					ScanIPs: []*string{
		// 						to.Ptr("example"),
		// 					},
		// 					Shape: to.Ptr("example"),
		// 					SSHPublicKeys: []*string{
		// 						to.Ptr("example"),
		// 					},
		// 					StorageVolumePerformanceMode: to.Ptr(armoracledatabase.StorageVolumePerformanceModeBalanced),
		// 					TimeZone: to.Ptr("2026-06-01T00:00:00Z"),
		// 					Version: to.Ptr("kphhahctsuk"),
		// 					ComputeModel: to.Ptr(armoracledatabase.ComputeModelECPU),
		// 					ComputeCount: to.Ptr[int32](19),
		// 					DataCollectionOptions: &armoracledatabase.DataCollectionOptions{
		// 						IsDiagnosticsEventsEnabled: to.Ptr(true),
		// 						IsHealthMonitoringEnabled: to.Ptr(true),
		// 						IsIncidentLogsEnabled: to.Ptr(true),
		// 					},
		// 					CharacterSet: to.Ptr("example"),
		// 					NcharacterSet: to.Ptr("example"),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("example"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key3781": to.Ptr("y"),
		// 				},
		// 				Location: to.Ptr("kvdtzeiu"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resources/resource1"),
		// 				Name: to.Ptr("resource1"),
		// 				Type: to.Ptr("Oracle.Database/resource"),
		// 				SystemData: &armoracledatabase.SystemData{
		// 					CreatedBy: to.Ptr("example"),
		// 					CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2026, time.July, 28, 22, 26, 48, 756000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("example"),
		// 					LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2026, time.July, 28, 22, 26, 48, 756000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

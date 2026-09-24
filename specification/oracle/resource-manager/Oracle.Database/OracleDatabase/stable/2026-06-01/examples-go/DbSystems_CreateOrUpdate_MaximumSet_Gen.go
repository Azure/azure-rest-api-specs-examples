package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v3"
)

// Generated from example definition: 2026-06-01/DbSystems_CreateOrUpdate_MaximumSet_Gen.json
func ExampleDbSystemsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDbSystemsClient().BeginCreateOrUpdate(ctx, "rgopenapi", "resource1", armoracledatabase.DbSystem{
		Properties: &armoracledatabase.DbSystemProperties{
			DatabaseEdition:            to.Ptr(armoracledatabase.DbSystemDatabaseEditionTypeStandardEdition),
			AdminPassword:              to.Ptr("********"),
			DbVersion:                  to.Ptr("example"),
			PdbName:                    to.Ptr("resource1"),
			Source:                     to.Ptr(armoracledatabase.DbSystemSourceTypeNone),
			ResourceAnchorID:           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resourceAnchors/anchor1"),
			NetworkAnchorID:            to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/networkAnchors/networkanchor1"),
			ClusterName:                to.Ptr("puw"),
			DisplayName:                to.Ptr("resource1"),
			InitialDataStorageSizeInGb: to.Ptr[int32](40),
			DbSystemOptions: &armoracledatabase.DbSystemOptions{
				StorageManagement: to.Ptr(armoracledatabase.StorageManagementTypeLVM),
			},
			DiskRedundancy: to.Ptr(armoracledatabase.DiskRedundancyTypeHigh),
			DomainV2:       to.Ptr("l"),
			GridImageOcid:  to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
			Hostname:       to.Ptr("b"),
			Ocid:           to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
			LicenseModelV2: to.Ptr(armoracledatabase.LicenseModelLicenseIncluded),
			LifecycleState: to.Ptr(armoracledatabase.DbSystemLifecycleStateProvisioning),
			NodeCount:      to.Ptr[int32](11),
			Shape:          to.Ptr("example"),
			SSHPublicKeys: []*string{
				to.Ptr("example"),
			},
			StorageVolumePerformanceMode: to.Ptr(armoracledatabase.StorageVolumePerformanceModeBalanced),
			TimeZone:                     to.Ptr("2026-06-01T00:00:00Z"),
			ComputeModel:                 to.Ptr(armoracledatabase.ComputeModelECPU),
			ComputeCount:                 to.Ptr[int32](10),
			DataCollectionOptions: &armoracledatabase.DataCollectionOptions{
				IsDiagnosticsEventsEnabled: to.Ptr(true),
				IsHealthMonitoringEnabled:  to.Ptr(true),
				IsIncidentLogsEnabled:      to.Ptr(true),
			},
			CharacterSet:  to.Ptr("example"),
			NcharacterSet: to.Ptr("fkdieg"),
		},
		Zones: []*string{
			to.Ptr("example"),
		},
		Tags: map[string]*string{
			"key1855": to.Ptr("hczjcgfrxqk"),
		},
		Location: to.Ptr("eastus"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armoracledatabase.DbSystemsClientCreateOrUpdateResponse{
	// 	DbSystem: armoracledatabase.DbSystem{
	// 		Properties: &armoracledatabase.DbSystemProperties{
	// 			DatabaseEdition: to.Ptr(armoracledatabase.DbSystemDatabaseEditionTypeStandardEdition),
	// 			DbVersion: to.Ptr("example"),
	// 			Source: to.Ptr(armoracledatabase.DbSystemSourceTypeNone),
	// 			ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 			OciURL: to.Ptr("https://microsoft.com/a"),
	// 			ResourceAnchorID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resourceAnchors/anchor1"),
	// 			NetworkAnchorID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/networkAnchors/networkanchor1"),
	// 			ClusterName: to.Ptr("puw"),
	// 			DisplayName: to.Ptr("resource1"),
	// 			DataStorageSizeInGbs: to.Ptr[int32](14),
	// 			DbSystemOptions: &armoracledatabase.DbSystemOptions{
	// 				StorageManagement: to.Ptr(armoracledatabase.StorageManagementTypeLVM),
	// 			},
	// 			DiskRedundancy: to.Ptr(armoracledatabase.DiskRedundancyTypeHigh),
	// 			DomainV2: to.Ptr("l"),
	// 			GridImageOcid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 			Hostname: to.Ptr("b"),
	// 			Ocid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 			LicenseModelV2: to.Ptr(armoracledatabase.LicenseModelLicenseIncluded),
	// 			LifecycleDetails: to.Ptr("tmbpykeb"),
	// 			LifecycleState: to.Ptr(armoracledatabase.DbSystemLifecycleStateProvisioning),
	// 			ListenerPort: to.Ptr[int32](16),
	// 			MemorySizeInGbs: to.Ptr[int32](17),
	// 			NodeCount: to.Ptr[int32](11),
	// 			ScanDNSName: to.Ptr("resource1"),
	// 			ScanIPs: []*string{
	// 				to.Ptr("example"),
	// 			},
	// 			Shape: to.Ptr("example"),
	// 			SSHPublicKeys: []*string{
	// 				to.Ptr("example"),
	// 			},
	// 			StorageVolumePerformanceMode: to.Ptr(armoracledatabase.StorageVolumePerformanceModeBalanced),
	// 			TimeZone: to.Ptr("2026-06-01T00:00:00Z"),
	// 			Version: to.Ptr("avgqty"),
	// 			ComputeModel: to.Ptr(armoracledatabase.ComputeModelECPU),
	// 			ComputeCount: to.Ptr[int32](10),
	// 			DataCollectionOptions: &armoracledatabase.DataCollectionOptions{
	// 				IsDiagnosticsEventsEnabled: to.Ptr(true),
	// 				IsHealthMonitoringEnabled: to.Ptr(true),
	// 				IsIncidentLogsEnabled: to.Ptr(true),
	// 			},
	// 			CharacterSet: to.Ptr("example"),
	// 			NcharacterSet: to.Ptr("fkdieg"),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("example"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1855": to.Ptr("hczjcgfrxqk"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resources/resource1"),
	// 		Name: to.Ptr("kqnjrlmq"),
	// 		Type: to.Ptr("Oracle.Database/resource"),
	// 		SystemData: &armoracledatabase.SystemData{
	// 			CreatedBy: to.Ptr("ns"),
	// 			CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("example"),
	// 			LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
	// 		},
	// 	},
	// }
}

package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
)

// Generated from example definition: 2025-09-01/autonomousDatabase_changeDisasterRecoveryConfiguration.json
func ExampleAutonomousDatabasesClient_BeginChangeDisasterRecoveryConfiguration_autonomousDatabasesChangeDisasterRecoveryConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutonomousDatabasesClient().BeginChangeDisasterRecoveryConfiguration(ctx, "rg000", "databasedb1", armoracledatabase.DisasterRecoveryConfigurationDetails{
		DisasterRecoveryType:        to.Ptr(armoracledatabase.DisasterRecoveryTypeAdg),
		IsReplicateAutomaticBackups: to.Ptr(false),
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
	// res = armoracledatabase.AutonomousDatabasesClientChangeDisasterRecoveryConfigurationResponse{
	// 	AutonomousDatabase: &armoracledatabase.AutonomousDatabase{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1"),
	// 		Type: to.Ptr("Oracle.Database/autonomousDatabases"),
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 			"tagK1": to.Ptr("tagV1"),
	// 		},
	// 		Properties: &armoracledatabase.AutonomousDatabaseProperties{
	// 			AutonomousDatabaseID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1"),
	// 			AutonomousMaintenanceScheduleType: to.Ptr(armoracledatabase.AutonomousMaintenanceScheduleTypeRegular),
	// 			CharacterSet: to.Ptr("AL32UTF8"),
	// 			NcharacterSet: to.Ptr("AL16UTF16"),
	// 			ComputeCount: to.Ptr[float32](2),
	// 			ComputeModel: to.Ptr(armoracledatabase.ComputeModelECPU),
	// 			CPUCoreCount: to.Ptr[int32](1),
	// 			DataStorageSizeInGbs: to.Ptr[int32](1024),
	// 			DataStorageSizeInTbs: to.Ptr[int32](1),
	// 			DatabaseEdition: to.Ptr(armoracledatabase.DatabaseEditionTypeEnterpriseEdition),
	// 			DataBaseType: to.Ptr(armoracledatabase.DataBaseTypeRegular),
	// 			DbVersion: to.Ptr("19c"),
	// 			DisplayName: to.Ptr("example_autonomous_databasedb1"),
	// 			IsAutoScalingEnabled: to.Ptr(false),
	// 			IsAutoScalingForStorageEnabled: to.Ptr(false),
	// 			IsLocalDataGuardEnabled: to.Ptr(false),
	// 			IsRemoteDataGuardEnabled: to.Ptr(true),
	// 			TimeDisasterRecoveryRoleChanged: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-30T18:22:10.970Z"); return t}()),
	// 			TimeDataGuardRoleChanged: to.Ptr("2024-07-30T18:22:10.970Z"),
	// 			TimeLocalDataGuardEnabled: to.Ptr("2024-07-04T01:02:36.782Z"),
	// 			LocalDisasterRecoveryType: to.Ptr(armoracledatabase.DisasterRecoveryTypeBackupBased),
	// 			RemoteDisasterRecoveryConfiguration: &armoracledatabase.DisasterRecoveryConfigurationDetails{
	// 				DisasterRecoveryType: to.Ptr(armoracledatabase.DisasterRecoveryTypeBackupBased),
	// 				IsReplicateAutomaticBackups: to.Ptr(false),
	// 			},
	// 			Role: to.Ptr(armoracledatabase.RoleTypeBackupCopy),
	// 			PeerDbIDs: []*string{
	// 				to.Ptr("ocid1.bbbbb"),
	// 			},
	// 			IsMtlsConnectionRequired: to.Ptr(true),
	// 			LicenseModel: to.Ptr(armoracledatabase.LicenseModelBringYourOwnLicense),
	// 			LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseLifecycleStateUpdating),
	// 			SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 			VnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 			ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateProvisioning),
	// 			OciURL: to.Ptr("https://fake"),
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-20T21:20:08.070Z"); return t}()),
	// 			Ocid: to.Ptr("ocid1..aaaaa"),
	// 			WhitelistedIPs: []*string{
	// 				to.Ptr("1.1.1.1"),
	// 				to.Ptr("1.1.1.0/24"),
	// 				to.Ptr("1.1.2.25"),
	// 			},
	// 		},
	// 	},
	// }
}

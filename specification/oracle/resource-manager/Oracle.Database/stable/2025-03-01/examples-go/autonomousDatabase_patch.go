package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: 2025-03-01/autonomousDatabase_patch.json
func ExampleAutonomousDatabasesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutonomousDatabasesClient().BeginUpdate(ctx, "rg000", "databasedb1", armoracledatabase.AutonomousDatabaseUpdate{}, nil)
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
	// res = armoracledatabase.AutonomousDatabasesClientUpdateResponse{
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
	// 			DbVersion: to.Ptr("18.4.0.0"),
	// 			DisplayName: to.Ptr("example_autonomous_databasedb1"),
	// 			IsAutoScalingEnabled: to.Ptr(false),
	// 			IsAutoScalingForStorageEnabled: to.Ptr(false),
	// 			IsLocalDataGuardEnabled: to.Ptr(false),
	// 			IsMtlsConnectionRequired: to.Ptr(true),
	// 			LicenseModel: to.Ptr(armoracledatabase.LicenseModelBringYourOwnLicense),
	// 			LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseLifecycleState("Succeeded")),
	// 			LifecycleDetails: to.Ptr("success"),
	// 			SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 			VnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 			ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 			OciURL: to.Ptr("https://fake"),
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-09T20:44:09.466Z"); return t}()),
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

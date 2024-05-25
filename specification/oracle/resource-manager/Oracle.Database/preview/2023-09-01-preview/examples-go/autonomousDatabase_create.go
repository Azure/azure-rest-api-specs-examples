package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/520e274d7d95fc6d1002dd3c1fcaf8d55d27f63e/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/autonomousDatabase_create.json
func ExampleAutonomousDatabasesClient_BeginCreateOrUpdate_createAutonomousDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutonomousDatabasesClient().BeginCreateOrUpdate(ctx, "rg000", "databasedb1", armoracledatabase.AutonomousDatabase{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"tagK1": to.Ptr("tagV1"),
		},
		Properties: &armoracledatabase.AutonomousDatabaseProperties{
			AdminPassword:        to.Ptr(""),
			CharacterSet:         to.Ptr("AL32UTF8"),
			ComputeCount:         to.Ptr[float32](2),
			ComputeModel:         to.Ptr(armoracledatabase.ComputeModelECPU),
			DataBaseType:         to.Ptr(armoracledatabase.DataBaseTypeRegular),
			DataStorageSizeInTbs: to.Ptr[int32](1),
			DbVersion:            to.Ptr("18.4.0.0"),
			DisplayName:          to.Ptr("example_autonomous_databasedb1"),
			NcharacterSet:        to.Ptr("AL16UTF16"),
			SubnetID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
			VnetID:               to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
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
	// res.AutonomousDatabase = armoracledatabase.AutonomousDatabase{
	// 	Type: to.Ptr("Oracle.Database/autonomousDatabases"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tagK1": to.Ptr("tagV1"),
	// 	},
	// 	Properties: &armoracledatabase.AutonomousDatabaseProperties{
	// 		AutonomousDatabaseID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1"),
	// 		AutonomousMaintenanceScheduleType: to.Ptr(armoracledatabase.AutonomousMaintenanceScheduleTypeRegular),
	// 		CharacterSet: to.Ptr("AL32UTF8"),
	// 		ComputeCount: to.Ptr[float32](2),
	// 		ComputeModel: to.Ptr(armoracledatabase.ComputeModelECPU),
	// 		CPUCoreCount: to.Ptr[int32](1),
	// 		DataBaseType: to.Ptr(armoracledatabase.DataBaseTypeRegular),
	// 		DataStorageSizeInGbs: to.Ptr[int32](1024),
	// 		DataStorageSizeInTbs: to.Ptr[int32](1),
	// 		DatabaseEdition: to.Ptr(armoracledatabase.DatabaseEditionTypeEnterpriseEdition),
	// 		DbVersion: to.Ptr("18.4.0.0"),
	// 		DisplayName: to.Ptr("example_autonomous_databasedb1"),
	// 		IsAutoScalingEnabled: to.Ptr(false),
	// 		IsAutoScalingForStorageEnabled: to.Ptr(false),
	// 		IsLocalDataGuardEnabled: to.Ptr(false),
	// 		IsMtlsConnectionRequired: to.Ptr(true),
	// 		LicenseModel: to.Ptr(armoracledatabase.LicenseModelBringYourOwnLicense),
	// 		LifecycleDetails: to.Ptr("success"),
	// 		LifecycleState: to.Ptr(armoracledatabase.AutonomousDatabaseLifecycleState("Succeeded")),
	// 		NcharacterSet: to.Ptr("AL16UTF16"),
	// 		OciURL: to.Ptr("https://fake"),
	// 		Ocid: to.Ptr("ocid1..aaaaa"),
	// 		ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 		SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-09T20:44:09.466Z"); return t}()),
	// 		VnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 		WhitelistedIPs: []*string{
	// 			to.Ptr("1.1.1.1"),
	// 			to.Ptr("1.1.1.0/24"),
	// 			to.Ptr("1.1.2.25")},
	// 		},
	// 	}
}

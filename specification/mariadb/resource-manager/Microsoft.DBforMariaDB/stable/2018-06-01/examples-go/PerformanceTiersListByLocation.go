package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/PerformanceTiersListByLocation.json
func ExampleLocationBasedPerformanceTierClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationBasedPerformanceTierClient().NewListPager("WestUS", nil)
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
		// page.PerformanceTierListResult = armmariadb.PerformanceTierListResult{
		// 	Value: []*armmariadb.PerformanceTierProperties{
		// 		{
		// 			ID: to.Ptr("Basic"),
		// 			MaxBackupRetentionDays: to.Ptr[int32](35),
		// 			MaxLargeStorageMB: to.Ptr[int32](0),
		// 			MaxStorageMB: to.Ptr[int32](2097152),
		// 			MinBackupRetentionDays: to.Ptr[int32](7),
		// 			MinLargeStorageMB: to.Ptr[int32](0),
		// 			MinStorageMB: to.Ptr[int32](5120),
		// 			ServiceLevelObjectives: []*armmariadb.PerformanceTierServiceLevelObjectives{
		// 				{
		// 					Edition: to.Ptr("Basic"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("B_Gen5_1"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](1048576),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](1),
		// 				},
		// 				{
		// 					Edition: to.Ptr("Basic"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("B_Gen5_2"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](1048576),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](2),
		// 			}},
		// 		},
		// 		{
		// 			ID: to.Ptr("GeneralPurpose"),
		// 			MaxBackupRetentionDays: to.Ptr[int32](35),
		// 			MaxLargeStorageMB: to.Ptr[int32](16777216),
		// 			MaxStorageMB: to.Ptr[int32](16777216),
		// 			MinBackupRetentionDays: to.Ptr[int32](7),
		// 			MinLargeStorageMB: to.Ptr[int32](0),
		// 			MinStorageMB: to.Ptr[int32](5120),
		// 			ServiceLevelObjectives: []*armmariadb.PerformanceTierServiceLevelObjectives{
		// 				{
		// 					Edition: to.Ptr("GeneralPurpose"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("GP_Gen5_2"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](2),
		// 				},
		// 				{
		// 					Edition: to.Ptr("GeneralPurpose"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("GP_Gen5_4"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](4),
		// 				},
		// 				{
		// 					Edition: to.Ptr("GeneralPurpose"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("GP_Gen5_8"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](8),
		// 				},
		// 				{
		// 					Edition: to.Ptr("GeneralPurpose"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("GP_Gen5_16"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](16),
		// 				},
		// 				{
		// 					Edition: to.Ptr("GeneralPurpose"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("GP_Gen5_32"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](32),
		// 				},
		// 				{
		// 					Edition: to.Ptr("GeneralPurpose"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("GP_Gen5_64"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](32),
		// 			}},
		// 		},
		// 		{
		// 			ID: to.Ptr("MemoryOptimized"),
		// 			MaxBackupRetentionDays: to.Ptr[int32](35),
		// 			MaxLargeStorageMB: to.Ptr[int32](16777216),
		// 			MaxStorageMB: to.Ptr[int32](16777216),
		// 			MinBackupRetentionDays: to.Ptr[int32](7),
		// 			MinLargeStorageMB: to.Ptr[int32](0),
		// 			MinStorageMB: to.Ptr[int32](5120),
		// 			ServiceLevelObjectives: []*armmariadb.PerformanceTierServiceLevelObjectives{
		// 				{
		// 					Edition: to.Ptr("MemoryOptimized"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("MO_Gen5_2"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](2),
		// 				},
		// 				{
		// 					Edition: to.Ptr("MemoryOptimized"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("MO_Gen5_4"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](4),
		// 				},
		// 				{
		// 					Edition: to.Ptr("MemoryOptimized"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("MO_Gen5_8"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](8),
		// 				},
		// 				{
		// 					Edition: to.Ptr("MemoryOptimized"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("MO_Gen5_16"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](16),
		// 				},
		// 				{
		// 					Edition: to.Ptr("MemoryOptimized"),
		// 					HardwareGeneration: to.Ptr("Gen5"),
		// 					ID: to.Ptr("MO_Gen5_32"),
		// 					MaxBackupRetentionDays: to.Ptr[int32](35),
		// 					MaxStorageMB: to.Ptr[int32](2097152),
		// 					MinBackupRetentionDays: to.Ptr[int32](7),
		// 					MinStorageMB: to.Ptr[int32](5120),
		// 					VCore: to.Ptr[int32](32),
		// 			}},
		// 	}},
		// }
	}
}

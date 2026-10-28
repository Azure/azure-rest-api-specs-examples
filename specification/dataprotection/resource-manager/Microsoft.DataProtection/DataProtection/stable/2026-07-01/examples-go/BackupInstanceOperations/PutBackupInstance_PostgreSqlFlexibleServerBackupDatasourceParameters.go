package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v4"
)

// Generated from example definition: 2026-07-01/BackupInstanceOperations/PutBackupInstance_PostgreSqlFlexibleServerBackupDatasourceParameters.json
func ExampleBackupInstancesClient_BeginCreateOrUpdate_createBackupInstanceWithPostgreSqlFlexibleServerBackupDatasourceParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("62b829ee-7936-40c9-a1c9-47a93f9f3965", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupInstancesClient().BeginCreateOrUpdate(ctx, "pgflexrg", "pgflexvault", "pgflexbi", armdataprotection.BackupInstanceResource{
		Properties: &armdataprotection.BackupInstance{
			DataSourceInfo: &armdataprotection.Datasource{
				DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
				ObjectType:       to.Ptr("Datasource"),
				ResourceID:       to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver"),
				ResourceLocation: to.Ptr("eastus2euap"),
				ResourceName:     to.Ptr("pgflexserver"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
				ResourceURI:      to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver"),
			},
			DataSourceSetInfo: &armdataprotection.DatasourceSet{
				DatasourceType:   to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
				ObjectType:       to.Ptr("DatasourceSet"),
				ResourceID:       to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver"),
				ResourceLocation: to.Ptr("eastus2euap"),
				ResourceName:     to.Ptr("pgflexserver"),
				ResourceType:     to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
				ResourceURI:      to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver"),
			},
			FriendlyName: to.Ptr("pgflexbi"),
			ObjectType:   to.Ptr("BackupInstance"),
			PolicyInfo: &armdataprotection.PolicyInfo{
				PolicyID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DataProtection/BackupVaults/pgflexvault/backupPolicies/pgflexpolicy"),
				PolicyParameters: &armdataprotection.PolicyParameters{
					BackupDatasourceParametersList: []armdataprotection.BackupDatasourceParametersClassification{
						&armdataprotection.PostgreSQLFlexibleServerBackupDatasourceParameters{
							ObjectType:         to.Ptr("PostgreSqlFlexibleServerBackupDatasourceParameters"),
							BackupSolutionType: to.Ptr(armdataprotection.BackupSolutionTypePhysicalBackup),
						},
					},
				},
			},
		},
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
	// res = armdataprotection.BackupInstancesClientCreateOrUpdateResponse{
	// 	BackupInstanceResource: armdataprotection.BackupInstanceResource{
	// 		Name: to.Ptr("pgflexbi"),
	// 		Type: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances"),
	// 		ID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DataProtection/backupVaults/pgflexvault/backupInstances/pgflexbi"),
	// 		Properties: &armdataprotection.BackupInstance{
	// 			DataSourceInfo: &armdataprotection.Datasource{
	// 				DatasourceType: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 				ObjectType: to.Ptr("Datasource"),
	// 				ResourceID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver"),
	// 				ResourceLocation: to.Ptr("eastus2euap"),
	// 				ResourceName: to.Ptr("pgflexserver"),
	// 				ResourceType: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 				ResourceURI: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver"),
	// 			},
	// 			DataSourceSetInfo: &armdataprotection.DatasourceSet{
	// 				DatasourceType: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 				ObjectType: to.Ptr("DatasourceSet"),
	// 				ResourceID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver"),
	// 				ResourceLocation: to.Ptr("eastus2euap"),
	// 				ResourceName: to.Ptr("pgflexserver"),
	// 				ResourceType: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 			},
	// 			FriendlyName: to.Ptr("pgflexbi"),
	// 			ObjectType: to.Ptr("BackupInstance"),
	// 			PolicyInfo: &armdataprotection.PolicyInfo{
	// 				PolicyID: to.Ptr("/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DataProtection/backupVaults/pgflexvault/backupPolicies/pgflexpolicy"),
	// 				PolicyParameters: &armdataprotection.PolicyParameters{
	// 					BackupDatasourceParametersList: []armdataprotection.BackupDatasourceParametersClassification{
	// 						&armdataprotection.PostgreSQLFlexibleServerBackupDatasourceParameters{
	// 							ObjectType: to.Ptr("PostgreSqlFlexibleServerBackupDatasourceParameters"),
	// 							BackupSolutionType: to.Ptr(armdataprotection.BackupSolutionTypePhysicalBackup),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProtectionStatus: &armdataprotection.ProtectionStatusDetails{
	// 				Status: to.Ptr(armdataprotection.Status("NotProtected")),
	// 			},
	// 			ProvisioningState: to.Ptr("Provisioned"),
	// 		},
	// 	},
	// }
}

package armcosmosforpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/ClusterList.json
func ExampleClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmosforpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListPager(nil)
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
		// page.ClusterListResult = armcosmosforpostgresql.ClusterListResult{
		// 	Value: []*armcosmosforpostgresql.Cluster{
		// 		{
		// 			Name: to.Ptr("testcluster1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster1"),
		// 			SystemData: &armcosmosforpostgresql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"additionalProp1": to.Ptr("string"),
		// 			},
		// 			Properties: &armcosmosforpostgresql.ClusterProperties{
		// 				AdministratorLogin: to.Ptr("citus"),
		// 				CitusVersion: to.Ptr("11.1"),
		// 				CoordinatorEnablePublicIPAccess: to.Ptr(true),
		// 				CoordinatorServerEdition: to.Ptr("GeneralPurpose"),
		// 				CoordinatorStorageQuotaInMb: to.Ptr[int32](2097152),
		// 				CoordinatorVCores: to.Ptr[int32](4),
		// 				EarliestRestoreTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-14T00:00:37.467Z"); return t}()),
		// 				EnableHa: to.Ptr(false),
		// 				EnableShardsOnCoordinator: to.Ptr(false),
		// 				MaintenanceWindow: &armcosmosforpostgresql.MaintenanceWindow{
		// 					CustomWindow: to.Ptr("Disabled"),
		// 					DayOfWeek: to.Ptr[int32](0),
		// 					StartHour: to.Ptr[int32](0),
		// 					StartMinute: to.Ptr[int32](0),
		// 				},
		// 				NodeCount: to.Ptr[int32](2),
		// 				NodeEnablePublicIPAccess: to.Ptr(false),
		// 				NodeServerEdition: to.Ptr("MemoryOptimized"),
		// 				NodeStorageQuotaInMb: to.Ptr[int32](2097152),
		// 				NodeVCores: to.Ptr[int32](8),
		// 				PostgresqlVersion: to.Ptr("14"),
		// 				PreferredPrimaryZone: to.Ptr("1"),
		// 				PrivateEndpointConnections: []*armcosmosforpostgresql.SimplePrivateEndpointConnection{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ReadReplicas: []*string{
		// 					to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSQL/serverGroupsv2/testreadreplica-01"),
		// 					to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSQL/serverGroupsv2/testreadreplica-02")},
		// 					ServerNames: []*armcosmosforpostgresql.ServerNameItem{
		// 						{
		// 							Name: to.Ptr("testcluster1-c"),
		// 							FullyQualifiedDomainName: to.Ptr("testcluster1-c.postgres.database.azure.com"),
		// 						},
		// 						{
		// 							Name: to.Ptr("testcluster1-w0"),
		// 							FullyQualifiedDomainName: to.Ptr("testcluster1-w0.postgres.database.azure.com"),
		// 						},
		// 						{
		// 							Name: to.Ptr("testcluster1-w1"),
		// 							FullyQualifiedDomainName: to.Ptr("testcluster1-w1.postgres.database.azure.com"),
		// 					}},
		// 					State: to.Ptr("Ready"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testcluster2"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup2/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster2"),
		// 				SystemData: &armcosmosforpostgresql.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"additionalProp1": to.Ptr("string"),
		// 				},
		// 				Properties: &armcosmosforpostgresql.ClusterProperties{
		// 					AdministratorLogin: to.Ptr("citus"),
		// 					CitusVersion: to.Ptr("11.1"),
		// 					CoordinatorEnablePublicIPAccess: to.Ptr(true),
		// 					CoordinatorServerEdition: to.Ptr("GeneralPurpose"),
		// 					CoordinatorStorageQuotaInMb: to.Ptr[int32](524288),
		// 					CoordinatorVCores: to.Ptr[int32](4),
		// 					EarliestRestoreTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-15T00:01:32.317Z"); return t}()),
		// 					EnableHa: to.Ptr(true),
		// 					EnableShardsOnCoordinator: to.Ptr(false),
		// 					MaintenanceWindow: &armcosmosforpostgresql.MaintenanceWindow{
		// 						CustomWindow: to.Ptr("Disabled"),
		// 						DayOfWeek: to.Ptr[int32](0),
		// 						StartHour: to.Ptr[int32](0),
		// 						StartMinute: to.Ptr[int32](0),
		// 					},
		// 					NodeCount: to.Ptr[int32](1),
		// 					NodeEnablePublicIPAccess: to.Ptr(false),
		// 					NodeServerEdition: to.Ptr("MemoryOptimized"),
		// 					NodeStorageQuotaInMb: to.Ptr[int32](524288),
		// 					NodeVCores: to.Ptr[int32](8),
		// 					PostgresqlVersion: to.Ptr("15"),
		// 					PreferredPrimaryZone: to.Ptr("1"),
		// 					PrivateEndpointConnections: []*armcosmosforpostgresql.SimplePrivateEndpointConnection{
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ReadReplicas: []*string{
		// 					},
		// 					ServerNames: []*armcosmosforpostgresql.ServerNameItem{
		// 						{
		// 							Name: to.Ptr("testcluster2-c"),
		// 							FullyQualifiedDomainName: to.Ptr("testcluster2-c.postgres.database.azure.com"),
		// 						},
		// 						{
		// 							Name: to.Ptr("testcluster2-w0"),
		// 							FullyQualifiedDomainName: to.Ptr("testcluster2-w0.postgres.database.azure.com"),
		// 					}},
		// 					State: to.Ptr("Ready"),
		// 				},
		// 		}},
		// 	}
	}
}

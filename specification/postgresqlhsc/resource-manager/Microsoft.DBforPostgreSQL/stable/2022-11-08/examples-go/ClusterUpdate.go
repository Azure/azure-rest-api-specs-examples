package armcosmosforpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterUpdate.json
func ExampleClustersClient_BeginUpdate_updateMultipleConfigurationSettingsOfTheCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmosforpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdate(ctx, "TestGroup", "testcluster", armcosmosforpostgresql.ClusterForUpdate{
		Properties: &armcosmosforpostgresql.ClusterPropertiesForUpdate{
			AdministratorLoginPassword: to.Ptr("newpassword"),
			CoordinatorVCores:          to.Ptr[int32](16),
			NodeCount:                  to.Ptr[int32](4),
			NodeVCores:                 to.Ptr[int32](16),
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
	// res.Cluster = armcosmosforpostgresql.Cluster{
	// 	Name: to.Ptr("testcluster1"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster1"),
	// 	SystemData: &armcosmosforpostgresql.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"additionalProp1": to.Ptr("string"),
	// 		"additionalProp2": to.Ptr("string"),
	// 		"additionalProp3": to.Ptr("string"),
	// 	},
	// 	Properties: &armcosmosforpostgresql.ClusterProperties{
	// 		AdministratorLogin: to.Ptr("citus"),
	// 		CitusVersion: to.Ptr("11.1"),
	// 		CoordinatorEnablePublicIPAccess: to.Ptr(true),
	// 		CoordinatorServerEdition: to.Ptr("GeneralPurpose"),
	// 		CoordinatorStorageQuotaInMb: to.Ptr[int32](2097152),
	// 		CoordinatorVCores: to.Ptr[int32](16),
	// 		EarliestRestoreTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-14T00:00:37.467Z"); return t}()),
	// 		EnableHa: to.Ptr(false),
	// 		EnableShardsOnCoordinator: to.Ptr(false),
	// 		MaintenanceWindow: &armcosmosforpostgresql.MaintenanceWindow{
	// 			CustomWindow: to.Ptr("Disabled"),
	// 			DayOfWeek: to.Ptr[int32](0),
	// 			StartHour: to.Ptr[int32](0),
	// 			StartMinute: to.Ptr[int32](0),
	// 		},
	// 		NodeCount: to.Ptr[int32](4),
	// 		NodeEnablePublicIPAccess: to.Ptr(false),
	// 		NodeServerEdition: to.Ptr("MemoryOptimized"),
	// 		NodeStorageQuotaInMb: to.Ptr[int32](2097152),
	// 		NodeVCores: to.Ptr[int32](16),
	// 		PostgresqlVersion: to.Ptr("14"),
	// 		PreferredPrimaryZone: to.Ptr("1"),
	// 		PrivateEndpointConnections: []*armcosmosforpostgresql.SimplePrivateEndpointConnection{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ReadReplicas: []*string{
	// 			to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSQL/serverGroupsv2/testreadreplica-01"),
	// 			to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSQL/serverGroupsv2/testreadreplica-02")},
	// 			ServerNames: []*armcosmosforpostgresql.ServerNameItem{
	// 				{
	// 					Name: to.Ptr("testcluster1-c"),
	// 					FullyQualifiedDomainName: to.Ptr("testcluster1-c.postgres.database.azure.com"),
	// 				},
	// 				{
	// 					Name: to.Ptr("testcluster1-w0"),
	// 					FullyQualifiedDomainName: to.Ptr("testcluster1-w0.postgres.database.azure.com"),
	// 				},
	// 				{
	// 					Name: to.Ptr("testcluster1-w1"),
	// 					FullyQualifiedDomainName: to.Ptr("testcluster1-w1.postgres.database.azure.com"),
	// 				},
	// 				{
	// 					Name: to.Ptr("testcluster1-w2"),
	// 					FullyQualifiedDomainName: to.Ptr("testcluster1-w2.postgres.database.azure.com"),
	// 				},
	// 				{
	// 					Name: to.Ptr("testcluster1-w3"),
	// 					FullyQualifiedDomainName: to.Ptr("testcluster1-w2.postgres.database.azure.com"),
	// 			}},
	// 			State: to.Ptr("Ready"),
	// 		},
	// 	}
}

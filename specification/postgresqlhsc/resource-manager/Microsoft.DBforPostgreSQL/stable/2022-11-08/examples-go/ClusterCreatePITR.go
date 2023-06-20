package armcosmosforpostgresql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterCreatePITR.json
func ExampleClustersClient_BeginCreate_createANewClusterAsAPointInTimeRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmosforpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreate(ctx, "TestGroup", "testcluster", armcosmosforpostgresql.Cluster{
		Location: to.Ptr("westus"),
		Properties: &armcosmosforpostgresql.ClusterProperties{
			PointInTimeUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-14T00:00:37.467Z"); return t }()),
			SourceLocation:   to.Ptr("westus"),
			SourceResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/source-cluster"),
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
	// 	Name: to.Ptr("testcluster"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster"),
	// 	SystemData: &armcosmosforpostgresql.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcosmosforpostgresql.ClusterProperties{
	// 		AdministratorLogin: to.Ptr("citus"),
	// 		CitusVersion: to.Ptr("11.1"),
	// 		CoordinatorEnablePublicIPAccess: to.Ptr(true),
	// 		CoordinatorServerEdition: to.Ptr("GeneralPurpose"),
	// 		CoordinatorStorageQuotaInMb: to.Ptr[int32](524288),
	// 		CoordinatorVCores: to.Ptr[int32](4),
	// 		EnableHa: to.Ptr(true),
	// 		EnableShardsOnCoordinator: to.Ptr(false),
	// 		MaintenanceWindow: &armcosmosforpostgresql.MaintenanceWindow{
	// 			CustomWindow: to.Ptr("Disabled"),
	// 			DayOfWeek: to.Ptr[int32](0),
	// 			StartHour: to.Ptr[int32](0),
	// 			StartMinute: to.Ptr[int32](0),
	// 		},
	// 		NodeCount: to.Ptr[int32](1),
	// 		NodeEnablePublicIPAccess: to.Ptr(false),
	// 		NodeServerEdition: to.Ptr("MemoryOptimized"),
	// 		NodeStorageQuotaInMb: to.Ptr[int32](524288),
	// 		NodeVCores: to.Ptr[int32](8),
	// 		PostgresqlVersion: to.Ptr("15"),
	// 		PreferredPrimaryZone: to.Ptr("1"),
	// 		PrivateEndpointConnections: []*armcosmosforpostgresql.SimplePrivateEndpointConnection{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ReadReplicas: []*string{
	// 		},
	// 		ServerNames: []*armcosmosforpostgresql.ServerNameItem{
	// 			{
	// 				Name: to.Ptr("testcluster-c"),
	// 				FullyQualifiedDomainName: to.Ptr("c.testcluster.postgres.database.azure.com"),
	// 			},
	// 			{
	// 				Name: to.Ptr("testcluster-w0"),
	// 				FullyQualifiedDomainName: to.Ptr("w0.testcluster.postgres.database.azure.com"),
	// 		}},
	// 		SourceResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/source-cluster"),
	// 		State: to.Ptr("Provisioning"),
	// 	},
	// }
}

package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: 2023-03-02-preview/ServerGet.json
func ExampleServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().Get(ctx, "TestGroup", "testcluster1", "testcluster1-c", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpostgresqlhsc.ServersClientGetResponse{
	// 	ClusterServer: &armpostgresqlhsc.ClusterServer{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster1/servers/testcluster1-c"),
	// 		Name: to.Ptr("testcluster1-c"),
	// 		Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers"),
	// 		SystemData: &armpostgresqlhsc.SystemData{
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		},
	// 		Properties: &armpostgresqlhsc.ClusterServerProperties{
	// 			FullyQualifiedDomainName: to.Ptr("testcluster1-c.postgres.database.azure.com"),
	// 			PostgresqlVersion: to.Ptr("12"),
	// 			CitusVersion: to.Ptr("9.5"),
	// 			AvailabilityZone: to.Ptr("1"),
	// 			Role: to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
	// 			ServerEdition: to.Ptr("MemoryOptimized"),
	// 			StorageQuotaInMb: to.Ptr[int32](10000),
	// 			VCores: to.Ptr[int32](4),
	// 			EnableHa: to.Ptr(true),
	// 			EnablePublicIPAccess: to.Ptr(true),
	// 			State: to.Ptr("Ready"),
	// 			HaState: to.Ptr("Healthy"),
	// 			AdministratorLogin: to.Ptr("citus"),
	// 			IsReadOnly: to.Ptr(false),
	// 		},
	// 	},
	// }
}

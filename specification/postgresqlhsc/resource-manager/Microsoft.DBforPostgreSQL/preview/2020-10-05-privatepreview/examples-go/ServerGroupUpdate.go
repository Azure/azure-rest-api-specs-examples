package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupUpdate.json
func ExampleServerGroupsClient_BeginUpdate_updateTheServerGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerGroupsClient().BeginUpdate(ctx, "TestGroup", "hsctestsg", armpostgresqlhsc.ServerGroupForUpdate{
		Properties: &armpostgresqlhsc.ServerGroupPropertiesForUpdate{
			AdministratorLoginPassword: to.Ptr("secret"),
			BackupRetentionDays:        to.Ptr[int32](30),
			PostgresqlVersion:          to.Ptr(armpostgresqlhsc.PostgreSQLVersionTwelve),
			ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
				{
					EnableHa:         to.Ptr(false),
					ServerEdition:    to.Ptr(armpostgresqlhsc.ServerEditionGeneralPurpose),
					StorageQuotaInMb: to.Ptr[int64](1048576),
					VCores:           to.Ptr[int64](8),
					Name:             to.Ptr(""),
					Role:             to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
					ServerCount:      to.Ptr[int32](1),
				},
				{
					EnableHa:         to.Ptr(true),
					ServerEdition:    to.Ptr(armpostgresqlhsc.ServerEditionMemoryOptimized),
					StorageQuotaInMb: to.Ptr[int64](524288),
					VCores:           to.Ptr[int64](4),
					Name:             to.Ptr(""),
					Role:             to.Ptr(armpostgresqlhsc.ServerRoleWorker),
					ServerCount:      to.Ptr[int32](4),
				}},
		},
		Tags: map[string]*string{
			"ElasticServer": to.Ptr("2"),
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
	// res.ServerGroup = armpostgresqlhsc.ServerGroup{
	// 	Name: to.Ptr("hsctestsg"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg"),
	// 	Location: to.Ptr("westus2"),
	// 	Properties: &armpostgresqlhsc.ServerGroupProperties{
	// 	},
	// 	SystemData: &armpostgresqlhsc.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
	// 	},
	// }
}

package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupCreate.json
func ExampleServerGroupsClient_BeginCreateOrUpdate_createANewServerGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerGroupsClient().BeginCreateOrUpdate(ctx, "TestGroup", "hsctestsg", armpostgresqlhsc.ServerGroup{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"ElasticServer": to.Ptr("1"),
		},
		Properties: &armpostgresqlhsc.ServerGroupProperties{
			AdministratorLogin:         to.Ptr("citus"),
			AdministratorLoginPassword: to.Ptr("password"),
			AvailabilityZone:           to.Ptr("1"),
			BackupRetentionDays:        to.Ptr[int32](35),
			CitusVersion:               to.Ptr(armpostgresqlhsc.CitusVersionNine5),
			DelegatedSubnetArguments: &armpostgresqlhsc.ServerGroupPropertiesDelegatedSubnetArguments{
				SubnetArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-vnet-subnet"),
			},
			EnableMx:          to.Ptr(true),
			EnableZfs:         to.Ptr(false),
			PostgresqlVersion: to.Ptr(armpostgresqlhsc.PostgreSQLVersionTwelve),
			PrivateDNSZoneArguments: &armpostgresqlhsc.ServerGroupPropertiesPrivateDNSZoneArguments{
				PrivateDNSZoneArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/privateDnsZones/test-private-dns-zone"),
			},
			ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
				{
					EnableHa:         to.Ptr(true),
					ServerEdition:    to.Ptr(armpostgresqlhsc.ServerEditionGeneralPurpose),
					StorageQuotaInMb: to.Ptr[int64](524288),
					VCores:           to.Ptr[int64](4),
					Name:             to.Ptr(""),
					Role:             to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
					ServerCount:      to.Ptr[int32](1),
				},
				{
					EnableHa:         to.Ptr(false),
					ServerEdition:    to.Ptr(armpostgresqlhsc.ServerEditionMemoryOptimized),
					StorageQuotaInMb: to.Ptr[int64](524288),
					VCores:           to.Ptr[int64](4),
					Name:             to.Ptr(""),
					Role:             to.Ptr(armpostgresqlhsc.ServerRoleWorker),
					ServerCount:      to.Ptr[int32](3),
				}},
			StandbyAvailabilityZone: to.Ptr("2"),
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

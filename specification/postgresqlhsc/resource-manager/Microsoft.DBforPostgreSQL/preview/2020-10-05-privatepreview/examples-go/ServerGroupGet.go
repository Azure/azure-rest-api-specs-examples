package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupGet.json
func ExampleServerGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerGroupsClient().Get(ctx, "TestGroup", "hsctestsg1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerGroup = armpostgresqlhsc.ServerGroup{
	// 	Name: to.Ptr("hsctestsg1"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"additionalProp1": to.Ptr("string"),
	// 		"additionalProp2": to.Ptr("string"),
	// 		"additionalProp3": to.Ptr("string"),
	// 	},
	// 	Properties: &armpostgresqlhsc.ServerGroupProperties{
	// 		AdministratorLogin: to.Ptr("citus"),
	// 		AvailabilityZone: to.Ptr("1"),
	// 		CitusVersion: to.Ptr(armpostgresqlhsc.CitusVersionNine5),
	// 		DelegatedSubnetArguments: &armpostgresqlhsc.ServerGroupPropertiesDelegatedSubnetArguments{
	// 			SubnetArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-vnet-subnet"),
	// 		},
	// 		EarliestRestoreTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-14T00:00:37.467Z"); return t}()),
	// 		EnableMx: to.Ptr(true),
	// 		EnableZfs: to.Ptr(false),
	// 		MaintenanceWindow: &armpostgresqlhsc.MaintenanceWindow{
	// 			CustomWindow: to.Ptr("Disabled"),
	// 			DayOfWeek: to.Ptr[int32](0),
	// 			StartHour: to.Ptr[int32](0),
	// 			StartMinute: to.Ptr[int32](0),
	// 		},
	// 		PostgresqlVersion: to.Ptr(armpostgresqlhsc.PostgreSQLVersionTwelve),
	// 		PrivateDNSZoneArguments: &armpostgresqlhsc.ServerGroupPropertiesPrivateDNSZoneArguments{
	// 			PrivateDNSZoneArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/privateDnsZones/test-private-dns-zone"),
	// 		},
	// 		ReadReplicas: []*string{
	// 			to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSQL/serverGroupsv2/hsctestreadreplica-01"),
	// 			to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSQL/serverGroupsv2/hsctestreadreplica-02")},
	// 			ResourceProviderType: to.Ptr(armpostgresqlhsc.ResourceProviderTypeMeru),
	// 			ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
	// 				{
	// 					EnableHa: to.Ptr(true),
	// 					EnablePublicIP: to.Ptr(true),
	// 					ServerEdition: to.Ptr(armpostgresqlhsc.ServerEditionMemoryOptimized),
	// 					StorageQuotaInMb: to.Ptr[int64](10000),
	// 					VCores: to.Ptr[int64](4),
	// 					Name: to.Ptr(""),
	// 					Role: to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
	// 					ServerCount: to.Ptr[int32](1),
	// 					ServerNames: []*armpostgresqlhsc.ServerNameItem{
	// 						{
	// 							Name: to.Ptr("hsctestsg1-c"),
	// 							FullyQualifiedDomainName: to.Ptr("hsctestsg1-c.postgres.database.azure.com"),
	// 					}},
	// 				},
	// 				{
	// 					EnableHa: to.Ptr(false),
	// 					EnablePublicIP: to.Ptr(false),
	// 					ServerEdition: to.Ptr(armpostgresqlhsc.ServerEditionGeneralPurpose),
	// 					StorageQuotaInMb: to.Ptr[int64](10000),
	// 					VCores: to.Ptr[int64](8),
	// 					Name: to.Ptr(""),
	// 					Role: to.Ptr(armpostgresqlhsc.ServerRoleWorker),
	// 					ServerCount: to.Ptr[int32](3),
	// 					ServerNames: []*armpostgresqlhsc.ServerNameItem{
	// 						{
	// 							Name: to.Ptr("hsctestsg1-w0"),
	// 							FullyQualifiedDomainName: to.Ptr("hsctestsg1-w0.postgres.database.azure.com"),
	// 						},
	// 						{
	// 							Name: to.Ptr("hsctestsg1-w1"),
	// 							FullyQualifiedDomainName: to.Ptr("hsctestsg1-w1.postgres.database.azure.com"),
	// 						},
	// 						{
	// 							Name: to.Ptr("hsctestsg1-w2"),
	// 							FullyQualifiedDomainName: to.Ptr("hsctestsg1-w2.postgres.database.azure.com"),
	// 					}},
	// 			}},
	// 			StandbyAvailabilityZone: to.Ptr("2"),
	// 			State: to.Ptr(armpostgresqlhsc.ServerStateReady),
	// 		},
	// 		SystemData: &armpostgresqlhsc.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
	// 		},
	// 	}
}

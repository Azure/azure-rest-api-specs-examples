package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGroupList.json
func ExampleServerGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerGroupsClient().NewListPager(nil)
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
		// page.ServerGroupListResult = armpostgresqlhsc.ServerGroupListResult{
		// 	Value: []*armpostgresqlhsc.ServerGroup{
		// 		{
		// 			Name: to.Ptr("hsctestsg1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg1"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"additionalProp1": to.Ptr("string"),
		// 			},
		// 			Properties: &armpostgresqlhsc.ServerGroupProperties{
		// 				AdministratorLogin: to.Ptr("citus"),
		// 				AvailabilityZone: to.Ptr("1"),
		// 				CitusVersion: to.Ptr(armpostgresqlhsc.CitusVersionNine5),
		// 				DelegatedSubnetArguments: &armpostgresqlhsc.ServerGroupPropertiesDelegatedSubnetArguments{
		// 					SubnetArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-vnet-subnet"),
		// 				},
		// 				EarliestRestoreTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-14T00:00:37.467Z"); return t}()),
		// 				EnableMx: to.Ptr(true),
		// 				EnableZfs: to.Ptr(false),
		// 				PostgresqlVersion: to.Ptr(armpostgresqlhsc.PostgreSQLVersionTwelve),
		// 				PrivateDNSZoneArguments: &armpostgresqlhsc.ServerGroupPropertiesPrivateDNSZoneArguments{
		// 					PrivateDNSZoneArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/privateDnsZones/test-private-dns-zone"),
		// 				},
		// 				ReadReplicas: []*string{
		// 					to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg2")},
		// 					ResourceProviderType: to.Ptr(armpostgresqlhsc.ResourceProviderTypeMeru),
		// 					ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
		// 						{
		// 							EnableHa: to.Ptr(true),
		// 							EnablePublicIP: to.Ptr(true),
		// 							ServerEdition: to.Ptr(armpostgresqlhsc.ServerEditionMemoryOptimized),
		// 							StorageQuotaInMb: to.Ptr[int64](10000),
		// 							VCores: to.Ptr[int64](4),
		// 							Name: to.Ptr(""),
		// 							Role: to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
		// 							ServerCount: to.Ptr[int32](1),
		// 							ServerNames: []*armpostgresqlhsc.ServerNameItem{
		// 								{
		// 									Name: to.Ptr("hsctestsg1-c"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg1-c.postgres.database.azure.com"),
		// 							}},
		// 						},
		// 						{
		// 							EnableHa: to.Ptr(false),
		// 							EnablePublicIP: to.Ptr(false),
		// 							ServerEdition: to.Ptr(armpostgresqlhsc.ServerEditionGeneralPurpose),
		// 							StorageQuotaInMb: to.Ptr[int64](10000),
		// 							VCores: to.Ptr[int64](8),
		// 							Name: to.Ptr(""),
		// 							Role: to.Ptr(armpostgresqlhsc.ServerRoleWorker),
		// 							ServerCount: to.Ptr[int32](3),
		// 							ServerNames: []*armpostgresqlhsc.ServerNameItem{
		// 								{
		// 									Name: to.Ptr("hsctestsg1-w0"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg1-w0.postgres.database.azure.com"),
		// 								},
		// 								{
		// 									Name: to.Ptr("hsctestsg1-w1"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg1-w1.postgres.database.azure.com"),
		// 								},
		// 								{
		// 									Name: to.Ptr("hsctestsg1-w2"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg1-w2.postgres.database.azure.com"),
		// 							}},
		// 					}},
		// 					StandbyAvailabilityZone: to.Ptr("2"),
		// 					State: to.Ptr(armpostgresqlhsc.ServerStateReady),
		// 				},
		// 				SystemData: &armpostgresqlhsc.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("hsctestsg2"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg2"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"additionalProp2": to.Ptr("string"),
		// 				},
		// 				Properties: &armpostgresqlhsc.ServerGroupProperties{
		// 					AdministratorLogin: to.Ptr("citus"),
		// 					AvailabilityZone: to.Ptr("1"),
		// 					CitusVersion: to.Ptr(armpostgresqlhsc.CitusVersionNine5),
		// 					EarliestRestoreTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-14T00:00:37.467Z"); return t}()),
		// 					EnableMx: to.Ptr(true),
		// 					EnableZfs: to.Ptr(false),
		// 					PostgresqlVersion: to.Ptr(armpostgresqlhsc.PostgreSQLVersionTwelve),
		// 					ResourceProviderType: to.Ptr(armpostgresqlhsc.ResourceProviderTypeMeru),
		// 					ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
		// 						{
		// 							EnableHa: to.Ptr(true),
		// 							EnablePublicIP: to.Ptr(true),
		// 							ServerEdition: to.Ptr(armpostgresqlhsc.ServerEditionMemoryOptimized),
		// 							StorageQuotaInMb: to.Ptr[int64](10000),
		// 							VCores: to.Ptr[int64](8),
		// 							Name: to.Ptr(""),
		// 							Role: to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
		// 							ServerCount: to.Ptr[int32](1),
		// 							ServerNames: []*armpostgresqlhsc.ServerNameItem{
		// 								{
		// 									Name: to.Ptr("hsctestsg2-c"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg2-c.postgres.database.azure.com"),
		// 							}},
		// 						},
		// 						{
		// 							EnableHa: to.Ptr(false),
		// 							EnablePublicIP: to.Ptr(false),
		// 							ServerEdition: to.Ptr(armpostgresqlhsc.ServerEditionGeneralPurpose),
		// 							StorageQuotaInMb: to.Ptr[int64](10000),
		// 							VCores: to.Ptr[int64](4),
		// 							Name: to.Ptr(""),
		// 							Role: to.Ptr(armpostgresqlhsc.ServerRoleWorker),
		// 							ServerCount: to.Ptr[int32](2),
		// 							ServerNames: []*armpostgresqlhsc.ServerNameItem{
		// 								{
		// 									Name: to.Ptr("hsctestsg2-w0"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg2-w0.postgres.database.azure.com"),
		// 								},
		// 								{
		// 									Name: to.Ptr("hsctestsg2-w1"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg2-w1.postgres.database.azure.com"),
		// 							}},
		// 					}},
		// 					SourceServerGroup: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg1"),
		// 					StandbyAvailabilityZone: to.Ptr("2"),
		// 					State: to.Ptr(armpostgresqlhsc.ServerStateReady),
		// 				},
		// 				SystemData: &armpostgresqlhsc.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("hsctestsg3"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg3"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"additionalProp3": to.Ptr("string"),
		// 				},
		// 				Properties: &armpostgresqlhsc.ServerGroupProperties{
		// 					AdministratorLogin: to.Ptr("citus"),
		// 					AvailabilityZone: to.Ptr("1"),
		// 					CitusVersion: to.Ptr(armpostgresqlhsc.CitusVersionNine5),
		// 					DelegatedSubnetArguments: &armpostgresqlhsc.ServerGroupPropertiesDelegatedSubnetArguments{
		// 						SubnetArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-vnet-subnet"),
		// 					},
		// 					EarliestRestoreTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-14T00:00:37.467Z"); return t}()),
		// 					EnableMx: to.Ptr(true),
		// 					EnableZfs: to.Ptr(false),
		// 					PostgresqlVersion: to.Ptr(armpostgresqlhsc.PostgreSQLVersionTwelve),
		// 					PrivateDNSZoneArguments: &armpostgresqlhsc.ServerGroupPropertiesPrivateDNSZoneArguments{
		// 						PrivateDNSZoneArmResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/privateDnsZones/test-private-dns-zone"),
		// 					},
		// 					ResourceProviderType: to.Ptr(armpostgresqlhsc.ResourceProviderTypeMeru),
		// 					ServerRoleGroups: []*armpostgresqlhsc.ServerRoleGroup{
		// 						{
		// 							EnableHa: to.Ptr(true),
		// 							EnablePublicIP: to.Ptr(true),
		// 							ServerEdition: to.Ptr(armpostgresqlhsc.ServerEditionMemoryOptimized),
		// 							StorageQuotaInMb: to.Ptr[int64](10000),
		// 							VCores: to.Ptr[int64](4),
		// 							Name: to.Ptr(""),
		// 							Role: to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
		// 							ServerCount: to.Ptr[int32](1),
		// 							ServerNames: []*armpostgresqlhsc.ServerNameItem{
		// 								{
		// 									Name: to.Ptr("hsctestsg3-c"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg3-c.postgres.database.azure.com"),
		// 							}},
		// 						},
		// 						{
		// 							EnableHa: to.Ptr(false),
		// 							EnablePublicIP: to.Ptr(false),
		// 							ServerEdition: to.Ptr(armpostgresqlhsc.ServerEditionGeneralPurpose),
		// 							StorageQuotaInMb: to.Ptr[int64](10000),
		// 							VCores: to.Ptr[int64](8),
		// 							Name: to.Ptr(""),
		// 							Role: to.Ptr(armpostgresqlhsc.ServerRoleWorker),
		// 							ServerCount: to.Ptr[int32](3),
		// 							ServerNames: []*armpostgresqlhsc.ServerNameItem{
		// 								{
		// 									Name: to.Ptr("hsctestsg3-w0"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg3-w0.postgres.database.azure.com"),
		// 								},
		// 								{
		// 									Name: to.Ptr("hsctestsg3-w1"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg3-w1.postgres.database.azure.com"),
		// 								},
		// 								{
		// 									Name: to.Ptr("hsctestsg3-w2"),
		// 									FullyQualifiedDomainName: to.Ptr("hsctestsg3-w2.postgres.database.azure.com"),
		// 							}},
		// 					}},
		// 					StandbyAvailabilityZone: to.Ptr("2"),
		// 					State: to.Ptr(armpostgresqlhsc.ServerStateReady),
		// 				},
		// 				SystemData: &armpostgresqlhsc.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}

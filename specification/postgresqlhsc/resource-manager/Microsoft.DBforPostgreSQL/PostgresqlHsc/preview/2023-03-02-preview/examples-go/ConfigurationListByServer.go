package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: 2023-03-02-preview/ConfigurationListByServer.json
func ExampleConfigurationsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationsClient().NewListByServerPager("TestResourceGroup", "testcluster", "testserver", nil)
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
		// page = armpostgresqlhsc.ConfigurationsClientListByServerResponse{
		// 	ServerConfigurationListResult: armpostgresqlhsc.ServerConfigurationListResult{
		// 		Value: []*armpostgresqlhsc.ServerConfiguration{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/servers/testserver/configurations/application_name"),
		// 				Name: to.Ptr("application_name"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/configurations"),
		// 				SystemData: &armpostgresqlhsc.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				},
		// 				Properties: &armpostgresqlhsc.ServerConfigurationProperties{
		// 					Value: to.Ptr(""),
		// 					Description: to.Ptr("Sets the application name to be reported in statistics and logs."),
		// 					DefaultValue: to.Ptr(""),
		// 					DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataType("String")),
		// 					AllowedValues: to.Ptr("[A-Za-z0-9._-]*"),
		// 					Source: to.Ptr("system-default"),
		// 					ProvisioningState: to.Ptr(armpostgresqlhsc.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/servers/testserver/configurations/array_nulls"),
		// 				Name: to.Ptr("array_nulls"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/configurations"),
		// 				SystemData: &armpostgresqlhsc.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				},
		// 				Properties: &armpostgresqlhsc.ServerConfigurationProperties{
		// 					Value: to.Ptr("on"),
		// 					Description: to.Ptr("Enables input of NULL elements in arrays."),
		// 					DefaultValue: to.Ptr("on"),
		// 					DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataTypeBoolean),
		// 					AllowedValues: to.Ptr("on,off"),
		// 					Source: to.Ptr("system-default"),
		// 					ProvisioningState: to.Ptr(armpostgresqlhsc.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/servers/testserver/configurations/autovacuum"),
		// 				Name: to.Ptr("autovacuum"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/configurations"),
		// 				SystemData: &armpostgresqlhsc.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				},
		// 				Properties: &armpostgresqlhsc.ServerConfigurationProperties{
		// 					Value: to.Ptr("on"),
		// 					Description: to.Ptr("Starts the autovacuum subprocess."),
		// 					DefaultValue: to.Ptr("on"),
		// 					DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataTypeBoolean),
		// 					AllowedValues: to.Ptr("on,off"),
		// 					Source: to.Ptr("system-default"),
		// 					ProvisioningState: to.Ptr(armpostgresqlhsc.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

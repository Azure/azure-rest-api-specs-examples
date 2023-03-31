package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ConfigurationListByServer.json
func ExampleConfigurationsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationsClient().NewListByServerPager("TestResourceGroup", "hsctestsg", "testserver", nil)
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
		// page.ServerConfigurationListResult = armpostgresqlhsc.ServerConfigurationListResult{
		// 	Value: []*armpostgresqlhsc.ServerConfiguration{
		// 		{
		// 			Name: to.Ptr("application_name"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg/servers/testserver/configurations/application_name"),
		// 			Properties: &armpostgresqlhsc.ServerConfigurationProperties{
		// 				Description: to.Ptr("Sets the application name to be reported in statistics and logs."),
		// 				AllowedValues: to.Ptr("[A-Za-z0-9._-]*"),
		// 				DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataType("String")),
		// 				DefaultValue: to.Ptr(""),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr(""),
		// 			},
		// 			SystemData: &armpostgresqlhsc.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("array_nulls"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg/servers/testserver/configurations/array_nulls"),
		// 			Properties: &armpostgresqlhsc.ServerConfigurationProperties{
		// 				Description: to.Ptr("Enables input of NULL elements in arrays."),
		// 				AllowedValues: to.Ptr("on,off"),
		// 				DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataTypeBoolean),
		// 				DefaultValue: to.Ptr("on"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("on"),
		// 			},
		// 			SystemData: &armpostgresqlhsc.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("autovacuum"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg/servers/testserver/configurations/autovacuum"),
		// 			Properties: &armpostgresqlhsc.ServerConfigurationProperties{
		// 				Description: to.Ptr("Starts the autovacuum subprocess."),
		// 				AllowedValues: to.Ptr("on,off"),
		// 				DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataTypeBoolean),
		// 				DefaultValue: to.Ptr("on"),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("on"),
		// 			},
		// 			SystemData: &armpostgresqlhsc.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

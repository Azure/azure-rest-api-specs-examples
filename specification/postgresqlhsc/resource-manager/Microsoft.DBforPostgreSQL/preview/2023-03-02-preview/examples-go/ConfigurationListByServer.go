package armcosmosforpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/ConfigurationListByServer.json
func ExampleConfigurationsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmosforpostgresql.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ServerConfigurationListResult = armcosmosforpostgresql.ServerConfigurationListResult{
		// 	Value: []*armcosmosforpostgresql.ServerConfiguration{
		// 		{
		// 			Name: to.Ptr("application_name"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/servers/testserver/configurations/application_name"),
		// 			SystemData: &armcosmosforpostgresql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 			},
		// 			Properties: &armcosmosforpostgresql.ServerConfigurationProperties{
		// 				Description: to.Ptr("Sets the application name to be reported in statistics and logs."),
		// 				AllowedValues: to.Ptr("[A-Za-z0-9._-]*"),
		// 				DataType: to.Ptr(armcosmosforpostgresql.ConfigurationDataType("String")),
		// 				DefaultValue: to.Ptr(""),
		// 				ProvisioningState: to.Ptr(armcosmosforpostgresql.ProvisioningStateSucceeded),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("array_nulls"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/servers/testserver/configurations/array_nulls"),
		// 			SystemData: &armcosmosforpostgresql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 			},
		// 			Properties: &armcosmosforpostgresql.ServerConfigurationProperties{
		// 				Description: to.Ptr("Enables input of NULL elements in arrays."),
		// 				AllowedValues: to.Ptr("on,off"),
		// 				DataType: to.Ptr(armcosmosforpostgresql.ConfigurationDataTypeBoolean),
		// 				DefaultValue: to.Ptr("on"),
		// 				ProvisioningState: to.Ptr(armcosmosforpostgresql.ProvisioningStateSucceeded),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("on"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("autovacuum"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/servers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/servers/testserver/configurations/autovacuum"),
		// 			SystemData: &armcosmosforpostgresql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 			},
		// 			Properties: &armcosmosforpostgresql.ServerConfigurationProperties{
		// 				Description: to.Ptr("Starts the autovacuum subprocess."),
		// 				AllowedValues: to.Ptr("on,off"),
		// 				DataType: to.Ptr(armcosmosforpostgresql.ConfigurationDataTypeBoolean),
		// 				DefaultValue: to.Ptr("on"),
		// 				ProvisioningState: to.Ptr(armcosmosforpostgresql.ProvisioningStateSucceeded),
		// 				Source: to.Ptr("system-default"),
		// 				Value: to.Ptr("on"),
		// 			},
		// 	}},
		// }
	}
}

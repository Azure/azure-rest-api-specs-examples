package armcosmosforpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/ConfigurationListByCluster.json
func ExampleConfigurationsClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmosforpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationsClient().NewListByClusterPager("TestResourceGroup", "testcluster", nil)
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
		// page.ClusterConfigurationListResult = armcosmosforpostgresql.ClusterConfigurationListResult{
		// 	Value: []*armcosmosforpostgresql.Configuration{
		// 		{
		// 			Name: to.Ptr("array_nulls"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/configurations/array_nulls"),
		// 			SystemData: &armcosmosforpostgresql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 			},
		// 			Properties: &armcosmosforpostgresql.ConfigurationProperties{
		// 				Description: to.Ptr("Enable input of NULL elements in arrays."),
		// 				AllowedValues: to.Ptr("on,off"),
		// 				DataType: to.Ptr(armcosmosforpostgresql.ConfigurationDataTypeBoolean),
		// 				ProvisioningState: to.Ptr(armcosmosforpostgresql.ProvisioningStateSucceeded),
		// 				ServerRoleGroupConfigurations: []*armcosmosforpostgresql.ServerRoleGroupConfiguration{
		// 					{
		// 						DefaultValue: to.Ptr("on"),
		// 						Role: to.Ptr(armcosmosforpostgresql.ServerRoleCoordinator),
		// 						Source: to.Ptr("system-default"),
		// 						Value: to.Ptr("on"),
		// 					},
		// 					{
		// 						DefaultValue: to.Ptr("on"),
		// 						Role: to.Ptr(armcosmosforpostgresql.ServerRoleWorker),
		// 						Source: to.Ptr("user-override"),
		// 						Value: to.Ptr("off"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("backslash_quote"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/configurations/backslash_quote"),
		// 			SystemData: &armcosmosforpostgresql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 			},
		// 			Properties: &armcosmosforpostgresql.ConfigurationProperties{
		// 				Description: to.Ptr("Sets whether \"\\'\" is allowed in string literals."),
		// 				AllowedValues: to.Ptr("safe_encoding,on,off"),
		// 				DataType: to.Ptr(armcosmosforpostgresql.ConfigurationDataTypeEnumeration),
		// 				ProvisioningState: to.Ptr(armcosmosforpostgresql.ProvisioningStateSucceeded),
		// 				ServerRoleGroupConfigurations: []*armcosmosforpostgresql.ServerRoleGroupConfiguration{
		// 					{
		// 						DefaultValue: to.Ptr("safe_encoding"),
		// 						Role: to.Ptr(armcosmosforpostgresql.ServerRoleCoordinator),
		// 						Source: to.Ptr("user-override"),
		// 						Value: to.Ptr("off"),
		// 					},
		// 					{
		// 						DefaultValue: to.Ptr("safe_encoding"),
		// 						Role: to.Ptr(armcosmosforpostgresql.ServerRoleWorker),
		// 						Source: to.Ptr("system-default"),
		// 						Value: to.Ptr("safe_encoding"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("work_mem"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/configurations/work_mem"),
		// 			SystemData: &armcosmosforpostgresql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 			},
		// 			Properties: &armcosmosforpostgresql.ConfigurationProperties{
		// 				Description: to.Ptr("Sets the amount of memory to be used by internal sort operations and hash tables before writing to temporary disk files."),
		// 				AllowedValues: to.Ptr("4096-2097151"),
		// 				DataType: to.Ptr(armcosmosforpostgresql.ConfigurationDataTypeInteger),
		// 				ProvisioningState: to.Ptr(armcosmosforpostgresql.ProvisioningStateSucceeded),
		// 				ServerRoleGroupConfigurations: []*armcosmosforpostgresql.ServerRoleGroupConfiguration{
		// 					{
		// 						DefaultValue: to.Ptr("158720"),
		// 						Role: to.Ptr(armcosmosforpostgresql.ServerRoleCoordinator),
		// 						Source: to.Ptr("system-default"),
		// 						Value: to.Ptr("158720"),
		// 					},
		// 					{
		// 						DefaultValue: to.Ptr("115712"),
		// 						Role: to.Ptr(armcosmosforpostgresql.ServerRoleWorker),
		// 						Source: to.Ptr("system-default"),
		// 						Value: to.Ptr("115712"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("client_encoding"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/configurations/client_encoding"),
		// 			SystemData: &armcosmosforpostgresql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
		// 			},
		// 			Properties: &armcosmosforpostgresql.ConfigurationProperties{
		// 				Description: to.Ptr("Sets the client's character set encoding."),
		// 				AllowedValues: to.Ptr("BIG5,EUC_CN,EUC_JP,EUC_JIS_2004,EUC_KR,EUC_TW,GB18030,GBK,ISO_8859_5,ISO_8859_6,ISO_8859_7,ISO_8859_8,JOHAB,KOI8R,KOI8U,LATIN1,LATIN2,LATIN3,LATIN4,LATIN5,LATIN6,LATIN7,LATIN8,LATIN9,LATIN10,MULE_INTERNAL,SJIS,SHIFT_JIS_2004,SQL_ASCII,UHC,UTF8,WIN866,WIN874,WIN1250,WIN1251,WIN1252,WIN1253,WIN1254,WIN1255,WIN1256,WIN1257,WIN1258"),
		// 				DataType: to.Ptr(armcosmosforpostgresql.ConfigurationDataTypeEnumeration),
		// 				ProvisioningState: to.Ptr(armcosmosforpostgresql.ProvisioningStateSucceeded),
		// 				ServerRoleGroupConfigurations: []*armcosmosforpostgresql.ServerRoleGroupConfiguration{
		// 					{
		// 						DefaultValue: to.Ptr("sql_ascii"),
		// 						Role: to.Ptr(armcosmosforpostgresql.ServerRoleCoordinator),
		// 						Source: to.Ptr("user-override"),
		// 						Value: to.Ptr("ISO_8859_7"),
		// 					},
		// 					{
		// 						DefaultValue: to.Ptr("sql_ascii"),
		// 						Role: to.Ptr(armcosmosforpostgresql.ServerRoleWorker),
		// 						Source: to.Ptr("user-override"),
		// 						Value: to.Ptr("ISO_8859_7"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

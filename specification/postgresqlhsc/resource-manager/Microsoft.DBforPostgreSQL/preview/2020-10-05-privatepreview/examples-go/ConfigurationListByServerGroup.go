package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ConfigurationListByServerGroup.json
func ExampleConfigurationsClient_NewListByServerGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationsClient().NewListByServerGroupPager("TestResourceGroup", "hsctestsg", nil)
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
		// page.ServerGroupConfigurationListResult = armpostgresqlhsc.ServerGroupConfigurationListResult{
		// 	Value: []*armpostgresqlhsc.ServerGroupConfiguration{
		// 		{
		// 			Name: to.Ptr("array_nulls"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg/configurations/array_nulls"),
		// 			Properties: &armpostgresqlhsc.ServerGroupConfigurationProperties{
		// 				Description: to.Ptr("Enable input of NULL elements in arrays."),
		// 				AllowedValues: to.Ptr("on,off"),
		// 				DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataTypeBoolean),
		// 				ServerRoleGroupConfigurations: []*armpostgresqlhsc.ServerRoleGroupConfiguration{
		// 					{
		// 						DefaultValue: to.Ptr("on"),
		// 						Role: to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
		// 						Source: to.Ptr("system-default"),
		// 						Value: to.Ptr("on"),
		// 					},
		// 					{
		// 						DefaultValue: to.Ptr("on"),
		// 						Role: to.Ptr(armpostgresqlhsc.ServerRoleWorker),
		// 						Source: to.Ptr("user-override"),
		// 						Value: to.Ptr("off"),
		// 				}},
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
		// 			Name: to.Ptr("backslash_quote"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg/configurations/backslash_quote"),
		// 			Properties: &armpostgresqlhsc.ServerGroupConfigurationProperties{
		// 				Description: to.Ptr("Sets whether \"\\'\" is allowed in string literals."),
		// 				AllowedValues: to.Ptr("safe_encoding,on,off"),
		// 				DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataTypeEnumeration),
		// 				ServerRoleGroupConfigurations: []*armpostgresqlhsc.ServerRoleGroupConfiguration{
		// 					{
		// 						DefaultValue: to.Ptr("safe_encoding"),
		// 						Role: to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
		// 						Source: to.Ptr("user-override"),
		// 						Value: to.Ptr("off"),
		// 					},
		// 					{
		// 						DefaultValue: to.Ptr("safe_encoding"),
		// 						Role: to.Ptr(armpostgresqlhsc.ServerRoleWorker),
		// 						Source: to.Ptr("system-default"),
		// 						Value: to.Ptr("safe_encoding"),
		// 				}},
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
		// 			Name: to.Ptr("work_mem"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg/configurations/work_mem"),
		// 			Properties: &armpostgresqlhsc.ServerGroupConfigurationProperties{
		// 				Description: to.Ptr("Sets the amount of memory to be used by internal sort operations and hash tables before writing to temporary disk files."),
		// 				AllowedValues: to.Ptr("4096-2097151"),
		// 				DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataTypeInteger),
		// 				ServerRoleGroupConfigurations: []*armpostgresqlhsc.ServerRoleGroupConfiguration{
		// 					{
		// 						DefaultValue: to.Ptr("158720"),
		// 						Role: to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
		// 						Source: to.Ptr("system-default"),
		// 						Value: to.Ptr("158720"),
		// 					},
		// 					{
		// 						DefaultValue: to.Ptr("115712"),
		// 						Role: to.Ptr(armpostgresqlhsc.ServerRoleWorker),
		// 						Source: to.Ptr("system-default"),
		// 						Value: to.Ptr("115712"),
		// 				}},
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
		// 			Name: to.Ptr("client_encoding"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg/configurations/client_encoding"),
		// 			Properties: &armpostgresqlhsc.ServerGroupConfigurationProperties{
		// 				Description: to.Ptr("Sets the client's character set encoding."),
		// 				AllowedValues: to.Ptr("BIG5,EUC_CN,EUC_JP,EUC_JIS_2004,EUC_KR,EUC_TW,GB18030,GBK,ISO_8859_5,ISO_8859_6,ISO_8859_7,ISO_8859_8,JOHAB,KOI8R,KOI8U,LATIN1,LATIN2,LATIN3,LATIN4,LATIN5,LATIN6,LATIN7,LATIN8,LATIN9,LATIN10,MULE_INTERNAL,SJIS,SHIFT_JIS_2004,SQL_ASCII,UHC,UTF8,WIN866,WIN874,WIN1250,WIN1251,WIN1252,WIN1253,WIN1254,WIN1255,WIN1256,WIN1257,WIN1258"),
		// 				DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataTypeEnumeration),
		// 				ServerRoleGroupConfigurations: []*armpostgresqlhsc.ServerRoleGroupConfiguration{
		// 					{
		// 						DefaultValue: to.Ptr("sql_ascii"),
		// 						Role: to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
		// 						Source: to.Ptr("user-override"),
		// 						Value: to.Ptr("ISO_8859_7"),
		// 					},
		// 					{
		// 						DefaultValue: to.Ptr("sql_ascii"),
		// 						Role: to.Ptr(armpostgresqlhsc.ServerRoleWorker),
		// 						Source: to.Ptr("user-override"),
		// 						Value: to.Ptr("ISO_8859_7"),
		// 				}},
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

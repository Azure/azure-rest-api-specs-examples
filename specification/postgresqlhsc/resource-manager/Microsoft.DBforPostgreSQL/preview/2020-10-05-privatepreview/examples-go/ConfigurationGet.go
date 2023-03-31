package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ConfigurationGet.json
func ExampleConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationsClient().Get(ctx, "TestResourceGroup", "hsctestsg", "array_nulls", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerGroupConfiguration = armpostgresqlhsc.ServerGroupConfiguration{
	// 	Name: to.Ptr("array_nulls"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/hsctestsg/configurations/array_nulls"),
	// 	Properties: &armpostgresqlhsc.ServerGroupConfigurationProperties{
	// 		Description: to.Ptr("Enable input of NULL elements in arrays."),
	// 		AllowedValues: to.Ptr("on,off"),
	// 		DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataTypeBoolean),
	// 		ServerRoleGroupConfigurations: []*armpostgresqlhsc.ServerRoleGroupConfiguration{
	// 			{
	// 				DefaultValue: to.Ptr("on"),
	// 				Role: to.Ptr(armpostgresqlhsc.ServerRoleCoordinator),
	// 				Source: to.Ptr("system-default"),
	// 				Value: to.Ptr("on"),
	// 			},
	// 			{
	// 				DefaultValue: to.Ptr("on"),
	// 				Role: to.Ptr(armpostgresqlhsc.ServerRoleWorker),
	// 				Source: to.Ptr("user-override"),
	// 				Value: to.Ptr("off"),
	// 		}},
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

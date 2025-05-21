package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ConfigurationGet.json
func ExampleConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationsClient().Get(ctx, "testrg", "testserver", "array_nulls", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Configuration = armpostgresqlflexibleservers.Configuration{
	// 	Name: to.Ptr("array_nulls"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/configurations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testserver/configurations/array_nulls"),
	// 	Properties: &armpostgresqlflexibleservers.ConfigurationProperties{
	// 		Description: to.Ptr("Enables input of NULL (case insensitive) to be considered as NULL value rather than the literal String 'NULL'."),
	// 		AllowedValues: to.Ptr("on,off"),
	// 		DataType: to.Ptr(armpostgresqlflexibleservers.ConfigurationDataTypeBoolean),
	// 		DefaultValue: to.Ptr("on"),
	// 		DocumentationLink: to.Ptr("https://www.postgresql.org/docs/13/runtime-config-compatible.html#GUC-ARRAY-NULLS"),
	// 		IsConfigPendingRestart: to.Ptr(false),
	// 		IsDynamicConfig: to.Ptr(true),
	// 		IsReadOnly: to.Ptr(false),
	// 		Source: to.Ptr("system-default"),
	// 		Value: to.Ptr("on"),
	// 	},
	// }
}

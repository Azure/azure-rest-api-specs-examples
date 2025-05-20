package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ConfigurationUpdate.json
func ExampleConfigurationsClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationsClient().BeginPut(ctx, "testrg", "testserver", "constraint_exclusion", armpostgresqlflexibleservers.Configuration{
		Properties: &armpostgresqlflexibleservers.ConfigurationProperties{
			Source: to.Ptr("user-override"),
			Value:  to.Ptr("on"),
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
	// res.Configuration = armpostgresqlflexibleservers.Configuration{
	// 	Name: to.Ptr("constraint_exclusion"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/configurations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testserver/configurations/constraint_exclusion"),
	// 	Properties: &armpostgresqlflexibleservers.ConfigurationProperties{
	// 		Description: to.Ptr("Controls the query planner's use of table constraints to optimize queries."),
	// 		AllowedValues: to.Ptr("partition,on,off"),
	// 		DataType: to.Ptr(armpostgresqlflexibleservers.ConfigurationDataTypeEnumeration),
	// 		DefaultValue: to.Ptr("partition"),
	// 		DocumentationLink: to.Ptr("https://www.postgresql.org/docs/13/runtime-config-query.html#GUC-CONSTRAINT-EXCLUSION"),
	// 		IsConfigPendingRestart: to.Ptr(false),
	// 		IsDynamicConfig: to.Ptr(true),
	// 		IsReadOnly: to.Ptr(false),
	// 		Source: to.Ptr("user-override"),
	// 		Value: to.Ptr("on"),
	// 	},
	// }
}

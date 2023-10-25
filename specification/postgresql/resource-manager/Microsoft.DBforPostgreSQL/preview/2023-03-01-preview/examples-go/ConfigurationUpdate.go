package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/ConfigurationUpdate.json
func ExampleConfigurationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationsClient().BeginUpdate(ctx, "testrg", "testserver", "event_scheduler", armpostgresqlflexibleservers.ConfigurationForUpdate{
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
	// 	Name: to.Ptr("event_scheduler"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/configurations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testserver/configurations/event_scheduler"),
	// 	Properties: &armpostgresqlflexibleservers.ConfigurationProperties{
	// 		Description: to.Ptr("Indicates the status of the Event Scheduler."),
	// 		AllowedValues: to.Ptr("ON,OFF,DISABLED"),
	// 		DataType: to.Ptr(armpostgresqlflexibleservers.ConfigurationDataTypeEnumeration),
	// 		DefaultValue: to.Ptr("OFF"),
	// 		Source: to.Ptr("user-override"),
	// 		Value: to.Ptr("ON"),
	// 	},
	// }
}

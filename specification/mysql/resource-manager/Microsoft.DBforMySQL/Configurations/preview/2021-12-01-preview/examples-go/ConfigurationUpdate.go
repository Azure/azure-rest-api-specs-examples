package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/mysql/resource-manager/Microsoft.DBforMySQL/Configurations/preview/2021-12-01-preview/examples/ConfigurationUpdate.json
func ExampleConfigurationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationsClient().BeginUpdate(ctx, "testrg", "testserver", "event_scheduler", armmysqlflexibleservers.Configuration{
		Properties: &armmysqlflexibleservers.ConfigurationProperties{
			Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceUserOverride),
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
	// res.Configuration = armmysqlflexibleservers.Configuration{
	// 	Name: to.Ptr("event_scheduler"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/configurations/event_scheduler"),
	// 	Properties: &armmysqlflexibleservers.ConfigurationProperties{
	// 		Description: to.Ptr("Indicates the status of the Event Scheduler. It is always OFF for a replica server to keep the replication consistency."),
	// 		AllowedValues: to.Ptr("ON,OFF"),
	// 		DataType: to.Ptr("Enumeration"),
	// 		DefaultValue: to.Ptr("OFF"),
	// 		IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
	// 		IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
	// 		IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
	// 		Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceUserOverride),
	// 		Value: to.Ptr("ON"),
	// 	},
	// }
}

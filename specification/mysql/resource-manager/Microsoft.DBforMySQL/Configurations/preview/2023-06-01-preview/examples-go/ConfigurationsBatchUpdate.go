package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/mysql/resource-manager/Microsoft.DBforMySQL/Configurations/preview/2023-06-01-preview/examples/ConfigurationsBatchUpdate.json
func ExampleConfigurationsClient_BeginBatchUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationsClient().BeginBatchUpdate(ctx, "testrg", "mysqltestserver", armmysqlflexibleservers.ConfigurationListForBatchUpdate{
		ResetAllToDefault: to.Ptr(armmysqlflexibleservers.ResetAllToDefaultFalse),
		Value: []*armmysqlflexibleservers.ConfigurationForBatchUpdate{
			{
				Name: to.Ptr("event_scheduler"),
				Properties: &armmysqlflexibleservers.ConfigurationForBatchUpdateProperties{
					Value: to.Ptr("OFF"),
				},
			},
			{
				Name: to.Ptr("div_precision_increment"),
				Properties: &armmysqlflexibleservers.ConfigurationForBatchUpdateProperties{
					Value: to.Ptr("8"),
				},
			}},
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
	// res.ConfigurationListResult = armmysqlflexibleservers.ConfigurationListResult{
	// 	Value: []*armmysqlflexibleservers.Configuration{
	// 		{
	// 			Name: to.Ptr("event_scheduler"),
	// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
	// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/event_scheduler"),
	// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
	// 				Description: to.Ptr("Indicates the status of the Event Scheduler. It is always OFF for a replica server to keep the replication consistency."),
	// 				AllowedValues: to.Ptr("ON,OFF"),
	// 				DataType: to.Ptr("Enumeration"),
	// 				DefaultValue: to.Ptr("OFF"),
	// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
	// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
	// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
	// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceUserOverride),
	// 				Value: to.Ptr("ON"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("div_precision_increment"),
	// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
	// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/div_precision_increment"),
	// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
	// 				Description: to.Ptr("Number of digits by which to increase the scale of the result of division operations."),
	// 				AllowedValues: to.Ptr("0-30"),
	// 				DataType: to.Ptr("Integer"),
	// 				DefaultValue: to.Ptr("4"),
	// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
	// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
	// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
	// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceUserOverride),
	// 				Value: to.Ptr("8"),
	// 			},
	// 	}},
	// }
}

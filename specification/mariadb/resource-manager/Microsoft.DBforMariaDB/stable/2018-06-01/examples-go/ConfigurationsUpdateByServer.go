package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ConfigurationsUpdateByServer.json
func ExampleServerParametersClient_BeginListUpdateConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerParametersClient().BeginListUpdateConfigurations(ctx, "testrg", "mariadbtestsvc1", armmariadb.ConfigurationListResult{}, nil)
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
	// res.ConfigurationListResult = armmariadb.ConfigurationListResult{
	// 	Value: []*armmariadb.Configuration{
	// 		{
	// 			Name: to.Ptr("event_scheduler"),
	// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers/configurations"),
	// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMariaDB/servers/mariadbtestsvc1/configurations/event_scheduler"),
	// 			Properties: &armmariadb.ConfigurationProperties{
	// 				Description: to.Ptr("Indicates the status of the Event Scheduler."),
	// 				AllowedValues: to.Ptr("ON,OFF,DISABLED"),
	// 				DataType: to.Ptr("Enumeration"),
	// 				DefaultValue: to.Ptr("OFF"),
	// 				Source: to.Ptr("system-default"),
	// 				Value: to.Ptr("OFF"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("div_precision_increment"),
	// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers/configurations"),
	// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMariaDB/servers/mariadbtestsvc1/configurations/div_precision_increment"),
	// 			Properties: &armmariadb.ConfigurationProperties{
	// 				Description: to.Ptr("Number of digits by which to increase the scale of the result of division operations."),
	// 				AllowedValues: to.Ptr("0-30"),
	// 				DataType: to.Ptr("Integer"),
	// 				DefaultValue: to.Ptr("4"),
	// 				Source: to.Ptr("system-default"),
	// 				Value: to.Ptr("4"),
	// 			},
	// 	}},
	// }
}

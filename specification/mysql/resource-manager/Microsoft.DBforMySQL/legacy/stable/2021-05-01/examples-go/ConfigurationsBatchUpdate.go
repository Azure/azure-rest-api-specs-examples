package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationsBatchUpdate.json
func ExampleConfigurationsClient_BeginBatchUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmysqlflexibleservers.NewConfigurationsClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginBatchUpdate(ctx,
		"testrg",
		"mysqltestserver",
		armmysqlflexibleservers.ConfigurationListForBatchUpdate{
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

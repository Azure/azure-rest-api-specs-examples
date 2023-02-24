package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBDatabaseAccountGetMetricDefinitions.json
func ExampleDatabaseAccountsClient_NewListMetricDefinitionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewDatabaseAccountsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListMetricDefinitionsPager("rg1", "ddb1", nil)
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
		// page.MetricDefinitionsListResult = armcosmos.MetricDefinitionsListResult{
		// 	Value: []*armcosmos.MetricDefinition{
		// 		{
		// 			Name: &armcosmos.MetricName{
		// 				LocalizedValue: to.Ptr("Total Requests"),
		// 				Value: to.Ptr("Total Requests"),
		// 			},
		// 			MetricAvailabilities: []*armcosmos.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P2D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P60D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armcosmos.PrimaryAggregationTypeTotal),
		// 			ResourceURI: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1"),
		// 			Unit: to.Ptr(armcosmos.UnitTypeCount),
		// 	}},
		// }
	}
}

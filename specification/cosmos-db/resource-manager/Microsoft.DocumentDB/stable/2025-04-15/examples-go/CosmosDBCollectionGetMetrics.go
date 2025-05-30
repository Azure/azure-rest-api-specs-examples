package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/011ecc5633300a5eefe43dde748f269d39e96458/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBCollectionGetMetrics.json
func ExampleCollectionClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCollectionClient().NewListMetricsPager("rg1", "ddb1", "databaseRid", "collectionRid", "$filter=(name.value eq 'Total Requests') and timeGrain eq duration'PT5M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T00:13:55.2780000Z", nil)
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
		// page.MetricListResult = armcosmos.MetricListResult{
		// 	Value: []*armcosmos.Metric{
		// 		{
		// 			Name: &armcosmos.MetricName{
		// 				LocalizedValue: to.Ptr("Total Requests"),
		// 				Value: to.Ptr("Total Requests"),
		// 			},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-20T00:13:55.278Z"); return t}()),
		// 			MetricValues: []*armcosmos.MetricValue{
		// 				{
		// 					Count: to.Ptr[int32](0),
		// 					Average: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:53:55.278Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 				},
		// 				{
		// 					Count: to.Ptr[int32](0),
		// 					Average: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:58:55.278Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 				},
		// 				{
		// 					Count: to.Ptr[int32](0),
		// 					Average: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-20T00:03:55.278Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 				},
		// 				{
		// 					Count: to.Ptr[int32](0),
		// 					Average: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-20T00:08:55.278Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 			}},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:53:55.278Z"); return t}()),
		// 			TimeGrain: to.Ptr("PT5M"),
		// 			Unit: to.Ptr(armcosmos.UnitTypeCount),
		// 	}},
		// }
	}
}

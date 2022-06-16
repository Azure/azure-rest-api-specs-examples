package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetTimeseries.json
func ExampleReportsClient_GetTimeseries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewReportsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetTimeseries(ctx,
		"MyResourceGroup",
		"MyProfile",
		"MyExperiment",
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-21T17:32:28Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-21T17:32:28Z"); return t }(),
		armfrontdoor.TimeseriesAggregationIntervalHourly,
		armfrontdoor.TimeseriesTypeMeasurementCounts,
		&armfrontdoor.ReportsClientGetTimeseriesOptions{Endpoint: nil,
			Country: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

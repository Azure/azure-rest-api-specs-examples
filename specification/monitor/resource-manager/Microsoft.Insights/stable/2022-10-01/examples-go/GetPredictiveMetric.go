package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/GetPredictiveMetric.json
func ExamplePredictiveMetricClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPredictiveMetricClient().Get(ctx, "myRG", "vmss1-Autoscale-775", "2021-10-14T22:00:00.000Z/2021-10-16T22:00:00.000Z", "PT1H", "Microsoft.Compute/virtualMachineScaleSets", "PercentageCPU", "Total", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PredictiveResponse = armmonitor.PredictiveResponse{
	// 	Data: []*armmonitor.PredictiveValue{
	// 		{
	// 			TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-14T22:00:00.000Z"); return t}()),
	// 			Value: to.Ptr[float64](123),
	// 		},
	// 		{
	// 			TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-14T22:02:00.000Z"); return t}()),
	// 			Value: to.Ptr[float64](120),
	// 		},
	// 		{
	// 			TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-14T22:02:00.000Z"); return t}()),
	// 			Value: to.Ptr[float64](88),
	// 		},
	// 		{
	// 			TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-14T22:03:00.000Z"); return t}()),
	// 			Value: to.Ptr[float64](100),
	// 	}},
	// 	Interval: to.Ptr("PT1H"),
	// 	MetricName: to.Ptr("PercentageCPU"),
	// 	TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1"),
	// 	Timespan: to.Ptr("2021-10-14T22:00:00.000Z/2021-10-16T22:00:00.000Z"),
	// }
}

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2019-03-01/examples/metricBaselines.json
func ExampleBaselinesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBaselinesClient().NewListPager("subscriptions/b368ca2f-e298-46b7-b0ab-012281956afa/resourceGroups/vms/providers/Microsoft.Compute/virtualMachines/vm1", &armmonitor.BaselinesClientListOptions{Metricnames: nil,
		Metricnamespace: nil,
		Timespan:        to.Ptr("2019-03-12T11:00:00.000Z/2019-03-12T12:00:00.000Z"),
		Interval:        to.Ptr("PT1H"),
		Aggregation:     to.Ptr("average"),
		Sensitivities:   to.Ptr("Low,Medium"),
		Filter:          nil,
		ResultType:      nil,
	})
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
		// page.MetricBaselinesResponse = armmonitor.MetricBaselinesResponse{
		// 	Value: []*armmonitor.SingleMetricBaseline{
		// 		{
		// 			Name: to.Ptr("Percentage CPU"),
		// 			Type: to.Ptr("microsoft.insights/metricBaselines"),
		// 			ID: to.Ptr("/subscriptions/11aeb0ed-456b-4ca0-8df5-b9fbdc63d0d3/resourceGroups/SmartAnalytics-DEV-VM/providers/Microsoft.Compute/virtualMachines/DemoVM1/providers/microsoft.insights/metricBaselines/Percentage CPU"),
		// 			Properties: &armmonitor.MetricBaselinesProperties{
		// 				Baselines: []*armmonitor.TimeSeriesBaseline{
		// 					{
		// 						Aggregation: to.Ptr("average"),
		// 						Data: []*armmonitor.SingleBaseline{
		// 							{
		// 								HighThresholds: []*float64{
		// 									to.Ptr[float64](90.3453),
		// 									to.Ptr[float64](91.3453)},
		// 									LowThresholds: []*float64{
		// 										to.Ptr[float64](30),
		// 										to.Ptr[float64](31.1)},
		// 										Sensitivity: to.Ptr(armmonitor.BaselineSensitivityLow),
		// 									},
		// 									{
		// 										HighThresholds: []*float64{
		// 											to.Ptr[float64](70.3453),
		// 											to.Ptr[float64](71.3453)},
		// 											LowThresholds: []*float64{
		// 												to.Ptr[float64](50),
		// 												to.Ptr[float64](51.1)},
		// 												Sensitivity: to.Ptr(armmonitor.BaselineSensitivityMedium),
		// 										}},
		// 										Dimensions: []*armmonitor.MetricSingleDimension{
		// 										},
		// 										MetadataValues: []*armmonitor.BaselineMetadata{
		// 											{
		// 												Name: to.Ptr("ErrorType"),
		// 												Value: to.Ptr("TooManyAnomalies"),
		// 											},
		// 											{
		// 												Name: to.Ptr("SeasonalityFrequency"),
		// 												Value: to.Ptr("288"),
		// 										}},
		// 										Timestamps: []*time.Time{
		// 											to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-12T11:00:00.000Z"); return t}()),
		// 											to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-12T12:00:00.000Z"); return t}())},
		// 									}},
		// 									Interval: to.Ptr("PT1H"),
		// 									Namespace: to.Ptr("microsoft.compute/virtualmachines"),
		// 									Timespan: to.Ptr("2019-03-12T11:00:00.000Z/2019-03-12T12:00:00.000Z"),
		// 								},
		// 						}},
		// 					}
	}
}

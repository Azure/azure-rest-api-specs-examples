package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getWebTestMetricAlert.json
func ExampleMetricAlertsClient_Get_getAWebTestAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricAlertsClient().Get(ctx, "rg-example", "webtest-name-example", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetricAlertResource = armmonitor.MetricAlertResource{
	// 	Name: to.Ptr("webtest-name-example"),
	// 	Type: to.Ptr("Microsoft.Insights/metricAlerts"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/metricalerts/webtest-name-example"),
	// 	Location: to.Ptr("global"),
	// 	Tags: map[string]*string{
	// 		"hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/component-example": to.Ptr("Resource"),
	// 		"hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/webtest-name-example": to.Ptr("Resource"),
	// 	},
	// 	Properties: &armmonitor.MetricAlertProperties{
	// 		Description: to.Ptr("Automatically created alert rule for availability test \"webtest-name-example\""),
	// 		Actions: []*armmonitor.MetricAlertAction{
	// 		},
	// 		Criteria: &armmonitor.WebtestLocationAvailabilityCriteria{
	// 			ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria),
	// 			ComponentID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/component-example"),
	// 			FailedLocationCount: to.Ptr[float32](2),
	// 			WebTestID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/webtest-name-example"),
	// 		},
	// 		Enabled: to.Ptr(true),
	// 		EvaluationFrequency: to.Ptr("PT1M"),
	// 		Scopes: []*string{
	// 			to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/webtest-name-example"),
	// 			to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/component-example")},
	// 			Severity: to.Ptr[int32](4),
	// 			WindowSize: to.Ptr("PT15M"),
	// 		},
	// 	}
}

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2026-01-01/createOrUpdateWebTestMetricAlert.json
func ExampleMetricAlertsClient_CreateOrUpdate_createOrUpdateAWebTestAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("12345678-1234-1234-1234-123456789101", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricAlertsClient().CreateOrUpdate(ctx, "rg-example", "webtest-name-example", armmonitor.MetricAlertResource{
		Location: to.Ptr("global"),
		Properties: &armmonitor.MetricAlertProperties{
			Description: to.Ptr("Automatically created alert rule for availability test \"component-example\" a"),
			Actions:     []*armmonitor.MetricAlertAction{},
			Criteria: &armmonitor.WebtestLocationAvailabilityCriteria{
				ComponentID:         to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example"),
				FailedLocationCount: to.Ptr[float32](2),
				ODataType:           to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria),
				WebTestID:           to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example"),
			},
			Enabled:             to.Ptr(true),
			EvaluationFrequency: to.Ptr("PT1M"),
			Scopes: []*string{
				to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example"),
				to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example"),
			},
			Severity:   to.Ptr[int32](4),
			WindowSize: to.Ptr("PT15M"),
		},
		Tags: map[string]*string{
			"hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example": to.Ptr("Resource"),
			"hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example":      to.Ptr("Resource"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.MetricAlertsClientCreateOrUpdateResponse{
	// 	MetricAlertResource: armmonitor.MetricAlertResource{
	// 		Name: to.Ptr("webtest-name-example"),
	// 		Type: to.Ptr("Microsoft.Insights/metricAlerts"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/metricalerts/webtest-name-example"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armmonitor.MetricAlertProperties{
	// 			Description: to.Ptr("Automatically created alert rule for availability test \"webtest-name-example\" a"),
	// 			Actions: []*armmonitor.MetricAlertAction{
	// 			},
	// 			Criteria: &armmonitor.WebtestLocationAvailabilityCriteria{
	// 				ComponentID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/component-example"),
	// 				FailedLocationCount: to.Ptr[float32](2),
	// 				ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria),
	// 				WebTestID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/webtest-name-example"),
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			EvaluationFrequency: to.Ptr("PT1M"),
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/webtest-name-example"),
	// 				to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/component-example"),
	// 			},
	// 			Severity: to.Ptr[int32](4),
	// 			WindowSize: to.Ptr("PT15M"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/component-example": to.Ptr("Resource"),
	// 			"hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/webtest-name-example": to.Ptr("Resource"),
	// 		},
	// 	},
	// }
}

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateWebTestMetricAlert.json
func ExampleMetricAlertsClient_CreateOrUpdate_createOrUpdateAWebTestAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewMetricAlertsClient("12345678-1234-1234-1234-123456789101", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "rg-example", "webtest-name-example", armmonitor.MetricAlertResource{
		Location: to.Ptr("global"),
		Tags: map[string]*string{
			"hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example": to.Ptr("Resource"),
			"hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example":      to.Ptr("Resource"),
		},
		Properties: &armmonitor.MetricAlertProperties{
			Description: to.Ptr("Automatically created alert rule for availability test \"component-example\" a"),
			Actions:     []*armmonitor.MetricAlertAction{},
			Criteria: &armmonitor.WebtestLocationAvailabilityCriteria{
				ODataType:           to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria),
				ComponentID:         to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example"),
				FailedLocationCount: to.Ptr[float32](2),
				WebTestID:           to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example"),
			},
			Enabled:             to.Ptr(true),
			EvaluationFrequency: to.Ptr("PT1M"),
			Scopes: []*string{
				to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example"),
				to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example")},
			Severity:   to.Ptr[int32](4),
			WindowSize: to.Ptr("PT15M"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

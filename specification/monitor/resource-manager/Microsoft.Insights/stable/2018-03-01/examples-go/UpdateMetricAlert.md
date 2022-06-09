```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/UpdateMetricAlert.json
func ExampleMetricAlertsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewMetricAlertsClient("14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"gigtest",
		"chiricutin",
		armmonitor.MetricAlertResourcePatch{
			Properties: &armmonitor.MetricAlertPropertiesPatch{
				Description: to.Ptr("This is the description of the rule1"),
				Actions: []*armmonitor.MetricAlertAction{
					{
						ActionGroupID: to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
						WebHookProperties: map[string]*string{
							"key11": to.Ptr("value11"),
							"key12": to.Ptr("value12"),
						},
					}},
				AutoMitigate: to.Ptr(true),
				Criteria: &armmonitor.MetricAlertSingleResourceMultipleMetricCriteria{
					ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria),
					AllOf: []*armmonitor.MetricCriteria{
						{
							Name:            to.Ptr("High_CPU_80"),
							CriterionType:   to.Ptr(armmonitor.CriterionTypeStaticThresholdCriterion),
							Dimensions:      []*armmonitor.MetricDimension{},
							MetricName:      to.Ptr("\\Processor(_Total)\\% Processor Time"),
							TimeAggregation: to.Ptr(armmonitor.AggregationTypeEnumAverage),
							Operator:        to.Ptr(armmonitor.OperatorGreaterThan),
							Threshold:       to.Ptr[float64](80.5),
						}},
				},
				Enabled:             to.Ptr(true),
				EvaluationFrequency: to.Ptr("Pt1m"),
				Scopes: []*string{
					to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme")},
				Severity:   to.Ptr[int32](3),
				WindowSize: to.Ptr("Pt15m"),
			},
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.7.0/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.6.0/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateDynamicMetricAlertMultipleResource.json
func ExampleMetricAlertsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmonitor.NewMetricAlertsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<rule-name>",
		armmonitor.MetricAlertResource{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armmonitor.MetricAlertProperties{
				Description: to.Ptr("<description>"),
				Actions: []*armmonitor.MetricAlertAction{
					{
						ActionGroupID: to.Ptr("<action-group-id>"),
						WebHookProperties: map[string]*string{
							"key11": to.Ptr("value11"),
							"key12": to.Ptr("value12"),
						},
					}},
				AutoMitigate: to.Ptr(true),
				Criteria: &armmonitor.MetricAlertMultipleResourceMultipleMetricCriteria{
					ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria),
					AllOf: []armmonitor.MultiMetricCriteriaClassification{
						&armmonitor.DynamicMetricCriteria{
							Name:             to.Ptr("<name>"),
							CriterionType:    to.Ptr(armmonitor.CriterionTypeDynamicThresholdCriterion),
							Dimensions:       []*armmonitor.MetricDimension{},
							MetricName:       to.Ptr("<metric-name>"),
							MetricNamespace:  to.Ptr("<metric-namespace>"),
							TimeAggregation:  to.Ptr(armmonitor.AggregationTypeEnumAverage),
							AlertSensitivity: to.Ptr(armmonitor.DynamicThresholdSensitivityMedium),
							FailingPeriods: &armmonitor.DynamicThresholdFailingPeriods{
								MinFailingPeriodsToAlert:  to.Ptr[float32](4),
								NumberOfEvaluationPeriods: to.Ptr[float32](4),
							},
							Operator: to.Ptr(armmonitor.DynamicThresholdOperatorGreaterOrLessThan),
						}},
				},
				Enabled:             to.Ptr(true),
				EvaluationFrequency: to.Ptr("<evaluation-frequency>"),
				Scopes: []*string{
					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1"),
					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2")},
				Severity:             to.Ptr[int32](3),
				TargetResourceRegion: to.Ptr("<target-resource-region>"),
				TargetResourceType:   to.Ptr("<target-resource-type>"),
				WindowSize:           to.Ptr("<window-size>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.3.0/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateDynamicMetricAlertMultipleResource.json
func ExampleMetricAlertsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewMetricAlertsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<rule-name>",
		armmonitor.MetricAlertResource{
			Resource: armmonitor.Resource{
				Location: to.StringPtr("<location>"),
				Tags:     map[string]*string{},
			},
			Properties: &armmonitor.MetricAlertProperties{
				Description: to.StringPtr("<description>"),
				Actions: []*armmonitor.MetricAlertAction{
					{
						ActionGroupID: to.StringPtr("<action-group-id>"),
						WebHookProperties: map[string]*string{
							"key11": to.StringPtr("value11"),
							"key12": to.StringPtr("value12"),
						},
					}},
				AutoMitigate: to.BoolPtr(true),
				Criteria: &armmonitor.MetricAlertMultipleResourceMultipleMetricCriteria{
					MetricAlertCriteria: armmonitor.MetricAlertCriteria{
						ODataType: armmonitor.OdatatypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria.ToPtr(),
					},
					AllOf: []armmonitor.MultiMetricCriteriaClassification{
						&armmonitor.DynamicMetricCriteria{
							MultiMetricCriteria: armmonitor.MultiMetricCriteria{
								Name:            to.StringPtr("<name>"),
								CriterionType:   armmonitor.CriterionTypeDynamicThresholdCriterion.ToPtr(),
								Dimensions:      []*armmonitor.MetricDimension{},
								MetricName:      to.StringPtr("<metric-name>"),
								MetricNamespace: to.StringPtr("<metric-namespace>"),
								TimeAggregation: armmonitor.AggregationTypeEnumAverage.ToPtr(),
							},
							AlertSensitivity: armmonitor.DynamicThresholdSensitivityMedium.ToPtr(),
							FailingPeriods: &armmonitor.DynamicThresholdFailingPeriods{
								MinFailingPeriodsToAlert:  to.Float32Ptr(4),
								NumberOfEvaluationPeriods: to.Float32Ptr(4),
							},
							Operator: armmonitor.DynamicThresholdOperatorGreaterOrLessThan.ToPtr(),
						}},
				},
				Enabled:             to.BoolPtr(true),
				EvaluationFrequency: to.StringPtr("<evaluation-frequency>"),
				Scopes: []*string{
					to.StringPtr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1"),
					to.StringPtr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2")},
				Severity:             to.Int32Ptr(3),
				TargetResourceRegion: to.StringPtr("<target-resource-region>"),
				TargetResourceType:   to.StringPtr("<target-resource-type>"),
				WindowSize:           to.StringPtr("<window-size>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("MetricAlertResource.ID: %s\n", *res.ID)
}
```

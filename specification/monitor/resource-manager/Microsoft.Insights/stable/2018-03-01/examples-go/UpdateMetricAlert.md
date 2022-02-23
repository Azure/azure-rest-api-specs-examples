Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.4.1/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/UpdateMetricAlert.json
func ExampleMetricAlertsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewMetricAlertsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<rule-name>",
		armmonitor.MetricAlertResourcePatch{
			Properties: &armmonitor.MetricAlertPropertiesPatch{
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
				Criteria: &armmonitor.MetricAlertSingleResourceMultipleMetricCriteria{
					ODataType: armmonitor.Odatatype("Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria").ToPtr(),
					AllOf: []*armmonitor.MetricCriteria{
						{
							Name:            to.StringPtr("<name>"),
							CriterionType:   armmonitor.CriterionType("StaticThresholdCriterion").ToPtr(),
							Dimensions:      []*armmonitor.MetricDimension{},
							MetricName:      to.StringPtr("<metric-name>"),
							TimeAggregation: armmonitor.AggregationTypeEnum("Average").ToPtr(),
							Operator:        armmonitor.Operator("GreaterThan").ToPtr(),
							Threshold:       to.Float64Ptr(80.5),
						}},
				},
				Enabled:             to.BoolPtr(true),
				EvaluationFrequency: to.StringPtr("<evaluation-frequency>"),
				Scopes: []*string{
					to.StringPtr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme")},
				Severity:   to.Int32Ptr(3),
				WindowSize: to.StringPtr("<window-size>"),
			},
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MetricAlertsClientUpdateResult)
}
```

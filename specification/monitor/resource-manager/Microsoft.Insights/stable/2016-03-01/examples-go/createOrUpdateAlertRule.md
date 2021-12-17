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

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/createOrUpdateAlertRule.json
func ExampleAlertRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewAlertRulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<rule-name>",
		armmonitor.AlertRuleResource{
			Resource: armmonitor.Resource{
				Location: to.StringPtr("<location>"),
				Tags:     map[string]*string{},
			},
			Properties: &armmonitor.AlertRule{
				Name:        to.StringPtr("<name>"),
				Description: to.StringPtr("<description>"),
				Actions:     []armmonitor.RuleActionClassification{},
				Condition: &armmonitor.ThresholdRuleCondition{
					RuleCondition: armmonitor.RuleCondition{
						DataSource: &armmonitor.RuleMetricDataSource{
							RuleDataSource: armmonitor.RuleDataSource{
								ODataType:   to.StringPtr("<odata-type>"),
								ResourceURI: to.StringPtr("<resource-uri>"),
							},
							MetricName: to.StringPtr("<metric-name>"),
						},
						ODataType: to.StringPtr("<odata-type>"),
					},
					Operator:        armmonitor.ConditionOperatorGreaterThan.ToPtr(),
					Threshold:       to.Float64Ptr(3),
					TimeAggregation: armmonitor.TimeAggregationOperatorTotal.ToPtr(),
					WindowSize:      to.StringPtr("<window-size>"),
				},
				IsEnabled: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AlertRuleResource.ID: %s\n", *res.ID)
}
```

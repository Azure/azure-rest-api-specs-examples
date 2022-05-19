Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.7.0/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/createOrUpdateAlertRule.json
func ExampleAlertRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewAlertRulesClient("b67f7fec-69fc-4974-9099-a26bd6ffeda3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"Rac46PostSwapRG",
		"chiricutin",
		armmonitor.AlertRuleResource{
			Location: to.Ptr("West US"),
			Tags:     map[string]*string{},
			Properties: &armmonitor.AlertRule{
				Name:        to.Ptr("chiricutin"),
				Description: to.Ptr("Pura Vida"),
				Actions:     []armmonitor.RuleActionClassification{},
				Condition: &armmonitor.ThresholdRuleCondition{
					DataSource: &armmonitor.RuleMetricDataSource{
						ODataType:   to.Ptr("Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource"),
						ResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest"),
						MetricName:  to.Ptr("Requests"),
					},
					ODataType:       to.Ptr("Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition"),
					Operator:        to.Ptr(armmonitor.ConditionOperatorGreaterThan),
					Threshold:       to.Ptr[float64](3),
					TimeAggregation: to.Ptr(armmonitor.TimeAggregationOperatorTotal),
					WindowSize:      to.Ptr("PT5M"),
				},
				IsEnabled: to.Ptr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

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

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/createOrUpdateScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewScheduledQueryRulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<rule-name>",
		armmonitor.LogSearchRuleResource{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armmonitor.LogSearchRule{
				Description: to.StringPtr("<description>"),
				Action: &armmonitor.AlertingAction{
					ODataType: to.StringPtr("<odata-type>"),
					AznsAction: &armmonitor.AzNsActionGroup{
						ActionGroup:          []*string{},
						CustomWebhookPayload: to.StringPtr("<custom-webhook-payload>"),
						EmailSubject:         to.StringPtr("<email-subject>"),
					},
					Severity: armmonitor.AlertSeverity("1").ToPtr(),
					Trigger: &armmonitor.TriggerCondition{
						MetricTrigger: &armmonitor.LogMetricTrigger{
							MetricColumn:      to.StringPtr("<metric-column>"),
							MetricTriggerType: armmonitor.MetricTriggerType("Consecutive").ToPtr(),
							Threshold:         to.Float64Ptr(5),
							ThresholdOperator: armmonitor.ConditionalOperator("GreaterThan").ToPtr(),
						},
						Threshold:         to.Float64Ptr(3),
						ThresholdOperator: armmonitor.ConditionalOperator("GreaterThan").ToPtr(),
					},
				},
				Enabled: armmonitor.Enabled("true").ToPtr(),
				Schedule: &armmonitor.Schedule{
					FrequencyInMinutes:  to.Int32Ptr(15),
					TimeWindowInMinutes: to.Int32Ptr(15),
				},
				Source: &armmonitor.Source{
					DataSourceID: to.StringPtr("<data-source-id>"),
					Query:        to.StringPtr("<query>"),
					QueryType:    armmonitor.QueryType("ResultCount").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ScheduledQueryRulesClientCreateOrUpdateResult)
}
```

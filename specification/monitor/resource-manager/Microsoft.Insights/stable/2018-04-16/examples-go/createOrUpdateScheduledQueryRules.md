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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/examples/createOrUpdateScheduledQueryRules.json
func ExampleScheduledQueryRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewScheduledQueryRulesClient("b67f7fec-69fc-4974-9099-a26bd6ffeda3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"Rac46PostSwapRG",
		"logalertfoo",
		armmonitor.LogSearchRuleResource{
			Location: to.Ptr("eastus"),
			Tags:     map[string]*string{},
			Properties: &armmonitor.LogSearchRule{
				Description: to.Ptr("log alert description"),
				Action: &armmonitor.AlertingAction{
					ODataType: to.Ptr("Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.AlertingAction"),
					AznsAction: &armmonitor.AzNsActionGroup{
						ActionGroup:          []*string{},
						CustomWebhookPayload: to.Ptr("{}"),
						EmailSubject:         to.Ptr("Email Header"),
					},
					Severity: to.Ptr(armmonitor.AlertSeverityOne),
					Trigger: &armmonitor.TriggerCondition{
						MetricTrigger: &armmonitor.LogMetricTrigger{
							MetricColumn:      to.Ptr("Computer"),
							MetricTriggerType: to.Ptr(armmonitor.MetricTriggerTypeConsecutive),
							Threshold:         to.Ptr[float64](5),
							ThresholdOperator: to.Ptr(armmonitor.ConditionalOperatorGreaterThan),
						},
						Threshold:         to.Ptr[float64](3),
						ThresholdOperator: to.Ptr(armmonitor.ConditionalOperatorGreaterThan),
					},
				},
				Enabled: to.Ptr(armmonitor.EnabledTrue),
				Schedule: &armmonitor.Schedule{
					FrequencyInMinutes:  to.Ptr[int32](15),
					TimeWindowInMinutes: to.Ptr[int32](15),
				},
				Source: &armmonitor.Source{
					DataSourceID: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.OperationalInsights/workspaces/sampleWorkspace"),
					Query:        to.Ptr("Heartbeat | summarize AggregatedValue = count() by bin(TimeGenerated, 5m)"),
					QueryType:    to.Ptr(armmonitor.QueryTypeResultCount),
				},
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

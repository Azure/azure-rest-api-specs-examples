package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/getAlertRule.json
func ExampleAlertRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertRulesClient().Get(ctx, "Rac46PostSwapRG", "chiricutin", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertRuleResource = armmonitor.AlertRuleResource{
	// 	Name: to.Ptr("chiricutin"),
	// 	Type: to.Ptr("Microsoft.Insights/alertRules"),
	// 	ID: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/alertrules/chiricutin"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"$type": to.Ptr("Microsoft.WindowsAzure.Management.Common.Storage.CasePreservedDictionary, Microsoft.WindowsAzure.Management.Common.Storage"),
	// 		"hidden-link:/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest": to.Ptr("Resource"),
	// 	},
	// 	Properties: &armmonitor.AlertRule{
	// 		Name: to.Ptr("chiricutin"),
	// 		Description: to.Ptr("Pura Vida"),
	// 		Actions: []armmonitor.RuleActionClassification{
	// 		},
	// 		Condition: &armmonitor.ThresholdRuleCondition{
	// 			DataSource: &armmonitor.RuleMetricDataSource{
	// 				ODataType: to.Ptr("Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource"),
	// 				ResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest"),
	// 				MetricName: to.Ptr("Requests"),
	// 			},
	// 			ODataType: to.Ptr("Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition"),
	// 			Operator: to.Ptr(armmonitor.ConditionOperatorGreaterThan),
	// 			Threshold: to.Ptr[float64](3),
	// 			TimeAggregation: to.Ptr(armmonitor.TimeAggregationOperatorTotal),
	// 			WindowSize: to.Ptr("PT5M"),
	// 		},
	// 		IsEnabled: to.Ptr(true),
	// 		LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-11-23T21:23:52.022Z"); return t}()),
	// 	},
	// }
}

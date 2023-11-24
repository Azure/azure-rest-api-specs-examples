package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/listAlertRule.json
func ExampleAlertRulesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertRulesClient().NewListByResourceGroupPager("Rac46PostSwapRG", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AlertRuleResourceCollection = armmonitor.AlertRuleResourceCollection{
		// 	Value: []*armmonitor.AlertRuleResource{
		// 		{
		// 			Name: to.Ptr("myRuleName"),
		// 			Type: to.Ptr("Microsoft.Insights/alertRules"),
		// 			ID: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/alertrules/myRuleName"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"$type": to.Ptr("Microsoft.WindowsAzure.Management.Common.Storage.CasePreservedDictionary, Microsoft.WindowsAzure.Management.Common.Storage"),
		// 				"hidden-link:/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest": to.Ptr("Resource"),
		// 			},
		// 			Properties: &armmonitor.AlertRule{
		// 				Name: to.Ptr("myRuleName"),
		// 				Description: to.Ptr("Pura Vida"),
		// 				Actions: []armmonitor.RuleActionClassification{
		// 					&armmonitor.RuleEmailAction{
		// 						ODataType: to.Ptr("Microsoft.Azure.Management.Insights.Models.RuleEmailAction"),
		// 						CustomEmails: []*string{
		// 							to.Ptr("gu@ms.com"),
		// 							to.Ptr("su@ms.net")},
		// 							SendToServiceOwners: to.Ptr(true),
		// 					}},
		// 					Condition: &armmonitor.ThresholdRuleCondition{
		// 						DataSource: &armmonitor.RuleMetricDataSource{
		// 							ODataType: to.Ptr("Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource"),
		// 							ResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest"),
		// 							MetricName: to.Ptr("Requests"),
		// 						},
		// 						ODataType: to.Ptr("Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition"),
		// 						Operator: to.Ptr(armmonitor.ConditionOperatorGreaterThan),
		// 						Threshold: to.Ptr[float64](2),
		// 						TimeAggregation: to.Ptr(armmonitor.TimeAggregationOperatorTotal),
		// 						WindowSize: to.Ptr("PT5M"),
		// 					},
		// 					IsEnabled: to.Ptr(true),
		// 					LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-11-10T21:04:39.000Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("chiricutin0"),
		// 				Type: to.Ptr("Microsoft.Insights/alertRules"),
		// 				ID: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/alertrules/chiricutin0"),
		// 				Location: to.Ptr("West US"),
		// 				Tags: map[string]*string{
		// 					"$type": to.Ptr("Microsoft.WindowsAzure.Management.Common.Storage.CasePreservedDictionary, Microsoft.WindowsAzure.Management.Common.Storage"),
		// 					"hidden-link:/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest": to.Ptr("Resource"),
		// 				},
		// 				Properties: &armmonitor.AlertRule{
		// 					Name: to.Ptr("chiricutin0"),
		// 					Description: to.Ptr("Pura Vida 0"),
		// 					Actions: []armmonitor.RuleActionClassification{
		// 					},
		// 					Condition: &armmonitor.ThresholdRuleCondition{
		// 						DataSource: &armmonitor.RuleMetricDataSource{
		// 							ODataType: to.Ptr("Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource"),
		// 							ResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest"),
		// 							MetricName: to.Ptr("Requests"),
		// 						},
		// 						ODataType: to.Ptr("Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition"),
		// 						Operator: to.Ptr(armmonitor.ConditionOperatorGreaterThan),
		// 						Threshold: to.Ptr[float64](2),
		// 						TimeAggregation: to.Ptr(armmonitor.TimeAggregationOperatorTotal),
		// 						WindowSize: to.Ptr("PT5M"),
		// 					},
		// 					IsEnabled: to.Ptr(true),
		// 					LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-11-10T21:04:39.108Z"); return t}()),
		// 				},
		// 		}},
		// 	}
	}
}

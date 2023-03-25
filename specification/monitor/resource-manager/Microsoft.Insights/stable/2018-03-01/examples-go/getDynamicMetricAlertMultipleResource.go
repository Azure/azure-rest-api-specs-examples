package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getDynamicMetricAlertMultipleResource.json
func ExampleMetricAlertsClient_Get_getADynamicAlertRuleForMultipleResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricAlertsClient().Get(ctx, "gigtest", "MetricAlertOnMultipleResources", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetricAlertResource = armmonitor.MetricAlertResource{
	// 	Type: to.Ptr("Microsoft.Insights/metricAlerts"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/providers/microsoft.insights/metricalerts/MetricAlertOnMultipleResources"),
	// 	Location: to.Ptr("global"),
	// 	Tags: map[string]*string{
	// 		"hidden-link:/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest": to.Ptr("Resource"),
	// 	},
	// 	Properties: &armmonitor.MetricAlertProperties{
	// 		Description: to.Ptr("This is the description of the rule1"),
	// 		Actions: []*armmonitor.MetricAlertAction{
	// 			{
	// 				ActionGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
	// 				WebHookProperties: map[string]*string{
	// 					"key11": to.Ptr("value11"),
	// 					"key12": to.Ptr("value12"),
	// 				},
	// 		}},
	// 		AutoMitigate: to.Ptr(false),
	// 		Criteria: &armmonitor.MetricAlertMultipleResourceMultipleMetricCriteria{
	// 			ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria),
	// 			AllOf: []armmonitor.MultiMetricCriteriaClassification{
	// 				&armmonitor.DynamicMetricCriteria{
	// 					Name: to.Ptr("High_CPU_80"),
	// 					CriterionType: to.Ptr(armmonitor.CriterionTypeDynamicThresholdCriterion),
	// 					Dimensions: []*armmonitor.MetricDimension{
	// 					},
	// 					MetricName: to.Ptr("Percentage CPU"),
	// 					MetricNamespace: to.Ptr("microsoft.compute/virtualmachines"),
	// 					TimeAggregation: to.Ptr(armmonitor.AggregationTypeEnumAverage),
	// 					AlertSensitivity: to.Ptr(armmonitor.DynamicThresholdSensitivityMedium),
	// 					FailingPeriods: &armmonitor.DynamicThresholdFailingPeriods{
	// 						MinFailingPeriodsToAlert: to.Ptr[float32](4),
	// 						NumberOfEvaluationPeriods: to.Ptr[float32](4),
	// 					},
	// 					Operator: to.Ptr(armmonitor.DynamicThresholdOperatorGreaterOrLessThan),
	// 			}},
	// 		},
	// 		Enabled: to.Ptr(true),
	// 		EvaluationFrequency: to.Ptr("PT1M"),
	// 		Scopes: []*string{
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1"),
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2")},
	// 			Severity: to.Ptr[int32](3),
	// 			TargetResourceRegion: to.Ptr("southcentralus"),
	// 			TargetResourceType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 			WindowSize: to.Ptr("PT15M"),
	// 		},
	// 	}
}

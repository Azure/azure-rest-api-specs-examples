package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/listMetricAlert.json
func ExampleMetricAlertsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetricAlertsClient().NewListBySubscriptionPager(nil)
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
		// page.MetricAlertResourceCollection = armmonitor.MetricAlertResourceCollection{
		// 	Value: []*armmonitor.MetricAlertResource{
		// 		{
		// 			Type: to.Ptr("Microsoft.Insights/metricAlerts"),
		// 			ID: to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/providers/microsoft.insights/metricalerts/highcpu"),
		// 			Location: to.Ptr("global"),
		// 			Tags: map[string]*string{
		// 				"hidden-link:/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest": to.Ptr("Resource"),
		// 			},
		// 			Properties: &armmonitor.MetricAlertProperties{
		// 				Description: to.Ptr("This is the description of the rule1"),
		// 				Actions: []*armmonitor.MetricAlertAction{
		// 					{
		// 						ActionGroupID: to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
		// 						WebHookProperties: map[string]*string{
		// 							"key11": to.Ptr("value11"),
		// 							"key12": to.Ptr("value12"),
		// 						},
		// 				}},
		// 				Criteria: &armmonitor.MetricAlertSingleResourceMultipleMetricCriteria{
		// 					ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria),
		// 					AllOf: []*armmonitor.MetricCriteria{
		// 						{
		// 							Name: to.Ptr("High_CPU_80"),
		// 							CriterionType: to.Ptr(armmonitor.CriterionTypeStaticThresholdCriterion),
		// 							Dimensions: []*armmonitor.MetricDimension{
		// 							},
		// 							MetricName: to.Ptr("CustomMetric_\\Processor(_Total)\\% Processor Time"),
		// 							TimeAggregation: to.Ptr(armmonitor.AggregationTypeEnumAverage),
		// 							Operator: to.Ptr(armmonitor.OperatorGreaterThan),
		// 							Threshold: to.Ptr[float64](80.5),
		// 					}},
		// 				},
		// 				Enabled: to.Ptr(true),
		// 				EvaluationFrequency: to.Ptr("Pt1m"),
		// 				Scopes: []*string{
		// 					to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme")},
		// 					Severity: to.Ptr[int32](3),
		// 					WindowSize: to.Ptr("Pt15m"),
		// 				},
		// 			},
		// 			{
		// 				Type: to.Ptr("Microsoft.Insights/metricAlerts"),
		// 				ID: to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/providers/microsoft.insights/metricalerts/mediumcpu"),
		// 				Location: to.Ptr("global"),
		// 				Tags: map[string]*string{
		// 					"hidden-link:/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest": to.Ptr("Resource"),
		// 				},
		// 				Properties: &armmonitor.MetricAlertProperties{
		// 					Description: to.Ptr("Med cpu alert"),
		// 					Actions: []*armmonitor.MetricAlertAction{
		// 						{
		// 							ActionGroupID: to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
		// 							WebHookProperties: map[string]*string{
		// 								"key11": to.Ptr("value11"),
		// 								"key12": to.Ptr("value12"),
		// 							},
		// 					}},
		// 					Criteria: &armmonitor.MetricAlertSingleResourceMultipleMetricCriteria{
		// 						ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria),
		// 						AllOf: []*armmonitor.MetricCriteria{
		// 							{
		// 								Name: to.Ptr("Medium_CPU_50"),
		// 								CriterionType: to.Ptr(armmonitor.CriterionTypeStaticThresholdCriterion),
		// 								Dimensions: []*armmonitor.MetricDimension{
		// 								},
		// 								MetricName: to.Ptr("CustomMetric_\\Processor(_Total)\\% Processor Time"),
		// 								TimeAggregation: to.Ptr(armmonitor.AggregationTypeEnumAverage),
		// 								Operator: to.Ptr(armmonitor.OperatorGreaterThan),
		// 								Threshold: to.Ptr[float64](50.5),
		// 						}},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					EvaluationFrequency: to.Ptr("Pt1m"),
		// 					Scopes: []*string{
		// 						to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme")},
		// 						Severity: to.Ptr[int32](4),
		// 						WindowSize: to.Ptr("Pt15m"),
		// 					},
		// 			}},
		// 		}
	}
}

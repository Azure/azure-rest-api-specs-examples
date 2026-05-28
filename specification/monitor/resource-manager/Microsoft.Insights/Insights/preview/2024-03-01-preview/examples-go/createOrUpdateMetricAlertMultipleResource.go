package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-03-01-preview/createOrUpdateMetricAlertMultipleResource.json
func ExampleMetricAlertsClient_CreateOrUpdate_createOrUpdateAnAlertRuleForMultipleResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricAlertsClient().CreateOrUpdate(ctx, "gigtest", "MetricAlertOnMultipleResources", armmonitor.MetricAlertResource{
		Location: to.Ptr("global"),
		Properties: &armmonitor.MetricAlertProperties{
			Description: to.Ptr("This is the description of the rule1"),
			Actions: []*armmonitor.MetricAlertAction{
				{
					ActionGroupID: to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
					WebHookProperties: map[string]*string{
						"key11": to.Ptr("value11"),
						"key12": to.Ptr("value12"),
					},
				},
			},
			AutoMitigate: to.Ptr(true),
			Criteria: &armmonitor.MetricAlertMultipleResourceMultipleMetricCriteria{
				AllOf: []armmonitor.MultiMetricCriteriaClassification{
					&armmonitor.MetricCriteria{
						Name:            to.Ptr("High_CPU_80"),
						CriterionType:   to.Ptr(armmonitor.CriterionTypeStaticThresholdCriterion),
						Dimensions:      []*armmonitor.MetricDimension{},
						MetricName:      to.Ptr("Percentage CPU"),
						MetricNamespace: to.Ptr("microsoft.compute/virtualmachines"),
						Operator:        to.Ptr(armmonitor.OperatorGreaterThan),
						Threshold:       to.Ptr[float64](80.5),
						TimeAggregation: to.Ptr(armmonitor.AggregationTypeEnumAverage),
					},
				},
				ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria),
			},
			Enabled:             to.Ptr(true),
			EvaluationFrequency: to.Ptr("PT1M"),
			Scopes: []*string{
				to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1"),
				to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2"),
			},
			Severity:             to.Ptr[int32](3),
			TargetResourceRegion: to.Ptr("southcentralus"),
			TargetResourceType:   to.Ptr("Microsoft.Compute/virtualMachines"),
			WindowSize:           to.Ptr("PT15M"),
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.MetricAlertsClientCreateOrUpdateResponse{
	// 	MetricAlertResource: armmonitor.MetricAlertResource{
	// 		Type: to.Ptr("Microsoft.Insights/metricAlerts"),
	// 		ID: to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/providers/microsoft.insights/metricalerts/MetricAlertOnMultipleResources"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armmonitor.MetricAlertProperties{
	// 			Description: to.Ptr("This is the description of the rule1"),
	// 			Actions: []*armmonitor.MetricAlertAction{
	// 				{
	// 					ActionGroupID: to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
	// 					WebHookProperties: map[string]*string{
	// 						"key11": to.Ptr("value11"),
	// 						"key12": to.Ptr("value12"),
	// 					},
	// 				},
	// 			},
	// 			AutoMitigate: to.Ptr(true),
	// 			Criteria: &armmonitor.MetricAlertMultipleResourceMultipleMetricCriteria{
	// 				AllOf: []armmonitor.MultiMetricCriteriaClassification{
	// 					&armmonitor.MetricCriteria{
	// 						Name: to.Ptr("High_CPU_80"),
	// 						CriterionType: to.Ptr(armmonitor.CriterionTypeStaticThresholdCriterion),
	// 						Dimensions: []*armmonitor.MetricDimension{
	// 						},
	// 						MetricName: to.Ptr("Percentage CPU"),
	// 						MetricNamespace: to.Ptr("microsoft.compute/virtualmachines"),
	// 						Operator: to.Ptr(armmonitor.OperatorGreaterThan),
	// 						Threshold: to.Ptr[float64](80.5),
	// 						TimeAggregation: to.Ptr(armmonitor.AggregationTypeEnumAverage),
	// 					},
	// 				},
	// 				ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria),
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			EvaluationFrequency: to.Ptr("PT1M"),
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1"),
	// 				to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2"),
	// 			},
	// 			Severity: to.Ptr[int32](3),
	// 			TargetResourceRegion: to.Ptr("southcentralus"),
	// 			TargetResourceType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 			WindowSize: to.Ptr("PT15M"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"hidden-link:/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest": to.Ptr("Resource"),
	// 		},
	// 	},
	// }
}

package armmonitor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateDynamicMetricAlertSingleResource.json
func ExampleMetricAlertsClient_CreateOrUpdate_createOrUpdateADynamicAlertRuleForSingleResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewMetricAlertsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "gigtest", "chiricutin", armmonitor.MetricAlertResource{
		Location: to.Ptr("global"),
		Tags:     map[string]*string{},
		Properties: &armmonitor.MetricAlertProperties{
			Description: to.Ptr("This is the description of the rule1"),
			Actions: []*armmonitor.MetricAlertAction{
				{
					ActionGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
					WebHookProperties: map[string]*string{
						"key11": to.Ptr("value11"),
						"key12": to.Ptr("value12"),
					},
				}},
			AutoMitigate: to.Ptr(true),
			Criteria: &armmonitor.MetricAlertMultipleResourceMultipleMetricCriteria{
				ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria),
				AllOf: []armmonitor.MultiMetricCriteriaClassification{
					&armmonitor.DynamicMetricCriteria{
						Name:             to.Ptr("High_CPU_80"),
						CriterionType:    to.Ptr(armmonitor.CriterionTypeDynamicThresholdCriterion),
						Dimensions:       []*armmonitor.MetricDimension{},
						MetricName:       to.Ptr("Percentage CPU"),
						MetricNamespace:  to.Ptr("microsoft.compute/virtualmachines"),
						TimeAggregation:  to.Ptr(armmonitor.AggregationTypeEnumAverage),
						AlertSensitivity: to.Ptr(armmonitor.DynamicThresholdSensitivityMedium),
						FailingPeriods: &armmonitor.DynamicThresholdFailingPeriods{
							MinFailingPeriodsToAlert:  to.Ptr[float32](4),
							NumberOfEvaluationPeriods: to.Ptr[float32](4),
						},
						IgnoreDataBefore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-04T21:00:00.000Z"); return t }()),
						Operator:         to.Ptr(armmonitor.DynamicThresholdOperatorGreaterOrLessThan),
					}},
			},
			Enabled:             to.Ptr(true),
			EvaluationFrequency: to.Ptr("PT1M"),
			Scopes: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme")},
			Severity:   to.Ptr[int32](3),
			WindowSize: to.Ptr("PT15M"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateMetricAlertWithDimensions.json
func ExampleMetricAlertsClient_CreateOrUpdate_createOrUpdateAnAlertRulesWithDimensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewMetricAlertsClient("14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "gigtest", "MetricAlertOnMultipleDimensions", armmonitor.MetricAlertResource{
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
					&armmonitor.MetricCriteria{
						Name:          to.Ptr("Metric1"),
						CriterionType: to.Ptr(armmonitor.CriterionTypeStaticThresholdCriterion),
						Dimensions: []*armmonitor.MetricDimension{
							{
								Name:     to.Ptr("ActivityName"),
								Operator: to.Ptr("Include"),
								Values: []*string{
									to.Ptr("*")},
							},
							{
								Name:     to.Ptr("StatusCode"),
								Operator: to.Ptr("Include"),
								Values: []*string{
									to.Ptr("200")},
							}},
						MetricName:      to.Ptr("Availability"),
						MetricNamespace: to.Ptr("Microsoft.KeyVault/vaults"),
						TimeAggregation: to.Ptr(armmonitor.AggregationTypeEnumAverage),
						Operator:        to.Ptr(armmonitor.OperatorGreaterThan),
						Threshold:       to.Ptr[float64](55),
					}},
			},
			Enabled:             to.Ptr(true),
			EvaluationFrequency: to.Ptr("PT1H"),
			Scopes: []*string{
				to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.KeyVault/vaults/keyVaultResource")},
			Severity:   to.Ptr[int32](3),
			WindowSize: to.Ptr("P1D"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

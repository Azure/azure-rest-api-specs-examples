package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2022-08-01-preview/examples/createOrUpdateScheduledQueryRuleSubscription.json
func ExampleScheduledQueryRulesClient_CreateOrUpdate_createOrUpdateAScheduledQueryRuleOnSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledQueryRulesClient().CreateOrUpdate(ctx, "QueryResourceGroupName", "perf", armmonitor.ScheduledQueryRuleResource{
		Location: to.Ptr("eastus"),
		Properties: &armmonitor.ScheduledQueryRuleProperties{
			Description: to.Ptr("Performance rule"),
			Actions: &armmonitor.Actions{
				ActionGroups: []*string{
					to.Ptr("/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup")},
				CustomProperties: map[string]*string{
					"key11": to.Ptr("value11"),
					"key12": to.Ptr("value12"),
				},
			},
			CheckWorkspaceAlertsStorageConfigured: to.Ptr(true),
			Criteria: &armmonitor.ScheduledQueryRuleCriteria{
				AllOf: []*armmonitor.Condition{
					{
						Dimensions: []*armmonitor.Dimension{
							{
								Name:     to.Ptr("ComputerIp"),
								Operator: to.Ptr(armmonitor.DimensionOperatorExclude),
								Values: []*string{
									to.Ptr("192.168.1.1")},
							},
							{
								Name:     to.Ptr("OSType"),
								Operator: to.Ptr(armmonitor.DimensionOperatorInclude),
								Values: []*string{
									to.Ptr("*")},
							}},
						FailingPeriods: &armmonitor.ConditionFailingPeriods{
							MinFailingPeriodsToAlert:  to.Ptr[int64](1),
							NumberOfEvaluationPeriods: to.Ptr[int64](1),
						},
						MetricMeasureColumn: to.Ptr("% Processor Time"),
						Operator:            to.Ptr(armmonitor.ConditionOperatorGreaterThan),
						Query:               to.Ptr("Perf | where ObjectName == \"Processor\""),
						ResourceIDColumn:    to.Ptr("resourceId"),
						Threshold:           to.Ptr[float64](70),
						TimeAggregation:     to.Ptr(armmonitor.TimeAggregationAverage),
					}},
			},
			Enabled:             to.Ptr(true),
			EvaluationFrequency: to.Ptr("PT5M"),
			MuteActionsDuration: to.Ptr("PT30M"),
			RuleResolveConfiguration: &armmonitor.RuleResolveConfiguration{
				AutoResolved:  to.Ptr(true),
				TimeToResolve: to.Ptr("PT10M"),
			},
			Scopes: []*string{
				to.Ptr("/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147")},
			Severity:            to.Ptr(armmonitor.AlertSeverity(4)),
			SkipQueryValidation: to.Ptr(true),
			TargetResourceTypes: []*string{
				to.Ptr("Microsoft.Compute/virtualMachines")},
			WindowSize: to.Ptr("PT10M"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScheduledQueryRuleResource = armmonitor.ScheduledQueryRuleResource{
	// 	Name: to.Ptr("perf"),
	// 	Type: to.Ptr("microsoft.insights/scheduledqueryrules"),
	// 	ID: to.Ptr("/subscriptions/dd4bfc94-a096-412b-9c43-4bd13e35afbc/resourcegroups/QueryResourceGroupName/providers/microsoft.insights/scheduledqueryrules/perf"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armmonitor.ScheduledQueryRuleProperties{
	// 		Description: to.Ptr("Performance rule"),
	// 		Actions: &armmonitor.Actions{
	// 			ActionGroups: []*string{
	// 				to.Ptr("/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup")},
	// 				CustomProperties: map[string]*string{
	// 					"key11": to.Ptr("value11"),
	// 					"key12": to.Ptr("value12"),
	// 				},
	// 			},
	// 			CheckWorkspaceAlertsStorageConfigured: to.Ptr(true),
	// 			Criteria: &armmonitor.ScheduledQueryRuleCriteria{
	// 				AllOf: []*armmonitor.Condition{
	// 					{
	// 						Dimensions: []*armmonitor.Dimension{
	// 							{
	// 								Name: to.Ptr("ComputerIp"),
	// 								Operator: to.Ptr(armmonitor.DimensionOperatorExclude),
	// 								Values: []*string{
	// 									to.Ptr("192.168.1.1")},
	// 								},
	// 								{
	// 									Name: to.Ptr("OSType"),
	// 									Operator: to.Ptr(armmonitor.DimensionOperatorInclude),
	// 									Values: []*string{
	// 										to.Ptr("*")},
	// 								}},
	// 								FailingPeriods: &armmonitor.ConditionFailingPeriods{
	// 									MinFailingPeriodsToAlert: to.Ptr[int64](1),
	// 									NumberOfEvaluationPeriods: to.Ptr[int64](1),
	// 								},
	// 								MetricMeasureColumn: to.Ptr("% Processor Time"),
	// 								Operator: to.Ptr(armmonitor.ConditionOperatorGreaterThan),
	// 								Query: to.Ptr("Perf | where ObjectName == \"Processor\""),
	// 								ResourceIDColumn: to.Ptr("resourceId"),
	// 								Threshold: to.Ptr[float64](70),
	// 								TimeAggregation: to.Ptr(armmonitor.TimeAggregationAverage),
	// 						}},
	// 					},
	// 					Enabled: to.Ptr(true),
	// 					EvaluationFrequency: to.Ptr("PT5M"),
	// 					IsWorkspaceAlertsStorageConfigured: to.Ptr(true),
	// 					MuteActionsDuration: to.Ptr("PT30M"),
	// 					RuleResolveConfiguration: &armmonitor.RuleResolveConfiguration{
	// 						AutoResolved: to.Ptr(true),
	// 						TimeToResolve: to.Ptr("PT10M"),
	// 					},
	// 					Scopes: []*string{
	// 						to.Ptr("/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147")},
	// 						Severity: to.Ptr(armmonitor.AlertSeverity(4)),
	// 						SkipQueryValidation: to.Ptr(true),
	// 						TargetResourceTypes: []*string{
	// 							to.Ptr("Microsoft.Compute/virtualMachines")},
	// 							WindowSize: to.Ptr("PT10M"),
	// 						},
	// 						Tags: map[string]*string{
	// 						},
	// 					}
}

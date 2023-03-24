package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/getAutoscaleSetting.json
func ExampleAutoscaleSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutoscaleSettingsClient().Get(ctx, "TestingMetricsScaleSet", "MySetting", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AutoscaleSettingResource = armmonitor.AutoscaleSettingResource{
	// 	Name: to.Ptr("MySetting"),
	// 	Type: to.Ptr("Microsoft.Insights/autoscaleSettings"),
	// 	ID: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/microsoft.insights/autoscalesettings/MySetting"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armmonitor.AutoscaleSetting{
	// 		Name: to.Ptr("MySetting"),
	// 		Enabled: to.Ptr(true),
	// 		Notifications: []*armmonitor.AutoscaleNotification{
	// 			{
	// 				Email: &armmonitor.EmailNotification{
	// 					CustomEmails: []*string{
	// 						to.Ptr("gu@ms.com"),
	// 						to.Ptr("ge@ns.net")},
	// 						SendToSubscriptionAdministrator: to.Ptr(true),
	// 						SendToSubscriptionCoAdministrators: to.Ptr(true),
	// 					},
	// 					Operation: to.Ptr("Scale"),
	// 					Webhooks: []*armmonitor.WebhookNotification{
	// 						{
	// 							Properties: map[string]*string{
	// 							},
	// 							ServiceURI: to.Ptr("http://myservice.com"),
	// 					}},
	// 			}},
	// 			PredictiveAutoscalePolicy: &armmonitor.PredictiveAutoscalePolicy{
	// 				ScaleMode: to.Ptr(armmonitor.PredictiveAutoscalePolicyScaleModeEnabled),
	// 			},
	// 			Profiles: []*armmonitor.AutoscaleProfile{
	// 				{
	// 					Name: to.Ptr("adios"),
	// 					Capacity: &armmonitor.ScaleCapacity{
	// 						Default: to.Ptr("1"),
	// 						Maximum: to.Ptr("10"),
	// 						Minimum: to.Ptr("1"),
	// 					},
	// 					FixedDate: &armmonitor.TimeWindow{
	// 						End: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-05T14:30:00Z"); return t}()),
	// 						Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-05T14:00:00Z"); return t}()),
	// 						TimeZone: to.Ptr("UTC"),
	// 					},
	// 					Rules: []*armmonitor.ScaleRule{
	// 						{
	// 							MetricTrigger: &armmonitor.MetricTrigger{
	// 								DividePerInstance: to.Ptr(false),
	// 								MetricName: to.Ptr("Percentage CPU"),
	// 								MetricResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
	// 								Operator: to.Ptr(armmonitor.ComparisonOperationTypeGreaterThan),
	// 								Statistic: to.Ptr(armmonitor.MetricStatisticTypeAverage),
	// 								Threshold: to.Ptr[float64](10),
	// 								TimeAggregation: to.Ptr(armmonitor.TimeAggregationTypeAverage),
	// 								TimeGrain: to.Ptr("PT1M"),
	// 								TimeWindow: to.Ptr("PT5M"),
	// 							},
	// 							ScaleAction: &armmonitor.ScaleAction{
	// 								Type: to.Ptr(armmonitor.ScaleTypeChangeCount),
	// 								Cooldown: to.Ptr("PT5M"),
	// 								Direction: to.Ptr(armmonitor.ScaleDirectionIncrease),
	// 								Value: to.Ptr("1"),
	// 							},
	// 						},
	// 						{
	// 							MetricTrigger: &armmonitor.MetricTrigger{
	// 								DividePerInstance: to.Ptr(false),
	// 								MetricName: to.Ptr("Percentage CPU"),
	// 								MetricResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
	// 								Operator: to.Ptr(armmonitor.ComparisonOperationTypeGreaterThan),
	// 								Statistic: to.Ptr(armmonitor.MetricStatisticTypeAverage),
	// 								Threshold: to.Ptr[float64](15),
	// 								TimeAggregation: to.Ptr(armmonitor.TimeAggregationTypeAverage),
	// 								TimeGrain: to.Ptr("PT2M"),
	// 								TimeWindow: to.Ptr("PT5M"),
	// 							},
	// 							ScaleAction: &armmonitor.ScaleAction{
	// 								Type: to.Ptr(armmonitor.ScaleTypeChangeCount),
	// 								Cooldown: to.Ptr("PT6M"),
	// 								Direction: to.Ptr(armmonitor.ScaleDirectionDecrease),
	// 								Value: to.Ptr("2"),
	// 							},
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("saludos"),
	// 					Capacity: &armmonitor.ScaleCapacity{
	// 						Default: to.Ptr("1"),
	// 						Maximum: to.Ptr("10"),
	// 						Minimum: to.Ptr("1"),
	// 					},
	// 					Recurrence: &armmonitor.Recurrence{
	// 						Frequency: to.Ptr(armmonitor.RecurrenceFrequencyWeek),
	// 						Schedule: &armmonitor.RecurrentSchedule{
	// 							Days: []*string{
	// 								to.Ptr("1")},
	// 								Hours: []*int32{
	// 									to.Ptr[int32](5)},
	// 									Minutes: []*int32{
	// 										to.Ptr[int32](15)},
	// 										TimeZone: to.Ptr("UTC"),
	// 									},
	// 								},
	// 								Rules: []*armmonitor.ScaleRule{
	// 									{
	// 										MetricTrigger: &armmonitor.MetricTrigger{
	// 											DividePerInstance: to.Ptr(false),
	// 											MetricName: to.Ptr("Percentage CPU"),
	// 											MetricResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
	// 											Operator: to.Ptr(armmonitor.ComparisonOperationTypeGreaterThan),
	// 											Statistic: to.Ptr(armmonitor.MetricStatisticTypeAverage),
	// 											Threshold: to.Ptr[float64](10),
	// 											TimeAggregation: to.Ptr(armmonitor.TimeAggregationTypeAverage),
	// 											TimeGrain: to.Ptr("PT1M"),
	// 											TimeWindow: to.Ptr("PT5M"),
	// 										},
	// 										ScaleAction: &armmonitor.ScaleAction{
	// 											Type: to.Ptr(armmonitor.ScaleTypeChangeCount),
	// 											Cooldown: to.Ptr("PT5M"),
	// 											Direction: to.Ptr(armmonitor.ScaleDirectionIncrease),
	// 											Value: to.Ptr("1"),
	// 										},
	// 									},
	// 									{
	// 										MetricTrigger: &armmonitor.MetricTrigger{
	// 											DividePerInstance: to.Ptr(false),
	// 											MetricName: to.Ptr("Percentage CPU"),
	// 											MetricResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
	// 											Operator: to.Ptr(armmonitor.ComparisonOperationTypeGreaterThan),
	// 											Statistic: to.Ptr(armmonitor.MetricStatisticTypeAverage),
	// 											Threshold: to.Ptr[float64](15),
	// 											TimeAggregation: to.Ptr(armmonitor.TimeAggregationTypeAverage),
	// 											TimeGrain: to.Ptr("PT2M"),
	// 											TimeWindow: to.Ptr("PT5M"),
	// 										},
	// 										ScaleAction: &armmonitor.ScaleAction{
	// 											Type: to.Ptr(armmonitor.ScaleTypeChangeCount),
	// 											Cooldown: to.Ptr("PT6M"),
	// 											Direction: to.Ptr(armmonitor.ScaleDirectionDecrease),
	// 											Value: to.Ptr("2"),
	// 										},
	// 								}},
	// 						}},
	// 						TargetResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
	// 					},
	// 				}
}

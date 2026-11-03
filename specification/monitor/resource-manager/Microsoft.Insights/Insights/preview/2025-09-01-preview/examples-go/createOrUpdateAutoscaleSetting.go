package armmonitor_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2025-09-01-preview/createOrUpdateAutoscaleSetting.json
func ExampleAutoscaleSettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("b67f7fec-69fc-4974-9099-a26bd6ffeda3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutoscaleSettingsClient().CreateOrUpdate(ctx, "TestingMetricsScaleSet", "MySetting", armmonitor.AutoscaleSettingResource{
		Location: to.Ptr("West US"),
		Identity: &armmonitor.ManagedServiceIdentity{
			Type: to.Ptr(armmonitor.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armmonitor.UserAssignedIdentity{
				"/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity": {},
			},
		},
		Properties: &armmonitor.AutoscaleSetting{
			Enabled: to.Ptr(true),
			Notifications: []*armmonitor.AutoscaleNotification{
				{
					Email: &armmonitor.EmailNotification{
						CustomEmails: []*string{
							to.Ptr("gu@ms.com"),
							to.Ptr("ge@ns.net"),
						},
						SendToSubscriptionAdministrator:    to.Ptr(true),
						SendToSubscriptionCoAdministrators: to.Ptr(true),
					},
					Operation: to.Ptr("Scale"),
					Webhooks: []*armmonitor.WebhookNotification{
						{
							Properties: map[string]*string{},
							ServiceURI: to.Ptr("http://myservice.com"),
						},
					},
				},
			},
			PredictiveAutoscalePolicy: &armmonitor.PredictiveAutoscalePolicy{
				ScaleMode: to.Ptr(armmonitor.PredictiveAutoscalePolicyScaleModeEnabled),
			},
			Profiles: []*armmonitor.AutoscaleProfile{
				{
					Name: to.Ptr("adios"),
					Capacity: &armmonitor.ScaleCapacity{
						Default: to.Ptr("1"),
						Maximum: to.Ptr("10"),
						Minimum: to.Ptr("1"),
					},
					FixedDate: &armmonitor.TimeWindow{
						End:      to.Ptr(time.Date(2015, time.March, 5, 14, 30, 0, 0, time.UTC)),
						Start:    to.Ptr(time.Date(2015, time.March, 5, 14, 0, 0, 0, time.UTC)),
						TimeZone: to.Ptr("UTC"),
					},
					Rules: []*armmonitor.ScaleRule{
						{
							MetricTrigger: &armmonitor.MetricTrigger{
								DividePerInstance: to.Ptr(false),
								MetricName:        to.Ptr("Percentage CPU"),
								MetricResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
								Operator:          to.Ptr(armmonitor.ComparisonOperationTypeGreaterThan),
								Statistic:         to.Ptr(armmonitor.MetricStatisticTypeAverage),
								Threshold:         to.Ptr[float64](10),
								TimeAggregation:   to.Ptr(armmonitor.TimeAggregationTypeAverage),
								TimeGrain:         to.Ptr("PT1M"),
								TimeWindow:        to.Ptr("PT5M"),
							},
							ScaleAction: &armmonitor.ScaleAction{
								Type:      to.Ptr(armmonitor.ScaleTypeChangeCount),
								Cooldown:  to.Ptr("PT5M"),
								Direction: to.Ptr(armmonitor.ScaleDirectionIncrease),
								Value:     to.Ptr("1"),
							},
						},
						{
							MetricTrigger: &armmonitor.MetricTrigger{
								DividePerInstance: to.Ptr(false),
								MetricName:        to.Ptr("Percentage CPU"),
								MetricResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
								Operator:          to.Ptr(armmonitor.ComparisonOperationTypeGreaterThan),
								Statistic:         to.Ptr(armmonitor.MetricStatisticTypeAverage),
								Threshold:         to.Ptr[float64](15),
								TimeAggregation:   to.Ptr(armmonitor.TimeAggregationTypeAverage),
								TimeGrain:         to.Ptr("PT2M"),
								TimeWindow:        to.Ptr("PT5M"),
							},
							ScaleAction: &armmonitor.ScaleAction{
								Type:      to.Ptr(armmonitor.ScaleTypeChangeCount),
								Cooldown:  to.Ptr("PT6M"),
								Direction: to.Ptr(armmonitor.ScaleDirectionDecrease),
								Value:     to.Ptr("2"),
							},
						},
					},
				},
				{
					Name: to.Ptr("saludos"),
					Capacity: &armmonitor.ScaleCapacity{
						Default: to.Ptr("1"),
						Maximum: to.Ptr("10"),
						Minimum: to.Ptr("1"),
					},
					Recurrence: &armmonitor.Recurrence{
						Frequency: to.Ptr(armmonitor.RecurrenceFrequencyWeek),
						Schedule: &armmonitor.RecurrentSchedule{
							Days: []*string{
								to.Ptr("1"),
							},
							Hours: []*int32{
								to.Ptr[int32](5),
							},
							Minutes: []*int32{
								to.Ptr[int32](15),
							},
							TimeZone: to.Ptr("UTC"),
						},
					},
					Rules: []*armmonitor.ScaleRule{
						{
							MetricTrigger: &armmonitor.MetricTrigger{
								DividePerInstance: to.Ptr(false),
								MetricName:        to.Ptr("Percentage CPU"),
								MetricResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
								Operator:          to.Ptr(armmonitor.ComparisonOperationTypeGreaterThan),
								Statistic:         to.Ptr(armmonitor.MetricStatisticTypeAverage),
								Threshold:         to.Ptr[float64](10),
								TimeAggregation:   to.Ptr(armmonitor.TimeAggregationTypeAverage),
								TimeGrain:         to.Ptr("PT1M"),
								TimeWindow:        to.Ptr("PT5M"),
							},
							ScaleAction: &armmonitor.ScaleAction{
								Type:      to.Ptr(armmonitor.ScaleTypeChangeCount),
								Cooldown:  to.Ptr("PT5M"),
								Direction: to.Ptr(armmonitor.ScaleDirectionIncrease),
								Value:     to.Ptr("1"),
							},
						},
						{
							MetricTrigger: &armmonitor.MetricTrigger{
								DividePerInstance: to.Ptr(false),
								MetricName:        to.Ptr("Percentage CPU"),
								MetricResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
								Operator:          to.Ptr(armmonitor.ComparisonOperationTypeGreaterThan),
								Statistic:         to.Ptr(armmonitor.MetricStatisticTypeAverage),
								Threshold:         to.Ptr[float64](15),
								TimeAggregation:   to.Ptr(armmonitor.TimeAggregationTypeAverage),
								TimeGrain:         to.Ptr("PT2M"),
								TimeWindow:        to.Ptr("PT5M"),
							},
							ScaleAction: &armmonitor.ScaleAction{
								Type:      to.Ptr(armmonitor.ScaleTypeChangeCount),
								Cooldown:  to.Ptr("PT6M"),
								Direction: to.Ptr(armmonitor.ScaleDirectionDecrease),
								Value:     to.Ptr("2"),
							},
						},
					},
				},
			},
			TargetResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.AutoscaleSettingsClientCreateOrUpdateResponse{
	// 	AutoscaleSettingResource: armmonitor.AutoscaleSettingResource{
	// 		Name: to.Ptr("MySetting"),
	// 		Type: to.Ptr("Microsoft.Insights/autoscaleSettings"),
	// 		ID: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/microsoft.insights/autoscalesettings/MySetting"),
	// 		Location: to.Ptr("West US"),
	// 		Identity: &armmonitor.ManagedServiceIdentity{
	// 			Type: to.Ptr(armmonitor.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armmonitor.UserAssignedIdentity{
	// 				"/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity": &armmonitor.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("3b5b3f1a-9d3e-4a2b-8c7d-1e2f3a4b5c6d"),
	// 					ClientID: to.Ptr("7c9e6679-7425-40de-944b-e07fc1f90ae7"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armmonitor.AutoscaleSetting{
	// 			Name: to.Ptr("MySetting"),
	// 			Enabled: to.Ptr(true),
	// 			Notifications: []*armmonitor.AutoscaleNotification{
	// 				{
	// 					Email: &armmonitor.EmailNotification{
	// 						CustomEmails: []*string{
	// 							to.Ptr("gu@ms.com"),
	// 							to.Ptr("ge@ns.net"),
	// 						},
	// 						SendToSubscriptionAdministrator: to.Ptr(true),
	// 						SendToSubscriptionCoAdministrators: to.Ptr(true),
	// 					},
	// 					Operation: to.Ptr("Scale"),
	// 					Webhooks: []*armmonitor.WebhookNotification{
	// 						{
	// 							Properties: map[string]*string{
	// 							},
	// 							ServiceURI: to.Ptr("http://myservice.com"),
	// 						},
	// 					},
	// 				},
	// 			},
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
	// 						End: to.Ptr(time.Date(2015, time.March, 5, 14, 30, 0, 0, time.UTC)),
	// 						Start: to.Ptr(time.Date(2015, time.March, 5, 14, 0, 0, 0, time.UTC)),
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
	// 						},
	// 					},
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
	// 								to.Ptr("1"),
	// 							},
	// 							Hours: []*int32{
	// 								to.Ptr[int32](5),
	// 							},
	// 							Minutes: []*int32{
	// 								to.Ptr[int32](15),
	// 							},
	// 							TimeZone: to.Ptr("UTC"),
	// 						},
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
	// 						},
	// 					},
	// 				},
	// 			},
	// 			TargetResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}

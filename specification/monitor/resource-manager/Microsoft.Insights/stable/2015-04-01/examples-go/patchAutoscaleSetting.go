package armmonitor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/patchAutoscaleSetting.json
func ExampleAutoscaleSettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewAutoscaleSettingsClient("b67f7fec-69fc-4974-9099-a26bd6ffeda3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"TestingMetricsScaleSet",
		"MySetting",
		armmonitor.AutoscaleSettingResourcePatch{
			Properties: &armmonitor.AutoscaleSetting{
				Enabled: to.Ptr(true),
				Notifications: []*armmonitor.AutoscaleNotification{
					{
						Email: &armmonitor.EmailNotification{
							CustomEmails: []*string{
								to.Ptr("gu@ms.com"),
								to.Ptr("ge@ns.net")},
							SendToSubscriptionAdministrator:    to.Ptr(true),
							SendToSubscriptionCoAdministrators: to.Ptr(true),
						},
						Operation: to.Ptr("Scale"),
						Webhooks: []*armmonitor.WebhookNotification{
							{
								Properties: map[string]*string{},
								ServiceURI: to.Ptr("http://myservice.com"),
							}},
					}},
				Profiles: []*armmonitor.AutoscaleProfile{
					{
						Name: to.Ptr("adios"),
						Capacity: &armmonitor.ScaleCapacity{
							Default: to.Ptr("1"),
							Maximum: to.Ptr("10"),
							Minimum: to.Ptr("1"),
						},
						FixedDate: &armmonitor.TimeWindow{
							End:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-05T14:30:00Z"); return t }()),
							Start:    to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-05T14:00:00Z"); return t }()),
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
							}},
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
									to.Ptr("1")},
								Hours: []*int32{
									to.Ptr[int32](5)},
								Minutes: []*int32{
									to.Ptr[int32](15)},
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
							}},
					}},
				TargetResourceURI: to.Ptr("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
			},
			Tags: map[string]*string{
				"$type": to.Ptr("Microsoft.WindowsAzure.Management.Common.Storage.CasePreservedDictionary"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

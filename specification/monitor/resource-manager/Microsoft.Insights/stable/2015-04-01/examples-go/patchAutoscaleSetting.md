Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.4.1/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/patchAutoscaleSetting.json
func ExampleAutoscaleSettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewAutoscaleSettingsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<autoscale-setting-name>",
		armmonitor.AutoscaleSettingResourcePatch{
			Properties: &armmonitor.AutoscaleSetting{
				Enabled: to.BoolPtr(true),
				Notifications: []*armmonitor.AutoscaleNotification{
					{
						Email: &armmonitor.EmailNotification{
							CustomEmails: []*string{
								to.StringPtr("gu@ms.com"),
								to.StringPtr("ge@ns.net")},
							SendToSubscriptionAdministrator:    to.BoolPtr(true),
							SendToSubscriptionCoAdministrators: to.BoolPtr(true),
						},
						Operation: to.StringPtr("<operation>"),
						Webhooks: []*armmonitor.WebhookNotification{
							{
								Properties: map[string]*string{},
								ServiceURI: to.StringPtr("<service-uri>"),
							}},
					}},
				Profiles: []*armmonitor.AutoscaleProfile{
					{
						Name: to.StringPtr("<name>"),
						Capacity: &armmonitor.ScaleCapacity{
							Default: to.StringPtr("<default>"),
							Maximum: to.StringPtr("<maximum>"),
							Minimum: to.StringPtr("<minimum>"),
						},
						FixedDate: &armmonitor.TimeWindow{
							End:      to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-05T14:30:00Z"); return t }()),
							Start:    to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-03-05T14:00:00Z"); return t }()),
							TimeZone: to.StringPtr("<time-zone>"),
						},
						Rules: []*armmonitor.ScaleRule{
							{
								MetricTrigger: &armmonitor.MetricTrigger{
									DividePerInstance: to.BoolPtr(false),
									MetricName:        to.StringPtr("<metric-name>"),
									MetricResourceURI: to.StringPtr("<metric-resource-uri>"),
									Operator:          armmonitor.ComparisonOperationTypeGreaterThan.ToPtr(),
									Statistic:         armmonitor.MetricStatisticTypeAverage.ToPtr(),
									Threshold:         to.Float64Ptr(10),
									TimeAggregation:   armmonitor.TimeAggregationTypeAverage.ToPtr(),
									TimeGrain:         to.StringPtr("<time-grain>"),
									TimeWindow:        to.StringPtr("<time-window>"),
								},
								ScaleAction: &armmonitor.ScaleAction{
									Type:      armmonitor.ScaleTypeChangeCount.ToPtr(),
									Cooldown:  to.StringPtr("<cooldown>"),
									Direction: armmonitor.ScaleDirectionIncrease.ToPtr(),
									Value:     to.StringPtr("<value>"),
								},
							},
							{
								MetricTrigger: &armmonitor.MetricTrigger{
									DividePerInstance: to.BoolPtr(false),
									MetricName:        to.StringPtr("<metric-name>"),
									MetricResourceURI: to.StringPtr("<metric-resource-uri>"),
									Operator:          armmonitor.ComparisonOperationTypeGreaterThan.ToPtr(),
									Statistic:         armmonitor.MetricStatisticTypeAverage.ToPtr(),
									Threshold:         to.Float64Ptr(15),
									TimeAggregation:   armmonitor.TimeAggregationTypeAverage.ToPtr(),
									TimeGrain:         to.StringPtr("<time-grain>"),
									TimeWindow:        to.StringPtr("<time-window>"),
								},
								ScaleAction: &armmonitor.ScaleAction{
									Type:      armmonitor.ScaleTypeChangeCount.ToPtr(),
									Cooldown:  to.StringPtr("<cooldown>"),
									Direction: armmonitor.ScaleDirectionDecrease.ToPtr(),
									Value:     to.StringPtr("<value>"),
								},
							}},
					},
					{
						Name: to.StringPtr("<name>"),
						Capacity: &armmonitor.ScaleCapacity{
							Default: to.StringPtr("<default>"),
							Maximum: to.StringPtr("<maximum>"),
							Minimum: to.StringPtr("<minimum>"),
						},
						Recurrence: &armmonitor.Recurrence{
							Frequency: armmonitor.RecurrenceFrequencyWeek.ToPtr(),
							Schedule: &armmonitor.RecurrentSchedule{
								Days: []*string{
									to.StringPtr("1")},
								Hours: []*int32{
									to.Int32Ptr(5)},
								Minutes: []*int32{
									to.Int32Ptr(15)},
								TimeZone: to.StringPtr("<time-zone>"),
							},
						},
						Rules: []*armmonitor.ScaleRule{
							{
								MetricTrigger: &armmonitor.MetricTrigger{
									DividePerInstance: to.BoolPtr(false),
									MetricName:        to.StringPtr("<metric-name>"),
									MetricResourceURI: to.StringPtr("<metric-resource-uri>"),
									Operator:          armmonitor.ComparisonOperationTypeGreaterThan.ToPtr(),
									Statistic:         armmonitor.MetricStatisticTypeAverage.ToPtr(),
									Threshold:         to.Float64Ptr(10),
									TimeAggregation:   armmonitor.TimeAggregationTypeAverage.ToPtr(),
									TimeGrain:         to.StringPtr("<time-grain>"),
									TimeWindow:        to.StringPtr("<time-window>"),
								},
								ScaleAction: &armmonitor.ScaleAction{
									Type:      armmonitor.ScaleTypeChangeCount.ToPtr(),
									Cooldown:  to.StringPtr("<cooldown>"),
									Direction: armmonitor.ScaleDirectionIncrease.ToPtr(),
									Value:     to.StringPtr("<value>"),
								},
							},
							{
								MetricTrigger: &armmonitor.MetricTrigger{
									DividePerInstance: to.BoolPtr(false),
									MetricName:        to.StringPtr("<metric-name>"),
									MetricResourceURI: to.StringPtr("<metric-resource-uri>"),
									Operator:          armmonitor.ComparisonOperationTypeGreaterThan.ToPtr(),
									Statistic:         armmonitor.MetricStatisticTypeAverage.ToPtr(),
									Threshold:         to.Float64Ptr(15),
									TimeAggregation:   armmonitor.TimeAggregationTypeAverage.ToPtr(),
									TimeGrain:         to.StringPtr("<time-grain>"),
									TimeWindow:        to.StringPtr("<time-window>"),
								},
								ScaleAction: &armmonitor.ScaleAction{
									Type:      armmonitor.ScaleTypeChangeCount.ToPtr(),
									Cooldown:  to.StringPtr("<cooldown>"),
									Direction: armmonitor.ScaleDirectionDecrease.ToPtr(),
									Value:     to.StringPtr("<value>"),
								},
							}},
					}},
				TargetResourceURI: to.StringPtr("<target-resource-uri>"),
			},
			Tags: map[string]*string{
				"$type": to.StringPtr("Microsoft.WindowsAzure.Management.Common.Storage.CasePreservedDictionary"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AutoscaleSettingsClientUpdateResult)
}
```

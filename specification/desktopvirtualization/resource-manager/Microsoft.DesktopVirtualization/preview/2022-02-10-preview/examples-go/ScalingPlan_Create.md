```go
package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2022-02-10-preview/examples/ScalingPlan_Create.json
func ExampleScalingPlansClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewScalingPlansClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<scaling-plan-name>",
		armdesktopvirtualization.ScalingPlan{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Properties: &armdesktopvirtualization.ScalingPlanProperties{
				Description:  to.Ptr("<description>"),
				ExclusionTag: to.Ptr("<exclusion-tag>"),
				FriendlyName: to.Ptr("<friendly-name>"),
				HostPoolReferences: []*armdesktopvirtualization.ScalingHostPoolReference{
					{
						HostPoolArmPath:    to.Ptr("<host-pool-arm-path>"),
						ScalingPlanEnabled: to.Ptr(true),
					}},
				HostPoolType: to.Ptr(armdesktopvirtualization.ScalingHostPoolTypePooled),
				Schedules: []*armdesktopvirtualization.ScalingSchedule{
					{
						Name: to.Ptr("<name>"),
						DaysOfWeek: []*armdesktopvirtualization.ScalingScheduleDaysOfWeekItem{
							to.Ptr(armdesktopvirtualization.ScalingScheduleDaysOfWeekItemMonday),
							to.Ptr(armdesktopvirtualization.ScalingScheduleDaysOfWeekItemTuesday),
							to.Ptr(armdesktopvirtualization.ScalingScheduleDaysOfWeekItemWednesday),
							to.Ptr(armdesktopvirtualization.ScalingScheduleDaysOfWeekItemThursday),
							to.Ptr(armdesktopvirtualization.ScalingScheduleDaysOfWeekItemFriday)},
						OffPeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
						OffPeakStartTime: &armdesktopvirtualization.Time{
							Hour:   to.Ptr[int32](20),
							Minute: to.Ptr[int32](0),
						},
						PeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmBreadthFirst),
						PeakStartTime: &armdesktopvirtualization.Time{
							Hour:   to.Ptr[int32](8),
							Minute: to.Ptr[int32](0),
						},
						RampDownCapacityThresholdPct:   to.Ptr[int32](50),
						RampDownForceLogoffUsers:       to.Ptr(true),
						RampDownLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
						RampDownMinimumHostsPct:        to.Ptr[int32](20),
						RampDownNotificationMessage:    to.Ptr("<ramp-down-notification-message>"),
						RampDownStartTime: &armdesktopvirtualization.Time{
							Hour:   to.Ptr[int32](18),
							Minute: to.Ptr[int32](0),
						},
						RampDownWaitTimeMinutes:      to.Ptr[int32](30),
						RampUpCapacityThresholdPct:   to.Ptr[int32](80),
						RampUpLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
						RampUpMinimumHostsPct:        to.Ptr[int32](20),
						RampUpStartTime: &armdesktopvirtualization.Time{
							Hour:   to.Ptr[int32](6),
							Minute: to.Ptr[int32](0),
						},
					}},
				TimeZone: to.Ptr("<time-zone>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdesktopvirtualization%2Farmdesktopvirtualization%2Fv0.4.0/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

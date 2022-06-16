package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ScalingPlan_Create.json
func ExampleScalingPlansClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewScalingPlansClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<scaling-plan-name>",
		armdesktopvirtualization.ScalingPlan{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
			Properties: &armdesktopvirtualization.ScalingPlanProperties{
				Description:  to.StringPtr("<description>"),
				ExclusionTag: to.StringPtr("<exclusion-tag>"),
				FriendlyName: to.StringPtr("<friendly-name>"),
				HostPoolReferences: []*armdesktopvirtualization.ScalingHostPoolReference{
					{
						HostPoolArmPath:    to.StringPtr("<host-pool-arm-path>"),
						ScalingPlanEnabled: to.BoolPtr(true),
					}},
				HostPoolType: armdesktopvirtualization.ScalingHostPoolType("Pooled").ToPtr(),
				Schedules: []*armdesktopvirtualization.ScalingSchedule{
					{
						Name: to.StringPtr("<name>"),
						DaysOfWeek: []*armdesktopvirtualization.ScalingScheduleDaysOfWeekItem{
							armdesktopvirtualization.ScalingScheduleDaysOfWeekItem("Monday").ToPtr(),
							armdesktopvirtualization.ScalingScheduleDaysOfWeekItem("Tuesday").ToPtr(),
							armdesktopvirtualization.ScalingScheduleDaysOfWeekItem("Wednesday").ToPtr(),
							armdesktopvirtualization.ScalingScheduleDaysOfWeekItem("Thursday").ToPtr(),
							armdesktopvirtualization.ScalingScheduleDaysOfWeekItem("Friday").ToPtr()},
						OffPeakLoadBalancingAlgorithm: armdesktopvirtualization.SessionHostLoadBalancingAlgorithm("DepthFirst").ToPtr(),
						OffPeakStartTime: &armdesktopvirtualization.Time{
							Hour:   to.Int32Ptr(20),
							Minute: to.Int32Ptr(0),
						},
						PeakLoadBalancingAlgorithm: armdesktopvirtualization.SessionHostLoadBalancingAlgorithm("BreadthFirst").ToPtr(),
						PeakStartTime: &armdesktopvirtualization.Time{
							Hour:   to.Int32Ptr(8),
							Minute: to.Int32Ptr(0),
						},
						RampDownCapacityThresholdPct:   to.Int32Ptr(50),
						RampDownForceLogoffUsers:       to.BoolPtr(true),
						RampDownLoadBalancingAlgorithm: armdesktopvirtualization.SessionHostLoadBalancingAlgorithm("DepthFirst").ToPtr(),
						RampDownMinimumHostsPct:        to.Int32Ptr(20),
						RampDownNotificationMessage:    to.StringPtr("<ramp-down-notification-message>"),
						RampDownStartTime: &armdesktopvirtualization.Time{
							Hour:   to.Int32Ptr(18),
							Minute: to.Int32Ptr(0),
						},
						RampDownWaitTimeMinutes:      to.Int32Ptr(30),
						RampUpCapacityThresholdPct:   to.Int32Ptr(80),
						RampUpLoadBalancingAlgorithm: armdesktopvirtualization.SessionHostLoadBalancingAlgorithm("DepthFirst").ToPtr(),
						RampUpMinimumHostsPct:        to.Int32Ptr(20),
						RampUpStartTime: &armdesktopvirtualization.Time{
							Hour:   to.Int32Ptr(6),
							Minute: to.Int32Ptr(0),
						},
					}},
				TimeZone: to.StringPtr("<time-zone>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ScalingPlansClientCreateResult)
}

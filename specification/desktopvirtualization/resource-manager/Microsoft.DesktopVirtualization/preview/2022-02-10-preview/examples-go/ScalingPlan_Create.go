package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2022-02-10-preview/examples/ScalingPlan_Create.json
func ExampleScalingPlansClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewScalingPlansClient("daefabc0-95b4-48b3-b645-8a753a63c4fa", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"resourceGroup1",
		"scalingPlan1",
		armdesktopvirtualization.ScalingPlan{
			Location: to.Ptr("centralus"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Properties: &armdesktopvirtualization.ScalingPlanProperties{
				Description:  to.Ptr("Description of Scaling Plan"),
				ExclusionTag: to.Ptr("value"),
				FriendlyName: to.Ptr("Scaling Plan 1"),
				HostPoolReferences: []*armdesktopvirtualization.ScalingHostPoolReference{
					{
						HostPoolArmPath:    to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1"),
						ScalingPlanEnabled: to.Ptr(true),
					}},
				HostPoolType: to.Ptr(armdesktopvirtualization.ScalingHostPoolTypePooled),
				Schedules: []*armdesktopvirtualization.ScalingSchedule{
					{
						Name: to.Ptr("schedule1"),
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
						RampDownNotificationMessage:    to.Ptr("message"),
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
				TimeZone: to.Ptr("Central Standard Time"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

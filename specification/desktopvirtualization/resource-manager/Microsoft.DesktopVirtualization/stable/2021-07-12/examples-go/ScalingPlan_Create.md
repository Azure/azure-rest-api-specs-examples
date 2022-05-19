Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdesktopvirtualization%2Farmdesktopvirtualization%2Fv1.0.0/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```go
package armdesktopvirtualization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2021-07-12/examples/ScalingPlan_Create.json
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
				Description:  to.Ptr("des1"),
				ExclusionTag: to.Ptr("value"),
				FriendlyName: to.Ptr("friendly"),
				HostPoolReferences: []*armdesktopvirtualization.ScalingHostPoolReference{
					{
						HostPoolArmPath:    to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1"),
						ScalingPlanEnabled: to.Ptr(true),
					}},
				HostPoolType: to.Ptr(armdesktopvirtualization.HostPoolTypePersonal),
				Schedules: []*armdesktopvirtualization.ScalingSchedule{
					{
						Name: to.Ptr("schedule1"),
						DaysOfWeek: []*armdesktopvirtualization.ScalingScheduleDaysOfWeekItem{
							to.Ptr(armdesktopvirtualization.ScalingScheduleDaysOfWeekItemMonday),
							to.Ptr(armdesktopvirtualization.ScalingScheduleDaysOfWeekItemTuesday),
							to.Ptr(armdesktopvirtualization.ScalingScheduleDaysOfWeekItemWednesday),
							to.Ptr(armdesktopvirtualization.ScalingScheduleDaysOfWeekItemThursday),
							to.Ptr(armdesktopvirtualization.ScalingScheduleDaysOfWeekItemFriday)},
						OffPeakLoadBalancingAlgorithm:  to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
						OffPeakStartTime:               to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-10T20:00:00.000Z"); return t }()),
						PeakLoadBalancingAlgorithm:     to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmBreadthFirst),
						PeakStartTime:                  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-10T08:00:00.000Z"); return t }()),
						RampDownCapacityThresholdPct:   to.Ptr[int32](50),
						RampDownForceLogoffUsers:       to.Ptr(true),
						RampDownLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
						RampDownMinimumHostsPct:        to.Ptr[int32](20),
						RampDownNotificationMessage:    to.Ptr("message"),
						RampDownStartTime:              to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-10T18:00:00.000Z"); return t }()),
						RampDownWaitTimeMinutes:        to.Ptr[int32](30),
						RampUpCapacityThresholdPct:     to.Ptr[int32](80),
						RampUpLoadBalancingAlgorithm:   to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
						RampUpMinimumHostsPct:          to.Ptr[int32](20),
						RampUpStartTime:                to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-10T06:00:00.000Z"); return t }()),
					}},
				TimeZone: to.Ptr(""),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

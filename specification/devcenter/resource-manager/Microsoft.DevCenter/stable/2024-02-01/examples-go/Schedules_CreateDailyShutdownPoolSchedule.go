package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Schedules_CreateDailyShutdownPoolSchedule.json
func ExampleSchedulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSchedulesClient().BeginCreateOrUpdate(ctx, "rg1", "DevProject", "DevPool", "autoShutdown", armdevcenter.Schedule{
		Properties: &armdevcenter.ScheduleProperties{
			Type:      to.Ptr(armdevcenter.ScheduledTypeStopDevBox),
			Frequency: to.Ptr(armdevcenter.ScheduledFrequencyDaily),
			State:     to.Ptr(armdevcenter.ScheduleEnableStatusEnabled),
			Time:      to.Ptr("17:30"),
			TimeZone:  to.Ptr("America/Los_Angeles"),
		},
	}, &armdevcenter.SchedulesClientBeginCreateOrUpdateOptions{Top: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Schedule = armdevcenter.Schedule{
	// 	Name: to.Ptr("autoShutdown"),
	// 	Type: to.Ptr("Microsoft.DevCenter/pools/schedules"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/TestProject/pools/DevPool/schedules/autoShutdown"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user1"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Properties: &armdevcenter.ScheduleProperties{
	// 		Type: to.Ptr(armdevcenter.ScheduledTypeStopDevBox),
	// 		Frequency: to.Ptr(armdevcenter.ScheduledFrequencyDaily),
	// 		State: to.Ptr(armdevcenter.ScheduleEnableStatusEnabled),
	// 		Time: to.Ptr("17:30"),
	// 		TimeZone: to.Ptr("America/Los_Angeles"),
	// 		ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 	},
	// }
}

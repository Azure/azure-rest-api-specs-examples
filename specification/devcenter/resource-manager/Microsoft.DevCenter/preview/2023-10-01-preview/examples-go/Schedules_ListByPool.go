package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Schedules_ListByPool.json
func ExampleSchedulesClient_NewListByPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSchedulesClient().NewListByPoolPager("rg1", "TestProject", "DevPool", &armdevcenter.SchedulesClientListByPoolOptions{Top: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ScheduleListResult = armdevcenter.ScheduleListResult{
		// 	Value: []*armdevcenter.Schedule{
		// 		{
		// 			Name: to.Ptr("autoShutdown"),
		// 			Type: to.Ptr("Microsoft.DevCenter/pools/schedules"),
		// 			ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/TestProject/pools/DevPool/schedules/autoShutdown"),
		// 			SystemData: &armdevcenter.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user1"),
		// 				LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 			},
		// 			Properties: &armdevcenter.ScheduleProperties{
		// 				Type: to.Ptr(armdevcenter.ScheduledTypeStopDevBox),
		// 				Frequency: to.Ptr(armdevcenter.ScheduledFrequencyDaily),
		// 				State: to.Ptr(armdevcenter.ScheduleEnableStatusEnabled),
		// 				Time: to.Ptr("17:30"),
		// 				TimeZone: to.Ptr("America/Los_Angeles"),
		// 				ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

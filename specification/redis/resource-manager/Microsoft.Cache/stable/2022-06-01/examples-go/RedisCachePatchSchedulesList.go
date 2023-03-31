package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/065033d1c4087a2b009e71c0b3f0666718354ebd/specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCachePatchSchedulesList.json
func ExamplePatchSchedulesClient_NewListByRedisResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPatchSchedulesClient().NewListByRedisResourcePager("rg1", "cache1", nil)
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
		// page.PatchScheduleListResult = armredis.PatchScheduleListResult{
		// 	Value: []*armredis.PatchSchedule{
		// 		{
		// 			Name: to.Ptr("cache1/default"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis/PatchSchedules"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/patchSchedules/default"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armredis.ScheduleEntries{
		// 				ScheduleEntries: []*armredis.ScheduleEntry{
		// 					{
		// 						DayOfWeek: to.Ptr(armredis.DayOfWeekMonday),
		// 						MaintenanceWindow: to.Ptr("PT5H"),
		// 						StartHourUTC: to.Ptr[int32](12),
		// 					},
		// 					{
		// 						DayOfWeek: to.Ptr(armredis.DayOfWeekTuesday),
		// 						StartHourUTC: to.Ptr[int32](12),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

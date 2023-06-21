package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCachePatchSchedulesGet.json
func ExamplePatchSchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPatchSchedulesClient().Get(ctx, "rg1", "cache1", armredis.DefaultNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PatchSchedule = armredis.PatchSchedule{
	// 	Name: to.Ptr("cache1/default"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/PatchSchedules"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/patchSchedules/default"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armredis.ScheduleEntries{
	// 		ScheduleEntries: []*armredis.ScheduleEntry{
	// 			{
	// 				DayOfWeek: to.Ptr(armredis.DayOfWeekMonday),
	// 				MaintenanceWindow: to.Ptr("PT5H"),
	// 				StartHourUTC: to.Ptr[int32](12),
	// 			},
	// 			{
	// 				DayOfWeek: to.Ptr(armredis.DayOfWeekTuesday),
	// 				StartHourUTC: to.Ptr[int32](12),
	// 		}},
	// 	},
	// }
}

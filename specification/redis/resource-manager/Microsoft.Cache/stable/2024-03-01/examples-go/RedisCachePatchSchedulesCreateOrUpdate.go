package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCachePatchSchedulesCreateOrUpdate.json
func ExamplePatchSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPatchSchedulesClient().CreateOrUpdate(ctx, "rg1", "cache1", armredis.DefaultNameDefault, armredis.PatchSchedule{
		Properties: &armredis.ScheduleEntries{
			ScheduleEntries: []*armredis.ScheduleEntry{
				{
					DayOfWeek:         to.Ptr(armredis.DayOfWeekMonday),
					MaintenanceWindow: to.Ptr("PT5H"),
					StartHourUTC:      to.Ptr[int32](12),
				},
				{
					DayOfWeek:    to.Ptr(armredis.DayOfWeekTuesday),
					StartHourUTC: to.Ptr[int32](12),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PatchSchedule = armredis.PatchSchedule{
	// 	Name: to.Ptr("cachename1/default"),
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

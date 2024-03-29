package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redis/resource-manager/Microsoft.Cache/stable/2022-05-01/examples/RedisCachePatchSchedulesCreateOrUpdate.json
func ExamplePatchSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredis.NewPatchSchedulesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "rg1", "cache1", armredis.DefaultNameDefault, armredis.PatchSchedule{
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
	// TODO: use response item
	_ = res
}

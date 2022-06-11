```go
package armstorsimple8000series_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupSchedulesCreateOrUpdate.json
func ExampleBackupSchedulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple8000series.NewBackupSchedulesClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"Device05ForSDKTest",
		"BkUpPolicy01ForSDKTest",
		"schedule2",
		"ResourceGroupForSDKTest",
		"ManagerForSDKTest1",
		armstorsimple8000series.BackupSchedule{
			Kind: to.Ptr("Series8000"),
			Properties: &armstorsimple8000series.BackupScheduleProperties{
				BackupType:     to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
				RetentionCount: to.Ptr[int64](1),
				ScheduleRecurrence: &armstorsimple8000series.ScheduleRecurrence{
					RecurrenceType:  to.Ptr(armstorsimple8000series.RecurrenceTypeWeekly),
					RecurrenceValue: to.Ptr[int32](1),
					WeeklyDaysList: []*armstorsimple8000series.DayOfWeek{
						to.Ptr(armstorsimple8000series.DayOfWeekFriday),
						to.Ptr(armstorsimple8000series.DayOfWeekThursday),
						to.Ptr(armstorsimple8000series.DayOfWeekMonday)},
				},
				ScheduleStatus: to.Ptr(armstorsimple8000series.ScheduleStatusEnabled),
				StartTime:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-24T01:00:00Z"); return t }()),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorsimple8000series%2Farmstorsimple8000series%2Fv1.0.0/sdk/resourcemanager/storsimple8000series/armstorsimple8000series/README.md) on how to add the SDK to your project and authenticate.

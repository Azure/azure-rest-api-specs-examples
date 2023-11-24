package armstorsimple8000series_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupSchedulesCreateOrUpdate.json
func ExampleBackupSchedulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupSchedulesClient().BeginCreateOrUpdate(ctx, "Device05ForSDKTest", "BkUpPolicy01ForSDKTest", "schedule2", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.BackupSchedule{
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
			StartTime:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-24T01:00:00.000Z"); return t }()),
		},
	}, nil)
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
	// res.BackupSchedule = armstorsimple8000series.BackupSchedule{
	// 	Name: to.Ptr("schedule2"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/schedules"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/backupPolicies/BkUpPolicy01ForSDKTest/schedules/schedule2"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.BackupScheduleProperties{
	// 		BackupType: to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
	// 		RetentionCount: to.Ptr[int64](1),
	// 		ScheduleRecurrence: &armstorsimple8000series.ScheduleRecurrence{
	// 			RecurrenceType: to.Ptr(armstorsimple8000series.RecurrenceTypeWeekly),
	// 			RecurrenceValue: to.Ptr[int32](1),
	// 			WeeklyDaysList: []*armstorsimple8000series.DayOfWeek{
	// 				to.Ptr(armstorsimple8000series.DayOfWeekMonday),
	// 				to.Ptr(armstorsimple8000series.DayOfWeekThursday),
	// 				to.Ptr(armstorsimple8000series.DayOfWeekFriday)},
	// 			},
	// 			ScheduleStatus: to.Ptr(armstorsimple8000series.ScheduleStatusEnabled),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-24T01:00:00.000Z"); return t}()),
	// 		},
	// 	}
}

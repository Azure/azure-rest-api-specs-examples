package armbulkactions_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-08-06-preview/ScheduledActions_Update_MaximumSet_Gen.json
func ExampleScheduledActionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScheduledActionsClient().BeginUpdate(ctx, "rgcompute", "myScheduledAction", armbulkactions.ScheduledActionUpdate{
		Properties: &armbulkactions.ScheduledActionUpdateProperties{
			ResourceType: to.Ptr(armbulkactions.ResourceTypeVirtualMachine),
			ActionType:   to.Ptr(armbulkactions.ScheduledActionTypeStart),
			StartTime:    to.Ptr(time.Date(2025, time.April, 17, 0, 23, 58, 149000000, time.UTC)),
			EndTime:      to.Ptr(time.Date(2025, time.April, 17, 0, 23, 58, 149000000, time.UTC)),
			Schedule: &armbulkactions.ScheduledActionsScheduleUpdate{
				ScheduledTime: to.Ptr(time.Date(0, time.January, 1, 19, 0, 0, 0, time.UTC)),
				TimeZone:      to.Ptr("America/Los_Angeles"),
				RequestedWeekDays: []*armbulkactions.WeekDay{
					to.Ptr(armbulkactions.WeekDayMonday),
				},
				RequestedMonths: []*armbulkactions.Month{
					to.Ptr(armbulkactions.MonthJanuary),
				},
				RequestedDaysOfTheMonth: []*int32{
					to.Ptr[int32](15),
				},
				ExecutionParameters: &armbulkactions.ScheduledActionsExecutionParameters{
					OptimizationPreference: to.Ptr(armbulkactions.OptimizationPreferenceCost),
					RetryPolicy: &armbulkactions.ScheduledActionsRetryPolicy{
						RetryCount:           to.Ptr[int32](17),
						RetryWindowInMinutes: to.Ptr[int32](29),
					},
				},
				DeadlineType: to.Ptr(armbulkactions.ScheduledActionsDeadlineTypeUnknown),
			},
			NotificationSettings: []*armbulkactions.NotificationProperties{
				{
					Destination: to.Ptr("admin@contoso.com"),
					Type:        to.Ptr(armbulkactions.NotificationTypeEmail),
					Language:    to.Ptr(armbulkactions.LanguageEnUs),
					Disabled:    to.Ptr(true),
				},
			},
			Disabled: to.Ptr(true),
		},
		Tags: map[string]*string{
			"key9989": to.Ptr("myTagValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.ScheduledActionsClientUpdateResponse{
	// }
}

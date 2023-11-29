package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateSchedule.json
func ExampleScheduleClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduleClient().Update(ctx, "rg", "myAutomationAccount33", "mySchedule", armautomation.ScheduleUpdateParameters{
		Name: to.Ptr("mySchedule"),
		Properties: &armautomation.ScheduleUpdateProperties{
			Description: to.Ptr("my updated description of schedule goes here"),
			IsEnabled:   to.Ptr(false),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Schedule = armautomation.Schedule{
	// 	Name: to.Ptr("mySchedule"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/schedules/mySchedule"),
	// 	Properties: &armautomation.ScheduleProperties{
	// 		Description: to.Ptr("my updated description of schedule goes here"),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T16:59:22.697Z"); return t}()),
	// 		ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-01T17:28:00.000Z"); return t}()),
	// 		ExpiryTimeOffsetMinutes: to.Ptr[float64](0),
	// 		Frequency: to.Ptr(armautomation.ScheduleFrequencyHour),
	// 		Interval: float64(1),
	// 		IsEnabled: to.Ptr(false),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T16:59:22.697Z"); return t}()),
	// 		NextRun: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T17:28:00.000Z"); return t}()),
	// 		NextRunOffsetMinutes: to.Ptr[float64](0),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T17:28:00.000Z"); return t}()),
	// 		StartTimeOffsetMinutes: to.Ptr[float64](0),
	// 		TimeZone: to.Ptr("UTC"),
	// 	},
	// }
}

package armautomation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateSchedule.json
func ExampleScheduleClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armautomation.NewScheduleClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg",
		"myAutomationAccount33",
		"mySchedule",
		armautomation.ScheduleCreateOrUpdateParameters{
			Name: to.Ptr("mySchedule"),
			Properties: &armautomation.ScheduleCreateOrUpdateProperties{
				Description:      to.Ptr("my description of schedule goes here"),
				AdvancedSchedule: &armautomation.AdvancedSchedule{},
				ExpiryTime:       to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-01T17:28:57.2494819Z"); return t }()),
				Frequency:        to.Ptr(armautomation.ScheduleFrequencyHour),
				Interval:         float64(1),
				StartTime:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T17:28:57.2494819Z"); return t }()),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

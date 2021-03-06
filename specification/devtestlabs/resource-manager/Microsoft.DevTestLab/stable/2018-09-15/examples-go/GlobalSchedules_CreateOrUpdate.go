package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_CreateOrUpdate.json
func ExampleGlobalSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewGlobalSchedulesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroupName",
		"labvmautostart",
		armdevtestlabs.Schedule{
			Properties: &armdevtestlabs.ScheduleProperties{
				Status:     to.Ptr(armdevtestlabs.EnableStatusEnabled),
				TaskType:   to.Ptr("LabVmsStartupTask"),
				TimeZoneID: to.Ptr("Hawaiian Standard Time"),
				WeeklyRecurrence: &armdevtestlabs.WeekDetails{
					Time: to.Ptr("0700"),
					Weekdays: []*string{
						to.Ptr("Monday"),
						to.Ptr("Tuesday"),
						to.Ptr("Wednesday"),
						to.Ptr("Thursday"),
						to.Ptr("Friday"),
						to.Ptr("Saturday")},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

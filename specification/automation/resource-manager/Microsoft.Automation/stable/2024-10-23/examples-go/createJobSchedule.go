package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/createJobSchedule.json
func ExampleJobScheduleClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewJobScheduleClient().Create(ctx, "rg", "ContoseAutomationAccount", "0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc", armautomation.JobScheduleCreateParameters{
		Properties: &armautomation.JobScheduleCreateProperties{
			Parameters: map[string]*string{
				"jobscheduletag01": to.Ptr("jobschedulevalue01"),
				"jobscheduletag02": to.Ptr("jobschedulevalue02"),
			},
			Runbook: &armautomation.RunbookAssociationProperty{
				Name: to.Ptr("TestRunbook"),
			},
			Schedule: &armautomation.ScheduleAssociationProperty{
				Name: to.Ptr("ScheduleNameGoesHere332204b5-debe-4348-a5c7-6357457189f2"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getJobSchedule.json
func ExampleJobScheduleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobScheduleClient().Get(ctx, "rg", "ContoseAutomationAccount", "0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobSchedule = armautomation.JobSchedule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobSchedules/0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc"),
	// 	Properties: &armautomation.JobScheduleProperties{
	// 		JobScheduleID: to.Ptr("0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc"),
	// 		Parameters: map[string]*string{
	// 			"jobscheduletag01": to.Ptr("jobschedulevalue01"),
	// 			"jobscheduletag02": to.Ptr("jobschedulevalue02"),
	// 		},
	// 		Runbook: &armautomation.RunbookAssociationProperty{
	// 			Name: to.Ptr("TestRunbook"),
	// 		},
	// 		Schedule: &armautomation.ScheduleAssociationProperty{
	// 			Name: to.Ptr("ScheduleNameGoesHere332204b5-debe-4348-a5c7-6357457189f2"),
	// 		},
	// 	},
	// }
}

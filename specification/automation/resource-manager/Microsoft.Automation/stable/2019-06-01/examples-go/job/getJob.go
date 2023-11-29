package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/getJob.json
func ExampleJobClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobClient().Get(ctx, "mygroup", "ContoseAutomationAccount", "foo", &armautomation.JobClientGetOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armautomation.Job{
	// 	Name: to.Ptr("foo"),
	// 	Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Jobs"),
	// 	ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobs/jobName"),
	// 	Properties: &armautomation.JobProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-01T05:53:30.243Z"); return t}()),
	// 		JobID: to.Ptr("5b8a3960-e8ab-45f6-bec6-567df8467d1a"),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-01T05:53:30.243Z"); return t}()),
	// 		LastStatusModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-01T05:53:30.243Z"); return t}()),
	// 		Parameters: map[string]*string{
	// 			"tag01": to.Ptr("value01"),
	// 			"tag02": to.Ptr("value02"),
	// 		},
	// 		ProvisioningState: to.Ptr(armautomation.JobProvisioningStateSucceeded),
	// 		RunOn: to.Ptr(""),
	// 		Runbook: &armautomation.RunbookAssociationProperty{
	// 			Name: to.Ptr("TestRunbook"),
	// 		},
	// 		Status: to.Ptr(armautomation.JobStatusNew),
	// 		StatusDetails: to.Ptr("None"),
	// 	},
	// }
}

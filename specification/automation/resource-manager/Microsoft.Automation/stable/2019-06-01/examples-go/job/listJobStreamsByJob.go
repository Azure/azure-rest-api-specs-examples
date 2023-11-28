package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/listJobStreamsByJob.json
func ExampleJobStreamClient_NewListByJobPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobStreamClient().NewListByJobPager("mygroup", "ContoseAutomationAccount", "foo", &armautomation.JobStreamClientListByJobOptions{Filter: nil,
		ClientRequestID: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.JobStreamListResult = armautomation.JobStreamListResult{
		// 	Value: []*armautomation.JobStream{
		// 		{
		// 			ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourcegroups/mygroup/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobs/jobName/streams/24456a8a-2857-4af6-932c-3455f38bd05e_00636535675981232703_00000000000000000001"),
		// 			Properties: &armautomation.JobStreamProperties{
		// 				JobStreamID: to.Ptr("24456a8a-2857-4af6-932c-3455f38bd05e_00636535675981232703_00000000000000000001"),
		// 				StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
		// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-07T02:33:18.123Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourcegroups/mygroup/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobs/jobName/streams/24456a8a-2857-4af6-932c-3455f38bd05e_00636535675984691350_00000000000000000002"),
		// 			Properties: &armautomation.JobStreamProperties{
		// 				JobStreamID: to.Ptr("24456a8a-2857-4af6-932c-3455f38bd05e_00636535675984691350_00000000000000000002"),
		// 				StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
		// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-07T02:33:18.469Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

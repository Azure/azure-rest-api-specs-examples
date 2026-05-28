package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/ListJobExecutionSteps.json
func ExampleJobStepExecutionsClient_NewListByJobExecutionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobStepExecutionsClient().NewListByJobExecutionPager("group1", "server1", "agent1", "job1", "5A86BF65-43AC-F258-2524-9E92992F97CA", nil)
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
		// page = armsql.JobStepExecutionsClientListByJobExecutionResponse{
		// 	JobExecutionListResult: armsql.JobExecutionListResult{
		// 		Value: []*armsql.JobExecution{
		// 			{
		// 				Properties: &armsql.JobExecutionProperties{
		// 					JobVersion: to.Ptr[int32](1),
		// 					StepName: to.Ptr("step1"),
		// 					StepID: to.Ptr[int32](1),
		// 					JobExecutionID: to.Ptr("5A86BF65-43AC-F258-2524-9E92992F97CA"),
		// 					Lifecycle: to.Ptr(armsql.JobExecutionLifecycleSucceeded),
		// 					ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
		// 					CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-03T04:33:15.7189151Z"); return t}()),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-03T04:33:16.176937Z"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-03T04:33:19.0600862Z"); return t}()),
		// 					CurrentAttempts: to.Ptr[int32](1),
		// 					CurrentAttemptStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-03T04:33:17.4840068Z"); return t}()),
		// 					LastMessage: to.Ptr("Step 1 succeeded."),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1/executions/5555-6666-7777-8888-999999999999/steps/step1"),
		// 				Name: to.Ptr("step1"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/jobAgents/jobs/executions/steps"),
		// 			},
		// 		},
		// 	},
		// }
	}
}

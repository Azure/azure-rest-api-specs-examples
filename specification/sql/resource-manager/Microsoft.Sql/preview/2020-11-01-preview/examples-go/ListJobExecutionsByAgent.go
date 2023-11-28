package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobExecutionsByAgent.json
func ExampleJobExecutionsClient_NewListByAgentPager_listAllJobExecutionsInAJobAgent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobExecutionsClient().NewListByAgentPager("group1", "server1", "agent1", &armsql.JobExecutionsClientListByAgentOptions{CreateTimeMin: nil,
		CreateTimeMax: nil,
		EndTimeMin:    nil,
		EndTimeMax:    nil,
		IsActive:      nil,
		Skip:          nil,
		Top:           nil,
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
		// page.JobExecutionListResult = armsql.JobExecutionListResult{
		// 	Value: []*armsql.JobExecution{
		// 		{
		// 			Name: to.Ptr("5555-6666-7777-8888-999999999999"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAgents/jobs/executions"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1/executions/5555-6666-7777-8888-999999999999"),
		// 			Properties: &armsql.JobExecutionProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:02:00.870Z"); return t}()),
		// 				CurrentAttemptStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:12:00.870Z"); return t}()),
		// 				CurrentAttempts: to.Ptr[int32](0),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:22:00.870Z"); return t}()),
		// 				JobExecutionID: to.Ptr("5A86BF65-43AC-F258-2524-9E92992F97CA"),
		// 				JobVersion: to.Ptr[int32](1),
		// 				LastMessage: to.Ptr("Job execution created."),
		// 				Lifecycle: to.Ptr(armsql.JobExecutionLifecycleCreated),
		// 				ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:12:00.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAgents/jobs/executions"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job2/executions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 			Properties: &armsql.JobExecutionProperties{
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-15T19:02:00.870Z"); return t}()),
		// 				CurrentAttempts: to.Ptr[int32](1),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-15T19:22:00.870Z"); return t}()),
		// 				JobExecutionID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 				JobVersion: to.Ptr[int32](1),
		// 				LastMessage: to.Ptr("Job execution succeeded."),
		// 				Lifecycle: to.Ptr(armsql.JobExecutionLifecycleSucceeded),
		// 				ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-15T19:12:00.870Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

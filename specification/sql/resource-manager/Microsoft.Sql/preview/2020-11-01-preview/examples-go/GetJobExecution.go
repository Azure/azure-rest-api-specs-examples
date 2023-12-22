package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetJobExecution.json
func ExampleJobExecutionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobExecutionsClient().Get(ctx, "group1", "server1", "agent1", "job1", "5A86BF65-43AC-F258-2524-9E92992F97CA", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobExecution = armsql.JobExecution{
	// 	Name: to.Ptr("5555-6666-7777-8888-999999999999"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/jobAgents/executions"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1/executions/5555-6666-7777-8888-999999999999"),
	// 	Properties: &armsql.JobExecutionProperties{
	// 		CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-21T19:02:00.870Z"); return t}()),
	// 		CurrentAttemptStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-13T19:12:00.870Z"); return t}()),
	// 		CurrentAttempts: to.Ptr[int32](0),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-21T19:22:00.870Z"); return t}()),
	// 		JobExecutionID: to.Ptr("5A86BF65-43AC-F258-2524-9E92992F97CA"),
	// 		JobVersion: to.Ptr[int32](1),
	// 		LastMessage: to.Ptr("Job execution created."),
	// 		Lifecycle: to.Ptr(armsql.JobExecutionLifecycleCreated),
	// 		ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-21T19:12:00.870Z"); return t}()),
	// 	},
	// }
}

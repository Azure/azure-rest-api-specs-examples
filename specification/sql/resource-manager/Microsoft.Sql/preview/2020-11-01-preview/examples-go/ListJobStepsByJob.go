package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobStepsByJob.json
func ExampleJobStepsClient_NewListByJobPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobStepsClient().NewListByJobPager("group1", "server1", "agent1", "job1", nil)
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
		// page.JobStepListResult = armsql.JobStepListResult{
		// 	Value: []*armsql.JobStep{
		// 		{
		// 			Name: to.Ptr("step1"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAgents/jobs/steps"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1/steps/step1"),
		// 			Properties: &armsql.JobStepProperties{
		// 				Action: &armsql.JobStepAction{
		// 					Type: to.Ptr(armsql.JobStepActionTypeTSQL),
		// 					Source: to.Ptr(armsql.JobStepActionSourceInline),
		// 					Value: to.Ptr("select 2"),
		// 				},
		// 				Credential: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred1"),
		// 				ExecutionOptions: &armsql.JobStepExecutionOptions{
		// 					InitialRetryIntervalSeconds: to.Ptr[int32](11),
		// 					MaximumRetryIntervalSeconds: to.Ptr[int32](222),
		// 					RetryAttempts: to.Ptr[int32](42),
		// 					RetryIntervalBackoffMultiplier: to.Ptr[float32](3),
		// 					TimeoutSeconds: to.Ptr[int32](1234),
		// 				},
		// 				Output: &armsql.JobStepOutput{
		// 					Type: to.Ptr(armsql.JobStepOutputTypeSQLDatabase),
		// 					Credential: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0"),
		// 					DatabaseName: to.Ptr("database3"),
		// 					ResourceGroupName: to.Ptr("group3"),
		// 					SchemaName: to.Ptr("myschema1234"),
		// 					ServerName: to.Ptr("server3"),
		// 					SubscriptionID: to.Ptr("3501b905-a848-4b5d-96e8-b253f62d735a"),
		// 					TableName: to.Ptr("mytable5678"),
		// 				},
		// 				StepID: to.Ptr[int32](1),
		// 				TargetGroup: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("step2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAgents/jobs/steps"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1/steps/step2"),
		// 			Properties: &armsql.JobStepProperties{
		// 				Action: &armsql.JobStepAction{
		// 					Type: to.Ptr(armsql.JobStepActionTypeTSQL),
		// 					Source: to.Ptr(armsql.JobStepActionSourceInline),
		// 					Value: to.Ptr("select 2"),
		// 				},
		// 				Credential: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred1"),
		// 				ExecutionOptions: &armsql.JobStepExecutionOptions{
		// 					InitialRetryIntervalSeconds: to.Ptr[int32](11),
		// 					MaximumRetryIntervalSeconds: to.Ptr[int32](222),
		// 					RetryAttempts: to.Ptr[int32](42),
		// 					RetryIntervalBackoffMultiplier: to.Ptr[float32](3),
		// 					TimeoutSeconds: to.Ptr[int32](1234),
		// 				},
		// 				Output: &armsql.JobStepOutput{
		// 					Type: to.Ptr(armsql.JobStepOutputTypeSQLDatabase),
		// 					Credential: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0"),
		// 					DatabaseName: to.Ptr("database3"),
		// 					ResourceGroupName: to.Ptr("group3"),
		// 					SchemaName: to.Ptr("myschema1234"),
		// 					ServerName: to.Ptr("server3"),
		// 					SubscriptionID: to.Ptr("3501b905-a848-4b5d-96e8-b253f62d735a"),
		// 					TableName: to.Ptr("mytable5678"),
		// 				},
		// 				StepID: to.Ptr[int32](2),
		// 				TargetGroup: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup1"),
		// 			},
		// 	}},
		// }
	}
}

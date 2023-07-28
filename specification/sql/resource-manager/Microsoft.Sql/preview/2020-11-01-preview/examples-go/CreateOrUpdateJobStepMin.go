package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobStepMin.json
func ExampleJobStepsClient_CreateOrUpdate_createOrUpdateAJobStepWithMinimalPropertiesSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobStepsClient().CreateOrUpdate(ctx, "group1", "server1", "agent1", "job1", "step1", armsql.JobStep{
		Properties: &armsql.JobStepProperties{
			Action: &armsql.JobStepAction{
				Value: to.Ptr("select 1"),
			},
			Credential:  to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0"),
			TargetGroup: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobStep = armsql.JobStep{
	// 	Name: to.Ptr("step1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/jobAgents/jobs/steps"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1/steps/step1"),
	// 	Properties: &armsql.JobStepProperties{
	// 		Action: &armsql.JobStepAction{
	// 			Type: to.Ptr(armsql.JobStepActionTypeTSQL),
	// 			Source: to.Ptr(armsql.JobStepActionSourceInline),
	// 			Value: to.Ptr("select 1"),
	// 		},
	// 		Credential: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0"),
	// 		ExecutionOptions: &armsql.JobStepExecutionOptions{
	// 			InitialRetryIntervalSeconds: to.Ptr[int32](1),
	// 			MaximumRetryIntervalSeconds: to.Ptr[int32](120),
	// 			RetryAttempts: to.Ptr[int32](10),
	// 			RetryIntervalBackoffMultiplier: to.Ptr[float32](2),
	// 			TimeoutSeconds: to.Ptr[int32](43200),
	// 		},
	// 		StepID: to.Ptr[int32](1),
	// 		TargetGroup: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup0"),
	// 	},
	// }
}

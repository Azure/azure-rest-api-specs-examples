```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobStepMax.json
func ExampleJobStepsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewJobStepsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"group1",
		"server1",
		"agent1",
		"job1",
		"step1",
		armsql.JobStep{
			Properties: &armsql.JobStepProperties{
				Action: &armsql.JobStepAction{
					Type:   to.Ptr(armsql.JobStepActionTypeTSQL),
					Source: to.Ptr(armsql.JobStepActionSourceInline),
					Value:  to.Ptr("select 2"),
				},
				Credential: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred1"),
				ExecutionOptions: &armsql.JobStepExecutionOptions{
					InitialRetryIntervalSeconds:    to.Ptr[int32](11),
					MaximumRetryIntervalSeconds:    to.Ptr[int32](222),
					RetryAttempts:                  to.Ptr[int32](42),
					RetryIntervalBackoffMultiplier: to.Ptr[float32](3),
					TimeoutSeconds:                 to.Ptr[int32](1234),
				},
				Output: &armsql.JobStepOutput{
					Type:              to.Ptr(armsql.JobStepOutputTypeSQLDatabase),
					Credential:        to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0"),
					DatabaseName:      to.Ptr("database3"),
					ResourceGroupName: to.Ptr("group3"),
					SchemaName:        to.Ptr("myschema1234"),
					ServerName:        to.Ptr("server3"),
					SubscriptionID:    to.Ptr("3501b905-a848-4b5d-96e8-b253f62d735a"),
					TableName:         to.Ptr("mytable5678"),
				},
				StepID:      to.Ptr[int32](1),
				TargetGroup: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup1"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv1.0.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

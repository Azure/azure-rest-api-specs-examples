Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.5.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armsql.NewJobStepsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<job-agent-name>",
		"<job-name>",
		"<step-name>",
		armsql.JobStep{
			Properties: &armsql.JobStepProperties{
				Action: &armsql.JobStepAction{
					Type:   to.Ptr(armsql.JobStepActionTypeTSQL),
					Source: to.Ptr(armsql.JobStepActionSourceInline),
					Value:  to.Ptr("<value>"),
				},
				Credential: to.Ptr("<credential>"),
				ExecutionOptions: &armsql.JobStepExecutionOptions{
					InitialRetryIntervalSeconds:    to.Ptr[int32](11),
					MaximumRetryIntervalSeconds:    to.Ptr[int32](222),
					RetryAttempts:                  to.Ptr[int32](42),
					RetryIntervalBackoffMultiplier: to.Ptr[float32](3),
					TimeoutSeconds:                 to.Ptr[int32](1234),
				},
				Output: &armsql.JobStepOutput{
					Type:              to.Ptr(armsql.JobStepOutputTypeSQLDatabase),
					Credential:        to.Ptr("<credential>"),
					DatabaseName:      to.Ptr("<database-name>"),
					ResourceGroupName: to.Ptr("<resource-group-name>"),
					SchemaName:        to.Ptr("<schema-name>"),
					ServerName:        to.Ptr("<server-name>"),
					SubscriptionID:    to.Ptr("<subscription-id>"),
					TableName:         to.Ptr("<table-name>"),
				},
				StepID:      to.Ptr[int32](1),
				TargetGroup: to.Ptr("<target-group>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

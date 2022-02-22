Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobStepMax.json
func ExampleJobStepsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewJobStepsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<job-agent-name>",
		"<job-name>",
		"<step-name>",
		armsql.JobStep{
			Properties: &armsql.JobStepProperties{
				Action: &armsql.JobStepAction{
					Type:   armsql.JobStepActionType("TSql").ToPtr(),
					Source: armsql.JobStepActionSource("Inline").ToPtr(),
					Value:  to.StringPtr("<value>"),
				},
				Credential: to.StringPtr("<credential>"),
				ExecutionOptions: &armsql.JobStepExecutionOptions{
					InitialRetryIntervalSeconds:    to.Int32Ptr(11),
					MaximumRetryIntervalSeconds:    to.Int32Ptr(222),
					RetryAttempts:                  to.Int32Ptr(42),
					RetryIntervalBackoffMultiplier: to.Float32Ptr(3),
					TimeoutSeconds:                 to.Int32Ptr(1234),
				},
				Output: &armsql.JobStepOutput{
					Type:              armsql.JobStepOutputType("SqlDatabase").ToPtr(),
					Credential:        to.StringPtr("<credential>"),
					DatabaseName:      to.StringPtr("<database-name>"),
					ResourceGroupName: to.StringPtr("<resource-group-name>"),
					SchemaName:        to.StringPtr("<schema-name>"),
					ServerName:        to.StringPtr("<server-name>"),
					SubscriptionID:    to.StringPtr("<subscription-id>"),
					TableName:         to.StringPtr("<table-name>"),
				},
				StepID:      to.Int32Ptr(1),
				TargetGroup: to.StringPtr("<target-group>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JobStepsClientCreateOrUpdateResult)
}
```

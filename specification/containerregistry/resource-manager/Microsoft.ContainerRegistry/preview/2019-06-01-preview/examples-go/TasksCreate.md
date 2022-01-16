Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerregistry%2Farmcontainerregistry%2Fv0.3.0/sdk/resourcemanager/containerregistry/armcontainerregistry/README.md) on how to add the SDK to your project and authenticate.

```go
package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TasksCreate.json
func ExampleTasksClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerregistry.NewTasksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<task-name>",
		armcontainerregistry.Task{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"testkey": to.StringPtr("value"),
			},
			Identity: &armcontainerregistry.IdentityProperties{
				Type: armcontainerregistry.ResourceIdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &armcontainerregistry.TaskProperties{
				AgentConfiguration: &armcontainerregistry.AgentProperties{
					CPU: to.Int32Ptr(2),
				},
				IsSystemTask: to.BoolPtr(false),
				LogTemplate:  to.StringPtr("<log-template>"),
				Platform: &armcontainerregistry.PlatformProperties{
					Architecture: armcontainerregistry.Architecture("amd64").ToPtr(),
					OS:           armcontainerregistry.OS("Linux").ToPtr(),
				},
				Status: armcontainerregistry.TaskStatus("Enabled").ToPtr(),
				Step: &armcontainerregistry.DockerBuildStep{
					Type:        armcontainerregistry.StepType("Docker").ToPtr(),
					ContextPath: to.StringPtr("<context-path>"),
					Arguments: []*armcontainerregistry.Argument{
						{
							Name:     to.StringPtr("<name>"),
							IsSecret: to.BoolPtr(false),
							Value:    to.StringPtr("<value>"),
						},
						{
							Name:     to.StringPtr("<name>"),
							IsSecret: to.BoolPtr(true),
							Value:    to.StringPtr("<value>"),
						}},
					DockerFilePath: to.StringPtr("<docker-file-path>"),
					ImageNames: []*string{
						to.StringPtr("azurerest:testtag")},
					IsPushEnabled: to.BoolPtr(true),
					NoCache:       to.BoolPtr(false),
				},
				Trigger: &armcontainerregistry.TriggerProperties{
					BaseImageTrigger: &armcontainerregistry.BaseImageTrigger{
						Name:                     to.StringPtr("<name>"),
						BaseImageTriggerType:     armcontainerregistry.BaseImageTriggerType("Runtime").ToPtr(),
						UpdateTriggerEndpoint:    to.StringPtr("<update-trigger-endpoint>"),
						UpdateTriggerPayloadType: armcontainerregistry.UpdateTriggerPayloadType("Token").ToPtr(),
					},
					SourceTriggers: []*armcontainerregistry.SourceTrigger{
						{
							Name: to.StringPtr("<name>"),
							SourceRepository: &armcontainerregistry.SourceProperties{
								Branch:        to.StringPtr("<branch>"),
								RepositoryURL: to.StringPtr("<repository-url>"),
								SourceControlAuthProperties: &armcontainerregistry.AuthInfo{
									Token:     to.StringPtr("<token>"),
									TokenType: armcontainerregistry.TokenType("PAT").ToPtr(),
								},
								SourceControlType: armcontainerregistry.SourceControlType("Github").ToPtr(),
							},
							SourceTriggerEvents: []*armcontainerregistry.SourceTriggerEvent{
								armcontainerregistry.SourceTriggerEvent("commit").ToPtr()},
						}},
					TimerTriggers: []*armcontainerregistry.TimerTrigger{
						{
							Name:     to.StringPtr("<name>"),
							Schedule: to.StringPtr("<schedule>"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TasksClientCreateResult)
}
```

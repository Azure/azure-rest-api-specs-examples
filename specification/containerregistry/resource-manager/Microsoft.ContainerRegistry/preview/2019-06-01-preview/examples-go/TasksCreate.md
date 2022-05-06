Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerregistry%2Farmcontainerregistry%2Fv0.5.0/sdk/resourcemanager/containerregistry/armcontainerregistry/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TasksCreate.json
func ExampleTasksClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewTasksClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<task-name>",
		armcontainerregistry.Task{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"testkey": to.Ptr("value"),
			},
			Identity: &armcontainerregistry.IdentityProperties{
				Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssigned),
			},
			Properties: &armcontainerregistry.TaskProperties{
				AgentConfiguration: &armcontainerregistry.AgentProperties{
					CPU: to.Ptr[int32](2),
				},
				IsSystemTask: to.Ptr(false),
				LogTemplate:  to.Ptr("<log-template>"),
				Platform: &armcontainerregistry.PlatformProperties{
					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
					OS:           to.Ptr(armcontainerregistry.OSLinux),
				},
				Status: to.Ptr(armcontainerregistry.TaskStatusEnabled),
				Step: &armcontainerregistry.DockerBuildStep{
					Type:        to.Ptr(armcontainerregistry.StepTypeDocker),
					ContextPath: to.Ptr("<context-path>"),
					Arguments: []*armcontainerregistry.Argument{
						{
							Name:     to.Ptr("<name>"),
							IsSecret: to.Ptr(false),
							Value:    to.Ptr("<value>"),
						},
						{
							Name:     to.Ptr("<name>"),
							IsSecret: to.Ptr(true),
							Value:    to.Ptr("<value>"),
						}},
					DockerFilePath: to.Ptr("<docker-file-path>"),
					ImageNames: []*string{
						to.Ptr("azurerest:testtag")},
					IsPushEnabled: to.Ptr(true),
					NoCache:       to.Ptr(false),
				},
				Trigger: &armcontainerregistry.TriggerProperties{
					BaseImageTrigger: &armcontainerregistry.BaseImageTrigger{
						Name:                     to.Ptr("<name>"),
						BaseImageTriggerType:     to.Ptr(armcontainerregistry.BaseImageTriggerTypeRuntime),
						UpdateTriggerEndpoint:    to.Ptr("<update-trigger-endpoint>"),
						UpdateTriggerPayloadType: to.Ptr(armcontainerregistry.UpdateTriggerPayloadTypeToken),
					},
					SourceTriggers: []*armcontainerregistry.SourceTrigger{
						{
							Name: to.Ptr("<name>"),
							SourceRepository: &armcontainerregistry.SourceProperties{
								Branch:        to.Ptr("<branch>"),
								RepositoryURL: to.Ptr("<repository-url>"),
								SourceControlAuthProperties: &armcontainerregistry.AuthInfo{
									Token:     to.Ptr("<token>"),
									TokenType: to.Ptr(armcontainerregistry.TokenTypePAT),
								},
								SourceControlType: to.Ptr(armcontainerregistry.SourceControlTypeGithub),
							},
							SourceTriggerEvents: []*armcontainerregistry.SourceTriggerEvent{
								to.Ptr(armcontainerregistry.SourceTriggerEventCommit)},
						}},
					TimerTriggers: []*armcontainerregistry.TimerTrigger{
						{
							Name:     to.Ptr("<name>"),
							Schedule: to.Ptr("<schedule>"),
						}},
				},
			},
		},
		&armcontainerregistry.TasksClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

```go
package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TasksCreate.json
func ExampleTasksClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewTasksClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"myResourceGroup",
		"myRegistry",
		"mytTask",
		armcontainerregistry.Task{
			Location: to.Ptr("eastus"),
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
				LogTemplate:  to.Ptr("acr/tasks:{{.Run.OS}}"),
				Platform: &armcontainerregistry.PlatformProperties{
					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
					OS:           to.Ptr(armcontainerregistry.OSLinux),
				},
				Status: to.Ptr(armcontainerregistry.TaskStatusEnabled),
				Step: &armcontainerregistry.DockerBuildStep{
					Type:        to.Ptr(armcontainerregistry.StepTypeDocker),
					ContextPath: to.Ptr("src"),
					Arguments: []*armcontainerregistry.Argument{
						{
							Name:     to.Ptr("mytestargument"),
							IsSecret: to.Ptr(false),
							Value:    to.Ptr("mytestvalue"),
						},
						{
							Name:     to.Ptr("mysecrettestargument"),
							IsSecret: to.Ptr(true),
							Value:    to.Ptr("mysecrettestvalue"),
						}},
					DockerFilePath: to.Ptr("src/DockerFile"),
					ImageNames: []*string{
						to.Ptr("azurerest:testtag")},
					IsPushEnabled: to.Ptr(true),
					NoCache:       to.Ptr(false),
				},
				Trigger: &armcontainerregistry.TriggerProperties{
					BaseImageTrigger: &armcontainerregistry.BaseImageTrigger{
						Name:                     to.Ptr("myBaseImageTrigger"),
						BaseImageTriggerType:     to.Ptr(armcontainerregistry.BaseImageTriggerTypeRuntime),
						UpdateTriggerEndpoint:    to.Ptr("https://user:pass@mycicd.webhook.com?token=foo"),
						UpdateTriggerPayloadType: to.Ptr(armcontainerregistry.UpdateTriggerPayloadTypeToken),
					},
					SourceTriggers: []*armcontainerregistry.SourceTrigger{
						{
							Name: to.Ptr("mySourceTrigger"),
							SourceRepository: &armcontainerregistry.SourceProperties{
								Branch:        to.Ptr("master"),
								RepositoryURL: to.Ptr("https://github.com/Azure/azure-rest-api-specs"),
								SourceControlAuthProperties: &armcontainerregistry.AuthInfo{
									Token:     to.Ptr("xxxxx"),
									TokenType: to.Ptr(armcontainerregistry.TokenTypePAT),
								},
								SourceControlType: to.Ptr(armcontainerregistry.SourceControlTypeGithub),
							},
							SourceTriggerEvents: []*armcontainerregistry.SourceTriggerEvent{
								to.Ptr(armcontainerregistry.SourceTriggerEventCommit)},
						}},
					TimerTriggers: []*armcontainerregistry.TimerTrigger{
						{
							Name:     to.Ptr("myTimerTrigger"),
							Schedule: to.Ptr("30 9 * * 1-5"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerregistry%2Farmcontainerregistry%2Fv0.6.0/sdk/resourcemanager/containerregistry/armcontainerregistry/README.md) on how to add the SDK to your project and authenticate.

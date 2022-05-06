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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TasksUpdate.json
func ExampleTasksClient_BeginUpdate() {
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
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<task-name>",
		armcontainerregistry.TaskUpdateParameters{
			Properties: &armcontainerregistry.TaskPropertiesUpdateParameters{
				AgentConfiguration: &armcontainerregistry.AgentProperties{
					CPU: to.Ptr[int32](3),
				},
				Credentials: &armcontainerregistry.Credentials{
					CustomRegistries: map[string]*armcontainerregistry.CustomRegistryCredentials{
						"myregistry.azurecr.io": {
							Identity: to.Ptr("<identity>"),
							Password: &armcontainerregistry.SecretObject{
								Type:  to.Ptr(armcontainerregistry.SecretObjectTypeVaultsecret),
								Value: to.Ptr("<value>"),
							},
							UserName: &armcontainerregistry.SecretObject{
								Type:  to.Ptr(armcontainerregistry.SecretObjectTypeOpaque),
								Value: to.Ptr("<value>"),
							},
						},
					},
				},
				LogTemplate: to.Ptr("<log-template>"),
				Status:      to.Ptr(armcontainerregistry.TaskStatusEnabled),
				Step: &armcontainerregistry.DockerBuildStepUpdateParameters{
					Type:           to.Ptr(armcontainerregistry.StepTypeDocker),
					DockerFilePath: to.Ptr("<docker-file-path>"),
					ImageNames: []*string{
						to.Ptr("azurerest:testtag1")},
				},
				Trigger: &armcontainerregistry.TriggerUpdateParameters{
					SourceTriggers: []*armcontainerregistry.SourceTriggerUpdateParameters{
						{
							Name: to.Ptr("<name>"),
							SourceRepository: &armcontainerregistry.SourceUpdateParameters{
								SourceControlAuthProperties: &armcontainerregistry.AuthInfoUpdateParameters{
									Token:     to.Ptr("<token>"),
									TokenType: to.Ptr(armcontainerregistry.TokenTypePAT),
								},
							},
							SourceTriggerEvents: []*armcontainerregistry.SourceTriggerEvent{
								to.Ptr(armcontainerregistry.SourceTriggerEventCommit)},
						}},
				},
			},
			Tags: map[string]*string{
				"testkey": to.Ptr("value"),
			},
		},
		&armcontainerregistry.TasksClientBeginUpdateOptions{ResumeToken: ""})
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

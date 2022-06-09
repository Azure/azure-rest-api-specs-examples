```go
package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TasksUpdate.json
func ExampleTasksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewTasksClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myRegistry",
		"myTask",
		armcontainerregistry.TaskUpdateParameters{
			Properties: &armcontainerregistry.TaskPropertiesUpdateParameters{
				AgentConfiguration: &armcontainerregistry.AgentProperties{
					CPU: to.Ptr[int32](3),
				},
				Credentials: &armcontainerregistry.Credentials{
					CustomRegistries: map[string]*armcontainerregistry.CustomRegistryCredentials{
						"myregistry.azurecr.io": {
							Identity: to.Ptr("[system]"),
							Password: &armcontainerregistry.SecretObject{
								Type:  to.Ptr(armcontainerregistry.SecretObjectTypeVaultsecret),
								Value: to.Ptr("https://myacbvault.vault.azure.net/secrets/password"),
							},
							UserName: &armcontainerregistry.SecretObject{
								Type:  to.Ptr(armcontainerregistry.SecretObjectTypeOpaque),
								Value: to.Ptr("username"),
							},
						},
					},
				},
				LogTemplate: to.Ptr("acr/tasks:{{.Run.OS}}"),
				Status:      to.Ptr(armcontainerregistry.TaskStatusEnabled),
				Step: &armcontainerregistry.DockerBuildStepUpdateParameters{
					Type:           to.Ptr(armcontainerregistry.StepTypeDocker),
					DockerFilePath: to.Ptr("src/DockerFile"),
					ImageNames: []*string{
						to.Ptr("azurerest:testtag1")},
				},
				Trigger: &armcontainerregistry.TriggerUpdateParameters{
					SourceTriggers: []*armcontainerregistry.SourceTriggerUpdateParameters{
						{
							Name: to.Ptr("mySourceTrigger"),
							SourceRepository: &armcontainerregistry.SourceUpdateParameters{
								SourceControlAuthProperties: &armcontainerregistry.AuthInfoUpdateParameters{
									Token:     to.Ptr("xxxxx"),
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

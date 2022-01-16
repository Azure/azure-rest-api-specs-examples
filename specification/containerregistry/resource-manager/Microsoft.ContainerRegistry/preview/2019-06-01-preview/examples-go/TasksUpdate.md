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

// x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TasksUpdate.json
func ExampleTasksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerregistry.NewTasksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<task-name>",
		armcontainerregistry.TaskUpdateParameters{
			Properties: &armcontainerregistry.TaskPropertiesUpdateParameters{
				AgentConfiguration: &armcontainerregistry.AgentProperties{
					CPU: to.Int32Ptr(3),
				},
				Credentials: &armcontainerregistry.Credentials{
					CustomRegistries: map[string]*armcontainerregistry.CustomRegistryCredentials{
						"myregistry.azurecr.io": {
							Identity: to.StringPtr("<identity>"),
							Password: &armcontainerregistry.SecretObject{
								Type:  armcontainerregistry.SecretObjectType("Vaultsecret").ToPtr(),
								Value: to.StringPtr("<value>"),
							},
							UserName: &armcontainerregistry.SecretObject{
								Type:  armcontainerregistry.SecretObjectType("Opaque").ToPtr(),
								Value: to.StringPtr("<value>"),
							},
						},
					},
				},
				LogTemplate: to.StringPtr("<log-template>"),
				Status:      armcontainerregistry.TaskStatus("Enabled").ToPtr(),
				Step: &armcontainerregistry.DockerBuildStepUpdateParameters{
					Type:           armcontainerregistry.StepType("Docker").ToPtr(),
					DockerFilePath: to.StringPtr("<docker-file-path>"),
					ImageNames: []*string{
						to.StringPtr("azurerest:testtag1")},
				},
				Trigger: &armcontainerregistry.TriggerUpdateParameters{
					SourceTriggers: []*armcontainerregistry.SourceTriggerUpdateParameters{
						{
							Name: to.StringPtr("<name>"),
							SourceRepository: &armcontainerregistry.SourceUpdateParameters{
								SourceControlAuthProperties: &armcontainerregistry.AuthInfoUpdateParameters{
									Token:     to.StringPtr("<token>"),
									TokenType: armcontainerregistry.TokenType("PAT").ToPtr(),
								},
							},
							SourceTriggerEvents: []*armcontainerregistry.SourceTriggerEvent{
								armcontainerregistry.SourceTriggerEvent("commit").ToPtr()},
						}},
				},
			},
			Tags: map[string]*string{
				"testkey": to.StringPtr("value"),
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
	log.Printf("Response result: %#v\n", res.TasksClientUpdateResult)
}
```

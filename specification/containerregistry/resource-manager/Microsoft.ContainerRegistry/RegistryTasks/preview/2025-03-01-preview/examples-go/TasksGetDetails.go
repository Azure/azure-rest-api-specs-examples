package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/523ccabf440d8cf1c5b0ea18a8ad1ffedf4902ac/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/TasksGetDetails.json
func ExampleTasksClient_GetDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTasksClient().GetDetails(ctx, "myResourceGroup", "myRegistry", "myTask", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Task = armcontainerregistry.Task{
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"testkey": to.Ptr("value"),
	// 	},
	// 	Identity: &armcontainerregistry.IdentityProperties{
	// 		Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("fa153151-b9fd-46f4-9088-5e6600f2689v"),
	// 		TenantID: to.Ptr("f686d426-8d16-42db-81b7-abu4gm510ccd"),
	// 	},
	// 	Properties: &armcontainerregistry.TaskProperties{
	// 		AgentConfiguration: &armcontainerregistry.AgentProperties{
	// 			CPU: to.Ptr[int32](2),
	// 		},
	// 		Credentials: &armcontainerregistry.Credentials{
	// 			CustomRegistries: map[string]*armcontainerregistry.CustomRegistryCredentials{
	// 				"myregistry.azure-test.io": &armcontainerregistry.CustomRegistryCredentials{
	// 					Identity: to.Ptr("[system]"),
	// 					Password: &armcontainerregistry.SecretObject{
	// 						Type: to.Ptr(armcontainerregistry.SecretObjectTypeVaultsecret),
	// 						Value: to.Ptr("https://myacbvault.vault.azure.net/secrets/username"),
	// 					},
	// 					UserName: &armcontainerregistry.SecretObject{
	// 						Type: to.Ptr(armcontainerregistry.SecretObjectTypeOpaque),
	// 						Value: to.Ptr("username"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		IsSystemTask: to.Ptr(false),
	// 		Platform: &armcontainerregistry.PlatformProperties{
	// 			Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 			OS: to.Ptr(armcontainerregistry.OSLinux),
	// 		},
	// 		Status: to.Ptr(armcontainerregistry.TaskStatusEnabled),
	// 		Step: &armcontainerregistry.DockerBuildStep{
	// 			Type: to.Ptr(armcontainerregistry.StepTypeDocker),
	// 			ContextPath: to.Ptr("src"),
	// 			Arguments: []*armcontainerregistry.Argument{
	// 				{
	// 					Name: to.Ptr("mytestargument"),
	// 					IsSecret: to.Ptr(false),
	// 					Value: to.Ptr("mytestvalue"),
	// 				},
	// 				{
	// 					Name: to.Ptr("mysecrettestargument"),
	// 					IsSecret: to.Ptr(true),
	// 					Value: to.Ptr("mysecrettestvalue"),
	// 			}},
	// 			DockerFilePath: to.Ptr("src/DockerFile"),
	// 			ImageNames: []*string{
	// 				to.Ptr("azurerest:testtag")},
	// 				IsPushEnabled: to.Ptr(true),
	// 				NoCache: to.Ptr(false),
	// 			},
	// 			Trigger: &armcontainerregistry.TriggerProperties{
	// 				BaseImageTrigger: &armcontainerregistry.BaseImageTrigger{
	// 					Name: to.Ptr("myBaseImageTrigger"),
	// 					BaseImageTriggerType: to.Ptr(armcontainerregistry.BaseImageTriggerTypeRuntime),
	// 					Status: to.Ptr(armcontainerregistry.TriggerStatusEnabled),
	// 					UpdateTriggerEndpoint: to.Ptr("https://user:pass@mycicd.webhook.com?token=foo"),
	// 					UpdateTriggerPayloadType: to.Ptr(armcontainerregistry.UpdateTriggerPayloadTypeToken),
	// 				},
	// 				SourceTriggers: []*armcontainerregistry.SourceTrigger{
	// 					{
	// 						Name: to.Ptr("mySourceTrigger"),
	// 						SourceRepository: &armcontainerregistry.SourceProperties{
	// 							Branch: to.Ptr("master"),
	// 							RepositoryURL: to.Ptr("https://github.com/Azure/azure-rest-api-specs"),
	// 							SourceControlAuthProperties: &armcontainerregistry.AuthInfo{
	// 								Token: to.Ptr("xxxxx"),
	// 								TokenType: to.Ptr(armcontainerregistry.TokenTypePAT),
	// 							},
	// 							SourceControlType: to.Ptr(armcontainerregistry.SourceControlTypeGithub),
	// 						},
	// 						SourceTriggerEvents: []*armcontainerregistry.SourceTriggerEvent{
	// 							to.Ptr(armcontainerregistry.SourceTriggerEventCommit)},
	// 							Status: to.Ptr(armcontainerregistry.TriggerStatusEnabled),
	// 					}},
	// 					TimerTriggers: []*armcontainerregistry.TimerTrigger{
	// 						{
	// 							Name: to.Ptr("myTimerTrigger"),
	// 							Schedule: to.Ptr("30 9 * * 1-5"),
	// 							Status: to.Ptr(armcontainerregistry.TriggerStatusEnabled),
	// 					}},
	// 				},
	// 			},
	// 		}
}

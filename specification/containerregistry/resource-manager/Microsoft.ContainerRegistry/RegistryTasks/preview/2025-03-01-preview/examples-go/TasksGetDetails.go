package armcontainerregistrytasks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistrytasks"
)

// Generated from example definition: 2025-03-01-preview/TasksGetDetails.json
func ExampleTasksClient_GetDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistrytasks.NewClientFactory("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
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
	// res = armcontainerregistrytasks.TasksClientGetDetailsResponse{
	// 	Task: &armcontainerregistrytasks.Task{
	// 		Identity: &armcontainerregistrytasks.IdentityProperties{
	// 			Type: to.Ptr(armcontainerregistrytasks.ResourceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("fa153151-b9fd-46f4-9088-5e6600f2689v"),
	// 			TenantID: to.Ptr("f686d426-8d16-42db-81b7-abu4gm510ccd"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armcontainerregistrytasks.TaskProperties{
	// 			AgentConfiguration: &armcontainerregistrytasks.AgentProperties{
	// 				CPU: to.Ptr[int32](2),
	// 			},
	// 			Credentials: &armcontainerregistrytasks.Credentials{
	// 				CustomRegistries: map[string]*armcontainerregistrytasks.CustomRegistryCredentials{
	// 					"myregistry.azure-test.io": &armcontainerregistrytasks.CustomRegistryCredentials{
	// 						Identity: to.Ptr("[system]"),
	// 						Password: &armcontainerregistrytasks.SecretObject{
	// 							Type: to.Ptr(armcontainerregistrytasks.SecretObjectTypeVaultsecret),
	// 							Value: to.Ptr("https://myacbvault.vault.azure.net/secrets/username"),
	// 						},
	// 						UserName: &armcontainerregistrytasks.SecretObject{
	// 							Type: to.Ptr(armcontainerregistrytasks.SecretObjectTypeOpaque),
	// 							Value: to.Ptr("username"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			IsSystemTask: to.Ptr(false),
	// 			Platform: &armcontainerregistrytasks.PlatformProperties{
	// 				Architecture: to.Ptr(armcontainerregistrytasks.ArchitectureAmd64),
	// 				OS: to.Ptr(armcontainerregistrytasks.OSLinux),
	// 			},
	// 			Status: to.Ptr(armcontainerregistrytasks.TaskStatusEnabled),
	// 			Step: &armcontainerregistrytasks.DockerBuildStep{
	// 				Type: to.Ptr(armcontainerregistrytasks.StepTypeDocker),
	// 				Arguments: []*armcontainerregistrytasks.Argument{
	// 					{
	// 						Name: to.Ptr("mytestargument"),
	// 						IsSecret: to.Ptr(false),
	// 						Value: to.Ptr("mytestvalue"),
	// 					},
	// 					{
	// 						Name: to.Ptr("mysecrettestargument"),
	// 						IsSecret: to.Ptr(true),
	// 						Value: to.Ptr("mysecrettestvalue"),
	// 					},
	// 				},
	// 				ContextPath: to.Ptr("src"),
	// 				DockerFilePath: to.Ptr("src/DockerFile"),
	// 				ImageNames: []*string{
	// 					to.Ptr("azurerest:testtag"),
	// 				},
	// 				IsPushEnabled: to.Ptr(true),
	// 				NoCache: to.Ptr(false),
	// 			},
	// 			Trigger: &armcontainerregistrytasks.TriggerProperties{
	// 				BaseImageTrigger: &armcontainerregistrytasks.BaseImageTrigger{
	// 					Name: to.Ptr("myBaseImageTrigger"),
	// 					BaseImageTriggerType: to.Ptr(armcontainerregistrytasks.BaseImageTriggerTypeRuntime),
	// 					Status: to.Ptr(armcontainerregistrytasks.TriggerStatusEnabled),
	// 					UpdateTriggerEndpoint: to.Ptr("https://user:pass@mycicd.webhook.com?token=foo"),
	// 					UpdateTriggerPayloadType: to.Ptr(armcontainerregistrytasks.UpdateTriggerPayloadTypeToken),
	// 				},
	// 				SourceTriggers: []*armcontainerregistrytasks.SourceTrigger{
	// 					{
	// 						Name: to.Ptr("mySourceTrigger"),
	// 						SourceRepository: &armcontainerregistrytasks.SourceProperties{
	// 							Branch: to.Ptr("master"),
	// 							RepositoryURL: to.Ptr("https://github.com/Azure/azure-rest-api-specs"),
	// 							SourceControlAuthProperties: &armcontainerregistrytasks.AuthInfo{
	// 								Token: to.Ptr("xxxxx"),
	// 								TokenType: to.Ptr(armcontainerregistrytasks.TokenTypePAT),
	// 							},
	// 							SourceControlType: to.Ptr(armcontainerregistrytasks.SourceControlTypeGithub),
	// 						},
	// 						SourceTriggerEvents: []*armcontainerregistrytasks.SourceTriggerEvent{
	// 							to.Ptr(armcontainerregistrytasks.SourceTriggerEventCommit),
	// 						},
	// 						Status: to.Ptr(armcontainerregistrytasks.TriggerStatusEnabled),
	// 					},
	// 				},
	// 				TimerTriggers: []*armcontainerregistrytasks.TimerTrigger{
	// 					{
	// 						Name: to.Ptr("myTimerTrigger"),
	// 						Schedule: to.Ptr("30 9 * * 1-5"),
	// 						Status: to.Ptr(armcontainerregistrytasks.TriggerStatusEnabled),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"testkey": to.Ptr("value"),
	// 		},
	// 	},
	// }
}

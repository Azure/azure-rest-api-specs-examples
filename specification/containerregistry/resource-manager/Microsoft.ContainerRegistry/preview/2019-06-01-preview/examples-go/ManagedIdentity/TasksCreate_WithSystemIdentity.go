package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/ManagedIdentity/TasksCreate_WithSystemIdentity.json
func ExampleTasksClient_BeginCreate_tasksCreateWithUserIdentitiesWithSystemIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTasksClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "mytTask", armcontainerregistry.Task{
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
					Name:                 to.Ptr("myBaseImageTrigger"),
					BaseImageTriggerType: to.Ptr(armcontainerregistry.BaseImageTriggerTypeRuntime),
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Task = armcontainerregistry.Task{
	// 	Name: to.Ptr("myTask"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/tasks"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tasks/myTask"),
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
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T06:54:23.536Z"); return t}()),
	// 		IsSystemTask: to.Ptr(false),
	// 		Platform: &armcontainerregistry.PlatformProperties{
	// 			Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 			OS: to.Ptr(armcontainerregistry.OSLinux),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		Status: to.Ptr(armcontainerregistry.TaskStatusEnabled),
	// 		Step: &armcontainerregistry.DockerBuildStep{
	// 			Type: to.Ptr(armcontainerregistry.StepTypeDocker),
	// 			ContextPath: to.Ptr("src"),
	// 			Arguments: []*armcontainerregistry.Argument{
	// 				{
	// 					Name: to.Ptr("mytestargument"),
	// 					IsSecret: to.Ptr(false),
	// 					Value: to.Ptr("mytestvalue"),
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
	// 				},
	// 				SourceTriggers: []*armcontainerregistry.SourceTrigger{
	// 					{
	// 						Name: to.Ptr("mySourceTrigger"),
	// 						SourceRepository: &armcontainerregistry.SourceProperties{
	// 							Branch: to.Ptr("master"),
	// 							RepositoryURL: to.Ptr("https://github.com/Azure/azure-rest-api-specs"),
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

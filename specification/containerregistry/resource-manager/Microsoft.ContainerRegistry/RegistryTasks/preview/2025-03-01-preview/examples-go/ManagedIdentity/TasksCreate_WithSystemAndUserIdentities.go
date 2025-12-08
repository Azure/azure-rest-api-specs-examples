package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/523ccabf440d8cf1c5b0ea18a8ad1ffedf4902ac/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/ManagedIdentity/TasksCreate_WithSystemAndUserIdentities.json
func ExampleTasksClient_Create_tasksCreateWithSystemAndUserIdentities() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTasksClient().Create(ctx, "myResourceGroup", "myRegistry", "mytTask", armcontainerregistry.Task{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"testkey": to.Ptr("value"),
		},
		Identity: &armcontainerregistry.IdentityProperties{
			Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
				"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": {},
			},
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
					Name:                     to.Ptr("myBaseImageTrigger"),
					BaseImageTriggerType:     to.Ptr(armcontainerregistry.BaseImageTriggerTypeRuntime),
					UpdateTriggerEndpoint:    to.Ptr("https://user:pass@mycicd.webhook.com?token=foo"),
					UpdateTriggerPayloadType: to.Ptr(armcontainerregistry.UpdateTriggerPayloadTypeDefault),
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
	// 		Type: to.Ptr(armcontainerregistry.ResourceIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("fc99dc67-1ad9-45c8-9ebc-e438081c8e30"),
	// 		TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 		UserAssignedIdentities: map[string]*armcontainerregistry.UserIdentityProperties{
	// 			"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": &armcontainerregistry.UserIdentityProperties{
	// 				ClientID: to.Ptr("e35621a5-f615-4a20-940e-de8a84b15abc"),
	// 				PrincipalID: to.Ptr("e45e3m7c-176e-416a-b466-0c5ec8298f8a"),
	// 			},
	// 		},
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
	// 					UpdateTriggerPayloadType: to.Ptr(armcontainerregistry.UpdateTriggerPayloadTypeDefault),
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

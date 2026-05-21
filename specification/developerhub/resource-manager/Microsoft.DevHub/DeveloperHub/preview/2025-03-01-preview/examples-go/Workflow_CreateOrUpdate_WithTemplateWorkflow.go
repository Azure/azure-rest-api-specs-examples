package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: 2025-03-01-preview/Workflow_CreateOrUpdate_WithTemplateWorkflow.json
func ExampleWorkflowClient_CreateOrUpdate_createTemplateWorkflow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowClient().CreateOrUpdate(ctx, "resourceGroup1", "workflow1", armdevhub.Workflow{
		Location: to.Ptr("location1"),
		Properties: &armdevhub.WorkflowProperties{
			TemplateWorkflowProfile: &armdevhub.TemplateWorkflowProfile{
				DeploymentTemplate: &armdevhub.TemplateReference{
					Destination: to.Ptr("."),
					Parameters: map[string]*string{
						"key1": to.Ptr("value1"),
					},
					TemplateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DevHub/templates/test-template/versions/0.0.1"),
				},
				DockerfileTemplate: &armdevhub.TemplateReference{
					Destination: to.Ptr("."),
					Parameters: map[string]*string{
						"key1": to.Ptr("value1"),
					},
					TemplateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DevHub/templates/test-template/versions/0.0.1"),
				},
				GitHubProviderProfile: &armdevhub.GitHubProviderProfile{
					OidcCredentials: &armdevhub.OidcCredentials{
						AzureClientID: to.Ptr("test-client-id"),
						AzureTenantID: to.Ptr("test"),
					},
					Repository: &armdevhub.GitHubRepository{
						BranchName:      to.Ptr("main"),
						RepositoryName:  to.Ptr("test-repo"),
						RepositoryOwner: to.Ptr("test-owner"),
					},
				},
				ManifestTemplates: []*armdevhub.TemplateReference{
					{
						Destination: to.Ptr("."),
						Parameters: map[string]*string{
							"key1": to.Ptr("value1"),
						},
						TemplateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DevHub/templates/test-template/versions/0.0.1"),
					},
				},
				RepositoryProvider: to.Ptr(armdevhub.RepositoryProviderTypeGithub),
				WorkflowTemplate: &armdevhub.TemplateReference{
					Destination: to.Ptr("."),
					Parameters: map[string]*string{
						"key1": to.Ptr("value1"),
					},
					TemplateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DevHub/templates/test-template/versions/0.0.1"),
				},
			},
		},
		Tags: map[string]*string{
			"appname": to.Ptr("testApp"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevhub.WorkflowClientCreateOrUpdateResponse{
	// 	Workflow: armdevhub.Workflow{
	// 		Name: to.Ptr("workflow1"),
	// 		Type: to.Ptr("Micfosoft.DevHub/Workflow"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DevHub/workflow/workflow1"),
	// 		Location: to.Ptr("location1"),
	// 		Properties: &armdevhub.WorkflowProperties{
	// 			TemplateWorkflowProfile: &armdevhub.TemplateWorkflowProfile{
	// 				AuthStatus: to.Ptr(armdevhub.AuthorizationStatusAuthorized),
	// 				DeploymentTemplate: &armdevhub.TemplateReference{
	// 					Destination: to.Ptr("."),
	// 					Parameters: map[string]*string{
	// 						"key1": to.Ptr("value1"),
	// 					},
	// 					TemplateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DevHub/templates/test-template/versions/0.0.1"),
	// 				},
	// 				DockerfileTemplate: &armdevhub.TemplateReference{
	// 					Destination: to.Ptr("."),
	// 					Parameters: map[string]*string{
	// 						"key1": to.Ptr("value1"),
	// 					},
	// 					TemplateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DevHub/templates/test-template/versions/0.0.1"),
	// 				},
	// 				GitHubProviderProfile: &armdevhub.GitHubProviderProfile{
	// 					OidcCredentials: &armdevhub.OidcCredentials{
	// 						AzureClientID: to.Ptr("test-client-id"),
	// 						AzureTenantID: to.Ptr("test"),
	// 					},
	// 					Repository: &armdevhub.GitHubRepository{
	// 						BranchName: to.Ptr("main"),
	// 						RepositoryName: to.Ptr("test-repo"),
	// 						RepositoryOwner: to.Ptr("test-owner"),
	// 					},
	// 				},
	// 				LastWorkflowRun: &armdevhub.WorkflowRun{
	// 					LastRunAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-01T00:00:00Z"); return t}()),
	// 					Succeeded: to.Ptr(false),
	// 					WorkflowRunStatus: to.Ptr(armdevhub.WorkflowRunStatusQueued),
	// 					WorkflowRunURL: to.Ptr("testurl"),
	// 				},
	// 				ManifestTemplates: []*armdevhub.TemplateReference{
	// 					{
	// 						Destination: to.Ptr("."),
	// 						Parameters: map[string]*string{
	// 							"key1": to.Ptr("value1"),
	// 						},
	// 						TemplateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DevHub/templates/test-template/versions/0.0.1"),
	// 					},
	// 				},
	// 				PullRequest: &armdevhub.PullRequest{
	// 					PrStatus: to.Ptr(armdevhub.PullRequestStatusSubmitted),
	// 					PrURL: to.Ptr("testurl"),
	// 					PullNumber: to.Ptr[int32](1234),
	// 				},
	// 				RepositoryProvider: to.Ptr(armdevhub.RepositoryProviderTypeGithub),
	// 				WorkflowTemplate: &armdevhub.TemplateReference{
	// 					Destination: to.Ptr("."),
	// 					Parameters: map[string]*string{
	// 						"key1": to.Ptr("value1"),
	// 					},
	// 					TemplateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DevHub/templates/test-template/versions/0.0.1"),
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armdevhub.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 			CreatedBy: to.Ptr("foo@contoso.com"),
	// 			CreatedByType: to.Ptr(armdevhub.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 			LastModifiedBy: to.Ptr("foo@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armdevhub.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"appname": to.Ptr("testapp"),
	// 		},
	// 	},
	// }
}

package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/Workflow_UpdateTags.json
func ExampleWorkflowClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowClient().UpdateTags(ctx, "resourceGroup1", "workflow1", armdevhub.TagsObject{
		Tags: map[string]*string{
			"promote":     to.Ptr("false"),
			"resourceEnv": to.Ptr("testing"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workflow = armdevhub.Workflow{
	// 	Name: to.Ptr("workflow1"),
	// 	Type: to.Ptr("Micfosoft.DevHub/Workflow"),
	// 	ID: to.Ptr("/subscription/subscriptionId1/resourceGroups/resourceGroup1/providers/Microsoft.DevHub/workflow/workflow1"),
	// 	SystemData: &armdevhub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("foo@contoso.com"),
	// 		CreatedByType: to.Ptr(armdevhub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("foo@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armdevhub.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location1"),
	// 	Tags: map[string]*string{
	// 		"previousValue": to.Ptr("stillExists"),
	// 		"promote": to.Ptr("false"),
	// 		"resourceEnv": to.Ptr("testing"),
	// 	},
	// 	Properties: &armdevhub.WorkflowProperties{
	// 		GithubWorkflowProfile: &armdevhub.GitHubWorkflowProfile{
	// 			Acr: &armdevhub.ACR{
	// 				AcrRegistryName: to.Ptr("registry1"),
	// 				AcrRepositoryName: to.Ptr("repo1"),
	// 				AcrResourceGroup: to.Ptr("resourceGroup1"),
	// 				AcrSubscriptionID: to.Ptr("subscriptionId1"),
	// 			},
	// 			AksResourceID: to.Ptr("/subscriptions/subscriptionId1/resourcegroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1"),
	// 			AuthStatus: to.Ptr(armdevhub.AuthorizationStatusAuthorized),
	// 			BranchName: to.Ptr("branch1"),
	// 			DeploymentProperties: &armdevhub.DeploymentProperties{
	// 				ManifestType: to.Ptr(armdevhub.ManifestTypeKube),
	// 				Overrides: map[string]*string{
	// 					"key1": to.Ptr("value1"),
	// 				},
	// 			},
	// 			DockerBuildContext: to.Ptr("repo1/src/"),
	// 			Dockerfile: to.Ptr("repo1/images/Dockerfile"),
	// 			LastWorkflowRun: &armdevhub.WorkflowRun{
	// 				LastRunAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T12:34:56.000Z"); return t}()),
	// 				Succeeded: to.Ptr(true),
	// 				WorkflowRunStatus: to.Ptr(armdevhub.WorkflowRunStatusCompleted),
	// 				WorkflowRunURL: to.Ptr("https://github.com/User/repo1/actions/runs/1820640230"),
	// 			},
	// 			OidcCredentials: &armdevhub.GitHubWorkflowProfileOidcCredentials{
	// 				AzureClientID: to.Ptr("12345678-3456-7890-5678-012345678901"),
	// 				AzureTenantID: to.Ptr("66666666-3456-7890-5678-012345678901"),
	// 			},
	// 			PrStatus: to.Ptr(armdevhub.PullRequestStatusMerged),
	// 			PrURL: to.Ptr("https://github.com/owner1/repo1/pull/6567"),
	// 			PullNumber: to.Ptr[int32](6567),
	// 			RepositoryName: to.Ptr("repo1"),
	// 			RepositoryOwner: to.Ptr("owner1"),
	// 		},
	// 	},
	// }
}

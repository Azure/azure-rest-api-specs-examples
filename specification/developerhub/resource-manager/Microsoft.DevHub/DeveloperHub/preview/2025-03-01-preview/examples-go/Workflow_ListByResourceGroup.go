package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: 2025-03-01-preview/Workflow_ListByResourceGroup.json
func ExampleWorkflowClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowClient().NewListByResourceGroupPager("resourceGroup1", &armdevhub.WorkflowClientListByResourceGroupOptions{
		ManagedClusterResource: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1")})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armdevhub.WorkflowClientListByResourceGroupResponse{
		// 	WorkflowListResult: armdevhub.WorkflowListResult{
		// 		NextLink: to.Ptr("https://management.azure.com:443/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DevHub/workflows?api-version=2018-05-05-preview&timeRange=1&$skipToken=1"),
		// 		Value: []*armdevhub.Workflow{
		// 			{
		// 				Name: to.Ptr("workflow1"),
		// 				Type: to.Ptr("Micfosoft.DevHub/Workflow"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DevHub/workflow/workflow1"),
		// 				Location: to.Ptr("location1"),
		// 				Properties: &armdevhub.WorkflowProperties{
		// 					GithubWorkflowProfile: &armdevhub.GitHubWorkflowProfile{
		// 						Acr: &armdevhub.ACR{
		// 							AcrRegistryName: to.Ptr("registry1"),
		// 							AcrRepositoryName: to.Ptr("repo1"),
		// 							AcrResourceGroup: to.Ptr("resourceGroup1"),
		// 							AcrSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						},
		// 						AksResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1"),
		// 						AuthStatus: to.Ptr(armdevhub.AuthorizationStatusAuthorized),
		// 						BranchName: to.Ptr("branch1"),
		// 						DeploymentProperties: &armdevhub.Deployment{
		// 							ManifestType: to.Ptr(armdevhub.ManifestTypeKube),
		// 							Overrides: map[string]*string{
		// 								"key1": to.Ptr("value1"),
		// 							},
		// 						},
		// 						DockerBuildContext: to.Ptr("repo1/src/"),
		// 						Dockerfile: to.Ptr("repo1/images/Dockerfile"),
		// 						LastWorkflowRun: &armdevhub.WorkflowRun{
		// 							LastRunAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T12:34:56.000Z"); return t}()),
		// 							Succeeded: to.Ptr(true),
		// 							WorkflowRunStatus: to.Ptr(armdevhub.WorkflowRunStatusCompleted),
		// 							WorkflowRunURL: to.Ptr("https://github.com/User/repo1/actions/runs/1820640230"),
		// 						},
		// 						OidcCredentials: &armdevhub.GitHubWorkflowProfileOidcCredentials{
		// 							AzureClientID: to.Ptr("12345678-3456-7890-5678-012345678901"),
		// 							AzureTenantID: to.Ptr("66666666-3456-7890-5678-012345678901"),
		// 						},
		// 						PrStatus: to.Ptr(armdevhub.PullRequestStatusMerged),
		// 						PrURL: to.Ptr("https://github.com/User/repo1/pull/6567"),
		// 						PullNumber: to.Ptr[int32](6567),
		// 						RepositoryName: to.Ptr("repo1"),
		// 						RepositoryOwner: to.Ptr("owner1"),
		// 					},
		// 				},
		// 				SystemData: &armdevhub.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 					CreatedBy: to.Ptr("foo@contoso.com"),
		// 					CreatedByType: to.Ptr(armdevhub.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 					LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armdevhub.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 					"appname": to.Ptr("testapp"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

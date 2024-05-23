package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/ListGitHubRepos_example.json
func ExampleGitHubReposClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGitHubReposClient().NewListPager("myRg", "mySecurityConnectorName", "myGitHubOwner", nil)
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
		// page.GitHubRepositoryListResponse = armsecurity.GitHubRepositoryListResponse{
		// 	Value: []*armsecurity.GitHubRepository{
		// 		{
		// 			Name: to.Ptr("myGitHubRepo"),
		// 			Type: to.Ptr("Microsoft.Security/securityConnectors/devops/gitHubOwners/repos"),
		// 			ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/gitHubOwners/myGitHubOwner/repos/myGitHubRepo"),
		// 			Properties: &armsecurity.GitHubRepositoryProperties{
		// 				OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
		// 				ParentOwnerName: to.Ptr("myGitHubOwner"),
		// 				ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
		// 				RepoFullName: to.Ptr("myGitHubOwner/myGitHubRepo"),
		// 				RepoName: to.Ptr("myGitHubRepo"),
		// 				RepoURL: to.Ptr("https://github.com/myGitHubOwner/myGitHubRepo"),
		// 			},
		// 	}},
		// }
	}
}

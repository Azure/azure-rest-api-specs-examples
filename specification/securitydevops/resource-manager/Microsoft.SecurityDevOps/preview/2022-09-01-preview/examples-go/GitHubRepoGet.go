package armsecuritydevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoGet.json
func ExampleGitHubRepoClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitHubRepoClient().Get(ctx, "westusrg", "testconnector", "Azure", "39093389", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubRepo = armsecuritydevops.GitHubRepo{
	// 	Name: to.Ptr("azure-rest-api-specs"),
	// 	Type: to.Ptr("Microsoft.SecurityDevOps/owners/repos"),
	// 	ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/gitHubConnectors/testconnector/owners/Azure/repos/azure-rest-api-specs"),
	// 	Properties: &armsecuritydevops.GitHubRepoProperties{
	// 		AccountID: to.Ptr[int64](6844498),
	// 		ProvisioningState: to.Ptr(armsecuritydevops.ProvisioningStateSucceeded),
	// 		RepoURL: to.Ptr("https://github.com/Azure/azure-rest-api-specs"),
	// 	},
	// }
}

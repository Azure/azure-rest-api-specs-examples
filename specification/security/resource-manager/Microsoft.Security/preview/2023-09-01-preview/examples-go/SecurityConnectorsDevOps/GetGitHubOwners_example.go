package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/GetGitHubOwners_example.json
func ExampleGitHubOwnersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitHubOwnersClient().Get(ctx, "myRg", "mySecurityConnectorName", "myGitHubOwner", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubOwner = armsecurity.GitHubOwner{
	// 	Name: to.Ptr("myGitHubOwner"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors/devops/gitHubOwners"),
	// 	ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/gitHubOwners/myGitHubOwner"),
	// 	Properties: &armsecurity.GitHubOwnerProperties{
	// 		OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
	// 		OwnerURL: to.Ptr("https://github.com/myGitHubOwner"),
	// 		ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 	},
	// }
}

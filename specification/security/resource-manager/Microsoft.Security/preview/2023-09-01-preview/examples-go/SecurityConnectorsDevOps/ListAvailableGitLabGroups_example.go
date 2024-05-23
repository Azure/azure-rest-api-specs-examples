package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/ListAvailableGitLabGroups_example.json
func ExampleGitLabGroupsClient_ListAvailable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitLabGroupsClient().ListAvailable(ctx, "myRg", "mySecurityConnectorName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitLabGroupListResponse = armsecurity.GitLabGroupListResponse{
	// 	Value: []*armsecurity.GitLabGroup{
	// 		{
	// 			Name: to.Ptr("myGitLabGroup$mySubGroup"),
	// 			Type: to.Ptr("Microsoft.Security/securityConnectors/devops/gitLabGroups"),
	// 			ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/gitLabGroups/myGitLabGroup$mySubGroup"),
	// 			Properties: &armsecurity.GitLabGroupProperties{
	// 				FullyQualifiedName: to.Ptr("myGitLabGroup$mySubGroup"),
	// 				OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
	// 				ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 				URL: to.Ptr("https://gitlab.example.com/myGitLabGroup/mySubGroup"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("myGitLabGroup"),
	// 			Type: to.Ptr("Microsoft.Security/securityConnectors/devops/gitLabGroups"),
	// 			ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/gitLabGroups/myGitLabGroup"),
	// 			Properties: &armsecurity.GitLabGroupProperties{
	// 				FullyQualifiedName: to.Ptr("myGitLabGroup"),
	// 				OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboardedByOtherConnector),
	// 				ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 				URL: to.Ptr("https://gitlab.example.com/myGitLabGroup"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("anotherGitLabGroup"),
	// 			Type: to.Ptr("Microsoft.Security/securityConnectors/devops/gitLabGroups"),
	// 			ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/gitLabGroups/anotherGitLabGroup"),
	// 			Properties: &armsecurity.GitLabGroupProperties{
	// 				FullyQualifiedName: to.Ptr("anotherGitLabGroup"),
	// 				OnboardingState: to.Ptr(armsecurity.OnboardingStateNotOnboarded),
	// 				URL: to.Ptr("https://gitlab.example.com/anotherGitLabGroup"),
	// 			},
	// 	}},
	// }
}

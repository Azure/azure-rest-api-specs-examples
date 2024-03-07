package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/ListGitLabGroups_example.json
func ExampleGitLabGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGitLabGroupsClient().NewListPager("myRg", "mySecurityConnectorName", nil)
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
		// page.GitLabGroupListResponse = armsecurity.GitLabGroupListResponse{
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
		// 	}},
		// }
	}
}

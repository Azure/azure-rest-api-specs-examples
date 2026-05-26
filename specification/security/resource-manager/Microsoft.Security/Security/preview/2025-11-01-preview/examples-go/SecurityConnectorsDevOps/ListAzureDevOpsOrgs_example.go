package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2025-11-01-preview/SecurityConnectorsDevOps/ListAzureDevOpsOrgs_example.json
func ExampleAzureDevOpsOrgsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("0806e1cd-cfda-4ff8-b99c-2b0af42cffd3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureDevOpsOrgsClient().NewListPager("myRg", "mySecurityConnectorName", nil)
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
		// page = armsecurity.AzureDevOpsOrgsClientListResponse{
		// 	AzureDevOpsOrgListResponse: armsecurity.AzureDevOpsOrgListResponse{
		// 		Value: []*armsecurity.AzureDevOpsOrg{
		// 			{
		// 				Name: to.Ptr("myAzDevOpsOrg"),
		// 				Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs"),
		// 				ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg"),
		// 				Properties: &armsecurity.AzureDevOpsOrgProperties{
		// 					ActionableRemediation: &armsecurity.ActionableRemediation{
		// 						State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
		// 					},
		// 					OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
		// 					ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

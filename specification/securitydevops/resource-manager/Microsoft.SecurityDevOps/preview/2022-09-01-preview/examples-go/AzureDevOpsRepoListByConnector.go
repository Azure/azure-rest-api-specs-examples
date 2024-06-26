package armsecuritydevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsRepoListByConnector.json
func ExampleAzureDevOpsRepoClient_NewListByConnectorPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureDevOpsRepoClient().NewListByConnectorPager("westusrg", "testconnector", nil)
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
		// page.AzureDevOpsRepoListResponse = armsecuritydevops.AzureDevOpsRepoListResponse{
		// 	Value: []*armsecuritydevops.AzureDevOpsRepo{
		// 		{
		// 			Type: to.Ptr("microsoft.securitydevops/azureDevOpsConnectors/githubrepo"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/testconnector/orgs/myOrg1/projects/myProject/repos/myRepo"),
		// 			Properties: &armsecuritydevops.AzureDevOpsRepoProperties{
		// 				RepoURL: to.Ptr("https://dev.azure.com/myOrg/myProject/myRepo"),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("microsoft.securitydevops/azureDevOpsConnectors/githubrepo"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/testconnector/orgs/myOrg2/projects/myProject/repos/myRepo"),
		// 			Properties: &armsecuritydevops.AzureDevOpsRepoProperties{
		// 				RepoURL: to.Ptr("https://dev.azure.com/myOrg2/myProject/myRepo"),
		// 			},
		// 	}},
		// }
	}
}

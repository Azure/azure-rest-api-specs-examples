package armsecuritydevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsProjectList.json
func ExampleAzureDevOpsProjectClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureDevOpsProjectClient().NewListPager("westusrg", "testconnector", "myOrg", nil)
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
		// page.AzureDevOpsProjectListResponse = armsecuritydevops.AzureDevOpsProjectListResponse{
		// 	Value: []*armsecuritydevops.AzureDevOpsProject{
		// 		{
		// 			Type: to.Ptr("microsoft.securitydevops/azureDevOpsConnectors/orgs/projects"),
		// 			ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/testconnector/orgs/myOrg/projects/myProject"),
		// 			Properties: &armsecuritydevops.AzureDevOpsProjectProperties{
		// 				AutoDiscovery: to.Ptr(armsecuritydevops.AutoDiscoveryDisabled),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("microsoft.securitydevops/azureDevOpsConnectors/orgs/projects"),
		// 			ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/testconnector/orgs/myOrg/projects/myProjectNew"),
		// 			Properties: &armsecuritydevops.AzureDevOpsProjectProperties{
		// 				AutoDiscovery: to.Ptr(armsecuritydevops.AutoDiscoveryDisabled),
		// 			},
		// 	}},
		// }
	}
}

package armsecuritydevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorCreateOrUpdate.json
func ExampleGitHubConnectorClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGitHubConnectorClient().BeginCreateOrUpdate(ctx, "westusrg", "testconnector", armsecuritydevops.GitHubConnector{
		Location: to.Ptr("West US"),
		Properties: &armsecuritydevops.GitHubConnectorProperties{
			Code: to.Ptr("00000000000000000000"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubConnector = armsecuritydevops.GitHubConnector{
	// 	Name: to.Ptr("testconnector"),
	// 	Type: to.Ptr("microsoft.securitydevops/githubconnectors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/gitHubConnectors/testconnector"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsecuritydevops.GitHubConnectorProperties{
	// 		ProvisioningState: to.Ptr(armsecuritydevops.ProvisioningStateSucceeded),
	// 	},
	// }
}

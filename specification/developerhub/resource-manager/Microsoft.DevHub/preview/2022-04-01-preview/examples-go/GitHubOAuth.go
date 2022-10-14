package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-04-01-preview/examples/GitHubOAuth.json
func ExampleDeveloperHubServiceClient_GitHubOAuth() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevhub.NewDeveloperHubServiceClient("subscriptionId1", "<code>", "<state>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GitHubOAuth(ctx, "eastus2euap", &armdevhub.DeveloperHubServiceClientGitHubOAuthOptions{Parameters: &armdevhub.GitHubOAuthCallRequest{
		RedirectURL: to.Ptr("https://ms.portal.azure.com/aks"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

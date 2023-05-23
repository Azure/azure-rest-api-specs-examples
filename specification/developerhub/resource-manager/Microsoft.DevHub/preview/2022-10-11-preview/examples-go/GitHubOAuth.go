package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/GitHubOAuth.json
func ExampleDeveloperHubServiceClient_GitHubOAuth() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeveloperHubServiceClient().GitHubOAuth(ctx, "eastus2euap", &armdevhub.DeveloperHubServiceClientGitHubOAuthOptions{Parameters: &armdevhub.GitHubOAuthCallRequest{
		RedirectURL: to.Ptr("https://ms.portal.azure.com/aks"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubOAuthInfoResponse = armdevhub.GitHubOAuthInfoResponse{
	// 	AuthURL: to.Ptr("https://github.com/login/oauth/authorize?client_id=11111111&redirect_uri=https://management.azure.com/subscriptions/subscriptionId1/providers/Microsoft.DevHub/locations/eastus2euap/githuboauth&state=2345678-3456-7890-5678-012345678901&scope=repo,user:email,workflow"),
	// 	Token: to.Ptr(""),
	// }
}

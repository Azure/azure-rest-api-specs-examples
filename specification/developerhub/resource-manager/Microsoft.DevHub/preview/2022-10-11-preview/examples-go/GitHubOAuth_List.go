package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/GitHubOAuth_List.json
func ExampleDeveloperHubServiceClient_ListGitHubOAuth() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeveloperHubServiceClient().ListGitHubOAuth(ctx, "eastus2euap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubOAuthListResponse = armdevhub.GitHubOAuthListResponse{
	// 	Value: []*armdevhub.GitHubOAuthResponse{
	// 		{
	// 			Name: to.Ptr("default"),
	// 			Type: to.Ptr("Microsoft.DevHub/locations/githuboauth"),
	// 			ID: to.Ptr("/subscriptions/subscriptionId1/providers/Microsoft.DevHub/locations/eastus2euap/githuboauth/default"),
	// 			Properties: &armdevhub.GitHubOAuthProperties{
	// 				Username: to.Ptr("user"),
	// 			},
	// 	}},
	// }
}

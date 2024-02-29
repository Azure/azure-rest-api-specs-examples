package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_GetGitHubAccessToken.json
func ExampleFactoriesClient_GetGitHubAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFactoriesClient().GetGitHubAccessToken(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.GitHubAccessTokenRequest{
		GitHubAccessCode:         to.Ptr("some"),
		GitHubAccessTokenBaseURL: to.Ptr("some"),
		GitHubClientID:           to.Ptr("some"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubAccessTokenResponse = armdatafactory.GitHubAccessTokenResponse{
	// 	GitHubAccessToken: to.Ptr("myAccessTokenExample"),
	// }
}

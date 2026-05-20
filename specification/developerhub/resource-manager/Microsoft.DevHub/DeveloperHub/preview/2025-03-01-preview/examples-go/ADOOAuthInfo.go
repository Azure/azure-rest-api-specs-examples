package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: 2025-03-01-preview/ADOOAuthInfo.json
func ExampleDeveloperHubServiceClient_GetADOOAuthInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeveloperHubServiceClient().GetADOOAuthInfo(ctx, "eastus2euap", &armdevhub.DeveloperHubServiceClientGetADOOAuthInfoOptions{
		Parameters: &armdevhub.ADOOAuthCallRequest{
			RedirectURL: to.Ptr("https://ms.portal.azure.com/aks"),
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevhub.DeveloperHubServiceClientGetADOOAuthInfoResponse{
	// 	ADOOAuthInfoResponse: armdevhub.ADOOAuthInfoResponse{
	// 		AuthURL: to.Ptr("https://login.microsoftonline.com/common/oauth2/authorize?client_id=11111111&response_type=code&state=1234567-9123-4567-8912-234567890123&scope=offline_access%20openid%20vso.profile&redirect_uri=https://management.azure.com/subscriptions/subscriptionId1/providers/Microsoft.DevHub/locations/eastus2euap/adooauth"),
	// 		Token: to.Ptr(""),
	// 	},
	// }
}

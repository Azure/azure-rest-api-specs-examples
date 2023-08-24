package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetSsoToken.json
func ExampleServiceClient_GetSsoToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceClient().GetSsoToken(ctx, "rg1", "apimService1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceGetSsoTokenResult = armapimanagement.ServiceGetSsoTokenResult{
	// 	RedirectURI: to.Ptr("https://apimService1.portal.azure-api.net:443/signin-sso?token=1%26201705301929%26eIkr3%2fnfaLs1GVJ0OVbzkJjAcwPFkEZAPM8VUXvXPf7cJ6lWsB9oUwsk2zln9x0KLkn21txCPJWWheSPq7SNeA%3d%3d"),
	// }
}

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetAuthorizationServer.json
func ExampleAuthorizationServerClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationServerClient().Get(ctx, "rg1", "apimService1", "newauthServer2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationServerContract = armapimanagement.AuthorizationServerContract{
	// 	Name: to.Ptr("newauthServer2"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationServers"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationServers/newauthServer2"),
	// 	Properties: &armapimanagement.AuthorizationServerContractProperties{
	// 		Description: to.Ptr("test server"),
	// 		AuthorizationMethods: []*armapimanagement.AuthorizationMethod{
	// 			to.Ptr(armapimanagement.AuthorizationMethodGET)},
	// 			BearerTokenSendingMethods: []*armapimanagement.BearerTokenSendingMethod{
	// 				to.Ptr(armapimanagement.BearerTokenSendingMethodAuthorizationHeader)},
	// 				ClientAuthenticationMethod: []*armapimanagement.ClientAuthenticationMethod{
	// 					to.Ptr(armapimanagement.ClientAuthenticationMethodBasic)},
	// 					DefaultScope: to.Ptr("read write"),
	// 					ResourceOwnerPassword: to.Ptr("pwd"),
	// 					ResourceOwnerUsername: to.Ptr("un"),
	// 					SupportState: to.Ptr(true),
	// 					TokenEndpoint: to.Ptr("https://www.contoso.com/oauth2/token"),
	// 					AuthorizationEndpoint: to.Ptr("https://www.contoso.com/oauth2/auth"),
	// 					ClientID: to.Ptr("1"),
	// 					ClientRegistrationEndpoint: to.Ptr("https://www.contoso.com/apps"),
	// 					DisplayName: to.Ptr("test3"),
	// 					GrantTypes: []*armapimanagement.GrantType{
	// 						to.Ptr(armapimanagement.GrantTypeAuthorizationCode),
	// 						to.Ptr(armapimanagement.GrantTypeImplicit)},
	// 						UseInAPIDocumentation: to.Ptr(true),
	// 						UseInTestConsole: to.Ptr(false),
	// 					},
	// 				}
}

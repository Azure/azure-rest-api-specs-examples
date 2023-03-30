package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListAuthorizationServers.json
func ExampleAuthorizationServerClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthorizationServerClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.AuthorizationServerClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
	})
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
		// page.AuthorizationServerCollection = armapimanagement.AuthorizationServerCollection{
		// 	Value: []*armapimanagement.AuthorizationServerContract{
		// 		{
		// 			Name: to.Ptr("newauthServer"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationServers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationServers/newauthServer"),
		// 			Properties: &armapimanagement.AuthorizationServerContractProperties{
		// 				Description: to.Ptr("test server"),
		// 				AuthorizationMethods: []*armapimanagement.AuthorizationMethod{
		// 					to.Ptr(armapimanagement.AuthorizationMethodGET)},
		// 					BearerTokenSendingMethods: []*armapimanagement.BearerTokenSendingMethod{
		// 						to.Ptr(armapimanagement.BearerTokenSendingMethodAuthorizationHeader)},
		// 						DefaultScope: to.Ptr("read write"),
		// 						ResourceOwnerPassword: to.Ptr("pwd"),
		// 						ResourceOwnerUsername: to.Ptr("un"),
		// 						SupportState: to.Ptr(true),
		// 						TokenEndpoint: to.Ptr("https://www.contoso.com/oauth2/token"),
		// 						AuthorizationEndpoint: to.Ptr("https://www.contoso.com/oauth2/auth"),
		// 						ClientID: to.Ptr("1"),
		// 						ClientRegistrationEndpoint: to.Ptr("https://www.contoso.com/apps"),
		// 						DisplayName: to.Ptr("test2"),
		// 						GrantTypes: []*armapimanagement.GrantType{
		// 							to.Ptr(armapimanagement.GrantTypeAuthorizationCode),
		// 							to.Ptr(armapimanagement.GrantTypeImplicit)},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("newauthServer2"),
		// 						Type: to.Ptr("Microsoft.ApiManagement/service/authorizationServers"),
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationServers/newauthServer2"),
		// 						Properties: &armapimanagement.AuthorizationServerContractProperties{
		// 							Description: to.Ptr("test server"),
		// 							AuthorizationMethods: []*armapimanagement.AuthorizationMethod{
		// 								to.Ptr(armapimanagement.AuthorizationMethodGET)},
		// 								BearerTokenSendingMethods: []*armapimanagement.BearerTokenSendingMethod{
		// 									to.Ptr(armapimanagement.BearerTokenSendingMethodAuthorizationHeader)},
		// 									ClientAuthenticationMethod: []*armapimanagement.ClientAuthenticationMethod{
		// 										to.Ptr(armapimanagement.ClientAuthenticationMethodBasic)},
		// 										DefaultScope: to.Ptr("read write"),
		// 										ResourceOwnerPassword: to.Ptr("pwd"),
		// 										ResourceOwnerUsername: to.Ptr("un"),
		// 										SupportState: to.Ptr(true),
		// 										TokenEndpoint: to.Ptr("https://www.contoso.com/oauth2/token"),
		// 										AuthorizationEndpoint: to.Ptr("https://www.contoso.com/oauth2/auth"),
		// 										ClientID: to.Ptr("1"),
		// 										ClientRegistrationEndpoint: to.Ptr("https://www.contoso.com/apps"),
		// 										DisplayName: to.Ptr("test3"),
		// 										GrantTypes: []*armapimanagement.GrantType{
		// 											to.Ptr(armapimanagement.GrantTypeAuthorizationCode),
		// 											to.Ptr(armapimanagement.GrantTypeImplicit)},
		// 										},
		// 								}},
		// 							}
	}
}

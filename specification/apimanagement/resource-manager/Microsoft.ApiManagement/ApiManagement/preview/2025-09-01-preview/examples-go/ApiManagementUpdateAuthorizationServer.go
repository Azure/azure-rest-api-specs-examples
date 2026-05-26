package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementUpdateAuthorizationServer.json
func ExampleAuthorizationServerClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationServerClient().Update(ctx, "rg1", "apimService1", "newauthServer", "*", armapimanagement.AuthorizationServerUpdateContract{
		Properties: &armapimanagement.AuthorizationServerUpdateContractProperties{
			ClientID:              to.Ptr("update"),
			ClientSecret:          to.Ptr("updated"),
			UseInAPIDocumentation: to.Ptr(true),
			UseInTestConsole:      to.Ptr(false),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.AuthorizationServerClientUpdateResponse{
	// 	AuthorizationServerContract: armapimanagement.AuthorizationServerContract{
	// 		Name: to.Ptr("newauthServer"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/authorizationServers"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationServers/newauthServer"),
	// 		Properties: &armapimanagement.AuthorizationServerContractProperties{
	// 			Description: to.Ptr("test server"),
	// 			AuthorizationEndpoint: to.Ptr("https://www.contoso.com/oauth2/auth"),
	// 			AuthorizationMethods: []*armapimanagement.AuthorizationMethod{
	// 				to.Ptr(armapimanagement.AuthorizationMethodGET),
	// 			},
	// 			BearerTokenSendingMethods: []*armapimanagement.BearerTokenSendingMethod{
	// 				to.Ptr(armapimanagement.BearerTokenSendingMethodAuthorizationHeader),
	// 			},
	// 			ClientAuthenticationMethod: []*armapimanagement.ClientAuthenticationMethod{
	// 				to.Ptr(armapimanagement.ClientAuthenticationMethodBasic),
	// 			},
	// 			ClientID: to.Ptr("updated"),
	// 			ClientRegistrationEndpoint: to.Ptr("https://www.contoso.com/apps"),
	// 			DefaultScope: to.Ptr("read write"),
	// 			DisplayName: to.Ptr("test3"),
	// 			GrantTypes: []*armapimanagement.GrantType{
	// 				to.Ptr(armapimanagement.GrantTypeAuthorizationCode),
	// 				to.Ptr(armapimanagement.GrantTypeImplicit),
	// 			},
	// 			ResourceOwnerPassword: to.Ptr("pwd"),
	// 			ResourceOwnerUsername: to.Ptr("un"),
	// 			SupportState: to.Ptr(true),
	// 			TokenEndpoint: to.Ptr("https://www.contoso.com/oauth2/token"),
	// 			UseInAPIDocumentation: to.Ptr(true),
	// 			UseInTestConsole: to.Ptr(false),
	// 		},
	// 	},
	// }
}

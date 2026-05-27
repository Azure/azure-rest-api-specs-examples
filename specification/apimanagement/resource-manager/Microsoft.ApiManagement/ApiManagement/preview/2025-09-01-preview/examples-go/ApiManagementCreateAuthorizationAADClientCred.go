package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateAuthorizationAADClientCred.json
func ExampleAuthorizationClient_CreateOrUpdate_apiManagementCreateAuthorizationAadClientCred() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationClient().CreateOrUpdate(ctx, "rg1", "apimService1", "aadwithclientcred", "authz1", armapimanagement.AuthorizationContract{
		Properties: &armapimanagement.AuthorizationContractProperties{
			AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
			OAuth2GrantType:   to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
			Parameters: map[string]*string{
				"clientId":     to.Ptr("clientsecretid"),
				"clientSecret": to.Ptr("clientsecretvalue"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.AuthorizationClientCreateOrUpdateResponse{
	// 	AuthorizationContract: armapimanagement.AuthorizationContract{
	// 		Name: to.Ptr("authz1"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithclientcred/authorizations/authz1"),
	// 		Properties: &armapimanagement.AuthorizationContractProperties{
	// 			AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
	// 			OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeClientCredentials),
	// 			Parameters: map[string]*string{
	// 				"clientId": to.Ptr("clientsecretid"),
	// 			},
	// 			Status: to.Ptr("Connected"),
	// 		},
	// 	},
	// }
}

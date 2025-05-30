package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListAuthorizationsAuthCode.json
func ExampleAuthorizationClient_NewListByAuthorizationProviderPager_apiManagementListAuthorizationsAuthCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthorizationClient().NewListByAuthorizationProviderPager("rg1", "apimService1", "aadwithauthcode", &armapimanagement.AuthorizationClientListByAuthorizationProviderOptions{Filter: nil,
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
		// page.AuthorizationCollection = armapimanagement.AuthorizationCollection{
		// 	Value: []*armapimanagement.AuthorizationContract{
		// 		{
		// 			Name: to.Ptr("authz1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz1"),
		// 			Properties: &armapimanagement.AuthorizationContractProperties{
		// 				AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
		// 				OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
		// 				Status: to.Ptr("Connected"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("authz2"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz2"),
		// 			Properties: &armapimanagement.AuthorizationContractProperties{
		// 				AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
		// 				Error: &armapimanagement.AuthorizationError{
		// 					Code: to.Ptr("Unauthenticated"),
		// 					Message: to.Ptr("This connection is not authenticated."),
		// 				},
		// 				OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
		// 				Status: to.Ptr("Error"),
		// 			},
		// 	}},
		// }
	}
}

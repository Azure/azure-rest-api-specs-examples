package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListAuthorizationsAuthCode.json
func ExampleAuthorizationClient_NewListByAuthorizationProviderPager_apiManagementListAuthorizationsAuthCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthorizationClient().NewListByAuthorizationProviderPager("rg1", "apimService1", "aadwithauthcode", nil)
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
		// page = armapimanagement.AuthorizationClientListByAuthorizationProviderResponse{
		// 	AuthorizationCollection: armapimanagement.AuthorizationCollection{
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.AuthorizationContract{
		// 			{
		// 				Name: to.Ptr("authz1"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz1"),
		// 				Properties: &armapimanagement.AuthorizationContractProperties{
		// 					AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
		// 					OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
		// 					Status: to.Ptr("Connected"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("authz2"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz2"),
		// 				Properties: &armapimanagement.AuthorizationContractProperties{
		// 					AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
		// 					Error: &armapimanagement.AuthorizationError{
		// 						Code: to.Ptr("Unauthenticated"),
		// 						Message: to.Ptr("This connection is not authenticated."),
		// 					},
		// 					OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
		// 					Status: to.Ptr("Error"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

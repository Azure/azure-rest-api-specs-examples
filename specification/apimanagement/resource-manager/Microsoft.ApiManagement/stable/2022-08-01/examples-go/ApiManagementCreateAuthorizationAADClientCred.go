package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationAADClientCred.json
func ExampleAuthorizationClient_CreateOrUpdate_apiManagementCreateAuthorizationAadClientCred() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationClient().CreateOrUpdate(ctx, "rg1", "apimService1", "aadwithclientcred", "authz1", armapimanagement.AuthorizationContract{
		Properties: &armapimanagement.AuthorizationContractProperties{
			AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
			OAuth2GrantType:   to.Ptr(armapimanagement.OAuth2GrantTypeAuthorizationCode),
			Parameters: map[string]*string{
				"clientId":     to.Ptr("53790925-fdd3-4b80-bc7a-4c3aaf25801d"),
				"clientSecret": to.Ptr("FcJkQ3iPSaKAQRA7Ft8Q~fZ1X5vKmqzUAfJagcJ8"),
			},
		},
	}, &armapimanagement.AuthorizationClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationContract = armapimanagement.AuthorizationContract{
	// 	Name: to.Ptr("authz1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithclientcred/authorizations/authz1"),
	// 	Properties: &armapimanagement.AuthorizationContractProperties{
	// 		AuthorizationType: to.Ptr(armapimanagement.AuthorizationTypeOAuth2),
	// 		OAuth2GrantType: to.Ptr(armapimanagement.OAuth2GrantTypeClientCredentials),
	// 		Parameters: map[string]*string{
	// 			"clientId": to.Ptr("53790925-fdd3-4b80-bc7a-4c3aaf25801d"),
	// 		},
	// 		Status: to.Ptr("Connected"),
	// 	},
	// }
}

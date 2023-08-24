package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetIdentityProvider.json
func ExampleIdentityProviderClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIdentityProviderClient().Get(ctx, "rg1", "apimService1", armapimanagement.IdentityProviderTypeAADB2C, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IdentityProviderContract = armapimanagement.IdentityProviderContract{
	// 	Name: to.Ptr("AadB2C"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/identityProviders"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/identityProviders/AadB2C"),
	// 	Properties: &armapimanagement.IdentityProviderContractProperties{
	// 		Type: to.Ptr(armapimanagement.IdentityProviderTypeAADB2C),
	// 		AllowedTenants: []*string{
	// 			to.Ptr("contosoaadb2c.onmicrosoft.com"),
	// 			to.Ptr("contoso2aadb2c.onmicrosoft.com")},
	// 			Authority: to.Ptr("login.microsoftonline.com"),
	// 			SigninPolicyName: to.Ptr("B2C_1_policy-signin"),
	// 			SigninTenant: to.Ptr("contosoaadb2c.onmicrosoft.com"),
	// 			SignupPolicyName: to.Ptr("B2C_1_policy-signup"),
	// 			ClientID: to.Ptr("f02dafe2-b8b8-48ec-a38e-27e5c16c51e5"),
	// 		},
	// 	}
}

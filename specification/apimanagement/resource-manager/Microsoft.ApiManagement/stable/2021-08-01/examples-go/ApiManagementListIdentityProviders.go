package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListIdentityProviders.json
func ExampleIdentityProviderClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIdentityProviderClient().NewListByServicePager("rg1", "apimService1", nil)
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
		// page.IdentityProviderList = armapimanagement.IdentityProviderList{
		// 	Count: to.Ptr[int64](3),
		// 	Value: []*armapimanagement.IdentityProviderContract{
		// 		{
		// 			Name: to.Ptr("Google"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/identityProviders"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/identityProviders/Google"),
		// 			Properties: &armapimanagement.IdentityProviderContractProperties{
		// 				Type: to.Ptr(armapimanagement.IdentityProviderTypeGoogle),
		// 				ClientID: to.Ptr("googleId"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Aad"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/identityProviders"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/identityProviders/Aad"),
		// 			Properties: &armapimanagement.IdentityProviderContractProperties{
		// 				Type: to.Ptr(armapimanagement.IdentityProviderTypeAAD),
		// 				AllowedTenants: []*string{
		// 					to.Ptr("samiraad.onmicrosoft.com")},
		// 					ClientID: to.Ptr("aadapplicationid"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("AadB2C"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/identityProviders"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/identityProviders/AadB2C"),
		// 				Properties: &armapimanagement.IdentityProviderContractProperties{
		// 					Type: to.Ptr(armapimanagement.IdentityProviderTypeAADB2C),
		// 					AllowedTenants: []*string{
		// 						to.Ptr("samirtestbc.onmicrosoft.com")},
		// 						SigninPolicyName: to.Ptr("B2C_1_Signin_Default"),
		// 						SignupPolicyName: to.Ptr("B2C_1_Signup_Default"),
		// 						ClientID: to.Ptr("aadb2clientId"),
		// 					},
		// 			}},
		// 		}
	}
}

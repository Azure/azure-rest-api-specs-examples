package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListIdentityProviders.json
func ExampleIdentityProviderClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armapimanagement.IdentityProviderClientListByServiceResponse{
		// 	IdentityProviderList: armapimanagement.IdentityProviderList{
		// 		Count: to.Ptr[int64](3),
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.IdentityProviderContract{
		// 			{
		// 				Name: to.Ptr("Google"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/identityProviders"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/identityProviders/Google"),
		// 				Properties: &armapimanagement.IdentityProviderContractProperties{
		// 					Type: to.Ptr(armapimanagement.IdentityProviderTypeGoogle),
		// 					ClientID: to.Ptr("googleId"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Aad"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/identityProviders"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/identityProviders/Aad"),
		// 				Properties: &armapimanagement.IdentityProviderContractProperties{
		// 					Type: to.Ptr(armapimanagement.IdentityProviderTypeAAD),
		// 					AllowedTenants: []*string{
		// 						to.Ptr("samiraad.onmicrosoft.com"),
		// 					},
		// 					ClientID: to.Ptr("aadapplicationid"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("AadB2C"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/identityProviders"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/identityProviders/AadB2C"),
		// 				Properties: &armapimanagement.IdentityProviderContractProperties{
		// 					Type: to.Ptr(armapimanagement.IdentityProviderTypeAADB2C),
		// 					AllowedTenants: []*string{
		// 						to.Ptr("samirtestbc.onmicrosoft.com"),
		// 					},
		// 					ClientID: to.Ptr("aadb2clientId"),
		// 					SigninPolicyName: to.Ptr("B2C_1_Signin_Default"),
		// 					SignupPolicyName: to.Ptr("B2C_1_Signup_Default"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListAuthorizationAccessPolicies.json
func ExampleAuthorizationAccessPolicyClient_NewListByAuthorizationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthorizationAccessPolicyClient().NewListByAuthorizationPager("rg1", "apimService1", "aadwithauthcode", "authz1", &armapimanagement.AuthorizationAccessPolicyClientListByAuthorizationOptions{Filter: nil,
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
		// page.AuthorizationAccessPolicyCollection = armapimanagement.AuthorizationAccessPolicyCollection{
		// 	Value: []*armapimanagement.AuthorizationAccessPolicyContract{
		// 		{
		// 			Name: to.Ptr("fe0bed83-631f-4149-bd0b-0464b1bc7cab"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations/accessPolicies"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz1/accessPolicies"),
		// 			Properties: &armapimanagement.AuthorizationAccessPolicyContractProperties{
		// 				ObjectID: to.Ptr("fe0bed83-631f-4149-bd0b-0464b1bc7cab"),
		// 				TenantID: to.Ptr("13932a0d-5c63-4d37-901d-1df9c97722ff"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5585d6cd-2289-42e9-ab9b-3e2e23d74b4a"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations/accessPolicies"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz1/accessPolicies"),
		// 			Properties: &armapimanagement.AuthorizationAccessPolicyContractProperties{
		// 				ObjectID: to.Ptr("5585d6cd-2289-42e9-ab9b-3e2e23d74b4a"),
		// 				TenantID: to.Ptr("13932a0d-5c63-4d37-901d-1df9c97722ff"),
		// 			},
		// 	}},
		// }
	}
}

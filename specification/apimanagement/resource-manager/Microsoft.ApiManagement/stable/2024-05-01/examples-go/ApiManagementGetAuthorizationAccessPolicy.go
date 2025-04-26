package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetAuthorizationAccessPolicy.json
func ExampleAuthorizationAccessPolicyClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationAccessPolicyClient().Get(ctx, "rg1", "apimService1", "aadwithauthcode", "authz1", "fe0bed83-631f-4149-bd0b-0464b1bc7cab", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationAccessPolicyContract = armapimanagement.AuthorizationAccessPolicyContract{
	// 	Name: to.Ptr("fe0bed83-631f-4149-bd0b-0464b1bc7cab"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders/authorizations/accessPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode/authorizations/authz1/accessPolicies"),
	// 	Properties: &armapimanagement.AuthorizationAccessPolicyContractProperties{
	// 		AppIDs: []*string{
	// 			to.Ptr("d5f04bb0-ba78-4878-a43e-35a0b74fe315")},
	// 			ObjectID: to.Ptr("fe0bed83-631f-4149-bd0b-0464b1bc7cab"),
	// 			TenantID: to.Ptr("13932a0d-5c63-4d37-901d-1df9c97722ff"),
	// 		},
	// 	}
}

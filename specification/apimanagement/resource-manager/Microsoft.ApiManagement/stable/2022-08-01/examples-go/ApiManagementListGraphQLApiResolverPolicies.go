package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGraphQLApiResolverPolicies.json
func ExampleGraphQLAPIResolverPolicyClient_NewListByResolverPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGraphQLAPIResolverPolicyClient().NewListByResolverPager("rg1", "apimService1", "599e2953193c3c0bd0b3e2fa", "599e29ab193c3c0bd0b3e2fb", nil)
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
		// page.PolicyCollection = armapimanagement.PolicyCollection{
		// 	Count: to.Ptr[int64](1),
		// 	Value: []*armapimanagement.PolicyContract{
		// 		{
		// 			Name: to.Ptr("policy"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/apis/resolvers/policies"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/599e2953193c3c0bd0b3e2fa/resolvers/599e29ab193c3c0bd0b3e2fb/policies/policy"),
		// 			Properties: &armapimanagement.PolicyContractProperties{
		// 				Value: to.Ptr("<!--\r\n    IMPORTANT:\r\n    - Resolver elements can appear only within the <http-request>, <http-response>, <backend> section elements.\r\n    - Only the <forward-request> policy element can appear within the <backend> section element.\r\n    - To apply a policy to the query request (before it is forwarded to the backend service), place a corresponding policy element within the <http-request> section element.\r\n    - To apply a policy to the resolver response (before it is sent back to the document executor), place a corresponding policy element within the <http-response> section element.\r\n    - To add a policy position the cursor at the desired insertion point and click on the round button associated with the policy.\r\n    - To remove a policy, delete the corresponding policy statement from the policy document.\r\n  - Policies are applied in the order of their appearance, from the top down.\r\n-->\r\n<http-data-source>\r\n  <http-request>\r\n    <set-method>GET</set-method>\r\n<set-backend-service base-url=\"https://some.service.com\" />\r\n<set-url>/api/users</set-url>  </http-request>\r\n  <backend>\r\n    <forward-request />\r\n  </backend>\r\n  <http-response>\r\n    <set-body template=\"liquid\">\r\n { \"id\": \"{{body.id}}\" }  </http-response>\r\n</http-data-source>"),
		// 			},
		// 	}},
		// }
	}
}

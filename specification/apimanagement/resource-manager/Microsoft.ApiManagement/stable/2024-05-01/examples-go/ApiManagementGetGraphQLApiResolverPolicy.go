package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetGraphQLApiResolverPolicy.json
func ExampleGraphQLAPIResolverPolicyClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGraphQLAPIResolverPolicyClient().Get(ctx, "rg1", "apimService1", "5600b539c53f5b0062040001", "5600b53ac53f5b0062080006", armapimanagement.PolicyIDNamePolicy, &armapimanagement.GraphQLAPIResolverPolicyClientGetOptions{Format: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PolicyContract = armapimanagement.PolicyContract{
	// 	Name: to.Ptr("policy"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/resolvers/policies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/5600b539c53f5b0062040001/resolvers/5600b53ac53f5b0062080006/policies/policy"),
	// 	Properties: &armapimanagement.PolicyContractProperties{
	// 		Value: to.Ptr("<http-data-source>\r\n  <http-request>\r\n   <set-method>GET</set-method>\r\n<set-backend-service base-url=\"https://some.service.com\" />\r\n<set-url>/api/users</set-url>\r\n</http-request>\r\n</http-data-source>"),
	// 	},
	// }
}

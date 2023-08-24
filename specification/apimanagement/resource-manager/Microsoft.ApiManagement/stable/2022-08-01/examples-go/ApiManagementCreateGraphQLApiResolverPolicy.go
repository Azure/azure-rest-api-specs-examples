package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGraphQLApiResolverPolicy.json
func ExampleGraphQLAPIResolverPolicyClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGraphQLAPIResolverPolicyClient().CreateOrUpdate(ctx, "rg1", "apimService1", "5600b57e7e8880006a040001", "5600b57e7e8880006a080001", armapimanagement.PolicyIDNamePolicy, armapimanagement.PolicyContract{
		Properties: &armapimanagement.PolicyContractProperties{
			Format: to.Ptr(armapimanagement.PolicyContentFormatXML),
			Value:  to.Ptr("<http-data-source><http-request><set-method>GET</set-method><set-backend-service base-url=\"https://some.service.com\" /><set-url>/api/users</set-url></http-request></http-data-source>"),
		},
	}, &armapimanagement.GraphQLAPIResolverPolicyClientCreateOrUpdateOptions{IfMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PolicyContract = armapimanagement.PolicyContract{
	// 	Name: to.Ptr("policy"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/resolvers/policies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/5600b57e7e8880006a040001/resolvers/5600b57e7e8880006a080001/policies/policy"),
	// 	Properties: &armapimanagement.PolicyContractProperties{
	// 		Value: to.Ptr("<http-data-source>\r\n  <http-request>\r\n   <set-method>GET</set-method>\r\n<set-backend-service base-url=\"https://some.service.com\" />\r\n<set-url>/api/users</set-url>\r\n</http-request>\r\n</http-data-source>"),
	// 	},
	// }
}

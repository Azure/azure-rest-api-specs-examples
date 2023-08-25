package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGraphQLApiResolvers.json
func ExampleGraphQLAPIResolverClient_NewListByAPIPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGraphQLAPIResolverClient().NewListByAPIPager("rg1", "apimService1", "57d2ef278aa04f0888cba3f3", &armapimanagement.GraphQLAPIResolverClientListByAPIOptions{Filter: nil,
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
		// page.ResolverCollection = armapimanagement.ResolverCollection{
		// 	Count: to.Ptr[int64](3),
		// 	Value: []*armapimanagement.ResolverContract{
		// 		{
		// 			Name: to.Ptr("57d2ef278aa04f0ad01d6cdc"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/apis/resolvers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/57d2ef278aa04f0888cba3f3/resolvers/57d2ef278aa04f0ad01d6cdc"),
		// 			Properties: &armapimanagement.ResolverEntityBaseContract{
		// 				Path: to.Ptr("Query/users"),
		// 				Description: to.Ptr("A GraphQL Resolver example"),
		// 				DisplayName: to.Ptr("Query Users"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("57d2ef278aa04f0ad01d6cda"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/apis/resolvers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/57d2ef278aa04f0888cba3f3/resolvers/57d2ef278aa04f0ad01d6cda"),
		// 			Properties: &armapimanagement.ResolverEntityBaseContract{
		// 				Path: to.Ptr("Mutation/makeUser"),
		// 				Description: to.Ptr("A GraphQL Resolver example"),
		// 				DisplayName: to.Ptr("Mutation makeUser"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("57d2ef278aa04f0ad01d6cd9"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/apis/resolvers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/57d2ef278aa04f0888cba3f3/resolvers/57d2ef278aa04f0ad01d6cd9"),
		// 			Properties: &armapimanagement.ResolverEntityBaseContract{
		// 				Path: to.Ptr("User/id"),
		// 				Description: to.Ptr("A GraphQL Resolver example"),
		// 				DisplayName: to.Ptr("Query for User Id field"),
		// 			},
		// 	}},
		// }
	}
}

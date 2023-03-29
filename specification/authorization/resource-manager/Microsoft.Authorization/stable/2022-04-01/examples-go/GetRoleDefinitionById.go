package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetRoleDefinitionById.json
func ExampleRoleDefinitionsClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleDefinitionsClient().GetByID(ctx, "roleDefinitionId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleDefinition = armauthorization.RoleDefinition{
	// 	Name: to.Ptr("roleDefinitionId"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleDefinitions"),
	// 	ID: to.Ptr("/subscriptions/subID/providers/Microsoft.Authorization/roleDefinitions/roleDefinitionId"),
	// 	Properties: &armauthorization.RoleDefinitionProperties{
	// 		RoleType: to.Ptr("roletype"),
	// 		Description: to.Ptr("Role description"),
	// 		AssignableScopes: []*string{
	// 			to.Ptr("/subscriptions/subId")},
	// 			Permissions: []*armauthorization.Permission{
	// 				{
	// 					Actions: []*string{
	// 						to.Ptr("action")},
	// 						DataActions: []*string{
	// 							to.Ptr("dataAction")},
	// 							NotActions: []*string{
	// 							},
	// 							NotDataActions: []*string{
	// 							},
	// 					}},
	// 					RoleName: to.Ptr("Role name"),
	// 				},
	// 			}
}

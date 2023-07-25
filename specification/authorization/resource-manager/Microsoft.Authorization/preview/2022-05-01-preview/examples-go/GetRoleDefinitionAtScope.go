package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-05-01-preview/examples/GetRoleDefinitionAtScope.json
func ExampleRoleDefinitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleDefinitionsClient().NewListPager("scope", &armauthorization.RoleDefinitionsClientListOptions{Filter: nil})
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
		// page.RoleDefinitionListResult = armauthorization.RoleDefinitionListResult{
		// 	Value: []*armauthorization.RoleDefinition{
		// 		{
		// 			Name: to.Ptr("roleDefinitionId"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleDefinitions"),
		// 			ID: to.Ptr("/subscriptions/subID/providers/Microsoft.Authorization/roleDefinitions/roleDefinitionId"),
		// 			Properties: &armauthorization.RoleDefinitionProperties{
		// 				RoleType: to.Ptr("roletype"),
		// 				Description: to.Ptr("Role description"),
		// 				AssignableScopes: []*string{
		// 					to.Ptr("/subscriptions/subId")},
		// 					Permissions: []*armauthorization.Permission{
		// 						{
		// 							Actions: []*string{
		// 								to.Ptr("action")},
		// 								DataActions: []*string{
		// 									to.Ptr("dataAction")},
		// 									NotActions: []*string{
		// 									},
		// 									NotDataActions: []*string{
		// 									},
		// 							}},
		// 							RoleName: to.Ptr("Role name"),
		// 						},
		// 				}},
		// 			}
	}
}

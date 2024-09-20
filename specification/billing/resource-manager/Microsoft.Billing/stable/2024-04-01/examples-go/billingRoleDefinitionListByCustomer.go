package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleDefinitionListByCustomer.json
func ExampleRoleDefinitionClient_NewListByCustomerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleDefinitionClient().NewListByCustomerPager("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "11111111-1111-1111-1111-111111111111", nil)
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
		// page.RoleDefinitionListResult = armbilling.RoleDefinitionListResult{
		// 	Value: []*armbilling.RoleDefinition{
		// 		{
		// 			Name: to.Ptr("30000000-aaaa-bbbb-cccc-100000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/customers/billingRoleDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/customers/11111111-1111-1111-1111-111111111111/billingRoleDefinitions/30000000-aaaa-bbbb-cccc-100000000000"),
		// 			Properties: &armbilling.RoleDefinitionProperties{
		// 				Description: to.Ptr("The Owner role gives the user all permissions including access management rights to the customer."),
		// 				Permissions: []*armbilling.Permission{
		// 					{
		// 						Actions: []*string{
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000000"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000001"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000003"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000005"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000006"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000008"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000011"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000012"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000013"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000016"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000017"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000018")},
		// 							NotActions: []*string{
		// 							},
		// 					}},
		// 					RoleName: to.Ptr("Customer owner"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("30000000-aaaa-bbbb-cccc-100000000001"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/customers/billingRoleDefinitions"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/customers/11111111-1111-1111-1111-111111111111/billingRoleDefinitions/30000000-aaaa-bbbb-cccc-100000000001"),
		// 				Properties: &armbilling.RoleDefinitionProperties{
		// 					Description: to.Ptr("The Contributor role gives the user all permissions except access management rights to the customer."),
		// 					Permissions: []*armbilling.Permission{
		// 						{
		// 							Actions: []*string{
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000001"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000003"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000005"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000006"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000008"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000011"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000012"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000013"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000016"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000017"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000018")},
		// 								NotActions: []*string{
		// 								},
		// 						}},
		// 						RoleName: to.Ptr("Customer contributor"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("30000000-aaaa-bbbb-cccc-100000000002"),
		// 					Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/customers/billingRoleDefinitions"),
		// 					ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/customers/11111111-1111-1111-1111-111111111111/billingRoleDefinitions/30000000-aaaa-bbbb-cccc-100000000002"),
		// 					Properties: &armbilling.RoleDefinitionProperties{
		// 						Description: to.Ptr("The reader role gives the user read only access to the customer."),
		// 						Permissions: []*armbilling.Permission{
		// 							{
		// 								Actions: []*string{
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000001"),
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000017")},
		// 									NotActions: []*string{
		// 									},
		// 							}},
		// 							RoleName: to.Ptr("Customer reader"),
		// 						},
		// 				}},
		// 			}
	}
}

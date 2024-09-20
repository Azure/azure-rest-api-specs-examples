package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleDefinitionListByBillingProfile.json
func ExampleRoleDefinitionClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleDefinitionClient().NewListByBillingProfilePager("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", nil)
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
		// 			Name: to.Ptr("40000000-aaaa-bbbb-cccc-100000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/billingRoleDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/billingRoleDefinitions/40000000-aaaa-bbbb-cccc-100000000000"),
		// 			Properties: &armbilling.RoleDefinitionProperties{
		// 				Description: to.Ptr("The Owner role gives the user all permissions including access management rights to the billing profile."),
		// 				Permissions: []*armbilling.Permission{
		// 					{
		// 						Actions: []*string{
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000000"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000001"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000002"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000003"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000004"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000005"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000011"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000012"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000013")},
		// 					}},
		// 					RoleName: to.Ptr("Billing profile owner"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("40000000-aaaa-bbbb-cccc-100000000001"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/billingRoleDefinitions"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/billingRoleDefinitions/40000000-aaaa-bbbb-cccc-100000000001"),
		// 				Properties: &armbilling.RoleDefinitionProperties{
		// 					Description: to.Ptr("The Contributor role gives the user all permissions except access management rights to the billing profile."),
		// 					Permissions: []*armbilling.Permission{
		// 						{
		// 							Actions: []*string{
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000001"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000002"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000003"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000004"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000005"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000011"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000012"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000013")},
		// 						}},
		// 						RoleName: to.Ptr("Billing profile contributor"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("40000000-aaaa-bbbb-cccc-100000000002"),
		// 					Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/billingRoleDefinitions"),
		// 					ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/billingRoleDefinitions/40000000-aaaa-bbbb-cccc-100000000002"),
		// 					Properties: &armbilling.RoleDefinitionProperties{
		// 						Description: to.Ptr("The Reader role gives the user read only access to the billing profile."),
		// 						Permissions: []*armbilling.Permission{
		// 							{
		// 								Actions: []*string{
		// 									to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
		// 									to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
		// 									to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
		// 									to.Ptr("40000000-aaaa-bbbb-cccc-200000000012")},
		// 							}},
		// 							RoleName: to.Ptr("Billing profile reader"),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("40000000-aaaa-bbbb-cccc-100000000004"),
		// 						Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/billingRoleDefinitions"),
		// 						ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/billingRoleDefinitions/40000000-aaaa-bbbb-cccc-100000000004"),
		// 						Properties: &armbilling.RoleDefinitionProperties{
		// 							Description: to.Ptr("The Invoice Manager role gives the user the ability to view and manage invoices."),
		// 							Permissions: []*armbilling.Permission{
		// 								{
		// 									Actions: []*string{
		// 										to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
		// 										to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
		// 										to.Ptr("40000000-aaaa-bbbb-cccc-200000000011"),
		// 										to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
		// 										to.Ptr("40000000-aaaa-bbbb-cccc-200000000012"),
		// 										to.Ptr("40000000-aaaa-bbbb-cccc-200000000013")},
		// 								}},
		// 								RoleName: to.Ptr("Invoice manager"),
		// 							},
		// 					}},
		// 				}
	}
}

package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/billingRoleDefinitionListByEnrollmentAccount.json
func ExampleRoleDefinitionClient_NewListByEnrollmentAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleDefinitionClient().NewListByEnrollmentAccountPager("123456", "4568789", nil)
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
		// page = armbilling.RoleDefinitionClientListByEnrollmentAccountResponse{
		// 	RoleDefinitionListResult: armbilling.RoleDefinitionListResult{
		// 		Value: []*armbilling.RoleDefinition{
		// 			{
		// 				Name: to.Ptr("30000000-aaaa-bbbb-cccc-100000000000"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingRoleDefinitions"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/enrollmentAccounts/4568789/billingRoleDefinitions/30000000-aaaa-bbbb-cccc-100000000000"),
		// 				Properties: &armbilling.RoleDefinitionProperties{
		// 					Description: to.Ptr("The account owner role gives the user all permissions including access management rights to an accounts."),
		// 					Permissions: []*armbilling.Permission{
		// 						{
		// 							Actions: []*string{
		// 								to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/read"),
		// 								to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/write"),
		// 								to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/action"),
		// 								to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/delete"),
		// 								to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingSubscriptions/read"),
		// 								to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingSubscriptions/write"),
		// 								to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingSubscriptions/action"),
		// 							},
		// 							NotActions: []*string{
		// 							},
		// 						},
		// 					},
		// 					RoleName: to.Ptr("Account Owner"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

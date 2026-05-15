package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/billingRoleDefinitionGetByEnrollmentAccount.json
func ExampleRoleDefinitionClient_GetByEnrollmentAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleDefinitionClient().GetByEnrollmentAccount(ctx, "123456", "4568789", "50000000-aaaa-bbbb-cccc-100000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.RoleDefinitionClientGetByEnrollmentAccountResponse{
	// 	RoleDefinition: armbilling.RoleDefinition{
	// 		Name: to.Ptr("50000000-aaaa-bbbb-cccc-100000000000"),
	// 		Type: to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingRoleDefinitions"),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/123456/enrollmentAccounts/4568789/billingRoleDefinitions/50000000-aaaa-bbbb-cccc-100000000000"),
	// 		Properties: &armbilling.RoleDefinitionProperties{
	// 			Permissions: []*armbilling.Permission{
	// 				{
	// 					Actions: []*string{
	// 						to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/read"),
	// 						to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/write"),
	// 						to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/action"),
	// 						to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/delete"),
	// 						to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingSubscriptions/read"),
	// 						to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingSubscriptions/write"),
	// 						to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingSubscriptions/action"),
	// 					},
	// 				},
	// 			},
	// 			RoleName: to.Ptr("Account Owner"),
	// 		},
	// 	},
	// }
}

package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/billingRoleAssignmentGetByBillingAccount.json
func ExampleRoleAssignmentsClient_GetByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().GetByBillingAccount(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", "10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.RoleAssignmentsClientGetByBillingAccountResponse{
	// 	RoleAssignment: armbilling.RoleAssignment{
	// 		Name: to.Ptr("10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
	// 		Type: to.Ptr("Microsoft.Billing/billingAccounts/billingRoleAssignments"),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingRoleAssignments/10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
	// 		Properties: &armbilling.RoleAssignmentProperties{
	// 			CreatedByPrincipalID: to.Ptr(""),
	// 			CreatedByPrincipalTenantID: to.Ptr("076915e7-de10-4323-bb34-a58c904068bb"),
	// 			CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-04T22:39:34.2606750Z"); return t}()),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			PrincipalType: to.Ptr(armbilling.PrincipalTypeUser),
	// 			RoleDefinitionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingRoleDefinitions/00000000-0000-0000-0000-000000000000"),
	// 			Scope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30"),
	// 		},
	// 	},
	// }
}
